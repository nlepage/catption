// +build linux freebsd

package main

import "os/exec"

func openFile(file string) error{
	return exec.Command("xdg-open", out).Start()
}
