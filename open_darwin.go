package main

import "os/exec"

func openFile(file string) error{
	return exec.Command("open", out).Start()
}