// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Define country table. Applies to FortiOS Version `>= 6.4.0`.
//
// ## Import
//
// Firewall Country can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/country:Country labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/country:Country labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Country struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Country ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Country name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Region ID list. The structure of `region` block is documented below.
	Regions CountryRegionArrayOutput `pulumi:"regions"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCountry registers a new resource with the given unique name, arguments, and options.
func NewCountry(ctx *pulumi.Context,
	name string, args *CountryArgs, opts ...pulumi.ResourceOption) (*Country, error) {
	if args == nil {
		args = &CountryArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Country
	err := ctx.RegisterResource("fortios:firewall/country:Country", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCountry gets an existing Country resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCountry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CountryState, opts ...pulumi.ResourceOption) (*Country, error) {
	var resource Country
	err := ctx.ReadResource("fortios:firewall/country:Country", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Country resources.
type countryState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Country ID.
	Fosid *int `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Region ID list. The structure of `region` block is documented below.
	Regions []CountryRegion `pulumi:"regions"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CountryState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Country ID.
	Fosid pulumi.IntPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Region ID list. The structure of `region` block is documented below.
	Regions CountryRegionArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CountryState) ElementType() reflect.Type {
	return reflect.TypeOf((*countryState)(nil)).Elem()
}

type countryArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Country ID.
	Fosid *int `pulumi:"fosid"`
	// Country name.
	Name *string `pulumi:"name"`
	// Region ID list. The structure of `region` block is documented below.
	Regions []CountryRegion `pulumi:"regions"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Country resource.
type CountryArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Country ID.
	Fosid pulumi.IntPtrInput
	// Country name.
	Name pulumi.StringPtrInput
	// Region ID list. The structure of `region` block is documented below.
	Regions CountryRegionArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CountryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*countryArgs)(nil)).Elem()
}

type CountryInput interface {
	pulumi.Input

	ToCountryOutput() CountryOutput
	ToCountryOutputWithContext(ctx context.Context) CountryOutput
}

func (*Country) ElementType() reflect.Type {
	return reflect.TypeOf((**Country)(nil)).Elem()
}

func (i *Country) ToCountryOutput() CountryOutput {
	return i.ToCountryOutputWithContext(context.Background())
}

func (i *Country) ToCountryOutputWithContext(ctx context.Context) CountryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CountryOutput)
}

// CountryArrayInput is an input type that accepts CountryArray and CountryArrayOutput values.
// You can construct a concrete instance of `CountryArrayInput` via:
//
//	CountryArray{ CountryArgs{...} }
type CountryArrayInput interface {
	pulumi.Input

	ToCountryArrayOutput() CountryArrayOutput
	ToCountryArrayOutputWithContext(context.Context) CountryArrayOutput
}

type CountryArray []CountryInput

func (CountryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Country)(nil)).Elem()
}

func (i CountryArray) ToCountryArrayOutput() CountryArrayOutput {
	return i.ToCountryArrayOutputWithContext(context.Background())
}

func (i CountryArray) ToCountryArrayOutputWithContext(ctx context.Context) CountryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CountryArrayOutput)
}

// CountryMapInput is an input type that accepts CountryMap and CountryMapOutput values.
// You can construct a concrete instance of `CountryMapInput` via:
//
//	CountryMap{ "key": CountryArgs{...} }
type CountryMapInput interface {
	pulumi.Input

	ToCountryMapOutput() CountryMapOutput
	ToCountryMapOutputWithContext(context.Context) CountryMapOutput
}

type CountryMap map[string]CountryInput

func (CountryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Country)(nil)).Elem()
}

func (i CountryMap) ToCountryMapOutput() CountryMapOutput {
	return i.ToCountryMapOutputWithContext(context.Background())
}

func (i CountryMap) ToCountryMapOutputWithContext(ctx context.Context) CountryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CountryMapOutput)
}

type CountryOutput struct{ *pulumi.OutputState }

func (CountryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Country)(nil)).Elem()
}

func (o CountryOutput) ToCountryOutput() CountryOutput {
	return o
}

func (o CountryOutput) ToCountryOutputWithContext(ctx context.Context) CountryOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o CountryOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Country) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Country ID.
func (o CountryOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Country) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Country name.
func (o CountryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Country) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Region ID list. The structure of `region` block is documented below.
func (o CountryOutput) Regions() CountryRegionArrayOutput {
	return o.ApplyT(func(v *Country) CountryRegionArrayOutput { return v.Regions }).(CountryRegionArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CountryOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Country) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CountryArrayOutput struct{ *pulumi.OutputState }

func (CountryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Country)(nil)).Elem()
}

func (o CountryArrayOutput) ToCountryArrayOutput() CountryArrayOutput {
	return o
}

func (o CountryArrayOutput) ToCountryArrayOutputWithContext(ctx context.Context) CountryArrayOutput {
	return o
}

func (o CountryArrayOutput) Index(i pulumi.IntInput) CountryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Country {
		return vs[0].([]*Country)[vs[1].(int)]
	}).(CountryOutput)
}

type CountryMapOutput struct{ *pulumi.OutputState }

func (CountryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Country)(nil)).Elem()
}

func (o CountryMapOutput) ToCountryMapOutput() CountryMapOutput {
	return o
}

func (o CountryMapOutput) ToCountryMapOutputWithContext(ctx context.Context) CountryMapOutput {
	return o
}

func (o CountryMapOutput) MapIndex(k pulumi.StringInput) CountryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Country {
		return vs[0].(map[string]*Country)[vs[1].(string)]
	}).(CountryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CountryInput)(nil)).Elem(), &Country{})
	pulumi.RegisterInputType(reflect.TypeOf((*CountryArrayInput)(nil)).Elem(), CountryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CountryMapInput)(nil)).Elem(), CountryMap{})
	pulumi.RegisterOutputType(CountryOutput{})
	pulumi.RegisterOutputType(CountryArrayOutput{})
	pulumi.RegisterOutputType(CountryMapOutput{})
}