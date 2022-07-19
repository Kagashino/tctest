// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Gateway extends pulumi.CustomResource {
    /**
     * Get an existing Gateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayState, opts?: pulumi.CustomResourceOptions): Gateway {
        return new Gateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Vpn/gateway:Gateway';

    /**
     * Returns true if the given object is an instance of Gateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Gateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Gateway.__pulumiType;
    }

    /**
     * The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include:
     * 5,10,20,50,100,200,500,1000. Default is 5. When charge type is `PREPAID`, bandwidth degradation operation is
     * unsupported.
     */
    public readonly bandwidth!: pulumi.Output<number | undefined>;
    /**
     * CDC instance ID.
     */
    public readonly cdcId!: pulumi.Output<string>;
    /**
     * Charge Type of the VPN gateway. Valid value: `PREPAID`, `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Create time of the VPN gateway.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Expired time of the VPN gateway when charge type is `PREPAID`.
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<string>;
    /**
     * Indicates whether ip address is blocked.
     */
    public /*out*/ readonly isAddressBlocked!: pulumi.Output<boolean>;
    /**
     * Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter
     * is only required for SSL VPN gateways.
     */
    public readonly maxConnection!: pulumi.Output<number>;
    /**
     * Name of the VPN gateway. The length of character is limited to 1-60.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The plan of new purchase. Valid value: `PREPAID_TO_POSTPAID`.
     */
    public /*out*/ readonly newPurchasePlan!: pulumi.Output<string>;
    /**
     * Period of instance to be prepaid. Valid value: `1`, `2`, `3`, `4`, `6`, `7`, `8`, `9`, `12`, `24`, `36`. The unit is
     * month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid
     * period. This para can only be set to take effect in create operation.
     */
    public readonly prepaidPeriod!: pulumi.Output<number | undefined>;
    /**
     * Flag indicates whether to renew or not. Valid value: `NOTIFY_AND_RENEW`, `NOTIFY_AND_AUTO_RENEW`,
     * `NOT_NOTIFY_AND_NOT_RENEW`. This para can only be set to take effect in create operation.
     */
    public readonly prepaidRenewFlag!: pulumi.Output<string | undefined>;
    /**
     * Public IP of the VPN gateway.
     */
    public /*out*/ readonly publicIpAddress!: pulumi.Output<string>;
    /**
     * Restrict state of gateway. Valid value: `PRETECIVELY_ISOLATED`, `NORMAL`.
     */
    public /*out*/ readonly restrictState!: pulumi.Output<string>;
    /**
     * State of the VPN gateway. Valid value: `PENDING`, `DELETING`, `AVAILABLE`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A list of tags used to associate different resources.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Type of gateway instance. Valid value: `IPSEC`, `SSL` and `CCN`. Note: CCN type is only for whitelist customer now.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * ID of the VPC. Required if vpn gateway is not in `CCN` type, and doesn't make sense for `CCN` vpn gateway.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;
    /**
     * Zone of the VPN gateway.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Gateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayArgs | GatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GatewayState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["cdcId"] = state ? state.cdcId : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["isAddressBlocked"] = state ? state.isAddressBlocked : undefined;
            resourceInputs["maxConnection"] = state ? state.maxConnection : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["newPurchasePlan"] = state ? state.newPurchasePlan : undefined;
            resourceInputs["prepaidPeriod"] = state ? state.prepaidPeriod : undefined;
            resourceInputs["prepaidRenewFlag"] = state ? state.prepaidRenewFlag : undefined;
            resourceInputs["publicIpAddress"] = state ? state.publicIpAddress : undefined;
            resourceInputs["restrictState"] = state ? state.restrictState : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as GatewayArgs | undefined;
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["cdcId"] = args ? args.cdcId : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["maxConnection"] = args ? args.maxConnection : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["prepaidPeriod"] = args ? args.prepaidPeriod : undefined;
            resourceInputs["prepaidRenewFlag"] = args ? args.prepaidRenewFlag : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["isAddressBlocked"] = undefined /*out*/;
            resourceInputs["newPurchasePlan"] = undefined /*out*/;
            resourceInputs["publicIpAddress"] = undefined /*out*/;
            resourceInputs["restrictState"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Gateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Gateway resources.
 */
export interface GatewayState {
    /**
     * The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include:
     * 5,10,20,50,100,200,500,1000. Default is 5. When charge type is `PREPAID`, bandwidth degradation operation is
     * unsupported.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * CDC instance ID.
     */
    cdcId?: pulumi.Input<string>;
    /**
     * Charge Type of the VPN gateway. Valid value: `PREPAID`, `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Create time of the VPN gateway.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Expired time of the VPN gateway when charge type is `PREPAID`.
     */
    expiredTime?: pulumi.Input<string>;
    /**
     * Indicates whether ip address is blocked.
     */
    isAddressBlocked?: pulumi.Input<boolean>;
    /**
     * Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter
     * is only required for SSL VPN gateways.
     */
    maxConnection?: pulumi.Input<number>;
    /**
     * Name of the VPN gateway. The length of character is limited to 1-60.
     */
    name?: pulumi.Input<string>;
    /**
     * The plan of new purchase. Valid value: `PREPAID_TO_POSTPAID`.
     */
    newPurchasePlan?: pulumi.Input<string>;
    /**
     * Period of instance to be prepaid. Valid value: `1`, `2`, `3`, `4`, `6`, `7`, `8`, `9`, `12`, `24`, `36`. The unit is
     * month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid
     * period. This para can only be set to take effect in create operation.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * Flag indicates whether to renew or not. Valid value: `NOTIFY_AND_RENEW`, `NOTIFY_AND_AUTO_RENEW`,
     * `NOT_NOTIFY_AND_NOT_RENEW`. This para can only be set to take effect in create operation.
     */
    prepaidRenewFlag?: pulumi.Input<string>;
    /**
     * Public IP of the VPN gateway.
     */
    publicIpAddress?: pulumi.Input<string>;
    /**
     * Restrict state of gateway. Valid value: `PRETECIVELY_ISOLATED`, `NORMAL`.
     */
    restrictState?: pulumi.Input<string>;
    /**
     * State of the VPN gateway. Valid value: `PENDING`, `DELETING`, `AVAILABLE`.
     */
    state?: pulumi.Input<string>;
    /**
     * A list of tags used to associate different resources.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Type of gateway instance. Valid value: `IPSEC`, `SSL` and `CCN`. Note: CCN type is only for whitelist customer now.
     */
    type?: pulumi.Input<string>;
    /**
     * ID of the VPC. Required if vpn gateway is not in `CCN` type, and doesn't make sense for `CCN` vpn gateway.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Zone of the VPN gateway.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Gateway resource.
 */
export interface GatewayArgs {
    /**
     * The maximum public network output bandwidth of VPN gateway (unit: Mbps), the available values include:
     * 5,10,20,50,100,200,500,1000. Default is 5. When charge type is `PREPAID`, bandwidth degradation operation is
     * unsupported.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * CDC instance ID.
     */
    cdcId?: pulumi.Input<string>;
    /**
     * Charge Type of the VPN gateway. Valid value: `PREPAID`, `POSTPAID_BY_HOUR`. The default is `POSTPAID_BY_HOUR`.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Maximum number of connected clients allowed for the SSL VPN gateway. Valid values: [5, 10, 20, 50, 100]. This parameter
     * is only required for SSL VPN gateways.
     */
    maxConnection?: pulumi.Input<number>;
    /**
     * Name of the VPN gateway. The length of character is limited to 1-60.
     */
    name?: pulumi.Input<string>;
    /**
     * Period of instance to be prepaid. Valid value: `1`, `2`, `3`, `4`, `6`, `7`, `8`, `9`, `12`, `24`, `36`. The unit is
     * month. Caution: when this para and renew_flag para are valid, the request means to renew several months more pre-paid
     * period. This para can only be set to take effect in create operation.
     */
    prepaidPeriod?: pulumi.Input<number>;
    /**
     * Flag indicates whether to renew or not. Valid value: `NOTIFY_AND_RENEW`, `NOTIFY_AND_AUTO_RENEW`,
     * `NOT_NOTIFY_AND_NOT_RENEW`. This para can only be set to take effect in create operation.
     */
    prepaidRenewFlag?: pulumi.Input<string>;
    /**
     * A list of tags used to associate different resources.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Type of gateway instance. Valid value: `IPSEC`, `SSL` and `CCN`. Note: CCN type is only for whitelist customer now.
     */
    type?: pulumi.Input<string>;
    /**
     * ID of the VPC. Required if vpn gateway is not in `CCN` type, and doesn't make sense for `CCN` vpn gateway.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * Zone of the VPN gateway.
     */
    zone: pulumi.Input<string>;
}
