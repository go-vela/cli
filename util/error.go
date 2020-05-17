package util

import "fmt"

// InvalidCommand returns a formatted error for improper flag usage
// with a CLI command
func InvalidCommand(f string) error {
	return fmt.Errorf("invalid command: Flag '--%s' is not set or is empty", f)
}

func InvalidFlag(f string) error {
	return fmt.Errorf("invalid command: Flag '--%s' is not valid", f)
}
