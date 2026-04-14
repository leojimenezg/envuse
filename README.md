# envuse

Package to easily load and use local environment files at runtime, avoiding the storage of credentials in the compiled binary.

## Summary
This package solves the problem of managing privileged information in a program, such as database credentials, API keys, and configuration values. It works with any text-based file but is designed specifically for `.env` formatted files.

## Features
The `envuse` package provides two simple functions to work with environment files.

### LoadFile
The `LoadFile(string)` function receives a relative path to locate and load the environment file at runtime. It parses the file content and stores all key-value pairs in an internal map.

The function accepts relative paths from the program's execution directory. While absolute paths are supported, relative paths are recommended since Go programs execute from the directory where they are invoked.

#### Example
```go
err := envuse.LoadFile("credentials.env")
if err != nil {
    log.Fatal(err)
}
```

### GetEnv
The `GetEnv(string)` function retrieves the value associated with a specified key. If the key exists, it returns its value; otherwise, it returns an empty string.

This function must be called after successfully loading a file with `LoadFile(string)`. The loaded data remains accessible throughout the program's lifecycle while staying encapsulated within the package scope.

#### Example
```go
host := envuse.GetEnv("host")
port := envuse.GetEnv("port")
```

## File Format
The environment file must follow these rules:
* Each key-value pair must be on a single line
* Keys and values are separated by an equal sign (`=`)
* Whitespace before/after keys, values, and the equal sign is automatically trimmed
* Lines without an equal sign are ignored
* Empty values are allowed (`KEY=` results in an empty string)
* Quoted values are supported (`KEY="value"`)
* If duplicate keys exist, the last occurrence takes precedence

### Example .env file
```
# This line is ignored (no equal sign)
DATABASE_HOST = localhost
DATABASE_PORT=5432
DATABASE_NAME = myapp_db
API_KEY="sk-1234567890abcdef"
EMPTY_VALUE=
```

## Errors
The `LoadFile(string)` function can return two types of errors.

### OpenFileError
This error indicates that the file could not be opened. Common causes include the file not existing, incorrect path, or insufficient permissions.

```go
type OpenFileError struct {
    File string
    Err error
}
```

Access error details using the struct fields or the `Error()` method.

### ReadFileError
This error indicates that the file was found and opened, but its content could not be read or parsed.

```go
type ReadFileError struct {
    File string
    Err error
}
```

Access error details using the struct fields or the `Error()` method.

## Notes
* I created this simple package because I needed to use environment file information without embedding sensitive data directly in the code or the compiled binary.
* I didn't like the existing solutions I found they seemed overcomplicated for such a straightforward task. More importantly, I didn't want to add a heavy dependency just to solve this problem.
* I don't plan to actively develop or extend this package in the near future, but I'm not ruling out improvements if needed.
