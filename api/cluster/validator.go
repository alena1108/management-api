package cluster

import (
	"github.com/rancher/norman/httperror"
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
)

const (
	specField = "spec"
)

func Validator(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}) error {
	spec, ok := data[specField]
	if !ok {
		return nil
	}

	keysToCheck := map[string]bool{
		"importedConfig":                true,
		"embeddedConfig":                true,
		"googleKubernetesEngineConfig":  true,
		"azureKubernetesServiceConfig":  true,
		"rancherKubernetesEngineConfig": true,
	}

	specData, _ := convert.EncodeToMap(spec)

	found := false
	for key, _ := range keysToCheck {
		val, ok := specData[key]
		if ok {
			configData, _ := convert.EncodeToMap(val)
			if len(configData) > 0 {
				found = true
				break
			}
		}
	}

	if found {
		return nil
	}
	return httperror.NewAPIError(httperror.MissingRequired, "Config field is required")

}
