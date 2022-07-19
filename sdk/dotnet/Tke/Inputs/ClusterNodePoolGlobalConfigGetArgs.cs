// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Tke.Inputs
{

    public sealed class ClusterNodePoolGlobalConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("expander")]
        public Input<string>? Expander { get; set; }

        [Input("ignoreDaemonSetsUtilization")]
        public Input<bool>? IgnoreDaemonSetsUtilization { get; set; }

        [Input("isScaleInEnabled")]
        public Input<bool>? IsScaleInEnabled { get; set; }

        [Input("maxConcurrentScaleIn")]
        public Input<int>? MaxConcurrentScaleIn { get; set; }

        [Input("scaleInDelay")]
        public Input<int>? ScaleInDelay { get; set; }

        [Input("scaleInUnneededTime")]
        public Input<int>? ScaleInUnneededTime { get; set; }

        [Input("scaleInUtilizationThreshold")]
        public Input<int>? ScaleInUtilizationThreshold { get; set; }

        [Input("skipNodesWithLocalStorage")]
        public Input<bool>? SkipNodesWithLocalStorage { get; set; }

        [Input("skipNodesWithSystemPods")]
        public Input<bool>? SkipNodesWithSystemPods { get; set; }

        public ClusterNodePoolGlobalConfigGetArgs()
        {
        }
    }
}
