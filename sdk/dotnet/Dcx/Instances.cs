// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Dcx
{
    public static class Instances
    {
        public static Task<InstancesResult> InvokeAsync(InstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstancesResult>("tctest:Dcx/instances:Instances", args ?? new InstancesArgs(), options.WithDefaults());

        public static Output<InstancesResult> Invoke(InstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstancesResult>("tctest:Dcx/instances:Instances", args ?? new InstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesArgs : Pulumi.InvokeArgs
    {
        [Input("dcxId")]
        public string? DcxId { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public InstancesArgs()
        {
        }
    }

    public sealed class InstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dcxId")]
        public Input<string>? DcxId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public InstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstancesResult
    {
        public readonly string? DcxId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.InstancesInstanceListResult> InstanceLists;
        public readonly string? Name;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private InstancesResult(
            string? dcxId,

            string id,

            ImmutableArray<Outputs.InstancesInstanceListResult> instanceLists,

            string? name,

            string? resultOutputFile)
        {
            DcxId = dcxId;
            Id = id;
            InstanceLists = instanceLists;
            Name = name;
            ResultOutputFile = resultOutputFile;
        }
    }
}
