// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commonschema

import (
	"reflect"
	"strings"

	"github.com/hashicorp/pandora/tools/data-api-sdk/v1/models"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/components/parser/internal"
	importerModels "github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
)

var _ customFieldMatcher = systemAssignedIdentityMatcher{}

type systemAssignedIdentityMatcher struct{}

func (systemAssignedIdentityMatcher) CustomFieldType() importerModels.CustomFieldType {
	return importerModels.CustomFieldTypeSystemAssignedIdentity
}

func (systemAssignedIdentityMatcher) IsMatch(_ importerModels.FieldDetails, definition importerModels.ObjectDefinition, known internal.ParseResult) bool {
	if definition.Type != importerModels.ObjectDefinitionReference {
		return false
	}

	// retrieve the model from the reference
	model, ok := known.Models[*definition.ReferenceName]
	if !ok {
		return false
	}

	hasMatchingType := false
	hasPrincipalId := false
	hasTenantId := false

	for fieldName, fieldVal := range model.Fields {
		if strings.EqualFold(fieldName, "PrincipalId") {
			hasPrincipalId = true
			continue
		}

		if strings.EqualFold(fieldName, "TenantId") {
			hasTenantId = true
			continue
		}

		if strings.EqualFold(fieldName, "Type") {
			if fieldVal.ObjectDefinition == nil || fieldVal.ObjectDefinition.Type != importerModels.ObjectDefinitionReference {
				continue
			}
			constant, ok := known.Constants[*fieldVal.ObjectDefinition.ReferenceName]
			if !ok {
				continue
			}
			expected := map[string]string{
				"SystemAssigned": "SystemAssigned",
			}
			hasMatchingType = validateIdentityConstantValues(constant, expected)
			continue
		}

		// other fields
		return false
	}

	return hasMatchingType && hasPrincipalId && hasTenantId
}

func validateIdentityConstantValues(input models.SDKConstant, expected map[string]string) bool {
	if input.Type != models.StringSDKConstantType {
		return false
	}

	// we can't guarantee the casing on these, so we should parse this insensitively since it'll be swapped
	// out anyway
	actual := make(map[string]string, 0)
	for k, v := range input.Values {
		// Some Identity Constants define 'None', others don't - we're not particularly concerned
		// for this comparison however, so remove them to normalize those values
		if strings.EqualFold(k, "None") {
			continue
		}

		actual[strings.ToLower(k)] = strings.ToLower(v)
	}

	normalizedExpected := make(map[string]string, 0)
	for k, v := range expected {
		normalizedExpected[strings.ToLower(k)] = strings.ToLower(v)
	}

	return reflect.DeepEqual(actual, normalizedExpected)
}
