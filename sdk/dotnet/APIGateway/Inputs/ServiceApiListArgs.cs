// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.APIGateway.Inputs
{

    public sealed class ServiceApiListArgs : Pulumi.ResourceArgs
    {
        [Input("apiDesc")]
        public Input<string>? ApiDesc { get; set; }

        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        [Input("apiName")]
        public Input<string>? ApiName { get; set; }

        [Input("method")]
        public Input<string>? Method { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        public ServiceApiListArgs()
        {
        }
    }
}
