// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cvm
{
    [TctestResourceType("tctest:Cvm/reservedInstance:ReservedInstance")]
    public partial class ReservedInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration ID of the reserved instance.
        /// </summary>
        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Expiry time of the RI.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Number of reserved instances to be purchased.
        /// </summary>
        [Output("instanceCount")]
        public Output<int> InstanceCount { get; private set; } = null!;

        /// <summary>
        /// Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
        /// Up to 60 characters (including pattern strings) are supported.
        /// </summary>
        [Output("reservedInstanceName")]
        public Output<string?> ReservedInstanceName { get; private set; } = null!;

        /// <summary>
        /// Start time of the RI.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// Status of the RI at the time of purchase.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ReservedInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReservedInstance(string name, ReservedInstanceArgs args, CustomResourceOptions? options = null)
            : base("tctest:Cvm/reservedInstance:ReservedInstance", name, args ?? new ReservedInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReservedInstance(string name, Input<string> id, ReservedInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Cvm/reservedInstance:ReservedInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReservedInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReservedInstance Get(string name, Input<string> id, ReservedInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ReservedInstance(name, id, state, options);
        }
    }

    public sealed class ReservedInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration ID of the reserved instance.
        /// </summary>
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// Number of reserved instances to be purchased.
        /// </summary>
        [Input("instanceCount", required: true)]
        public Input<int> InstanceCount { get; set; } = null!;

        /// <summary>
        /// Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
        /// Up to 60 characters (including pattern strings) are supported.
        /// </summary>
        [Input("reservedInstanceName")]
        public Input<string>? ReservedInstanceName { get; set; }

        public ReservedInstanceArgs()
        {
        }
    }

    public sealed class ReservedInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration ID of the reserved instance.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// Expiry time of the RI.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Number of reserved instances to be purchased.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// Reserved Instance display name. - If you do not specify an instance display name, 'Unnamed' is displayed by default. -
        /// Up to 60 characters (including pattern strings) are supported.
        /// </summary>
        [Input("reservedInstanceName")]
        public Input<string>? ReservedInstanceName { get; set; }

        /// <summary>
        /// Start time of the RI.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Status of the RI at the time of purchase.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public ReservedInstanceState()
        {
        }
    }
}
