package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xdg/desktop"
)

// ByPopularity implements sort.Interface for []desktop.Entry based on popularity.
type ByPopularity []desktop.Entry

func (a ByPopularity) Len() int      { return len(a) }
func (a ByPopularity) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPopularity) Less(i, j int) bool {
	iPop, err := getPop(a[i])
	if err != nil {
		return true
	}
	jPop, err := getPop(a[j])
	if err != nil {
		return false
	}
	return iPop > jPop
}

// popUp increments the popularity count for a desktop entry.
func popUp(entry desktop.Entry) error {
	pop, err := getPop(entry)
	if err != nil {
		return fmt.Errorf("failed reading popularity: %v", err)
	}
	// increment and write new popularity
	pop++
	err = setPop(entry, pop)
	if err != nil {
		return fmt.Errorf("failed setting popularity: %v", err)
	}
	return nil
}

// getPop returns the popularity for an entry
func getPop(entry desktop.Entry) (int, error) {
	var pop int
	// check if cache directory exists and create if missing
	d := filepath.Join(basedir.CacheHome, "dprint")
	if err := os.MkdirAll(d, 0755); err != nil {
		return pop, fmt.Errorf("failed to create cache directory: %v", err)
	}

	// read popularity of cached entry if it exists
	name := filepath.Join(d, entry.Name)
	s, err := slurp(name)
	if err != nil {
		return pop, fmt.Errorf("cache file could not be read, but exists: %v", err)
	}
	if s == "" {
		s = "0"
	}
	s = strings.TrimSuffix(s, "\n")
	pop, err = strconv.Atoi(string(s))
	if err != nil {
		return pop, fmt.Errorf("cache file is corrupt: you may need to manually edit or delete the file: %q %v", name, err)
	}
	return pop, nil
}

// set the popularity count for a desktop entry.
func setPop(entry desktop.Entry, pop int) error {
	// check if cache directory exists and create if missing
	d := filepath.Join(basedir.CacheHome, "dprint")
	if err := os.MkdirAll(d, 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %v", err)
	}
	name := filepath.Join(d, entry.Name)
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create cache file: %v", err)
	}
	_, err = f.Write([]byte(strconv.Itoa(pop)))
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

// slurp a file into a string or return a blank string if the file doesn't exist.
func slurp(filename string) (string, error) {
	f, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed opening file: %v", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("failed reading file: %v", err)
	}
	err = f.Close()
	if err != nil {
		return string(b), fmt.Errorf("failed closing file: %v", err)
	}
	return string(b), nil
}
