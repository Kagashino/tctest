// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function backups(args: BackupsArgs, opts?: pulumi.InvokeOptions): Promise<BackupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Sqlserver/backups:Backups", {
        "endTime": args.endTime,
        "instanceId": args.instanceId,
        "resultOutputFile": args.resultOutputFile,
        "startTime": args.startTime,
    }, opts);
}

/**
 * A collection of arguments for invoking Backups.
 */
export interface BackupsArgs {
    endTime: string;
    instanceId: string;
    resultOutputFile?: string;
    startTime: string;
}

/**
 * A collection of values returned by Backups.
 */
export interface BackupsResult {
    readonly endTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly lists: outputs.Sqlserver.BackupsList[];
    readonly resultOutputFile?: string;
    readonly startTime: string;
}

export function backupsOutput(args: BackupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BackupsResult> {
    return pulumi.output(args).apply(a => backups(a, opts))
}

/**
 * A collection of arguments for invoking Backups.
 */
export interface BackupsOutputArgs {
    endTime: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    startTime: pulumi.Input<string>;
}
