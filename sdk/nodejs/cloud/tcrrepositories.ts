// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function tcrrepositories(args: TcrrepositoriesArgs, opts?: pulumi.InvokeOptions): Promise<TcrrepositoriesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cloud/tCRRepositories:TCRRepositories", {
        "instanceId": args.instanceId,
        "namespaceName": args.namespaceName,
        "repositoryName": args.repositoryName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking TCRRepositories.
 */
export interface TcrrepositoriesArgs {
    instanceId: string;
    namespaceName: string;
    repositoryName?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by TCRRepositories.
 */
export interface TcrrepositoriesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly namespaceName: string;
    readonly repositoryLists: outputs.Cloud.TCRRepositoriesRepositoryList[];
    readonly repositoryName?: string;
    readonly resultOutputFile?: string;
}

export function tcrrepositoriesOutput(args: TcrrepositoriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TcrrepositoriesResult> {
    return pulumi.output(args).apply(a => tcrrepositories(a, opts))
}

/**
 * A collection of arguments for invoking TCRRepositories.
 */
export interface TcrrepositoriesOutputArgs {
    instanceId: pulumi.Input<string>;
    namespaceName: pulumi.Input<string>;
    repositoryName?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
