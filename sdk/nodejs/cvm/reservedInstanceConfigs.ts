// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function reservedInstanceConfigs(args?: ReservedInstanceConfigsArgs, opts?: pulumi.InvokeOptions): Promise<ReservedInstanceConfigsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cvm/reservedInstanceConfigs:ReservedInstanceConfigs", {
        "availabilityZone": args.availabilityZone,
        "duration": args.duration,
        "instanceType": args.instanceType,
        "offeringType": args.offeringType,
        "productDescription": args.productDescription,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking ReservedInstanceConfigs.
 */
export interface ReservedInstanceConfigsArgs {
    availabilityZone?: string;
    duration?: number;
    instanceType?: string;
    offeringType?: string;
    productDescription?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by ReservedInstanceConfigs.
 */
export interface ReservedInstanceConfigsResult {
    readonly availabilityZone?: string;
    readonly configLists: outputs.Cvm.ReservedInstanceConfigsConfigList[];
    readonly duration?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceType?: string;
    readonly offeringType?: string;
    readonly productDescription?: string;
    readonly resultOutputFile?: string;
}

export function reservedInstanceConfigsOutput(args?: ReservedInstanceConfigsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ReservedInstanceConfigsResult> {
    return pulumi.output(args).apply(a => reservedInstanceConfigs(a, opts))
}

/**
 * A collection of arguments for invoking ReservedInstanceConfigs.
 */
export interface ReservedInstanceConfigsOutputArgs {
    availabilityZone?: pulumi.Input<string>;
    duration?: pulumi.Input<number>;
    instanceType?: pulumi.Input<string>;
    offeringType?: pulumi.Input<string>;
    productDescription?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
