// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function tableGroups(args: TableGroupsArgs, opts?: pulumi.InvokeOptions): Promise<TableGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Tcaplus/tableGroups:TableGroups", {
        "clusterId": args.clusterId,
        "resultOutputFile": args.resultOutputFile,
        "tablegroupId": args.tablegroupId,
        "tablegroupName": args.tablegroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking TableGroups.
 */
export interface TableGroupsArgs {
    clusterId: string;
    resultOutputFile?: string;
    tablegroupId?: string;
    tablegroupName?: string;
}

/**
 * A collection of values returned by TableGroups.
 */
export interface TableGroupsResult {
    readonly clusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Tcaplus.TableGroupsList[];
    readonly resultOutputFile?: string;
    readonly tablegroupId?: string;
    readonly tablegroupName?: string;
}

export function tableGroupsOutput(args: TableGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TableGroupsResult> {
    return pulumi.output(args).apply(a => tableGroups(a, opts))
}

/**
 * A collection of arguments for invoking TableGroups.
 */
export interface TableGroupsOutputArgs {
    clusterId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    tablegroupId?: pulumi.Input<string>;
    tablegroupName?: pulumi.Input<string>;
}
