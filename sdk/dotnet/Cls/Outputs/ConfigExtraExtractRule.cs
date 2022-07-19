// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cls.Outputs
{

    [OutputType]
    public sealed class ConfigExtraExtractRule
    {
        public readonly int? Backtracking;
        public readonly string? BeginRegex;
        public readonly string? Delimiter;
        public readonly ImmutableArray<Outputs.ConfigExtraExtractRuleFilterKeyRegex> FilterKeyRegexes;
        public readonly ImmutableArray<string> Keys;
        public readonly string? LogRegex;
        public readonly string? TimeFormat;
        public readonly string? TimeKey;
        public readonly string? UnMatchLogKey;
        public readonly bool? UnMatchUpLoadSwitch;

        [OutputConstructor]
        private ConfigExtraExtractRule(
            int? backtracking,

            string? beginRegex,

            string? delimiter,

            ImmutableArray<Outputs.ConfigExtraExtractRuleFilterKeyRegex> filterKeyRegexes,

            ImmutableArray<string> keys,

            string? logRegex,

            string? timeFormat,

            string? timeKey,

            string? unMatchLogKey,

            bool? unMatchUpLoadSwitch)
        {
            Backtracking = backtracking;
            BeginRegex = beginRegex;
            Delimiter = delimiter;
            FilterKeyRegexes = filterKeyRegexes;
            Keys = keys;
            LogRegex = logRegex;
            TimeFormat = timeFormat;
            TimeKey = timeKey;
            UnMatchLogKey = unMatchLogKey;
            UnMatchUpLoadSwitch = unMatchUpLoadSwitch;
        }
    }
}
