package main

import (
	"fmt"
	"os"
	"os/exec"
)// Creating main file to be executed when running `docker run image <command> <parameters>

func main () {
  switch os.args[1] {
    case "run":
      run ()
    default:
      panic("command not found")
  }
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func must (err error) {
  if err != nill
    panic(err)
}
