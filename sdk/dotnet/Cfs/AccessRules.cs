// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cfs
{
    public static class AccessRules
    {
        public static Task<AccessRulesResult> InvokeAsync(AccessRulesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<AccessRulesResult>("tencentcloud:Cfs/accessRules:AccessRules", args ?? new AccessRulesArgs(), options.WithDefaults());

        public static Output<AccessRulesResult> Invoke(AccessRulesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<AccessRulesResult>("tencentcloud:Cfs/accessRules:AccessRules", args ?? new AccessRulesInvokeArgs(), options.WithDefaults());
    }


    public sealed class AccessRulesArgs : Pulumi.InvokeArgs
    {
        [Input("accessGroupId", required: true)]
        public string AccessGroupId { get; set; } = null!;

        [Input("accessRuleId")]
        public string? AccessRuleId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public AccessRulesArgs()
        {
        }
    }

    public sealed class AccessRulesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("accessGroupId", required: true)]
        public Input<string> AccessGroupId { get; set; } = null!;

        [Input("accessRuleId")]
        public Input<string>? AccessRuleId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public AccessRulesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class AccessRulesResult
    {
        public readonly string AccessGroupId;
        public readonly string? AccessRuleId;
        public readonly ImmutableArray<Outputs.AccessRulesAccessRuleListResult> AccessRuleLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private AccessRulesResult(
            string accessGroupId,

            string? accessRuleId,

            ImmutableArray<Outputs.AccessRulesAccessRuleListResult> accessRuleLists,

            string id,

            string? resultOutputFile)
        {
            AccessGroupId = accessGroupId;
            AccessRuleId = accessRuleId;
            AccessRuleLists = accessRuleLists;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}