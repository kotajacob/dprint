package main

import (
	"git.sr.ht/~sircmpwn/getopt"
	"os"
)

func main() {
	var dir, in, out string
	opts, optind, err := getopt.Getopts(os.Args, "d:i:o:")
	if err != nil {
		panic(err)
	}
	for _, opt := range opts {
		switch opt.Option {
		case 'd':
			println("Option -d specified: " + opt.Value)
			dir = opt.Value
		case 'i':
			println("Option -i specified: " + opt.Value)
			in = opt.Value
		case 'o':
			println("Option -o specified: " + opt.Value)
			out = opt.Value
		}
	}
	println(dir)
	println(in)
	println(out)
}
