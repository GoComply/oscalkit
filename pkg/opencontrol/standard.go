package opencontrol

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"

	"github.com/gocomply/oscalkit/types/oscal/catalog"
)

type Standard struct {
	Name     string   `yaml:"name"`
	Controls Controls `yaml:",inline"`
}

type Controls map[string]Control

type Control struct {
	Family      string `yaml:"family"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func NewStandard(catalog *catalog.Catalog) (*Standard, error) {
	if len(catalog.Controls) != 0 {
		return nil, fmt.Errorf("Direct controls in the catalog not implemented yet")
	}
	controls := map[string]Control{}

	for _, grp := range catalog.Groups {
		if len(grp.Groups) != 0 {
			return nil, fmt.Errorf("Groups inside groups in the catalog not implemented yet")
		}
		for _, ctrl := range grp.Controls {
			controls[strings.ToUpper(ctrl.Id)] = Control{
				Family:      strings.ToUpper(grp.Id),
				Name:        string(ctrl.Title),
				Description: "TODO",
			}
		}
	}

	return &Standard{
		Name:     string(catalog.Metadata.Title),
		Controls: controls,
	}, nil
}

func (std *Standard) SaveToFile(filename string) error {
	y, err := yaml.Marshal(std)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, y, 0644)
}
