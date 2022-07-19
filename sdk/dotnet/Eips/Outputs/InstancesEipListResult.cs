// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Eips.Outputs
{

    [OutputType]
    public sealed class InstancesEipListResult
    {
        public readonly string CreateTime;
        public readonly string EipId;
        public readonly string EipName;
        public readonly string EipType;
        public readonly string EniId;
        public readonly string InstanceId;
        public readonly string PublicIp;
        public readonly string Status;
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private InstancesEipListResult(
            string createTime,

            string eipId,

            string eipName,

            string eipType,

            string eniId,

            string instanceId,

            string publicIp,

            string status,

            ImmutableDictionary<string, object> tags)
        {
            CreateTime = createTime;
            EipId = eipId;
            EipName = eipName;
            EipType = eipType;
            EniId = eniId;
            InstanceId = instanceId;
            PublicIp = publicIp;
            Status = status;
            Tags = tags;
        }
    }
}
