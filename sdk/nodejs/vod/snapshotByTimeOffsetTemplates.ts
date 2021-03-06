// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function snapshotByTimeOffsetTemplates(args?: SnapshotByTimeOffsetTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<SnapshotByTimeOffsetTemplatesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Vod/snapshotByTimeOffsetTemplates:SnapshotByTimeOffsetTemplates", {
        "definition": args.definition,
        "resultOutputFile": args.resultOutputFile,
        "subAppId": args.subAppId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking SnapshotByTimeOffsetTemplates.
 */
export interface SnapshotByTimeOffsetTemplatesArgs {
    definition?: string;
    resultOutputFile?: string;
    subAppId?: number;
    type?: string;
}

/**
 * A collection of values returned by SnapshotByTimeOffsetTemplates.
 */
export interface SnapshotByTimeOffsetTemplatesResult {
    readonly definition?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly subAppId?: number;
    readonly templateLists: outputs.Vod.SnapshotByTimeOffsetTemplatesTemplateList[];
    readonly type?: string;
}

export function snapshotByTimeOffsetTemplatesOutput(args?: SnapshotByTimeOffsetTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SnapshotByTimeOffsetTemplatesResult> {
    return pulumi.output(args).apply(a => snapshotByTimeOffsetTemplates(a, opts))
}

/**
 * A collection of arguments for invoking SnapshotByTimeOffsetTemplates.
 */
export interface SnapshotByTimeOffsetTemplatesOutputArgs {
    definition?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subAppId?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
}
