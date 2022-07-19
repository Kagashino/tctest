// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cbs
{
    [TctestResourceType("tctest:Cbs/storageAttachment:StorageAttachment")]
    public partial class StorageAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the CVM instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// ID of the mounted CBS.
        /// </summary>
        [Output("storageId")]
        public Output<string> StorageId { get; private set; } = null!;


        /// <summary>
        /// Create a StorageAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StorageAttachment(string name, StorageAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tctest:Cbs/storageAttachment:StorageAttachment", name, args ?? new StorageAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StorageAttachment(string name, Input<string> id, StorageAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Cbs/storageAttachment:StorageAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StorageAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StorageAttachment Get(string name, Input<string> id, StorageAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new StorageAttachment(name, id, state, options);
        }
    }

    public sealed class StorageAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the CVM instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// ID of the mounted CBS.
        /// </summary>
        [Input("storageId", required: true)]
        public Input<string> StorageId { get; set; } = null!;

        public StorageAttachmentArgs()
        {
        }
    }

    public sealed class StorageAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the CVM instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// ID of the mounted CBS.
        /// </summary>
        [Input("storageId")]
        public Input<string>? StorageId { get; set; }

        public StorageAttachmentState()
        {
        }
    }
}
