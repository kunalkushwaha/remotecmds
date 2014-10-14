package plugins

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
)

//Define Plugin interface.
type WorkerPlugin interface {
	Init() bool
	Execute(string) (string, error)
	Cleanup() bool
}

//Default member variables of each plugin.
type WorkerType struct {
	Context      *ssh.Session
	OutputBuffer bytes.Buffer
}
