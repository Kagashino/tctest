// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func PolicyConditions(ctx *pulumi.Context, args *PolicyConditionsArgs, opts ...pulumi.InvokeOption) (*PolicyConditionsResult, error) {
	var rv PolicyConditionsResult
	err := ctx.Invoke("tencentcloud:Monitor/policyConditions:PolicyConditions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking PolicyConditions.
type PolicyConditionsArgs struct {
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by PolicyConditions.
type PolicyConditionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                 `pulumi:"id"`
	Lists            []PolicyConditionsList `pulumi:"lists"`
	Name             *string                `pulumi:"name"`
	ResultOutputFile *string                `pulumi:"resultOutputFile"`
}

func PolicyConditionsOutput(ctx *pulumi.Context, args PolicyConditionsOutputArgs, opts ...pulumi.InvokeOption) PolicyConditionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PolicyConditionsResult, error) {
			args := v.(PolicyConditionsArgs)
			r, err := PolicyConditions(ctx, &args, opts...)
			var s PolicyConditionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PolicyConditionsResultOutput)
}

// A collection of arguments for invoking PolicyConditions.
type PolicyConditionsOutputArgs struct {
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (PolicyConditionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyConditionsArgs)(nil)).Elem()
}

// A collection of values returned by PolicyConditions.
type PolicyConditionsResultOutput struct{ *pulumi.OutputState }

func (PolicyConditionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyConditionsResult)(nil)).Elem()
}

func (o PolicyConditionsResultOutput) ToPolicyConditionsResultOutput() PolicyConditionsResultOutput {
	return o
}

func (o PolicyConditionsResultOutput) ToPolicyConditionsResultOutputWithContext(ctx context.Context) PolicyConditionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o PolicyConditionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyConditionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o PolicyConditionsResultOutput) Lists() PolicyConditionsListArrayOutput {
	return o.ApplyT(func(v PolicyConditionsResult) []PolicyConditionsList { return v.Lists }).(PolicyConditionsListArrayOutput)
}

func (o PolicyConditionsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyConditionsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PolicyConditionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyConditionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyConditionsResultOutput{})
}