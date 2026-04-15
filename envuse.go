package envuse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OpenFileError struct {
	File string
	Err  error
}

func (e OpenFileError) Error() string {
	return fmt.Sprintf("failed to open env file %s: %v", e.File, e.Err)
}

type ReadFileError struct {
	File string
	Err  error
}

func (e ReadFileError) Error() string {
	return fmt.Sprintf("failed to read from env file %s: %v", e.File, e.Err)
}

// envMap stores key-value variables extracted from loaded file.
var envMap map[string]string = make(map[string]string)

// LoadEnvFile loads and parses environment variables from a file.
// The file must contain key-value pairs separated by "=" on each line,
// whitespace around keys and values is removed. Empty lines are skipped.
// Parameters:
//   - fileName: path to the environment file
//
// Returns nil on success, or OpenFileError/ReadFileError on failure.
func LoadEnvFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return OpenFileError{File: fileName, Err: err}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}

		value = strings.TrimSpace(value)
		envMap[key] = value
	}

	err = scanner.Err()
	if err != nil {
		return ReadFileError{File: fileName, Err: err}
	}

	return nil
}

// GetEnv gets the value of an environment variable by key.
// Must be called after LoadEnvFile has successfully loaded the environment file.
// Parameters:
//   - envKey: exact name of the environment variable
//
// Returns the key value if exists, or an empty string if the key doesn't exist or file hasn't been loaded.
func GetEnv(envKey string) string { return envMap[envKey] }
