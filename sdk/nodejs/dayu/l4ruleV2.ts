// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class L4RuleV2 extends pulumi.CustomResource {
    /**
     * Get an existing L4RuleV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L4RuleV2State, opts?: pulumi.CustomResourceOptions): L4RuleV2 {
        return new L4RuleV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Dayu/l4RuleV2:L4RuleV2';

    /**
     * Returns true if the given object is an instance of L4RuleV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L4RuleV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L4RuleV2.__pulumiType;
    }

    /**
     * Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
     */
    public readonly business!: pulumi.Output<string>;
    /**
     * Resource id.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * A list of layer 4 rules. Each element contains the following attributes:
     */
    public readonly rules!: pulumi.Output<outputs.Dayu.L4RuleV2Rules>;
    /**
     * The virtual port of the layer 4 rule.
     */
    public readonly virtualPort!: pulumi.Output<number>;
    /**
     * Resource vpn.
     */
    public readonly vpn!: pulumi.Output<string>;

    /**
     * Create a L4RuleV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L4RuleV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L4RuleV2Args | L4RuleV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as L4RuleV2State | undefined;
            resourceInputs["business"] = state ? state.business : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["virtualPort"] = state ? state.virtualPort : undefined;
            resourceInputs["vpn"] = state ? state.vpn : undefined;
        } else {
            const args = argsOrState as L4RuleV2Args | undefined;
            if ((!args || args.business === undefined) && !opts.urn) {
                throw new Error("Missing required property 'business'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            if ((!args || args.virtualPort === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualPort'");
            }
            if ((!args || args.vpn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpn'");
            }
            resourceInputs["business"] = args ? args.business : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["virtualPort"] = args ? args.virtualPort : undefined;
            resourceInputs["vpn"] = args ? args.vpn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(L4RuleV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L4RuleV2 resources.
 */
export interface L4RuleV2State {
    /**
     * Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
     */
    business?: pulumi.Input<string>;
    /**
     * Resource id.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * A list of layer 4 rules. Each element contains the following attributes:
     */
    rules?: pulumi.Input<inputs.Dayu.L4RuleV2Rules>;
    /**
     * The virtual port of the layer 4 rule.
     */
    virtualPort?: pulumi.Input<number>;
    /**
     * Resource vpn.
     */
    vpn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a L4RuleV2 resource.
 */
export interface L4RuleV2Args {
    /**
     * Business of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
     */
    business: pulumi.Input<string>;
    /**
     * Resource id.
     */
    resourceId: pulumi.Input<string>;
    /**
     * A list of layer 4 rules. Each element contains the following attributes:
     */
    rules: pulumi.Input<inputs.Dayu.L4RuleV2Rules>;
    /**
     * The virtual port of the layer 4 rule.
     */
    virtualPort: pulumi.Input<number>;
    /**
     * Resource vpn.
     */
    vpn: pulumi.Input<string>;
}
