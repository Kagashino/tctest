// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cloud
{
    public static class APIGatewayServices
    {
        public static Task<APIGatewayServicesResult> InvokeAsync(APIGatewayServicesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<APIGatewayServicesResult>("tctest:Cloud/aPIGatewayServices:APIGatewayServices", args ?? new APIGatewayServicesArgs(), options.WithDefaults());

        public static Output<APIGatewayServicesResult> Invoke(APIGatewayServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<APIGatewayServicesResult>("tctest:Cloud/aPIGatewayServices:APIGatewayServices", args ?? new APIGatewayServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class APIGatewayServicesArgs : Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("serviceId")]
        public string? ServiceId { get; set; }

        [Input("serviceName")]
        public string? ServiceName { get; set; }

        public APIGatewayServicesArgs()
        {
        }
    }

    public sealed class APIGatewayServicesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public APIGatewayServicesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class APIGatewayServicesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.APIGatewayServicesListResult> Lists;
        public readonly string? ResultOutputFile;
        public readonly string? ServiceId;
        public readonly string? ServiceName;

        [OutputConstructor]
        private APIGatewayServicesResult(
            string id,

            ImmutableArray<Outputs.APIGatewayServicesListResult> lists,

            string? resultOutputFile,

            string? serviceId,

            string? serviceName)
        {
            Id = id;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            ServiceId = serviceId;
            ServiceName = serviceName;
        }
    }
}
