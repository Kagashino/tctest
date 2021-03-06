// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Monitor
{
    public static class Data
    {
        public static Task<DataResult> InvokeAsync(DataArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<DataResult>("tctest:Monitor/data:Data", args ?? new DataArgs(), options.WithDefaults());

        public static Output<DataResult> Invoke(DataInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<DataResult>("tctest:Monitor/data:Data", args ?? new DataInvokeArgs(), options.WithDefaults());
    }


    public sealed class DataArgs : Pulumi.InvokeArgs
    {
        [Input("dimensions", required: true)]
        private List<Inputs.DataDimensionArgs>? _dimensions;
        public List<Inputs.DataDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new List<Inputs.DataDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("endTime", required: true)]
        public string EndTime { get; set; } = null!;

        [Input("metricName", required: true)]
        public string MetricName { get; set; } = null!;

        [Input("namespace", required: true)]
        public string Namespace { get; set; } = null!;

        [Input("period")]
        public int? Period { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public string StartTime { get; set; } = null!;

        public DataArgs()
        {
        }
    }

    public sealed class DataInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("dimensions", required: true)]
        private InputList<Inputs.DataDimensionInputArgs>? _dimensions;
        public InputList<Inputs.DataDimensionInputArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.DataDimensionInputArgs>());
            set => _dimensions = value;
        }

        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("period")]
        public Input<int>? Period { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public DataInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class DataResult
    {
        public readonly ImmutableArray<Outputs.DataDimensionResult> Dimensions;
        public readonly string EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.DataListResult> Lists;
        public readonly string MetricName;
        public readonly string Namespace;
        public readonly int? Period;
        public readonly string? ResultOutputFile;
        public readonly string StartTime;

        [OutputConstructor]
        private DataResult(
            ImmutableArray<Outputs.DataDimensionResult> dimensions,

            string endTime,

            string id,

            ImmutableArray<Outputs.DataListResult> lists,

            string metricName,

            string @namespace,

            int? period,

            string? resultOutputFile,

            string startTime)
        {
            Dimensions = dimensions;
            EndTime = endTime;
            Id = id;
            Lists = lists;
            MetricName = metricName;
            Namespace = @namespace;
            Period = period;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
        }
    }
}
