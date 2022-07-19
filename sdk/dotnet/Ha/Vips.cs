// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Ha
{
    public static class Vips
    {
        public static Task<VipsResult> InvokeAsync(VipsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<VipsResult>("tctest:Ha/vips:Vips", args ?? new VipsArgs(), options.WithDefaults());

        public static Output<VipsResult> Invoke(VipsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<VipsResult>("tctest:Ha/vips:Vips", args ?? new VipsInvokeArgs(), options.WithDefaults());
    }


    public sealed class VipsArgs : Pulumi.InvokeArgs
    {
        [Input("addressIp")]
        public string? AddressIp { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public string? SubnetId { get; set; }

        [Input("vpcId")]
        public string? VpcId { get; set; }

        public VipsArgs()
        {
        }
    }

    public sealed class VipsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("addressIp")]
        public Input<string>? AddressIp { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VipsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class VipsResult
    {
        public readonly string? AddressIp;
        public readonly ImmutableArray<Outputs.VipsHaVipListResult> HaVipLists;
        public readonly string? Id;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly string? SubnetId;
        public readonly string? VpcId;

        [OutputConstructor]
        private VipsResult(
            string? addressIp,

            ImmutableArray<Outputs.VipsHaVipListResult> haVipLists,

            string? id,

            string? name,

            string? resultOutputFile,

            string? subnetId,

            string? vpcId)
        {
            AddressIp = addressIp;
            HaVipLists = haVipLists;
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
