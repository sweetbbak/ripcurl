package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func dxecute(tmp *os.File, cmd ...string) {
	b := []string{"sh", "-c"}
	b = append(b, cmd...)
	b = append(b, tmp.Name())
	fmt.Println(b)
	exe := exec.Command(b[0], b[1], "'", "amy", "-f", tmp.Name(), "'")
	exe.Run()
}

func Execute(tmp *os.File) {
	exe := exec.Command("amy", "-f", tmp.Name())
	exe.Run()
}

func runCtx(cmd ...string) *os.Process {
	b := []string{"sh", "-c"}
	b = append(b, cmd...)
	exe := exec.Command(b[0], b[1:]...)
	exe.Run()
	return exe.Process
}

func pauseTTS(pr *os.Process) error {
	err := pr.Signal(syscall.SIGSTOP)
	return err
}

func continueTTS(pr *os.Process) error {
	err := pr.Signal(syscall.SIGCONT)
	return err
}

func stopTTS(pr *os.Process) error {
	err := pr.Signal(syscall.SIGKILL)
	return err
}

func startTTS(t string, c string) {
	tmp, _ := os.CreateTemp(".", "tts")
	tmp.Write([]byte(t))
	Execute(tmp)
}
