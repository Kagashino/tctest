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
    return pulumi.runtime.invoke("tctest:Instances/instances:Instances", {
        "availabilityZone": args.availabilityZone,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "projectId": args.projectId,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesArgs {
    availabilityZone?: string;
    instanceId?: string;
    instanceName?: string;
    projectId?: number;
    resultOutputFile?: string;
    subnetId?: string;
    tags?: {[key: string]: any};
    vpcId?: string;
}

/**
 * A collection of values returned by Instances.
 */
export interface InstancesResult {
    readonly availabilityZone?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly instanceLists: outputs.Instances.InstancesInstanceList[];
    readonly instanceName?: string;
    readonly projectId?: number;
    readonly resultOutputFile?: string;
    readonly subnetId?: string;
    readonly tags?: {[key: string]: any};
    readonly vpcId?: string;
}

export function instancesOutput(args?: InstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InstancesResult> {
    return pulumi.output(args).apply(a => instances(a, opts))
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesOutputArgs {
    availabilityZone?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    instanceName?: pulumi.Input<string>;
    projectId?: pulumi.Input<number>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    vpcId?: pulumi.Input<string>;
}
