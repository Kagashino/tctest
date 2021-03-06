// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class SAMLProvider extends pulumi.CustomResource {
    /**
     * Get an existing SAMLProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SAMLProviderState, opts?: pulumi.CustomResourceOptions): SAMLProvider {
        return new SAMLProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Cam/sAMLProvider:SAMLProvider';

    /**
     * Returns true if the given object is an instance of SAMLProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SAMLProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SAMLProvider.__pulumiType;
    }

    /**
     * The create time of the CAM SAML provider.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the CAM SAML provider.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The meta data document of the CAM SAML provider.
     */
    public readonly metaData!: pulumi.Output<string>;
    /**
     * Name of CAM SAML provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of the CAM SAML provider.
     */
    public /*out*/ readonly providerArn!: pulumi.Output<string>;
    /**
     * The last update time of the CAM SAML provider.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a SAMLProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SAMLProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SAMLProviderArgs | SAMLProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SAMLProviderState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["metaData"] = state ? state.metaData : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["providerArn"] = state ? state.providerArn : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as SAMLProviderArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.metaData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metaData'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["metaData"] = args ? args.metaData : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["providerArn"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SAMLProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SAMLProvider resources.
 */
export interface SAMLProviderState {
    /**
     * The create time of the CAM SAML provider.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The description of the CAM SAML provider.
     */
    description?: pulumi.Input<string>;
    /**
     * The meta data document of the CAM SAML provider.
     */
    metaData?: pulumi.Input<string>;
    /**
     * Name of CAM SAML provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the CAM SAML provider.
     */
    providerArn?: pulumi.Input<string>;
    /**
     * The last update time of the CAM SAML provider.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SAMLProvider resource.
 */
export interface SAMLProviderArgs {
    /**
     * The description of the CAM SAML provider.
     */
    description: pulumi.Input<string>;
    /**
     * The meta data document of the CAM SAML provider.
     */
    metaData: pulumi.Input<string>;
    /**
     * Name of CAM SAML provider.
     */
    name?: pulumi.Input<string>;
}
