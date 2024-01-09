package main

import (
	"fmt"
	"log"
	"os/exec"
)

type Device struct {
	name      string
	ipaddress string
}

func (d Device) ping() {
	command := "ping"
	args := []string{"-c", "3", d.ipaddress}

	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Error: %s\nOutput: %s", err, output)
	}

	fmt.Printf("Ping successful to %s (%s)\n", d.name, d.ipaddress)
}

func main() {
	d := Device{
		name:      "test",
		ipaddress: "google.com",
	}
	d.ping()
}
