// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package router

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure IPv6 multicast.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := router.NewMulticast6(ctx, "trname", &router.Multicast6Args{
//				MulticastPmtu:    pulumi.String("disable"),
//				MulticastRouting: pulumi.String("disable"),
//				PimSmGlobal: &router.Multicast6PimSmGlobalArgs{
//					RegisterRateLimit: pulumi.Int(0),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Router Multicast6 can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:router/multicast6:Multicast6 labelname RouterMulticast6
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:router/multicast6:Multicast6 labelname RouterMulticast6
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Multicast6 struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces Multicast6InterfaceArrayOutput `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu pulumi.StringOutput `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringOutput `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal Multicast6PimSmGlobalOutput `pulumi:"pimSmGlobal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewMulticast6 registers a new resource with the given unique name, arguments, and options.
func NewMulticast6(ctx *pulumi.Context,
	name string, args *Multicast6Args, opts ...pulumi.ResourceOption) (*Multicast6, error) {
	if args == nil {
		args = &Multicast6Args{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Multicast6
	err := ctx.RegisterResource("fortios:router/multicast6:Multicast6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMulticast6 gets an existing Multicast6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMulticast6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Multicast6State, opts ...pulumi.ResourceOption) (*Multicast6, error) {
	var resource Multicast6
	err := ctx.ReadResource("fortios:router/multicast6:Multicast6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Multicast6 resources.
type multicast6State struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces []Multicast6Interface `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu *string `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting *string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal *Multicast6PimSmGlobal `pulumi:"pimSmGlobal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Multicast6State struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces Multicast6InterfaceArrayInput
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu pulumi.StringPtrInput
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringPtrInput
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal Multicast6PimSmGlobalPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Multicast6State) ElementType() reflect.Type {
	return reflect.TypeOf((*multicast6State)(nil)).Elem()
}

type multicast6Args struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces []Multicast6Interface `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu *string `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting *string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal *Multicast6PimSmGlobal `pulumi:"pimSmGlobal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Multicast6 resource.
type Multicast6Args struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces Multicast6InterfaceArrayInput
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu pulumi.StringPtrInput
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringPtrInput
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal Multicast6PimSmGlobalPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Multicast6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*multicast6Args)(nil)).Elem()
}

type Multicast6Input interface {
	pulumi.Input

	ToMulticast6Output() Multicast6Output
	ToMulticast6OutputWithContext(ctx context.Context) Multicast6Output
}

func (*Multicast6) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicast6)(nil)).Elem()
}

func (i *Multicast6) ToMulticast6Output() Multicast6Output {
	return i.ToMulticast6OutputWithContext(context.Background())
}

func (i *Multicast6) ToMulticast6OutputWithContext(ctx context.Context) Multicast6Output {
	return pulumi.ToOutputWithContext(ctx, i).(Multicast6Output)
}

// Multicast6ArrayInput is an input type that accepts Multicast6Array and Multicast6ArrayOutput values.
// You can construct a concrete instance of `Multicast6ArrayInput` via:
//
//	Multicast6Array{ Multicast6Args{...} }
type Multicast6ArrayInput interface {
	pulumi.Input

	ToMulticast6ArrayOutput() Multicast6ArrayOutput
	ToMulticast6ArrayOutputWithContext(context.Context) Multicast6ArrayOutput
}

type Multicast6Array []Multicast6Input

func (Multicast6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicast6)(nil)).Elem()
}

func (i Multicast6Array) ToMulticast6ArrayOutput() Multicast6ArrayOutput {
	return i.ToMulticast6ArrayOutputWithContext(context.Background())
}

func (i Multicast6Array) ToMulticast6ArrayOutputWithContext(ctx context.Context) Multicast6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Multicast6ArrayOutput)
}

// Multicast6MapInput is an input type that accepts Multicast6Map and Multicast6MapOutput values.
// You can construct a concrete instance of `Multicast6MapInput` via:
//
//	Multicast6Map{ "key": Multicast6Args{...} }
type Multicast6MapInput interface {
	pulumi.Input

	ToMulticast6MapOutput() Multicast6MapOutput
	ToMulticast6MapOutputWithContext(context.Context) Multicast6MapOutput
}

type Multicast6Map map[string]Multicast6Input

func (Multicast6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicast6)(nil)).Elem()
}

func (i Multicast6Map) ToMulticast6MapOutput() Multicast6MapOutput {
	return i.ToMulticast6MapOutputWithContext(context.Background())
}

func (i Multicast6Map) ToMulticast6MapOutputWithContext(ctx context.Context) Multicast6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Multicast6MapOutput)
}

type Multicast6Output struct{ *pulumi.OutputState }

func (Multicast6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Multicast6)(nil)).Elem()
}

func (o Multicast6Output) ToMulticast6Output() Multicast6Output {
	return o
}

func (o Multicast6Output) ToMulticast6OutputWithContext(ctx context.Context) Multicast6Output {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o Multicast6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicast6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
func (o Multicast6Output) Interfaces() Multicast6InterfaceArrayOutput {
	return o.ApplyT(func(v *Multicast6) Multicast6InterfaceArrayOutput { return v.Interfaces }).(Multicast6InterfaceArrayOutput)
}

// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
func (o Multicast6Output) MulticastPmtu() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicast6) pulumi.StringOutput { return v.MulticastPmtu }).(pulumi.StringOutput)
}

// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
func (o Multicast6Output) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *Multicast6) pulumi.StringOutput { return v.MulticastRouting }).(pulumi.StringOutput)
}

// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
func (o Multicast6Output) PimSmGlobal() Multicast6PimSmGlobalOutput {
	return o.ApplyT(func(v *Multicast6) Multicast6PimSmGlobalOutput { return v.PimSmGlobal }).(Multicast6PimSmGlobalOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Multicast6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multicast6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Multicast6ArrayOutput struct{ *pulumi.OutputState }

func (Multicast6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Multicast6)(nil)).Elem()
}

func (o Multicast6ArrayOutput) ToMulticast6ArrayOutput() Multicast6ArrayOutput {
	return o
}

func (o Multicast6ArrayOutput) ToMulticast6ArrayOutputWithContext(ctx context.Context) Multicast6ArrayOutput {
	return o
}

func (o Multicast6ArrayOutput) Index(i pulumi.IntInput) Multicast6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Multicast6 {
		return vs[0].([]*Multicast6)[vs[1].(int)]
	}).(Multicast6Output)
}

type Multicast6MapOutput struct{ *pulumi.OutputState }

func (Multicast6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Multicast6)(nil)).Elem()
}

func (o Multicast6MapOutput) ToMulticast6MapOutput() Multicast6MapOutput {
	return o
}

func (o Multicast6MapOutput) ToMulticast6MapOutputWithContext(ctx context.Context) Multicast6MapOutput {
	return o
}

func (o Multicast6MapOutput) MapIndex(k pulumi.StringInput) Multicast6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Multicast6 {
		return vs[0].(map[string]*Multicast6)[vs[1].(string)]
	}).(Multicast6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Multicast6Input)(nil)).Elem(), &Multicast6{})
	pulumi.RegisterInputType(reflect.TypeOf((*Multicast6ArrayInput)(nil)).Elem(), Multicast6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Multicast6MapInput)(nil)).Elem(), Multicast6Map{})
	pulumi.RegisterOutputType(Multicast6Output{})
	pulumi.RegisterOutputType(Multicast6ArrayOutput{})
	pulumi.RegisterOutputType(Multicast6MapOutput{})
}