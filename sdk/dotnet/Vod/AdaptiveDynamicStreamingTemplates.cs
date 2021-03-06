// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod
{
    public static class AdaptiveDynamicStreamingTemplates
    {
        public static Task<AdaptiveDynamicStreamingTemplatesResult> InvokeAsync(AdaptiveDynamicStreamingTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<AdaptiveDynamicStreamingTemplatesResult>("tctest:Vod/adaptiveDynamicStreamingTemplates:AdaptiveDynamicStreamingTemplates", args ?? new AdaptiveDynamicStreamingTemplatesArgs(), options.WithDefaults());

        public static Output<AdaptiveDynamicStreamingTemplatesResult> Invoke(AdaptiveDynamicStreamingTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<AdaptiveDynamicStreamingTemplatesResult>("tctest:Vod/adaptiveDynamicStreamingTemplates:AdaptiveDynamicStreamingTemplates", args ?? new AdaptiveDynamicStreamingTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class AdaptiveDynamicStreamingTemplatesArgs : Pulumi.InvokeArgs
    {
        [Input("definition")]
        public string? Definition { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public int? SubAppId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public AdaptiveDynamicStreamingTemplatesArgs()
        {
        }
    }

    public sealed class AdaptiveDynamicStreamingTemplatesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("definition")]
        public Input<string>? Definition { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public AdaptiveDynamicStreamingTemplatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class AdaptiveDynamicStreamingTemplatesResult
    {
        public readonly string? Definition;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int? SubAppId;
        public readonly ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListResult> TemplateLists;
        public readonly string? Type;

        [OutputConstructor]
        private AdaptiveDynamicStreamingTemplatesResult(
            string? definition,

            string id,

            string? resultOutputFile,

            int? subAppId,

            ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListResult> templateLists,

            string? type)
        {
            Definition = definition;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SubAppId = subAppId;
            TemplateLists = templateLists;
            Type = type;
        }
    }
}
