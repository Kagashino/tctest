// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Mongodb.Outputs
{

    [OutputType]
    public sealed class ZoneConfigListResult
    {
        public readonly string AvailableZone;
        public readonly string ClusterType;
        public readonly int Cpu;
        public readonly int DefaultStorage;
        public readonly string EngineVersion;
        public readonly string MachineType;
        public readonly int MaxStorage;
        public readonly int Memory;
        public readonly int MinStorage;

        [OutputConstructor]
        private ZoneConfigListResult(
            string availableZone,

            string clusterType,

            int cpu,

            int defaultStorage,

            string engineVersion,

            string machineType,

            int maxStorage,

            int memory,

            int minStorage)
        {
            AvailableZone = availableZone;
            ClusterType = clusterType;
            Cpu = cpu;
            DefaultStorage = defaultStorage;
            EngineVersion = engineVersion;
            MachineType = machineType;
            MaxStorage = maxStorage;
            Memory = memory;
            MinStorage = minStorage;
        }
    }
}
