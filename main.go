package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run(os.Args[2:]...)
	case "child":
		child(os.Args[2:]...)
	default:
		log.Fatal("Unknown command. Use run <command_name>, like `run /bin/bash` or `run echo hello`")
	}
}

func run(command ...string) {
	log.Println("Executing", command, "from run")
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, command[0:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run child using namespaces. The command provided will be executed inside that.
	must(cmd.Run())
}

func child(command ...string) {
	log.Println("Executing", command, "from child")

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		fmt.Println("ErrorQ", err)
		panic(err)
	}
}
