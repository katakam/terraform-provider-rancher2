package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Types

type packetCredentialConfig struct {
	APIKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
}

//Schemas

func cloudCredentialPacketFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Packet api key",
		},
	}

	return s
}
