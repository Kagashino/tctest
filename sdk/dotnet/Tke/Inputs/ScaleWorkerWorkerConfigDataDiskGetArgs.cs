// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Tke.Inputs
{

    public sealed class ScaleWorkerWorkerConfigDataDiskGetArgs : Pulumi.ResourceArgs
    {
        [Input("autoFormatAndMount")]
        public Input<bool>? AutoFormatAndMount { get; set; }

        [Input("diskPartition")]
        public Input<string>? DiskPartition { get; set; }

        [Input("diskSize")]
        public Input<int>? DiskSize { get; set; }

        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        [Input("fileSystem")]
        public Input<string>? FileSystem { get; set; }

        [Input("mountTarget")]
        public Input<string>? MountTarget { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        public ScaleWorkerWorkerConfigDataDiskGetArgs()
        {
        }
    }
}
