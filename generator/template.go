package generator

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/ettle/strcase"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type TemplateManifest struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Render      []*RenderTarget `json:"render,omitempty"`
	Post        *string         `json:"post,omitempty"`
}

type RenderTarget struct {
	Path     string `json:"path,omitempty"`
	For      string `json:"for,omitempty"`
	Template string `json:"template,omitempty"`
}

func GetManifest(path string) (*TemplateManifest, error) {
	log.Debug().Msgf("Loading template manifest %s", path)
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := &TemplateManifest{}
	err = yaml.Unmarshal(buf, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func GetTemplate(path string) (*template.Template, error) {
	log.Debug().Msgf("Loading template files %s", path)
	t := template.New(filepath.Base(path))

	// copied from: https://github.com/helm/helm/blob/8648ccf5d35d682dcd5f7a9c2082f0aaf071e817/pkg/engine/engine.go#L147-L154
	fm := map[string]interface{}{
		"include": func(name string, data interface{}) (string, error) {
			// https://stackoverflow.com/a/71091339
			buf := bytes.NewBuffer(nil)
			if err := t.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}

			return buf.String(), nil
		},
		"json": func(data interface{}) (string, error) {
			buf, err := json.Marshal(data)
			if err != nil {
				return "", err
			}
			return string(buf), err
		},
		"comment": func(prefix string, data interface{}) (string, error) {
			str := data.(string)
			lines := strings.Split(str, "\n")
			for idx, line := range lines {
				if line != "" {
					lines[idx] = prefix + " " + line
				}
			}
			return strings.Join(lines, "\n"), nil
		},
		"toGoPascalCase": func(data interface{}) (string, error) {
			return strcase.ToGoPascal(data.(string)), nil
		},
		"toGoCamelCase": func(data interface{}) (string, error) {
			return strcase.ToGoCamel(data.(string)), nil
		},
		"debugf": func(msg string, vars []interface{}) (string, error) {
			log.Debug().Msgf(msg, vars...)
			return "", nil
		},
		"warnf": func(msg string, vars []interface{}) (string, error) {
			log.Warn().Msgf(msg, vars...)
			return "", nil
		},
		"warn": func(data interface{}) (string, error) {
			log.Warn().Msgf("Template warning: %s", data)
			return "", nil
		},
	}

	t, err := t.Funcs(sprig.TxtFuncMap()).Funcs(fm).ParseGlob(path)
	if err != nil {
		return nil, err
	}

	return t, nil
}
