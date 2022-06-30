// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GatewayInstances(ctx *pulumi.Context, args *GatewayInstancesArgs, opts ...pulumi.InvokeOption) (*GatewayInstancesResult, error) {
	var rv GatewayInstancesResult
	err := ctx.Invoke("tencentcloud:Dc/gatewayInstances:GatewayInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GatewayInstances.
type GatewayInstancesArgs struct {
	DcgId            *string `pulumi:"dcgId"`
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by GatewayInstances.
type GatewayInstancesResult struct {
	DcgId *string `pulumi:"dcgId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                         `pulumi:"id"`
	InstanceLists    []GatewayInstancesInstanceList `pulumi:"instanceLists"`
	Name             *string                        `pulumi:"name"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

func GatewayInstancesOutput(ctx *pulumi.Context, args GatewayInstancesOutputArgs, opts ...pulumi.InvokeOption) GatewayInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GatewayInstancesResult, error) {
			args := v.(GatewayInstancesArgs)
			r, err := GatewayInstances(ctx, &args, opts...)
			var s GatewayInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GatewayInstancesResultOutput)
}

// A collection of arguments for invoking GatewayInstances.
type GatewayInstancesOutputArgs struct {
	DcgId            pulumi.StringPtrInput `pulumi:"dcgId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GatewayInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayInstancesArgs)(nil)).Elem()
}

// A collection of values returned by GatewayInstances.
type GatewayInstancesResultOutput struct{ *pulumi.OutputState }

func (GatewayInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayInstancesResult)(nil)).Elem()
}

func (o GatewayInstancesResultOutput) ToGatewayInstancesResultOutput() GatewayInstancesResultOutput {
	return o
}

func (o GatewayInstancesResultOutput) ToGatewayInstancesResultOutputWithContext(ctx context.Context) GatewayInstancesResultOutput {
	return o
}

func (o GatewayInstancesResultOutput) DcgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayInstancesResult) *string { return v.DcgId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GatewayInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GatewayInstancesResultOutput) InstanceLists() GatewayInstancesInstanceListArrayOutput {
	return o.ApplyT(func(v GatewayInstancesResult) []GatewayInstancesInstanceList { return v.InstanceLists }).(GatewayInstancesInstanceListArrayOutput)
}

func (o GatewayInstancesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayInstancesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GatewayInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayInstancesResultOutput{})
}