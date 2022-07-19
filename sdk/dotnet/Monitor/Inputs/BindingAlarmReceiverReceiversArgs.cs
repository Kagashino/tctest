// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Monitor.Inputs
{

    public sealed class BindingAlarmReceiverReceiversArgs : Pulumi.ResourceArgs
    {
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("notifyWays", required: true)]
        private InputList<string>? _notifyWays;
        public InputList<string> NotifyWays
        {
            get => _notifyWays ?? (_notifyWays = new InputList<string>());
            set => _notifyWays = value;
        }

        [Input("receiveLanguage")]
        public Input<string>? ReceiveLanguage { get; set; }

        [Input("receiverGroupLists")]
        private InputList<int>? _receiverGroupLists;
        public InputList<int> ReceiverGroupLists
        {
            get => _receiverGroupLists ?? (_receiverGroupLists = new InputList<int>());
            set => _receiverGroupLists = value;
        }

        [Input("receiverType", required: true)]
        public Input<string> ReceiverType { get; set; } = null!;

        [Input("receiverUserLists")]
        private InputList<int>? _receiverUserLists;
        public InputList<int> ReceiverUserLists
        {
            get => _receiverUserLists ?? (_receiverUserLists = new InputList<int>());
            set => _receiverUserLists = value;
        }

        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        public BindingAlarmReceiverReceiversArgs()
        {
        }
    }
}
