package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestVersionHelper(t *testing.T) {
	suite := []struct {
		flags     []string
		isVersion bool
	}{
		{[]string{"-v"}, true},
		{[]string{"--version"}, true},
		{[]string{"-version"}, true},
		{[]string{"-h"}, false},
		{[]string{"--help"}, false},
	}

	for _, tt := range suite {
		cli := NewCli("app", "0.0.1")
		cli.Out = &bytes.Buffer{}

		cli.Args = tt.flags
		cli.Run()

		if cli.isVersion() != tt.isVersion {
			t.Errorf("Expected %v got %v", tt.isVersion, cli.isVersion())
		}
	}
}

func TestHelpHelper(t *testing.T) {
	suite := []struct {
		flags  []string
		isHelp bool
	}{
		{[]string{"-h"}, true},
		{[]string{"--help"}, true},
		{[]string{"-help"}, true},
		{[]string{"-v"}, false},
		{[]string{"--version"}, false},
	}

	for _, tt := range suite {
		cli := NewCli("app", "0.0.1")
		cli.Out = &bytes.Buffer{}

		cli.Args = tt.flags
		cli.Run()

		if cli.isHelp() != tt.isHelp {
			t.Errorf("Expected %v got %v", tt.isHelp, cli.isHelp())
		}
	}
}

func TestHelpMessage(t *testing.T) {
	cli := NewCli("app", "0.0.1")
	buffer := &bytes.Buffer{}
	cli.Out = buffer
	cli.Args = []string{"-h"}

	cli.Run()

	if !strings.Contains(buffer.String(), "usage: app [--version] [--help] <command> [<args>]") {
		t.Errorf("Invalid help message, got: %s", buffer.String())
	}
}

func TestHelpMessageWithCommandList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	command := NewMockCommand(ctrl)
	command.EXPECT().Synopsis().Return("help message")

	cli := NewCli("app", "0.0.1")
	buffer := &bytes.Buffer{}
	cli.Commands = map[string]CommandFactory{
		"foo": func() Command {
			return command
		},
	}

	cli.Out = buffer
	cli.Args = []string{"-h"}

	cli.Run()

	if !strings.Contains(buffer.String(), "usage: app [--version] [--help] <command> [<args>]") {
		t.Errorf("Invalid help message, got: %s", buffer.String())
	}

	if !strings.Contains(buffer.String(), "help message") {
		t.Errorf("Invalid help message, got: %s", buffer.String())
	}
}

func TestVersionMessage(t *testing.T) {
	cli := NewCli("app", "0.0.1")
	buffer := &bytes.Buffer{}
	cli.Out = buffer
	cli.Args = []string{"-v"}

	cli.Run()

	if !strings.Contains(buffer.String(), "0.0.1") {
		t.Errorf("Invalid version message, got: %s", buffer.String())
	}
}

func TestSubcommandHelpMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	command := NewMockCommand(ctrl)
	command.EXPECT().Help().Return("long help message")

	cli := NewCli("app", "0.0.1")
	buffer := &bytes.Buffer{}
	cli.Commands = map[string]CommandFactory{
		"foo": func() Command {
			return command
		},
	}

	cli.Out = buffer
	cli.Args = []string{"-h", "foo"}

	cli.Run()

	if !strings.Contains(buffer.String(), "long help message") {
		t.Errorf("Missing command help message")
	}
}

func TestExecuteCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	command := NewMockCommand(ctrl)
	command.EXPECT().Run([]string{"-x", "ok"}).Return(124)

	cli := NewCli("app", "0.0.1")
	buffer := &bytes.Buffer{}
	cli.Commands = map[string]CommandFactory{
		"foo": func() Command {
			return command
		},
	}

	cli.Out = buffer
	cli.Args = []string{"foo", "-x", "ok"}
	exitCode := cli.Run()

	if exitCode != 124 {
		t.Errorf("Invalid exit code, wants %d got %d", 124, exitCode)
	}
}
