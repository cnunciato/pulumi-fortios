// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure community lists.
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
//			_, err := fortios.NewRouterCommunitylist(ctx, "trname", &fortios.RouterCommunitylistArgs{
//				Rules: fortios.RouterCommunitylistRuleArray{
//					&fortios.RouterCommunitylistRuleArgs{
//						Action: pulumi.String("deny"),
//						Match:  pulumi.String("123:234 345:456"),
//						Regexp: pulumi.String("123:234 345:456"),
//					},
//				},
//				Type: pulumi.String("standard"),
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
// # Router CommunityList can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/routerCommunitylist:RouterCommunitylist labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/routerCommunitylist:RouterCommunitylist labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type RouterCommunitylist struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Community list name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Community list rule. The structure of `rule` block is documented below.
	Rules RouterCommunitylistRuleArrayOutput `pulumi:"rules"`
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRouterCommunitylist registers a new resource with the given unique name, arguments, and options.
func NewRouterCommunitylist(ctx *pulumi.Context,
	name string, args *RouterCommunitylistArgs, opts ...pulumi.ResourceOption) (*RouterCommunitylist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RouterCommunitylist
	err := ctx.RegisterResource("fortios:index/routerCommunitylist:RouterCommunitylist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterCommunitylist gets an existing RouterCommunitylist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterCommunitylist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterCommunitylistState, opts ...pulumi.ResourceOption) (*RouterCommunitylist, error) {
	var resource RouterCommunitylist
	err := ctx.ReadResource("fortios:index/routerCommunitylist:RouterCommunitylist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterCommunitylist resources.
type routerCommunitylistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Community list name.
	Name *string `pulumi:"name"`
	// Community list rule. The structure of `rule` block is documented below.
	Rules []RouterCommunitylistRule `pulumi:"rules"`
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RouterCommunitylistState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Community list name.
	Name pulumi.StringPtrInput
	// Community list rule. The structure of `rule` block is documented below.
	Rules RouterCommunitylistRuleArrayInput
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterCommunitylistState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerCommunitylistState)(nil)).Elem()
}

type routerCommunitylistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Community list name.
	Name *string `pulumi:"name"`
	// Community list rule. The structure of `rule` block is documented below.
	Rules []RouterCommunitylistRule `pulumi:"rules"`
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a RouterCommunitylist resource.
type RouterCommunitylistArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Community list name.
	Name pulumi.StringPtrInput
	// Community list rule. The structure of `rule` block is documented below.
	Rules RouterCommunitylistRuleArrayInput
	// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
	Type pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RouterCommunitylistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerCommunitylistArgs)(nil)).Elem()
}

type RouterCommunitylistInput interface {
	pulumi.Input

	ToRouterCommunitylistOutput() RouterCommunitylistOutput
	ToRouterCommunitylistOutputWithContext(ctx context.Context) RouterCommunitylistOutput
}

func (*RouterCommunitylist) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterCommunitylist)(nil)).Elem()
}

func (i *RouterCommunitylist) ToRouterCommunitylistOutput() RouterCommunitylistOutput {
	return i.ToRouterCommunitylistOutputWithContext(context.Background())
}

func (i *RouterCommunitylist) ToRouterCommunitylistOutputWithContext(ctx context.Context) RouterCommunitylistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterCommunitylistOutput)
}

// RouterCommunitylistArrayInput is an input type that accepts RouterCommunitylistArray and RouterCommunitylistArrayOutput values.
// You can construct a concrete instance of `RouterCommunitylistArrayInput` via:
//
//	RouterCommunitylistArray{ RouterCommunitylistArgs{...} }
type RouterCommunitylistArrayInput interface {
	pulumi.Input

	ToRouterCommunitylistArrayOutput() RouterCommunitylistArrayOutput
	ToRouterCommunitylistArrayOutputWithContext(context.Context) RouterCommunitylistArrayOutput
}

type RouterCommunitylistArray []RouterCommunitylistInput

func (RouterCommunitylistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterCommunitylist)(nil)).Elem()
}

func (i RouterCommunitylistArray) ToRouterCommunitylistArrayOutput() RouterCommunitylistArrayOutput {
	return i.ToRouterCommunitylistArrayOutputWithContext(context.Background())
}

func (i RouterCommunitylistArray) ToRouterCommunitylistArrayOutputWithContext(ctx context.Context) RouterCommunitylistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterCommunitylistArrayOutput)
}

// RouterCommunitylistMapInput is an input type that accepts RouterCommunitylistMap and RouterCommunitylistMapOutput values.
// You can construct a concrete instance of `RouterCommunitylistMapInput` via:
//
//	RouterCommunitylistMap{ "key": RouterCommunitylistArgs{...} }
type RouterCommunitylistMapInput interface {
	pulumi.Input

	ToRouterCommunitylistMapOutput() RouterCommunitylistMapOutput
	ToRouterCommunitylistMapOutputWithContext(context.Context) RouterCommunitylistMapOutput
}

type RouterCommunitylistMap map[string]RouterCommunitylistInput

func (RouterCommunitylistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterCommunitylist)(nil)).Elem()
}

func (i RouterCommunitylistMap) ToRouterCommunitylistMapOutput() RouterCommunitylistMapOutput {
	return i.ToRouterCommunitylistMapOutputWithContext(context.Background())
}

func (i RouterCommunitylistMap) ToRouterCommunitylistMapOutputWithContext(ctx context.Context) RouterCommunitylistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterCommunitylistMapOutput)
}

type RouterCommunitylistOutput struct{ *pulumi.OutputState }

func (RouterCommunitylistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterCommunitylist)(nil)).Elem()
}

func (o RouterCommunitylistOutput) ToRouterCommunitylistOutput() RouterCommunitylistOutput {
	return o
}

func (o RouterCommunitylistOutput) ToRouterCommunitylistOutputWithContext(ctx context.Context) RouterCommunitylistOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o RouterCommunitylistOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterCommunitylist) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Community list name.
func (o RouterCommunitylistOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterCommunitylist) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Community list rule. The structure of `rule` block is documented below.
func (o RouterCommunitylistOutput) Rules() RouterCommunitylistRuleArrayOutput {
	return o.ApplyT(func(v *RouterCommunitylist) RouterCommunitylistRuleArrayOutput { return v.Rules }).(RouterCommunitylistRuleArrayOutput)
}

// Community list type (standard or expanded). Valid values: `standard`, `expanded`.
func (o RouterCommunitylistOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterCommunitylist) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RouterCommunitylistOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterCommunitylist) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RouterCommunitylistArrayOutput struct{ *pulumi.OutputState }

func (RouterCommunitylistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterCommunitylist)(nil)).Elem()
}

func (o RouterCommunitylistArrayOutput) ToRouterCommunitylistArrayOutput() RouterCommunitylistArrayOutput {
	return o
}

func (o RouterCommunitylistArrayOutput) ToRouterCommunitylistArrayOutputWithContext(ctx context.Context) RouterCommunitylistArrayOutput {
	return o
}

func (o RouterCommunitylistArrayOutput) Index(i pulumi.IntInput) RouterCommunitylistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterCommunitylist {
		return vs[0].([]*RouterCommunitylist)[vs[1].(int)]
	}).(RouterCommunitylistOutput)
}

type RouterCommunitylistMapOutput struct{ *pulumi.OutputState }

func (RouterCommunitylistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterCommunitylist)(nil)).Elem()
}

func (o RouterCommunitylistMapOutput) ToRouterCommunitylistMapOutput() RouterCommunitylistMapOutput {
	return o
}

func (o RouterCommunitylistMapOutput) ToRouterCommunitylistMapOutputWithContext(ctx context.Context) RouterCommunitylistMapOutput {
	return o
}

func (o RouterCommunitylistMapOutput) MapIndex(k pulumi.StringInput) RouterCommunitylistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterCommunitylist {
		return vs[0].(map[string]*RouterCommunitylist)[vs[1].(string)]
	}).(RouterCommunitylistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterCommunitylistInput)(nil)).Elem(), &RouterCommunitylist{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterCommunitylistArrayInput)(nil)).Elem(), RouterCommunitylistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterCommunitylistMapInput)(nil)).Elem(), RouterCommunitylistMap{})
	pulumi.RegisterOutputType(RouterCommunitylistOutput{})
	pulumi.RegisterOutputType(RouterCommunitylistArrayOutput{})
	pulumi.RegisterOutputType(RouterCommunitylistMapOutput{})
}
