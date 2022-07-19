// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Proxy extends pulumi.CustomResource {
    /**
     * Get an existing Proxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyState, opts?: pulumi.CustomResourceOptions): Proxy {
        return new Proxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Gaap/proxy:Proxy';

    /**
     * Returns true if the given object is an instance of Proxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Proxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Proxy.__pulumiType;
    }

    /**
     * Access region of the GAAP proxy. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`,
     * `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`,
     * `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
     */
    public readonly accessRegion!: pulumi.Output<string>;
    /**
     * Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500` and `1000`.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`,
     * `80`, `90` and `100`.
     */
    public readonly concurrent!: pulumi.Output<number>;
    /**
     * Creation time of the GAAP proxy.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Access domain of the GAAP proxy.
     */
    public /*out*/ readonly domain!: pulumi.Output<string>;
    /**
     * Indicates whether GAAP proxy is enabled, default value is `true`.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * Forwarding IP of the GAAP proxy.
     */
    public /*out*/ readonly forwardIp!: pulumi.Output<string>;
    /**
     * Access IP of the GAAP proxy.
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * Name of the GAAP proxy, the maximum length is 30.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the project within the GAAP proxy, `0` means is default project.
     */
    public readonly projectId!: pulumi.Output<number | undefined>;
    /**
     * Region of the GAAP realserver. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`,
     * `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`,
     * `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
     */
    public readonly realserverRegion!: pulumi.Output<string>;
    /**
     * Indicates whether GAAP proxy can scalable.
     */
    public /*out*/ readonly scalable!: pulumi.Output<boolean>;
    /**
     * Status of the GAAP proxy.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Supported protocols of the GAAP proxy.
     */
    public /*out*/ readonly supportProtocols!: pulumi.Output<string[]>;
    /**
     * Tags of the GAAP proxy.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Proxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyArgs | ProxyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyState | undefined;
            resourceInputs["accessRegion"] = state ? state.accessRegion : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["concurrent"] = state ? state.concurrent : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["forwardIp"] = state ? state.forwardIp : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["realserverRegion"] = state ? state.realserverRegion : undefined;
            resourceInputs["scalable"] = state ? state.scalable : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["supportProtocols"] = state ? state.supportProtocols : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ProxyArgs | undefined;
            if ((!args || args.accessRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessRegion'");
            }
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.concurrent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'concurrent'");
            }
            if ((!args || args.realserverRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realserverRegion'");
            }
            resourceInputs["accessRegion"] = args ? args.accessRegion : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["concurrent"] = args ? args.concurrent : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["realserverRegion"] = args ? args.realserverRegion : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["forwardIp"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["scalable"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["supportProtocols"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Proxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Proxy resources.
 */
export interface ProxyState {
    /**
     * Access region of the GAAP proxy. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`,
     * `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`,
     * `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
     */
    accessRegion?: pulumi.Input<string>;
    /**
     * Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500` and `1000`.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`,
     * `80`, `90` and `100`.
     */
    concurrent?: pulumi.Input<number>;
    /**
     * Creation time of the GAAP proxy.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Access domain of the GAAP proxy.
     */
    domain?: pulumi.Input<string>;
    /**
     * Indicates whether GAAP proxy is enabled, default value is `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Forwarding IP of the GAAP proxy.
     */
    forwardIp?: pulumi.Input<string>;
    /**
     * Access IP of the GAAP proxy.
     */
    ip?: pulumi.Input<string>;
    /**
     * Name of the GAAP proxy, the maximum length is 30.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the project within the GAAP proxy, `0` means is default project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Region of the GAAP realserver. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`,
     * `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`,
     * `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
     */
    realserverRegion?: pulumi.Input<string>;
    /**
     * Indicates whether GAAP proxy can scalable.
     */
    scalable?: pulumi.Input<boolean>;
    /**
     * Status of the GAAP proxy.
     */
    status?: pulumi.Input<string>;
    /**
     * Supported protocols of the GAAP proxy.
     */
    supportProtocols?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags of the GAAP proxy.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Proxy resource.
 */
export interface ProxyArgs {
    /**
     * Access region of the GAAP proxy. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`,
     * `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`,
     * `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
     */
    accessRegion: pulumi.Input<string>;
    /**
     * Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500` and `1000`.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`,
     * `80`, `90` and `100`.
     */
    concurrent: pulumi.Input<number>;
    /**
     * Indicates whether GAAP proxy is enabled, default value is `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Name of the GAAP proxy, the maximum length is 30.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the project within the GAAP proxy, `0` means is default project.
     */
    projectId?: pulumi.Input<number>;
    /**
     * Region of the GAAP realserver. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`,
     * `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`,
     * `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
     */
    realserverRegion: pulumi.Input<string>;
    /**
     * Tags of the GAAP proxy.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
