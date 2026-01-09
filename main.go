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
		fmt.Fprintf(s.Stderr(), "invalid compiler:  %s\n\r", s.User())
		s.Close() ; return
	}
	
	var pre Pre; if pre, e = readPre(s); e != nil {
		fmt.Fprintf(s.Stderr(), "invalid compiler:  %s\n\r", s.User())
		s.Close() ; return
	}

	if e = comp.Func(s, pre); e != nil {
		fmt.Fprintf(s.Stderr(), "invalid compiler:  %s\n\r", s.User())
		s.Close() ; return
	}
	s.Close()
}
