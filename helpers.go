package main

import (
	"os"
	"io"
	"bufio"
	"math/big"
	"crypto/rand"
	"archive/tar"
	"path/filepath"
	"github.com/gliderlabs/ssh"
)

func readReq(s ssh.Session) (ReqDat, error) {
	pReader := bufio.NewReader(s)

	d, e := os.MkdirTemp("/tmp", "ssh_rem_comp--*")
	if e != nil { return ReqDat{}, e }
	if e = wr_tarball(pReader, d); e != nil { return ReqDat{}, e }
	dat := ReqDat{
		Args: s.Command(),
		Dir: d,
	}

	return dat, nil
}

func wr_tarball(r io.Reader, d string) error {
	tr := tar.NewReader(r)
	for {
		h, e := tr.Next()
		if e == io.EOF { break }
		if e != nil { return e }
		
		e = ex_tarFi(tr, h, d)
		if e != nil { return e }
	}

	return nil
}

func ex_tarFi(tr *tar.Reader, h *tar.Header, d string) error {
	fiP := filepath.Join(d, h.Name)
	if h.Typeflag == tar.TypeDir { 
		return os.MkdirAll(fiP, os.FileMode(h.Mode))
	}

	if e := os.MkdirAll(filepath.Dir(fiP), 0755); e != nil { return e }

	fi, e := os.OpenFile(fiP, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(h.Mode))
	if e != nil { return e }
	defer fi.Close()

	if _, e = io.Copy(fi, tr); e != nil { return e }

	if e := fi.Sync(); e != nil { return e }

	return nil
}

func ran() string {
	charSet := []string{
		"a", "b",	"c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z", "A", "B",
		"C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "0", "9", "8", "7",
		"6", "5", "4", "3", "2", "1",
	}

	var l int64 = 16
	
	//actually generate
	var res string
	var i int64
	for i = 0; i < l; i++ {
		//convert to big.Int (for crypto/rand) 
		bigInt := big.NewInt(int64(len(charSet)))

		//generate random integer
		in, err := rand.Int(rand.Reader, bigInt)
		if err != nil { return err.Error() }

		//convert to regular integer
		ranDig := int(in.Int64())

		//add char of random index
		//  to result
		res += charSet[ranDig]
	}

	//finally,
	//  return the result
	return res
}
