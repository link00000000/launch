package configurations

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GoConfigurationJSON struct {
	*BaseConfigurationJSON
	Request string            `json:"request"`
	Program string            `json:"program"`
	Args    []string          `json:"args"`
	Env     map[string]string `json:"env"`
	EnvFile string            `json:"envFile"`
}

type GoConfiguration struct {
	Request string
	Program string
	Args    []string
	Env     map[string]string
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

func NewGoConfigurationFromJSON(b []byte, vars Variables) (*GoConfigurationJSON, error) {
	var cfgJson GoConfigurationJSON

	err := json.Unmarshal(b, &cfgJson)
	if err != nil {
		return nil, err
	}

	var cfg GoConfiguration
	cfg.Program = SubstituteVariables(cfgJson.Program, vars)

	if cfgJson.EnvFile != "" {
		f, err := os.Open(cfgJson.EnvFile)

		if err != nil {
			return nil, err
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		lineNumber := 1
		for scanner.Scan() {
			line := scanner.Text()

			if len(line) == 0 || line[0] == '#' {
				continue
			}

			parts := strings.SplitN(line, "=", 2)

			if len(parts) != 2 {
				return nil, NewMalformedEnvFileError(cfgJson.EnvFile, lineNumber, "no separator")
			}

			cfg.Env[parts[0]] = parts[1]

			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return &cfgJson, nil
}

func (cfg *GoConfigurationJSON) Execute(cwd string) (int, error) {
	switch cfgJson.Request {
	case "launch":
		return cfgJson.launch(cwd)
	default:
		return 0, ErrUnsupportedRequest
	}
}

func (cfg *GoConfigurationJSON) launch(cwd string) (int, error) {
	if cfg.Program == "" {
		return 0, NewInvalidOptionError("program", "option required for \"launch\" request")
	}

	args := append([]string{"--allow-non-terminal-interactive=true", "debug", cfg.Program, "--"}, cfg.Args...)

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
