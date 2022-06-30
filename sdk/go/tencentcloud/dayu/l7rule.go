// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type L7Rule struct {
	pulumi.CustomResourceState

	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means
	// the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is
	// health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add
	// the corresponding values.
	HealthCheckCode pulumi.IntOutput `pulumi:"healthCheckCode"`
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3
	// consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum pulumi.IntOutput `pulumi:"healthCheckHealthNum"`
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval pulumi.IntOutput `pulumi:"healthCheckInterval"`
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod pulumi.StringOutput `pulumi:"healthCheckMethod"`
	// Path of health check. The default is `/`.
	HealthCheckPath pulumi.StringOutput `pulumi:"healthCheckPath"`
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch pulumi.BoolOutput `pulumi:"healthCheckSwitch"`
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times,
	// indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum pulumi.IntOutput `pulumi:"healthCheckUnhealthNum"`
	// Name of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// ID of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// ID of the layer 7 rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to
	// 16.
	SourceLists pulumi.StringArrayOutput `pulumi:"sourceLists"`
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType pulumi.IntOutput `pulumi:"sourceType"`
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SslId pulumi.StringPtrOutput `pulumi:"sslId"`
	// Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete
	// failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
	Status pulumi.IntOutput `pulumi:"status"`
	// Indicate the rule will take effect or not.
	Switch pulumi.BoolOutput `pulumi:"switch"`
}

// NewL7Rule registers a new resource with the given unique name, arguments, and options.
func NewL7Rule(ctx *pulumi.Context,
	name string, args *L7RuleArgs, opts ...pulumi.ResourceOption) (*L7Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.SourceLists == nil {
		return nil, errors.New("invalid value for required argument 'SourceLists'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	if args.Switch == nil {
		return nil, errors.New("invalid value for required argument 'Switch'")
	}
	var resource L7Rule
	err := ctx.RegisterResource("tencentcloud:Dayu/l7Rule:L7Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL7Rule gets an existing L7Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL7Rule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L7RuleState, opts ...pulumi.ResourceOption) (*L7Rule, error) {
	var resource L7Rule
	err := ctx.ReadResource("tencentcloud:Dayu/l7Rule:L7Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L7Rule resources.
type l7ruleState struct {
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain *string `pulumi:"domain"`
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means
	// the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is
	// health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add
	// the corresponding values.
	HealthCheckCode *int `pulumi:"healthCheckCode"`
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3
	// consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum *int `pulumi:"healthCheckHealthNum"`
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod *string `pulumi:"healthCheckMethod"`
	// Path of health check. The default is `/`.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch *bool `pulumi:"healthCheckSwitch"`
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times,
	// indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum *int `pulumi:"healthCheckUnhealthNum"`
	// Name of the rule.
	Name *string `pulumi:"name"`
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol *string `pulumi:"protocol"`
	// ID of the resource that the layer 7 rule works for.
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType *string `pulumi:"resourceType"`
	// ID of the layer 7 rule.
	RuleId *string `pulumi:"ruleId"`
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to
	// 16.
	SourceLists []string `pulumi:"sourceLists"`
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType *int `pulumi:"sourceType"`
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SslId *string `pulumi:"sslId"`
	// Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete
	// failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
	Status *int `pulumi:"status"`
	// Indicate the rule will take effect or not.
	Switch *bool `pulumi:"switch"`
}

type L7RuleState struct {
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain pulumi.StringPtrInput
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means
	// the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is
	// health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add
	// the corresponding values.
	HealthCheckCode pulumi.IntPtrInput
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3
	// consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum pulumi.IntPtrInput
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval pulumi.IntPtrInput
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod pulumi.StringPtrInput
	// Path of health check. The default is `/`.
	HealthCheckPath pulumi.StringPtrInput
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch pulumi.BoolPtrInput
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times,
	// indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum pulumi.IntPtrInput
	// Name of the rule.
	Name pulumi.StringPtrInput
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol pulumi.StringPtrInput
	// ID of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringPtrInput
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringPtrInput
	// ID of the layer 7 rule.
	RuleId pulumi.StringPtrInput
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to
	// 16.
	SourceLists pulumi.StringArrayInput
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType pulumi.IntPtrInput
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SslId pulumi.StringPtrInput
	// Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete
	// failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
	Status pulumi.IntPtrInput
	// Indicate the rule will take effect or not.
	Switch pulumi.BoolPtrInput
}

func (L7RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleState)(nil)).Elem()
}

type l7ruleArgs struct {
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain string `pulumi:"domain"`
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means
	// the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is
	// health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add
	// the corresponding values.
	HealthCheckCode *int `pulumi:"healthCheckCode"`
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3
	// consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum *int `pulumi:"healthCheckHealthNum"`
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod *string `pulumi:"healthCheckMethod"`
	// Path of health check. The default is `/`.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch *bool `pulumi:"healthCheckSwitch"`
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times,
	// indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum *int `pulumi:"healthCheckUnhealthNum"`
	// Name of the rule.
	Name *string `pulumi:"name"`
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol string `pulumi:"protocol"`
	// ID of the resource that the layer 7 rule works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType string `pulumi:"resourceType"`
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to
	// 16.
	SourceLists []string `pulumi:"sourceLists"`
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType int `pulumi:"sourceType"`
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SslId *string `pulumi:"sslId"`
	// Indicate the rule will take effect or not.
	Switch bool `pulumi:"switch"`
}

// The set of arguments for constructing a L7Rule resource.
type L7RuleArgs struct {
	// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
	Domain pulumi.StringInput
	// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means
	// the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is
	// health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add
	// the corresponding values.
	HealthCheckCode pulumi.IntPtrInput
	// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3
	// consecutive times, indicates that the forwarding is normal. The value range is [2-10].
	HealthCheckHealthNum pulumi.IntPtrInput
	// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
	HealthCheckInterval pulumi.IntPtrInput
	// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
	HealthCheckMethod pulumi.StringPtrInput
	// Path of health check. The default is `/`.
	HealthCheckPath pulumi.StringPtrInput
	// Indicates whether health check is enabled. The default is `false`.
	HealthCheckSwitch pulumi.BoolPtrInput
	// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times,
	// indicates that the forwarding is abnormal. The value range is [2-10].
	HealthCheckUnhealthNum pulumi.IntPtrInput
	// Name of the rule.
	Name pulumi.StringPtrInput
	// Protocol of the rule. Valid values: `http`, `https`.
	Protocol pulumi.StringInput
	// ID of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringInput
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringInput
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to
	// 16.
	SourceLists pulumi.StringArrayInput
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType pulumi.IntInput
	// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
	SslId pulumi.StringPtrInput
	// Indicate the rule will take effect or not.
	Switch pulumi.BoolInput
}

func (L7RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleArgs)(nil)).Elem()
}

type L7RuleInput interface {
	pulumi.Input

	ToL7RuleOutput() L7RuleOutput
	ToL7RuleOutputWithContext(ctx context.Context) L7RuleOutput
}

func (*L7Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**L7Rule)(nil)).Elem()
}

func (i *L7Rule) ToL7RuleOutput() L7RuleOutput {
	return i.ToL7RuleOutputWithContext(context.Background())
}

func (i *L7Rule) ToL7RuleOutputWithContext(ctx context.Context) L7RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleOutput)
}

// L7RuleArrayInput is an input type that accepts L7RuleArray and L7RuleArrayOutput values.
// You can construct a concrete instance of `L7RuleArrayInput` via:
//
//          L7RuleArray{ L7RuleArgs{...} }
type L7RuleArrayInput interface {
	pulumi.Input

	ToL7RuleArrayOutput() L7RuleArrayOutput
	ToL7RuleArrayOutputWithContext(context.Context) L7RuleArrayOutput
}

type L7RuleArray []L7RuleInput

func (L7RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7Rule)(nil)).Elem()
}

func (i L7RuleArray) ToL7RuleArrayOutput() L7RuleArrayOutput {
	return i.ToL7RuleArrayOutputWithContext(context.Background())
}

func (i L7RuleArray) ToL7RuleArrayOutputWithContext(ctx context.Context) L7RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleArrayOutput)
}

// L7RuleMapInput is an input type that accepts L7RuleMap and L7RuleMapOutput values.
// You can construct a concrete instance of `L7RuleMapInput` via:
//
//          L7RuleMap{ "key": L7RuleArgs{...} }
type L7RuleMapInput interface {
	pulumi.Input

	ToL7RuleMapOutput() L7RuleMapOutput
	ToL7RuleMapOutputWithContext(context.Context) L7RuleMapOutput
}

type L7RuleMap map[string]L7RuleInput

func (L7RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7Rule)(nil)).Elem()
}

func (i L7RuleMap) ToL7RuleMapOutput() L7RuleMapOutput {
	return i.ToL7RuleMapOutputWithContext(context.Background())
}

func (i L7RuleMap) ToL7RuleMapOutputWithContext(ctx context.Context) L7RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleMapOutput)
}

type L7RuleOutput struct{ *pulumi.OutputState }

func (L7RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L7Rule)(nil)).Elem()
}

func (o L7RuleOutput) ToL7RuleOutput() L7RuleOutput {
	return o
}

func (o L7RuleOutput) ToL7RuleOutputWithContext(ctx context.Context) L7RuleOutput {
	return o
}

// Domain that the layer 7 rule works for. Valid string length ranges from 0 to 80.
func (o L7RuleOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// HTTP Status Code. The default is `26`. Valid value ranges: [1~31]. `1` means the return value '1xx' is health. `2` means
// the return value '2xx' is health. `4` means the return value '3xx' is health. `8` means the return value '4xx' is
// health. `16` means the return value '5xx' is health. If you want multiple return codes to indicate health, need to add
// the corresponding values.
func (o L7RuleOutput) HealthCheckCode() pulumi.IntOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.IntOutput { return v.HealthCheckCode }).(pulumi.IntOutput)
}

// Health threshold of health check, and the default is `3`. If a success result is returned for the health check 3
// consecutive times, indicates that the forwarding is normal. The value range is [2-10].
func (o L7RuleOutput) HealthCheckHealthNum() pulumi.IntOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.IntOutput { return v.HealthCheckHealthNum }).(pulumi.IntOutput)
}

// Interval time of health check. Valid value ranges: [10~60]sec. The default is 15 sec.
func (o L7RuleOutput) HealthCheckInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.IntOutput { return v.HealthCheckInterval }).(pulumi.IntOutput)
}

// Methods of health check. The default is 'HEAD', the available value are 'HEAD' and 'GET'.
func (o L7RuleOutput) HealthCheckMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.HealthCheckMethod }).(pulumi.StringOutput)
}

// Path of health check. The default is `/`.
func (o L7RuleOutput) HealthCheckPath() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.HealthCheckPath }).(pulumi.StringOutput)
}

// Indicates whether health check is enabled. The default is `false`.
func (o L7RuleOutput) HealthCheckSwitch() pulumi.BoolOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.BoolOutput { return v.HealthCheckSwitch }).(pulumi.BoolOutput)
}

// Unhealthy threshold of health check, and the default is `3`. If the unhealthy result is returned 3 consecutive times,
// indicates that the forwarding is abnormal. The value range is [2-10].
func (o L7RuleOutput) HealthCheckUnhealthNum() pulumi.IntOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.IntOutput { return v.HealthCheckUnhealthNum }).(pulumi.IntOutput)
}

// Name of the rule.
func (o L7RuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Protocol of the rule. Valid values: `http`, `https`.
func (o L7RuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// ID of the resource that the layer 7 rule works for.
func (o L7RuleOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
func (o L7RuleOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// ID of the layer 7 rule.
func (o L7RuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to
// 16.
func (o L7RuleOutput) SourceLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringArrayOutput { return v.SourceLists }).(pulumi.StringArrayOutput)
}

// Source type, `1` for source of host, `2` for source of IP.
func (o L7RuleOutput) SourceType() pulumi.IntOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.IntOutput { return v.SourceType }).(pulumi.IntOutput)
}

// SSL ID, when the `protocol` is `https`, the field should be set with valid SSL id.
func (o L7RuleOutput) SslId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.StringPtrOutput { return v.SslId }).(pulumi.StringPtrOutput)
}

// Status of the rule. `0` for create/modify success, `2` for create/modify fail, `3` for delete success, `5` for delete
// failed, `6` for waiting to be created/modified, `7` for waiting to be deleted and 8 for waiting to get SSL ID.
func (o L7RuleOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Indicate the rule will take effect or not.
func (o L7RuleOutput) Switch() pulumi.BoolOutput {
	return o.ApplyT(func(v *L7Rule) pulumi.BoolOutput { return v.Switch }).(pulumi.BoolOutput)
}

type L7RuleArrayOutput struct{ *pulumi.OutputState }

func (L7RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7Rule)(nil)).Elem()
}

func (o L7RuleArrayOutput) ToL7RuleArrayOutput() L7RuleArrayOutput {
	return o
}

func (o L7RuleArrayOutput) ToL7RuleArrayOutputWithContext(ctx context.Context) L7RuleArrayOutput {
	return o
}

func (o L7RuleArrayOutput) Index(i pulumi.IntInput) L7RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *L7Rule {
		return vs[0].([]*L7Rule)[vs[1].(int)]
	}).(L7RuleOutput)
}

type L7RuleMapOutput struct{ *pulumi.OutputState }

func (L7RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7Rule)(nil)).Elem()
}

func (o L7RuleMapOutput) ToL7RuleMapOutput() L7RuleMapOutput {
	return o
}

func (o L7RuleMapOutput) ToL7RuleMapOutputWithContext(ctx context.Context) L7RuleMapOutput {
	return o
}

func (o L7RuleMapOutput) MapIndex(k pulumi.StringInput) L7RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *L7Rule {
		return vs[0].(map[string]*L7Rule)[vs[1].(string)]
	}).(L7RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*L7RuleInput)(nil)).Elem(), &L7Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7RuleArrayInput)(nil)).Elem(), L7RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7RuleMapInput)(nil)).Elem(), L7RuleMap{})
	pulumi.RegisterOutputType(L7RuleOutput{})
	pulumi.RegisterOutputType(L7RuleArrayOutput{})
	pulumi.RegisterOutputType(L7RuleMapOutput{})
}