// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function apigatewayIpStrategy(args: ApigatewayIpStrategyArgs, opts?: pulumi.InvokeOptions): Promise<ApigatewayIpStrategyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Cloud/aPIGatewayIpStrategy:APIGatewayIpStrategy", {
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
        "strategyName": args.strategyName,
    }, opts);
}

/**
 * A collection of arguments for invoking APIGatewayIpStrategy.
 */
export interface ApigatewayIpStrategyArgs {
    resultOutputFile?: string;
    serviceId: string;
    strategyName?: string;
}

/**
 * A collection of values returned by APIGatewayIpStrategy.
 */
export interface ApigatewayIpStrategyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lists: outputs.Cloud.APIGatewayIpStrategyList[];
    readonly resultOutputFile?: string;
    readonly serviceId: string;
    readonly strategyName?: string;
}

export function apigatewayIpStrategyOutput(args: ApigatewayIpStrategyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ApigatewayIpStrategyResult> {
    return pulumi.output(args).apply(a => apigatewayIpStrategy(a, opts))
}

/**
 * A collection of arguments for invoking APIGatewayIpStrategy.
 */
export interface ApigatewayIpStrategyOutputArgs {
    resultOutputFile?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
    strategyName?: pulumi.Input<string>;
}
