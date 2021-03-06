// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function apigatewayUsagePlanEnvironments(args: ApigatewayUsagePlanEnvironmentsArgs, opts?: pulumi.InvokeOptions): Promise<ApigatewayUsagePlanEnvironmentsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cloud/aPIGatewayUsagePlanEnvironments:APIGatewayUsagePlanEnvironments", {
        "bindType": args.bindType,
        "resultOutputFile": args.resultOutputFile,
        "usagePlanId": args.usagePlanId,
    }, opts);
}

/**
 * A collection of arguments for invoking APIGatewayUsagePlanEnvironments.
 */
export interface ApigatewayUsagePlanEnvironmentsArgs {
    bindType?: string;
    resultOutputFile?: string;
    usagePlanId: string;
}

/**
 * A collection of values returned by APIGatewayUsagePlanEnvironments.
 */
export interface ApigatewayUsagePlanEnvironmentsResult {
    readonly bindType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Cloud.APIGatewayUsagePlanEnvironmentsList[];
    readonly resultOutputFile?: string;
    readonly usagePlanId: string;
}

export function apigatewayUsagePlanEnvironmentsOutput(args: ApigatewayUsagePlanEnvironmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ApigatewayUsagePlanEnvironmentsResult> {
    return pulumi.output(args).apply(a => apigatewayUsagePlanEnvironments(a, opts))
}

/**
 * A collection of arguments for invoking APIGatewayUsagePlanEnvironments.
 */
export interface ApigatewayUsagePlanEnvironmentsOutputArgs {
    bindType?: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
    usagePlanId: pulumi.Input<string>;
}
