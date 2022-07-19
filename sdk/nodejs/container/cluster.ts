// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tctest:Container/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The network bandwidth of the node.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * The network type of the node.
     */
    public readonly bandwidthType!: pulumi.Output<string>;
    /**
     * The CIDR which the cluster is going to use.
     */
    public readonly clusterCidr!: pulumi.Output<string>;
    /**
     * The description of the cluster.
     */
    public readonly clusterDesc!: pulumi.Output<string | undefined>;
    /**
     * The name of the cluster.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * The kubernetes version of the cluster.
     */
    public readonly clusterVersion!: pulumi.Output<string | undefined>;
    /**
     * The cpu of the node.
     *
     * @deprecated It has been deprecated from version 1.16.0. Set 'instance_type' instead.
     */
    public readonly cpu!: pulumi.Output<number | undefined>;
    /**
     * The type of node needed by cvm.
     */
    public readonly cvmType!: pulumi.Output<string | undefined>;
    /**
     * The docker graph path is going to mounted.
     */
    public readonly dockerGraphPath!: pulumi.Output<string | undefined>;
    /**
     * The node number is going to create in the cluster.
     */
    public readonly goodsNum!: pulumi.Output<number>;
    /**
     * The name ot node.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * The instance type of the node needed by cvm.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Describe whether the node enable the gateway capability.
     */
    public readonly isVpcGateway!: pulumi.Output<number>;
    /**
     * The key_id of each node(if using key pair to access).
     */
    public readonly keyId!: pulumi.Output<string | undefined>;
    /**
     * The kubernetes version of the cluster.
     */
    public /*out*/ readonly kubernetesVersion!: pulumi.Output<string>;
    /**
     * The memory of the node.
     *
     * @deprecated It has been deprecated from version 1.16.0. Set 'instance_type' instead.
     */
    public readonly mem!: pulumi.Output<number | undefined>;
    /**
     * The path which volume is going to be mounted.
     */
    public readonly mountTarget!: pulumi.Output<string | undefined>;
    /**
     * The node number of the cluster.
     */
    public /*out*/ readonly nodesNum!: pulumi.Output<number>;
    /**
     * The node status of the cluster.
     */
    public /*out*/ readonly nodesStatus!: pulumi.Output<string>;
    /**
     * The system os name of the node.
     */
    public readonly osName!: pulumi.Output<string>;
    /**
     * The password of each node.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The puchase duration of the node needed by cvm.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * Indicate whether wan ip is needed.
     */
    public readonly requireWanIp!: pulumi.Output<number | undefined>;
    /**
     * The size of the root volume.
     */
    public readonly rootSize!: pulumi.Output<number>;
    /**
     * The type of the root volume. see more from CVM.
     */
    public readonly rootType!: pulumi.Output<string | undefined>;
    /**
     * The security group id.
     */
    public readonly sgId!: pulumi.Output<string | undefined>;
    /**
     * The size of the data volume.
     */
    public readonly storageSize!: pulumi.Output<number>;
    /**
     * The type of the data volume. see more from CVM.
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * The subnet id which the node stays in.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The total cpu of the cluster.
     */
    public /*out*/ readonly totalCpu!: pulumi.Output<number>;
    /**
     * The total memory of the cluster.
     */
    public /*out*/ readonly totalMem!: pulumi.Output<number>;
    /**
     * Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.
     */
    public readonly unschedulable!: pulumi.Output<number | undefined>;
    /**
     * User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from
     * CCS api documents.
     */
    public readonly userScript!: pulumi.Output<string | undefined>;
    /**
     * Specify vpc which the node(s) stay in.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The zone which the node stays in.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["bandwidthType"] = state ? state.bandwidthType : undefined;
            resourceInputs["clusterCidr"] = state ? state.clusterCidr : undefined;
            resourceInputs["clusterDesc"] = state ? state.clusterDesc : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["clusterVersion"] = state ? state.clusterVersion : undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["cvmType"] = state ? state.cvmType : undefined;
            resourceInputs["dockerGraphPath"] = state ? state.dockerGraphPath : undefined;
            resourceInputs["goodsNum"] = state ? state.goodsNum : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["isVpcGateway"] = state ? state.isVpcGateway : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
            resourceInputs["kubernetesVersion"] = state ? state.kubernetesVersion : undefined;
            resourceInputs["mem"] = state ? state.mem : undefined;
            resourceInputs["mountTarget"] = state ? state.mountTarget : undefined;
            resourceInputs["nodesNum"] = state ? state.nodesNum : undefined;
            resourceInputs["nodesStatus"] = state ? state.nodesStatus : undefined;
            resourceInputs["osName"] = state ? state.osName : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["requireWanIp"] = state ? state.requireWanIp : undefined;
            resourceInputs["rootSize"] = state ? state.rootSize : undefined;
            resourceInputs["rootType"] = state ? state.rootType : undefined;
            resourceInputs["sgId"] = state ? state.sgId : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["totalCpu"] = state ? state.totalCpu : undefined;
            resourceInputs["totalMem"] = state ? state.totalMem : undefined;
            resourceInputs["unschedulable"] = state ? state.unschedulable : undefined;
            resourceInputs["userScript"] = state ? state.userScript : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.bandwidthType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthType'");
            }
            if ((!args || args.clusterCidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterCidr'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.goodsNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'goodsNum'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.isVpcGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isVpcGateway'");
            }
            if ((!args || args.osName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'osName'");
            }
            if ((!args || args.rootSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rootSize'");
            }
            if ((!args || args.storageSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSize'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["bandwidthType"] = args ? args.bandwidthType : undefined;
            resourceInputs["clusterCidr"] = args ? args.clusterCidr : undefined;
            resourceInputs["clusterDesc"] = args ? args.clusterDesc : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["clusterVersion"] = args ? args.clusterVersion : undefined;
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["cvmType"] = args ? args.cvmType : undefined;
            resourceInputs["dockerGraphPath"] = args ? args.dockerGraphPath : undefined;
            resourceInputs["goodsNum"] = args ? args.goodsNum : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["isVpcGateway"] = args ? args.isVpcGateway : undefined;
            resourceInputs["keyId"] = args ? args.keyId : undefined;
            resourceInputs["mem"] = args ? args.mem : undefined;
            resourceInputs["mountTarget"] = args ? args.mountTarget : undefined;
            resourceInputs["osName"] = args ? args.osName : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["requireWanIp"] = args ? args.requireWanIp : undefined;
            resourceInputs["rootSize"] = args ? args.rootSize : undefined;
            resourceInputs["rootType"] = args ? args.rootType : undefined;
            resourceInputs["sgId"] = args ? args.sgId : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["unschedulable"] = args ? args.unschedulable : undefined;
            resourceInputs["userScript"] = args ? args.userScript : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["kubernetesVersion"] = undefined /*out*/;
            resourceInputs["nodesNum"] = undefined /*out*/;
            resourceInputs["nodesStatus"] = undefined /*out*/;
            resourceInputs["totalCpu"] = undefined /*out*/;
            resourceInputs["totalMem"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The network bandwidth of the node.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * The network type of the node.
     */
    bandwidthType?: pulumi.Input<string>;
    /**
     * The CIDR which the cluster is going to use.
     */
    clusterCidr?: pulumi.Input<string>;
    /**
     * The description of the cluster.
     */
    clusterDesc?: pulumi.Input<string>;
    /**
     * The name of the cluster.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * The kubernetes version of the cluster.
     */
    clusterVersion?: pulumi.Input<string>;
    /**
     * The cpu of the node.
     *
     * @deprecated It has been deprecated from version 1.16.0. Set 'instance_type' instead.
     */
    cpu?: pulumi.Input<number>;
    /**
     * The type of node needed by cvm.
     */
    cvmType?: pulumi.Input<string>;
    /**
     * The docker graph path is going to mounted.
     */
    dockerGraphPath?: pulumi.Input<string>;
    /**
     * The node number is going to create in the cluster.
     */
    goodsNum?: pulumi.Input<number>;
    /**
     * The name ot node.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The instance type of the node needed by cvm.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Describe whether the node enable the gateway capability.
     */
    isVpcGateway?: pulumi.Input<number>;
    /**
     * The key_id of each node(if using key pair to access).
     */
    keyId?: pulumi.Input<string>;
    /**
     * The kubernetes version of the cluster.
     */
    kubernetesVersion?: pulumi.Input<string>;
    /**
     * The memory of the node.
     *
     * @deprecated It has been deprecated from version 1.16.0. Set 'instance_type' instead.
     */
    mem?: pulumi.Input<number>;
    /**
     * The path which volume is going to be mounted.
     */
    mountTarget?: pulumi.Input<string>;
    /**
     * The node number of the cluster.
     */
    nodesNum?: pulumi.Input<number>;
    /**
     * The node status of the cluster.
     */
    nodesStatus?: pulumi.Input<string>;
    /**
     * The system os name of the node.
     */
    osName?: pulumi.Input<string>;
    /**
     * The password of each node.
     */
    password?: pulumi.Input<string>;
    /**
     * The puchase duration of the node needed by cvm.
     */
    period?: pulumi.Input<number>;
    /**
     * Indicate whether wan ip is needed.
     */
    requireWanIp?: pulumi.Input<number>;
    /**
     * The size of the root volume.
     */
    rootSize?: pulumi.Input<number>;
    /**
     * The type of the root volume. see more from CVM.
     */
    rootType?: pulumi.Input<string>;
    /**
     * The security group id.
     */
    sgId?: pulumi.Input<string>;
    /**
     * The size of the data volume.
     */
    storageSize?: pulumi.Input<number>;
    /**
     * The type of the data volume. see more from CVM.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The subnet id which the node stays in.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The total cpu of the cluster.
     */
    totalCpu?: pulumi.Input<number>;
    /**
     * The total memory of the cluster.
     */
    totalMem?: pulumi.Input<number>;
    /**
     * Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from
     * CCS api documents.
     */
    userScript?: pulumi.Input<string>;
    /**
     * Specify vpc which the node(s) stay in.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The zone which the node stays in.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The network bandwidth of the node.
     */
    bandwidth: pulumi.Input<number>;
    /**
     * The network type of the node.
     */
    bandwidthType: pulumi.Input<string>;
    /**
     * The CIDR which the cluster is going to use.
     */
    clusterCidr: pulumi.Input<string>;
    /**
     * The description of the cluster.
     */
    clusterDesc?: pulumi.Input<string>;
    /**
     * The name of the cluster.
     */
    clusterName: pulumi.Input<string>;
    /**
     * The kubernetes version of the cluster.
     */
    clusterVersion?: pulumi.Input<string>;
    /**
     * The cpu of the node.
     *
     * @deprecated It has been deprecated from version 1.16.0. Set 'instance_type' instead.
     */
    cpu?: pulumi.Input<number>;
    /**
     * The type of node needed by cvm.
     */
    cvmType?: pulumi.Input<string>;
    /**
     * The docker graph path is going to mounted.
     */
    dockerGraphPath?: pulumi.Input<string>;
    /**
     * The node number is going to create in the cluster.
     */
    goodsNum: pulumi.Input<number>;
    /**
     * The name ot node.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The instance type of the node needed by cvm.
     */
    instanceType: pulumi.Input<string>;
    /**
     * Describe whether the node enable the gateway capability.
     */
    isVpcGateway: pulumi.Input<number>;
    /**
     * The key_id of each node(if using key pair to access).
     */
    keyId?: pulumi.Input<string>;
    /**
     * The memory of the node.
     *
     * @deprecated It has been deprecated from version 1.16.0. Set 'instance_type' instead.
     */
    mem?: pulumi.Input<number>;
    /**
     * The path which volume is going to be mounted.
     */
    mountTarget?: pulumi.Input<string>;
    /**
     * The system os name of the node.
     */
    osName: pulumi.Input<string>;
    /**
     * The password of each node.
     */
    password?: pulumi.Input<string>;
    /**
     * The puchase duration of the node needed by cvm.
     */
    period?: pulumi.Input<number>;
    /**
     * Indicate whether wan ip is needed.
     */
    requireWanIp?: pulumi.Input<number>;
    /**
     * The size of the root volume.
     */
    rootSize: pulumi.Input<number>;
    /**
     * The type of the root volume. see more from CVM.
     */
    rootType?: pulumi.Input<string>;
    /**
     * The security group id.
     */
    sgId?: pulumi.Input<string>;
    /**
     * The size of the data volume.
     */
    storageSize: pulumi.Input<number>;
    /**
     * The type of the data volume. see more from CVM.
     */
    storageType?: pulumi.Input<string>;
    /**
     * The subnet id which the node stays in.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Determine whether the node will be schedulable. 0 is the default meaning node will be schedulable. 1 for unschedulable.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * User defined script in a base64-format. The script runs after the kubernetes component is ready on node. see more from
     * CCS api documents.
     */
    userScript?: pulumi.Input<string>;
    /**
     * Specify vpc which the node(s) stay in.
     */
    vpcId: pulumi.Input<string>;
    /**
     * The zone which the node stays in.
     */
    zoneId: pulumi.Input<string>;
}
