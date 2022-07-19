// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function policyGroups(args?: PolicyGroupsArgs, opts?: pulumi.InvokeOptions): Promise<PolicyGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Monitor/policyGroups:PolicyGroups", {
        "name": args.name,
        "policyViewNames": args.policyViewNames,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking PolicyGroups.
 */
export interface PolicyGroupsArgs {
    name?: string;
    policyViewNames?: string[];
    resultOutputFile?: string;
}

/**
 * A collection of values returned by PolicyGroups.
 */
export interface PolicyGroupsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Monitor.PolicyGroupsList[];
    readonly name?: string;
    readonly policyViewNames?: string[];
    readonly resultOutputFile?: string;
}

export function policyGroupsOutput(args?: PolicyGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<PolicyGroupsResult> {
    return pulumi.output(args).apply(a => policyGroups(a, opts))
}

/**
 * A collection of arguments for invoking PolicyGroups.
 */
export interface PolicyGroupsOutputArgs {
    name?: pulumi.Input<string>;
    policyViewNames?: pulumi.Input<pulumi.Input<string>[]>;
    resultOutputFile?: pulumi.Input<string>;
}
