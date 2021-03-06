// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function bucketObjects(args: BucketObjectsArgs, opts?: pulumi.InvokeOptions): Promise<BucketObjectsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cos/bucketObjects:BucketObjects", {
        "bucket": args.bucket,
        "key": args.key,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking BucketObjects.
 */
export interface BucketObjectsArgs {
    bucket: string;
    key: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by BucketObjects.
 */
export interface BucketObjectsResult {
    readonly bucket: string;
    readonly cacheControl: string;
    readonly contentDisposition: string;
    readonly contentEncoding: string;
    readonly contentType: string;
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    readonly lastModified: string;
    readonly resultOutputFile?: string;
    readonly storageClass: string;
}

export function bucketObjectsOutput(args: BucketObjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BucketObjectsResult> {
    return pulumi.output(args).apply(a => bucketObjects(a, opts))
}

/**
 * A collection of arguments for invoking BucketObjects.
 */
export interface BucketObjectsOutputArgs {
    bucket: pulumi.Input<string>;
    key: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
