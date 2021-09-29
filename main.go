// dprint - print specified values from desktop files to stdout
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~kota/xdg/basedir"
	"git.sr.ht/~kota/xdg/desktop"
	"git.sr.ht/~sircmpwn/getopt"
)

var (
	// Version is a semantic version number set at build time. It's configured
	// in config.mk
	Version string
	// Config represents the directory name under XDG_CONFIG_HOME where desktop
	// files are searched. This only matters if the -d option isn't used.
	// This value is normally set out build time and can be configured in
	// config.mk
	Config string
)

// usage prints some basic usage information
func usage() {
	log.Fatal("Usage: dprint [-v] [-d path] [-i key:val] [-o key]")
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
	// replace args if needed
	if dir == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			dir = scanner.Text()
		}
	}
	if in == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			in = scanner.Text()
		}
	}
	if out == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out = scanner.Text()
		}
	}
	// set dir to default XDG path if blank
	dir = getConfig(dir)
	// walk the directory to get an entries list
	entries, err := walk(dir)
	if err != nil {
		fmt.Printf("Failed getting entries: %q %v\n", dir, err)
	}
	// filter selection by key:value pair
	entries = filter(in, entries)
	// print output selections
	for _, entry := range entries {
		// print specified key
		if out != "" {
			fmt.Println(getOut(entry, out))
		} else {
			// print name as default
			fmt.Println(entry.Name)
		}
	}
}

// getConfig returns the config path.
func getConfig(d string) string {
	if d == "" {
		d = filepath.Join(basedir.ConfigHome, Config)
	}
	return d
}

// checkName checks that the file info is a desktop file.
func checkName(fi os.FileInfo) bool {
	if !fi.IsDir() {
		if filepath.Ext(fi.Name()) == ".desktop" {
			return true
		}
	}
	return false
}

// walk the file tree rooted at dir.
func walk(dir string) ([]desktop.Entry, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	infos, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	var entries []desktop.Entry
	for _, fi := range infos {
		if checkName(fi) {
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

// splitInput string on : into key and value strings.
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

// stripExec removes execution field codes from a string.
func stripExec(in string) string {
	out := strings.Replace(in, "%f", "", -1)
	out = strings.Replace(out, "%F", "", -1)
	out = strings.Replace(out, "%u", "", -1)
	out = strings.Replace(out, "%U", "", -1)
	out = strings.Replace(out, "%d", "", -1)
	out = strings.Replace(out, "%D", "", -1)
	out = strings.Replace(out, "%n", "", -1)
	out = strings.Replace(out, "%N", "", -1)
	out = strings.Replace(out, "%i", "", -1)
	out = strings.Replace(out, "%c", "", -1)
	out = strings.Replace(out, "%k", "", -1)
	out = strings.Replace(out, "%v", "", -1)
	out = strings.Replace(out, "%m", "", -1)
	return out
}

// checkKey returns true is the entry has the key and value.
func checkKey(entry desktop.Entry, key string, val string) bool {
	switch key {
	case "Version":
		if entry.Version == val {
			return true
		}
	case "Name":
		if entry.Name == val {
			return true
		}
	case "GenericName":
		if entry.GenericName == val {
			return true
		}
	case "Comment":
		if entry.Comment == val {
			return true
		}
	case "Icon":
		if entry.Icon == val {
			return true
		}
	case "URL":
		if entry.URL == val {
			return true
		}
	case "TryExec":
		if entry.TryExec == val {
			return true
		}
	case "Exec":
		if entry.Exec == val {
			return true
		}
	case "Path":
		if entry.Path == val {
			return true
		}
	}
	return false
}

// getOut returns the string tied to an output value.
func getOut(entry desktop.Entry, key string) string {
	switch key {
	case "Version":
		return entry.Version
	case "Name":
		return entry.Name
	case "GenericName":
		return entry.GenericName
	case "Comment":
		return entry.Comment
	case "Icon":
		return entry.Icon
	case "URL":
		return entry.URL
	case "TryExec":
		return entry.TryExec
	case "Exec":
		return entry.Exec
	case "StripExec":
		return stripExec(entry.Exec)
	case "Path":
		return entry.Path
	}
	return entry.Name
}

// filter selection by key:value pair.
func filter(in string, entries []desktop.Entry) []desktop.Entry {
	if in == "" {
		return entries
	}
	key, val := splitInput(in)
	var selection []desktop.Entry
	for _, entry := range entries {
		if checkKey(entry, key, val) {
			selection = append(selection, entry)
		}
	}
	return selection
}
