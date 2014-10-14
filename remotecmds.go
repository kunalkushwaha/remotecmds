package main

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"flag"
	"fmt"
	"github.com/kunalkushwaha/remotecmds/plugins"
)

/*
The basic idea of this remotecmds is instead of client server architecture,
execute commands on remote servers using ssh.
*/
//Implement various plugins for actions.

func main() {

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	//var wg sync.WaitGroup
	var session *ssh.Session
	var err error
	var buf bytes.Buffer

	//Command line parsing.
	node := flag.String("node", "", "node name or IP @ cmd to be executed")
	cmd := flag.String("cmd", "", "cmds { time, CPU, RAM, say, scr }")
	flag.Parse()

	session, err = connect_to_machine(*node + ":22")
	if err != nil {
		panic("Failed to connect: " + err.Error())
	}

	switch *cmd {
	case "say":
		fmt.Println("executing say ")
	case "time":
		time := plugins.TimePlugin{plugins.WorkerType{session, buf}}
		time.Init()
		out, _ := time.Execute("get time")
		fmt.Println(out)
	case "CPU", "cpu":
		fmt.Println("executing CPU")
	case "RAM", "ram":
		fmt.Println("executing RAM")
	case "scr":
		fmt.Println("executing scr cap")
		//	gnome-screenshot -f ~/Desktop/screen.png
	default:
		fmt.Println("<Usage>")
		fmt.Println("./remotecmds -node=<nodename> -cmd=<cmds>")
		flag.PrintDefaults()
	}

	session.Close()
}
