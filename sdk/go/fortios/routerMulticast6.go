// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IPv6 multicast.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewRouterMulticast6(ctx, "trname", &fortios.RouterMulticast6Args{
//				MulticastPmtu:    pulumi.String("disable"),
//				MulticastRouting: pulumi.String("disable"),
//				PimSmGlobal: &fortios.RouterMulticast6PimSmGlobalArgs{
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
//
// ## Import
//
// # Router Multicast6 can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerMulticast6:RouterMulticast6 labelname RouterMulticast6
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerMulticast6:RouterMulticast6 labelname RouterMulticast6
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterMulticast6 struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces RouterMulticast6InterfaceArrayOutput `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu pulumi.StringOutput `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringOutput `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal RouterMulticast6PimSmGlobalOutput `pulumi:"pimSmGlobal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterMulticast6 registers a new resource with the given unique name, arguments, and options.
func NewRouterMulticast6(ctx *pulumi.Context,
	name string, args *RouterMulticast6Args, opts ...pulumi.ResourceOption) (*RouterMulticast6, error) {
	if args == nil {
		args = &RouterMulticast6Args{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource RouterMulticast6
	err := ctx.RegisterResource("fortios:index/routerMulticast6:RouterMulticast6", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterMulticast6 gets an existing RouterMulticast6 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterMulticast6(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterMulticast6State, opts ...pulumi.ResourceOption) (*RouterMulticast6, error) {
	var resource RouterMulticast6
	err := ctx.ReadResource("fortios:index/routerMulticast6:RouterMulticast6", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterMulticast6 resources.
type routerMulticast6State struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces []RouterMulticast6Interface `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu *string `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting *string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal *RouterMulticast6PimSmGlobal `pulumi:"pimSmGlobal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterMulticast6State struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces RouterMulticast6InterfaceArrayInput
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu pulumi.StringPtrInput
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringPtrInput
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal RouterMulticast6PimSmGlobalPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterMulticast6State) ElementType() reflect.Type {
	return reflect.TypeOf((*routerMulticast6State)(nil)).Elem()
}

type routerMulticast6Args struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces []RouterMulticast6Interface `pulumi:"interfaces"`
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu *string `pulumi:"multicastPmtu"`
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting *string `pulumi:"multicastRouting"`
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal *RouterMulticast6PimSmGlobal `pulumi:"pimSmGlobal"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterMulticast6 resource.
type RouterMulticast6Args struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
	Interfaces RouterMulticast6InterfaceArrayInput
	// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
	MulticastPmtu pulumi.StringPtrInput
	// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
	MulticastRouting pulumi.StringPtrInput
	// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
	PimSmGlobal RouterMulticast6PimSmGlobalPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterMulticast6Args) ElementType() reflect.Type {
	return reflect.TypeOf((*routerMulticast6Args)(nil)).Elem()
}

type RouterMulticast6Input interface {
	pulumi.Input

	ToRouterMulticast6Output() RouterMulticast6Output
	ToRouterMulticast6OutputWithContext(ctx context.Context) RouterMulticast6Output
}

func (*RouterMulticast6) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterMulticast6)(nil)).Elem()
}

func (i *RouterMulticast6) ToRouterMulticast6Output() RouterMulticast6Output {
	return i.ToRouterMulticast6OutputWithContext(context.Background())
}

func (i *RouterMulticast6) ToRouterMulticast6OutputWithContext(ctx context.Context) RouterMulticast6Output {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticast6Output)
}

// RouterMulticast6ArrayInput is an input type that accepts RouterMulticast6Array and RouterMulticast6ArrayOutput values.
// You can construct a concrete instance of `RouterMulticast6ArrayInput` via:
//
//	RouterMulticast6Array{ RouterMulticast6Args{...} }
type RouterMulticast6ArrayInput interface {
	pulumi.Input

	ToRouterMulticast6ArrayOutput() RouterMulticast6ArrayOutput
	ToRouterMulticast6ArrayOutputWithContext(context.Context) RouterMulticast6ArrayOutput
}

type RouterMulticast6Array []RouterMulticast6Input

func (RouterMulticast6Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterMulticast6)(nil)).Elem()
}

func (i RouterMulticast6Array) ToRouterMulticast6ArrayOutput() RouterMulticast6ArrayOutput {
	return i.ToRouterMulticast6ArrayOutputWithContext(context.Background())
}

func (i RouterMulticast6Array) ToRouterMulticast6ArrayOutputWithContext(ctx context.Context) RouterMulticast6ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticast6ArrayOutput)
}

// RouterMulticast6MapInput is an input type that accepts RouterMulticast6Map and RouterMulticast6MapOutput values.
// You can construct a concrete instance of `RouterMulticast6MapInput` via:
//
//	RouterMulticast6Map{ "key": RouterMulticast6Args{...} }
type RouterMulticast6MapInput interface {
	pulumi.Input

	ToRouterMulticast6MapOutput() RouterMulticast6MapOutput
	ToRouterMulticast6MapOutputWithContext(context.Context) RouterMulticast6MapOutput
}

type RouterMulticast6Map map[string]RouterMulticast6Input

func (RouterMulticast6Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterMulticast6)(nil)).Elem()
}

func (i RouterMulticast6Map) ToRouterMulticast6MapOutput() RouterMulticast6MapOutput {
	return i.ToRouterMulticast6MapOutputWithContext(context.Background())
}

func (i RouterMulticast6Map) ToRouterMulticast6MapOutputWithContext(ctx context.Context) RouterMulticast6MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMulticast6MapOutput)
}

type RouterMulticast6Output struct{ *pulumi.OutputState }

func (RouterMulticast6Output) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterMulticast6)(nil)).Elem()
}

func (o RouterMulticast6Output) ToRouterMulticast6Output() RouterMulticast6Output {
	return o
}

func (o RouterMulticast6Output) ToRouterMulticast6OutputWithContext(ctx context.Context) RouterMulticast6Output {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o RouterMulticast6Output) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Protocol Independent Multicast (PIM) interfaces. The structure of `interface` block is documented below.
func (o RouterMulticast6Output) Interfaces() RouterMulticast6InterfaceArrayOutput {
	return o.ApplyT(func(v *RouterMulticast6) RouterMulticast6InterfaceArrayOutput { return v.Interfaces }).(RouterMulticast6InterfaceArrayOutput)
}

// Enable/disable PMTU for IPv6 multicast. Valid values: `enable`, `disable`.
func (o RouterMulticast6Output) MulticastPmtu() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringOutput { return v.MulticastPmtu }).(pulumi.StringOutput)
}

// Enable/disable IPv6 multicast routing. Valid values: `enable`, `disable`.
func (o RouterMulticast6Output) MulticastRouting() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringOutput { return v.MulticastRouting }).(pulumi.StringOutput)
}

// PIM sparse-mode global settings. The structure of `pimSmGlobal` block is documented below.
func (o RouterMulticast6Output) PimSmGlobal() RouterMulticast6PimSmGlobalOutput {
	return o.ApplyT(func(v *RouterMulticast6) RouterMulticast6PimSmGlobalOutput { return v.PimSmGlobal }).(RouterMulticast6PimSmGlobalOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterMulticast6Output) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterMulticast6) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterMulticast6ArrayOutput struct{ *pulumi.OutputState }

func (RouterMulticast6ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterMulticast6)(nil)).Elem()
}

func (o RouterMulticast6ArrayOutput) ToRouterMulticast6ArrayOutput() RouterMulticast6ArrayOutput {
	return o
}

func (o RouterMulticast6ArrayOutput) ToRouterMulticast6ArrayOutputWithContext(ctx context.Context) RouterMulticast6ArrayOutput {
	return o
}

func (o RouterMulticast6ArrayOutput) Index(i pulumi.IntInput) RouterMulticast6Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterMulticast6 {
		return vs[0].([]*RouterMulticast6)[vs[1].(int)]
	}).(RouterMulticast6Output)
}

type RouterMulticast6MapOutput struct{ *pulumi.OutputState }

func (RouterMulticast6MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterMulticast6)(nil)).Elem()
}

func (o RouterMulticast6MapOutput) ToRouterMulticast6MapOutput() RouterMulticast6MapOutput {
	return o
}

func (o RouterMulticast6MapOutput) ToRouterMulticast6MapOutputWithContext(ctx context.Context) RouterMulticast6MapOutput {
	return o
}

func (o RouterMulticast6MapOutput) MapIndex(k pulumi.StringInput) RouterMulticast6Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterMulticast6 {
		return vs[0].(map[string]*RouterMulticast6)[vs[1].(string)]
	}).(RouterMulticast6Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMulticast6Input)(nil)).Elem(), &RouterMulticast6{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMulticast6ArrayInput)(nil)).Elem(), RouterMulticast6Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMulticast6MapInput)(nil)).Elem(), RouterMulticast6Map{})
	pulumi.RegisterOutputType(RouterMulticast6Output{})
	pulumi.RegisterOutputType(RouterMulticast6ArrayOutput{})
	pulumi.RegisterOutputType(RouterMulticast6MapOutput{})
}
