package main

import (
	"os"

	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/internal/wrapper"
	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/pkg/plugin"
)

func main() {
	plugin := plugin.Load(os.Getenv("WRAPPER_PLUGIN_PATH"))
	wrapper.NewPluginWrapper(plugin).
		Prepare().
		Run().
		Exit()
}
