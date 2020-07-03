package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {

    // get `go` executable path
    goExecutable, _ := exec.LookPath( "claat" )

    // construct `go version` command
    cmdGoVer := &exec.Cmd {
        Path: goExecutable,
        Args: []string{ goExecutable, "export" },
        Stdout: os.Stdout,
        Stderr: os.Stdout,
    }

    // run `go version` command
    if err := cmdGoVer.Run(); err != nil {
        fmt.Println( "Error:", err );
    }

}