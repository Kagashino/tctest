// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function keys(args?: KeysArgs, opts?: pulumi.InvokeOptions): Promise<KeysResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Kms/keys:Keys", {
        "keyState": args.keyState,
        "keyUsage": args.keyUsage,
        "orderType": args.orderType,
        "origin": args.origin,
        "resultOutputFile": args.resultOutputFile,
        "role": args.role,
        "searchKeyAlias": args.searchKeyAlias,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking Keys.
 */
export interface KeysArgs {
    keyState?: number;
    keyUsage?: string;
    orderType?: number;
    origin?: string;
    resultOutputFile?: string;
    role?: number;
    searchKeyAlias?: string;
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by Keys.
 */
export interface KeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyLists: outputs.Kms.KeysKeyList[];
    readonly keyState?: number;
    readonly keyUsage?: string;
    readonly orderType?: number;
    readonly origin?: string;
    readonly resultOutputFile?: string;
    readonly role?: number;
    readonly searchKeyAlias?: string;
    readonly tags?: {[key: string]: any};
}

export function keysOutput(args?: KeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<KeysResult> {
    return pulumi.output(args).apply(a => keys(a, opts))
}

/**
 * A collection of arguments for invoking Keys.
 */
export interface KeysOutputArgs {
    keyState?: pulumi.Input<number>;
    keyUsage?: pulumi.Input<string>;
    orderType?: pulumi.Input<number>;
    origin?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    role?: pulumi.Input<number>;
    searchKeyAlias?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
}
