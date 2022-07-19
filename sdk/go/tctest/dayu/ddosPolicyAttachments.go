// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func DdosPolicyAttachments(ctx *pulumi.Context, args *DdosPolicyAttachmentsArgs, opts ...pulumi.InvokeOption) (*DdosPolicyAttachmentsResult, error) {
	var rv DdosPolicyAttachmentsResult
	err := ctx.Invoke("tctest:Dayu/ddosPolicyAttachments:DdosPolicyAttachments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking DdosPolicyAttachments.
type DdosPolicyAttachmentsArgs struct {
	PolicyId         *string `pulumi:"policyId"`
	ResourceId       *string `pulumi:"resourceId"`
	ResourceType     string  `pulumi:"resourceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by DdosPolicyAttachments.
type DdosPolicyAttachmentsResult struct {
	DayuDdosPolicyAttachmentLists []DdosPolicyAttachmentsDayuDdosPolicyAttachmentList `pulumi:"dayuDdosPolicyAttachmentLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	PolicyId         *string `pulumi:"policyId"`
	ResourceId       *string `pulumi:"resourceId"`
	ResourceType     string  `pulumi:"resourceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func DdosPolicyAttachmentsOutput(ctx *pulumi.Context, args DdosPolicyAttachmentsOutputArgs, opts ...pulumi.InvokeOption) DdosPolicyAttachmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DdosPolicyAttachmentsResult, error) {
			args := v.(DdosPolicyAttachmentsArgs)
			r, err := DdosPolicyAttachments(ctx, &args, opts...)
			var s DdosPolicyAttachmentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DdosPolicyAttachmentsResultOutput)
}

// A collection of arguments for invoking DdosPolicyAttachments.
type DdosPolicyAttachmentsOutputArgs struct {
	PolicyId         pulumi.StringPtrInput `pulumi:"policyId"`
	ResourceId       pulumi.StringPtrInput `pulumi:"resourceId"`
	ResourceType     pulumi.StringInput    `pulumi:"resourceType"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (DdosPolicyAttachmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosPolicyAttachmentsArgs)(nil)).Elem()
}

// A collection of values returned by DdosPolicyAttachments.
type DdosPolicyAttachmentsResultOutput struct{ *pulumi.OutputState }

func (DdosPolicyAttachmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DdosPolicyAttachmentsResult)(nil)).Elem()
}

func (o DdosPolicyAttachmentsResultOutput) ToDdosPolicyAttachmentsResultOutput() DdosPolicyAttachmentsResultOutput {
	return o
}

func (o DdosPolicyAttachmentsResultOutput) ToDdosPolicyAttachmentsResultOutputWithContext(ctx context.Context) DdosPolicyAttachmentsResultOutput {
	return o
}

func (o DdosPolicyAttachmentsResultOutput) DayuDdosPolicyAttachmentLists() DdosPolicyAttachmentsDayuDdosPolicyAttachmentListArrayOutput {
	return o.ApplyT(func(v DdosPolicyAttachmentsResult) []DdosPolicyAttachmentsDayuDdosPolicyAttachmentList {
		return v.DayuDdosPolicyAttachmentLists
	}).(DdosPolicyAttachmentsDayuDdosPolicyAttachmentListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o DdosPolicyAttachmentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DdosPolicyAttachmentsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o DdosPolicyAttachmentsResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DdosPolicyAttachmentsResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o DdosPolicyAttachmentsResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DdosPolicyAttachmentsResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o DdosPolicyAttachmentsResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v DdosPolicyAttachmentsResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o DdosPolicyAttachmentsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DdosPolicyAttachmentsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DdosPolicyAttachmentsResultOutput{})
}