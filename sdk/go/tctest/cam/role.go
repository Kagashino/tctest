// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Role struct {
	pulumi.CustomResourceState

	// Indicates whether the CAM role can login or not.
	ConsoleLogin pulumi.BoolPtrOutput `pulumi:"consoleLogin"`
	// Create time of the CAM role.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the CAM role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document pulumi.StringOutput `pulumi:"document"`
	// Name of CAM role.
	Name pulumi.StringOutput `pulumi:"name"`
	// The last update time of the CAM role.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Document == nil {
		return nil, errors.New("invalid value for required argument 'Document'")
	}
	var resource Role
	err := ctx.RegisterResource("tctest:Cam/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("tctest:Cam/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin *bool `pulumi:"consoleLogin"`
	// Create time of the CAM role.
	CreateTime *string `pulumi:"createTime"`
	// Description of the CAM role.
	Description *string `pulumi:"description"`
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document *string `pulumi:"document"`
	// Name of CAM role.
	Name *string `pulumi:"name"`
	// The last update time of the CAM role.
	UpdateTime *string `pulumi:"updateTime"`
}

type RoleState struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin pulumi.BoolPtrInput
	// Create time of the CAM role.
	CreateTime pulumi.StringPtrInput
	// Description of the CAM role.
	Description pulumi.StringPtrInput
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document pulumi.StringPtrInput
	// Name of CAM role.
	Name pulumi.StringPtrInput
	// The last update time of the CAM role.
	UpdateTime pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin *bool `pulumi:"consoleLogin"`
	// Description of the CAM role.
	Description *string `pulumi:"description"`
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document string `pulumi:"document"`
	// Name of CAM role.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// Indicates whether the CAM role can login or not.
	ConsoleLogin pulumi.BoolPtrInput
	// Description of the CAM role.
	Description pulumi.StringPtrInput
	// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
	// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
	// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
	// replaced with the uin it stands for.
	Document pulumi.StringInput
	// Name of CAM role.
	Name pulumi.StringPtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//          RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//          RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// Indicates whether the CAM role can login or not.
func (o RoleOutput) ConsoleLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.BoolPtrOutput { return v.ConsoleLogin }).(pulumi.BoolPtrOutput)
}

// Create time of the CAM role.
func (o RoleOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the CAM role.
func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Document of the CAM role. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604).
// There are some notes when using this para in terraform: 1. The elements in json claimed supporting two types as `string`
// and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when appears, it must be
// replaced with the uin it stands for.
func (o RoleOutput) Document() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Document }).(pulumi.StringOutput)
}

// Name of CAM role.
func (o RoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The last update time of the CAM role.
func (o RoleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}