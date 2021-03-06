// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function snapshotPolicies(args?: SnapshotPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<SnapshotPoliciesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cbs/snapshotPolicies:SnapshotPolicies", {
        "resultOutputFile": args.resultOutputFile,
        "snapshotPolicyId": args.snapshotPolicyId,
        "snapshotPolicyName": args.snapshotPolicyName,
    }, opts);
}

/**
 * A collection of arguments for invoking SnapshotPolicies.
 */
export interface SnapshotPoliciesArgs {
    resultOutputFile?: string;
    snapshotPolicyId?: string;
    snapshotPolicyName?: string;
}

/**
 * A collection of values returned by SnapshotPolicies.
 */
export interface SnapshotPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly snapshotPolicyId?: string;
    readonly snapshotPolicyLists: outputs.Cbs.SnapshotPoliciesSnapshotPolicyList[];
    readonly snapshotPolicyName?: string;
}

export function snapshotPoliciesOutput(args?: SnapshotPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SnapshotPoliciesResult> {
    return pulumi.output(args).apply(a => snapshotPolicies(a, opts))
}

/**
 * A collection of arguments for invoking SnapshotPolicies.
 */
export interface SnapshotPoliciesOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    snapshotPolicyId?: pulumi.Input<string>;
    snapshotPolicyName?: pulumi.Input<string>;
}
