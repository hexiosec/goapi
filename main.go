package main

import (
	"embed"
	"fmt"
	"os"
	"time"

	"github.com/hexiosec/goapi/generator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

//go:embed templates/*
var TemplateFS embed.FS

func main() {
	inputSchema := pflag.StringP("input", "i", "./openapi.yml", "Input schema")
	outputPath := pflag.StringP("output", "o", "./output", "Output folder")
	template := pflag.StringP("template", "t", "", "Generator template")
	templatesPath := pflag.String("templates-path", "", "Path to template library")
	verbose := pflag.BoolP("verbose", "v", false, "Turn on verbose messaging")
	help := pflag.Bool("help", false, "Show help")
	pflag.Parse()

	if help != nil && *help {
		pflag.PrintDefaults()
		os.Exit(0)
	}

	if *template == "" {
		fmt.Println("No template specified")
		pflag.PrintDefaults()
		os.Exit(1)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if *templatesPath == "" {
		templatesPath = nil
	}

	g := generator.NewGenerator(TemplateFS, templatesPath)

	if err := g.LoadSchema(*inputSchema); err != nil {
		log.Fatal().Msgf("Load schema failed: %s", err)
	}

	if err := g.RenderTemplate(*template, *outputPath); err != nil {
		log.Fatal().Msgf("Generate failed: %s", err)
	}
}
