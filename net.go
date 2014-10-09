package main

import "code.google.com/p/go.crypto/ssh"

func connect_to_machine(IP_name string) (*ssh.Session, error) {

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

	return session, err

}
