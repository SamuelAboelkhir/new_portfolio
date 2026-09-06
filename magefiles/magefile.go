//go:build mage

//mage:multiline

// Set the general description you want to have displayed with mage -l here.
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Dev() {
	mg.Deps(Tailwind, Templ)
}

func Tailwind() error {
	cmd := exec.Command(
		"./bin/tailwindcss-linux-x64",
		"-i", "./views/css/styles.css",
		"-o", "./public/styles.css",
		"--watch",
	)
	fmt.Println("Starting tailwind watcher")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Generate() error {
	cmd := exec.Command("templ", "generate")
	fmt.Println("Doing initial templ generation before running the server")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Templ() error {
	cmd := exec.Command("templ", "generate", "--watch", "--proxy", "http://localhost:8080", "--cmd", "air")
	fmt.Println("Starting templ watcher")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
