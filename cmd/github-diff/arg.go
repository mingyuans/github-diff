package main

import (
	"github.com/actions-go/toolkit/core"
	"os"
)

type ActionArg struct {
	LoggerLevel string `json:"logger_level"`
	Token       string `json:"token"`
	FileName    string `json:"filename"`
}

func ParseArg() ActionArg {
	var token = core.GetInputOrDefault("token", "")
	if len(token) == 0 {
		core.SetFailed("GitHub token is required. Please set the 'token' input.")
		os.Exit(1)
	}

	loggerLevel := core.GetInputOrDefault("logger-level", "info")
	fileName := core.GetInputOrDefault("file-name", "pr.diff")
	return ActionArg{
		LoggerLevel: loggerLevel,
		Token:       token,
		FileName:    fileName,
	}
}
