// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function vips(args?: VipsArgs, opts?: pulumi.InvokeOptions): Promise<VipsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Ha/vips:Vips", {
        "addressIp": args.addressIp,
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking Vips.
 */
export interface VipsArgs {
    addressIp?: string;
    id?: string;
    name?: string;
    resultOutputFile?: string;
    subnetId?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by Vips.
 */
export interface VipsResult {
    readonly addressIp?: string;
    readonly haVipLists: outputs.Ha.VipsHaVipList[];
    readonly id?: string;
    readonly name?: string;
    readonly resultOutputFile?: string;
    readonly subnetId?: string;
    readonly vpcId?: string;
}

export function vipsOutput(args?: VipsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VipsResult> {
    return pulumi.output(args).apply(a => vips(a, opts))
}

/**
 * A collection of arguments for invoking Vips.
 */
export interface VipsOutputArgs {
    addressIp?: pulumi.Input<string>;
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
