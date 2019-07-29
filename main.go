// dprint - print specified values from desktop files to stdout
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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

// readDir returns a slice of os.FileInfo's from dir
func readDir(dir string) ([]os.FileInfo, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// dirFi returns the file info's of directory dir
func dirFi(dir string) ([]os.FileInfo, error) {
	infos, err := readDir(dir)
	if err != nil {
		return nil, err
	}
	return infos, nil
}

// cName checks that the file info is a desktop file
func cName(fi os.FileInfo) bool {
	if fi.IsDir() == false {
		if r.MatchString(fi.Name()) {
			return true
		}
	}
	return false
}

// walk the file tree rooted at dir
func walk(dir string) ([]desktop.Entry, error) {
	infos, err := dirFi(dir)
	if err != nil {
		return nil, err
	}
	var entries []desktop.Entry
	for _, fi := range infos {
		if cName(fi) {
			path := filepath.Join(dir, fi.Name())
			dat, err := os.Open(path)
			if err != nil {
				return nil, err
			}
			entry, err := desktop.New(dat)
			if err != nil {
				return nil, err
			}
			entries = append(entries, *entry)
		}
	}
	return entries, nil
}

// split input string on : into key and value strings
func splitInput(in string) (string, string) {
	ins := strings.Split(in, ":")
	if len(ins) != 2 {
		fmt.Println("Error reading input key:value pair")
		os.Exit(1)
	}
	key := ins[0]
	val := ins[1]
	return key, val
}

// filter selection by key:value pair
func filter(in string, entries []desktop.Entry) []desktop.Entry {
	if in == "" {
		return entries
	}
	key, val := splitInput(in)
	// ignoring key for now and using Name
	fmt.Println(key)
	var selection []desktop.Entry
	for _, entry := range entries {
		if entry.Name == val {
			selection = append(selection, entry)
		}
	}
	return selection
}

func main() {
	// parse arguments in the getopt style
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
	// walk the directory to get an entries list
	entries, err := walk(dir)
	if err != nil {
		fmt.Printf("Failed getting entries: %q %v\n", dir, err)
	}
	// filter selection by key:value pair
	entries = filter(in, entries)
	// TEST
	for _, entry := range entries {
		fmt.Println(entry.Name)
	}
	fmt.Println(out)
}
