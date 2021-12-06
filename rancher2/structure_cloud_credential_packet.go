package rancher2

// Flatteners

func flattenCloudCredentialPacket(in *packetCredentialConfig, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	if len(in.APIKey) > 0 {
		obj["apiKey"] = in.APIKey
	}

	return []interface{}{obj}
}

// Expanders

func expandCloudCredentialPacket(p []interface{}) *packetCredentialConfig {
	obj := &packetCredentialConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["ApiKey"].(string); ok && len(v) > 0 {
		obj.APIKey = v
	}
	return obj
}
