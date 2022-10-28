package main

import (
	"github.com/gapidobri/radigo/cmd"
	"github.com/gapidobri/radigo/internal/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
