// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Sqlserver
{
    public static class PublishSubscribes
    {
        public static Task<PublishSubscribesResult> InvokeAsync(PublishSubscribesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<PublishSubscribesResult>("tctest:Sqlserver/publishSubscribes:PublishSubscribes", args ?? new PublishSubscribesArgs(), options.WithDefaults());

        public static Output<PublishSubscribesResult> Invoke(PublishSubscribesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<PublishSubscribesResult>("tctest:Sqlserver/publishSubscribes:PublishSubscribes", args ?? new PublishSubscribesInvokeArgs(), options.WithDefaults());
    }


    public sealed class PublishSubscribesArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("pubOrSubInstanceId")]
        public string? PubOrSubInstanceId { get; set; }

        [Input("pubOrSubInstanceIp")]
        public string? PubOrSubInstanceIp { get; set; }

        [Input("publishDatabase")]
        public string? PublishDatabase { get; set; }

        [Input("publishSubscribeId")]
        public int? PublishSubscribeId { get; set; }

        [Input("publishSubscribeName")]
        public string? PublishSubscribeName { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subscribeDatabase")]
        public string? SubscribeDatabase { get; set; }

        public PublishSubscribesArgs()
        {
        }
    }

    public sealed class PublishSubscribesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("pubOrSubInstanceId")]
        public Input<string>? PubOrSubInstanceId { get; set; }

        [Input("pubOrSubInstanceIp")]
        public Input<string>? PubOrSubInstanceIp { get; set; }

        [Input("publishDatabase")]
        public Input<string>? PublishDatabase { get; set; }

        [Input("publishSubscribeId")]
        public Input<int>? PublishSubscribeId { get; set; }

        [Input("publishSubscribeName")]
        public Input<string>? PublishSubscribeName { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subscribeDatabase")]
        public Input<string>? SubscribeDatabase { get; set; }

        public PublishSubscribesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class PublishSubscribesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? PubOrSubInstanceId;
        public readonly string? PubOrSubInstanceIp;
        public readonly string? PublishDatabase;
        public readonly int? PublishSubscribeId;
        public readonly ImmutableArray<Outputs.PublishSubscribesPublishSubscribeListResult> PublishSubscribeLists;
        public readonly string? PublishSubscribeName;
        public readonly string? ResultOutputFile;
        public readonly string? SubscribeDatabase;

        [OutputConstructor]
        private PublishSubscribesResult(
            string id,

            string instanceId,

            string? pubOrSubInstanceId,

            string? pubOrSubInstanceIp,

            string? publishDatabase,

            int? publishSubscribeId,

            ImmutableArray<Outputs.PublishSubscribesPublishSubscribeListResult> publishSubscribeLists,

            string? publishSubscribeName,

            string? resultOutputFile,

            string? subscribeDatabase)
        {
            Id = id;
            InstanceId = instanceId;
            PubOrSubInstanceId = pubOrSubInstanceId;
            PubOrSubInstanceIp = pubOrSubInstanceIp;
            PublishDatabase = publishDatabase;
            PublishSubscribeId = publishSubscribeId;
            PublishSubscribeLists = publishSubscribeLists;
            PublishSubscribeName = publishSubscribeName;
            ResultOutputFile = resultOutputFile;
            SubscribeDatabase = subscribeDatabase;
        }
    }
}
