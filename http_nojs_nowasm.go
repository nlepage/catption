// +build !js,!wasm

package main

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	httpCmd.RunE = func(cmd *cobra.Command, args []string) error {
		var l, err = net.Listen("tcp", ":")
		if err != nil {
			return err
		}

		var addr = fmt.Sprintf("http://%s", l.Addr())
		fmt.Printf("Listening at %s\n", addr)
		exec.Command(openCmd, addr).Start()

		return http.Serve(l, httpHandler)
	}
}
