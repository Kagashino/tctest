// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vpc
{
    [TctestResourceType("tctest:Vpc/routeEntry:RouteEntry")]
    public partial class RouteEntry : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the routing table entry.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Destination address block.
        /// </summary>
        [Output("destinationCidrBlock")]
        public Output<string> DestinationCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Whether the entry is disabled, default is `false`.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        /// </summary>
        [Output("nextHub")]
        public Output<string> NextHub { get; private set; } = null!;

        /// <summary>
        /// Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `SSLVPN`, `NAT`, `NORMAL_CVM`, `EIP`
        /// and `CCN`.
        /// </summary>
        [Output("nextType")]
        public Output<string> NextType { get; private set; } = null!;

        /// <summary>
        /// ID of routing table to which this entry belongs.
        /// </summary>
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;


        /// <summary>
        /// Create a RouteEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteEntry(string name, RouteEntryArgs args, CustomResourceOptions? options = null)
            : base("tctest:Vpc/routeEntry:RouteEntry", name, args ?? new RouteEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteEntry(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Vpc/routeEntry:RouteEntry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RouteEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteEntry Get(string name, Input<string> id, RouteEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteEntry(name, id, state, options);
        }
    }

    public sealed class RouteEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the routing table entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination address block.
        /// </summary>
        [Input("destinationCidrBlock", required: true)]
        public Input<string> DestinationCidrBlock { get; set; } = null!;

        /// <summary>
        /// Whether the entry is disabled, default is `false`.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        /// </summary>
        [Input("nextHub", required: true)]
        public Input<string> NextHub { get; set; } = null!;

        /// <summary>
        /// Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `SSLVPN`, `NAT`, `NORMAL_CVM`, `EIP`
        /// and `CCN`.
        /// </summary>
        [Input("nextType", required: true)]
        public Input<string> NextType { get; set; } = null!;

        /// <summary>
        /// ID of routing table to which this entry belongs.
        /// </summary>
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        public RouteEntryArgs()
        {
        }
    }

    public sealed class RouteEntryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the routing table entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination address block.
        /// </summary>
        [Input("destinationCidrBlock")]
        public Input<string>? DestinationCidrBlock { get; set; }

        /// <summary>
        /// Whether the entry is disabled, default is `false`.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
        /// </summary>
        [Input("nextHub")]
        public Input<string>? NextHub { get; set; }

        /// <summary>
        /// Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `SSLVPN`, `NAT`, `NORMAL_CVM`, `EIP`
        /// and `CCN`.
        /// </summary>
        [Input("nextType")]
        public Input<string>? NextType { get; set; }

        /// <summary>
        /// ID of routing table to which this entry belongs.
        /// </summary>
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        public RouteEntryState()
        {
        }
    }
}
