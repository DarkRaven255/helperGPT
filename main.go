package main

import (
	"helperGPT/config"
	"helperGPT/gui"
)

func main() {
	config.LoadConfig()
	gui.LoadGui()
}
