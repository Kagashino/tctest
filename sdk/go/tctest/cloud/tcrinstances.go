// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func TCRInstances(ctx *pulumi.Context, args *TCRInstancesArgs, opts ...pulumi.InvokeOption) (*TCRInstancesResult, error) {
	var rv TCRInstancesResult
	err := ctx.Invoke("tctest:Cloud/tCRInstances:TCRInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking TCRInstances.
type TCRInstancesArgs struct {
	InstanceId       *string `pulumi:"instanceId"`
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by TCRInstances.
type TCRInstancesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                     `pulumi:"id"`
	InstanceId       *string                    `pulumi:"instanceId"`
	InstanceLists    []TCRInstancesInstanceList `pulumi:"instanceLists"`
	Name             *string                    `pulumi:"name"`
	ResultOutputFile *string                    `pulumi:"resultOutputFile"`
}

func TCRInstancesOutput(ctx *pulumi.Context, args TCRInstancesOutputArgs, opts ...pulumi.InvokeOption) TCRInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TCRInstancesResult, error) {
			args := v.(TCRInstancesArgs)
			r, err := TCRInstances(ctx, &args, opts...)
			var s TCRInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TCRInstancesResultOutput)
}

// A collection of arguments for invoking TCRInstances.
type TCRInstancesOutputArgs struct {
	InstanceId       pulumi.StringPtrInput `pulumi:"instanceId"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (TCRInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TCRInstancesArgs)(nil)).Elem()
}

// A collection of values returned by TCRInstances.
type TCRInstancesResultOutput struct{ *pulumi.OutputState }

func (TCRInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TCRInstancesResult)(nil)).Elem()
}

func (o TCRInstancesResultOutput) ToTCRInstancesResultOutput() TCRInstancesResultOutput {
	return o
}

func (o TCRInstancesResultOutput) ToTCRInstancesResultOutputWithContext(ctx context.Context) TCRInstancesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o TCRInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TCRInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o TCRInstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TCRInstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o TCRInstancesResultOutput) InstanceLists() TCRInstancesInstanceListArrayOutput {
	return o.ApplyT(func(v TCRInstancesResult) []TCRInstancesInstanceList { return v.InstanceLists }).(TCRInstancesInstanceListArrayOutput)
}

func (o TCRInstancesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TCRInstancesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TCRInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TCRInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TCRInstancesResultOutput{})
}