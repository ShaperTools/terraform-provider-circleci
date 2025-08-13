package circleci

import (
	circleci "github.com/shapertools/terraform-provider-circleci/circleci/circleci-go"
)

type Config struct {
	AuthToken string
}

func (c *Config) NewClient() *circleci.Client {
	return circleci.NewClient(c.AuthToken)
}
