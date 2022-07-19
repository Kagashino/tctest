// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vpn
{
    public static class GatewayRoutes
    {
        public static Task<GatewayRoutesResult> InvokeAsync(GatewayRoutesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GatewayRoutesResult>("tctest:Vpn/gatewayRoutes:GatewayRoutes", args ?? new GatewayRoutesArgs(), options.WithDefaults());

        public static Output<GatewayRoutesResult> Invoke(GatewayRoutesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GatewayRoutesResult>("tctest:Vpn/gatewayRoutes:GatewayRoutes", args ?? new GatewayRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GatewayRoutesArgs : Pulumi.InvokeArgs
    {
        [Input("destinationCidr")]
        public string? DestinationCidr { get; set; }

        [Input("instanceId")]
        public string? InstanceId { get; set; }

        [Input("instanceType")]
        public string? InstanceType { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("vpnGatewayId", required: true)]
        public string VpnGatewayId { get; set; } = null!;

        public GatewayRoutesArgs()
        {
        }
    }

    public sealed class GatewayRoutesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("destinationCidr")]
        public Input<string>? DestinationCidr { get; set; }

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("vpnGatewayId", required: true)]
        public Input<string> VpnGatewayId { get; set; } = null!;

        public GatewayRoutesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GatewayRoutesResult
    {
        public readonly string? DestinationCidr;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly string? InstanceType;
        public readonly string? ResultOutputFile;
        public readonly string VpnGatewayId;
        public readonly ImmutableArray<Outputs.GatewayRoutesVpnGatewayRouteListResult> VpnGatewayRouteLists;

        [OutputConstructor]
        private GatewayRoutesResult(
            string? destinationCidr,

            string id,

            string? instanceId,

            string? instanceType,

            string? resultOutputFile,

            string vpnGatewayId,

            ImmutableArray<Outputs.GatewayRoutesVpnGatewayRouteListResult> vpnGatewayRouteLists)
        {
            DestinationCidr = destinationCidr;
            Id = id;
            InstanceId = instanceId;
            InstanceType = instanceType;
            ResultOutputFile = resultOutputFile;
            VpnGatewayId = vpnGatewayId;
            VpnGatewayRouteLists = vpnGatewayRouteLists;
        }
    }
}
