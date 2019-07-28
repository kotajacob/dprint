// dprint - print specified values from desktop files to stdout
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"git.sr.ht/~kota/xdg/basedir"
	"git.sr.ht/~sircmpwn/getopt"
)

// variables set by config.mk
var (
	Version string
	Config  string
)

// regex to check if the file is a desktop file
var r = regexp.MustCompile(`(?m)(.*)\.desktop`)

// usage prints some basic usage information
func usage() {
	log.Fatal("Usage: dprint [-v] [-d path] [-i key:val] [-o key]")
}

// visit is called on each file or directory in the specified path
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() == false {
		if r.MatchString(info.Name()) {
			fmt.Println(info.Name())
		}
	}
	return nil
}

// set the config path to the XDG standard location if not set with -d
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
	// set dir to default XDG path if blank
	dir = setConfig(dir)

	// walk the dir and call visit for each file or subdir
	err = filepath.Walk(dir, visit)
	if err != nil {
		fmt.Printf("Failed reading directory: %v\n", err)
		os.Exit(1)
	}

	// TEST
	fmt.Println(in, out)
}
