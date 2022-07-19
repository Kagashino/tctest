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
    public sealed class BindingObjectsListResult
    {
        public readonly string DimensionsJson;
        public readonly int IsShielded;
        public readonly string Region;
        public readonly string UniqueId;

        [OutputConstructor]
        private BindingObjectsListResult(
            string dimensionsJson,

            int isShielded,

            string region,

            string uniqueId)
        {
            DimensionsJson = dimensionsJson;
            IsShielded = isShielded;
            Region = region;
            UniqueId = uniqueId;
        }
    }
}
