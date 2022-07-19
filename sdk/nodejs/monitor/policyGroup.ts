// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class PolicyGroup extends pulumi.CustomResource {
    /**
     * Get an existing PolicyGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyGroupState, opts?: pulumi.CustomResourceOptions): PolicyGroup {
        return new PolicyGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Monitor/policyGroup:PolicyGroup';

    /**
     * Returns true if the given object is an instance of PolicyGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicyGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicyGroup.__pulumiType;
    }

    /**
     * A list binding objects(list only those in the `provider.region`). Each element contains the following attributes:
     */
    public /*out*/ readonly bindingObjects!: pulumi.Output<outputs.Monitor.PolicyGroupBindingObject[]>;
    /**
     * A list of threshold rules. Each element contains the following attributes:
     */
    public readonly conditions!: pulumi.Output<outputs.Monitor.PolicyGroupCondition[] | undefined>;
    /**
     * A list of dimensions for this policy group.
     */
    public /*out*/ readonly dimensionGroups!: pulumi.Output<string[]>;
    /**
     * A list of event rules. Each element contains the following attributes:
     */
    public readonly eventConditions!: pulumi.Output<outputs.Monitor.PolicyGroupEventCondition[] | undefined>;
    /**
     * Policy group name, length should between 1 and 20.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
     * alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
     */
    public readonly isUnionRule!: pulumi.Output<number | undefined>;
    /**
     * Recently edited user uin.
     */
    public /*out*/ readonly lastEditUin!: pulumi.Output<string>;
    /**
     * Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
     * `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
     */
    public readonly policyViewName!: pulumi.Output<string>;
    /**
     * The project id to which the policy group belongs, default is `0`.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * A list of receivers. Each element contains the following attributes:
     */
    public /*out*/ readonly receivers!: pulumi.Output<outputs.Monitor.PolicyGroupReceiver[]>;
    /**
     * Policy group's remark information.
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * Support regions this policy group.
     */
    public /*out*/ readonly supportRegions!: pulumi.Output<string[]>;
    /**
     * The policy group update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a PolicyGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyGroupArgs | PolicyGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyGroupState | undefined;
            resourceInputs["bindingObjects"] = state ? state.bindingObjects : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["dimensionGroups"] = state ? state.dimensionGroups : undefined;
            resourceInputs["eventConditions"] = state ? state.eventConditions : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["isUnionRule"] = state ? state.isUnionRule : undefined;
            resourceInputs["lastEditUin"] = state ? state.lastEditUin : undefined;
            resourceInputs["policyViewName"] = state ? state.policyViewName : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["receivers"] = state ? state.receivers : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["supportRegions"] = state ? state.supportRegions : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as PolicyGroupArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.policyViewName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyViewName'");
            }
            if ((!args || args.remark === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remark'");
            }
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["eventConditions"] = args ? args.eventConditions : undefined;
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["isUnionRule"] = args ? args.isUnionRule : undefined;
            resourceInputs["policyViewName"] = args ? args.policyViewName : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["bindingObjects"] = undefined /*out*/;
            resourceInputs["dimensionGroups"] = undefined /*out*/;
            resourceInputs["lastEditUin"] = undefined /*out*/;
            resourceInputs["receivers"] = undefined /*out*/;
            resourceInputs["supportRegions"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicyGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicyGroup resources.
 */
export interface PolicyGroupState {
    /**
     * A list binding objects(list only those in the `provider.region`). Each element contains the following attributes:
     */
    bindingObjects?: pulumi.Input<pulumi.Input<inputs.Monitor.PolicyGroupBindingObject>[]>;
    /**
     * A list of threshold rules. Each element contains the following attributes:
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.Monitor.PolicyGroupCondition>[]>;
    /**
     * A list of dimensions for this policy group.
     */
    dimensionGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of event rules. Each element contains the following attributes:
     */
    eventConditions?: pulumi.Input<pulumi.Input<inputs.Monitor.PolicyGroupEventCondition>[]>;
    /**
     * Policy group name, length should between 1 and 20.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
     * alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
     */
    isUnionRule?: pulumi.Input<number>;
    /**
     * Recently edited user uin.
     */
    lastEditUin?: pulumi.Input<string>;
    /**
     * Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
     * `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
     */
    policyViewName?: pulumi.Input<string>;
    /**
     * The project id to which the policy group belongs, default is `0`.
     */
    projectId?: pulumi.Input<number>;
    /**
     * A list of receivers. Each element contains the following attributes:
     */
    receivers?: pulumi.Input<pulumi.Input<inputs.Monitor.PolicyGroupReceiver>[]>;
    /**
     * Policy group's remark information.
     */
    remark?: pulumi.Input<string>;
    /**
     * Support regions this policy group.
     */
    supportRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The policy group update time.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PolicyGroup resource.
 */
export interface PolicyGroupArgs {
    /**
     * A list of threshold rules. Each element contains the following attributes:
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.Monitor.PolicyGroupCondition>[]>;
    /**
     * A list of event rules. Each element contains the following attributes:
     */
    eventConditions?: pulumi.Input<pulumi.Input<inputs.Monitor.PolicyGroupEventCondition>[]>;
    /**
     * Policy group name, length should between 1 and 20.
     */
    groupName: pulumi.Input<string>;
    /**
     * The and or relation of indicator alarm rule. Valid values: `0`, `1`. `0` represents or rule (if any rule is met, the
     * alarm will be raised), `1` represents and rule (if all rules are met, the alarm will be raised).The default is 0.
     */
    isUnionRule?: pulumi.Input<number>;
    /**
     * Policy view name, eg:`cvm_device`,`BANDWIDTHPACKAGE`, refer to
     * `data.tencentcloud_monitor_policy_conditions(policy_view_name)`.
     */
    policyViewName: pulumi.Input<string>;
    /**
     * The project id to which the policy group belongs, default is `0`.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Policy group's remark information.
     */
    remark: pulumi.Input<string>;
}
