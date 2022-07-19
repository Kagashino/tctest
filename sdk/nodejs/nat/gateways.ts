// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function gateways(args?: GatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GatewaysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Nat/gateways:Gateways", {
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysArgs {
    id?: string;
    name?: string;
    resultOutputFile?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by Gateways.
 */
export interface GatewaysResult {
    readonly id?: string;
    readonly name?: string;
    readonly nats: outputs.Nat.GatewaysNat[];
    readonly resultOutputFile?: string;
    readonly vpcId?: string;
}

export function gatewaysOutput(args?: GatewaysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewaysResult> {
    return pulumi.output(args).apply(a => gateways(a, opts))
}

/**
 * A collection of arguments for invoking Gateways.
 */
export interface GatewaysOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
