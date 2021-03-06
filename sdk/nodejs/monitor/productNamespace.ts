// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function productNamespace(args?: ProductNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<ProductNamespaceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Monitor/productNamespace:ProductNamespace", {
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking ProductNamespace.
 */
export interface ProductNamespaceArgs {
    name?: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by ProductNamespace.
 */
export interface ProductNamespaceResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Monitor.ProductNamespaceList[];
    readonly name?: string;
    readonly resultOutputFile?: string;
}

export function productNamespaceOutput(args?: ProductNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ProductNamespaceResult> {
    return pulumi.output(args).apply(a => productNamespace(a, opts))
}

/**
 * A collection of arguments for invoking ProductNamespace.
 */
export interface ProductNamespaceOutputArgs {
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
