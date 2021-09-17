package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func openFile(file string) error{
	path, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	return exec.Command("cmd", "/c", fmt.Sprintf("start %s", path)).Start()
}