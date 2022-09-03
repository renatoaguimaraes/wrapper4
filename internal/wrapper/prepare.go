package wrapper

import (
	"log"
	"os"
)

func (w *Wrapper) Prepare() *Wrapper {
	if w.IsPrepared() {
		log.Fatalln("wrapper", "already prepared")
	}

	log.Println("wrapper", "preparing")
	if len(os.Args) == 1 {
		log.Fatalln("There is nothing to wrap 🤷‍♂️")
	}
	w.cmd = os.Args[1]
	w.args = []string{}
	if len(os.Args) > 2 {
		w.args = os.Args[2:]
	}
	log.Println("wrapper", "command", "->", w.cmd, w.args)
	log.Println("wrapper", "prepared")
	return w
}