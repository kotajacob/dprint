// dprint - print specified values from desktop files to stdout
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"git.sr.ht/~kota/xdg/basedir"
	"git.sr.ht/~kota/xdg/desktop"
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

	// walk the dir and store file names
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Failed reaching path: %q %v\n", path, err)
			return err
		}
		if info.IsDir() == false {
			if r.MatchString(info.Name()) {
				// s = append(s, filepath.Join(dir, info.Name()))
				path := filepath.Join(dir, info.Name())
				dat, err := os.Open(path)
				if err != nil {
					fmt.Printf("Error opening desktop file: %v\n", err)
				}
				entry, err := desktop.New(dat)
				if err != nil {
					fmt.Printf("Error reading desktop file: %v\n", err)
				}
				fmt.Println(entry.Name)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Failed walking path: %q %v\n", dir, err)
	}

	// TEST
	fmt.Println(in, out)
}
