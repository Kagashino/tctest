// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Tke.Outputs
{

    [OutputType]
    public sealed class ClusterAttachmentWorkerConfigOverridesDataDisk
    {
        public readonly bool? AutoFormatAndMount;
        public readonly string? DiskPartition;
        public readonly int? DiskSize;
        public readonly string? DiskType;
        public readonly string? FileSystem;
        public readonly string? MountTarget;

        [OutputConstructor]
        private ClusterAttachmentWorkerConfigOverridesDataDisk(
            bool? autoFormatAndMount,

            string? diskPartition,

            int? diskSize,

            string? diskType,

            string? fileSystem,

            string? mountTarget)
        {
            AutoFormatAndMount = autoFormatAndMount;
            DiskPartition = diskPartition;
            DiskSize = diskSize;
            DiskType = diskType;
            FileSystem = fileSystem;
            MountTarget = mountTarget;
        }
    }
}
