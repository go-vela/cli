package util

import "fmt"

// InvalidCommand returns a formatted error for improper flag usage
// with a CLI command
func InvalidCommand(f string) error {
	return fmt.Errorf("invalid command: Flag '--%s' is not set or is empty", f)
}

func InvalidFlagValue(v string, f string) error {
	return fmt.Errorf("invalid value '%s' for flag '--%s'", v, f)
}
