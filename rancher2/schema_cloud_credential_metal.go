package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Types

type metalCredentialConfig struct {
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
}

//Schemas

func cloudCredentialMetalFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"apiKey": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Metal ApiKey",
		},
	}

	return s
}
