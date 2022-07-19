// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class VipEipAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VipEipAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VipEipAttachmentState, opts?: pulumi.CustomResourceOptions): VipEipAttachment {
        return new VipEipAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Ha/vipEipAttachment:VipEipAttachment';

    /**
     * Returns true if the given object is an instance of VipEipAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VipEipAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VipEipAttachment.__pulumiType;
    }

    /**
     * Public address of the EIP.
     */
    public readonly addressIp!: pulumi.Output<string>;
    /**
     * ID of the attached HA VIP.
     */
    public readonly havipId!: pulumi.Output<string>;

    /**
     * Create a VipEipAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VipEipAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VipEipAttachmentArgs | VipEipAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VipEipAttachmentState | undefined;
            resourceInputs["addressIp"] = state ? state.addressIp : undefined;
            resourceInputs["havipId"] = state ? state.havipId : undefined;
        } else {
            const args = argsOrState as VipEipAttachmentArgs | undefined;
            if ((!args || args.addressIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressIp'");
            }
            if ((!args || args.havipId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'havipId'");
            }
            resourceInputs["addressIp"] = args ? args.addressIp : undefined;
            resourceInputs["havipId"] = args ? args.havipId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VipEipAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VipEipAttachment resources.
 */
export interface VipEipAttachmentState {
    /**
     * Public address of the EIP.
     */
    addressIp?: pulumi.Input<string>;
    /**
     * ID of the attached HA VIP.
     */
    havipId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VipEipAttachment resource.
 */
export interface VipEipAttachmentArgs {
    /**
     * Public address of the EIP.
     */
    addressIp: pulumi.Input<string>;
    /**
     * ID of the attached HA VIP.
     */
    havipId: pulumi.Input<string>;
}
