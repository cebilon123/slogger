package config

import "github.com/Netflix/go-env"

var Environment environment

type environment struct {
	Port string `env:"PORT,required=true"`

	Extras env.EnvSet
}

func NewEnvironment() error {
	var e environment
	es, err := env.UnmarshalFromEnviron(&e)

	if err != nil {
		return err
	}

	e.Extras = es
	Environment = e

	return nil
}
