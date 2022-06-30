// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class Zone extends pulumi.CustomResource {
    /**
     * Get an existing Zone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZoneState, opts?: pulumi.CustomResourceOptions): Zone {
        return new Zone(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:PrivateDns/zone:Zone';

    /**
     * Returns true if the given object is an instance of Zone.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Zone {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Zone.__pulumiType;
    }

    /**
     * List of authorized accounts' VPCs to associate with the private domain.
     */
    public readonly accountVpcSets!: pulumi.Output<outputs.PrivateDns.ZoneAccountVpcSet[] | undefined>;
    /**
     * Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
     */
    public readonly dnsForwardStatus!: pulumi.Output<string | undefined>;
    /**
     * Domain name, which must be in the format of standard TLD.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Remarks.
     */
    public readonly remark!: pulumi.Output<string | undefined>;
    /**
     * Tags the private domain when it is created.
     *
     * @deprecated It has been deprecated from version 1.72.4. Use `tags` instead.
     */
    public readonly tagSets!: pulumi.Output<outputs.PrivateDns.ZoneTagSet[]>;
    /**
     * Tags of the private dns zone.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Associates the private domain to a VPC when it is created.
     */
    public readonly vpcSets!: pulumi.Output<outputs.PrivateDns.ZoneVpcSet[] | undefined>;

    /**
     * Create a Zone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZoneArgs | ZoneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZoneState | undefined;
            resourceInputs["accountVpcSets"] = state ? state.accountVpcSets : undefined;
            resourceInputs["dnsForwardStatus"] = state ? state.dnsForwardStatus : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["tagSets"] = state ? state.tagSets : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcSets"] = state ? state.vpcSets : undefined;
        } else {
            const args = argsOrState as ZoneArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["accountVpcSets"] = args ? args.accountVpcSets : undefined;
            resourceInputs["dnsForwardStatus"] = args ? args.dnsForwardStatus : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["tagSets"] = args ? args.tagSets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcSets"] = args ? args.vpcSets : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Zone.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Zone resources.
 */
export interface ZoneState {
    /**
     * List of authorized accounts' VPCs to associate with the private domain.
     */
    accountVpcSets?: pulumi.Input<pulumi.Input<inputs.PrivateDns.ZoneAccountVpcSet>[]>;
    /**
     * Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
     */
    dnsForwardStatus?: pulumi.Input<string>;
    /**
     * Domain name, which must be in the format of standard TLD.
     */
    domain?: pulumi.Input<string>;
    /**
     * Remarks.
     */
    remark?: pulumi.Input<string>;
    /**
     * Tags the private domain when it is created.
     *
     * @deprecated It has been deprecated from version 1.72.4. Use `tags` instead.
     */
    tagSets?: pulumi.Input<pulumi.Input<inputs.PrivateDns.ZoneTagSet>[]>;
    /**
     * Tags of the private dns zone.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Associates the private domain to a VPC when it is created.
     */
    vpcSets?: pulumi.Input<pulumi.Input<inputs.PrivateDns.ZoneVpcSet>[]>;
}

/**
 * The set of arguments for constructing a Zone resource.
 */
export interface ZoneArgs {
    /**
     * List of authorized accounts' VPCs to associate with the private domain.
     */
    accountVpcSets?: pulumi.Input<pulumi.Input<inputs.PrivateDns.ZoneAccountVpcSet>[]>;
    /**
     * Whether to enable subdomain recursive DNS. Valid values: ENABLED, DISABLED. Default value: DISABLED.
     */
    dnsForwardStatus?: pulumi.Input<string>;
    /**
     * Domain name, which must be in the format of standard TLD.
     */
    domain: pulumi.Input<string>;
    /**
     * Remarks.
     */
    remark?: pulumi.Input<string>;
    /**
     * Tags the private domain when it is created.
     *
     * @deprecated It has been deprecated from version 1.72.4. Use `tags` instead.
     */
    tagSets?: pulumi.Input<pulumi.Input<inputs.PrivateDns.ZoneTagSet>[]>;
    /**
     * Tags of the private dns zone.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Associates the private domain to a VPC when it is created.
     */
    vpcSets?: pulumi.Input<pulumi.Input<inputs.PrivateDns.ZoneVpcSet>[]>;
}