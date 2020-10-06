package opencontrol

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
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
	controls := Controls{}

	for _, grp := range catalog.Groups {
		if len(grp.Groups) != 0 {
			return nil, fmt.Errorf("Groups inside groups in the catalog not implemented yet")
		}
		for _, ctrl := range grp.Controls {
			controls.Add(&ctrl, strings.ToUpper(grp.Id))
		}
	}

	return &Standard{
		Name:     string(catalog.Metadata.Title.Raw),
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

func (controls Controls) Add(ctrl *catalog.Control, family string) {
	controls[oscalIdToOpencontrol(ctrl.Id)] = Control{
		Family:      family,
		Name:        string(ctrl.Title.Raw),
		Description: ctrl.StatementToMarkdown(),
	}
	for _, child := range ctrl.Controls {
		controls.Add(&child, family)
	}
}

func oscalIdToOpencontrol(id string) string {
	re := regexp.MustCompile(`^([a-z][a-z])-([0-9]+).([0-9]+)$`)
	match := re.FindStringSubmatch(id)
	if len(match) == 0 {
		return strings.ToUpper(id)
	}
	return strings.ToUpper(match[1]) + "-" + match[2] + " (" + match[3] + ")"
}
