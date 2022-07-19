// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstancesInstanceList struct {
	Bandwidth           int      `pulumi:"bandwidth"`
	BgpAsn              int      `pulumi:"bgpAsn"`
	BgpAuthKey          string   `pulumi:"bgpAuthKey"`
	CreateTime          string   `pulumi:"createTime"`
	CustomerAddress     string   `pulumi:"customerAddress"`
	DcId                string   `pulumi:"dcId"`
	DcgId               string   `pulumi:"dcgId"`
	DcxId               string   `pulumi:"dcxId"`
	Name                string   `pulumi:"name"`
	NetworkRegion       string   `pulumi:"networkRegion"`
	NetworkType         string   `pulumi:"networkType"`
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	RouteType           string   `pulumi:"routeType"`
	State               string   `pulumi:"state"`
	TencentAddress      string   `pulumi:"tencentAddress"`
	Vlan                int      `pulumi:"vlan"`
	VpcId               string   `pulumi:"vpcId"`
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
	Bandwidth           pulumi.IntInput         `pulumi:"bandwidth"`
	BgpAsn              pulumi.IntInput         `pulumi:"bgpAsn"`
	BgpAuthKey          pulumi.StringInput      `pulumi:"bgpAuthKey"`
	CreateTime          pulumi.StringInput      `pulumi:"createTime"`
	CustomerAddress     pulumi.StringInput      `pulumi:"customerAddress"`
	DcId                pulumi.StringInput      `pulumi:"dcId"`
	DcgId               pulumi.StringInput      `pulumi:"dcgId"`
	DcxId               pulumi.StringInput      `pulumi:"dcxId"`
	Name                pulumi.StringInput      `pulumi:"name"`
	NetworkRegion       pulumi.StringInput      `pulumi:"networkRegion"`
	NetworkType         pulumi.StringInput      `pulumi:"networkType"`
	RouteFilterPrefixes pulumi.StringArrayInput `pulumi:"routeFilterPrefixes"`
	RouteType           pulumi.StringInput      `pulumi:"routeType"`
	State               pulumi.StringInput      `pulumi:"state"`
	TencentAddress      pulumi.StringInput      `pulumi:"tencentAddress"`
	Vlan                pulumi.IntInput         `pulumi:"vlan"`
	VpcId               pulumi.StringInput      `pulumi:"vpcId"`
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

func (o InstancesInstanceListOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.Bandwidth }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) BgpAsn() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.BgpAsn }).(pulumi.IntOutput)
}

func (o InstancesInstanceListOutput) BgpAuthKey() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.BgpAuthKey }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) CustomerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.CustomerAddress }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) DcId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.DcId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) DcgId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.DcgId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) DcxId() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.DcxId }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.Name }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) NetworkRegion() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.NetworkRegion }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.NetworkType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) RouteFilterPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesInstanceList) []string { return v.RouteFilterPrefixes }).(pulumi.StringArrayOutput)
}

func (o InstancesInstanceListOutput) RouteType() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.RouteType }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.State }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) TencentAddress() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesInstanceList) string { return v.TencentAddress }).(pulumi.StringOutput)
}

func (o InstancesInstanceListOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesInstanceList) int { return v.Vlan }).(pulumi.IntOutput)
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListInput)(nil)).Elem(), InstancesInstanceListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesInstanceListArrayInput)(nil)).Elem(), InstancesInstanceListArray{})
	pulumi.RegisterOutputType(InstancesInstanceListOutput{})
	pulumi.RegisterOutputType(InstancesInstanceListArrayOutput{})
}