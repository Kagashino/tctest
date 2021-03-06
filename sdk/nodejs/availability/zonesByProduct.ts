// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function zonesByProduct(args: ZonesByProductArgs, opts?: pulumi.InvokeOptions): Promise<ZonesByProductResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Availability/zonesByProduct:ZonesByProduct", {
        "includeUnavailable": args.includeUnavailable,
        "name": args.name,
        "product": args.product,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking ZonesByProduct.
 */
export interface ZonesByProductArgs {
    includeUnavailable?: boolean;
    name?: string;
    product: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by ZonesByProduct.
 */
export interface ZonesByProductResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeUnavailable?: boolean;
    readonly name?: string;
    readonly product: string;
    readonly resultOutputFile?: string;
    readonly zones: outputs.Availability.ZonesByProductZone[];
}

export function zonesByProductOutput(args: ZonesByProductOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ZonesByProductResult> {
    return pulumi.output(args).apply(a => zonesByProduct(a, opts))
}

/**
 * A collection of arguments for invoking ZonesByProduct.
 */
export interface ZonesByProductOutputArgs {
    includeUnavailable?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    product: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
