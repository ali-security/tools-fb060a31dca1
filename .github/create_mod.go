//go:build ignore
// +build ignore

// create_mod builds a Go module zip from a directory.
// Usage: go run create_mod.go <dir> <modpath> <version> <outzip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <dir> <modpath> <version> <outzip>")
		os.Exit(2)
	}
	dir, modpath, version, outzip := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	f, err := os.Create(outzip)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := zip.CreateFromDir(f, module.Version{Path: modpath, Version: version}, dir); err != nil {
		f.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
