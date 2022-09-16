using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.Synapse.v2021_06_01.IntegrationRuntime;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum SsisObjectMetadataTypeConstant
{
    [Description("Environment")]
    Environment,

    [Description("Folder")]
    Folder,

    [Description("Package")]
    Package,

    [Description("Project")]
    Project,
}
