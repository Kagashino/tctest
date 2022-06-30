// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Container.Outputs
{

    [OutputType]
    public sealed class ClusterInstancesNodeResult
    {
        public readonly string AbnormalReason;
        public readonly int Cpu;
        public readonly string InstanceId;
        public readonly int IsNormal;
        public readonly string LanIp;
        public readonly int Mem;
        public readonly string WanIp;

        [OutputConstructor]
        private ClusterInstancesNodeResult(
            string abnormalReason,

            int cpu,

            string instanceId,

            int isNormal,

            string lanIp,

            int mem,

            string wanIp)
        {
            AbnormalReason = abnormalReason;
            Cpu = cpu;
            InstanceId = instanceId;
            IsNormal = isNormal;
            LanIp = lanIp;
            Mem = mem;
            WanIp = wanIp;
        }
    }
}