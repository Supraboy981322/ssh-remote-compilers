package main

import (
	"fmt"
//	"os/exec"
	"github.com/gliderlabs/ssh"
)

func c_compiler(s ssh.Session, pre Pre) error {
	for k, v := range pre {
		fmt.Fprintf(s.Stderr(), "k{%s} v{%s}\n", k, v)
	}
	return nil
}
