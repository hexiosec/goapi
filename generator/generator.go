package generator

import (
	"errors"
	"os"
	"path"
	"strings"

	specv31 "github.com/hexiosec/goapi/spec-v3-1"
	"github.com/invopop/yaml"
	"github.com/rs/zerolog/log"
)

type Generator struct {
	doc *specv31.Document
}

type TemplateContext struct {
	Doc    *specv31.Document
	Node   any
	Config *TemplateManifest
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) LoadSchema(path string) error {
	log.Debug().Msgf("Loading spec %s", path)
	buf, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	log.Debug().Msgf("Parsing spec %s", path)
	g.doc = &specv31.Document{}
	err = yaml.Unmarshal(buf, g.doc)
	if err != nil {
		return err
	}

	log.Info().Msgf("OpenAPI: v%s", g.doc.OpenAPI)

	if g.doc.Info != nil {
		log.Info().Msgf("Specification: %s v%s", g.doc.Info.Title, g.doc.Info.Version)
	} else {
		log.Warn().Msg("No info block found in spec")
	}

	return Validate(g.doc)
}

func (g *Generator) RenderTemplate(name string, outPath string) error {
	if info, err := os.Stat(outPath); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	} else if info.IsDir() {
		log.Info().Msgf("Removing existing output folder")
		os.RemoveAll(outPath)
	}

	manifestPath := path.Join("./templates", name, "manifest.yml")
	manifest, err := GetManifest(manifestPath)
	if err != nil {
		return err
	}

	log.Info().Msgf("Template: %s", manifest.Name)

	templatePath := path.Join("./templates", name, "*.tmpl")
	template, err := GetTemplate(templatePath)
	if err != nil {
		return err
	}

	log.Debug().Msgf("Defined: %s", strings.TrimPrefix(template.DefinedTemplates(), "; defined templates are: "))

	err = os.Mkdir(outPath, os.ModePerm)
	if err != nil {
		return err
	}

	for _, target := range manifest.Render {
		switch target.For {
		case "root":
			f, err := os.OpenFile(path.Join(outPath, target.Path), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
			if err != nil {
				return err
			}

			err = template.ExecuteTemplate(f, target.Template, &TemplateContext{Doc: g.doc})
			if err != nil {
				return err
			}

		case "tag":
			for _, tag := range g.doc.Tags {
				targetPath := strings.Replace(target.Path, "*", tag.Name, 1)
				f, err := os.OpenFile(path.Join(outPath, targetPath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
				if err != nil {
					return err
				}

				err = template.ExecuteTemplate(f, target.Template, &TemplateContext{Doc: g.doc, Node: tag})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
