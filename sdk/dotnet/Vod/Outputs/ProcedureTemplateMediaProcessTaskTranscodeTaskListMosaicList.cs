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
    public sealed class ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicList
    {
        public readonly string? CoordinateOrigin;
        public readonly double? EndTimeOffset;
        public readonly string? Height;
        public readonly double? StartTimeOffset;
        public readonly string? Width;
        public readonly string? XPos;
        public readonly string? YPos;

        [OutputConstructor]
        private ProcedureTemplateMediaProcessTaskTranscodeTaskListMosaicList(
            string? coordinateOrigin,

            double? endTimeOffset,

            string? height,

            double? startTimeOffset,

            string? width,

            string? xPos,

            string? yPos)
        {
            CoordinateOrigin = coordinateOrigin;
            EndTimeOffset = endTimeOffset;
            Height = height;
            StartTimeOffset = startTimeOffset;
            Width = width;
            XPos = xPos;
            YPos = yPos;
        }
    }
}
