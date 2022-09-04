package wrapper

import (
	"os/exec"

	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/pkg/plugin"
)

type Wrapper interface {
	Prepare() Wrapper
	Run() Wrapper
	Exit()

	IsPrepared() bool
	IsProcessed() bool
	HasError() bool
}

type pluginWrapper struct {
	cmd     string
	args    []string
	proc    *exec.Cmd
	procErr error
	plugin  plugin.PluginRunner
}

// NewPluginWrapper returns a new Wrapper whose the plugin function should be informed.
//
//  plugin := plugin.Load("plugin.so")
//  wrapper.NewPluginWrapper(plugin).
//  Prepare().
//  Run().
//  Exit()
func NewPluginWrapper(plugin plugin.PluginRunner) Wrapper {
	return &pluginWrapper{
		plugin: plugin,
	}
}

func (w *pluginWrapper) IsPrepared() bool {
	return w.cmd != ""
}

func (w *pluginWrapper) IsProcessed() bool {
	return w.proc != nil
}

func (w *pluginWrapper) HasError() bool {
	return w.procErr != nil
}
