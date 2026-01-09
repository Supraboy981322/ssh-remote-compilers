package main

import "github.com/gliderlabs/ssh"

type (
	Pre map[string]string
	ReqDat struct {
		Pre Pre
		Dir string
	}

	ServerConf struct {
		Port int
	}

	CompConf struct {
		Cmd string
		Args []string
		Func func(s ssh.Session, dat ReqDat) error
	}

	Conf struct {
		Server ServerConf
		Comp map[string]CompConf
	}
)
