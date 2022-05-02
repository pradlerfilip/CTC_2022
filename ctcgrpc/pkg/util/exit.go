package util

import (
	"fmt"
	"os"
)

func Exit(msg string) {
	_, _ = fmt.Fprint(os.Stderr, msg)
	os.Exit(1)
}

func ExitOnError(err error) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

// ExitOnErrorf error must be last argument (is checked if nil)
func ExitOnErrorf(fmts string, args ...interface{}) {
	if len(args) == 0 || args[len(args)-1] == nil {
		return
	}

	_, _ = fmt.Fprint(os.Stderr, fmt.Errorf(fmts, args...))
	os.Exit(1)
}
