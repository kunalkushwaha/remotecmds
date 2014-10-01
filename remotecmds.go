package main

import "code.google.com/p/go.crypto/ssh"
import "fmt"
import "bytes"

/*
The basic idea of this remotecmds is instead of client server architecture,
execute commands on remote servers using ssh.
*/

func main() {
	
	
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("root123"),
		},
	}
	client, err := ssh.Dial("tcp", "172.26.126.11:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/ls -lh"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
