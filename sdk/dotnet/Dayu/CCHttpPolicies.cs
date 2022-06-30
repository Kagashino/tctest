// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu
{
    public static class CCHttpPolicies
    {
        public static Task<CCHttpPoliciesResult> InvokeAsync(CCHttpPoliciesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<CCHttpPoliciesResult>("tencentcloud:Dayu/cCHttpPolicies:CCHttpPolicies", args ?? new CCHttpPoliciesArgs(), options.WithDefaults());

        public static Output<CCHttpPoliciesResult> Invoke(CCHttpPoliciesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<CCHttpPoliciesResult>("tencentcloud:Dayu/cCHttpPolicies:CCHttpPolicies", args ?? new CCHttpPoliciesInvokeArgs(), options.WithDefaults());
    }


    public sealed class CCHttpPoliciesArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("policyId")]
        public string? PolicyId { get; set; }

        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        [Input("resourceType", required: true)]
        public string ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public CCHttpPoliciesArgs()
        {
        }
    }

    public sealed class CCHttpPoliciesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public CCHttpPoliciesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class CCHttpPoliciesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.CCHttpPoliciesListResult> Lists;
        public readonly string? Name;
        public readonly string? PolicyId;
        public readonly string ResourceId;
        public readonly string ResourceType;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private CCHttpPoliciesResult(
            string id,

            ImmutableArray<Outputs.CCHttpPoliciesListResult> lists,

            string? name,

            string? policyId,

            string resourceId,

            string resourceType,

            string? resultOutputFile)
        {
            Id = id;
            Lists = lists;
            Name = name;
            PolicyId = policyId;
            ResourceId = resourceId;
            ResourceType = resourceType;
            ResultOutputFile = resultOutputFile;
        }
    }
}