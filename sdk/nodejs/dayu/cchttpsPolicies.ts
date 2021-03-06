// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function cchttpsPolicies(args: CchttpsPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<CchttpsPoliciesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Dayu/cCHttpsPolicies:CCHttpsPolicies", {
        "name": args.name,
        "policyId": args.policyId,
        "resourceId": args.resourceId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking CCHttpsPolicies.
 */
export interface CchttpsPoliciesArgs {
    name?: string;
    policyId?: string;
    resourceId: string;
    resourceType: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by CCHttpsPolicies.
 */
export interface CchttpsPoliciesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Dayu.CCHttpsPoliciesList[];
    readonly name?: string;
    readonly policyId?: string;
    readonly resourceId: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}

export function cchttpsPoliciesOutput(args: CchttpsPoliciesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CchttpsPoliciesResult> {
    return pulumi.output(args).apply(a => cchttpsPolicies(a, opts))
}

/**
 * A collection of arguments for invoking CCHttpsPolicies.
 */
export interface CchttpsPoliciesOutputArgs {
    name?: pulumi.Input<string>;
    policyId?: pulumi.Input<string>;
    resourceId: pulumi.Input<string>;
    resourceType: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
