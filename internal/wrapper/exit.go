package wrapper

import (
	"log"
	"os"
)

func (w *Wrapper) Exit() {
	if w.HasError() {
		log.Println(w.procErr.Error())
	}
	os.Exit(w.proc.ProcessState.ExitCode())
}
