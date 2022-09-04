package main

import (
	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/internal/callback"
	"github.com/renatoaguimaraes/wrapper4-k8s-jobs/internal/wrapper"
)

func main() {
	wrapper.NewWrapper(callback.Quit).
		Prepare().
		Process().
		Exit()
}
