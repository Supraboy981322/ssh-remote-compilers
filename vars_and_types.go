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

var (
	CONF = Conf{
		Server: ServerConf{
			Port: 7845,
		},
		Comp: map[string]CompConf{
			"c": CompConf{
				Cmd: "gcc",
				Args: []string{},
				Func: c_compiler,
			},
			"go": CompConf{
				Cmd: "go",
				Args: []string{},
				Func: go_compiler,
			},
			"zig": CompConf{
				Cmd: "zig",
				Args: []string{},
				Func: go_compiler,
			},
		},
	}
)
