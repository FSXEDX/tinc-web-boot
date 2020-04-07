package tincd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"tinc-web-boot/network"
	"tinc-web-boot/utils"
)

type netImpl struct {
	tincBin string
	ctx     context.Context

	done  chan struct{}
	stop  func()
	lock  sync.Mutex
	peers peersManager

	definition *network.Network
}

func (impl *netImpl) Start() {
	impl.lock.Lock()
	defer impl.lock.Unlock()
	impl.unsafeStop()

	ctx, cancel := context.WithCancel(impl.ctx)
	done := make(chan struct{})
	impl.stop = cancel
	impl.done = done
	impl.peers = peersManager{
		network: impl.definition,
	}
	go func() {
		defer cancel()
		defer close(done)
		err := impl.run(ctx)
		if err != nil {
			log.Println("failed run network", impl.definition.Name(), ":", err)
		}
	}()
}

func (impl *netImpl) Stop() {
	impl.lock.Lock()
	defer impl.lock.Unlock()
	impl.unsafeStop()
}

func (impl *netImpl) Peers() []*Peer {
	return impl.peers.List()
}

func (impl *netImpl) Definition() *network.Network {
	return impl.definition
}

func (impl *netImpl) IsRunning() bool {
	ch := impl.done
	if ch == nil {
		return false
	}
	select {
	case <-ch:
		return false
	default:
		return true
	}
}

func (impl *netImpl) unsafeStop() {
	v := impl.stop
	if v == nil {
		return
	}
	v()
	<-impl.done
	impl.stop = nil
}

func (impl *netImpl) run(global context.Context) error {
	if err := impl.definition.Configure(global, impl.tincBin); err != nil {
		return fmt.Errorf("configure: %w", err)
	}

	ctx, abort := context.WithCancel(global)
	defer abort()

	cmd := exec.CommandContext(ctx, impl.tincBin, "-D", "-d", "-d", "-d",
		"--pidfile", impl.definition.Pidfile(),
		"--logfile", impl.definition.Logfile(),
		"-c", ".")
	cmd.Dir = impl.definition.Root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utils.SetCmdAttrs(cmd)

	peers := make(chan peerReq)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer abort()
		err := cmd.Run()
		if err != nil {
			log.Println(impl.definition.Name(), "failed to run tinc:", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer abort()
		runAPI(ctx, peers, impl.definition)
		log.Println(impl.definition.Name(), "api stopped")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer abort()
		impl.peers.Run(ctx, peers)
	}()

	wg.Wait()
	return ctx.Err()
}
