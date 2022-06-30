// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcaplus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClustersList struct {
	ApiAccessId           string `pulumi:"apiAccessId"`
	ApiAccessIp           string `pulumi:"apiAccessIp"`
	ApiAccessPort         int    `pulumi:"apiAccessPort"`
	ClusterId             string `pulumi:"clusterId"`
	ClusterName           string `pulumi:"clusterName"`
	CreateTime            string `pulumi:"createTime"`
	IdlType               string `pulumi:"idlType"`
	NetworkType           string `pulumi:"networkType"`
	OldPasswordExpireTime string `pulumi:"oldPasswordExpireTime"`
	Password              string `pulumi:"password"`
	PasswordStatus        string `pulumi:"passwordStatus"`
	SubnetId              string `pulumi:"subnetId"`
	VpcId                 string `pulumi:"vpcId"`
}

// ClustersListInput is an input type that accepts ClustersListArgs and ClustersListOutput values.
// You can construct a concrete instance of `ClustersListInput` via:
//
//          ClustersListArgs{...}
type ClustersListInput interface {
	pulumi.Input

	ToClustersListOutput() ClustersListOutput
	ToClustersListOutputWithContext(context.Context) ClustersListOutput
}

type ClustersListArgs struct {
	ApiAccessId           pulumi.StringInput `pulumi:"apiAccessId"`
	ApiAccessIp           pulumi.StringInput `pulumi:"apiAccessIp"`
	ApiAccessPort         pulumi.IntInput    `pulumi:"apiAccessPort"`
	ClusterId             pulumi.StringInput `pulumi:"clusterId"`
	ClusterName           pulumi.StringInput `pulumi:"clusterName"`
	CreateTime            pulumi.StringInput `pulumi:"createTime"`
	IdlType               pulumi.StringInput `pulumi:"idlType"`
	NetworkType           pulumi.StringInput `pulumi:"networkType"`
	OldPasswordExpireTime pulumi.StringInput `pulumi:"oldPasswordExpireTime"`
	Password              pulumi.StringInput `pulumi:"password"`
	PasswordStatus        pulumi.StringInput `pulumi:"passwordStatus"`
	SubnetId              pulumi.StringInput `pulumi:"subnetId"`
	VpcId                 pulumi.StringInput `pulumi:"vpcId"`
}

func (ClustersListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClustersList)(nil)).Elem()
}

func (i ClustersListArgs) ToClustersListOutput() ClustersListOutput {
	return i.ToClustersListOutputWithContext(context.Background())
}

func (i ClustersListArgs) ToClustersListOutputWithContext(ctx context.Context) ClustersListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustersListOutput)
}

// ClustersListArrayInput is an input type that accepts ClustersListArray and ClustersListArrayOutput values.
// You can construct a concrete instance of `ClustersListArrayInput` via:
//
//          ClustersListArray{ ClustersListArgs{...} }
type ClustersListArrayInput interface {
	pulumi.Input

	ToClustersListArrayOutput() ClustersListArrayOutput
	ToClustersListArrayOutputWithContext(context.Context) ClustersListArrayOutput
}

type ClustersListArray []ClustersListInput

func (ClustersListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClustersList)(nil)).Elem()
}

func (i ClustersListArray) ToClustersListArrayOutput() ClustersListArrayOutput {
	return i.ToClustersListArrayOutputWithContext(context.Background())
}

func (i ClustersListArray) ToClustersListArrayOutputWithContext(ctx context.Context) ClustersListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClustersListArrayOutput)
}

type ClustersListOutput struct{ *pulumi.OutputState }

func (ClustersListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClustersList)(nil)).Elem()
}

func (o ClustersListOutput) ToClustersListOutput() ClustersListOutput {
	return o
}

func (o ClustersListOutput) ToClustersListOutputWithContext(ctx context.Context) ClustersListOutput {
	return o
}

func (o ClustersListOutput) ApiAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.ApiAccessId }).(pulumi.StringOutput)
}

func (o ClustersListOutput) ApiAccessIp() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.ApiAccessIp }).(pulumi.StringOutput)
}

func (o ClustersListOutput) ApiAccessPort() pulumi.IntOutput {
	return o.ApplyT(func(v ClustersList) int { return v.ApiAccessPort }).(pulumi.IntOutput)
}

func (o ClustersListOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o ClustersListOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.ClusterName }).(pulumi.StringOutput)
}

func (o ClustersListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o ClustersListOutput) IdlType() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.IdlType }).(pulumi.StringOutput)
}

func (o ClustersListOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.NetworkType }).(pulumi.StringOutput)
}

func (o ClustersListOutput) OldPasswordExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.OldPasswordExpireTime }).(pulumi.StringOutput)
}

func (o ClustersListOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.Password }).(pulumi.StringOutput)
}

func (o ClustersListOutput) PasswordStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.PasswordStatus }).(pulumi.StringOutput)
}

func (o ClustersListOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o ClustersListOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v ClustersList) string { return v.VpcId }).(pulumi.StringOutput)
}

type ClustersListArrayOutput struct{ *pulumi.OutputState }

func (ClustersListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClustersList)(nil)).Elem()
}

func (o ClustersListArrayOutput) ToClustersListArrayOutput() ClustersListArrayOutput {
	return o
}

func (o ClustersListArrayOutput) ToClustersListArrayOutputWithContext(ctx context.Context) ClustersListArrayOutput {
	return o
}

func (o ClustersListArrayOutput) Index(i pulumi.IntInput) ClustersListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClustersList {
		return vs[0].([]ClustersList)[vs[1].(int)]
	}).(ClustersListOutput)
}

type IdlTableInfo struct {
	Error             *string `pulumi:"error"`
	IndexKeySet       *string `pulumi:"indexKeySet"`
	KeyFields         *string `pulumi:"keyFields"`
	SumKeyFieldSize   *int    `pulumi:"sumKeyFieldSize"`
	SumValueFieldSize *int    `pulumi:"sumValueFieldSize"`
	TableName         *string `pulumi:"tableName"`
	ValueFields       *string `pulumi:"valueFields"`
}

// IdlTableInfoInput is an input type that accepts IdlTableInfoArgs and IdlTableInfoOutput values.
// You can construct a concrete instance of `IdlTableInfoInput` via:
//
//          IdlTableInfoArgs{...}
type IdlTableInfoInput interface {
	pulumi.Input

	ToIdlTableInfoOutput() IdlTableInfoOutput
	ToIdlTableInfoOutputWithContext(context.Context) IdlTableInfoOutput
}

type IdlTableInfoArgs struct {
	Error             pulumi.StringPtrInput `pulumi:"error"`
	IndexKeySet       pulumi.StringPtrInput `pulumi:"indexKeySet"`
	KeyFields         pulumi.StringPtrInput `pulumi:"keyFields"`
	SumKeyFieldSize   pulumi.IntPtrInput    `pulumi:"sumKeyFieldSize"`
	SumValueFieldSize pulumi.IntPtrInput    `pulumi:"sumValueFieldSize"`
	TableName         pulumi.StringPtrInput `pulumi:"tableName"`
	ValueFields       pulumi.StringPtrInput `pulumi:"valueFields"`
}

func (IdlTableInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdlTableInfo)(nil)).Elem()
}

func (i IdlTableInfoArgs) ToIdlTableInfoOutput() IdlTableInfoOutput {
	return i.ToIdlTableInfoOutputWithContext(context.Background())
}

func (i IdlTableInfoArgs) ToIdlTableInfoOutputWithContext(ctx context.Context) IdlTableInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdlTableInfoOutput)
}

// IdlTableInfoArrayInput is an input type that accepts IdlTableInfoArray and IdlTableInfoArrayOutput values.
// You can construct a concrete instance of `IdlTableInfoArrayInput` via:
//
//          IdlTableInfoArray{ IdlTableInfoArgs{...} }
type IdlTableInfoArrayInput interface {
	pulumi.Input

	ToIdlTableInfoArrayOutput() IdlTableInfoArrayOutput
	ToIdlTableInfoArrayOutputWithContext(context.Context) IdlTableInfoArrayOutput
}

type IdlTableInfoArray []IdlTableInfoInput

func (IdlTableInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdlTableInfo)(nil)).Elem()
}

func (i IdlTableInfoArray) ToIdlTableInfoArrayOutput() IdlTableInfoArrayOutput {
	return i.ToIdlTableInfoArrayOutputWithContext(context.Background())
}

func (i IdlTableInfoArray) ToIdlTableInfoArrayOutputWithContext(ctx context.Context) IdlTableInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdlTableInfoArrayOutput)
}

type IdlTableInfoOutput struct{ *pulumi.OutputState }

func (IdlTableInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdlTableInfo)(nil)).Elem()
}

func (o IdlTableInfoOutput) ToIdlTableInfoOutput() IdlTableInfoOutput {
	return o
}

func (o IdlTableInfoOutput) ToIdlTableInfoOutputWithContext(ctx context.Context) IdlTableInfoOutput {
	return o
}

func (o IdlTableInfoOutput) Error() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *string { return v.Error }).(pulumi.StringPtrOutput)
}

func (o IdlTableInfoOutput) IndexKeySet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *string { return v.IndexKeySet }).(pulumi.StringPtrOutput)
}

func (o IdlTableInfoOutput) KeyFields() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *string { return v.KeyFields }).(pulumi.StringPtrOutput)
}

func (o IdlTableInfoOutput) SumKeyFieldSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *int { return v.SumKeyFieldSize }).(pulumi.IntPtrOutput)
}

func (o IdlTableInfoOutput) SumValueFieldSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *int { return v.SumValueFieldSize }).(pulumi.IntPtrOutput)
}

func (o IdlTableInfoOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o IdlTableInfoOutput) ValueFields() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdlTableInfo) *string { return v.ValueFields }).(pulumi.StringPtrOutput)
}

type IdlTableInfoArrayOutput struct{ *pulumi.OutputState }

func (IdlTableInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdlTableInfo)(nil)).Elem()
}

func (o IdlTableInfoArrayOutput) ToIdlTableInfoArrayOutput() IdlTableInfoArrayOutput {
	return o
}

func (o IdlTableInfoArrayOutput) ToIdlTableInfoArrayOutputWithContext(ctx context.Context) IdlTableInfoArrayOutput {
	return o
}

func (o IdlTableInfoArrayOutput) Index(i pulumi.IntInput) IdlTableInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdlTableInfo {
		return vs[0].([]IdlTableInfo)[vs[1].(int)]
	}).(IdlTableInfoOutput)
}

type IdlsList struct {
	IdlId string `pulumi:"idlId"`
}

// IdlsListInput is an input type that accepts IdlsListArgs and IdlsListOutput values.
// You can construct a concrete instance of `IdlsListInput` via:
//
//          IdlsListArgs{...}
type IdlsListInput interface {
	pulumi.Input

	ToIdlsListOutput() IdlsListOutput
	ToIdlsListOutputWithContext(context.Context) IdlsListOutput
}

type IdlsListArgs struct {
	IdlId pulumi.StringInput `pulumi:"idlId"`
}

func (IdlsListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdlsList)(nil)).Elem()
}

func (i IdlsListArgs) ToIdlsListOutput() IdlsListOutput {
	return i.ToIdlsListOutputWithContext(context.Background())
}

func (i IdlsListArgs) ToIdlsListOutputWithContext(ctx context.Context) IdlsListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdlsListOutput)
}

// IdlsListArrayInput is an input type that accepts IdlsListArray and IdlsListArrayOutput values.
// You can construct a concrete instance of `IdlsListArrayInput` via:
//
//          IdlsListArray{ IdlsListArgs{...} }
type IdlsListArrayInput interface {
	pulumi.Input

	ToIdlsListArrayOutput() IdlsListArrayOutput
	ToIdlsListArrayOutputWithContext(context.Context) IdlsListArrayOutput
}

type IdlsListArray []IdlsListInput

func (IdlsListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdlsList)(nil)).Elem()
}

func (i IdlsListArray) ToIdlsListArrayOutput() IdlsListArrayOutput {
	return i.ToIdlsListArrayOutputWithContext(context.Background())
}

func (i IdlsListArray) ToIdlsListArrayOutputWithContext(ctx context.Context) IdlsListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdlsListArrayOutput)
}

type IdlsListOutput struct{ *pulumi.OutputState }

func (IdlsListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdlsList)(nil)).Elem()
}

func (o IdlsListOutput) ToIdlsListOutput() IdlsListOutput {
	return o
}

func (o IdlsListOutput) ToIdlsListOutputWithContext(ctx context.Context) IdlsListOutput {
	return o
}

func (o IdlsListOutput) IdlId() pulumi.StringOutput {
	return o.ApplyT(func(v IdlsList) string { return v.IdlId }).(pulumi.StringOutput)
}

type IdlsListArrayOutput struct{ *pulumi.OutputState }

func (IdlsListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdlsList)(nil)).Elem()
}

func (o IdlsListArrayOutput) ToIdlsListArrayOutput() IdlsListArrayOutput {
	return o
}

func (o IdlsListArrayOutput) ToIdlsListArrayOutputWithContext(ctx context.Context) IdlsListArrayOutput {
	return o
}

func (o IdlsListArrayOutput) Index(i pulumi.IntInput) IdlsListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdlsList {
		return vs[0].([]IdlsList)[vs[1].(int)]
	}).(IdlsListOutput)
}

type TableGroupsList struct {
	CreateTime     string `pulumi:"createTime"`
	TableCount     int    `pulumi:"tableCount"`
	TablegroupId   string `pulumi:"tablegroupId"`
	TablegroupName string `pulumi:"tablegroupName"`
	TotalSize      int    `pulumi:"totalSize"`
}

// TableGroupsListInput is an input type that accepts TableGroupsListArgs and TableGroupsListOutput values.
// You can construct a concrete instance of `TableGroupsListInput` via:
//
//          TableGroupsListArgs{...}
type TableGroupsListInput interface {
	pulumi.Input

	ToTableGroupsListOutput() TableGroupsListOutput
	ToTableGroupsListOutputWithContext(context.Context) TableGroupsListOutput
}

type TableGroupsListArgs struct {
	CreateTime     pulumi.StringInput `pulumi:"createTime"`
	TableCount     pulumi.IntInput    `pulumi:"tableCount"`
	TablegroupId   pulumi.StringInput `pulumi:"tablegroupId"`
	TablegroupName pulumi.StringInput `pulumi:"tablegroupName"`
	TotalSize      pulumi.IntInput    `pulumi:"totalSize"`
}

func (TableGroupsListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGroupsList)(nil)).Elem()
}

func (i TableGroupsListArgs) ToTableGroupsListOutput() TableGroupsListOutput {
	return i.ToTableGroupsListOutputWithContext(context.Background())
}

func (i TableGroupsListArgs) ToTableGroupsListOutputWithContext(ctx context.Context) TableGroupsListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGroupsListOutput)
}

// TableGroupsListArrayInput is an input type that accepts TableGroupsListArray and TableGroupsListArrayOutput values.
// You can construct a concrete instance of `TableGroupsListArrayInput` via:
//
//          TableGroupsListArray{ TableGroupsListArgs{...} }
type TableGroupsListArrayInput interface {
	pulumi.Input

	ToTableGroupsListArrayOutput() TableGroupsListArrayOutput
	ToTableGroupsListArrayOutputWithContext(context.Context) TableGroupsListArrayOutput
}

type TableGroupsListArray []TableGroupsListInput

func (TableGroupsListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableGroupsList)(nil)).Elem()
}

func (i TableGroupsListArray) ToTableGroupsListArrayOutput() TableGroupsListArrayOutput {
	return i.ToTableGroupsListArrayOutputWithContext(context.Background())
}

func (i TableGroupsListArray) ToTableGroupsListArrayOutputWithContext(ctx context.Context) TableGroupsListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGroupsListArrayOutput)
}

type TableGroupsListOutput struct{ *pulumi.OutputState }

func (TableGroupsListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGroupsList)(nil)).Elem()
}

func (o TableGroupsListOutput) ToTableGroupsListOutput() TableGroupsListOutput {
	return o
}

func (o TableGroupsListOutput) ToTableGroupsListOutputWithContext(ctx context.Context) TableGroupsListOutput {
	return o
}

func (o TableGroupsListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v TableGroupsList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o TableGroupsListOutput) TableCount() pulumi.IntOutput {
	return o.ApplyT(func(v TableGroupsList) int { return v.TableCount }).(pulumi.IntOutput)
}

func (o TableGroupsListOutput) TablegroupId() pulumi.StringOutput {
	return o.ApplyT(func(v TableGroupsList) string { return v.TablegroupId }).(pulumi.StringOutput)
}

func (o TableGroupsListOutput) TablegroupName() pulumi.StringOutput {
	return o.ApplyT(func(v TableGroupsList) string { return v.TablegroupName }).(pulumi.StringOutput)
}

func (o TableGroupsListOutput) TotalSize() pulumi.IntOutput {
	return o.ApplyT(func(v TableGroupsList) int { return v.TotalSize }).(pulumi.IntOutput)
}

type TableGroupsListArrayOutput struct{ *pulumi.OutputState }

func (TableGroupsListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableGroupsList)(nil)).Elem()
}

func (o TableGroupsListArrayOutput) ToTableGroupsListArrayOutput() TableGroupsListArrayOutput {
	return o
}

func (o TableGroupsListArrayOutput) ToTableGroupsListArrayOutputWithContext(ctx context.Context) TableGroupsListArrayOutput {
	return o
}

func (o TableGroupsListArrayOutput) Index(i pulumi.IntInput) TableGroupsListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TableGroupsList {
		return vs[0].([]TableGroupsList)[vs[1].(int)]
	}).(TableGroupsListOutput)
}

type TablesList struct {
	CreateTime      string `pulumi:"createTime"`
	Description     string `pulumi:"description"`
	Error           string `pulumi:"error"`
	IdlId           string `pulumi:"idlId"`
	ReservedReadCu  int    `pulumi:"reservedReadCu"`
	ReservedVolume  int    `pulumi:"reservedVolume"`
	ReservedWriteCu int    `pulumi:"reservedWriteCu"`
	Status          string `pulumi:"status"`
	TableId         string `pulumi:"tableId"`
	TableIdlType    string `pulumi:"tableIdlType"`
	TableName       string `pulumi:"tableName"`
	TableSize       int    `pulumi:"tableSize"`
	TableType       string `pulumi:"tableType"`
	TablegroupId    string `pulumi:"tablegroupId"`
}

// TablesListInput is an input type that accepts TablesListArgs and TablesListOutput values.
// You can construct a concrete instance of `TablesListInput` via:
//
//          TablesListArgs{...}
type TablesListInput interface {
	pulumi.Input

	ToTablesListOutput() TablesListOutput
	ToTablesListOutputWithContext(context.Context) TablesListOutput
}

type TablesListArgs struct {
	CreateTime      pulumi.StringInput `pulumi:"createTime"`
	Description     pulumi.StringInput `pulumi:"description"`
	Error           pulumi.StringInput `pulumi:"error"`
	IdlId           pulumi.StringInput `pulumi:"idlId"`
	ReservedReadCu  pulumi.IntInput    `pulumi:"reservedReadCu"`
	ReservedVolume  pulumi.IntInput    `pulumi:"reservedVolume"`
	ReservedWriteCu pulumi.IntInput    `pulumi:"reservedWriteCu"`
	Status          pulumi.StringInput `pulumi:"status"`
	TableId         pulumi.StringInput `pulumi:"tableId"`
	TableIdlType    pulumi.StringInput `pulumi:"tableIdlType"`
	TableName       pulumi.StringInput `pulumi:"tableName"`
	TableSize       pulumi.IntInput    `pulumi:"tableSize"`
	TableType       pulumi.StringInput `pulumi:"tableType"`
	TablegroupId    pulumi.StringInput `pulumi:"tablegroupId"`
}

func (TablesListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TablesList)(nil)).Elem()
}

func (i TablesListArgs) ToTablesListOutput() TablesListOutput {
	return i.ToTablesListOutputWithContext(context.Background())
}

func (i TablesListArgs) ToTablesListOutputWithContext(ctx context.Context) TablesListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablesListOutput)
}

// TablesListArrayInput is an input type that accepts TablesListArray and TablesListArrayOutput values.
// You can construct a concrete instance of `TablesListArrayInput` via:
//
//          TablesListArray{ TablesListArgs{...} }
type TablesListArrayInput interface {
	pulumi.Input

	ToTablesListArrayOutput() TablesListArrayOutput
	ToTablesListArrayOutputWithContext(context.Context) TablesListArrayOutput
}

type TablesListArray []TablesListInput

func (TablesListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TablesList)(nil)).Elem()
}

func (i TablesListArray) ToTablesListArrayOutput() TablesListArrayOutput {
	return i.ToTablesListArrayOutputWithContext(context.Background())
}

func (i TablesListArray) ToTablesListArrayOutputWithContext(ctx context.Context) TablesListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablesListArrayOutput)
}

type TablesListOutput struct{ *pulumi.OutputState }

func (TablesListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TablesList)(nil)).Elem()
}

func (o TablesListOutput) ToTablesListOutput() TablesListOutput {
	return o
}

func (o TablesListOutput) ToTablesListOutputWithContext(ctx context.Context) TablesListOutput {
	return o
}

func (o TablesListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o TablesListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.Description }).(pulumi.StringOutput)
}

func (o TablesListOutput) Error() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.Error }).(pulumi.StringOutput)
}

func (o TablesListOutput) IdlId() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.IdlId }).(pulumi.StringOutput)
}

func (o TablesListOutput) ReservedReadCu() pulumi.IntOutput {
	return o.ApplyT(func(v TablesList) int { return v.ReservedReadCu }).(pulumi.IntOutput)
}

func (o TablesListOutput) ReservedVolume() pulumi.IntOutput {
	return o.ApplyT(func(v TablesList) int { return v.ReservedVolume }).(pulumi.IntOutput)
}

func (o TablesListOutput) ReservedWriteCu() pulumi.IntOutput {
	return o.ApplyT(func(v TablesList) int { return v.ReservedWriteCu }).(pulumi.IntOutput)
}

func (o TablesListOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.Status }).(pulumi.StringOutput)
}

func (o TablesListOutput) TableId() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.TableId }).(pulumi.StringOutput)
}

func (o TablesListOutput) TableIdlType() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.TableIdlType }).(pulumi.StringOutput)
}

func (o TablesListOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.TableName }).(pulumi.StringOutput)
}

func (o TablesListOutput) TableSize() pulumi.IntOutput {
	return o.ApplyT(func(v TablesList) int { return v.TableSize }).(pulumi.IntOutput)
}

func (o TablesListOutput) TableType() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.TableType }).(pulumi.StringOutput)
}

func (o TablesListOutput) TablegroupId() pulumi.StringOutput {
	return o.ApplyT(func(v TablesList) string { return v.TablegroupId }).(pulumi.StringOutput)
}

type TablesListArrayOutput struct{ *pulumi.OutputState }

func (TablesListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TablesList)(nil)).Elem()
}

func (o TablesListArrayOutput) ToTablesListArrayOutput() TablesListArrayOutput {
	return o
}

func (o TablesListArrayOutput) ToTablesListArrayOutputWithContext(ctx context.Context) TablesListArrayOutput {
	return o
}

func (o TablesListArrayOutput) Index(i pulumi.IntInput) TablesListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TablesList {
		return vs[0].([]TablesList)[vs[1].(int)]
	}).(TablesListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClustersListInput)(nil)).Elem(), ClustersListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClustersListArrayInput)(nil)).Elem(), ClustersListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdlTableInfoInput)(nil)).Elem(), IdlTableInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdlTableInfoArrayInput)(nil)).Elem(), IdlTableInfoArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdlsListInput)(nil)).Elem(), IdlsListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdlsListArrayInput)(nil)).Elem(), IdlsListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableGroupsListInput)(nil)).Elem(), TableGroupsListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableGroupsListArrayInput)(nil)).Elem(), TableGroupsListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TablesListInput)(nil)).Elem(), TablesListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TablesListArrayInput)(nil)).Elem(), TablesListArray{})
	pulumi.RegisterOutputType(ClustersListOutput{})
	pulumi.RegisterOutputType(ClustersListArrayOutput{})
	pulumi.RegisterOutputType(IdlTableInfoOutput{})
	pulumi.RegisterOutputType(IdlTableInfoArrayOutput{})
	pulumi.RegisterOutputType(IdlsListOutput{})
	pulumi.RegisterOutputType(IdlsListArrayOutput{})
	pulumi.RegisterOutputType(TableGroupsListOutput{})
	pulumi.RegisterOutputType(TableGroupsListArrayOutput{})
	pulumi.RegisterOutputType(TablesListOutput{})
	pulumi.RegisterOutputType(TablesListArrayOutput{})
}