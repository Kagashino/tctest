// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function accounts(args: AccountsArgs, opts?: pulumi.InvokeOptions): Promise<AccountsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Sqlserver/accounts:Accounts", {
        "instanceId": args.instanceId,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsArgs {
    instanceId: string;
    name?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by Accounts.
 */
export interface AccountsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly lists: outputs.Sqlserver.AccountsList[];
    readonly name?: string;
    readonly resultOutputFile?: string;
}

export function accountsOutput(args: AccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AccountsResult> {
    return pulumi.output(args).apply(a => accounts(a, opts))
}

/**
 * A collection of arguments for invoking Accounts.
 */
export interface AccountsOutputArgs {
    instanceId: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
