// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Tke.Inputs
{

    public sealed class ClusterExistInstanceArgs : Pulumi.ResourceArgs
    {
        [Input("desiredPodNumbers")]
        private InputList<int>? _desiredPodNumbers;
        public InputList<int> DesiredPodNumbers
        {
            get => _desiredPodNumbers ?? (_desiredPodNumbers = new InputList<int>());
            set => _desiredPodNumbers = value;
        }

        [Input("instancesPara")]
        public Input<Inputs.ClusterExistInstanceInstancesParaArgs>? InstancesPara { get; set; }

        [Input("nodeRole")]
        public Input<string>? NodeRole { get; set; }

        public ClusterExistInstanceArgs()
        {
        }
    }
}
