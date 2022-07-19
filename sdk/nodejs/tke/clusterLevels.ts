// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function clusterLevels(args?: ClusterLevelsArgs, opts?: pulumi.InvokeOptions): Promise<ClusterLevelsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Tke/clusterLevels:ClusterLevels", {
        "clusterId": args.clusterId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking ClusterLevels.
 */
export interface ClusterLevelsArgs {
    clusterId?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by ClusterLevels.
 */
export interface ClusterLevelsResult {
    readonly clusterId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Tke.ClusterLevelsList[];
    readonly resultOutputFile?: string;
}

export function clusterLevelsOutput(args?: ClusterLevelsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ClusterLevelsResult> {
    return pulumi.output(args).apply(a => clusterLevels(a, opts))
}

/**
 * A collection of arguments for invoking ClusterLevels.
 */
export interface ClusterLevelsOutputArgs {
    clusterId?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
