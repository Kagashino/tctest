// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function adaptiveDynamicStreamingTemplates(args?: AdaptiveDynamicStreamingTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<AdaptiveDynamicStreamingTemplatesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Vod/adaptiveDynamicStreamingTemplates:AdaptiveDynamicStreamingTemplates", {
        "definition": args.definition,
        "resultOutputFile": args.resultOutputFile,
        "subAppId": args.subAppId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking AdaptiveDynamicStreamingTemplates.
 */
export interface AdaptiveDynamicStreamingTemplatesArgs {
    definition?: string;
    resultOutputFile?: string;
    subAppId?: number;
    type?: string;
}

/**
 * A collection of values returned by AdaptiveDynamicStreamingTemplates.
 */
export interface AdaptiveDynamicStreamingTemplatesResult {
    readonly definition?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    readonly subAppId?: number;
    readonly templateLists: outputs.Vod.AdaptiveDynamicStreamingTemplatesTemplateList[];
    readonly type?: string;
}

export function adaptiveDynamicStreamingTemplatesOutput(args?: AdaptiveDynamicStreamingTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<AdaptiveDynamicStreamingTemplatesResult> {
    return pulumi.output(args).apply(a => adaptiveDynamicStreamingTemplates(a, opts))
}

/**
 * A collection of arguments for invoking AdaptiveDynamicStreamingTemplates.
 */
export interface AdaptiveDynamicStreamingTemplatesOutputArgs {
    definition?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    subAppId?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
}
