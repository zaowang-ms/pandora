using System.Collections.Generic;
using Pandora.Definitions.Interfaces;


// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.


namespace Pandora.Definitions.ResourceManager.Cognitive.v2023_05_01.PrivateLinkResources;

internal class Definition : ResourceDefinition
{
    public string Name => "PrivateLinkResources";
    public IEnumerable<Interfaces.ApiOperation> Operations => new List<Interfaces.ApiOperation>
    {
        new ListOperation(),
    };
    public IEnumerable<System.Type> Constants => new List<System.Type>
    {

    };
    public IEnumerable<System.Type> Models => new List<System.Type>
    {
        typeof(PrivateLinkResourceModel),
        typeof(PrivateLinkResourceListResultModel),
        typeof(PrivateLinkResourcePropertiesModel),
    };
}
