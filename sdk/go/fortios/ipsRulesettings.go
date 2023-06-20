// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPS rule setting.
//
// ## Import
//
// # Ips RuleSettings can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/ipsRulesettings:IpsRulesettings labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/ipsRulesettings:IpsRulesettings labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type IpsRulesettings struct {
	pulumi.CustomResourceState

	// Rule ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIpsRulesettings registers a new resource with the given unique name, arguments, and options.
func NewIpsRulesettings(ctx *pulumi.Context,
	name string, args *IpsRulesettingsArgs, opts ...pulumi.ResourceOption) (*IpsRulesettings, error) {
	if args == nil {
		args = &IpsRulesettingsArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource IpsRulesettings
	err := ctx.RegisterResource("fortios:index/ipsRulesettings:IpsRulesettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpsRulesettings gets an existing IpsRulesettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpsRulesettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpsRulesettingsState, opts ...pulumi.ResourceOption) (*IpsRulesettings, error) {
	var resource IpsRulesettings
	err := ctx.ReadResource("fortios:index/ipsRulesettings:IpsRulesettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpsRulesettings resources.
type ipsRulesettingsState struct {
	// Rule ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IpsRulesettingsState struct {
	// Rule ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsRulesettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsRulesettingsState)(nil)).Elem()
}

type ipsRulesettingsArgs struct {
	// Rule ID.
	Fosid *int `pulumi:"fosid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a IpsRulesettings resource.
type IpsRulesettingsArgs struct {
	// Rule ID.
	Fosid pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IpsRulesettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipsRulesettingsArgs)(nil)).Elem()
}

type IpsRulesettingsInput interface {
	pulumi.Input

	ToIpsRulesettingsOutput() IpsRulesettingsOutput
	ToIpsRulesettingsOutputWithContext(ctx context.Context) IpsRulesettingsOutput
}

func (*IpsRulesettings) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsRulesettings)(nil)).Elem()
}

func (i *IpsRulesettings) ToIpsRulesettingsOutput() IpsRulesettingsOutput {
	return i.ToIpsRulesettingsOutputWithContext(context.Background())
}

func (i *IpsRulesettings) ToIpsRulesettingsOutputWithContext(ctx context.Context) IpsRulesettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRulesettingsOutput)
}

// IpsRulesettingsArrayInput is an input type that accepts IpsRulesettingsArray and IpsRulesettingsArrayOutput values.
// You can construct a concrete instance of `IpsRulesettingsArrayInput` via:
//
//	IpsRulesettingsArray{ IpsRulesettingsArgs{...} }
type IpsRulesettingsArrayInput interface {
	pulumi.Input

	ToIpsRulesettingsArrayOutput() IpsRulesettingsArrayOutput
	ToIpsRulesettingsArrayOutputWithContext(context.Context) IpsRulesettingsArrayOutput
}

type IpsRulesettingsArray []IpsRulesettingsInput

func (IpsRulesettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsRulesettings)(nil)).Elem()
}

func (i IpsRulesettingsArray) ToIpsRulesettingsArrayOutput() IpsRulesettingsArrayOutput {
	return i.ToIpsRulesettingsArrayOutputWithContext(context.Background())
}

func (i IpsRulesettingsArray) ToIpsRulesettingsArrayOutputWithContext(ctx context.Context) IpsRulesettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRulesettingsArrayOutput)
}

// IpsRulesettingsMapInput is an input type that accepts IpsRulesettingsMap and IpsRulesettingsMapOutput values.
// You can construct a concrete instance of `IpsRulesettingsMapInput` via:
//
//	IpsRulesettingsMap{ "key": IpsRulesettingsArgs{...} }
type IpsRulesettingsMapInput interface {
	pulumi.Input

	ToIpsRulesettingsMapOutput() IpsRulesettingsMapOutput
	ToIpsRulesettingsMapOutputWithContext(context.Context) IpsRulesettingsMapOutput
}

type IpsRulesettingsMap map[string]IpsRulesettingsInput

func (IpsRulesettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsRulesettings)(nil)).Elem()
}

func (i IpsRulesettingsMap) ToIpsRulesettingsMapOutput() IpsRulesettingsMapOutput {
	return i.ToIpsRulesettingsMapOutputWithContext(context.Background())
}

func (i IpsRulesettingsMap) ToIpsRulesettingsMapOutputWithContext(ctx context.Context) IpsRulesettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsRulesettingsMapOutput)
}

type IpsRulesettingsOutput struct{ *pulumi.OutputState }

func (IpsRulesettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpsRulesettings)(nil)).Elem()
}

func (o IpsRulesettingsOutput) ToIpsRulesettingsOutput() IpsRulesettingsOutput {
	return o
}

func (o IpsRulesettingsOutput) ToIpsRulesettingsOutputWithContext(ctx context.Context) IpsRulesettingsOutput {
	return o
}

// Rule ID.
func (o IpsRulesettingsOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *IpsRulesettings) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IpsRulesettingsOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpsRulesettings) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IpsRulesettingsArrayOutput struct{ *pulumi.OutputState }

func (IpsRulesettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpsRulesettings)(nil)).Elem()
}

func (o IpsRulesettingsArrayOutput) ToIpsRulesettingsArrayOutput() IpsRulesettingsArrayOutput {
	return o
}

func (o IpsRulesettingsArrayOutput) ToIpsRulesettingsArrayOutputWithContext(ctx context.Context) IpsRulesettingsArrayOutput {
	return o
}

func (o IpsRulesettingsArrayOutput) Index(i pulumi.IntInput) IpsRulesettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpsRulesettings {
		return vs[0].([]*IpsRulesettings)[vs[1].(int)]
	}).(IpsRulesettingsOutput)
}

type IpsRulesettingsMapOutput struct{ *pulumi.OutputState }

func (IpsRulesettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpsRulesettings)(nil)).Elem()
}

func (o IpsRulesettingsMapOutput) ToIpsRulesettingsMapOutput() IpsRulesettingsMapOutput {
	return o
}

func (o IpsRulesettingsMapOutput) ToIpsRulesettingsMapOutputWithContext(ctx context.Context) IpsRulesettingsMapOutput {
	return o
}

func (o IpsRulesettingsMapOutput) MapIndex(k pulumi.StringInput) IpsRulesettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpsRulesettings {
		return vs[0].(map[string]*IpsRulesettings)[vs[1].(string)]
	}).(IpsRulesettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRulesettingsInput)(nil)).Elem(), &IpsRulesettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRulesettingsArrayInput)(nil)).Elem(), IpsRulesettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpsRulesettingsMapInput)(nil)).Elem(), IpsRulesettingsMap{})
	pulumi.RegisterOutputType(IpsRulesettingsOutput{})
	pulumi.RegisterOutputType(IpsRulesettingsArrayOutput{})
	pulumi.RegisterOutputType(IpsRulesettingsMapOutput{})
}
