// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure dedicated management.
//
// ## Import
//
// # System DedicatedMgmt can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt labelname SystemDedicatedMgmt
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt labelname SystemDedicatedMgmt
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemDedicatedmgmt struct {
	pulumi.CustomResourceState

	// Default gateway for dedicated management interface.
	DefaultGateway pulumi.StringOutput `pulumi:"defaultGateway"`
	// DHCP end IP for dedicated management.
	DhcpEndIp pulumi.StringOutput `pulumi:"dhcpEndIp"`
	// DHCP netmask.
	DhcpNetmask pulumi.StringOutput `pulumi:"dhcpNetmask"`
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer pulumi.StringOutput `pulumi:"dhcpServer"`
	// DHCP start IP for dedicated management.
	DhcpStartIp pulumi.StringOutput `pulumi:"dhcpStartIp"`
	// Dedicated management interface.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemDedicatedmgmt registers a new resource with the given unique name, arguments, and options.
func NewSystemDedicatedmgmt(ctx *pulumi.Context,
	name string, args *SystemDedicatedmgmtArgs, opts ...pulumi.ResourceOption) (*SystemDedicatedmgmt, error) {
	if args == nil {
		args = &SystemDedicatedmgmtArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemDedicatedmgmt
	err := ctx.RegisterResource("fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemDedicatedmgmt gets an existing SystemDedicatedmgmt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemDedicatedmgmt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemDedicatedmgmtState, opts ...pulumi.ResourceOption) (*SystemDedicatedmgmt, error) {
	var resource SystemDedicatedmgmt
	err := ctx.ReadResource("fortios:index/systemDedicatedmgmt:SystemDedicatedmgmt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemDedicatedmgmt resources.
type systemDedicatedmgmtState struct {
	// Default gateway for dedicated management interface.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// DHCP end IP for dedicated management.
	DhcpEndIp *string `pulumi:"dhcpEndIp"`
	// DHCP netmask.
	DhcpNetmask *string `pulumi:"dhcpNetmask"`
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer *string `pulumi:"dhcpServer"`
	// DHCP start IP for dedicated management.
	DhcpStartIp *string `pulumi:"dhcpStartIp"`
	// Dedicated management interface.
	Interface *string `pulumi:"interface"`
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemDedicatedmgmtState struct {
	// Default gateway for dedicated management interface.
	DefaultGateway pulumi.StringPtrInput
	// DHCP end IP for dedicated management.
	DhcpEndIp pulumi.StringPtrInput
	// DHCP netmask.
	DhcpNetmask pulumi.StringPtrInput
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer pulumi.StringPtrInput
	// DHCP start IP for dedicated management.
	DhcpStartIp pulumi.StringPtrInput
	// Dedicated management interface.
	Interface pulumi.StringPtrInput
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDedicatedmgmtState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDedicatedmgmtState)(nil)).Elem()
}

type systemDedicatedmgmtArgs struct {
	// Default gateway for dedicated management interface.
	DefaultGateway *string `pulumi:"defaultGateway"`
	// DHCP end IP for dedicated management.
	DhcpEndIp *string `pulumi:"dhcpEndIp"`
	// DHCP netmask.
	DhcpNetmask *string `pulumi:"dhcpNetmask"`
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer *string `pulumi:"dhcpServer"`
	// DHCP start IP for dedicated management.
	DhcpStartIp *string `pulumi:"dhcpStartIp"`
	// Dedicated management interface.
	Interface *string `pulumi:"interface"`
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemDedicatedmgmt resource.
type SystemDedicatedmgmtArgs struct {
	// Default gateway for dedicated management interface.
	DefaultGateway pulumi.StringPtrInput
	// DHCP end IP for dedicated management.
	DhcpEndIp pulumi.StringPtrInput
	// DHCP netmask.
	DhcpNetmask pulumi.StringPtrInput
	// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
	DhcpServer pulumi.StringPtrInput
	// DHCP start IP for dedicated management.
	DhcpStartIp pulumi.StringPtrInput
	// Dedicated management interface.
	Interface pulumi.StringPtrInput
	// Enable/disable dedicated management. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemDedicatedmgmtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemDedicatedmgmtArgs)(nil)).Elem()
}

type SystemDedicatedmgmtInput interface {
	pulumi.Input

	ToSystemDedicatedmgmtOutput() SystemDedicatedmgmtOutput
	ToSystemDedicatedmgmtOutputWithContext(ctx context.Context) SystemDedicatedmgmtOutput
}

func (*SystemDedicatedmgmt) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDedicatedmgmt)(nil)).Elem()
}

func (i *SystemDedicatedmgmt) ToSystemDedicatedmgmtOutput() SystemDedicatedmgmtOutput {
	return i.ToSystemDedicatedmgmtOutputWithContext(context.Background())
}

func (i *SystemDedicatedmgmt) ToSystemDedicatedmgmtOutputWithContext(ctx context.Context) SystemDedicatedmgmtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDedicatedmgmtOutput)
}

// SystemDedicatedmgmtArrayInput is an input type that accepts SystemDedicatedmgmtArray and SystemDedicatedmgmtArrayOutput values.
// You can construct a concrete instance of `SystemDedicatedmgmtArrayInput` via:
//
//	SystemDedicatedmgmtArray{ SystemDedicatedmgmtArgs{...} }
type SystemDedicatedmgmtArrayInput interface {
	pulumi.Input

	ToSystemDedicatedmgmtArrayOutput() SystemDedicatedmgmtArrayOutput
	ToSystemDedicatedmgmtArrayOutputWithContext(context.Context) SystemDedicatedmgmtArrayOutput
}

type SystemDedicatedmgmtArray []SystemDedicatedmgmtInput

func (SystemDedicatedmgmtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDedicatedmgmt)(nil)).Elem()
}

func (i SystemDedicatedmgmtArray) ToSystemDedicatedmgmtArrayOutput() SystemDedicatedmgmtArrayOutput {
	return i.ToSystemDedicatedmgmtArrayOutputWithContext(context.Background())
}

func (i SystemDedicatedmgmtArray) ToSystemDedicatedmgmtArrayOutputWithContext(ctx context.Context) SystemDedicatedmgmtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDedicatedmgmtArrayOutput)
}

// SystemDedicatedmgmtMapInput is an input type that accepts SystemDedicatedmgmtMap and SystemDedicatedmgmtMapOutput values.
// You can construct a concrete instance of `SystemDedicatedmgmtMapInput` via:
//
//	SystemDedicatedmgmtMap{ "key": SystemDedicatedmgmtArgs{...} }
type SystemDedicatedmgmtMapInput interface {
	pulumi.Input

	ToSystemDedicatedmgmtMapOutput() SystemDedicatedmgmtMapOutput
	ToSystemDedicatedmgmtMapOutputWithContext(context.Context) SystemDedicatedmgmtMapOutput
}

type SystemDedicatedmgmtMap map[string]SystemDedicatedmgmtInput

func (SystemDedicatedmgmtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDedicatedmgmt)(nil)).Elem()
}

func (i SystemDedicatedmgmtMap) ToSystemDedicatedmgmtMapOutput() SystemDedicatedmgmtMapOutput {
	return i.ToSystemDedicatedmgmtMapOutputWithContext(context.Background())
}

func (i SystemDedicatedmgmtMap) ToSystemDedicatedmgmtMapOutputWithContext(ctx context.Context) SystemDedicatedmgmtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDedicatedmgmtMapOutput)
}

type SystemDedicatedmgmtOutput struct{ *pulumi.OutputState }

func (SystemDedicatedmgmtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDedicatedmgmt)(nil)).Elem()
}

func (o SystemDedicatedmgmtOutput) ToSystemDedicatedmgmtOutput() SystemDedicatedmgmtOutput {
	return o
}

func (o SystemDedicatedmgmtOutput) ToSystemDedicatedmgmtOutputWithContext(ctx context.Context) SystemDedicatedmgmtOutput {
	return o
}

// Default gateway for dedicated management interface.
func (o SystemDedicatedmgmtOutput) DefaultGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.DefaultGateway }).(pulumi.StringOutput)
}

// DHCP end IP for dedicated management.
func (o SystemDedicatedmgmtOutput) DhcpEndIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.DhcpEndIp }).(pulumi.StringOutput)
}

// DHCP netmask.
func (o SystemDedicatedmgmtOutput) DhcpNetmask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.DhcpNetmask }).(pulumi.StringOutput)
}

// Enable/disable DHCP server on management interface. Valid values: `enable`, `disable`.
func (o SystemDedicatedmgmtOutput) DhcpServer() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.DhcpServer }).(pulumi.StringOutput)
}

// DHCP start IP for dedicated management.
func (o SystemDedicatedmgmtOutput) DhcpStartIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.DhcpStartIp }).(pulumi.StringOutput)
}

// Dedicated management interface.
func (o SystemDedicatedmgmtOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Enable/disable dedicated management. Valid values: `enable`, `disable`.
func (o SystemDedicatedmgmtOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemDedicatedmgmtOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDedicatedmgmt) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemDedicatedmgmtArrayOutput struct{ *pulumi.OutputState }

func (SystemDedicatedmgmtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemDedicatedmgmt)(nil)).Elem()
}

func (o SystemDedicatedmgmtArrayOutput) ToSystemDedicatedmgmtArrayOutput() SystemDedicatedmgmtArrayOutput {
	return o
}

func (o SystemDedicatedmgmtArrayOutput) ToSystemDedicatedmgmtArrayOutputWithContext(ctx context.Context) SystemDedicatedmgmtArrayOutput {
	return o
}

func (o SystemDedicatedmgmtArrayOutput) Index(i pulumi.IntInput) SystemDedicatedmgmtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemDedicatedmgmt {
		return vs[0].([]*SystemDedicatedmgmt)[vs[1].(int)]
	}).(SystemDedicatedmgmtOutput)
}

type SystemDedicatedmgmtMapOutput struct{ *pulumi.OutputState }

func (SystemDedicatedmgmtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemDedicatedmgmt)(nil)).Elem()
}

func (o SystemDedicatedmgmtMapOutput) ToSystemDedicatedmgmtMapOutput() SystemDedicatedmgmtMapOutput {
	return o
}

func (o SystemDedicatedmgmtMapOutput) ToSystemDedicatedmgmtMapOutputWithContext(ctx context.Context) SystemDedicatedmgmtMapOutput {
	return o
}

func (o SystemDedicatedmgmtMapOutput) MapIndex(k pulumi.StringInput) SystemDedicatedmgmtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemDedicatedmgmt {
		return vs[0].(map[string]*SystemDedicatedmgmt)[vs[1].(string)]
	}).(SystemDedicatedmgmtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDedicatedmgmtInput)(nil)).Elem(), &SystemDedicatedmgmt{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDedicatedmgmtArrayInput)(nil)).Elem(), SystemDedicatedmgmtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDedicatedmgmtMapInput)(nil)).Elem(), SystemDedicatedmgmtMap{})
	pulumi.RegisterOutputType(SystemDedicatedmgmtOutput{})
	pulumi.RegisterOutputType(SystemDedicatedmgmtArrayOutput{})
	pulumi.RegisterOutputType(SystemDedicatedmgmtMapOutput{})
}