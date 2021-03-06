// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function gatewaySnats(args: GatewaySnatsArgs, opts?: pulumi.InvokeOptions): Promise<GatewaySnatsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Nat/gatewaySnats:GatewaySnats", {
        "description": args.description,
        "instanceId": args.instanceId,
        "natGatewayId": args.natGatewayId,
        "publicIpAddrs": args.publicIpAddrs,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
    }, opts);
}

/**
 * A collection of arguments for invoking GatewaySnats.
 */
export interface GatewaySnatsArgs {
    description?: string;
    instanceId?: string;
    natGatewayId: string;
    publicIpAddrs?: string[];
    resultOutputFile?: string;
    subnetId?: string;
}

/**
 * A collection of values returned by GatewaySnats.
 */
export interface GatewaySnatsResult {
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly natGatewayId: string;
    readonly publicIpAddrs?: string[];
    readonly resultOutputFile?: string;
    readonly snatLists: outputs.Nat.GatewaySnatsSnatList[];
    readonly subnetId?: string;
}

export function gatewaySnatsOutput(args: GatewaySnatsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewaySnatsResult> {
    return pulumi.output(args).apply(a => gatewaySnats(a, opts))
}

/**
 * A collection of arguments for invoking GatewaySnats.
 */
export interface GatewaySnatsOutputArgs {
    description?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    natGatewayId: pulumi.Input<string>;
    publicIpAddrs?: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
}
