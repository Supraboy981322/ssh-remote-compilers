package main

import (
	"fmt"
	"strconv"
	"github.com/gliderlabs/ssh"
	"github.com/charmbracelet/log"
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
				Func: c_compiler,
			},
			"go": CompConf{
				Cmd: "go",
				Args: []string{},
				Func: go_compiler,
			},
		},
	} 
)

func init() {
	log.Info("starting server...")
}

func main() {
	ssh.Handle(sshServer)

	log.Infof("Listening on port %d", conf.Server.Port)

	port := ":"+strconv.Itoa(conf.Server.Port)
	log.Fatal(ssh.ListenAndServe(port, nil))
}

func sshServer(s ssh.Session) {
	var e error

	comp := conf.Comp[s.User()]
	if comp.Func == nil {
		log.Errorf("invalid compiler (user) attempt: %s", s.User()) 
		fmt.Fprintf(s.Stderr(), "unknown compiler:  %s\n\r", s.User())
		s.Close() ; return
	}
	
	var dat ReqDat; if dat, e = readReq(s); e != nil {
		log.Errorf("failed to read data:  %v\n\r", e)
		fmt.Fprintf(s.Stderr(), "failed to read data:  %v\n\r", e)
		s.Close() ; return
	}

	if e = comp.Func(s, dat); e != nil {
		log.Errorf("failed to compile:  %v\n\r", e)
		fmt.Fprintf(s.Stderr(), "failed to compile:  %v\n\r", e)
		s.Close() ; return
	}

	s.Close()
}
