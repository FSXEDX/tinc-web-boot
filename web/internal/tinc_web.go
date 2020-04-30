// Code generated by jsonrpc2. DO NOT EDIT.
//go:generate jsonrpc2-gen -f ../../jsonrpc2.yaml -I TincWeb -I TincWebUI -I TincWebMajordomo
package internal

import (
	"encoding/json"
	jsonrpc2 "github.com/reddec/jsonrpc2"
	"time"
	network "tinc-web-boot/network"
	shared "tinc-web-boot/web/shared"
)

func RegisterTincWeb(router *jsonrpc2.Router, wrap shared.TincWeb) []string {
	router.RegisterFunc("TincWeb.Networks", func(params json.RawMessage, positional bool) (interface{}, error) {
		return wrap.Networks()
	})

	router.RegisterFunc("TincWeb.Network", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"name"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Network(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Create", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"name"`
			Arg1 string `json:"subnet"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Create(args.Arg0, args.Arg1)
	})

	router.RegisterFunc("TincWeb.Remove", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Remove(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Start", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Start(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Stop", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Stop(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Peers", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Peers(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Peer", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
			Arg1 string `json:"name"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Peer(args.Arg0, args.Arg1)
	})

	router.RegisterFunc("TincWeb.Import", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 shared.Sharing `json:"sharing"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Import(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Share", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Share(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Node", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string `json:"network"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Node(args.Arg0)
	})

	router.RegisterFunc("TincWeb.Upgrade", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string          `json:"network"`
			Arg1 network.Upgrade `json:"update"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Upgrade(args.Arg0, args.Arg1)
	})

	router.RegisterFunc("TincWeb.Majordomo", func(params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 string        `json:"network"`
			Arg1 time.Duration `json:"lifetime"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		return wrap.Majordomo(args.Arg0, args.Arg1)
	})

	return []string{"TincWeb.Networks", "TincWeb.Network", "TincWeb.Create", "TincWeb.Remove", "TincWeb.Start", "TincWeb.Stop", "TincWeb.Peers", "TincWeb.Peer", "TincWeb.Import", "TincWeb.Share", "TincWeb.Node", "TincWeb.Upgrade", "TincWeb.Majordomo"}
}
