// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class InstanceMultiZoneInfo
    {
        public readonly string AvailabilityZone;
        public readonly string SubnetId;

        [OutputConstructor]
        private InstanceMultiZoneInfo(
            string availabilityZone,

            string subnetId)
        {
            AvailabilityZone = availabilityZone;
            SubnetId = subnetId;
        }
    }
}
