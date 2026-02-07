package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	force := flag.Bool("force", false, "skip confirmation and delete immediately")
	dryRun := flag.Bool("dry-run", false, "list matching files without deleting")
	recursive := flag.Bool("recursive", false, "search subdirectories recursively")
	verbose := flag.Bool("verbose", false, "show each file as it is deleted")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: rmfilter [options] <directory> <pattern>\n\n")
		fmt.Fprintf(os.Stderr, "Removes files whose names contain <pattern>.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	dir := args[0]
	pattern := args[1]

	info, err := os.Stat(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot access %q: %v\n", dir, err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: %q is not a directory\n", dir)
		os.Exit(1)
	}

	// Collect matching files
	var matches []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: cannot access %q: %v\n", path, err)
			return nil
		}
		if info.IsDir() && !*recursive && path != dir {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.Contains(info.Name(), pattern) {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	if len(matches) == 0 {
		fmt.Println("No files found matching pattern:", pattern)
		return
	}

	fmt.Printf("Found %d file(s) matching %q:\n\n", len(matches), pattern)
	for _, f := range matches {
		fmt.Println("  ", f)
	}
	fmt.Println()

	if *dryRun {
		return
	}

	if !*force {
		fmt.Print("Delete these files? [y/N] ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if answer != "y" && answer != "yes" {
			fmt.Println("Aborted.")
			return
		}
	}

	deleted := 0
	for _, f := range matches {
		if err := os.Remove(f); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting %q: %v\n", f, err)
		} else {
			if *verbose {
				fmt.Println("Deleted:", f)
			}
			deleted++
		}
	}
	fmt.Printf("Deleted %d file(s).\n", deleted)
}
