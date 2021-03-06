// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function httpDomains(args: HttpDomainsArgs, opts?: pulumi.InvokeOptions): Promise<HttpDomainsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Gaap/httpDomains:HttpDomains", {
        "domain": args.domain,
        "listenerId": args.listenerId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking HttpDomains.
 */
export interface HttpDomainsArgs {
    domain: string;
    listenerId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by HttpDomains.
 */
export interface HttpDomainsResult {
    readonly domain: string;
    readonly domains: outputs.Gaap.HttpDomainsDomain[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listenerId: string;
    readonly resultOutputFile?: string;
}

export function httpDomainsOutput(args: HttpDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<HttpDomainsResult> {
    return pulumi.output(args).apply(a => httpDomains(a, opts))
}

/**
 * A collection of arguments for invoking HttpDomains.
 */
export interface HttpDomainsOutputArgs {
    domain: pulumi.Input<string>;
    listenerId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
