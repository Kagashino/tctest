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
    public sealed class AsScalingGroupAutoScalingGroupForwardBalancerId
    {
        public readonly string ListenerId;
        public readonly string LoadBalancerId;
        public readonly string? RuleId;
        public readonly ImmutableArray<Outputs.AsScalingGroupAutoScalingGroupForwardBalancerIdTargetAttribute> TargetAttributes;

        [OutputConstructor]
        private AsScalingGroupAutoScalingGroupForwardBalancerId(
            string listenerId,

            string loadBalancerId,

            string? ruleId,

            ImmutableArray<Outputs.AsScalingGroupAutoScalingGroupForwardBalancerIdTargetAttribute> targetAttributes)
        {
            ListenerId = listenerId;
            LoadBalancerId = loadBalancerId;
            RuleId = ruleId;
            TargetAttributes = targetAttributes;
        }
    }
}
