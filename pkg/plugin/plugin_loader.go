package plugin

import (
	"log"
	"plugin"
)

func Load(path string) PluginRunner {
	p, err := plugin.Open(path)
	if err != nil {
		log.Println("plugin:", "load fail:", path, err)
		return NewDummyPluginRunner()
	}
	wp, err := p.Lookup("GetPlugin")
	if err != nil {
		log.Println("plugin:", "GetPlugin function not found:", path)
		return NewDummyPluginRunner()
	}
	gp, ok := wp.(func() interface{})
	if !ok {
		log.Println("plugin:", "GetPlugin can't be cast to 'func() interface{}'")
		return NewDummyPluginRunner()
	}
	pr, ok := gp().(PluginRunner)
	if !ok {
		log.Println("plugin:", "GetPlugin doesn't return a PluginRunner instance")
		return NewDummyPluginRunner()
	}
	return pr
}
