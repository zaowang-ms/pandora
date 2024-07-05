// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataworkarounds

import (
	"fmt"
	sdkModels "github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/internal/components/apidefinitions/parser/dataworkarounds"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/internal/logging"
	importerModels "github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

func ApplyWorkarounds(input []importerModels.AzureApiDefinition) (*[]importerModels.AzureApiDefinition, error) {

	output := make([]importerModels.AzureApiDefinition, 0)
	logging.Tracef("Processing Swagger Data Workarounds..")
	for _, item := range input {

		// use a temporary `APIVersion` struct to bridge the logic between the old and new, until the full logic is refactored
		// this is intentional to allow removing of now unused code
		apiVersion := sdkModels.APIVersion{
			APIVersion: item.ApiVersion,
			Generate:   true,
			Preview:    item.IsPreviewVersion(),
			Resources:  item.Resources,
			Source:     sdkModels.AzureRestAPISpecsSourceDataOrigin,
		}
		fixed, err := dataworkarounds.Apply(item.ServiceName, apiVersion)
		if err != nil {
			return nil, fmt.Errorf("applying workarounds: %+v", err)
		}
		item.Resources = fixed.Resources

		output = append(output, item)
	}

	return &output, nil
}
