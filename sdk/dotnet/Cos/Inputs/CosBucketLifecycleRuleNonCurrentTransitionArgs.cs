// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cos.Inputs
{

    public sealed class CosBucketLifecycleRuleNonCurrentTransitionArgs : Pulumi.ResourceArgs
    {
        [Input("nonCurrentDays")]
        public Input<int>? NonCurrentDays { get; set; }

        [Input("storageClass", required: true)]
        public Input<string> StorageClass { get; set; } = null!;

        public CosBucketLifecycleRuleNonCurrentTransitionArgs()
        {
        }
    }
}
