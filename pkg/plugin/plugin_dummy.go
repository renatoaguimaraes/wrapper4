package plugin

type dummyPluginRunner struct{}

func NewDummyPluginRunner() PluginRunner {
	return &dummyPluginRunner{}
}

func (p dummyPluginRunner) Run() {}
