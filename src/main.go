//go:build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Process => %v [%d]\n", os.Args, os.Getpid())
	args := os.Args
	if len(args) < 2 {
		panic(fmt.Sprint("not defend command"))
	}
	switch args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic(fmt.Sprint(args[1], " not defined"))
	}
}
func run() {
	cmd := exec.Command(os.Args[0], append([]string{"child"}, os.Args[2])...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
func child() {
	cmd := exec.Command(os.Args[2])
	syscall.Sethostname([]byte("container"))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
