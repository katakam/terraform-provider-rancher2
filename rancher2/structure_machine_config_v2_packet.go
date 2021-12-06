package rancher2

import (
	norman "github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	machineConfigV2PacketKind         = "PacketMachineTemplate"
	machineConfigV2PacketAPIVersion   = "rke-machine-config.cattle.io/v1"
	machineConfigV2PacketAPIType      = "rke-machine-config.cattle.io.packetmachine"
	machineConfigV2PacketClusterIDsep = "."
)

//Types

type machineConfigV2Packet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	APIKey            string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
	AuthToken         string `json:"authToken,omitempty" yaml:"authToken,omitempty"`
	BillingCycle      string `json:"billingCycle,omitempty" yaml:"billingCycle,omitempty"`
	FacilityCode      string `json:"facilityCode,omitempty" yaml:"facilityCode,omitempty"`
	HwReservationID   string `json:"hwReservationId,omitempty" yaml:"hwReservationId,omitempty"`
	MetroCode         string `json:"metroCode,omitempty" yaml:"metroCode,omitempty"`
	Os                string `json:"os,omitempty" yaml:"os,omitempty"`
	Plan              string `json:"plan,omitempty" yaml:"plan,omitempty"`
	ProjectID         string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	SpotInstance      bool   `json:"spotInstance,omitempty" yaml:"spotInstance,omitempty"`
	SpotPriceMax      string `json:"spotPriceMax,omitempty" yaml:"spotPriceMax,omitempty"`
	TerminationTime   string `json:"terminationTime,omitempty" yaml:"terminationTime,omitempty"`
	UaPrefix          string `json:"uaPrefix,omitempty" yaml:"uaPrefix,omitempty"`
	Userdata          string `json:"userdata,omitempty" yaml:"userdata,omitempty"`
}

type MachineConfigV2Packet struct {
	norman.Resource
	machineConfigV2Packet
}

// Flatteners

func flattenMachineConfigV2Packet(in *MachineConfigV2Packet) []interface{} {
	if in == nil {
		return nil
	}
	obj := make(map[string]interface{})
	if len(in.APIKey) > 0 {
		obj["apiKey"] = in.APIKey
	}
	if len(in.AuthToken) > 0 {
		obj["authToken"] = in.AuthToken
	}
	obj["billingCycle"] = in.BillingCycle
	if len(in.FacilityCode) > 0 {
		obj["facilityCode"] = in.FacilityCode
	}
	if len(in.HwReservationID) > 0 {
		obj["hwReservationId"] = in.HwReservationID
	}
	if len(in.MetroCode) > 0 {
		obj["metroCode"] = in.MetroCode
	}
	if len(in.Os) > 0 {
		obj["os"] = in.Os
	}
	if len(in.Plan) > 0 {
		obj["plan"] = in.Plan
	}
	if len(in.ProjectID) > 0 {
		obj["projectId"] = in.ProjectID
	}

	if in.SpotInstance {
		obj["spotInstance"] = in.SpotInstance
	}

	if len(in.SpotPriceMax) > 0 {
		obj["spotPriceMax"] = in.SpotPriceMax
	}

	if len(in.TerminationTime) > 0 {
		obj["terminationTime"] = in.TerminationTime
	}

	if len(in.UaPrefix) > 0 {
		obj["uaprefix"] = in.UaPrefix
	}

	if len(in.Userdata) > 0 {
		obj["userdata"] = in.Userdata
	}
	return []interface{}{obj}
}

// Expanders

func expandMachineConfigV2Packet(p []interface{}, source *MachineConfigV2) *MachineConfigV2Packet {
	if p == nil || len(p) == 0 || p[0] == nil {
		return nil
	}
	obj := &MachineConfigV2Packet{}

	if len(source.ID) > 0 {
		obj.ID = source.ID
	}
	in := p[0].(map[string]interface{})

	obj.TypeMeta.Kind = machineConfigV2PacketKind
	obj.TypeMeta.APIVersion = machineConfigV2PacketAPIVersion
	source.TypeMeta = obj.TypeMeta
	obj.ObjectMeta = source.ObjectMeta
	if v, ok := in["apikey"].(string); ok && len(v) > 0 {
		obj.APIKey = v
	}
	if v, ok := in["authToken"].(string); ok && len(v) > 0 {
		obj.AuthToken = v
	}
	if v, ok := in["billingCycle"].(string); ok && len(v) > 0 {
		obj.BillingCycle = v
	}

	if v, ok := in["facilityCode"].(string); ok && len(v) > 0 {
		obj.FacilityCode = v
	}

	if v, ok := in["hwReservationId"].(string); ok && len(v) > 0 {
		obj.HwReservationID = v
	}

	if v, ok := in["metroCode"].(string); ok && len(v) > 0 {
		obj.MetroCode = v
	}
	if v, ok := in["os"].(string); ok && len(v) > 0 {
		obj.Os = v
	}

	if v, ok := in["plan"].(string); ok && len(v) > 0 {
		obj.Plan = v
	}

	if v, ok := in["projectId"].(string); ok && len(v) > 0 {
		obj.ProjectID = v
	}

	if v, ok := in["spotInstance"].(bool); ok {
		obj.SpotInstance = v
	}

	if v, ok := in["spotPriceMax"].(string); ok && len(v) > 0 {
		obj.SpotPriceMax = v
	}

	if v, ok := in["terminationTime"].(string); ok && len(v) > 0 {
		obj.TerminationTime = v
	}

	if v, ok := in["uaPrefix"].(string); ok && len(v) > 0 {
		obj.UaPrefix = v
	}
	if v, ok := in["userdata"].(string); ok && len(v) > 0 {
		obj.Userdata = v
	}

	return obj
}
