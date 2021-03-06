// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Clb
{
    public static class Listeners
    {
        public static Task<ListenersResult> InvokeAsync(ListenersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListenersResult>("tctest:Clb/listeners:Listeners", args ?? new ListenersArgs(), options.WithDefaults());

        public static Output<ListenersResult> Invoke(ListenersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ListenersResult>("tctest:Clb/listeners:Listeners", args ?? new ListenersInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListenersArgs : Pulumi.InvokeArgs
    {
        [Input("clbId", required: true)]
        public string ClbId { get; set; } = null!;

        [Input("listenerId")]
        public string? ListenerId { get; set; }

        [Input("port")]
        public int? Port { get; set; }

        [Input("protocol")]
        public string? Protocol { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public ListenersArgs()
        {
        }
    }

    public sealed class ListenersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("clbId", required: true)]
        public Input<string> ClbId { get; set; } = null!;

        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public ListenersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListenersResult
    {
        public readonly string ClbId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ListenerId;
        public readonly ImmutableArray<Outputs.ListenersListenerListResult> ListenerLists;
        public readonly int? Port;
        public readonly string? Protocol;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private ListenersResult(
            string clbId,

            string id,

            string? listenerId,

            ImmutableArray<Outputs.ListenersListenerListResult> listenerLists,

            int? port,

            string? protocol,

            string? resultOutputFile)
        {
            ClbId = clbId;
            Id = id;
            ListenerId = listenerId;
            ListenerLists = listenerLists;
            Port = port;
            Protocol = protocol;
            ResultOutputFile = resultOutputFile;
        }
    }
}
