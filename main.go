package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	commandName = "setxkbmap"
	timeout     = time.Second * 3
	icon        = "keyboard"
)

var (
	errNotEnoughArguments = fmt.Errorf("language argument not provided")
	errCouldNotEncode     = fmt.Errorf("could not encode json")
)

type jsonOutput struct {
	Icon  string `json:"icon,omitempty"`
	State string `json:"state,omitempty"`
	Text  string `json:"text,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		output := &jsonOutput{
			Icon:  icon,
			State: "Critical",
			Text:  errNotEnoughArguments.Error(),
		}

		// Ignore error
		_ = json.NewEncoder(os.Stdout).Encode(output)
		os.Exit(1)
	}

	lang := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, commandName, lang)

	err := cmd.Run()
	if err != nil {
		output := &jsonOutput{
			Icon:  icon,
			State: "Critical",
			Text:  err.Error(),
		}

		// Ignore error
		_ = json.NewEncoder(os.Stdout).Encode(output)
		os.Exit(1)
	}

	output := &jsonOutput{
		Icon:  icon,
		State: "",
		Text:  lang,
	}

	// Ignore error
	_ = json.NewEncoder(os.Stdout).Encode(output)
	os.Exit(0)
}
