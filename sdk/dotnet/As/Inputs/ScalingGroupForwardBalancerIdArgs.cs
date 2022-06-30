// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As.Inputs
{

    public sealed class ScalingGroupForwardBalancerIdArgs : Pulumi.ResourceArgs
    {
        [Input("listenerId", required: true)]
        public Input<string> ListenerId { get; set; } = null!;

        [Input("loadBalancerId", required: true)]
        public Input<string> LoadBalancerId { get; set; } = null!;

        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        [Input("targetAttributes", required: true)]
        private InputList<Inputs.ScalingGroupForwardBalancerIdTargetAttributeArgs>? _targetAttributes;
        public InputList<Inputs.ScalingGroupForwardBalancerIdTargetAttributeArgs> TargetAttributes
        {
            get => _targetAttributes ?? (_targetAttributes = new InputList<Inputs.ScalingGroupForwardBalancerIdTargetAttributeArgs>());
            set => _targetAttributes = value;
        }

        public ScalingGroupForwardBalancerIdArgs()
        {
        }
    }
}