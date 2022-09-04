package wrapper

import "os/exec"

type Wrapper interface {
	Prepare() Wrapper
	Run() Wrapper
	Exit()

	IsPrepared() bool
	IsProcessed() bool
	HasError() bool
}

// New Wrapper returns a new Wrapper whose the hook function should be informed.
//
//  wrapper.NewWrapper(hook.IstioProxyHalt).
//  Prepare().
//  Run().
//  Exit()
func NewWrapper(hook func()) Wrapper {
	return &wrapper{
		hook: hook,
	}
}

type wrapper struct {
	cmd     string
	args    []string
	proc    *exec.Cmd
	procErr error
	hook    func()
}

func (w *wrapper) IsPrepared() bool {
	return w.cmd != ""
}

func (w *wrapper) IsProcessed() bool {
	return w.proc != nil
}

func (w *wrapper) HasError() bool {
	return w.procErr != nil
}
