package internal

import (
	"context"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	GracefulShutdown time.Duration `name:"graceful-shutdown" env:"GRACEFUL_SHUTDOWN" description:"Interval before server shutdown" default:"15s"`
	Bind             string        `name:"bind" env:"BIND" description:"Address to where bind HTTP server" default:"127.0.0.1:8686"`
	TLS              bool          `name:"tls" env:"TLS" description:"Enable HTTPS serving with TLS"`
	CertFile         string        `name:"cert-file" env:"CERT_FILE" description:"Path to certificate for TLS" default:"server.crt"`
	KeyFile          string        `name:"key-file" env:"KEY_FILE" description:"Path to private key for TLS" default:"server.key"`
}

func (qs *HttpServer) Serve(globalCtx context.Context, handler http.Handler) error {

	server := http.Server{
		Addr:    qs.Bind,
		Handler: handler,
	}

	go func() {
		<-globalCtx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), qs.GracefulShutdown)
		defer cancel()
		server.Shutdown(ctx)
	}()
	log.Println("REST server is on", qs.Bind)
	if qs.TLS {
		return server.ListenAndServeTLS(qs.CertFile, qs.KeyFile)
	}
	return server.ListenAndServe()
}