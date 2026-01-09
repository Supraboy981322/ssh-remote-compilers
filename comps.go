package main

import (
	"io"
	"os"
	"fmt"
	"os/exec"
	"path/filepath"
	"github.com/gliderlabs/ssh"
)

func c_compiler(s ssh.Session, dat ReqDat) error {
	defer os.RemoveAll(dat.Dir)
	fmt.Fprintf(s.Stderr(), "TODO: C\n", dat.Dir) 
	return nil
}

func go_compiler(s ssh.Session, dat ReqDat) error {
	defer os.RemoveAll(dat.Dir)
	var e error

	binP := filepath.Join(dat.Dir, ran())
	fmt.Println(binP)
	cmd := exec.Command("go", "build", "-o", binP, ".")
	cmd.Stderr, cmd.Stdout, cmd.Dir= s.Stderr(), s.Stderr(), dat.Dir
	if e = cmd.Run(); e != nil { return e }

	var bin *os.File
	if bin, e = os.Open(binP); e != nil { return e }
	if _, e = io.Copy(s, bin); e != nil { return e }

	return nil
}

func zig_compiler(s ssh.Session, dat ReqDat) error {
	defer os.RemoveAll(dat.Dir)
	fmt.Fprintf(s.Stderr(), "TODO: Zig\n", dat.Dir) 
	return nil
}
