package main

import (
	"os"
	"strings"

	"github.com/hexiosec/goapi/generator"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.SetEnvPrefix("goapi")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()
	v.AddConfigPath(".")

	pflag.StringP("input", "i", "./openapi.yml", "Input schema")
	pflag.StringP("output", "o", "./output", "Output folder")
	pflag.StringP("template", "t", "", "Generator template")
	pflag.String("templates-path", "./templates", "Path to template library")
	pflag.Parse()

	v.BindPFlags(pflag.CommandLine)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	g := generator.NewGenerator()

	if err := g.LoadSchema(v.GetString("input")); err != nil {
		panic(err)
	}

	if err := g.RenderTemplate(v.GetString("templates-path"), v.GetString("template"), v.GetString("output")); err != nil {
		panic(err)
	}
}
