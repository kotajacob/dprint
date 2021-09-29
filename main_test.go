package main

import (
	"os"
	"testing"
)

type mockedFileInfo struct {
	// Embed this so we only need to add methods used by testable functions
	os.FileInfo
	name string
	dir  bool
}

func (m mockedFileInfo) IsDir() bool  { return m.dir }
func (m mockedFileInfo) Name() string { return m.name }

func TestCheckName(t *testing.T) {
	tests := []struct {
		name string
		dir  bool
		want bool
	}{
		{
			"hello.world",
			false,
			false,
		},
		{
			"hello.desktop",
			false,
			true,
		},
		{
			"hello.desktop",
			true,
			false,
		},
		{
			"desktop",
			false,
			false,
		},
		{
			".desktop",
			false,
			true,
		},
		{
			".",
			false,
			false,
		},
		{
			"",
			false,
			false,
		},
	}

	for _, test := range tests {
		var fi mockedFileInfo
		fi.name = test.name
		fi.dir = test.dir
		want := test.want
		got := checkName(fi)
		if want != got {
			t.Fatalf("checkName %s failed:\nwanted %v\ngot %v\n", test.name, want, got)
		}
	}
}
