package main

import "code.google.com/p/go.crypto/ssh"
import "fmt"
import "bytes"

//import "container/list"

/*
The basic idea of this remotecmds is instead of client server architecture,
execute commands on remote servers using ssh.
*/
/*
type job struct{
	param_list list
}

//Define Plugin interface.
type WorkerPlugin interface {
	init() bool
	execute() (string, error)
}
*/
//Implement various plugins for actions.

//Break whole program in modules, for connection management, plugins and main function.
//commands
//	gnome-screenshot -f ~/Desktop/screen.png

func main() {

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var session *ssh.Session
	var err error
	session, err = connect_to_machine("172.26.126.11:22")
	if err != nil {
		panic("Failed to connect: " + err.Error())
	}

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/ls -lh /"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	session.Close()
}


