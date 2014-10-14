package plugins

import (
	"fmt"
)

type TimePlugin struct {
	PluginCommon WorkerType
}

func (i *TimePlugin) Init() bool {
	i.PluginCommon.Context.Stdout = &i.PluginCommon.OutputBuffer
	return true
}

func (i *TimePlugin) Execute(str string) (string, error) {
	fmt.Printf("Time : ")
	var err error

	if err := i.PluginCommon.Context.Run("/usr/bin/date"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	return i.PluginCommon.OutputBuffer.String(), err
}

func (i *TimePlugin) Cleanup() bool {
	return true
}
