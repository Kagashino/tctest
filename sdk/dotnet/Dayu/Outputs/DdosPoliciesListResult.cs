// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Dayu.Outputs
{

    [OutputType]
    public sealed class DdosPoliciesListResult
    {
        public readonly ImmutableArray<string> BlackIps;
        public readonly string CreateTime;
        public readonly ImmutableArray<Outputs.DdosPoliciesListDropOptionResult> DropOptions;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.DdosPoliciesListPacketFilterResult> PacketFilters;
        public readonly string PolicyId;
        public readonly ImmutableArray<Outputs.DdosPoliciesListPortFilterResult> PortFilters;
        public readonly string SceneId;
        public readonly ImmutableArray<Outputs.DdosPoliciesListWatermarkFilterResult> WatermarkFilters;
        public readonly ImmutableArray<Outputs.DdosPoliciesListWatermarkKeyResult> WatermarkKeys;
        public readonly ImmutableArray<string> WhiteIps;

        [OutputConstructor]
        private DdosPoliciesListResult(
            ImmutableArray<string> blackIps,

            string createTime,

            ImmutableArray<Outputs.DdosPoliciesListDropOptionResult> dropOptions,

            string name,

            ImmutableArray<Outputs.DdosPoliciesListPacketFilterResult> packetFilters,

            string policyId,

            ImmutableArray<Outputs.DdosPoliciesListPortFilterResult> portFilters,

            string sceneId,

            ImmutableArray<Outputs.DdosPoliciesListWatermarkFilterResult> watermarkFilters,

            ImmutableArray<Outputs.DdosPoliciesListWatermarkKeyResult> watermarkKeys,

            ImmutableArray<string> whiteIps)
        {
            BlackIps = blackIps;
            CreateTime = createTime;
            DropOptions = dropOptions;
            Name = name;
            PacketFilters = packetFilters;
            PolicyId = policyId;
            PortFilters = portFilters;
            SceneId = sceneId;
            WatermarkFilters = watermarkFilters;
            WatermarkKeys = watermarkKeys;
            WhiteIps = whiteIps;
        }
    }
}
