// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function gatewayCCNRoutes(args: GatewayCCNRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GatewayCCNRoutesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tctest:Dc/gatewayCCNRoutes:GatewayCCNRoutes", {
        "dcgId": args.dcgId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking GatewayCCNRoutes.
 */
export interface GatewayCCNRoutesArgs {
    dcgId: string;
    resultOutputFile?: string;
}

/**
 * A collection of values returned by GatewayCCNRoutes.
 */
export interface GatewayCCNRoutesResult {
    readonly dcgId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceLists: outputs.Dc.GatewayCCNRoutesInstanceList[];
    readonly resultOutputFile?: string;
}

export function gatewayCCNRoutesOutput(args: GatewayCCNRoutesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GatewayCCNRoutesResult> {
    return pulumi.output(args).apply(a => gatewayCCNRoutes(a, opts))
}

/**
 * A collection of arguments for invoking GatewayCCNRoutes.
 */
export interface GatewayCCNRoutesOutputArgs {
    dcgId: pulumi.Input<string>;
    resultOutputFile?: pulumi.Input<string>;
}
