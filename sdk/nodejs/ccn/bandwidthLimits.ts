// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function bandwidthLimits(args: BandwidthLimitsArgs, opts?: pulumi.InvokeOptions): Promise<BandwidthLimitsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Ccn/bandwidthLimits:BandwidthLimits", {
        "ccnId": args.ccnId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking BandwidthLimits.
 */
export interface BandwidthLimitsArgs {
    ccnId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by BandwidthLimits.
 */
export interface BandwidthLimitsResult {
    readonly ccnId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limits: outputs.Ccn.BandwidthLimitsLimit[];
    readonly resultOutputFile?: string;
}

export function bandwidthLimitsOutput(args: BandwidthLimitsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BandwidthLimitsResult> {
    return pulumi.output(args).apply(a => bandwidthLimits(a, opts))
}

/**
 * A collection of arguments for invoking BandwidthLimits.
 */
export interface BandwidthLimitsOutputArgs {
    ccnId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
