package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func init() {
	// make sure we only have one process and that it runs on the main thread
	// (so that ideally, when we Exec, we keep our user switches and stuff)
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}

func main() {
	log.SetFlags(0) // no timestamps on our logs

	// Fetch environment
	env := os.Environ()

	// Lookup path for `sudo`
	binary, lookErr := exec.LookPath("sudo")
	if lookErr != nil {
		panic(lookErr)
	}

	// Prepare `sudo` args
	if len(os.Args) <= 2 {
		args := []string{"sudo", "-u", "-i"}
	} else {
		args := append([]string{"sudo", "-u", "--"}, os.Args[2:])
	}

	err := syscall.Exec(binary, args, env)
	if err != nil {
		log.Fatalf("error: exec failed: %v", err)
	}
}
