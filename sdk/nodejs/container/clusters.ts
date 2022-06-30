// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function clusters(args?: ClustersArgs, opts?: pulumi.InvokeOptions): Promise<ClustersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Container/clusters:Clusters", {
        "clusterId": args.clusterId,
        "limit": args.limit,
    }, opts);
}

/**
 * A collection of arguments for invoking Clusters.
 */
export interface ClustersArgs {
    clusterId?: string;
    limit?: number;
}

/**
 * A collection of values returned by Clusters.
 */
export interface ClustersResult {
    readonly clusterId?: string;
    readonly clusters: outputs.Container.ClustersCluster[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limit?: number;
    readonly totalCount: number;
}

export function clustersOutput(args?: ClustersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ClustersResult> {
    return pulumi.output(args).apply(a => clusters(a, opts))
}

/**
 * A collection of arguments for invoking Clusters.
 */
export interface ClustersOutputArgs {
    clusterId?: pulumi.Input<string>;
    limit?: pulumi.Input<number>;
}