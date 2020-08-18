package static

import (
	"github.com/markbates/pkger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ApplicationYaml struct {
	Slack struct {
		Url       string `yaml:"url"`
		Channel   string `yaml:"channel"`
		Messenger string `yaml:"messenger"`
		IconEmoji string `yaml:"icon_emoji"`
		IconUrl   string `yaml:"icon_url"`
	}
}

func ApplicationYamlLoad() (*ApplicationYaml, error) {
	f, err := pkger.Open("/config/application.yml")
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	yml := new(ApplicationYaml)
	if err := yaml.Unmarshal(b, yml); err != nil {
		return nil, err
	}
	return yml, nil
}
