// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Elasticsearch.Inputs
{

    public sealed class InstanceWebNodeTypeInfoArgs : Pulumi.ResourceArgs
    {
        [Input("nodeNum", required: true)]
        public Input<int> NodeNum { get; set; } = null!;

        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        public InstanceWebNodeTypeInfoArgs()
        {
        }
    }
}
