using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;
using Pandora.Definitions.Attributes;
using Pandora.Definitions.Attributes.Validation;
using Pandora.Definitions.CustomTypes;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Synapse.v2021_06_01.RecoverableSqlPools;


internal class RecoverableSqlPoolPropertiesModel
{
    [JsonPropertyName("edition")]
    public string? Edition { get; set; }

    [JsonPropertyName("elasticPoolName")]
    public string? ElasticPoolName { get; set; }

    [DateFormat(DateFormatAttribute.DateFormat.RFC3339)]
    [JsonPropertyName("lastAvailableBackupDate")]
    public DateTime? LastAvailableBackupDate { get; set; }

    [JsonPropertyName("serviceLevelObjective")]
    public string? ServiceLevelObjective { get; set; }
}
