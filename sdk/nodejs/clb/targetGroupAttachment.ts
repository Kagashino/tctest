// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class TargetGroupAttachment extends pulumi.CustomResource {
    /**
     * Get an existing TargetGroupAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetGroupAttachmentState, opts?: pulumi.CustomResourceOptions): TargetGroupAttachment {
        return new TargetGroupAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Clb/targetGroupAttachment:TargetGroupAttachment';

    /**
     * Returns true if the given object is an instance of TargetGroupAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetGroupAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetGroupAttachment.__pulumiType;
    }

    /**
     * ID of the CLB.
     */
    public readonly clbId!: pulumi.Output<string>;
    /**
     * ID of the CLB listener.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * ID of the CLB listener rule.
     */
    public readonly ruleId!: pulumi.Output<string | undefined>;
    /**
     * ID of the CLB target group.
     */
    public readonly targetGroupId!: pulumi.Output<string | undefined>;
    /**
     * ID of the CLB target group.
     *
     * @deprecated It has been deprecated from version 1.47.1. Use `target_group_id` instead.
     */
    public readonly targrtGroupId!: pulumi.Output<string | undefined>;

    /**
     * Create a TargetGroupAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetGroupAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetGroupAttachmentArgs | TargetGroupAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TargetGroupAttachmentState | undefined;
            resourceInputs["clbId"] = state ? state.clbId : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["targetGroupId"] = state ? state.targetGroupId : undefined;
            resourceInputs["targrtGroupId"] = state ? state.targrtGroupId : undefined;
        } else {
            const args = argsOrState as TargetGroupAttachmentArgs | undefined;
            if ((!args || args.clbId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clbId'");
            }
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["clbId"] = args ? args.clbId : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["targetGroupId"] = args ? args.targetGroupId : undefined;
            resourceInputs["targrtGroupId"] = args ? args.targrtGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TargetGroupAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetGroupAttachment resources.
 */
export interface TargetGroupAttachmentState {
    /**
     * ID of the CLB.
     */
    clbId?: pulumi.Input<string>;
    /**
     * ID of the CLB listener.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * ID of the CLB listener rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * ID of the CLB target group.
     */
    targetGroupId?: pulumi.Input<string>;
    /**
     * ID of the CLB target group.
     *
     * @deprecated It has been deprecated from version 1.47.1. Use `target_group_id` instead.
     */
    targrtGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetGroupAttachment resource.
 */
export interface TargetGroupAttachmentArgs {
    /**
     * ID of the CLB.
     */
    clbId: pulumi.Input<string>;
    /**
     * ID of the CLB listener.
     */
    listenerId: pulumi.Input<string>;
    /**
     * ID of the CLB listener rule.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * ID of the CLB target group.
     */
    targetGroupId?: pulumi.Input<string>;
    /**
     * ID of the CLB target group.
     *
     * @deprecated It has been deprecated from version 1.47.1. Use `target_group_id` instead.
     */
    targrtGroupId?: pulumi.Input<string>;
}
