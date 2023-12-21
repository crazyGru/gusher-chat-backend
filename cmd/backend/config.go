package main

import (
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

type Config struct {
	DevelopmentMode string         `yaml:"developmentMode"`
	SwaggerGui      string         `yaml:"swaggerGui"`
	AuthProviderUrl string         `yaml:"authProviderUrl"`
	DbAddr          string         `yaml:"dbAddr"`
	DbUser          string         `yaml:"dbUser"`
	DbPass          string         `yaml:"dbPass"`
	DbName          string         `yaml:"dbName"`
	MockIdentity    map[string]int `yaml:"mockIdentity"`
}

func (c *Config) ReadFromFile(filePath string) error {
	var err error
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// replace ENV variables using pattern ${VARNAME}
	m := regexp.MustCompile(`\$\{(.+)}`)

	bytes = m.ReplaceAllFunc(bytes, func(b []byte) []byte {
		return []byte(os.Getenv(string(m.FindSubmatch(b)[1])))
	})

	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}
