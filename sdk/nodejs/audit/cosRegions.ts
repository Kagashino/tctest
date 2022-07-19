// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function cosRegions(args?: CosRegionsArgs, opts?: pulumi.InvokeOptions): Promise<CosRegionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Audit/cosRegions:CosRegions", {
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking CosRegions.
 */
export interface CosRegionsArgs {
    resultOutputFile?: string;
}

/**
 * A collection of values returned by CosRegions.
 */
export interface CosRegionsResult {
    readonly auditCosRegionLists: outputs.Audit.CosRegionsAuditCosRegionList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}

export function cosRegionsOutput(args?: CosRegionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CosRegionsResult> {
    return pulumi.output(args).apply(a => cosRegions(a, opts))
}

/**
 * A collection of arguments for invoking CosRegions.
 */
export interface CosRegionsOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
}
