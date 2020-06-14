package main

import (
	wasmhttp "github.com/nlepage/go-wasm-http-server"
	"github.com/spf13/cobra"
)

func init() {
	httpCmd.Run = func(cmd *cobra.Command, args []string) {
		wasmhttp.Serve(httpHandler)

		select {}
	}
}
