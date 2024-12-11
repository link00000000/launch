package env

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type MalformedEnvFileError struct {
	File    string
	Line    int
	Message string
}

func NewMalformedEnvFileError(file string, line int, message string) *MalformedEnvFileError {
	return &MalformedEnvFileError{file, line, message}
}

// Implements [error]
func (err MalformedEnvFileError) Error() string {
	return fmt.Sprintf("malformed .env file at %s:%d: %s", err.File, err.Line, err.Message)
}

func Read(r io.Reader) (map[string]string, error) {

	scanner := bufio.NewScanner(r)
	env := make(map[string]string)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		line = trimWhitespace(line)
		line = stripComments(line)
		key, value := parseLine(line)

		if len(key) > 0 && len(value) > 0 {
			env[string(key)] = string(value)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}

func ReadFile(name string) (map[string]string, error) {
	f, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	return Read(f)
}

func stripComments(s string) string {
	inEscape := false
	inQuotes := false

	var strippedLine []byte

loop:
	for i := 0; i < len(s); i++ {
		c := rune(s[i])

		switch true {
		case c == '\\' && !inEscape:
			inEscape = true
		case (c == '"' || c == '\'') && !inEscape:
			inQuotes = !inQuotes
		case c == '#' && !inEscape && !inQuotes:
			break loop
		}

		strippedLine = append(strippedLine, byte(c))
	}

	return string(strippedLine)
}

func parseLine(s string) (string, string) {
	type part uint8
	const (
		keyPart part = iota
		equalsPart
		valuePart
	)

	valuePartIsNext := func(s string) bool {
		return regexp.Match("\\s*=")
	}

	inEscape := false
	inQuotes := false
	inPart := keyPart

	var key []byte
	var value []byte

	s = trimWhitespace(s)

	for i := 0; i < len(s); i++ {
		c := rune(s[i])

		switch true {
		case c == '\\' && !inEscape:
			inEscape = true
		case (c == '"' || c == '\'') && !inEscape:
			inQuotes = !inQuotes
		case !inEscape && !inQuotes && inPart == keyPart && regexp.Match("\\s*=", []byte(s)):
			inPart = equalsPart
		case c == '=' && !inEscape && !inQuotes && inPart == keyPart:
			inPart = valuePart
		case inPart == keyPart:
			inEscape = false
			key = append(key, byte(c))
		case inPart == valuePart:
			inEscape = false
			value = append(value, byte(c))
		}
	}

	return string(key), string(value)
}

func trimWhitespace(s string) string {
	return strings.Trim(s, " \t\r")
}
