package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	command := fmt.Sprintf("ssh-keygen -b 2048 -t rsa -f " + "test" + " -q -N \"\"\"\"")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
