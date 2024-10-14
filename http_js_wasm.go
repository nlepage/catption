package main

import (
	wasmhttp "github.com/nlepage/go-wasm-http-server/v2"
	"github.com/spf13/cobra"
)

func init() {
	httpCmd.Run = func(cmd *cobra.Command, args []string) {
		if _, err := wasmhttp.Serve(httpHandler); err != nil {
			panic(err)
		}

		select {}
	}
}
