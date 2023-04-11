using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Security.v2017_08_01_preview.InformationProtectionPolicies;


internal class InformationProtectionPolicyPropertiesModel
{
    [JsonPropertyName("informationTypes")]
    public Dictionary<string, InformationTypeModel>? InformationTypes { get; set; }

    [JsonPropertyName("labels")]
    public Dictionary<string, SensitivityLabelModel>? Labels { get; set; }

    [DateFormat(DateFormatAttribute.DateFormat.RFC3339)]
    [JsonPropertyName("lastModifiedUtc")]
    public DateTime? LastModifiedUtc { get; set; }

    [JsonPropertyName("version")]
    public string? Version { get; set; }
}
