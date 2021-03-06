// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function l7rules(args: L7rulesArgs, opts?: pulumi.InvokeOptions): Promise<L7rulesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Dayu/l7Rules:L7Rules", {
        "domain": args.domain,
        "resourceId": args.resourceId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
        "ruleId": args.ruleId,
    }, opts);
}

/**
 * A collection of arguments for invoking L7Rules.
 */
export interface L7rulesArgs {
    domain?: string;
    resourceId: string;
    resourceType: string;
    resultOutputFile?: string;
    ruleId?: string;
}

/**
 * A collection of values returned by L7Rules.
 */
export interface L7rulesResult {
    readonly domain?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Dayu.L7RulesList[];
    readonly resourceId: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
    readonly ruleId?: string;
}

export function l7rulesOutput(args: L7rulesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<L7rulesResult> {
    return pulumi.output(args).apply(a => l7rules(a, opts))
}

/**
 * A collection of arguments for invoking L7Rules.
 */
export interface L7rulesOutputArgs {
    domain?: pulumi.Input<string>;
    resourceId: pulumi.Input<string>;
    resourceType: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    ruleId?: pulumi.Input<string>;
}
