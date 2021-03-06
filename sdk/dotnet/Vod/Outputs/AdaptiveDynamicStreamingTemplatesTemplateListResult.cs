// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Vod.Outputs
{

    [OutputType]
    public sealed class AdaptiveDynamicStreamingTemplatesTemplateListResult
    {
        public readonly string Comment;
        public readonly string CreateTime;
        public readonly string Definition;
        public readonly bool DisableHigherVideoBitrate;
        public readonly bool DisableHigherVideoResolution;
        public readonly string DrmType;
        public readonly string Format;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoResult> StreamInfos;
        public readonly string Type;
        public readonly string UpdateTime;

        [OutputConstructor]
        private AdaptiveDynamicStreamingTemplatesTemplateListResult(
            string comment,

            string createTime,

            string definition,

            bool disableHigherVideoBitrate,

            bool disableHigherVideoResolution,

            string drmType,

            string format,

            string name,

            ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoResult> streamInfos,

            string type,

            string updateTime)
        {
            Comment = comment;
            CreateTime = createTime;
            Definition = definition;
            DisableHigherVideoBitrate = disableHigherVideoBitrate;
            DisableHigherVideoResolution = disableHigherVideoResolution;
            DrmType = drmType;
            Format = format;
            Name = name;
            StreamInfos = streamInfos;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
