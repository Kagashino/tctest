// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Gaap.Inputs
{

    public sealed class Layer4ListenerRealserverBindSetGetArgs : Pulumi.ResourceArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public Layer4ListenerRealserverBindSetGetArgs()
        {
        }
    }
}
