package main

import (
	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/internal/hook"
	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/internal/wrapper"
)

func main() {
	wrapper.NewWrapper(hook.IstioProxyHalt).
		Prepare().
		Run().
		Exit()
}
