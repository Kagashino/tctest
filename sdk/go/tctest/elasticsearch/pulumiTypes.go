// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceMultiZoneInfo struct {
	AvailabilityZone string `pulumi:"availabilityZone"`
	SubnetId         string `pulumi:"subnetId"`
}

// InstanceMultiZoneInfoInput is an input type that accepts InstanceMultiZoneInfoArgs and InstanceMultiZoneInfoOutput values.
// You can construct a concrete instance of `InstanceMultiZoneInfoInput` via:
//
//          InstanceMultiZoneInfoArgs{...}
type InstanceMultiZoneInfoInput interface {
	pulumi.Input

	ToInstanceMultiZoneInfoOutput() InstanceMultiZoneInfoOutput
	ToInstanceMultiZoneInfoOutputWithContext(context.Context) InstanceMultiZoneInfoOutput
}

type InstanceMultiZoneInfoArgs struct {
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	SubnetId         pulumi.StringInput `pulumi:"subnetId"`
}

func (InstanceMultiZoneInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMultiZoneInfo)(nil)).Elem()
}

func (i InstanceMultiZoneInfoArgs) ToInstanceMultiZoneInfoOutput() InstanceMultiZoneInfoOutput {
	return i.ToInstanceMultiZoneInfoOutputWithContext(context.Background())
}

func (i InstanceMultiZoneInfoArgs) ToInstanceMultiZoneInfoOutputWithContext(ctx context.Context) InstanceMultiZoneInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMultiZoneInfoOutput)
}

// InstanceMultiZoneInfoArrayInput is an input type that accepts InstanceMultiZoneInfoArray and InstanceMultiZoneInfoArrayOutput values.
// You can construct a concrete instance of `InstanceMultiZoneInfoArrayInput` via:
//
//          InstanceMultiZoneInfoArray{ InstanceMultiZoneInfoArgs{...} }
type InstanceMultiZoneInfoArrayInput interface {
	pulumi.Input

	ToInstanceMultiZoneInfoArrayOutput() InstanceMultiZoneInfoArrayOutput
	ToInstanceMultiZoneInfoArrayOutputWithContext(context.Context) InstanceMultiZoneInfoArrayOutput
}

type InstanceMultiZoneInfoArray []InstanceMultiZoneInfoInput

func (InstanceMultiZoneInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMultiZoneInfo)(nil)).Elem()
}

func (i InstanceMultiZoneInfoArray) ToInstanceMultiZoneInfoArrayOutput() InstanceMultiZoneInfoArrayOutput {
	return i.ToInstanceMultiZoneInfoArrayOutputWithContext(context.Background())
}

func (i InstanceMultiZoneInfoArray) ToInstanceMultiZoneInfoArrayOutputWithContext(ctx context.Context) InstanceMultiZoneInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMultiZoneInfoArrayOutput)
}

type InstanceMultiZoneInfoOutput struct{ *pulumi.OutputState }

func (InstanceMultiZoneInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceMultiZoneInfo)(nil)).Elem()
}

func (o InstanceMultiZoneInfoOutput) ToInstanceMultiZoneInfoOutput() InstanceMultiZoneInfoOutput {
	return o
}

func (o InstanceMultiZoneInfoOutput) ToInstanceMultiZoneInfoOutputWithContext(ctx context.Context) InstanceMultiZoneInfoOutput {
	return o
}

func (o InstanceMultiZoneInfoOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceMultiZoneInfo) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o InstanceMultiZoneInfoOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceMultiZoneInfo) string { return v.SubnetId }).(pulumi.StringOutput)
}

type InstanceMultiZoneInfoArrayOutput struct{ *pulumi.OutputState }

func (InstanceMultiZoneInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceMultiZoneInfo)(nil)).Elem()
}

func (o InstanceMultiZoneInfoArrayOutput) ToInstanceMultiZoneInfoArrayOutput() InstanceMultiZoneInfoArrayOutput {
	return o
}

func (o InstanceMultiZoneInfoArrayOutput) ToInstanceMultiZoneInfoArrayOutputWithContext(ctx context.Context) InstanceMultiZoneInfoArrayOutput {
	return o
}

func (o InstanceMultiZoneInfoArrayOutput) Index(i pulumi.IntInput) InstanceMultiZoneInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceMultiZoneInfo {
		return vs[0].([]InstanceMultiZoneInfo)[vs[1].(int)]
	}).(InstanceMultiZoneInfoOutput)
}

type InstanceNodeInfoList struct {
	DiskSize *int    `pulumi:"diskSize"`
	DiskType *string `pulumi:"diskType"`
	Encrypt  *bool   `pulumi:"encrypt"`
	NodeNum  int     `pulumi:"nodeNum"`
	NodeType string  `pulumi:"nodeType"`
	Type     *string `pulumi:"type"`
}

// InstanceNodeInfoListInput is an input type that accepts InstanceNodeInfoListArgs and InstanceNodeInfoListOutput values.
// You can construct a concrete instance of `InstanceNodeInfoListInput` via:
//
//          InstanceNodeInfoListArgs{...}
type InstanceNodeInfoListInput interface {
	pulumi.Input

	ToInstanceNodeInfoListOutput() InstanceNodeInfoListOutput
	ToInstanceNodeInfoListOutputWithContext(context.Context) InstanceNodeInfoListOutput
}

type InstanceNodeInfoListArgs struct {
	DiskSize pulumi.IntPtrInput    `pulumi:"diskSize"`
	DiskType pulumi.StringPtrInput `pulumi:"diskType"`
	Encrypt  pulumi.BoolPtrInput   `pulumi:"encrypt"`
	NodeNum  pulumi.IntInput       `pulumi:"nodeNum"`
	NodeType pulumi.StringInput    `pulumi:"nodeType"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (InstanceNodeInfoListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNodeInfoList)(nil)).Elem()
}

func (i InstanceNodeInfoListArgs) ToInstanceNodeInfoListOutput() InstanceNodeInfoListOutput {
	return i.ToInstanceNodeInfoListOutputWithContext(context.Background())
}

func (i InstanceNodeInfoListArgs) ToInstanceNodeInfoListOutputWithContext(ctx context.Context) InstanceNodeInfoListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNodeInfoListOutput)
}

// InstanceNodeInfoListArrayInput is an input type that accepts InstanceNodeInfoListArray and InstanceNodeInfoListArrayOutput values.
// You can construct a concrete instance of `InstanceNodeInfoListArrayInput` via:
//
//          InstanceNodeInfoListArray{ InstanceNodeInfoListArgs{...} }
type InstanceNodeInfoListArrayInput interface {
	pulumi.Input

	ToInstanceNodeInfoListArrayOutput() InstanceNodeInfoListArrayOutput
	ToInstanceNodeInfoListArrayOutputWithContext(context.Context) InstanceNodeInfoListArrayOutput
}

type InstanceNodeInfoListArray []InstanceNodeInfoListInput

func (InstanceNodeInfoListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceNodeInfoList)(nil)).Elem()
}

func (i InstanceNodeInfoListArray) ToInstanceNodeInfoListArrayOutput() InstanceNodeInfoListArrayOutput {
	return i.ToInstanceNodeInfoListArrayOutputWithContext(context.Background())
}

func (i InstanceNodeInfoListArray) ToInstanceNodeInfoListArrayOutputWithContext(ctx context.Context) InstanceNodeInfoListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceNodeInfoListArrayOutput)
}

type InstanceNodeInfoListOutput struct{ *pulumi.OutputState }

func (InstanceNodeInfoListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceNodeInfoList)(nil)).Elem()
}

func (o InstanceNodeInfoListOutput) ToInstanceNodeInfoListOutput() InstanceNodeInfoListOutput {
	return o
}

func (o InstanceNodeInfoListOutput) ToInstanceNodeInfoListOutputWithContext(ctx context.Context) InstanceNodeInfoListOutput {
	return o
}

func (o InstanceNodeInfoListOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InstanceNodeInfoList) *int { return v.DiskSize }).(pulumi.IntPtrOutput)
}

func (o InstanceNodeInfoListOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNodeInfoList) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o InstanceNodeInfoListOutput) Encrypt() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InstanceNodeInfoList) *bool { return v.Encrypt }).(pulumi.BoolPtrOutput)
}

func (o InstanceNodeInfoListOutput) NodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceNodeInfoList) int { return v.NodeNum }).(pulumi.IntOutput)
}

func (o InstanceNodeInfoListOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceNodeInfoList) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o InstanceNodeInfoListOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceNodeInfoList) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type InstanceNodeInfoListArrayOutput struct{ *pulumi.OutputState }

func (InstanceNodeInfoListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceNodeInfoList)(nil)).Elem()
}

func (o InstanceNodeInfoListArrayOutput) ToInstanceNodeInfoListArrayOutput() InstanceNodeInfoListArrayOutput {
	return o
}

func (o InstanceNodeInfoListArrayOutput) ToInstanceNodeInfoListArrayOutputWithContext(ctx context.Context) InstanceNodeInfoListArrayOutput {
	return o
}

func (o InstanceNodeInfoListArrayOutput) Index(i pulumi.IntInput) InstanceNodeInfoListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceNodeInfoList {
		return vs[0].([]InstanceNodeInfoList)[vs[1].(int)]
	}).(InstanceNodeInfoListOutput)
}

type InstanceWebNodeTypeInfo struct {
	NodeNum  int    `pulumi:"nodeNum"`
	NodeType string `pulumi:"nodeType"`
}

// InstanceWebNodeTypeInfoInput is an input type that accepts InstanceWebNodeTypeInfoArgs and InstanceWebNodeTypeInfoOutput values.
// You can construct a concrete instance of `InstanceWebNodeTypeInfoInput` via:
//
//          InstanceWebNodeTypeInfoArgs{...}
type InstanceWebNodeTypeInfoInput interface {
	pulumi.Input

	ToInstanceWebNodeTypeInfoOutput() InstanceWebNodeTypeInfoOutput
	ToInstanceWebNodeTypeInfoOutputWithContext(context.Context) InstanceWebNodeTypeInfoOutput
}

type InstanceWebNodeTypeInfoArgs struct {
	NodeNum  pulumi.IntInput    `pulumi:"nodeNum"`
	NodeType pulumi.StringInput `pulumi:"nodeType"`
}

func (InstanceWebNodeTypeInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceWebNodeTypeInfo)(nil)).Elem()
}

func (i InstanceWebNodeTypeInfoArgs) ToInstanceWebNodeTypeInfoOutput() InstanceWebNodeTypeInfoOutput {
	return i.ToInstanceWebNodeTypeInfoOutputWithContext(context.Background())
}

func (i InstanceWebNodeTypeInfoArgs) ToInstanceWebNodeTypeInfoOutputWithContext(ctx context.Context) InstanceWebNodeTypeInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceWebNodeTypeInfoOutput)
}

// InstanceWebNodeTypeInfoArrayInput is an input type that accepts InstanceWebNodeTypeInfoArray and InstanceWebNodeTypeInfoArrayOutput values.
// You can construct a concrete instance of `InstanceWebNodeTypeInfoArrayInput` via:
//
//          InstanceWebNodeTypeInfoArray{ InstanceWebNodeTypeInfoArgs{...} }
type InstanceWebNodeTypeInfoArrayInput interface {
	pulumi.Input

	ToInstanceWebNodeTypeInfoArrayOutput() InstanceWebNodeTypeInfoArrayOutput
	ToInstanceWebNodeTypeInfoArrayOutputWithContext(context.Context) InstanceWebNodeTypeInfoArrayOutput
}

type InstanceWebNodeTypeInfoArray []InstanceWebNodeTypeInfoInput

func (InstanceWebNodeTypeInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceWebNodeTypeInfo)(nil)).Elem()
}

func (i InstanceWebNodeTypeInfoArray) ToInstanceWebNodeTypeInfoArrayOutput() InstanceWebNodeTypeInfoArrayOutput {
	return i.ToInstanceWebNodeTypeInfoArrayOutputWithContext(context.Background())
}

func (i InstanceWebNodeTypeInfoArray) ToInstanceWebNodeTypeInfoArrayOutputWithContext(ctx context.Context) InstanceWebNodeTypeInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceWebNodeTypeInfoArrayOutput)
}

type InstanceWebNodeTypeInfoOutput struct{ *pulumi.OutputState }

func (InstanceWebNodeTypeInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceWebNodeTypeInfo)(nil)).Elem()
}

func (o InstanceWebNodeTypeInfoOutput) ToInstanceWebNodeTypeInfoOutput() InstanceWebNodeTypeInfoOutput {
	return o
}

func (o InstanceWebNodeTypeInfoOutput) ToInstanceWebNodeTypeInfoOutputWithContext(ctx context.Context) InstanceWebNodeTypeInfoOutput {
	return o
}

func (o InstanceWebNodeTypeInfoOutput) NodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceWebNodeTypeInfo) int { return v.NodeNum }).(pulumi.IntOutput)
}

func (o InstanceWebNodeTypeInfoOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceWebNodeTypeInfo) string { return v.NodeType }).(pulumi.StringOutput)
}

type InstanceWebNodeTypeInfoArrayOutput struct{ *pulumi.OutputState }

func (InstanceWebNodeTypeInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceWebNodeTypeInfo)(nil)).Elem()
}

func (o InstanceWebNodeTypeInfoArrayOutput) ToInstanceWebNodeTypeInfoArrayOutput() InstanceWebNodeTypeInfoArrayOutput {
	return o
}

func (o InstanceWebNodeTypeInfoArrayOutput) ToInstanceWebNodeTypeInfoArrayOutputWithContext(ctx context.Context) InstanceWebNodeTypeInfoArrayOutput {
	return o
}

func (o InstanceWebNodeTypeInfoArrayOutput) Index(i pulumi.IntInput) InstanceWebNodeTypeInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceWebNodeTypeInfo {
		return vs[0].([]InstanceWebNodeTypeInfo)[vs[1].(int)]
	}).(InstanceWebNodeTypeInfoOutput)
}

type InstancesInstanceList struct {
	AvailabilityZone    string                               `pulumi:"availabilityZone"`
	BasicSecurityType   int                                  `pulumi:"basicSecurityType"`
	ChargeType          string                               `pulumi:"chargeType"`
	CreateTime          string                               `pulumi:"createTime"`
	DeployMode          int                                  `pulumi:"deployMode"`
	ElasticsearchDomain string                               `pulumi:"elasticsearchDomain"`
	ElasticsearchPort   int                                  `pulumi:"elasticsearchPort"`
	ElasticsearchVip    string                               `pulumi:"elasticsearchVip"`
	InstanceId          string                               `pulumi:"instanceId"`
	InstanceName        string                               `pulumi:"instanceName"`
	KibanaUrl           string                               `pulumi:"kibanaUrl"`
	LicenseType         string                               `pulumi:"licenseType"`
	MultiZoneInfos      []InstancesInstanceListMultiZoneInfo `pulumi:"multiZoneInfos"`
	NodeInfoLists       []InstancesInstanceListNodeInfoList  `pulumi:"nodeInfoLists"`
	SubnetId            string                               `pulumi:"subnetId"`
	Tags                map[string]interface{}               `pulumi:"tags"`
	Version             string                               `pulumi:"version"`
	VpcId               string                               `pulumi:"vpcId"`
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
	AvailabilityZone    pulumi.StringInput                           `pulumi:"availabilityZone"`
	BasicSecurityType   pulumi.IntInput                              `pulumi:"basicSecurityType"`
	ChargeType          pulumi.StringInput                           `pulumi:"chargeType"`
	CreateTime          pulumi.StringInput                           `pulumi:"createTime"`
	DeployMode          pulumi.IntInput                              `pulumi:"deployMode"`
	ElasticsearchDomain pulumi.StringInput                           `pulumi:"elasticsearchDomain"`
	ElasticsearchPort   pulumi.IntInput                              `pulumi:"elasticsearchPort"`
	ElasticsearchVip    pulumi.StringInput                           `pulumi:"elasticsearchVip"`
	InstanceId          pulumi.StringInput                           `pulumi:"instanceId"`
	InstanceName        pulumi.StringInput                           `pulumi:"instanceName"`
	KibanaUrl           pulumi.StringInput                           `pulumi:"kibanaUrl"`
	LicenseType         pulumi.StringInput                           `pulumi:"licenseType"`
	MultiZoneInfos      InstancesInstanceListMultiZoneInfoArrayInput `pulumi:"multiZoneInfos"`
	NodeInfoLists       InstancesInstanceListNodeInfoListArrayInput  `pulumi:"nodeInfoLists"`
	SubnetId            pulumi.StringInput                           `pulumi:"subnetId"`
	Tags                pulumi.MapInput                              `pulumi:"tags"`
	Version             pulumi.StringInput                           `pulumi:"version"`
	VpcId               pulumi.StringInput                           `pulumi:"vpcId"`
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

func (o InstancesInstanceListOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) BasicSecurityType() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.BasicSecurityType }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.ChargeType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) DeployMode() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.DeployMode }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) ElasticsearchDomain() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.ElasticsearchDomain }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) ElasticsearchPort() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.ElasticsearchPort }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) ElasticsearchVip() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.ElasticsearchVip }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) KibanaUrl() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.KibanaUrl }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.LicenseType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) MultiZoneInfos() InstancesInstanceListMultiZoneInfoArrayOutput {
	return o.ApplyT(func(v InstancesInstanceList) []InstancesInstanceListMultiZoneInfo { return v.MultiZoneInfos }).(InstancesInstanceListMultiZoneInfoArrayOutput)
}

func (o InstancesInstanceListOutput) NodeInfoLists() InstancesInstanceListNodeInfoListArrayOutput {
	return o.ApplyT(func(v InstancesInstanceList) []InstancesInstanceListNodeInfoList { return v.NodeInfoLists }).(InstancesInstanceListNodeInfoListArrayOutput)
}

func (o InstancesInstanceListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v InstancesInstanceList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func (o InstancesInstanceListOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Version }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.VpcId }).(pulumi.StringOutput)
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

type InstancesInstanceListMultiZoneInfo struct {
	AvailabilityZone string `pulumi:"availabilityZone"`
	SubnetId         string `pulumi:"subnetId"`
}

// InstancesInstanceListMultiZoneInfoInput is an input type that accepts InstancesInstanceListMultiZoneInfoArgs and InstancesInstanceListMultiZoneInfoOutput values.
// You can construct a concrete instance of `InstancesInstanceListMultiZoneInfoInput` via:
//
//          InstancesInstanceListMultiZoneInfoArgs{...}
type InstancesInstanceListMultiZoneInfoInput interface {
	pulumi.Input

	ToInstancesInstanceListMultiZoneInfoOutput() InstancesInstanceListMultiZoneInfoOutput
	ToInstancesInstanceListMultiZoneInfoOutputWithContext(context.Context) InstancesInstanceListMultiZoneInfoOutput
}

type InstancesInstanceListMultiZoneInfoArgs struct {
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	SubnetId         pulumi.StringInput `pulumi:"subnetId"`
}

func (InstancesInstanceListMultiZoneInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListMultiZoneInfo)(nil)).Elem()
}

func (i InstancesInstanceListMultiZoneInfoArgs) ToInstancesInstanceListMultiZoneInfoOutput() InstancesInstanceListMultiZoneInfoOutput {
	return i.ToInstancesInstanceListMultiZoneInfoOutputWithContext(context.Background())
}

func (i InstancesInstanceListMultiZoneInfoArgs) ToInstancesInstanceListMultiZoneInfoOutputWithContext(ctx context.Context) InstancesInstanceListMultiZoneInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListMultiZoneInfoOutput)
}

// InstancesInstanceListMultiZoneInfoArrayInput is an input type that accepts InstancesInstanceListMultiZoneInfoArray and InstancesInstanceListMultiZoneInfoArrayOutput values.
// You can construct a concrete instance of `InstancesInstanceListMultiZoneInfoArrayInput` via:
//
//          InstancesInstanceListMultiZoneInfoArray{ InstancesInstanceListMultiZoneInfoArgs{...} }
type InstancesInstanceListMultiZoneInfoArrayInput interface {
	pulumi.Input

	ToInstancesInstanceListMultiZoneInfoArrayOutput() InstancesInstanceListMultiZoneInfoArrayOutput
	ToInstancesInstanceListMultiZoneInfoArrayOutputWithContext(context.Context) InstancesInstanceListMultiZoneInfoArrayOutput
}

type InstancesInstanceListMultiZoneInfoArray []InstancesInstanceListMultiZoneInfoInput

func (InstancesInstanceListMultiZoneInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListMultiZoneInfo)(nil)).Elem()
}

func (i InstancesInstanceListMultiZoneInfoArray) ToInstancesInstanceListMultiZoneInfoArrayOutput() InstancesInstanceListMultiZoneInfoArrayOutput {
	return i.ToInstancesInstanceListMultiZoneInfoArrayOutputWithContext(context.Background())
}

func (i InstancesInstanceListMultiZoneInfoArray) ToInstancesInstanceListMultiZoneInfoArrayOutputWithContext(ctx context.Context) InstancesInstanceListMultiZoneInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListMultiZoneInfoArrayOutput)
}

type InstancesInstanceListMultiZoneInfoOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListMultiZoneInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListMultiZoneInfo)(nil)).Elem()
}

func (o InstancesInstanceListMultiZoneInfoOutput) ToInstancesInstanceListMultiZoneInfoOutput() InstancesInstanceListMultiZoneInfoOutput {
	return o
}

func (o InstancesInstanceListMultiZoneInfoOutput) ToInstancesInstanceListMultiZoneInfoOutputWithContext(ctx context.Context) InstancesInstanceListMultiZoneInfoOutput {
	return o
}

func (o InstancesInstanceListMultiZoneInfoOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListMultiZoneInfo) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o InstancesInstanceListMultiZoneInfoOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListMultiZoneInfo) string { return v.SubnetId }).(pulumi.StringOutput)
}

type InstancesInstanceListMultiZoneInfoArrayOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListMultiZoneInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListMultiZoneInfo)(nil)).Elem()
}

func (o InstancesInstanceListMultiZoneInfoArrayOutput) ToInstancesInstanceListMultiZoneInfoArrayOutput() InstancesInstanceListMultiZoneInfoArrayOutput {
	return o
}

func (o InstancesInstanceListMultiZoneInfoArrayOutput) ToInstancesInstanceListMultiZoneInfoArrayOutputWithContext(ctx context.Context) InstancesInstanceListMultiZoneInfoArrayOutput {
	return o
}

func (o InstancesInstanceListMultiZoneInfoArrayOutput) Index(i pulumi.IntInput) InstancesInstanceListMultiZoneInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesInstanceListMultiZoneInfo {
		return vs[0].([]InstancesInstanceListMultiZoneInfo)[vs[1].(int)]
	}).(InstancesInstanceListMultiZoneInfoOutput)
}

type InstancesInstanceListNodeInfoList struct {
	DiskSize int    `pulumi:"diskSize"`
	DiskType string `pulumi:"diskType"`
	Encrypt  bool   `pulumi:"encrypt"`
	NodeNum  int    `pulumi:"nodeNum"`
	NodeType string `pulumi:"nodeType"`
	Type     string `pulumi:"type"`
}

// InstancesInstanceListNodeInfoListInput is an input type that accepts InstancesInstanceListNodeInfoListArgs and InstancesInstanceListNodeInfoListOutput values.
// You can construct a concrete instance of `InstancesInstanceListNodeInfoListInput` via:
//
//          InstancesInstanceListNodeInfoListArgs{...}
type InstancesInstanceListNodeInfoListInput interface {
	pulumi.Input

	ToInstancesInstanceListNodeInfoListOutput() InstancesInstanceListNodeInfoListOutput
	ToInstancesInstanceListNodeInfoListOutputWithContext(context.Context) InstancesInstanceListNodeInfoListOutput
}

type InstancesInstanceListNodeInfoListArgs struct {
	DiskSize pulumi.IntInput    `pulumi:"diskSize"`
	DiskType pulumi.StringInput `pulumi:"diskType"`
	Encrypt  pulumi.BoolInput   `pulumi:"encrypt"`
	NodeNum  pulumi.IntInput    `pulumi:"nodeNum"`
	NodeType pulumi.StringInput `pulumi:"nodeType"`
	Type     pulumi.StringInput `pulumi:"type"`
}

func (InstancesInstanceListNodeInfoListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListNodeInfoList)(nil)).Elem()
}

func (i InstancesInstanceListNodeInfoListArgs) ToInstancesInstanceListNodeInfoListOutput() InstancesInstanceListNodeInfoListOutput {
	return i.ToInstancesInstanceListNodeInfoListOutputWithContext(context.Background())
}

func (i InstancesInstanceListNodeInfoListArgs) ToInstancesInstanceListNodeInfoListOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListNodeInfoListOutput)
}

// InstancesInstanceListNodeInfoListArrayInput is an input type that accepts InstancesInstanceListNodeInfoListArray and InstancesInstanceListNodeInfoListArrayOutput values.
// You can construct a concrete instance of `InstancesInstanceListNodeInfoListArrayInput` via:
//
//          InstancesInstanceListNodeInfoListArray{ InstancesInstanceListNodeInfoListArgs{...} }
type InstancesInstanceListNodeInfoListArrayInput interface {
	pulumi.Input

	ToInstancesInstanceListNodeInfoListArrayOutput() InstancesInstanceListNodeInfoListArrayOutput
	ToInstancesInstanceListNodeInfoListArrayOutputWithContext(context.Context) InstancesInstanceListNodeInfoListArrayOutput
}

type InstancesInstanceListNodeInfoListArray []InstancesInstanceListNodeInfoListInput

func (InstancesInstanceListNodeInfoListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListNodeInfoList)(nil)).Elem()
}

func (i InstancesInstanceListNodeInfoListArray) ToInstancesInstanceListNodeInfoListArrayOutput() InstancesInstanceListNodeInfoListArrayOutput {
	return i.ToInstancesInstanceListNodeInfoListArrayOutputWithContext(context.Background())
}

func (i InstancesInstanceListNodeInfoListArray) ToInstancesInstanceListNodeInfoListArrayOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListNodeInfoListArrayOutput)
}

type InstancesInstanceListNodeInfoListOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListNodeInfoListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListNodeInfoList)(nil)).Elem()
}

func (o InstancesInstanceListNodeInfoListOutput) ToInstancesInstanceListNodeInfoListOutput() InstancesInstanceListNodeInfoListOutput {
	return o
}

func (o InstancesInstanceListNodeInfoListOutput) ToInstancesInstanceListNodeInfoListOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoListOutput {
	return o
}

func (o InstancesInstanceListNodeInfoListOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfoList) int { return v.DiskSize }).(pulumi.IntOutput)
}

func (o InstancesInstanceListNodeInfoListOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfoList) string { return v.DiskType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListNodeInfoListOutput) Encrypt() pulumi.BoolOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfoList) bool { return v.Encrypt }).(pulumi.BoolOutput)
}

func (o InstancesInstanceListNodeInfoListOutput) NodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfoList) int { return v.NodeNum }).(pulumi.IntOutput)
}

func (o InstancesInstanceListNodeInfoListOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfoList) string { return v.NodeType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListNodeInfoListOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListNodeInfoList) string { return v.Type }).(pulumi.StringOutput)
}

type InstancesInstanceListNodeInfoListArrayOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListNodeInfoListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListNodeInfoList)(nil)).Elem()
}

func (o InstancesInstanceListNodeInfoListArrayOutput) ToInstancesInstanceListNodeInfoListArrayOutput() InstancesInstanceListNodeInfoListArrayOutput {
	return o
}

func (o InstancesInstanceListNodeInfoListArrayOutput) ToInstancesInstanceListNodeInfoListArrayOutputWithContext(ctx context.Context) InstancesInstanceListNodeInfoListArrayOutput {
	return o
}

func (o InstancesInstanceListNodeInfoListArrayOutput) Index(i pulumi.IntInput) InstancesInstanceListNodeInfoListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesInstanceListNodeInfoList {
		return vs[0].([]InstancesInstanceListNodeInfoList)[vs[1].(int)]
	}).(InstancesInstanceListNodeInfoListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMultiZoneInfoInput)(nil)).Elem(), InstanceMultiZoneInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMultiZoneInfoArrayInput)(nil)).Elem(), InstanceMultiZoneInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceNodeInfoListInput)(nil)).Elem(), InstanceNodeInfoListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceNodeInfoListArrayInput)(nil)).Elem(), InstanceNodeInfoListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceWebNodeTypeInfoInput)(nil)).Elem(), InstanceWebNodeTypeInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceWebNodeTypeInfoArrayInput)(nil)).Elem(), InstanceWebNodeTypeInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListInput)(nil)).Elem(), InstancesInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListArrayInput)(nil)).Elem(), InstancesInstanceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListMultiZoneInfoInput)(nil)).Elem(), InstancesInstanceListMultiZoneInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListMultiZoneInfoArrayInput)(nil)).Elem(), InstancesInstanceListMultiZoneInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListNodeInfoListInput)(nil)).Elem(), InstancesInstanceListNodeInfoListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListNodeInfoListArrayInput)(nil)).Elem(), InstancesInstanceListNodeInfoListArray{})
	pulumi.RegisterOutputType(InstanceMultiZoneInfoOutput{})
	pulumi.RegisterOutputType(InstanceMultiZoneInfoArrayOutput{})
	pulumi.RegisterOutputType(InstanceNodeInfoListOutput{})
	pulumi.RegisterOutputType(InstanceNodeInfoListArrayOutput{})
	pulumi.RegisterOutputType(InstanceWebNodeTypeInfoOutput{})
	pulumi.RegisterOutputType(InstanceWebNodeTypeInfoArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListMultiZoneInfoOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListMultiZoneInfoArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListNodeInfoListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListNodeInfoListArrayOutput{})
}
