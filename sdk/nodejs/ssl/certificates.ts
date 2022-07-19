// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function certificates(args?: CertificatesArgs, opts?: pulumi.InvokeOptions): Promise<CertificatesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Ssl/certificates:Certificates", {
        "id": args.id,
        "name": args.name,
        "resultOutputFile": args.resultOutputFile,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking Certificates.
 */
export interface CertificatesArgs {
    id?: string;
    name?: string;
    resultOutputFile?: string;
    type?: string;
}

/**
 * A collection of values returned by Certificates.
 */
export interface CertificatesResult {
    readonly certificates: outputs.Ssl.CertificatesCertificate[];
    readonly id?: string;
    readonly name?: string;
    readonly resultOutputFile?: string;
    readonly type?: string;
}

export function certificatesOutput(args?: CertificatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CertificatesResult> {
    return pulumi.output(args).apply(a => certificates(a, opts))
}

/**
 * A collection of arguments for invoking Certificates.
 */
export interface CertificatesOutputArgs {
    id?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
}
