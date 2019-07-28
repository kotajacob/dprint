package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"git.sr.ht/~kota/xdg/basedir"
	"git.sr.ht/~sircmpwn/getopt"
)

var (
	Version string
	Config  string
)

func usage() {
	log.Fatal("Usage: dprint [-v] [-d path] [-i key:val] [-o key]")
}

func setConfig(d string) string {
	if d == "" {
		d = filepath.Join(basedir.ConfigHome, Config)
	}
	return d
}

func main() {
	var dir, in, out string
	opts, optind, err := getopt.Getopts(os.Args, "vd:i:o:")
	if err != nil {
		log.Print(err)
		usage()
		return
	}
	for _, opt := range opts {
		switch opt.Option {
		case 'v':
			fmt.Println("dprint " + Version)
			return
		case 'd':
			dir = opt.Value
		case 'i':
			in = opt.Value
		case 'o':
			out = opt.Value
		}
	}
	args := os.Args[optind:]
	if len(args) > 0 {
		usage()
		return
	}
	dir = setConfig(dir)
	println(dir)
	println(in)
	println(out)
}
