package rancher2

// Flatteners

func flattenCloudCredentialMetal(in *metalCredentialConfig, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	if len(in.ApiKey) > 0 {
		obj["apiKey"] = in.ApiKey
	}

	return []interface{}{obj}
}

// Expanders

func expandCloudCredentialMetal(p []interface{}) *metalCredentialConfig {
	obj := &metalCredentialConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["ApiKey"].(string); ok && len(v) > 0 {
		obj.ApiKey = v
	}
	return obj
}
