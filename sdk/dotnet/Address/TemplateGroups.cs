// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Address
{
    public static class TemplateGroups
    {
        public static Task<TemplateGroupsResult> InvokeAsync(TemplateGroupsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<TemplateGroupsResult>("tctest:Address/templateGroups:TemplateGroups", args ?? new TemplateGroupsArgs(), options.WithDefaults());

        public static Output<TemplateGroupsResult> Invoke(TemplateGroupsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<TemplateGroupsResult>("tctest:Address/templateGroups:TemplateGroups", args ?? new TemplateGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class TemplateGroupsArgs : Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public TemplateGroupsArgs()
        {
        }
    }

    public sealed class TemplateGroupsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public TemplateGroupsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class TemplateGroupsResult
    {
        public readonly ImmutableArray<Outputs.TemplateGroupsGroupListResult> GroupLists;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private TemplateGroupsResult(
            ImmutableArray<Outputs.TemplateGroupsGroupListResult> groupLists,

            string? id,

            string? name,

            string? resultOutputFile)
        {
            GroupLists = groupLists;
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}
