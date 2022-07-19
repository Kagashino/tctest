// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cynosdb
{
    public static class Instances
    {
        public static Task<InstancesResult> InvokeAsync(InstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstancesResult>("tctest:Cynosdb/instances:Instances", args ?? new InstancesArgs(), options.WithDefaults());

        public static Output<InstancesResult> Invoke(InstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstancesResult>("tctest:Cynosdb/instances:Instances", args ?? new InstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        [Input("dbType")]
        public string? DbType { get; set; }

        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("instanceName")]
        public string? InstanceName { get; set; }

        [Input("projectId")]
        public int? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public InstancesArgs()
        {
        }
    }

    public sealed class InstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("dbType")]
        public Input<string>? DbType { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public InstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstancesResult
    {
        public readonly string? ClusterId;
        public readonly string? DbType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly ImmutableArray<Outputs.InstancesInstanceListResult> InstanceLists;
        public readonly string? InstanceName;
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private InstancesResult(
            string? clusterId,

            string? dbType,

            string id,

            string? instanceId,

            ImmutableArray<Outputs.InstancesInstanceListResult> instanceLists,

            string? instanceName,

            int? projectId,

            string? resultOutputFile)
        {
            ClusterId = clusterId;
            DbType = dbType;
            Id = id;
            InstanceId = instanceId;
            InstanceLists = instanceLists;
            InstanceName = instanceName;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
