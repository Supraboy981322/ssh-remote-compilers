package main

type (
	ReqDat struct {
		Args []string
		Dir string
	}

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
	CONF = Conf{
		Server: ServerConf{
			Port: 7845,
		},
		Comp: map[string]CompConf{
			"c": CompConf{
				Cmd: "gcc",
				Args: []string{"-o"},
			},
			"go": CompConf{
				Cmd: "go",
				Args: []string{"build", "-o"},
			},
			"zig": CompConf{
				Cmd: "zig",
				Args: []string{},
			},
		},
	}
)
