package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"

	b "github.com/joeblew99/gopherjs-electron/backend"
)

func main() {

	// start backend
	fmt.Println("starting backend... ")

	go func() {
		http.HandleFunc("/", b.HomeHandler)
		http.HandleFunc("/hello/", b.HelloHandler)
		http.HandleFunc("/api/test/", b.ApiTestHandler)

		fmt.Println("backend listening on localhost:8080")

		http.ListenAndServe(":8080", nil)

	}()

	// start frontend
	fmt.Println("starting frontend .... ")

	cmd := exec.Command("electron", "../frontend/.")
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		printError(err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}

}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
