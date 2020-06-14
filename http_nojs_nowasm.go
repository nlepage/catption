// +build !js,!wasm

package main

import (
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	httpCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return http.ListenAndServe(":8888", httpHandler)
	}
}
