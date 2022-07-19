// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Dayu.Inputs
{

    public sealed class CCPolicyV2CcPrecisionReqLimitGetArgs : Pulumi.ResourceArgs
    {
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("level", required: true)]
        public Input<string> Level { get; set; } = null!;

        [Input("policys", required: true)]
        private InputList<Inputs.CCPolicyV2CcPrecisionReqLimitPolicyGetArgs>? _policys;
        public InputList<Inputs.CCPolicyV2CcPrecisionReqLimitPolicyGetArgs> Policys
        {
            get => _policys ?? (_policys = new InputList<Inputs.CCPolicyV2CcPrecisionReqLimitPolicyGetArgs>());
            set => _policys = value;
        }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public CCPolicyV2CcPrecisionReqLimitGetArgs()
        {
        }
    }
}
