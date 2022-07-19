// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function instances(args?: InstancesArgs, opts?: pulumi.InvokeOptions): Promise<InstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Sqlserver/instances:Instances", {
        "id": args.id,
        "name": args.name,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesArgs {
    id?: string;
    name?: string;
    projectId?: number;
    resultOutputFile?: string;
    subnetId?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by Instances.
 */
export interface InstancesResult {
    readonly id?: string;
    readonly instanceLists: outputs.Sqlserver.InstancesInstanceList[];
    readonly name?: string;
    readonly projectId?: number;
    readonly resultOutputFile?: string;
    readonly subnetId?: string;
    readonly vpcId?: string;
}

export function instancesOutput(args?: InstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InstancesResult> {
    return pulumi.output(args).apply(a => instances(a, opts))
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    projectId?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
