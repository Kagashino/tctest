// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CustomizedConfig extends pulumi.CustomResource {
    /**
     * Get an existing CustomizedConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomizedConfigState, opts?: pulumi.CustomResourceOptions): CustomizedConfig {
        return new CustomizedConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Clb/customizedConfig:CustomizedConfig';

    /**
     * Returns true if the given object is an instance of CustomizedConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomizedConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomizedConfig.__pulumiType;
    }

    /**
     * Content of Customized Config.
     */
    public readonly configContent!: pulumi.Output<string>;
    /**
     * Name of Customized Config.
     */
    public readonly configName!: pulumi.Output<string>;
    /**
     * Create time of Customized Config.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * List of LoadBalancer Ids.
     */
    public readonly loadBalancerIds!: pulumi.Output<string[] | undefined>;
    /**
     * Update time of Customized Config.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a CustomizedConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomizedConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomizedConfigArgs | CustomizedConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomizedConfigState | undefined;
            resourceInputs["configContent"] = state ? state.configContent : undefined;
            resourceInputs["configName"] = state ? state.configName : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["loadBalancerIds"] = state ? state.loadBalancerIds : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as CustomizedConfigArgs | undefined;
            if ((!args || args.configContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configContent'");
            }
            if ((!args || args.configName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configName'");
            }
            resourceInputs["configContent"] = args ? args.configContent : undefined;
            resourceInputs["configName"] = args ? args.configName : undefined;
            resourceInputs["loadBalancerIds"] = args ? args.loadBalancerIds : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomizedConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomizedConfig resources.
 */
export interface CustomizedConfigState {
    /**
     * Content of Customized Config.
     */
    configContent?: pulumi.Input<string>;
    /**
     * Name of Customized Config.
     */
    configName?: pulumi.Input<string>;
    /**
     * Create time of Customized Config.
     */
    createTime?: pulumi.Input<string>;
    /**
     * List of LoadBalancer Ids.
     */
    loadBalancerIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Update time of Customized Config.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomizedConfig resource.
 */
export interface CustomizedConfigArgs {
    /**
     * Content of Customized Config.
     */
    configContent: pulumi.Input<string>;
    /**
     * Name of Customized Config.
     */
    configName: pulumi.Input<string>;
    /**
     * List of LoadBalancer Ids.
     */
    loadBalancerIds?: pulumi.Input<pulumi.Input<string>[]>;
}