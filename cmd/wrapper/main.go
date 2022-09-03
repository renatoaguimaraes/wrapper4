package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	log.Println("wrapper init")
	cmd, args := prepare()
	proc, err := process(cmd, args)
	defer func() { exit(proc, err) }()
	quit()
}

func prepare() (string, []string) {
	if len(os.Args) == 1 {
		log.Fatalln("There is nothing to wrap 🤷‍♂️")
	}
	cmd := os.Args[1]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	log.Println("cmd:", cmd, args)
	return cmd, args
}

func process(cmd string, args []string) (*exec.Cmd, error) {
	proc := exec.Command(cmd, args...)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	err := proc.Run()

	return proc, err
}

func exit(proc *exec.Cmd, err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(proc.ProcessState.ExitCode())
}

func quit() {
	req, err := http.NewRequest(http.MethodPost, "http://localhost:15000/quitquitquit", bytes.NewReader([]byte{}))
	if err != nil {
		log.Println("Invalid quit request", err)
		return
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	if _, err = client.Do(req); err != nil {
		log.Println("No istio-proxy running to quit", err)
	} else {
		log.Println("Success to send quit message to istio-proxy")
	}
}
