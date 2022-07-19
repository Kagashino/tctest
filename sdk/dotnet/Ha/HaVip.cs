// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Ha
{
    [TctestResourceType("tctest:Ha/haVip:HaVip")]
    public partial class HaVip : Pulumi.CustomResource
    {
        /// <summary>
        /// EIP that is associated.
        /// </summary>
        [Output("addressIp")]
        public Output<string> AddressIp { get; private set; } = null!;

        /// <summary>
        /// Create time of the HA VIP.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Instance ID that is associated.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Name of the HA VIP. The length of character is limited to 1-60.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network interface ID that is associated.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// State of the HA VIP. Valid value: `AVAILABLE`, `UNBIND`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Subnet ID.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after
        /// resource created automatically.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a HaVip resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HaVip(string name, HaVipArgs args, CustomResourceOptions? options = null)
            : base("tctest:Ha/haVip:HaVip", name, args ?? new HaVipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HaVip(string name, Input<string> id, HaVipState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Ha/haVip:HaVip", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HaVip resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HaVip Get(string name, Input<string> id, HaVipState? state = null, CustomResourceOptions? options = null)
        {
            return new HaVip(name, id, state, options);
        }
    }

    public sealed class HaVipArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the HA VIP. The length of character is limited to 1-60.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Subnet ID.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after
        /// resource created automatically.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public HaVipArgs()
        {
        }
    }

    public sealed class HaVipState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// EIP that is associated.
        /// </summary>
        [Input("addressIp")]
        public Input<string>? AddressIp { get; set; }

        /// <summary>
        /// Create time of the HA VIP.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Instance ID that is associated.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of the HA VIP. The length of character is limited to 1-60.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network interface ID that is associated.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// State of the HA VIP. Valid value: `AVAILABLE`, `UNBIND`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Subnet ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after
        /// resource created automatically.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public HaVipState()
        {
        }
    }
}
