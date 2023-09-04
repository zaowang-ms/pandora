using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.ContainerRegistry.v2023_06_01_preview.Tokens;

[ConstantType(ConstantTypeAttribute.ConstantType.String)]
internal enum TokenPasswordNameConstant
{
    [Description("password1")]
    PasswordOne,

    [Description("password2")]
    PasswordTwo,
}
