package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoInput(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	require.NotNil(t, err)

	exitErr, ok := err.(*exec.ExitError)
	require.True(t, ok)
	require.Equal(t, 1, exitErr.ExitCode())

	require.True(t, strings.Contains(stdout.String(), "The command is intended to work with pipes."))
	require.True(t, strings.Contains(stdout.String(), "Usage: echo 'Hello, world!' | cowsay"))
	require.True(t, strings.Contains(stderr.String(), "exit status 1"))
}

func TestCowsay(t *testing.T) {
	expected := "Hello, world!\n"

	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdin = strings.NewReader(expected)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	require.NoError(t, err)

	require.Equal(t, expected, stdout.String())
}
