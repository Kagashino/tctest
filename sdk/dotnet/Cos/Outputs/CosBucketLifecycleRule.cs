// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cos.Outputs
{

    [OutputType]
    public sealed class CosBucketLifecycleRule
    {
        public readonly Outputs.CosBucketLifecycleRuleExpiration? Expiration;
        public readonly string FilterPrefix;
        public readonly string? Id;
        public readonly Outputs.CosBucketLifecycleRuleNonCurrentExpiration? NonCurrentExpiration;
        public readonly ImmutableArray<Outputs.CosBucketLifecycleRuleNonCurrentTransition> NonCurrentTransitions;
        public readonly ImmutableArray<Outputs.CosBucketLifecycleRuleTransition> Transitions;

        [OutputConstructor]
        private CosBucketLifecycleRule(
            Outputs.CosBucketLifecycleRuleExpiration? expiration,

            string filterPrefix,

            string? id,

            Outputs.CosBucketLifecycleRuleNonCurrentExpiration? nonCurrentExpiration,

            ImmutableArray<Outputs.CosBucketLifecycleRuleNonCurrentTransition> nonCurrentTransitions,

            ImmutableArray<Outputs.CosBucketLifecycleRuleTransition> transitions)
        {
            Expiration = expiration;
            FilterPrefix = filterPrefix;
            Id = id;
            NonCurrentExpiration = nonCurrentExpiration;
            NonCurrentTransitions = nonCurrentTransitions;
            Transitions = transitions;
        }
    }
}
