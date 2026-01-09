package main

import "github.com/gliderlabs/ssh"

type (
	Pre map[string]string

	ServerConf struct {
		Port int
	}

	CompConf struct {
		Cmd string
		Args []string
		Func func(s ssh.Session, pre Pre) error
	}

	Conf struct {
		Server ServerConf
		Comp map[string]CompConf
	}
)
