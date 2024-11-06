package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/ettle/strcase"
	"github.com/invopop/yaml"
	"github.com/rs/zerolog/log"
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

func (g *Generator) GetManifest(name string) (*TemplateManifest, error) {
	log.Debug().Msgf("Loading template manifest %s", name)
	var buf []byte
	var err error

	if g.extTemplates == nil {
		buf, err = g.defaultTemplates.ReadFile(path.Join("templates", name, "manifest.yml"))
	} else {
		buf, err = os.ReadFile(path.Join(*g.extTemplates, name, "manifest.yml"))
	}

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

func (g *Generator) GetTemplate(name string) (*template.Template, error) {
	log.Debug().Msgf("Loading template files %s", name)
	t := template.New("root")

	t.Funcs(sprig.TxtFuncMap())

	t.Funcs(map[string]interface{}{
		"include": func(name string, data interface{}) (string, error) {
			// copied from: https://github.com/helm/helm/blob/8648ccf5d35d682dcd5f7a9c2082f0aaf071e817/pkg/engine/engine.go#L147-L154
			// https://stackoverflow.com/a/71091339
			buf := bytes.NewBuffer(nil)
			if err := t.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}

			return buf.String(), nil
		},
		"deref": func(data interface{}) (string, error) {
			if ptr, ok := data.(*float64); ok {
				val := *ptr
				if float64(int(val)) == val {
					return fmt.Sprintf("%d", int(val)), nil
				}
				return fmt.Sprintf("%f", *ptr), nil
			} else if ptr, ok := data.(*int); ok {
				return fmt.Sprintf("%d", *ptr), nil
			}
			return "", fmt.Errorf("deref: unrecognised %s", data)
		},
		"json": func(data interface{}) (string, error) {
			buf, err := json.Marshal(data)
			if err != nil {
				return "", err
			}
			return string(buf), err
		},
		"yaml": func(data interface{}) (string, error) {
			buf, err := yaml.Marshal(data)
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
		"toSnakeCase": func(data interface{}) (string, error) {
			return strcase.ToSnake(data.(string)), nil
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
	})

	if g.extTemplates == nil {
		_, err := t.ParseFS(g.defaultTemplates, "templates/"+name+"/*.tmpl")
		return t, err
	}

	_, err := t.ParseGlob(*g.extTemplates + "/" + name + "/*.tmpl")
	return t, err
}
