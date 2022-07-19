// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cvm
{
    public static class InstancesTypes
    {
        public static Task<InstancesTypesResult> InvokeAsync(InstancesTypesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstancesTypesResult>("tctest:Cvm/instancesTypes:InstancesTypes", args ?? new InstancesTypesArgs(), options.WithDefaults());

        public static Output<InstancesTypesResult> Invoke(InstancesTypesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstancesTypesResult>("tctest:Cvm/instancesTypes:InstancesTypes", args ?? new InstancesTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstancesTypesArgs : Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public string? AvailabilityZone { get; set; }

        [Input("cpuCoreCount")]
        public int? CpuCoreCount { get; set; }

        [Input("excludeSoldOut")]
        public bool? ExcludeSoldOut { get; set; }

        [Input("filters")]
        private List<Inputs.InstancesTypesFilterArgs>? _filters;
        public List<Inputs.InstancesTypesFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.InstancesTypesFilterArgs>());
            set => _filters = value;
        }

        [Input("gpuCoreCount")]
        public int? GpuCoreCount { get; set; }

        [Input("memorySize")]
        public int? MemorySize { get; set; }

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public InstancesTypesArgs()
        {
        }
    }

    public sealed class InstancesTypesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("cpuCoreCount")]
        public Input<int>? CpuCoreCount { get; set; }

        [Input("excludeSoldOut")]
        public Input<bool>? ExcludeSoldOut { get; set; }

        [Input("filters")]
        private InputList<Inputs.InstancesTypesFilterInputArgs>? _filters;
        public InputList<Inputs.InstancesTypesFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.InstancesTypesFilterInputArgs>());
            set => _filters = value;
        }

        [Input("gpuCoreCount")]
        public Input<int>? GpuCoreCount { get; set; }

        [Input("memorySize")]
        public Input<int>? MemorySize { get; set; }

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public InstancesTypesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstancesTypesResult
    {
        public readonly string? AvailabilityZone;
        public readonly int? CpuCoreCount;
        public readonly bool? ExcludeSoldOut;
        public readonly ImmutableArray<Outputs.InstancesTypesFilterResult> Filters;
        public readonly int? GpuCoreCount;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.InstancesTypesInstanceTypeResult> InstanceTypes;
        public readonly int? MemorySize;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private InstancesTypesResult(
            string? availabilityZone,

            int? cpuCoreCount,

            bool? excludeSoldOut,

            ImmutableArray<Outputs.InstancesTypesFilterResult> filters,

            int? gpuCoreCount,

            string id,

            ImmutableArray<Outputs.InstancesTypesInstanceTypeResult> instanceTypes,

            int? memorySize,

            string? resultOutputFile)
        {
            AvailabilityZone = availabilityZone;
            CpuCoreCount = cpuCoreCount;
            ExcludeSoldOut = excludeSoldOut;
            Filters = filters;
            GpuCoreCount = gpuCoreCount;
            Id = id;
            InstanceTypes = instanceTypes;
            MemorySize = memorySize;
            ResultOutputFile = resultOutputFile;
        }
    }
}
