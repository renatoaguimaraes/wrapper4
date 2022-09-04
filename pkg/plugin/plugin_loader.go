package plugin

import (
	"log"
	"plugin"
)

func Load(path string) PluginRunner {
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	plugin, err := p.Lookup("GetPlugin")
	if err != nil {
		log.Fatalln(err)
	}
	return plugin.(func() interface{})().(PluginRunner)
}
