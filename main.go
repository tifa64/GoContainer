package main

import "os"
// Creating main file to be executed when running `docker run image <command> <parameters>

func main () {
  switch os.args[1] {
    case "run":
      run ()
    default:
      panic("command not found")
  }
}

func run () {
  fmt.Printf("Running %v\n", os.args[2:])
}

func must (err error) {
  if err != nill
    panic(err)
}
