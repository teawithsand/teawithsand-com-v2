package main

import (
	"embed"

	"github.com/teawithsand/webpage/cmd"
	"github.com/teawithsand/webpage/domain/webapp"
)

//go:embed __dist/*
var Assets embed.FS

func main() {
	webapp.EmbeddedAssets = Assets
	cmd.Execute()
}
