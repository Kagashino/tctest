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
    public sealed class AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoResult
    {
        public readonly ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoAudioResult> Audios;
        public readonly bool RemoveAudio;
        public readonly ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoVideoResult> Videos;

        [OutputConstructor]
        private AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoResult(
            ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoAudioResult> audios,

            bool removeAudio,

            ImmutableArray<Outputs.AdaptiveDynamicStreamingTemplatesTemplateListStreamInfoVideoResult> videos)
        {
            Audios = audios;
            RemoveAudio = removeAudio;
            Videos = videos;
        }
    }
}
