package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func system(cmd string) int {
	c := exec.Command("sh", "-c", cmd)
	c.Env = os.Environ()
	// c.Env = append(c.Env, env_vars)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err == nil {
		return 0
	}

	// Figure out the exit code
	if ws, ok := c.ProcessState.Sys().(syscall.WaitStatus); ok {
		if ws.Exited() {
			return ws.ExitStatus()
		}

		if ws.Signaled() {
			return -int(ws.Signal())
		}
	}
	return -1
}

func startTTS(t string, c string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	tmp, _ := os.CreateTemp(dir, "tts_")
	tmp.Write([]byte(t))
	tmp.Close()
	// cmd := strings.ReplaceAll(command, "{{file}}", tmp.Name())
	// defer os.Remove(tmp.Name())

	out := exec.Command("balcon", "-i", "-n", "Amy")
	out.Env = os.Environ()
	p, err := out.StdinPipe()
	p.Write([]byte(t))
	out.Stdout = os.Stdout
	out.Stderr = os.Stderr
	out.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)
}

func tts_stdin(t string, c string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	tmp, _ := os.CreateTemp(dir, "tts_")
	tmp.Write([]byte(t))
	tmp.Close()
	defer os.Remove(tmp.Name())

	cmd := strings.Split(command, " ")
	out := exec.Command(cmd[0], cmd[1:]...)

	// fmt.Println("Command: ", command)
	// fmt.Println("exec: ", cmd)

	out.Env = os.Environ()
	p, err := out.StdinPipe()
	p.Write([]byte(t))
	out.Stdout = os.Stdout
	out.Stderr = os.Stderr
	out.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)
}
