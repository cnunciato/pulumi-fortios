// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch Auto-Config custom QoS policy.
//
// ## Import
//
// # SwitchControllerAutoConfig Custom can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerautoconfigCustom struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Auto-Config FortiLink or ISL/ICL interface name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Switch binding list. The structure of `switchBinding` block is documented below.
	SwitchBindings SwitchcontrollerautoconfigCustomSwitchBindingArrayOutput `pulumi:"switchBindings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerautoconfigCustom registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerautoconfigCustom(ctx *pulumi.Context,
	name string, args *SwitchcontrollerautoconfigCustomArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerautoconfigCustom, error) {
	if args == nil {
		args = &SwitchcontrollerautoconfigCustomArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerautoconfigCustom
	err := ctx.RegisterResource("fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerautoconfigCustom gets an existing SwitchcontrollerautoconfigCustom resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerautoconfigCustom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerautoconfigCustomState, opts ...pulumi.ResourceOption) (*SwitchcontrollerautoconfigCustom, error) {
	var resource SwitchcontrollerautoconfigCustom
	err := ctx.ReadResource("fortios:index/switchcontrollerautoconfigCustom:SwitchcontrollerautoconfigCustom", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerautoconfigCustom resources.
type switchcontrollerautoconfigCustomState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Auto-Config FortiLink or ISL/ICL interface name.
	Name *string `pulumi:"name"`
	// Switch binding list. The structure of `switchBinding` block is documented below.
	SwitchBindings []SwitchcontrollerautoconfigCustomSwitchBinding `pulumi:"switchBindings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerautoconfigCustomState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Auto-Config FortiLink or ISL/ICL interface name.
	Name pulumi.StringPtrInput
	// Switch binding list. The structure of `switchBinding` block is documented below.
	SwitchBindings SwitchcontrollerautoconfigCustomSwitchBindingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerautoconfigCustomState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerautoconfigCustomState)(nil)).Elem()
}

type switchcontrollerautoconfigCustomArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Auto-Config FortiLink or ISL/ICL interface name.
	Name *string `pulumi:"name"`
	// Switch binding list. The structure of `switchBinding` block is documented below.
	SwitchBindings []SwitchcontrollerautoconfigCustomSwitchBinding `pulumi:"switchBindings"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerautoconfigCustom resource.
type SwitchcontrollerautoconfigCustomArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Auto-Config FortiLink or ISL/ICL interface name.
	Name pulumi.StringPtrInput
	// Switch binding list. The structure of `switchBinding` block is documented below.
	SwitchBindings SwitchcontrollerautoconfigCustomSwitchBindingArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerautoconfigCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerautoconfigCustomArgs)(nil)).Elem()
}

type SwitchcontrollerautoconfigCustomInput interface {
	pulumi.Input

	ToSwitchcontrollerautoconfigCustomOutput() SwitchcontrollerautoconfigCustomOutput
	ToSwitchcontrollerautoconfigCustomOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomOutput
}

func (*SwitchcontrollerautoconfigCustom) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerautoconfigCustom)(nil)).Elem()
}

func (i *SwitchcontrollerautoconfigCustom) ToSwitchcontrollerautoconfigCustomOutput() SwitchcontrollerautoconfigCustomOutput {
	return i.ToSwitchcontrollerautoconfigCustomOutputWithContext(context.Background())
}

func (i *SwitchcontrollerautoconfigCustom) ToSwitchcontrollerautoconfigCustomOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerautoconfigCustomOutput)
}

// SwitchcontrollerautoconfigCustomArrayInput is an input type that accepts SwitchcontrollerautoconfigCustomArray and SwitchcontrollerautoconfigCustomArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerautoconfigCustomArrayInput` via:
//
//	SwitchcontrollerautoconfigCustomArray{ SwitchcontrollerautoconfigCustomArgs{...} }
type SwitchcontrollerautoconfigCustomArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerautoconfigCustomArrayOutput() SwitchcontrollerautoconfigCustomArrayOutput
	ToSwitchcontrollerautoconfigCustomArrayOutputWithContext(context.Context) SwitchcontrollerautoconfigCustomArrayOutput
}

type SwitchcontrollerautoconfigCustomArray []SwitchcontrollerautoconfigCustomInput

func (SwitchcontrollerautoconfigCustomArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerautoconfigCustom)(nil)).Elem()
}

func (i SwitchcontrollerautoconfigCustomArray) ToSwitchcontrollerautoconfigCustomArrayOutput() SwitchcontrollerautoconfigCustomArrayOutput {
	return i.ToSwitchcontrollerautoconfigCustomArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerautoconfigCustomArray) ToSwitchcontrollerautoconfigCustomArrayOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerautoconfigCustomArrayOutput)
}

// SwitchcontrollerautoconfigCustomMapInput is an input type that accepts SwitchcontrollerautoconfigCustomMap and SwitchcontrollerautoconfigCustomMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerautoconfigCustomMapInput` via:
//
//	SwitchcontrollerautoconfigCustomMap{ "key": SwitchcontrollerautoconfigCustomArgs{...} }
type SwitchcontrollerautoconfigCustomMapInput interface {
	pulumi.Input

	ToSwitchcontrollerautoconfigCustomMapOutput() SwitchcontrollerautoconfigCustomMapOutput
	ToSwitchcontrollerautoconfigCustomMapOutputWithContext(context.Context) SwitchcontrollerautoconfigCustomMapOutput
}

type SwitchcontrollerautoconfigCustomMap map[string]SwitchcontrollerautoconfigCustomInput

func (SwitchcontrollerautoconfigCustomMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerautoconfigCustom)(nil)).Elem()
}

func (i SwitchcontrollerautoconfigCustomMap) ToSwitchcontrollerautoconfigCustomMapOutput() SwitchcontrollerautoconfigCustomMapOutput {
	return i.ToSwitchcontrollerautoconfigCustomMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerautoconfigCustomMap) ToSwitchcontrollerautoconfigCustomMapOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerautoconfigCustomMapOutput)
}

type SwitchcontrollerautoconfigCustomOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerautoconfigCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerautoconfigCustom)(nil)).Elem()
}

func (o SwitchcontrollerautoconfigCustomOutput) ToSwitchcontrollerautoconfigCustomOutput() SwitchcontrollerautoconfigCustomOutput {
	return o
}

func (o SwitchcontrollerautoconfigCustomOutput) ToSwitchcontrollerautoconfigCustomOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchcontrollerautoconfigCustomOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerautoconfigCustom) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Auto-Config FortiLink or ISL/ICL interface name.
func (o SwitchcontrollerautoconfigCustomOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerautoconfigCustom) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Switch binding list. The structure of `switchBinding` block is documented below.
func (o SwitchcontrollerautoconfigCustomOutput) SwitchBindings() SwitchcontrollerautoconfigCustomSwitchBindingArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerautoconfigCustom) SwitchcontrollerautoconfigCustomSwitchBindingArrayOutput {
		return v.SwitchBindings
	}).(SwitchcontrollerautoconfigCustomSwitchBindingArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerautoconfigCustomOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerautoconfigCustom) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerautoconfigCustomArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerautoconfigCustomArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerautoconfigCustom)(nil)).Elem()
}

func (o SwitchcontrollerautoconfigCustomArrayOutput) ToSwitchcontrollerautoconfigCustomArrayOutput() SwitchcontrollerautoconfigCustomArrayOutput {
	return o
}

func (o SwitchcontrollerautoconfigCustomArrayOutput) ToSwitchcontrollerautoconfigCustomArrayOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomArrayOutput {
	return o
}

func (o SwitchcontrollerautoconfigCustomArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerautoconfigCustomOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerautoconfigCustom {
		return vs[0].([]*SwitchcontrollerautoconfigCustom)[vs[1].(int)]
	}).(SwitchcontrollerautoconfigCustomOutput)
}

type SwitchcontrollerautoconfigCustomMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerautoconfigCustomMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerautoconfigCustom)(nil)).Elem()
}

func (o SwitchcontrollerautoconfigCustomMapOutput) ToSwitchcontrollerautoconfigCustomMapOutput() SwitchcontrollerautoconfigCustomMapOutput {
	return o
}

func (o SwitchcontrollerautoconfigCustomMapOutput) ToSwitchcontrollerautoconfigCustomMapOutputWithContext(ctx context.Context) SwitchcontrollerautoconfigCustomMapOutput {
	return o
}

func (o SwitchcontrollerautoconfigCustomMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerautoconfigCustomOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerautoconfigCustom {
		return vs[0].(map[string]*SwitchcontrollerautoconfigCustom)[vs[1].(string)]
	}).(SwitchcontrollerautoconfigCustomOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerautoconfigCustomInput)(nil)).Elem(), &SwitchcontrollerautoconfigCustom{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerautoconfigCustomArrayInput)(nil)).Elem(), SwitchcontrollerautoconfigCustomArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerautoconfigCustomMapInput)(nil)).Elem(), SwitchcontrollerautoconfigCustomMap{})
	pulumi.RegisterOutputType(SwitchcontrollerautoconfigCustomOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerautoconfigCustomArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerautoconfigCustomMapOutput{})
}