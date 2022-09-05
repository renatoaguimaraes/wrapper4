package wrapper

import (
	"log"
	"os"
)

func (w *pluginWrapper) Prepare() Wrapper {
	if w.IsPrepared() {
		log.Println("wrapper:", "already prepared")
		return w
	}

	log.Println("wrapper:", "preparing command")
	log.Println("wrapper:", os.Args[1:])
	if len(os.Args) == 1 {
		log.Println("wrapper:", "there's nothing to wrap 🤷‍♂️")
		return w
	}
	w.cmd = os.Args[1]
	w.args = []string{}
	if len(os.Args) > 2 {
		w.args = os.Args[2:]
	}
	log.Println("wrapper:", "prepared")
	return w
}
