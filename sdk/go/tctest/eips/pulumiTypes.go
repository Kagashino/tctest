// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eips

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstancesEipList struct {
	CreateTime string                 `pulumi:"createTime"`
	EipId      string                 `pulumi:"eipId"`
	EipName    string                 `pulumi:"eipName"`
	EipType    string                 `pulumi:"eipType"`
	EniId      string                 `pulumi:"eniId"`
	InstanceId string                 `pulumi:"instanceId"`
	PublicIp   string                 `pulumi:"publicIp"`
	Status     string                 `pulumi:"status"`
	Tags       map[string]interface{} `pulumi:"tags"`
}

// InstancesEipListInput is an input type that accepts InstancesEipListArgs and InstancesEipListOutput values.
// You can construct a concrete instance of `InstancesEipListInput` via:
//
//          InstancesEipListArgs{...}
type InstancesEipListInput interface {
	pulumi.Input

	ToInstancesEipListOutput() InstancesEipListOutput
	ToInstancesEipListOutputWithContext(context.Context) InstancesEipListOutput
}

type InstancesEipListArgs struct {
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	EipId      pulumi.StringInput `pulumi:"eipId"`
	EipName    pulumi.StringInput `pulumi:"eipName"`
	EipType    pulumi.StringInput `pulumi:"eipType"`
	EniId      pulumi.StringInput `pulumi:"eniId"`
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	PublicIp   pulumi.StringInput `pulumi:"publicIp"`
	Status     pulumi.StringInput `pulumi:"status"`
	Tags       pulumi.MapInput    `pulumi:"tags"`
}

func (InstancesEipListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesEipList)(nil)).Elem()
}

func (i InstancesEipListArgs) ToInstancesEipListOutput() InstancesEipListOutput {
	return i.ToInstancesEipListOutputWithContext(context.Background())
}

func (i InstancesEipListArgs) ToInstancesEipListOutputWithContext(ctx context.Context) InstancesEipListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesEipListOutput)
}

// InstancesEipListArrayInput is an input type that accepts InstancesEipListArray and InstancesEipListArrayOutput values.
// You can construct a concrete instance of `InstancesEipListArrayInput` via:
//
//          InstancesEipListArray{ InstancesEipListArgs{...} }
type InstancesEipListArrayInput interface {
	pulumi.Input

	ToInstancesEipListArrayOutput() InstancesEipListArrayOutput
	ToInstancesEipListArrayOutputWithContext(context.Context) InstancesEipListArrayOutput
}

type InstancesEipListArray []InstancesEipListInput

func (InstancesEipListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesEipList)(nil)).Elem()
}

func (i InstancesEipListArray) ToInstancesEipListArrayOutput() InstancesEipListArrayOutput {
	return i.ToInstancesEipListArrayOutputWithContext(context.Background())
}

func (i InstancesEipListArray) ToInstancesEipListArrayOutputWithContext(ctx context.Context) InstancesEipListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesEipListArrayOutput)
}

type InstancesEipListOutput struct{ *pulumi.OutputState }

func (InstancesEipListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesEipList)(nil)).Elem()
}

func (o InstancesEipListOutput) ToInstancesEipListOutput() InstancesEipListOutput {
	return o
}

func (o InstancesEipListOutput) ToInstancesEipListOutputWithContext(ctx context.Context) InstancesEipListOutput {
	return o
}

func (o InstancesEipListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) EipId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.EipId }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) EipName() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.EipName }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) EipType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.EipType }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) EniId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.EniId }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.PublicIp }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesEipList) string { return v.Status }).(pulumi.StringOutput)
}

func (o InstancesEipListOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v InstancesEipList) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

type InstancesEipListArrayOutput struct{ *pulumi.OutputState }

func (InstancesEipListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesEipList)(nil)).Elem()
}

func (o InstancesEipListArrayOutput) ToInstancesEipListArrayOutput() InstancesEipListArrayOutput {
	return o
}

func (o InstancesEipListArrayOutput) ToInstancesEipListArrayOutputWithContext(ctx context.Context) InstancesEipListArrayOutput {
	return o
}

func (o InstancesEipListArrayOutput) Index(i pulumi.IntInput) InstancesEipListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesEipList {
		return vs[0].([]InstancesEipList)[vs[1].(int)]
	}).(InstancesEipListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesEipListInput)(nil)).Elem(), InstancesEipListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesEipListArrayInput)(nil)).Elem(), InstancesEipListArray{})
	pulumi.RegisterOutputType(InstancesEipListOutput{})
	pulumi.RegisterOutputType(InstancesEipListArrayOutput{})
}
