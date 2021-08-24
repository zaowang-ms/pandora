package main

import (
	"fmt"

	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/models"
	"github.com/hashicorp/pandora/tools/importer-rest-api-specs/parser"
)

func parseSwaggerFiles(input RunInput, debug bool) (*[]parsedData, error) {
	parsed := make(map[string]parsedData, 0)
	for _, file := range input.SwaggerFiles {
		swaggerFile, err := parser.Load(input.SwaggerDirectory, file, debug)
		if err != nil {
			return nil, fmt.Errorf("parsing file %q: %+v", file, err)
		}

		definition, err := swaggerFile.Parse(input.ServiceName, input.ApiVersion)
		if err != nil {
			return nil, fmt.Errorf("parsing definition: %+v", err)
		}

		data := parsedData{
			ServiceName: input.ServiceName,
			ApiVersion:  input.ApiVersion,
			Resources:   definition.Resources,
		}

		if existing, ok := parsed[data.Key()]; ok {
			// it's possible for Swagger tags to exist in multiple files, as EventHubs has DeleteAuthorizationRule which
			// lives in the AuthorizationRule json, but is technically part of the EventHubs namespace - as such we need
			// to combine the items rather than overwriting the key
			resources, err := data.combineResourcesWith(existing.Resources)
			if err != nil {
				return nil, fmt.Errorf("combining resources for %q: %+v", existing.Key(), err)
			}
			data.Resources = *resources
		}

		parsed[data.Key()] = data
	}

	out := make([]parsedData, 0)
	for _, v := range parsed {
		out = append(out, v)
	}
	return &out, nil
}

type parsedData struct {
	ServiceName string
	ApiVersion  string
	Resources   map[string]models.AzureApiResource
}

func (d parsedData) Key() string {
	return fmt.Sprintf("%s-%s", d.ServiceName, d.ApiVersion)
}

func (d parsedData) combineResourcesWith(other map[string]models.AzureApiResource) (*map[string]models.AzureApiResource, error) {
	resources := make(map[string]models.AzureApiResource)
	for k, v := range d.Resources {
		resources[k] = v
	}

	for k, v := range other {
		existing, ok := resources[k]
		if !ok {
			resources[k] = v
			continue
		}

		constants, err := combineConstants(existing.Constants, v.Constants)
		if err != nil {
			return nil, fmt.Errorf("combining constants: %+v", err)
		}
		existing.Constants = *constants

		models, err := combineModels(existing.Models, v.Models)
		if err != nil {
			return nil, fmt.Errorf("combining models: %+v", err)
		}
		existing.Models = *models

		operations, err := combineOperations(existing.Operations, v.Operations)
		if err != nil {
			return nil, fmt.Errorf("combining operations: %+v", err)
		}
		existing.Operations = *operations

		resourceIds, err := combineMaps(existing.ResourceIds, v.ResourceIds)
		if err != nil {
			return nil, fmt.Errorf("combining resource ids: %+v", err)
		}
		existing.ResourceIds = *resourceIds

		resources[k] = existing
	}

	return &resources, nil
}

func combineConstants(first map[string]models.ConstantDetails, second map[string]models.ConstantDetails) (*map[string]models.ConstantDetails, error) {
	constants := make(map[string]models.ConstantDetails, 0)
	for k, v := range first {
		constants[k] = v
	}

	for k, v := range second {
		existingConst, ok := constants[k]
		if !ok {
			constants[k] = v
			continue
		}

		if existingConst.FieldType != v.FieldType {
			return nil, fmt.Errorf("combining constant %q - multiple field types defined as %q and %q", k, string(existingConst.FieldType), string(v.FieldType))
		}

		vals, err := combineMaps(existingConst.Values, v.Values)
		if err != nil {
			return nil, fmt.Errorf("combining maps: %+v", err)
		}
		existingConst.Values = *vals
		constants[k] = existingConst
	}

	return &constants, nil
}

func combineModels(first map[string]models.ModelDetails, second map[string]models.ModelDetails) (*map[string]models.ModelDetails, error) {
	output := make(map[string]models.ModelDetails)

	for k, v := range first {
		output[k] = v
	}

	for k, v := range second {
		existing, ok := output[k]
		if ok && len(existing.Fields) != len(v.Fields) {
			return nil, fmt.Errorf("duplicate models named %q with different fields - first %d - second %d", k, len(existing.Fields), len(v.Fields))
		}

		output[k] = v
	}

	return &output, nil
}

func combineOperations(first map[string]models.OperationDetails, second map[string]models.OperationDetails) (*map[string]models.OperationDetails, error) {
	output := make(map[string]models.OperationDetails, 0)

	for k, v := range first {
		output[k] = v
	}

	for k, v := range second {
		// if there's duplicate operations named the same thing in different Swaggers, this is likely a data issue
		_, ok := output[k]
		if ok {
			return nil, fmt.Errorf("duplicate operations named %q", k)
		}

		output[k] = v
	}

	return &output, nil
}

func combineMaps(first map[string]string, second map[string]string) (*map[string]string, error) {
	vals := make(map[string]string, 0)
	for k, v := range first {
		vals[k] = v
	}
	for k, v := range second {
		existing, ok := vals[k]
		if !ok {
			vals[k] = v
			continue
		}

		if existing != v {
			return nil, fmt.Errorf("duplicate key %q with different value - first %q - second %q", k, existing, v)
		}
	}

	return &vals, nil
}
