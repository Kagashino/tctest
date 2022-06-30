// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tke.Outputs
{

    [OutputType]
    public sealed class NodePoolAutoScalingConfigDataDisk
    {
        public readonly bool? DeleteWithInstance;
        public readonly int? DiskSize;
        public readonly string? DiskType;
        public readonly string? SnapshotId;

        [OutputConstructor]
        private NodePoolAutoScalingConfigDataDisk(
            bool? deleteWithInstance,

            int? diskSize,

            string? diskType,

            string? snapshotId)
        {
            DeleteWithInstance = deleteWithInstance;
            DiskSize = diskSize;
            DiskType = diskType;
            SnapshotId = snapshotId;
        }
    }
}