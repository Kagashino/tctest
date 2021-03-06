// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Address.Outputs
{

    [OutputType]
    public sealed class TemplateGroupsGroupListResult
    {
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<string> TemplateIds;

        [OutputConstructor]
        private TemplateGroupsGroupListResult(
            string id,

            string name,

            ImmutableArray<string> templateIds)
        {
            Id = id;
            Name = name;
            TemplateIds = templateIds;
        }
    }
}
