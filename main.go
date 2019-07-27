package main

import (
	"git.sr.ht/~sircmpwn/getopt"
	"os"
)

func main() {
	var dir, in, out string
	opts, _, err := getopt.Getopts(os.Args, "d:i:o:")
	if err != nil {
		panic(err)
	}
	for _, opt := range opts {
		switch opt.Option {
		case 'd':
			dir = opt.Value
		case 'i':
			in = opt.Value
		case 'o':
			out = opt.Value
		}
	}
	println(dir)
	println(in)
	println(out)
}
