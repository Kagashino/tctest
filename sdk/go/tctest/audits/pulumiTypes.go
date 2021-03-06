// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package audits

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstancesAuditList struct {
	AuditSwitch   bool   `pulumi:"auditSwitch"`
	CosBucket     string `pulumi:"cosBucket"`
	Id            string `pulumi:"id"`
	LogFilePrefix string `pulumi:"logFilePrefix"`
	Name          string `pulumi:"name"`
}

// InstancesAuditListInput is an input type that accepts InstancesAuditListArgs and InstancesAuditListOutput values.
// You can construct a concrete instance of `InstancesAuditListInput` via:
//
//          InstancesAuditListArgs{...}
type InstancesAuditListInput interface {
	pulumi.Input

	ToInstancesAuditListOutput() InstancesAuditListOutput
	ToInstancesAuditListOutputWithContext(context.Context) InstancesAuditListOutput
}

type InstancesAuditListArgs struct {
	AuditSwitch   pulumi.BoolInput   `pulumi:"auditSwitch"`
	CosBucket     pulumi.StringInput `pulumi:"cosBucket"`
	Id            pulumi.StringInput `pulumi:"id"`
	LogFilePrefix pulumi.StringInput `pulumi:"logFilePrefix"`
	Name          pulumi.StringInput `pulumi:"name"`
}

func (InstancesAuditListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesAuditList)(nil)).Elem()
}

func (i InstancesAuditListArgs) ToInstancesAuditListOutput() InstancesAuditListOutput {
	return i.ToInstancesAuditListOutputWithContext(context.Background())
}

func (i InstancesAuditListArgs) ToInstancesAuditListOutputWithContext(ctx context.Context) InstancesAuditListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesAuditListOutput)
}

// InstancesAuditListArrayInput is an input type that accepts InstancesAuditListArray and InstancesAuditListArrayOutput values.
// You can construct a concrete instance of `InstancesAuditListArrayInput` via:
//
//          InstancesAuditListArray{ InstancesAuditListArgs{...} }
type InstancesAuditListArrayInput interface {
	pulumi.Input

	ToInstancesAuditListArrayOutput() InstancesAuditListArrayOutput
	ToInstancesAuditListArrayOutputWithContext(context.Context) InstancesAuditListArrayOutput
}

type InstancesAuditListArray []InstancesAuditListInput

func (InstancesAuditListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesAuditList)(nil)).Elem()
}

func (i InstancesAuditListArray) ToInstancesAuditListArrayOutput() InstancesAuditListArrayOutput {
	return i.ToInstancesAuditListArrayOutputWithContext(context.Background())
}

func (i InstancesAuditListArray) ToInstancesAuditListArrayOutputWithContext(ctx context.Context) InstancesAuditListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesAuditListArrayOutput)
}

type InstancesAuditListOutput struct{ *pulumi.OutputState }

func (InstancesAuditListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesAuditList)(nil)).Elem()
}

func (o InstancesAuditListOutput) ToInstancesAuditListOutput() InstancesAuditListOutput {
	return o
}

func (o InstancesAuditListOutput) ToInstancesAuditListOutputWithContext(ctx context.Context) InstancesAuditListOutput {
	return o
}

func (o InstancesAuditListOutput) AuditSwitch() pulumi.BoolOutput {
	return o.ApplyT(func(v InstancesAuditList) bool { return v.AuditSwitch }).(pulumi.BoolOutput)
}

func (o InstancesAuditListOutput) CosBucket() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesAuditList) string { return v.CosBucket }).(pulumi.StringOutput)
}

func (o InstancesAuditListOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesAuditList) string { return v.Id }).(pulumi.StringOutput)
}

func (o InstancesAuditListOutput) LogFilePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesAuditList) string { return v.LogFilePrefix }).(pulumi.StringOutput)
}

func (o InstancesAuditListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesAuditList) string { return v.Name }).(pulumi.StringOutput)
}

type InstancesAuditListArrayOutput struct{ *pulumi.OutputState }

func (InstancesAuditListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstancesAuditList)(nil)).Elem()
}

func (o InstancesAuditListArrayOutput) ToInstancesAuditListArrayOutput() InstancesAuditListArrayOutput {
	return o
}

func (o InstancesAuditListArrayOutput) ToInstancesAuditListArrayOutputWithContext(ctx context.Context) InstancesAuditListArrayOutput {
	return o
}

func (o InstancesAuditListArrayOutput) Index(i pulumi.IntInput) InstancesAuditListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstancesAuditList {
		return vs[0].([]InstancesAuditList)[vs[1].(int)]
	}).(InstancesAuditListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesAuditListInput)(nil)).Elem(), InstancesAuditListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesAuditListArrayInput)(nil)).Elem(), InstancesAuditListArray{})
	pulumi.RegisterOutputType(InstancesAuditListOutput{})
	pulumi.RegisterOutputType(InstancesAuditListArrayOutput{})
}
