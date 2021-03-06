// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod
{
    public static class ProcedureTemplates
    {
        public static Task<ProcedureTemplatesResult> InvokeAsync(ProcedureTemplatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ProcedureTemplatesResult>("tctest:Vod/procedureTemplates:ProcedureTemplates", args ?? new ProcedureTemplatesArgs(), options.WithDefaults());

        public static Output<ProcedureTemplatesResult> Invoke(ProcedureTemplatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ProcedureTemplatesResult>("tctest:Vod/procedureTemplates:ProcedureTemplates", args ?? new ProcedureTemplatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class ProcedureTemplatesArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public int? SubAppId { get; set; }

        [Input("type")]
        public string? Type { get; set; }

        public ProcedureTemplatesArgs()
        {
        }
    }

    public sealed class ProcedureTemplatesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("subAppId")]
        public Input<int>? SubAppId { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProcedureTemplatesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ProcedureTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        public readonly string? ResultOutputFile;
        public readonly int? SubAppId;
        public readonly ImmutableArray<Outputs.ProcedureTemplatesTemplateListResult> TemplateLists;
        public readonly string? Type;

        [OutputConstructor]
        private ProcedureTemplatesResult(
            string id,

            string? name,

            string? resultOutputFile,

            int? subAppId,

            ImmutableArray<Outputs.ProcedureTemplatesTemplateListResult> templateLists,

            string? type)
        {
            Id = id;
            Name = name;
            ResultOutputFile = resultOutputFile;
            SubAppId = subAppId;
            TemplateLists = templateLists;
            Type = type;
        }
    }
}
