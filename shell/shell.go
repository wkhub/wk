package shell

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/wkhub/wk/fs"
)

func Current() Shell {
	// out, _ := exec.Command("ps", "-p $$", "-ocomm=").Output()
	// shell := string(out[:])
	// fmt.Print(os.Getppid())
	// shell := os.Getenv("SHELL")
	return ZSH
}

type Shell interface {
	Run(session Session)
	Exec(session Session) int
	Eval(session Session)
	Rc() string
}

type Helper struct {
	Name string
}

func (sh Helper) configDir() string {
	wkHome := fs.Home()
	return filepath.Join(wkHome, "shells", sh.Name)
}

func (sh Helper) ensureConfigDir() {
	if err := os.MkdirAll(sh.configDir(), 0755); err != nil {
		panic(err)
	}
}

func (sh Helper) configFile(name string) string {
	return filepath.Join(sh.configDir(), name)
}

// Command build a command binded to a given session
// stdin, stdout and stder are binded to os ones
func Command(cmd string, session Session) *exec.Cmd {
	command := exec.Command(cmd)
	command.Env = session.Env.Environ()
	command.Dir = session.Cwd
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return command
}

// RunWithExitCode run a command and returns the exit code
func RunWithExitCode(cmd *exec.Cmd) int {
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.Sys().(syscall.WaitStatus).ExitStatus()
		}
		return -1
	}
	return cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
}
