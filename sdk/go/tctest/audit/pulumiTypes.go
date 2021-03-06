// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package audit

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CosRegionsAuditCosRegionList struct {
	CosRegion     string `pulumi:"cosRegion"`
	CosRegionName string `pulumi:"cosRegionName"`
}

// CosRegionsAuditCosRegionListInput is an input type that accepts CosRegionsAuditCosRegionListArgs and CosRegionsAuditCosRegionListOutput values.
// You can construct a concrete instance of `CosRegionsAuditCosRegionListInput` via:
//
//          CosRegionsAuditCosRegionListArgs{...}
type CosRegionsAuditCosRegionListInput interface {
	pulumi.Input

	ToCosRegionsAuditCosRegionListOutput() CosRegionsAuditCosRegionListOutput
	ToCosRegionsAuditCosRegionListOutputWithContext(context.Context) CosRegionsAuditCosRegionListOutput
}

type CosRegionsAuditCosRegionListArgs struct {
	CosRegion     pulumi.StringInput `pulumi:"cosRegion"`
	CosRegionName pulumi.StringInput `pulumi:"cosRegionName"`
}

func (CosRegionsAuditCosRegionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CosRegionsAuditCosRegionList)(nil)).Elem()
}

func (i CosRegionsAuditCosRegionListArgs) ToCosRegionsAuditCosRegionListOutput() CosRegionsAuditCosRegionListOutput {
	return i.ToCosRegionsAuditCosRegionListOutputWithContext(context.Background())
}

func (i CosRegionsAuditCosRegionListArgs) ToCosRegionsAuditCosRegionListOutputWithContext(ctx context.Context) CosRegionsAuditCosRegionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosRegionsAuditCosRegionListOutput)
}

// CosRegionsAuditCosRegionListArrayInput is an input type that accepts CosRegionsAuditCosRegionListArray and CosRegionsAuditCosRegionListArrayOutput values.
// You can construct a concrete instance of `CosRegionsAuditCosRegionListArrayInput` via:
//
//          CosRegionsAuditCosRegionListArray{ CosRegionsAuditCosRegionListArgs{...} }
type CosRegionsAuditCosRegionListArrayInput interface {
	pulumi.Input

	ToCosRegionsAuditCosRegionListArrayOutput() CosRegionsAuditCosRegionListArrayOutput
	ToCosRegionsAuditCosRegionListArrayOutputWithContext(context.Context) CosRegionsAuditCosRegionListArrayOutput
}

type CosRegionsAuditCosRegionListArray []CosRegionsAuditCosRegionListInput

func (CosRegionsAuditCosRegionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CosRegionsAuditCosRegionList)(nil)).Elem()
}

func (i CosRegionsAuditCosRegionListArray) ToCosRegionsAuditCosRegionListArrayOutput() CosRegionsAuditCosRegionListArrayOutput {
	return i.ToCosRegionsAuditCosRegionListArrayOutputWithContext(context.Background())
}

func (i CosRegionsAuditCosRegionListArray) ToCosRegionsAuditCosRegionListArrayOutputWithContext(ctx context.Context) CosRegionsAuditCosRegionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosRegionsAuditCosRegionListArrayOutput)
}

type CosRegionsAuditCosRegionListOutput struct{ *pulumi.OutputState }

func (CosRegionsAuditCosRegionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CosRegionsAuditCosRegionList)(nil)).Elem()
}

func (o CosRegionsAuditCosRegionListOutput) ToCosRegionsAuditCosRegionListOutput() CosRegionsAuditCosRegionListOutput {
	return o
}

func (o CosRegionsAuditCosRegionListOutput) ToCosRegionsAuditCosRegionListOutputWithContext(ctx context.Context) CosRegionsAuditCosRegionListOutput {
	return o
}

func (o CosRegionsAuditCosRegionListOutput) CosRegion() pulumi.StringOutput {
	return o.ApplyT(func(v CosRegionsAuditCosRegionList) string { return v.CosRegion }).(pulumi.StringOutput)
}

func (o CosRegionsAuditCosRegionListOutput) CosRegionName() pulumi.StringOutput {
	return o.ApplyT(func(v CosRegionsAuditCosRegionList) string { return v.CosRegionName }).(pulumi.StringOutput)
}

type CosRegionsAuditCosRegionListArrayOutput struct{ *pulumi.OutputState }

func (CosRegionsAuditCosRegionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CosRegionsAuditCosRegionList)(nil)).Elem()
}

func (o CosRegionsAuditCosRegionListArrayOutput) ToCosRegionsAuditCosRegionListArrayOutput() CosRegionsAuditCosRegionListArrayOutput {
	return o
}

func (o CosRegionsAuditCosRegionListArrayOutput) ToCosRegionsAuditCosRegionListArrayOutputWithContext(ctx context.Context) CosRegionsAuditCosRegionListArrayOutput {
	return o
}

func (o CosRegionsAuditCosRegionListArrayOutput) Index(i pulumi.IntInput) CosRegionsAuditCosRegionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CosRegionsAuditCosRegionList {
		return vs[0].([]CosRegionsAuditCosRegionList)[vs[1].(int)]
	}).(CosRegionsAuditCosRegionListOutput)
}

type KeyAliasAuditKeyAliasList struct {
	KeyAlias string `pulumi:"keyAlias"`
	KeyId    string `pulumi:"keyId"`
}

// KeyAliasAuditKeyAliasListInput is an input type that accepts KeyAliasAuditKeyAliasListArgs and KeyAliasAuditKeyAliasListOutput values.
// You can construct a concrete instance of `KeyAliasAuditKeyAliasListInput` via:
//
//          KeyAliasAuditKeyAliasListArgs{...}
type KeyAliasAuditKeyAliasListInput interface {
	pulumi.Input

	ToKeyAliasAuditKeyAliasListOutput() KeyAliasAuditKeyAliasListOutput
	ToKeyAliasAuditKeyAliasListOutputWithContext(context.Context) KeyAliasAuditKeyAliasListOutput
}

type KeyAliasAuditKeyAliasListArgs struct {
	KeyAlias pulumi.StringInput `pulumi:"keyAlias"`
	KeyId    pulumi.StringInput `pulumi:"keyId"`
}

func (KeyAliasAuditKeyAliasListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (i KeyAliasAuditKeyAliasListArgs) ToKeyAliasAuditKeyAliasListOutput() KeyAliasAuditKeyAliasListOutput {
	return i.ToKeyAliasAuditKeyAliasListOutputWithContext(context.Background())
}

func (i KeyAliasAuditKeyAliasListArgs) ToKeyAliasAuditKeyAliasListOutputWithContext(ctx context.Context) KeyAliasAuditKeyAliasListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAliasAuditKeyAliasListOutput)
}

// KeyAliasAuditKeyAliasListArrayInput is an input type that accepts KeyAliasAuditKeyAliasListArray and KeyAliasAuditKeyAliasListArrayOutput values.
// You can construct a concrete instance of `KeyAliasAuditKeyAliasListArrayInput` via:
//
//          KeyAliasAuditKeyAliasListArray{ KeyAliasAuditKeyAliasListArgs{...} }
type KeyAliasAuditKeyAliasListArrayInput interface {
	pulumi.Input

	ToKeyAliasAuditKeyAliasListArrayOutput() KeyAliasAuditKeyAliasListArrayOutput
	ToKeyAliasAuditKeyAliasListArrayOutputWithContext(context.Context) KeyAliasAuditKeyAliasListArrayOutput
}

type KeyAliasAuditKeyAliasListArray []KeyAliasAuditKeyAliasListInput

func (KeyAliasAuditKeyAliasListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (i KeyAliasAuditKeyAliasListArray) ToKeyAliasAuditKeyAliasListArrayOutput() KeyAliasAuditKeyAliasListArrayOutput {
	return i.ToKeyAliasAuditKeyAliasListArrayOutputWithContext(context.Background())
}

func (i KeyAliasAuditKeyAliasListArray) ToKeyAliasAuditKeyAliasListArrayOutputWithContext(ctx context.Context) KeyAliasAuditKeyAliasListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyAliasAuditKeyAliasListArrayOutput)
}

type KeyAliasAuditKeyAliasListOutput struct{ *pulumi.OutputState }

func (KeyAliasAuditKeyAliasListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (o KeyAliasAuditKeyAliasListOutput) ToKeyAliasAuditKeyAliasListOutput() KeyAliasAuditKeyAliasListOutput {
	return o
}

func (o KeyAliasAuditKeyAliasListOutput) ToKeyAliasAuditKeyAliasListOutputWithContext(ctx context.Context) KeyAliasAuditKeyAliasListOutput {
	return o
}

func (o KeyAliasAuditKeyAliasListOutput) KeyAlias() pulumi.StringOutput {
	return o.ApplyT(func(v KeyAliasAuditKeyAliasList) string { return v.KeyAlias }).(pulumi.StringOutput)
}

func (o KeyAliasAuditKeyAliasListOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v KeyAliasAuditKeyAliasList) string { return v.KeyId }).(pulumi.StringOutput)
}

type KeyAliasAuditKeyAliasListArrayOutput struct{ *pulumi.OutputState }

func (KeyAliasAuditKeyAliasListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyAliasAuditKeyAliasList)(nil)).Elem()
}

func (o KeyAliasAuditKeyAliasListArrayOutput) ToKeyAliasAuditKeyAliasListArrayOutput() KeyAliasAuditKeyAliasListArrayOutput {
	return o
}

func (o KeyAliasAuditKeyAliasListArrayOutput) ToKeyAliasAuditKeyAliasListArrayOutputWithContext(ctx context.Context) KeyAliasAuditKeyAliasListArrayOutput {
	return o
}

func (o KeyAliasAuditKeyAliasListArrayOutput) Index(i pulumi.IntInput) KeyAliasAuditKeyAliasListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyAliasAuditKeyAliasList {
		return vs[0].([]KeyAliasAuditKeyAliasList)[vs[1].(int)]
	}).(KeyAliasAuditKeyAliasListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CosRegionsAuditCosRegionListInput)(nil)).Elem(), CosRegionsAuditCosRegionListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CosRegionsAuditCosRegionListArrayInput)(nil)).Elem(), CosRegionsAuditCosRegionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAliasAuditKeyAliasListInput)(nil)).Elem(), KeyAliasAuditKeyAliasListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyAliasAuditKeyAliasListArrayInput)(nil)).Elem(), KeyAliasAuditKeyAliasListArray{})
	pulumi.RegisterOutputType(CosRegionsAuditCosRegionListOutput{})
	pulumi.RegisterOutputType(CosRegionsAuditCosRegionListArrayOutput{})
	pulumi.RegisterOutputType(KeyAliasAuditKeyAliasListOutput{})
	pulumi.RegisterOutputType(KeyAliasAuditKeyAliasListArrayOutput{})
}
