// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function fileSystems(args?: FileSystemsArgs, opts?: pulumi.InvokeOptions): Promise<FileSystemsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cfs/fileSystems:FileSystems", {
        "availabilityZone": args.availabilityZone,
        "fileSystemId": args.fileSystemId,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "subnetId": args.subnetId,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking FileSystems.
 */
export interface FileSystemsArgs {
    availabilityZone?: string;
    fileSystemId?: string;
    name?: string;
    resultOutputFile?: string;
    subnetId?: string;
    vpcId?: string;
}

/**
 * A collection of values returned by FileSystems.
 */
export interface FileSystemsResult {
    readonly availabilityZone?: string;
    readonly fileSystemId?: string;
    readonly fileSystemLists: outputs.Cfs.FileSystemsFileSystemList[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly resultOutputFile?: string;
    readonly subnetId?: string;
    readonly vpcId?: string;
}

export function fileSystemsOutput(args?: FileSystemsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FileSystemsResult> {
    return pulumi.output(args).apply(a => fileSystems(a, opts))
}

/**
 * A collection of arguments for invoking FileSystems.
 */
export interface FileSystemsOutputArgs {
    availabilityZone?: pulumi.Input<string>;
    fileSystemId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subnetId?: pulumi.Input<string>;
    vpcId?: pulumi.Input<string>;
}
