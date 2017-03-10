package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
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

	// Args that we pass to sudo
	var args []string

	var shell string
	var ext string

	// The command that was executed
	cmd := os.Args[0]

	ext = strings.TrimLeft(filepath.Ext(cmd), ".")
	// If no extension, default to bash
	if ext == "" {
		ext = "bash"
	}

	// Resolve extension to a shell
	shellFound, shellPathErr := exec.LookPath(ext)
	if shellPathErr != nil {
		log.Fatalf("error: find to find shell %v: %v", ext, shellPathErr)
	}
	shell = shellFound

	// Shell is always launched as current user
	user, userErr := user.Current()
	if userErr != nil {
		log.Fatalf("error: unable to determine current user: %v", userErr)
	}

	// Fetch environment
	env := os.Environ()

	// Lookup path for `sudo`
	binary, lookPathErr := exec.LookPath("sudo")
	if lookPathErr != nil {
		log.Fatalf("error: find to find sudo: %v", lookPathErr)
	}

	// Prepare `sudo` args
	if len(os.Args) < 2 {
		args = []string{"sudo", "-E", "-u", user.Username, "--", shell, "-l"}
	} else {
		args = append([]string{"sudo", "-E", "-u", user.Username, "-s", shell, "-c", "--"}, os.Args[1:]...)
	}

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatalf("error: exec failed: %v", execErr)
	}
}
