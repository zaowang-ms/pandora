using Pandora.Definitions.Attributes;
using System.ComponentModel;

namespace Pandora.Definitions.ResourceManager.FrontDoor.v2020_04_01.FrontDoors
{
    [ConstantType(ConstantTypeAttribute.ConstantType.String)]
    internal enum MatchProcessingBehaviorConstant
    {
        [Description("Continue")]
        Continue,

        [Description("Stop")]
        Stop,
    }
}
