package configurations

import (
	"fmt"
	"os"
	"os/exec"
)

type GoConfiguration struct {
	*BaseConfiguration
	Request     string            `json:"request"`
	Program     string            `json:"program"`
	RuntimeArgs []string          `json:"runtimeArgs"` // Replace with args?
	Env         map[string]string `json:"env"`
}

type InvalidOptionError struct {
	Option string
	Reason string
}

func NewInvalidOptionError(option, reason string) InvalidOptionError {
	return InvalidOptionError{Option: option, Reason: reason}
}

func (err InvalidOptionError) Error() string {
	return fmt.Sprintf("invalid option %s: %s", err.Option, err.Reason)
}

func (cfg *GoConfiguration) Execute(cwd string) (int, error) {
	switch cfg.Request {
	case "launch":
		return cfg.launch(cwd)
	default:
		return 0, ErrUnsupportedRequest
	}
}

func (cfg *GoConfiguration) launch(cwd string) (int, error) {
	if cfg.Program == "" {
		return 0, NewInvalidOptionError("program", "option required for \"launch\" request")
	}

	args := append([]string{"--allow-non-terminal-interactive=true", "debug", cfg.Program, "--"}, cfg.RuntimeArgs...)

	cmd := exec.Command("dlv", args...)
	cmd.Dir = cwd
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	for k, v := range cfg.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	err := cmd.Run()

	if err, ok := err.(*exec.ExitError); ok {
		return err.ExitCode(), nil
	}

	if err != nil {
		return 0, err
	}

	return cmd.ProcessState.ExitCode(), nil
}
