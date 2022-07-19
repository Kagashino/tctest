// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cvm.Outputs
{

    [OutputType]
    public sealed class InstanceDataDisk
    {
        public readonly string? DataDiskId;
        public readonly int DataDiskSize;
        public readonly string? DataDiskSnapshotId;
        public readonly string DataDiskType;
        public readonly bool? DeleteWithInstance;
        public readonly bool? Encrypt;
        public readonly int? ThroughputPerformance;

        [OutputConstructor]
        private InstanceDataDisk(
            string? dataDiskId,

            int dataDiskSize,

            string? dataDiskSnapshotId,

            string dataDiskType,

            bool? deleteWithInstance,

            bool? encrypt,

            int? throughputPerformance)
        {
            DataDiskId = dataDiskId;
            DataDiskSize = dataDiskSize;
            DataDiskSnapshotId = dataDiskSnapshotId;
            DataDiskType = dataDiskType;
            DeleteWithInstance = deleteWithInstance;
            Encrypt = encrypt;
            ThroughputPerformance = throughputPerformance;
        }
    }
}
