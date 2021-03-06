// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Ckafka
{
    public static class Topics
    {
        public static Task<TopicsResult> InvokeAsync(TopicsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<TopicsResult>("tctest:Ckafka/topics:Topics", args ?? new TopicsArgs(), options.WithDefaults());

        public static Output<TopicsResult> Invoke(TopicsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<TopicsResult>("tctest:Ckafka/topics:Topics", args ?? new TopicsInvokeArgs(), options.WithDefaults());
    }


    public sealed class TopicsArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("topicName")]
        public string? TopicName { get; set; }

        public TopicsArgs()
        {
        }
    }

    public sealed class TopicsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("topicName")]
        public Input<string>? TopicName { get; set; }

        public TopicsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class TopicsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly ImmutableArray<Outputs.TopicsInstanceListResult> InstanceLists;
        public readonly string? ResultOutputFile;
        public readonly string? TopicName;

        [OutputConstructor]
        private TopicsResult(
            string id,

            string instanceId,

            ImmutableArray<Outputs.TopicsInstanceListResult> instanceLists,

            string? resultOutputFile,

            string? topicName)
        {
            Id = id;
            InstanceId = instanceId;
            InstanceLists = instanceLists;
            ResultOutputFile = resultOutputFile;
            TopicName = topicName;
        }
    }
}
