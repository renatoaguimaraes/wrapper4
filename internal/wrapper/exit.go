package wrapper

import (
	"log"
	"os"
)

func (w *wrapper) Exit() {
	if w.HasError() {
		log.Println(w.procErr.Error())
	}
	os.Exit(w.proc.ProcessState.ExitCode())
}
