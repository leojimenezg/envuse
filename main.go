package envuse

import (
	"os"
	"fmt"
	"bufio"
	"slices"
)

type OpenFileError struct {
	File string
	Err error
}
func (e OpenFileError) Error() string {
	return fmt.Sprintf("failed to open env file %s: %v", e.File, e.Err)
}

type ReadFileError struct {
	File string
	Err error
}
func (e ReadFileError) Error() string {
	return fmt.Sprintf("failed to read from env file %s: %v", e.File, e.Err)
}

var envMap map[string]string = map[string]string{}

// LoadEnvFile loads and parses environment variables from a file.
// The file must contain key-value pairs separated by "=" on each line,
// whitespace around keys and values is ignored. Empty lines are skipped.
// Parameters:
//   - fileName: path to the environment file
// Returns nil on success, or OpenFileError/ReadFileError on failure.
func LoadEnvFile(fileName string) error {
	file, errFile := os.Open(fileName)
	if errFile != nil { return OpenFileError{ File: fileName, Err: errFile } }
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)  // Already done by default
	for fileScanner.Scan() {
		fileLine := []rune(fileScanner.Text())
		signIndex := slices.Index(fileLine, '=')
		if signIndex < 0 { continue }
		key := string(fileLine[:signIndex])
		if key == "" { continue }
		if (signIndex + 1) >= len(fileLine) { envMap[key] = "" }
		value := string(fileLine[(signIndex + 1):])
		envMap[key] = value
	}
	fileScannerErr := fileScanner.Err()
	if fileScannerErr != nil { return ReadFileError{ File: fileName, Err: fileScannerErr } }
	return nil
}

// GetEnv gets the value of an environment variable by key.
// Must be called after LoadEnvFile has successfully loaded the environment file.
// Parameters:
//   - envKey: exact name of the environment variable
// Returns the key value if exists, or an empty string if the key doesn't exist or file hasn't been loaded.
func GetEnv(envKey string) string { return envMap[envKey] }
