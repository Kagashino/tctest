// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Address
{
    [TctestResourceType("tctest:Address/templateGroup:TemplateGroup")]
    public partial class TemplateGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the address template group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Template ID list.
        /// </summary>
        [Output("templateIds")]
        public Output<ImmutableArray<string>> TemplateIds { get; private set; } = null!;


        /// <summary>
        /// Create a TemplateGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TemplateGroup(string name, TemplateGroupArgs args, CustomResourceOptions? options = null)
            : base("tctest:Address/templateGroup:TemplateGroup", name, args ?? new TemplateGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TemplateGroup(string name, Input<string> id, TemplateGroupState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Address/templateGroup:TemplateGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TemplateGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TemplateGroup Get(string name, Input<string> id, TemplateGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new TemplateGroup(name, id, state, options);
        }
    }

    public sealed class TemplateGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the address template group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("templateIds", required: true)]
        private InputList<string>? _templateIds;

        /// <summary>
        /// Template ID list.
        /// </summary>
        public InputList<string> TemplateIds
        {
            get => _templateIds ?? (_templateIds = new InputList<string>());
            set => _templateIds = value;
        }

        public TemplateGroupArgs()
        {
        }
    }

    public sealed class TemplateGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the address template group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("templateIds")]
        private InputList<string>? _templateIds;

        /// <summary>
        /// Template ID list.
        /// </summary>
        public InputList<string> TemplateIds
        {
            get => _templateIds ?? (_templateIds = new InputList<string>());
            set => _templateIds = value;
        }

        public TemplateGroupState()
        {
        }
    }
}
