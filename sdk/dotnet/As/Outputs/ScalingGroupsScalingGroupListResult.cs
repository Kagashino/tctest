// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.As.Outputs
{

    [OutputType]
    public sealed class ScalingGroupsScalingGroupListResult
    {
        public readonly string ConfigurationId;
        public readonly string CreateTime;
        public readonly int DefaultCooldown;
        public readonly int DesiredCapacity;
        public readonly ImmutableArray<Outputs.ScalingGroupsScalingGroupListForwardBalancerIdResult> ForwardBalancerIds;
        public readonly int InstanceCount;
        public readonly ImmutableArray<string> LoadBalancerIds;
        public readonly int MaxSize;
        public readonly int MinSize;
        public readonly string MultiZoneSubnetPolicy;
        public readonly int ProjectId;
        public readonly string RetryPolicy;
        public readonly string ScalingGroupId;
        public readonly string ScalingGroupName;
        public readonly string Status;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly ImmutableArray<string> TerminationPolicies;
        public readonly string VpcId;
        public readonly ImmutableArray<string> Zones;

        [OutputConstructor]
        private ScalingGroupsScalingGroupListResult(
            string configurationId,

            string createTime,

            int defaultCooldown,

            int desiredCapacity,

            ImmutableArray<Outputs.ScalingGroupsScalingGroupListForwardBalancerIdResult> forwardBalancerIds,

            int instanceCount,

            ImmutableArray<string> loadBalancerIds,

            int maxSize,

            int minSize,

            string multiZoneSubnetPolicy,

            int projectId,

            string retryPolicy,

            string scalingGroupId,

            string scalingGroupName,

            string status,

            ImmutableArray<string> subnetIds,

            ImmutableDictionary<string, object> tags,

            ImmutableArray<string> terminationPolicies,

            string vpcId,

            ImmutableArray<string> zones)
        {
            ConfigurationId = configurationId;
            CreateTime = createTime;
            DefaultCooldown = defaultCooldown;
            DesiredCapacity = desiredCapacity;
            ForwardBalancerIds = forwardBalancerIds;
            InstanceCount = instanceCount;
            LoadBalancerIds = loadBalancerIds;
            MaxSize = maxSize;
            MinSize = minSize;
            MultiZoneSubnetPolicy = multiZoneSubnetPolicy;
            ProjectId = projectId;
            RetryPolicy = retryPolicy;
            ScalingGroupId = scalingGroupId;
            ScalingGroupName = scalingGroupName;
            Status = status;
            SubnetIds = subnetIds;
            Tags = tags;
            TerminationPolicies = terminationPolicies;
            VpcId = vpcId;
            Zones = zones;
        }
    }
}
