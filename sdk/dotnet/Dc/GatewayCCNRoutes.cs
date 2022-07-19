// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Dc
{
    public static class GatewayCCNRoutes
    {
        public static Task<GatewayCCNRoutesResult> InvokeAsync(GatewayCCNRoutesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GatewayCCNRoutesResult>("tctest:Dc/gatewayCCNRoutes:GatewayCCNRoutes", args ?? new GatewayCCNRoutesArgs(), options.WithDefaults());

        public static Output<GatewayCCNRoutesResult> Invoke(GatewayCCNRoutesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GatewayCCNRoutesResult>("tctest:Dc/gatewayCCNRoutes:GatewayCCNRoutes", args ?? new GatewayCCNRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GatewayCCNRoutesArgs : Pulumi.InvokeArgs
    {
        [Input("dcgId", required: true)]
        public string DcgId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GatewayCCNRoutesArgs()
        {
        }
    }

    public sealed class GatewayCCNRoutesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dcgId", required: true)]
        public Input<string> DcgId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GatewayCCNRoutesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GatewayCCNRoutesResult
    {
        public readonly string DcgId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GatewayCCNRoutesInstanceListResult> InstanceLists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GatewayCCNRoutesResult(
            string dcgId,

            string id,

            ImmutableArray<Outputs.GatewayCCNRoutesInstanceListResult> instanceLists,

            string? resultOutputFile)
        {
            DcgId = dcgId;
            Id = id;
            InstanceLists = instanceLists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
