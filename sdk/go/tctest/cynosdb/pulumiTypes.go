// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterParamItem struct {
	CurrentValue string  `pulumi:"currentValue"`
	Name         string  `pulumi:"name"`
	OldValue     *string `pulumi:"oldValue"`
}

// ClusterParamItemInput is an input type that accepts ClusterParamItemArgs and ClusterParamItemOutput values.
// You can construct a concrete instance of `ClusterParamItemInput` via:
//
//          ClusterParamItemArgs{...}
type ClusterParamItemInput interface {
	pulumi.Input

	ToClusterParamItemOutput() ClusterParamItemOutput
	ToClusterParamItemOutputWithContext(context.Context) ClusterParamItemOutput
}

type ClusterParamItemArgs struct {
	CurrentValue pulumi.StringInput    `pulumi:"currentValue"`
	Name         pulumi.StringInput    `pulumi:"name"`
	OldValue     pulumi.StringPtrInput `pulumi:"oldValue"`
}

func (ClusterParamItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParamItem)(nil)).Elem()
}

func (i ClusterParamItemArgs) ToClusterParamItemOutput() ClusterParamItemOutput {
	return i.ToClusterParamItemOutputWithContext(context.Background())
}

func (i ClusterParamItemArgs) ToClusterParamItemOutputWithContext(ctx context.Context) ClusterParamItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParamItemOutput)
}

// ClusterParamItemArrayInput is an input type that accepts ClusterParamItemArray and ClusterParamItemArrayOutput values.
// You can construct a concrete instance of `ClusterParamItemArrayInput` via:
//
//          ClusterParamItemArray{ ClusterParamItemArgs{...} }
type ClusterParamItemArrayInput interface {
	pulumi.Input

	ToClusterParamItemArrayOutput() ClusterParamItemArrayOutput
	ToClusterParamItemArrayOutputWithContext(context.Context) ClusterParamItemArrayOutput
}

type ClusterParamItemArray []ClusterParamItemInput

func (ClusterParamItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParamItem)(nil)).Elem()
}

func (i ClusterParamItemArray) ToClusterParamItemArrayOutput() ClusterParamItemArrayOutput {
	return i.ToClusterParamItemArrayOutputWithContext(context.Background())
}

func (i ClusterParamItemArray) ToClusterParamItemArrayOutputWithContext(ctx context.Context) ClusterParamItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParamItemArrayOutput)
}

type ClusterParamItemOutput struct{ *pulumi.OutputState }

func (ClusterParamItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParamItem)(nil)).Elem()
}

func (o ClusterParamItemOutput) ToClusterParamItemOutput() ClusterParamItemOutput {
	return o
}

func (o ClusterParamItemOutput) ToClusterParamItemOutputWithContext(ctx context.Context) ClusterParamItemOutput {
	return o
}

func (o ClusterParamItemOutput) CurrentValue() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParamItem) string { return v.CurrentValue }).(pulumi.StringOutput)
}

func (o ClusterParamItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParamItem) string { return v.Name }).(pulumi.StringOutput)
}

func (o ClusterParamItemOutput) OldValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterParamItem) *string { return v.OldValue }).(pulumi.StringPtrOutput)
}

type ClusterParamItemArrayOutput struct{ *pulumi.OutputState }

func (ClusterParamItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParamItem)(nil)).Elem()
}

func (o ClusterParamItemArrayOutput) ToClusterParamItemArrayOutput() ClusterParamItemArrayOutput {
	return o
}

func (o ClusterParamItemArrayOutput) ToClusterParamItemArrayOutputWithContext(ctx context.Context) ClusterParamItemArrayOutput {
	return o
}

func (o ClusterParamItemArrayOutput) Index(i pulumi.IntInput) ClusterParamItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterParamItem {
		return vs[0].([]ClusterParamItem)[vs[1].(int)]
	}).(ClusterParamItemOutput)
}

type ClusterRoGroupAddr struct {
	Ip   *string `pulumi:"ip"`
	Port *int    `pulumi:"port"`
}

// ClusterRoGroupAddrInput is an input type that accepts ClusterRoGroupAddrArgs and ClusterRoGroupAddrOutput values.
// You can construct a concrete instance of `ClusterRoGroupAddrInput` via:
//
//          ClusterRoGroupAddrArgs{...}
type ClusterRoGroupAddrInput interface {
	pulumi.Input

	ToClusterRoGroupAddrOutput() ClusterRoGroupAddrOutput
	ToClusterRoGroupAddrOutputWithContext(context.Context) ClusterRoGroupAddrOutput
}

type ClusterRoGroupAddrArgs struct {
	Ip   pulumi.StringPtrInput `pulumi:"ip"`
	Port pulumi.IntPtrInput    `pulumi:"port"`
}

func (ClusterRoGroupAddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoGroupAddr)(nil)).Elem()
}

func (i ClusterRoGroupAddrArgs) ToClusterRoGroupAddrOutput() ClusterRoGroupAddrOutput {
	return i.ToClusterRoGroupAddrOutputWithContext(context.Background())
}

func (i ClusterRoGroupAddrArgs) ToClusterRoGroupAddrOutputWithContext(ctx context.Context) ClusterRoGroupAddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoGroupAddrOutput)
}

// ClusterRoGroupAddrArrayInput is an input type that accepts ClusterRoGroupAddrArray and ClusterRoGroupAddrArrayOutput values.
// You can construct a concrete instance of `ClusterRoGroupAddrArrayInput` via:
//
//          ClusterRoGroupAddrArray{ ClusterRoGroupAddrArgs{...} }
type ClusterRoGroupAddrArrayInput interface {
	pulumi.Input

	ToClusterRoGroupAddrArrayOutput() ClusterRoGroupAddrArrayOutput
	ToClusterRoGroupAddrArrayOutputWithContext(context.Context) ClusterRoGroupAddrArrayOutput
}

type ClusterRoGroupAddrArray []ClusterRoGroupAddrInput

func (ClusterRoGroupAddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRoGroupAddr)(nil)).Elem()
}

func (i ClusterRoGroupAddrArray) ToClusterRoGroupAddrArrayOutput() ClusterRoGroupAddrArrayOutput {
	return i.ToClusterRoGroupAddrArrayOutputWithContext(context.Background())
}

func (i ClusterRoGroupAddrArray) ToClusterRoGroupAddrArrayOutputWithContext(ctx context.Context) ClusterRoGroupAddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoGroupAddrArrayOutput)
}

type ClusterRoGroupAddrOutput struct{ *pulumi.OutputState }

func (ClusterRoGroupAddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoGroupAddr)(nil)).Elem()
}

func (o ClusterRoGroupAddrOutput) ToClusterRoGroupAddrOutput() ClusterRoGroupAddrOutput {
	return o
}

func (o ClusterRoGroupAddrOutput) ToClusterRoGroupAddrOutputWithContext(ctx context.Context) ClusterRoGroupAddrOutput {
	return o
}

func (o ClusterRoGroupAddrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterRoGroupAddr) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ClusterRoGroupAddrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterRoGroupAddr) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ClusterRoGroupAddrArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoGroupAddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRoGroupAddr)(nil)).Elem()
}

func (o ClusterRoGroupAddrArrayOutput) ToClusterRoGroupAddrArrayOutput() ClusterRoGroupAddrArrayOutput {
	return o
}

func (o ClusterRoGroupAddrArrayOutput) ToClusterRoGroupAddrArrayOutputWithContext(ctx context.Context) ClusterRoGroupAddrArrayOutput {
	return o
}

func (o ClusterRoGroupAddrArrayOutput) Index(i pulumi.IntInput) ClusterRoGroupAddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRoGroupAddr {
		return vs[0].([]ClusterRoGroupAddr)[vs[1].(int)]
	}).(ClusterRoGroupAddrOutput)
}

type ClusterRoGroupInstance struct {
	InstanceId   *string `pulumi:"instanceId"`
	InstanceName *string `pulumi:"instanceName"`
}

// ClusterRoGroupInstanceInput is an input type that accepts ClusterRoGroupInstanceArgs and ClusterRoGroupInstanceOutput values.
// You can construct a concrete instance of `ClusterRoGroupInstanceInput` via:
//
//          ClusterRoGroupInstanceArgs{...}
type ClusterRoGroupInstanceInput interface {
	pulumi.Input

	ToClusterRoGroupInstanceOutput() ClusterRoGroupInstanceOutput
	ToClusterRoGroupInstanceOutputWithContext(context.Context) ClusterRoGroupInstanceOutput
}

type ClusterRoGroupInstanceArgs struct {
	InstanceId   pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
}

func (ClusterRoGroupInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoGroupInstance)(nil)).Elem()
}

func (i ClusterRoGroupInstanceArgs) ToClusterRoGroupInstanceOutput() ClusterRoGroupInstanceOutput {
	return i.ToClusterRoGroupInstanceOutputWithContext(context.Background())
}

func (i ClusterRoGroupInstanceArgs) ToClusterRoGroupInstanceOutputWithContext(ctx context.Context) ClusterRoGroupInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoGroupInstanceOutput)
}

// ClusterRoGroupInstanceArrayInput is an input type that accepts ClusterRoGroupInstanceArray and ClusterRoGroupInstanceArrayOutput values.
// You can construct a concrete instance of `ClusterRoGroupInstanceArrayInput` via:
//
//          ClusterRoGroupInstanceArray{ ClusterRoGroupInstanceArgs{...} }
type ClusterRoGroupInstanceArrayInput interface {
	pulumi.Input

	ToClusterRoGroupInstanceArrayOutput() ClusterRoGroupInstanceArrayOutput
	ToClusterRoGroupInstanceArrayOutputWithContext(context.Context) ClusterRoGroupInstanceArrayOutput
}

type ClusterRoGroupInstanceArray []ClusterRoGroupInstanceInput

func (ClusterRoGroupInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRoGroupInstance)(nil)).Elem()
}

func (i ClusterRoGroupInstanceArray) ToClusterRoGroupInstanceArrayOutput() ClusterRoGroupInstanceArrayOutput {
	return i.ToClusterRoGroupInstanceArrayOutputWithContext(context.Background())
}

func (i ClusterRoGroupInstanceArray) ToClusterRoGroupInstanceArrayOutputWithContext(ctx context.Context) ClusterRoGroupInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRoGroupInstanceArrayOutput)
}

type ClusterRoGroupInstanceOutput struct{ *pulumi.OutputState }

func (ClusterRoGroupInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRoGroupInstance)(nil)).Elem()
}

func (o ClusterRoGroupInstanceOutput) ToClusterRoGroupInstanceOutput() ClusterRoGroupInstanceOutput {
	return o
}

func (o ClusterRoGroupInstanceOutput) ToClusterRoGroupInstanceOutputWithContext(ctx context.Context) ClusterRoGroupInstanceOutput {
	return o
}

func (o ClusterRoGroupInstanceOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterRoGroupInstance) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o ClusterRoGroupInstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterRoGroupInstance) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

type ClusterRoGroupInstanceArrayOutput struct{ *pulumi.OutputState }

func (ClusterRoGroupInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRoGroupInstance)(nil)).Elem()
}

func (o ClusterRoGroupInstanceArrayOutput) ToClusterRoGroupInstanceArrayOutput() ClusterRoGroupInstanceArrayOutput {
	return o
}

func (o ClusterRoGroupInstanceArrayOutput) ToClusterRoGroupInstanceArrayOutputWithContext(ctx context.Context) ClusterRoGroupInstanceArrayOutput {
	return o
}

func (o ClusterRoGroupInstanceArrayOutput) Index(i pulumi.IntInput) ClusterRoGroupInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRoGroupInstance {
		return vs[0].([]ClusterRoGroupInstance)[vs[1].(int)]
	}).(ClusterRoGroupInstanceOutput)
}

type ClusterRwGroupAddr struct {
	Ip   *string `pulumi:"ip"`
	Port *int    `pulumi:"port"`
}

// ClusterRwGroupAddrInput is an input type that accepts ClusterRwGroupAddrArgs and ClusterRwGroupAddrOutput values.
// You can construct a concrete instance of `ClusterRwGroupAddrInput` via:
//
//          ClusterRwGroupAddrArgs{...}
type ClusterRwGroupAddrInput interface {
	pulumi.Input

	ToClusterRwGroupAddrOutput() ClusterRwGroupAddrOutput
	ToClusterRwGroupAddrOutputWithContext(context.Context) ClusterRwGroupAddrOutput
}

type ClusterRwGroupAddrArgs struct {
	Ip   pulumi.StringPtrInput `pulumi:"ip"`
	Port pulumi.IntPtrInput    `pulumi:"port"`
}

func (ClusterRwGroupAddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRwGroupAddr)(nil)).Elem()
}

func (i ClusterRwGroupAddrArgs) ToClusterRwGroupAddrOutput() ClusterRwGroupAddrOutput {
	return i.ToClusterRwGroupAddrOutputWithContext(context.Background())
}

func (i ClusterRwGroupAddrArgs) ToClusterRwGroupAddrOutputWithContext(ctx context.Context) ClusterRwGroupAddrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRwGroupAddrOutput)
}

// ClusterRwGroupAddrArrayInput is an input type that accepts ClusterRwGroupAddrArray and ClusterRwGroupAddrArrayOutput values.
// You can construct a concrete instance of `ClusterRwGroupAddrArrayInput` via:
//
//          ClusterRwGroupAddrArray{ ClusterRwGroupAddrArgs{...} }
type ClusterRwGroupAddrArrayInput interface {
	pulumi.Input

	ToClusterRwGroupAddrArrayOutput() ClusterRwGroupAddrArrayOutput
	ToClusterRwGroupAddrArrayOutputWithContext(context.Context) ClusterRwGroupAddrArrayOutput
}

type ClusterRwGroupAddrArray []ClusterRwGroupAddrInput

func (ClusterRwGroupAddrArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRwGroupAddr)(nil)).Elem()
}

func (i ClusterRwGroupAddrArray) ToClusterRwGroupAddrArrayOutput() ClusterRwGroupAddrArrayOutput {
	return i.ToClusterRwGroupAddrArrayOutputWithContext(context.Background())
}

func (i ClusterRwGroupAddrArray) ToClusterRwGroupAddrArrayOutputWithContext(ctx context.Context) ClusterRwGroupAddrArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRwGroupAddrArrayOutput)
}

type ClusterRwGroupAddrOutput struct{ *pulumi.OutputState }

func (ClusterRwGroupAddrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRwGroupAddr)(nil)).Elem()
}

func (o ClusterRwGroupAddrOutput) ToClusterRwGroupAddrOutput() ClusterRwGroupAddrOutput {
	return o
}

func (o ClusterRwGroupAddrOutput) ToClusterRwGroupAddrOutputWithContext(ctx context.Context) ClusterRwGroupAddrOutput {
	return o
}

func (o ClusterRwGroupAddrOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterRwGroupAddr) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o ClusterRwGroupAddrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterRwGroupAddr) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ClusterRwGroupAddrArrayOutput struct{ *pulumi.OutputState }

func (ClusterRwGroupAddrArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRwGroupAddr)(nil)).Elem()
}

func (o ClusterRwGroupAddrArrayOutput) ToClusterRwGroupAddrArrayOutput() ClusterRwGroupAddrArrayOutput {
	return o
}

func (o ClusterRwGroupAddrArrayOutput) ToClusterRwGroupAddrArrayOutputWithContext(ctx context.Context) ClusterRwGroupAddrArrayOutput {
	return o
}

func (o ClusterRwGroupAddrArrayOutput) Index(i pulumi.IntInput) ClusterRwGroupAddrOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRwGroupAddr {
		return vs[0].([]ClusterRwGroupAddr)[vs[1].(int)]
	}).(ClusterRwGroupAddrOutput)
}

type ClusterRwGroupInstance struct {
	InstanceId   *string `pulumi:"instanceId"`
	InstanceName *string `pulumi:"instanceName"`
}

// ClusterRwGroupInstanceInput is an input type that accepts ClusterRwGroupInstanceArgs and ClusterRwGroupInstanceOutput values.
// You can construct a concrete instance of `ClusterRwGroupInstanceInput` via:
//
//          ClusterRwGroupInstanceArgs{...}
type ClusterRwGroupInstanceInput interface {
	pulumi.Input

	ToClusterRwGroupInstanceOutput() ClusterRwGroupInstanceOutput
	ToClusterRwGroupInstanceOutputWithContext(context.Context) ClusterRwGroupInstanceOutput
}

type ClusterRwGroupInstanceArgs struct {
	InstanceId   pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
}

func (ClusterRwGroupInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRwGroupInstance)(nil)).Elem()
}

func (i ClusterRwGroupInstanceArgs) ToClusterRwGroupInstanceOutput() ClusterRwGroupInstanceOutput {
	return i.ToClusterRwGroupInstanceOutputWithContext(context.Background())
}

func (i ClusterRwGroupInstanceArgs) ToClusterRwGroupInstanceOutputWithContext(ctx context.Context) ClusterRwGroupInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRwGroupInstanceOutput)
}

// ClusterRwGroupInstanceArrayInput is an input type that accepts ClusterRwGroupInstanceArray and ClusterRwGroupInstanceArrayOutput values.
// You can construct a concrete instance of `ClusterRwGroupInstanceArrayInput` via:
//
//          ClusterRwGroupInstanceArray{ ClusterRwGroupInstanceArgs{...} }
type ClusterRwGroupInstanceArrayInput interface {
	pulumi.Input

	ToClusterRwGroupInstanceArrayOutput() ClusterRwGroupInstanceArrayOutput
	ToClusterRwGroupInstanceArrayOutputWithContext(context.Context) ClusterRwGroupInstanceArrayOutput
}

type ClusterRwGroupInstanceArray []ClusterRwGroupInstanceInput

func (ClusterRwGroupInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRwGroupInstance)(nil)).Elem()
}

func (i ClusterRwGroupInstanceArray) ToClusterRwGroupInstanceArrayOutput() ClusterRwGroupInstanceArrayOutput {
	return i.ToClusterRwGroupInstanceArrayOutputWithContext(context.Background())
}

func (i ClusterRwGroupInstanceArray) ToClusterRwGroupInstanceArrayOutputWithContext(ctx context.Context) ClusterRwGroupInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterRwGroupInstanceArrayOutput)
}

type ClusterRwGroupInstanceOutput struct{ *pulumi.OutputState }

func (ClusterRwGroupInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterRwGroupInstance)(nil)).Elem()
}

func (o ClusterRwGroupInstanceOutput) ToClusterRwGroupInstanceOutput() ClusterRwGroupInstanceOutput {
	return o
}

func (o ClusterRwGroupInstanceOutput) ToClusterRwGroupInstanceOutputWithContext(ctx context.Context) ClusterRwGroupInstanceOutput {
	return o
}

func (o ClusterRwGroupInstanceOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterRwGroupInstance) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o ClusterRwGroupInstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterRwGroupInstance) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

type ClusterRwGroupInstanceArrayOutput struct{ *pulumi.OutputState }

func (ClusterRwGroupInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterRwGroupInstance)(nil)).Elem()
}

func (o ClusterRwGroupInstanceArrayOutput) ToClusterRwGroupInstanceArrayOutput() ClusterRwGroupInstanceArrayOutput {
	return o
}

func (o ClusterRwGroupInstanceArrayOutput) ToClusterRwGroupInstanceArrayOutputWithContext(ctx context.Context) ClusterRwGroupInstanceArrayOutput {
	return o
}

func (o ClusterRwGroupInstanceArrayOutput) Index(i pulumi.IntInput) ClusterRwGroupInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterRwGroupInstance {
		return vs[0].([]ClusterRwGroupInstance)[vs[1].(int)]
	}).(ClusterRwGroupInstanceOutput)
}

type ClustersClusterList struct {
	AutoRenewFlag int    `pulumi:"autoRenewFlag"`
	AvailableZone string `pulumi:"availableZone"`
	ChargeType    string `pulumi:"chargeType"`
	ClusterId     string `pulumi:"clusterId"`
	ClusterLimit  int    `pulumi:"clusterLimit"`
	ClusterName   string `pulumi:"clusterName"`
	ClusterStatus string `pulumi:"clusterStatus"`
	CreateTime    string `pulumi:"createTime"`
	DbType        string `pulumi:"dbType"`
	DbVersion     string `pulumi:"dbVersion"`
	Port          int    `pulumi:"port"`
	ProjectId     int    `pulumi:"projectId"`
	SubnetId      string `pulumi:"subnetId"`
	VpcId         string `pulumi:"vpcId"`
}

// ClustersClusterListInput is an input type that accepts ClustersClusterListArgs and ClustersClusterListOutput values.
// You can construct a concrete instance of `ClustersClusterListInput` via:
//
//          ClustersClusterListArgs{...}
type ClustersClusterListInput interface {
	pulumi.Input

	ToClustersClusterListOutput() ClustersClusterListOutput
	ToClustersClusterListOutputWithContext(context.Context) ClustersClusterListOutput
}

type ClustersClusterListArgs struct {
	AutoRenewFlag pulumi.IntInput    `pulumi:"autoRenewFlag"`
	AvailableZone pulumi.StringInput `pulumi:"availableZone"`
	ChargeType    pulumi.StringInput `pulumi:"chargeType"`
	ClusterId     pulumi.StringInput `pulumi:"clusterId"`
	ClusterLimit  pulumi.IntInput    `pulumi:"clusterLimit"`
	ClusterName   pulumi.StringInput `pulumi:"clusterName"`
	ClusterStatus pulumi.StringInput `pulumi:"clusterStatus"`
	CreateTime    pulumi.StringInput `pulumi:"createTime"`
	DbType        pulumi.StringInput `pulumi:"dbType"`
	DbVersion     pulumi.StringInput `pulumi:"dbVersion"`
	Port          pulumi.IntInput    `pulumi:"port"`
	ProjectId     pulumi.IntInput    `pulumi:"projectId"`
	SubnetId      pulumi.StringInput `pulumi:"subnetId"`
	VpcId         pulumi.StringInput `pulumi:"vpcId"`
}

func (ClustersClusterListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClustersClusterList)(nil)).Elem()
}

func (i ClustersClusterListArgs) ToClustersClusterListOutput() ClustersClusterListOutput {
	return i.ToClustersClusterListOutputWithContext(context.Background())
}

func (i ClustersClusterListArgs) ToClustersClusterListOutputWithContext(ctx context.Context) ClustersClusterListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustersClusterListOutput)
}

// ClustersClusterListArrayInput is an input type that accepts ClustersClusterListArray and ClustersClusterListArrayOutput values.
// You can construct a concrete instance of `ClustersClusterListArrayInput` via:
//
//          ClustersClusterListArray{ ClustersClusterListArgs{...} }
type ClustersClusterListArrayInput interface {
	pulumi.Input

	ToClustersClusterListArrayOutput() ClustersClusterListArrayOutput
	ToClustersClusterListArrayOutputWithContext(context.Context) ClustersClusterListArrayOutput
}

type ClustersClusterListArray []ClustersClusterListInput

func (ClustersClusterListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClustersClusterList)(nil)).Elem()
}

func (i ClustersClusterListArray) ToClustersClusterListArrayOutput() ClustersClusterListArrayOutput {
	return i.ToClustersClusterListArrayOutputWithContext(context.Background())
}

func (i ClustersClusterListArray) ToClustersClusterListArrayOutputWithContext(ctx context.Context) ClustersClusterListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustersClusterListArrayOutput)
}

type ClustersClusterListOutput struct{ *pulumi.OutputState }

func (ClustersClusterListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClustersClusterList)(nil)).Elem()
}

func (o ClustersClusterListOutput) ToClustersClusterListOutput() ClustersClusterListOutput {
	return o
}

func (o ClustersClusterListOutput) ToClustersClusterListOutputWithContext(ctx context.Context) ClustersClusterListOutput {
	return o
}

func (o ClustersClusterListOutput) AutoRenewFlag() pulumi.IntOutput {
	return o.ApplyT(func(v ClustersClusterList) int { return v.AutoRenewFlag }).(pulumi.IntOutput)
}

func (o ClustersClusterListOutput) AvailableZone() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.AvailableZone }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.ChargeType }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) ClusterLimit() pulumi.IntOutput {
	return o.ApplyT(func(v ClustersClusterList) int { return v.ClusterLimit }).(pulumi.IntOutput)
}

func (o ClustersClusterListOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.ClusterName }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) ClusterStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.ClusterStatus }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) DbType() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.DbType }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) DbVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.DbVersion }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ClustersClusterList) int { return v.Port }).(pulumi.IntOutput)
}

func (o ClustersClusterListOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v ClustersClusterList) int { return v.ProjectId }).(pulumi.IntOutput)
}

func (o ClustersClusterListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o ClustersClusterListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersClusterList) string { return v.VpcId }).(pulumi.StringOutput)
}

type ClustersClusterListArrayOutput struct{ *pulumi.OutputState }

func (ClustersClusterListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClustersClusterList)(nil)).Elem()
}

func (o ClustersClusterListArrayOutput) ToClustersClusterListArrayOutput() ClustersClusterListArrayOutput {
	return o
}

func (o ClustersClusterListArrayOutput) ToClustersClusterListArrayOutputWithContext(ctx context.Context) ClustersClusterListArrayOutput {
	return o
}

func (o ClustersClusterListArrayOutput) Index(i pulumi.IntInput) ClustersClusterListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClustersClusterList {
		return vs[0].([]ClustersClusterList)[vs[1].(int)]
	}).(ClustersClusterListOutput)
}

type InstancesInstanceList struct {
	ClusterId           *string `pulumi:"clusterId"`
	CreateTime          string  `pulumi:"createTime"`
	InstanceCpuCore     int     `pulumi:"instanceCpuCore"`
	InstanceId          *string `pulumi:"instanceId"`
	InstanceMemorySize  int     `pulumi:"instanceMemorySize"`
	InstanceName        string  `pulumi:"instanceName"`
	InstanceStatus      string  `pulumi:"instanceStatus"`
	InstanceStorageSize int     `pulumi:"instanceStorageSize"`
	InstanceType        string  `pulumi:"instanceType"`
}

// InstancesInstanceListInput is an input type that accepts InstancesInstanceListArgs and InstancesInstanceListOutput values.
// You can construct a concrete instance of `InstancesInstanceListInput` via:
//
//          InstancesInstanceListArgs{...}
type InstancesInstanceListInput interface {
	pulumi.Input

	ToInstancesInstanceListOutput() InstancesInstanceListOutput
	ToInstancesInstanceListOutputWithContext(context.Context) InstancesInstanceListOutput
}

type InstancesInstanceListArgs struct {
	ClusterId           pulumi.StringPtrInput `pulumi:"clusterId"`
	CreateTime          pulumi.StringInput    `pulumi:"createTime"`
	InstanceCpuCore     pulumi.IntInput       `pulumi:"instanceCpuCore"`
	InstanceId          pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceMemorySize  pulumi.IntInput       `pulumi:"instanceMemorySize"`
	InstanceName        pulumi.StringInput    `pulumi:"instanceName"`
	InstanceStatus      pulumi.StringInput    `pulumi:"instanceStatus"`
	InstanceStorageSize pulumi.IntInput       `pulumi:"instanceStorageSize"`
	InstanceType        pulumi.StringInput    `pulumi:"instanceType"`
}

func (InstancesInstanceListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceList)(nil)).Elem()
}

func (i InstancesInstanceListArgs) ToInstancesInstanceListOutput() InstancesInstanceListOutput {
	return i.ToInstancesInstanceListOutputWithContext(context.Background())
}

func (i InstancesInstanceListArgs) ToInstancesInstanceListOutputWithContext(ctx context.Context) InstancesInstanceListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListOutput)
}

// InstancesInstanceListArrayInput is an input type that accepts InstancesInstanceListArray and InstancesInstanceListArrayOutput values.
// You can construct a concrete instance of `InstancesInstanceListArrayInput` via:
//
//          InstancesInstanceListArray{ InstancesInstanceListArgs{...} }
type InstancesInstanceListArrayInput interface {
	pulumi.Input

	ToInstancesInstanceListArrayOutput() InstancesInstanceListArrayOutput
	ToInstancesInstanceListArrayOutputWithContext(context.Context) InstancesInstanceListArrayOutput
}

type InstancesInstanceListArray []InstancesInstanceListInput

func (InstancesInstanceListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceList)(nil)).Elem()
}

func (i InstancesInstanceListArray) ToInstancesInstanceListArrayOutput() InstancesInstanceListArrayOutput {
	return i.ToInstancesInstanceListArrayOutputWithContext(context.Background())
}

func (i InstancesInstanceListArray) ToInstancesInstanceListArrayOutputWithContext(ctx context.Context) InstancesInstanceListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListArrayOutput)
}

type InstancesInstanceListOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceList)(nil)).Elem()
}

func (o InstancesInstanceListOutput) ToInstancesInstanceListOutput() InstancesInstanceListOutput {
	return o
}

func (o InstancesInstanceListOutput) ToInstancesInstanceListOutputWithContext(ctx context.Context) InstancesInstanceListOutput {
	return o
}

func (o InstancesInstanceListOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesInstanceList) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o InstancesInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) InstanceCpuCore() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.InstanceCpuCore }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesInstanceList) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o InstancesInstanceListOutput) InstanceMemorySize() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.InstanceMemorySize }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) InstanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.InstanceStatus }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) InstanceStorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.InstanceStorageSize }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.InstanceType }).(pulumi.StringOutput)
}

type InstancesInstanceListArrayOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceList)(nil)).Elem()
}

func (o InstancesInstanceListArrayOutput) ToInstancesInstanceListArrayOutput() InstancesInstanceListArrayOutput {
	return o
}

func (o InstancesInstanceListArrayOutput) ToInstancesInstanceListArrayOutputWithContext(ctx context.Context) InstancesInstanceListArrayOutput {
	return o
}

func (o InstancesInstanceListArrayOutput) Index(i pulumi.IntInput) InstancesInstanceListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesInstanceList {
		return vs[0].([]InstancesInstanceList)[vs[1].(int)]
	}).(InstancesInstanceListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParamItemInput)(nil)).Elem(), ClusterParamItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParamItemArrayInput)(nil)).Elem(), ClusterParamItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoGroupAddrInput)(nil)).Elem(), ClusterRoGroupAddrArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoGroupAddrArrayInput)(nil)).Elem(), ClusterRoGroupAddrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoGroupInstanceInput)(nil)).Elem(), ClusterRoGroupInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRoGroupInstanceArrayInput)(nil)).Elem(), ClusterRoGroupInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRwGroupAddrInput)(nil)).Elem(), ClusterRwGroupAddrArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRwGroupAddrArrayInput)(nil)).Elem(), ClusterRwGroupAddrArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRwGroupInstanceInput)(nil)).Elem(), ClusterRwGroupInstanceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterRwGroupInstanceArrayInput)(nil)).Elem(), ClusterRwGroupInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustersClusterListInput)(nil)).Elem(), ClustersClusterListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustersClusterListArrayInput)(nil)).Elem(), ClustersClusterListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListInput)(nil)).Elem(), InstancesInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListArrayInput)(nil)).Elem(), InstancesInstanceListArray{})
	pulumi.RegisterOutputType(ClusterParamItemOutput{})
	pulumi.RegisterOutputType(ClusterParamItemArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoGroupAddrOutput{})
	pulumi.RegisterOutputType(ClusterRoGroupAddrArrayOutput{})
	pulumi.RegisterOutputType(ClusterRoGroupInstanceOutput{})
	pulumi.RegisterOutputType(ClusterRoGroupInstanceArrayOutput{})
	pulumi.RegisterOutputType(ClusterRwGroupAddrOutput{})
	pulumi.RegisterOutputType(ClusterRwGroupAddrArrayOutput{})
	pulumi.RegisterOutputType(ClusterRwGroupInstanceOutput{})
	pulumi.RegisterOutputType(ClusterRwGroupInstanceArrayOutput{})
	pulumi.RegisterOutputType(ClustersClusterListOutput{})
	pulumi.RegisterOutputType(ClustersClusterListArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListArrayOutput{})
}