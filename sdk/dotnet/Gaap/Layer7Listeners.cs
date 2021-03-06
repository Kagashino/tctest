// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Gaap
{
    public static class Layer7Listeners
    {
        public static Task<Layer7ListenersResult> InvokeAsync(Layer7ListenersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<Layer7ListenersResult>("tctest:Gaap/layer7Listeners:Layer7Listeners", args ?? new Layer7ListenersArgs(), options.WithDefaults());

        public static Output<Layer7ListenersResult> Invoke(Layer7ListenersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<Layer7ListenersResult>("tctest:Gaap/layer7Listeners:Layer7Listeners", args ?? new Layer7ListenersInvokeArgs(), options.WithDefaults());
    }


    public sealed class Layer7ListenersArgs : Pulumi.InvokeArgs
    {
        [Input("listenerId")]
        public string? ListenerId { get; set; }

        [Input("listenerName")]
        public string? ListenerName { get; set; }

        [Input("port")]
        public int? Port { get; set; }

        [Input("protocol", required: true)]
        public string Protocol { get; set; } = null!;

        [Input("proxyId")]
        public string? ProxyId { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public Layer7ListenersArgs()
        {
        }
    }

    public sealed class Layer7ListenersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("listenerId")]
        public Input<string>? ListenerId { get; set; }

        [Input("listenerName")]
        public Input<string>? ListenerName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("proxyId")]
        public Input<string>? ProxyId { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public Layer7ListenersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class Layer7ListenersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ListenerId;
        public readonly string? ListenerName;
        public readonly ImmutableArray<Outputs.Layer7ListenersListenerResult> Listeners;
        public readonly int? Port;
        public readonly string Protocol;
        public readonly string? ProxyId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private Layer7ListenersResult(
            string id,

            string? listenerId,

            string? listenerName,

            ImmutableArray<Outputs.Layer7ListenersListenerResult> listeners,

            int? port,

            string protocol,

            string? proxyId,

            string? resultOutputFile)
        {
            Id = id;
            ListenerId = listenerId;
            ListenerName = listenerName;
            Listeners = listeners;
            Port = port;
            Protocol = protocol;
            ProxyId = proxyId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
