//go:build mage

//mage:multiline

// Set the general description you want to have displayed with mage -l here.
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Starts the tailwind and templ watchers, generating go files from templ files and starts the server
func Dev() {
	mg.Deps(Tailwind, Templ, Air)
}

// Generates styles.css and starts the tailwind watcher
func Tailwind() error {
	cmd := exec.Command(
		"tailwindcss",
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

// Generates go files from templ files
func Generate() error {
	cmd := exec.Command("templ", "generate")
	fmt.Println("Doing initial templ generation before running the server")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Starts the templ generator and watcher using AIR to watch the code base for file changes
// Used port 7331 by default a the proxy server for templ
func Templ() error {
	godotenv.Load(".env")
	url := fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))
	cmd := exec.Command("templ", "generate", "--watch", "--proxy", url)
	fmt.Println("Starting templ watcher")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Air() error {
	cmd := exec.Command("air")

	fmt.Println("Starting air")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
