// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecretVersionsSecretVersionList struct {
	SecretBinary string `pulumi:"secretBinary"`
	SecretString string `pulumi:"secretString"`
	VersionId    string `pulumi:"versionId"`
}

// SecretVersionsSecretVersionListInput is an input type that accepts SecretVersionsSecretVersionListArgs and SecretVersionsSecretVersionListOutput values.
// You can construct a concrete instance of `SecretVersionsSecretVersionListInput` via:
//
//          SecretVersionsSecretVersionListArgs{...}
type SecretVersionsSecretVersionListInput interface {
	pulumi.Input

	ToSecretVersionsSecretVersionListOutput() SecretVersionsSecretVersionListOutput
	ToSecretVersionsSecretVersionListOutputWithContext(context.Context) SecretVersionsSecretVersionListOutput
}

type SecretVersionsSecretVersionListArgs struct {
	SecretBinary pulumi.StringInput `pulumi:"secretBinary"`
	SecretString pulumi.StringInput `pulumi:"secretString"`
	VersionId    pulumi.StringInput `pulumi:"versionId"`
}

func (SecretVersionsSecretVersionListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretVersionsSecretVersionList)(nil)).Elem()
}

func (i SecretVersionsSecretVersionListArgs) ToSecretVersionsSecretVersionListOutput() SecretVersionsSecretVersionListOutput {
	return i.ToSecretVersionsSecretVersionListOutputWithContext(context.Background())
}

func (i SecretVersionsSecretVersionListArgs) ToSecretVersionsSecretVersionListOutputWithContext(ctx context.Context) SecretVersionsSecretVersionListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionsSecretVersionListOutput)
}

// SecretVersionsSecretVersionListArrayInput is an input type that accepts SecretVersionsSecretVersionListArray and SecretVersionsSecretVersionListArrayOutput values.
// You can construct a concrete instance of `SecretVersionsSecretVersionListArrayInput` via:
//
//          SecretVersionsSecretVersionListArray{ SecretVersionsSecretVersionListArgs{...} }
type SecretVersionsSecretVersionListArrayInput interface {
	pulumi.Input

	ToSecretVersionsSecretVersionListArrayOutput() SecretVersionsSecretVersionListArrayOutput
	ToSecretVersionsSecretVersionListArrayOutputWithContext(context.Context) SecretVersionsSecretVersionListArrayOutput
}

type SecretVersionsSecretVersionListArray []SecretVersionsSecretVersionListInput

func (SecretVersionsSecretVersionListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretVersionsSecretVersionList)(nil)).Elem()
}

func (i SecretVersionsSecretVersionListArray) ToSecretVersionsSecretVersionListArrayOutput() SecretVersionsSecretVersionListArrayOutput {
	return i.ToSecretVersionsSecretVersionListArrayOutputWithContext(context.Background())
}

func (i SecretVersionsSecretVersionListArray) ToSecretVersionsSecretVersionListArrayOutputWithContext(ctx context.Context) SecretVersionsSecretVersionListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionsSecretVersionListArrayOutput)
}

type SecretVersionsSecretVersionListOutput struct{ *pulumi.OutputState }

func (SecretVersionsSecretVersionListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretVersionsSecretVersionList)(nil)).Elem()
}

func (o SecretVersionsSecretVersionListOutput) ToSecretVersionsSecretVersionListOutput() SecretVersionsSecretVersionListOutput {
	return o
}

func (o SecretVersionsSecretVersionListOutput) ToSecretVersionsSecretVersionListOutputWithContext(ctx context.Context) SecretVersionsSecretVersionListOutput {
	return o
}

func (o SecretVersionsSecretVersionListOutput) SecretBinary() pulumi.StringOutput {
	return o.ApplyT(func(v SecretVersionsSecretVersionList) string { return v.SecretBinary }).(pulumi.StringOutput)
}

func (o SecretVersionsSecretVersionListOutput) SecretString() pulumi.StringOutput {
	return o.ApplyT(func(v SecretVersionsSecretVersionList) string { return v.SecretString }).(pulumi.StringOutput)
}

func (o SecretVersionsSecretVersionListOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v SecretVersionsSecretVersionList) string { return v.VersionId }).(pulumi.StringOutput)
}

type SecretVersionsSecretVersionListArrayOutput struct{ *pulumi.OutputState }

func (SecretVersionsSecretVersionListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretVersionsSecretVersionList)(nil)).Elem()
}

func (o SecretVersionsSecretVersionListArrayOutput) ToSecretVersionsSecretVersionListArrayOutput() SecretVersionsSecretVersionListArrayOutput {
	return o
}

func (o SecretVersionsSecretVersionListArrayOutput) ToSecretVersionsSecretVersionListArrayOutputWithContext(ctx context.Context) SecretVersionsSecretVersionListArrayOutput {
	return o
}

func (o SecretVersionsSecretVersionListArrayOutput) Index(i pulumi.IntInput) SecretVersionsSecretVersionListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretVersionsSecretVersionList {
		return vs[0].([]SecretVersionsSecretVersionList)[vs[1].(int)]
	}).(SecretVersionsSecretVersionListOutput)
}

type SecretsSecretList struct {
	CreateTime  int    `pulumi:"createTime"`
	CreateUin   int    `pulumi:"createUin"`
	DeleteTime  int    `pulumi:"deleteTime"`
	Description string `pulumi:"description"`
	KmsKeyId    string `pulumi:"kmsKeyId"`
	SecretName  string `pulumi:"secretName"`
	Status      string `pulumi:"status"`
}

// SecretsSecretListInput is an input type that accepts SecretsSecretListArgs and SecretsSecretListOutput values.
// You can construct a concrete instance of `SecretsSecretListInput` via:
//
//          SecretsSecretListArgs{...}
type SecretsSecretListInput interface {
	pulumi.Input

	ToSecretsSecretListOutput() SecretsSecretListOutput
	ToSecretsSecretListOutputWithContext(context.Context) SecretsSecretListOutput
}

type SecretsSecretListArgs struct {
	CreateTime  pulumi.IntInput    `pulumi:"createTime"`
	CreateUin   pulumi.IntInput    `pulumi:"createUin"`
	DeleteTime  pulumi.IntInput    `pulumi:"deleteTime"`
	Description pulumi.StringInput `pulumi:"description"`
	KmsKeyId    pulumi.StringInput `pulumi:"kmsKeyId"`
	SecretName  pulumi.StringInput `pulumi:"secretName"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (SecretsSecretListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretsSecretList)(nil)).Elem()
}

func (i SecretsSecretListArgs) ToSecretsSecretListOutput() SecretsSecretListOutput {
	return i.ToSecretsSecretListOutputWithContext(context.Background())
}

func (i SecretsSecretListArgs) ToSecretsSecretListOutputWithContext(ctx context.Context) SecretsSecretListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsSecretListOutput)
}

// SecretsSecretListArrayInput is an input type that accepts SecretsSecretListArray and SecretsSecretListArrayOutput values.
// You can construct a concrete instance of `SecretsSecretListArrayInput` via:
//
//          SecretsSecretListArray{ SecretsSecretListArgs{...} }
type SecretsSecretListArrayInput interface {
	pulumi.Input

	ToSecretsSecretListArrayOutput() SecretsSecretListArrayOutput
	ToSecretsSecretListArrayOutputWithContext(context.Context) SecretsSecretListArrayOutput
}

type SecretsSecretListArray []SecretsSecretListInput

func (SecretsSecretListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretsSecretList)(nil)).Elem()
}

func (i SecretsSecretListArray) ToSecretsSecretListArrayOutput() SecretsSecretListArrayOutput {
	return i.ToSecretsSecretListArrayOutputWithContext(context.Background())
}

func (i SecretsSecretListArray) ToSecretsSecretListArrayOutputWithContext(ctx context.Context) SecretsSecretListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsSecretListArrayOutput)
}

type SecretsSecretListOutput struct{ *pulumi.OutputState }

func (SecretsSecretListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretsSecretList)(nil)).Elem()
}

func (o SecretsSecretListOutput) ToSecretsSecretListOutput() SecretsSecretListOutput {
	return o
}

func (o SecretsSecretListOutput) ToSecretsSecretListOutputWithContext(ctx context.Context) SecretsSecretListOutput {
	return o
}

func (o SecretsSecretListOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v SecretsSecretList) int { return v.CreateTime }).(pulumi.IntOutput)
}

func (o SecretsSecretListOutput) CreateUin() pulumi.IntOutput {
	return o.ApplyT(func(v SecretsSecretList) int { return v.CreateUin }).(pulumi.IntOutput)
}

func (o SecretsSecretListOutput) DeleteTime() pulumi.IntOutput {
	return o.ApplyT(func(v SecretsSecretList) int { return v.DeleteTime }).(pulumi.IntOutput)
}

func (o SecretsSecretListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v SecretsSecretList) string { return v.Description }).(pulumi.StringOutput)
}

func (o SecretsSecretListOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v SecretsSecretList) string { return v.KmsKeyId }).(pulumi.StringOutput)
}

func (o SecretsSecretListOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v SecretsSecretList) string { return v.SecretName }).(pulumi.StringOutput)
}

func (o SecretsSecretListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SecretsSecretList) string { return v.Status }).(pulumi.StringOutput)
}

type SecretsSecretListArrayOutput struct{ *pulumi.OutputState }

func (SecretsSecretListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretsSecretList)(nil)).Elem()
}

func (o SecretsSecretListArrayOutput) ToSecretsSecretListArrayOutput() SecretsSecretListArrayOutput {
	return o
}

func (o SecretsSecretListArrayOutput) ToSecretsSecretListArrayOutputWithContext(ctx context.Context) SecretsSecretListArrayOutput {
	return o
}

func (o SecretsSecretListArrayOutput) Index(i pulumi.IntInput) SecretsSecretListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretsSecretList {
		return vs[0].([]SecretsSecretList)[vs[1].(int)]
	}).(SecretsSecretListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionsSecretVersionListInput)(nil)).Elem(), SecretVersionsSecretVersionListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionsSecretVersionListArrayInput)(nil)).Elem(), SecretVersionsSecretVersionListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsSecretListInput)(nil)).Elem(), SecretsSecretListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsSecretListArrayInput)(nil)).Elem(), SecretsSecretListArray{})
	pulumi.RegisterOutputType(SecretVersionsSecretVersionListOutput{})
	pulumi.RegisterOutputType(SecretVersionsSecretVersionListArrayOutput{})
	pulumi.RegisterOutputType(SecretsSecretListOutput{})
	pulumi.RegisterOutputType(SecretsSecretListArrayOutput{})
}