package wrapper

import (
	"log"
	"os"
)

func (w *pluginWrapper) Exit() {
	if w.HasError() {
		log.Println(w.procErr.Error())
	}
	os.Exit(w.proc.ProcessState.ExitCode())
}
