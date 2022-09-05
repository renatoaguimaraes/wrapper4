package wrapper

import (
	"log"
	"os"
)

func (w *pluginWrapper) Exit() {
	if !w.IsProcessed() {
		log.Println("wrapper:", "cmd not found or not prepared")
		os.Exit(1)
	}
	if w.HasError() {
		log.Println(w.procErr.Error())
	}
	os.Exit(w.proc.ProcessState.ExitCode())
}
