package main

import (
	"os"
	"net/http"
)

type (
	ServerConf struct {
		Port int
	}
	CompConf struct {
		Cmd string
		Args []string
	}
	Conf struct {
		Server ServerConf
		Comp map[string]CompConf
	}
)

var (
	conf = Conf{
		Server: ServerConf{
			Port: 7845,
		},
		Comp: map[string]CompConf{
			"c": CompConf{
				Cmd: "gcc",
				Args: []string{},
			},
		},
	} 
)

func main() {
}
