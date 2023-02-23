package tests

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func CheckError(t *testing.T, expectedErr error, expectedErrPattern string, actualErr error) {
	if expectedErr != nil || len(expectedErrPattern) > 0 {
		if actualErr == nil {
			t.Errorf("error was not returned as expected")
		}
		if expectedErr != nil && !errors.Is(actualErr, expectedErr) {
			t.Errorf("wrong error was returned: expected `%v`, got `%v`", expectedErr, actualErr)
		}
		if len(expectedErrPattern) > 0 && !strings.Contains(actualErr.Error(), expectedErrPattern) {
			t.Errorf("wrong error was returned: expected pattern `%v`, got `%v`", expectedErrPattern, actualErr.Error())
		}
	} else if actualErr != nil {
		t.Errorf("error was returned: %v", actualErr)
	}
}

// https://go.dev/play/p/Fzj1k7jul7z

func SilentExecute(execute func(args []string, version string) error, args []string, version string) (stdout string, stderr string, err error) {

	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		return
	}
	stderrReader, stderrWriter, err := os.Pipe()
	if err != nil {
		return
	}

	{
		origStdout := os.Stdout
		os.Stdout = stdoutWriter
		defer func() { os.Stdout = origStdout }()
	}
	{
		origStderr := os.Stderr
		os.Stderr = stderrWriter
		defer func() { os.Stderr = origStderr }()
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		var b bytes.Buffer
		defer wg.Done()
		if _, err := io.Copy(&b, stdoutReader); err != nil {
			return
		}
		stdout = b.String()
	}()
	wg.Add(1)
	go func() {
		var b bytes.Buffer
		defer wg.Done()
		if _, err := io.Copy(&b, stderrReader); err != nil {
			return
		}
		stderr = b.String()
	}()

	err = execute(args, version)
	stderrWriter.Close()
	stderrReader.Close()
	wg.Wait()
	return

}
