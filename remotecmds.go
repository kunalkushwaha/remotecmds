package main

import (
	"code.google.com/p/go.crypto/ssh"
	//"fmt"
	"bytes"
	"log"
	//"errors"
	"sync"
	"github.com/kunalkushwaha/remotecmds/plugins"
)

//import "container/list"

/*
The basic idea of this remotecmds is instead of client server architecture,
execute commands on remote servers using ssh.
*/


//Implement various plugins for actions.

//Break whole program in modules, for connection management, plugins and main function.
//commands
//	gnome-screenshot -f ~/Desktop/screen.png

func main() {

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var wg sync.WaitGroup
	var session *ssh.Session
	var err error
	session, err = connect_to_machine("172.26.126.11:22")
	if err != nil {
		panic("Failed to connect: " + err.Error())
	}

	
	var buf bytes.Buffer
	//session.Stdout = &b
	//if err := session.Run("/usr/bin/ls -lh /"); err != nil {
	//	panic("Failed to run: " + err.Error())
	//}




	//fmt.Print(&buf)
	ab := plugins.InfoPlugin{plugins.WorkerType{session, buf}, "test"}
	//ac := plugins.InfoPlugin{plugins.WorkerType{session, buf},"232"}
	PluginList := []plugins.WorkerPlugin{&ab}
	for i, _ := range PluginList {
		wg.Add(1)
		PluginList[i].Init()
		PluginList[i].Execute("executing plugin")
		//	defer wg.Done()
		
		PluginList[i].Cleanup()
		
		logger := log.New(&buf, "remotecmds: ", log.Lshortfile)
		logger.Print(buf.String())
	}
	 //wg.Wait()
	session.Close()
}


