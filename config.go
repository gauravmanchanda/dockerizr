package main

import (
	"reflect"
)

type Config struct {
	AppPort           string
	ContainerMemory   string `json:"containerMemory"`
	TechStack         string `json:"techStack"`
	ImageVersion      string `json:"imageVersion"`
	Framework         string `json:"framework"`
	QuayOrganization  string `json:"quayOrganization"`
	ServiceName       string `json:"serviceName"`
	ServiceVisibility string `json:"serviceVisibility"`
	SlackChannel      string `json:"slackChannel"`
	Template          string `json:"template"`
}

func (c *Config) setAppPort() {
	if c.ServiceVisibility == "private" {
		c.AppPort = "8080"
	} else {
		c.AppPort = "80"
	}
}

func (c Config) getField(field string) string {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}
