// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Sqlserver.Outputs
{

    [OutputType]
    public sealed class ReadonlyGroupsListResult
    {
        public readonly string Id;
        public readonly int IsOfflineDelay;
        public readonly string MasterInstanceId;
        public readonly int MaxDelayTime;
        public readonly int MinInstances;
        public readonly string Name;
        public readonly ImmutableArray<string> ReadonlyInstanceSets;
        public readonly int Status;
        public readonly string Vip;
        public readonly int Vport;

        [OutputConstructor]
        private ReadonlyGroupsListResult(
            string id,

            int isOfflineDelay,

            string masterInstanceId,

            int maxDelayTime,

            int minInstances,

            string name,

            ImmutableArray<string> readonlyInstanceSets,

            int status,

            string vip,

            int vport)
        {
            Id = id;
            IsOfflineDelay = isOfflineDelay;
            MasterInstanceId = masterInstanceId;
            MaxDelayTime = maxDelayTime;
            MinInstances = minInstances;
            Name = name;
            ReadonlyInstanceSets = readonlyInstanceSets;
            Status = status;
            Vip = vip;
            Vport = vport;
        }
    }
}
