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
	username := ""

	// user.Current() may not be implemented on some linux distros (e.g. alpine)
	user, userErr := user.Current()
	if userErr == nil {
		username = user.Username
	}

	// Fallback to fetching the `LOGNAME` env
	if username == "" {
		username = os.Getenv("LOGNAME")
	}

	// Fallback to fetching the `USER` env
	if username == "" {
		username = os.Getenv("USER")
	}

	// Fallback to fetching `USERNAME` env
	if username == "" {
		username = os.Getenv("USERNAME")
	}

	// Fallback to calling `whoami` command
	if username == "" {
		whoami := exec.Command("whoami")
		whoamiStdout, whoamiErr := whoami.Output()
		if whoamiErr != nil {
			log.Fatalf("error: unable to determine current user: %v", whoamiErr)
		}
		username = strings.TrimSpace(string(whoamiStdout))
	}

	// Give up
	if username == "" {
		log.Fatalf("error: unable to determine current user: %v", userErr)
	}

	// Set default shell (do not set to `sudosh`; it may cause infinite loops)
	os.Setenv("SHELL", shell)

	// Fetch environment
	env := os.Environ()

	// Lookup path for `sudo`
	binary, sudoPathErr := exec.LookPath("sudo")
	if sudoPathErr != nil {
		log.Fatalf("error: find to find sudo: %v", sudoPathErr)
	}

	// Prepare `sudo` args
	if len(os.Args) < 2 {
		args = []string{"sudo", "-E", "-u", username, shell, "-l"}
	} else {
		args = append([]string{"sudo", "-E", "-u", username, shell}, os.Args[1:]...)
	}

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		log.Fatalf("error: exec failed: %v", execErr)
	}
}
