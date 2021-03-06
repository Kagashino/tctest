// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Monitor.Outputs
{

    [OutputType]
    public sealed class PolicyConditionsListEventMetricResult
    {
        public readonly int EventId;
        public readonly string EventShowName;
        public readonly bool NeedRecovered;

        [OutputConstructor]
        private PolicyConditionsListEventMetricResult(
            int eventId,

            string eventShowName,

            bool needRecovered)
        {
            EventId = eventId;
            EventShowName = eventShowName;
            NeedRecovered = needRecovered;
        }
    }
}
