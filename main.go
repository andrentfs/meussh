package main

import (
	"fmt"

	sshclient "github.com/helloyi/go-sshclient"
)

func main() {

	client, err := sshclient.DialWithPasswd("192.168.0.251:22", "digistar", "digistar")
	if err != nil {
		fmt.Println(err)
	}

	var script = `
  enable
  reload
`
	// It's as same as Cmd
	client.Script(script).Run()
	client.Script(script).Output()
	out2, err := client.Script(script).SmartOutput()

	if err != nil {
		// the 'out' is stderr output
		fmt.Println(err, out2)
	}
	// the 'out' is stdout output
	fmt.Println(string(out2))

	defer client.Close()
}
