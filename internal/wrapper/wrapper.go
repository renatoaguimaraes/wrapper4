package wrapper

import "os/exec"

type Wrapper struct {
	cmd      string
	args     []string
	proc     *exec.Cmd
	procErr  error
	callback func()
}

func NewWrapper(callback func()) *Wrapper {
	return &Wrapper{
		callback: callback,
	}
}

func (w *Wrapper) IsPrepared() bool {
	return w.cmd != ""
}

func (w *Wrapper) IsProcessed() bool {
	return w.proc != nil
}

func (w *Wrapper) HasError() bool {
	return w.procErr != nil
}
