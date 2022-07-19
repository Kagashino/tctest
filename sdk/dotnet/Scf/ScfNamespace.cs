// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Scf
{
    [TctestResourceType("tctest:Scf/scfNamespace:ScfNamespace")]
    public partial class ScfNamespace : Pulumi.CustomResource
    {
        /// <summary>
        /// SCF namespace creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the SCF namespace.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// SCF namespace last modified time.
        /// </summary>
        [Output("modifyTime")]
        public Output<string> ModifyTime { get; private set; } = null!;

        /// <summary>
        /// Name of the SCF namespace.
        /// </summary>
        [Output("namespace")]
        public Output<string> Namespace { get; private set; } = null!;

        /// <summary>
        /// SCF namespace type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ScfNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScfNamespace(string name, ScfNamespaceArgs args, CustomResourceOptions? options = null)
            : base("tctest:Scf/scfNamespace:ScfNamespace", name, args ?? new ScfNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScfNamespace(string name, Input<string> id, ScfNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Scf/scfNamespace:ScfNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScfNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScfNamespace Get(string name, Input<string> id, ScfNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new ScfNamespace(name, id, state, options);
        }
    }

    public sealed class ScfNamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the SCF namespace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the SCF namespace.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public ScfNamespaceArgs()
        {
        }
    }

    public sealed class ScfNamespaceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// SCF namespace creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the SCF namespace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// SCF namespace last modified time.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        /// <summary>
        /// Name of the SCF namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// SCF namespace type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ScfNamespaceState()
        {
        }
    }
}
