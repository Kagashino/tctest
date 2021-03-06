// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cdh.Outputs
{

    [OutputType]
    public sealed class InstancesCdhInstanceListResult
    {
        public readonly string AvailabilityZone;
        public readonly string CageId;
        public readonly string ChargeType;
        public readonly string CreateTime;
        public readonly ImmutableArray<string> CvmInstanceIds;
        public readonly string ExpiredTime;
        public readonly string HostId;
        public readonly string HostName;
        public readonly ImmutableArray<Outputs.InstancesCdhInstanceListHostResourceResult> HostResources;
        public readonly string HostState;
        public readonly string HostType;
        public readonly string PrepaidRenewFlag;
        public readonly int ProjectId;

        [OutputConstructor]
        private InstancesCdhInstanceListResult(
            string availabilityZone,

            string cageId,

            string chargeType,

            string createTime,

            ImmutableArray<string> cvmInstanceIds,

            string expiredTime,

            string hostId,

            string hostName,

            ImmutableArray<Outputs.InstancesCdhInstanceListHostResourceResult> hostResources,

            string hostState,

            string hostType,

            string prepaidRenewFlag,

            int projectId)
        {
            AvailabilityZone = availabilityZone;
            CageId = cageId;
            ChargeType = chargeType;
            CreateTime = createTime;
            CvmInstanceIds = cvmInstanceIds;
            ExpiredTime = expiredTime;
            HostId = hostId;
            HostName = hostName;
            HostResources = hostResources;
            HostState = hostState;
            HostType = hostType;
            PrepaidRenewFlag = prepaidRenewFlag;
            ProjectId = projectId;
        }
    }
}
