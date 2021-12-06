package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	packetConfigDriver = "packet"
)

//Types

type packetConfig struct {
	APIKey          string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
	AuthToken       string `json:"authToken,omitempty" yaml:"authToken,omitempty"`
	BillingCycle    string `json:"billingCycle,omitempty" yaml:"billingCycle,omitempty"`
	FacilityCode    string `json:"facilityCode,omitempty" yaml:"facilityCode,omitempty"`
	HwReservationID string `json:"hwReservationId,omitempty" yaml:"hwReservationId,omitempty"`
	MetroCode       string `json:"metroCode,omitempty" yaml:"metroCode,omitempty"`
	Os              string `json:"os,omitempty" yaml:"os,omitempty"`
	Plan            string `json:"plan,omitempty" yaml:"plan,omitempty"`
	ProjectID       string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	SpotInstance    bool   `json:"spotInstance,omitempty" yaml:"spotInstance,omitempty"`
	SpotPriceMax    string `json:"spotPriceMax,omitempty" yaml:"spotPriceMax,omitempty"`
	TerminationTime string `json:"terminationTime,omitempty" yaml:"terminationTime,omitempty"`
	UaPrefix        string `json:"uaPrefix,omitempty" yaml:"uaPrefix,omitempty"`
	Userdata        string `json:"userdata,omitempty" yaml:"userdata,omitempty"`
}

//Schemas

func packetConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"api_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Authentication Key (deprecated name, use Auth Token)",
		},
		"auth_token": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Equinix Metal Authentication Token",
		},
		"billing_cycle": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "hourly",
			Description: "Equinix Metal billing cycle, hourly or monthly",
		},
		"facility_code": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Equinix Metal facility code",
		},
		"hw_reservation_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Equinix Metal Reserved hardware ID",
		},
		"metro_code": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Equinix Metal metro code (\"dc\" is used if empty and facility is not set)",
		},
		"os": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "ubuntu_20_04",
			Description: "Equinix Metal OS",
		},
		"plan": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "c3.small.x86",
			Description: "Equinix Metal Server Plan",
		},
		"project_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Equinix Metal Project Id",
		},
		"spot_instance": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Request a Equinix Metal Spot Instance",
		},
		"spot_price_max": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The maximum Equinix Metal Spot Price",
		},
		"termination_time": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Request a Equinix Metal Spot Instance",
		},
		"ua_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Prefix the User-Agent in Equinix Metal API calls with some 'product/version' 0.6.0 packet",
		},
		"userdata": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "File contents for userdata",
		},
	}

	return s
}
