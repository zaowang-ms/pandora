using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.ContainerRegistry.v2023_06_01_preview.Tokens;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum TokenCertificateNameConstant
{
    [Description("certificate1")]
    CertificateOne,

    [Description("certificate2")]
    CertificateTwo,
}
