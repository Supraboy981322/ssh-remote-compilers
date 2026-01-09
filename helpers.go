package main

import (
	"io"
	"fmt"
	"bufio"
	"strings"
	"github.com/gliderlabs/ssh"
)

func readPre(s ssh.Session) (Pre, error) {
	pReader := bufio.NewReader(s)
	pre := Pre{}
	for {
		l, e := pReader.ReadString('\n')
		l = strings.TrimSpace(l)
		if e != nil {
			if e == io.EOF { break }
			return pre, fmt.Errorf("err reading p:  %v\n\r", e)
		} ; if l == "~~~" { break }

		p := strings.Split(l, ":")
		if len(p) < 2 { return pre, fmt.Errorf("invalid p:  %s", l) }

		for i, h := range p { p[i] = strings.TrimSpace(h)	}
		pre[p[0]] = p[1]
	}

	return pre, nil
}
