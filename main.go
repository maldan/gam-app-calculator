package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-calculator/internal/app/calculator"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
