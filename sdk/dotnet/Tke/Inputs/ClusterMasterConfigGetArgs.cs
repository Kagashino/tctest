// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Tke.Inputs
{

    public sealed class ClusterMasterConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        [Input("camRoleName")]
        public Input<string>? CamRoleName { get; set; }

        [Input("count")]
        public Input<int>? Count { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.ClusterMasterConfigDataDiskGetArgs>? _dataDisks;
        public InputList<Inputs.ClusterMasterConfigDataDiskGetArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.ClusterMasterConfigDataDiskGetArgs>());
            set => _dataDisks = value;
        }

        [Input("desiredPodNum")]
        public Input<int>? DesiredPodNum { get; set; }

        [Input("disasterRecoverGroupIds")]
        public Input<string>? DisasterRecoverGroupIds { get; set; }

        [Input("enhancedMonitorService")]
        public Input<bool>? EnhancedMonitorService { get; set; }

        [Input("enhancedSecurityService")]
        public Input<bool>? EnhancedSecurityService { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("imgId")]
        public Input<string>? ImgId { get; set; }

        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        [Input("instanceChargeTypePrepaidPeriod")]
        public Input<int>? InstanceChargeTypePrepaidPeriod { get; set; }

        [Input("instanceChargeTypePrepaidRenewFlag")]
        public Input<string>? InstanceChargeTypePrepaidRenewFlag { get; set; }

        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        [Input("keyIds")]
        public Input<string>? KeyIds { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("publicIpAssigned")]
        public Input<bool>? PublicIpAssigned { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("systemDiskSize")]
        public Input<int>? SystemDiskSize { get; set; }

        [Input("systemDiskType")]
        public Input<string>? SystemDiskType { get; set; }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public ClusterMasterConfigGetArgs()
        {
        }
    }
}
