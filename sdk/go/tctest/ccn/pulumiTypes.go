// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ccn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BandwidthLimitsLimit struct {
	BandwidthLimit int    `pulumi:"bandwidthLimit"`
	DstRegion      string `pulumi:"dstRegion"`
	Region         string `pulumi:"region"`
}

// BandwidthLimitsLimitInput is an input type that accepts BandwidthLimitsLimitArgs and BandwidthLimitsLimitOutput values.
// You can construct a concrete instance of `BandwidthLimitsLimitInput` via:
//
//          BandwidthLimitsLimitArgs{...}
type BandwidthLimitsLimitInput interface {
	pulumi.Input

	ToBandwidthLimitsLimitOutput() BandwidthLimitsLimitOutput
	ToBandwidthLimitsLimitOutputWithContext(context.Context) BandwidthLimitsLimitOutput
}

type BandwidthLimitsLimitArgs struct {
	BandwidthLimit pulumi.IntInput    `pulumi:"bandwidthLimit"`
	DstRegion      pulumi.StringInput `pulumi:"dstRegion"`
	Region         pulumi.StringInput `pulumi:"region"`
}

func (BandwidthLimitsLimitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthLimitsLimit)(nil)).Elem()
}

func (i BandwidthLimitsLimitArgs) ToBandwidthLimitsLimitOutput() BandwidthLimitsLimitOutput {
	return i.ToBandwidthLimitsLimitOutputWithContext(context.Background())
}

func (i BandwidthLimitsLimitArgs) ToBandwidthLimitsLimitOutputWithContext(ctx context.Context) BandwidthLimitsLimitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthLimitsLimitOutput)
}

// BandwidthLimitsLimitArrayInput is an input type that accepts BandwidthLimitsLimitArray and BandwidthLimitsLimitArrayOutput values.
// You can construct a concrete instance of `BandwidthLimitsLimitArrayInput` via:
//
//          BandwidthLimitsLimitArray{ BandwidthLimitsLimitArgs{...} }
type BandwidthLimitsLimitArrayInput interface {
	pulumi.Input

	ToBandwidthLimitsLimitArrayOutput() BandwidthLimitsLimitArrayOutput
	ToBandwidthLimitsLimitArrayOutputWithContext(context.Context) BandwidthLimitsLimitArrayOutput
}

type BandwidthLimitsLimitArray []BandwidthLimitsLimitInput

func (BandwidthLimitsLimitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BandwidthLimitsLimit)(nil)).Elem()
}

func (i BandwidthLimitsLimitArray) ToBandwidthLimitsLimitArrayOutput() BandwidthLimitsLimitArrayOutput {
	return i.ToBandwidthLimitsLimitArrayOutputWithContext(context.Background())
}

func (i BandwidthLimitsLimitArray) ToBandwidthLimitsLimitArrayOutputWithContext(ctx context.Context) BandwidthLimitsLimitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthLimitsLimitArrayOutput)
}

type BandwidthLimitsLimitOutput struct{ *pulumi.OutputState }

func (BandwidthLimitsLimitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthLimitsLimit)(nil)).Elem()
}

func (o BandwidthLimitsLimitOutput) ToBandwidthLimitsLimitOutput() BandwidthLimitsLimitOutput {
	return o
}

func (o BandwidthLimitsLimitOutput) ToBandwidthLimitsLimitOutputWithContext(ctx context.Context) BandwidthLimitsLimitOutput {
	return o
}

func (o BandwidthLimitsLimitOutput) BandwidthLimit() pulumi.IntOutput {
	return o.ApplyT(func(v BandwidthLimitsLimit) int { return v.BandwidthLimit }).(pulumi.IntOutput)
}

func (o BandwidthLimitsLimitOutput) DstRegion() pulumi.StringOutput {
	return o.ApplyT(func(v BandwidthLimitsLimit) string { return v.DstRegion }).(pulumi.StringOutput)
}

func (o BandwidthLimitsLimitOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v BandwidthLimitsLimit) string { return v.Region }).(pulumi.StringOutput)
}

type BandwidthLimitsLimitArrayOutput struct{ *pulumi.OutputState }

func (BandwidthLimitsLimitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BandwidthLimitsLimit)(nil)).Elem()
}

func (o BandwidthLimitsLimitArrayOutput) ToBandwidthLimitsLimitArrayOutput() BandwidthLimitsLimitArrayOutput {
	return o
}

func (o BandwidthLimitsLimitArrayOutput) ToBandwidthLimitsLimitArrayOutputWithContext(ctx context.Context) BandwidthLimitsLimitArrayOutput {
	return o
}

func (o BandwidthLimitsLimitArrayOutput) Index(i pulumi.IntInput) BandwidthLimitsLimitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BandwidthLimitsLimit {
		return vs[0].([]BandwidthLimitsLimit)[vs[1].(int)]
	}).(BandwidthLimitsLimitOutput)
}

type InstancesInstanceList struct {
	AttachmentLists    []InstancesInstanceListAttachmentList `pulumi:"attachmentLists"`
	BandwidthLimitType string                                `pulumi:"bandwidthLimitType"`
	CcnId              string                                `pulumi:"ccnId"`
	ChargeType         string                                `pulumi:"chargeType"`
	CreateTime         string                                `pulumi:"createTime"`
	Description        string                                `pulumi:"description"`
	Name               string                                `pulumi:"name"`
	Qos                string                                `pulumi:"qos"`
	State              string                                `pulumi:"state"`
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
	AttachmentLists    InstancesInstanceListAttachmentListArrayInput `pulumi:"attachmentLists"`
	BandwidthLimitType pulumi.StringInput                            `pulumi:"bandwidthLimitType"`
	CcnId              pulumi.StringInput                            `pulumi:"ccnId"`
	ChargeType         pulumi.StringInput                            `pulumi:"chargeType"`
	CreateTime         pulumi.StringInput                            `pulumi:"createTime"`
	Description        pulumi.StringInput                            `pulumi:"description"`
	Name               pulumi.StringInput                            `pulumi:"name"`
	Qos                pulumi.StringInput                            `pulumi:"qos"`
	State              pulumi.StringInput                            `pulumi:"state"`
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

func (o InstancesInstanceListOutput) AttachmentLists() InstancesInstanceListAttachmentListArrayOutput {
	return o.ApplyT(func(v InstancesInstanceList) []InstancesInstanceListAttachmentList { return v.AttachmentLists }).(InstancesInstanceListAttachmentListArrayOutput)
}

func (o InstancesInstanceListOutput) BandwidthLimitType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.BandwidthLimitType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) CcnId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CcnId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.ChargeType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Description }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Name }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Qos() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Qos }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.State }).(pulumi.StringOutput)
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

type InstancesInstanceListAttachmentList struct {
	AttachedTime   string   `pulumi:"attachedTime"`
	CidrBlocks     []string `pulumi:"cidrBlocks"`
	InstanceId     string   `pulumi:"instanceId"`
	InstanceRegion string   `pulumi:"instanceRegion"`
	InstanceType   string   `pulumi:"instanceType"`
	State          string   `pulumi:"state"`
}

// InstancesInstanceListAttachmentListInput is an input type that accepts InstancesInstanceListAttachmentListArgs and InstancesInstanceListAttachmentListOutput values.
// You can construct a concrete instance of `InstancesInstanceListAttachmentListInput` via:
//
//          InstancesInstanceListAttachmentListArgs{...}
type InstancesInstanceListAttachmentListInput interface {
	pulumi.Input

	ToInstancesInstanceListAttachmentListOutput() InstancesInstanceListAttachmentListOutput
	ToInstancesInstanceListAttachmentListOutputWithContext(context.Context) InstancesInstanceListAttachmentListOutput
}

type InstancesInstanceListAttachmentListArgs struct {
	AttachedTime   pulumi.StringInput      `pulumi:"attachedTime"`
	CidrBlocks     pulumi.StringArrayInput `pulumi:"cidrBlocks"`
	InstanceId     pulumi.StringInput      `pulumi:"instanceId"`
	InstanceRegion pulumi.StringInput      `pulumi:"instanceRegion"`
	InstanceType   pulumi.StringInput      `pulumi:"instanceType"`
	State          pulumi.StringInput      `pulumi:"state"`
}

func (InstancesInstanceListAttachmentListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListAttachmentList)(nil)).Elem()
}

func (i InstancesInstanceListAttachmentListArgs) ToInstancesInstanceListAttachmentListOutput() InstancesInstanceListAttachmentListOutput {
	return i.ToInstancesInstanceListAttachmentListOutputWithContext(context.Background())
}

func (i InstancesInstanceListAttachmentListArgs) ToInstancesInstanceListAttachmentListOutputWithContext(ctx context.Context) InstancesInstanceListAttachmentListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListAttachmentListOutput)
}

// InstancesInstanceListAttachmentListArrayInput is an input type that accepts InstancesInstanceListAttachmentListArray and InstancesInstanceListAttachmentListArrayOutput values.
// You can construct a concrete instance of `InstancesInstanceListAttachmentListArrayInput` via:
//
//          InstancesInstanceListAttachmentListArray{ InstancesInstanceListAttachmentListArgs{...} }
type InstancesInstanceListAttachmentListArrayInput interface {
	pulumi.Input

	ToInstancesInstanceListAttachmentListArrayOutput() InstancesInstanceListAttachmentListArrayOutput
	ToInstancesInstanceListAttachmentListArrayOutputWithContext(context.Context) InstancesInstanceListAttachmentListArrayOutput
}

type InstancesInstanceListAttachmentListArray []InstancesInstanceListAttachmentListInput

func (InstancesInstanceListAttachmentListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListAttachmentList)(nil)).Elem()
}

func (i InstancesInstanceListAttachmentListArray) ToInstancesInstanceListAttachmentListArrayOutput() InstancesInstanceListAttachmentListArrayOutput {
	return i.ToInstancesInstanceListAttachmentListArrayOutputWithContext(context.Background())
}

func (i InstancesInstanceListAttachmentListArray) ToInstancesInstanceListAttachmentListArrayOutputWithContext(ctx context.Context) InstancesInstanceListAttachmentListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesInstanceListAttachmentListArrayOutput)
}

type InstancesInstanceListAttachmentListOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListAttachmentListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesInstanceListAttachmentList)(nil)).Elem()
}

func (o InstancesInstanceListAttachmentListOutput) ToInstancesInstanceListAttachmentListOutput() InstancesInstanceListAttachmentListOutput {
	return o
}

func (o InstancesInstanceListAttachmentListOutput) ToInstancesInstanceListAttachmentListOutputWithContext(ctx context.Context) InstancesInstanceListAttachmentListOutput {
	return o
}

func (o InstancesInstanceListAttachmentListOutput) AttachedTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListAttachmentList) string { return v.AttachedTime }).(pulumi.StringOutput)
}

func (o InstancesInstanceListAttachmentListOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesInstanceListAttachmentList) []string { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

func (o InstancesInstanceListAttachmentListOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListAttachmentList) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListAttachmentListOutput) InstanceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListAttachmentList) string { return v.InstanceRegion }).(pulumi.StringOutput)
}

func (o InstancesInstanceListAttachmentListOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListAttachmentList) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListAttachmentListOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceListAttachmentList) string { return v.State }).(pulumi.StringOutput)
}

type InstancesInstanceListAttachmentListArrayOutput struct{ *pulumi.OutputState }

func (InstancesInstanceListAttachmentListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesInstanceListAttachmentList)(nil)).Elem()
}

func (o InstancesInstanceListAttachmentListArrayOutput) ToInstancesInstanceListAttachmentListArrayOutput() InstancesInstanceListAttachmentListArrayOutput {
	return o
}

func (o InstancesInstanceListAttachmentListArrayOutput) ToInstancesInstanceListAttachmentListArrayOutputWithContext(ctx context.Context) InstancesInstanceListAttachmentListArrayOutput {
	return o
}

func (o InstancesInstanceListAttachmentListArrayOutput) Index(i pulumi.IntInput) InstancesInstanceListAttachmentListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesInstanceListAttachmentList {
		return vs[0].([]InstancesInstanceListAttachmentList)[vs[1].(int)]
	}).(InstancesInstanceListAttachmentListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthLimitsLimitInput)(nil)).Elem(), BandwidthLimitsLimitArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthLimitsLimitArrayInput)(nil)).Elem(), BandwidthLimitsLimitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListInput)(nil)).Elem(), InstancesInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListArrayInput)(nil)).Elem(), InstancesInstanceListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListAttachmentListInput)(nil)).Elem(), InstancesInstanceListAttachmentListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListAttachmentListArrayInput)(nil)).Elem(), InstancesInstanceListAttachmentListArray{})
	pulumi.RegisterOutputType(BandwidthLimitsLimitOutput{})
	pulumi.RegisterOutputType(BandwidthLimitsLimitArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListArrayOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListAttachmentListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListAttachmentListArrayOutput{})
}
