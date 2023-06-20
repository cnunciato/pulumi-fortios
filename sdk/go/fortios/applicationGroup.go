// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure firewall application groups.
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
//			_, err := fortios.NewApplicationGroup(ctx, "trname", &fortios.ApplicationGroupArgs{
//				Categories: fortios.ApplicationGroupCategoryArray{
//					&fortios.ApplicationGroupCategoryArgs{
//						Id: pulumi.Int(2),
//					},
//				},
//				Comment: pulumi.String("group1 test"),
//				Type:    pulumi.String("category"),
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
// # Application Group can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/applicationGroup:ApplicationGroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/applicationGroup:ApplicationGroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type ApplicationGroup struct {
	pulumi.CustomResourceState

	// Application ID list. The structure of `application` block is documented below.
	Applications ApplicationGroupApplicationArrayOutput `pulumi:"applications"`
	// Application behavior filter.
	Behavior pulumi.StringOutput `pulumi:"behavior"`
	// Application category ID list. The structure of `category` block is documented below.
	Categories ApplicationGroupCategoryArrayOutput `pulumi:"categories"`
	// Comment
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Application group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
	Popularity pulumi.StringOutput `pulumi:"popularity"`
	// Application protocol filter.
	Protocols pulumi.StringOutput `pulumi:"protocols"`
	// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
	Risks ApplicationGroupRiskArrayOutput `pulumi:"risks"`
	// Application technology filter.
	Technology pulumi.StringOutput `pulumi:"technology"`
	// Application group type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Application vendor filter.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
}

// NewApplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		args = &ApplicationGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationGroup
	err := ctx.RegisterResource("fortios:index/applicationGroup:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGroup gets an existing ApplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("fortios:index/applicationGroup:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGroup resources.
type applicationGroupState struct {
	// Application ID list. The structure of `application` block is documented below.
	Applications []ApplicationGroupApplication `pulumi:"applications"`
	// Application behavior filter.
	Behavior *string `pulumi:"behavior"`
	// Application category ID list. The structure of `category` block is documented below.
	Categories []ApplicationGroupCategory `pulumi:"categories"`
	// Comment
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Application group name.
	Name *string `pulumi:"name"`
	// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
	Popularity *string `pulumi:"popularity"`
	// Application protocol filter.
	Protocols *string `pulumi:"protocols"`
	// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
	Risks []ApplicationGroupRisk `pulumi:"risks"`
	// Application technology filter.
	Technology *string `pulumi:"technology"`
	// Application group type.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Application vendor filter.
	Vendor *string `pulumi:"vendor"`
}

type ApplicationGroupState struct {
	// Application ID list. The structure of `application` block is documented below.
	Applications ApplicationGroupApplicationArrayInput
	// Application behavior filter.
	Behavior pulumi.StringPtrInput
	// Application category ID list. The structure of `category` block is documented below.
	Categories ApplicationGroupCategoryArrayInput
	// Comment
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Application group name.
	Name pulumi.StringPtrInput
	// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
	Popularity pulumi.StringPtrInput
	// Application protocol filter.
	Protocols pulumi.StringPtrInput
	// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
	Risks ApplicationGroupRiskArrayInput
	// Application technology filter.
	Technology pulumi.StringPtrInput
	// Application group type.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Application vendor filter.
	Vendor pulumi.StringPtrInput
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	// Application ID list. The structure of `application` block is documented below.
	Applications []ApplicationGroupApplication `pulumi:"applications"`
	// Application behavior filter.
	Behavior *string `pulumi:"behavior"`
	// Application category ID list. The structure of `category` block is documented below.
	Categories []ApplicationGroupCategory `pulumi:"categories"`
	// Comment
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Application group name.
	Name *string `pulumi:"name"`
	// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
	Popularity *string `pulumi:"popularity"`
	// Application protocol filter.
	Protocols *string `pulumi:"protocols"`
	// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
	Risks []ApplicationGroupRisk `pulumi:"risks"`
	// Application technology filter.
	Technology *string `pulumi:"technology"`
	// Application group type.
	Type *string `pulumi:"type"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Application vendor filter.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a ApplicationGroup resource.
type ApplicationGroupArgs struct {
	// Application ID list. The structure of `application` block is documented below.
	Applications ApplicationGroupApplicationArrayInput
	// Application behavior filter.
	Behavior pulumi.StringPtrInput
	// Application category ID list. The structure of `category` block is documented below.
	Categories ApplicationGroupCategoryArrayInput
	// Comment
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Application group name.
	Name pulumi.StringPtrInput
	// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
	Popularity pulumi.StringPtrInput
	// Application protocol filter.
	Protocols pulumi.StringPtrInput
	// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
	Risks ApplicationGroupRiskArrayInput
	// Application technology filter.
	Technology pulumi.StringPtrInput
	// Application group type.
	Type pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Application vendor filter.
	Vendor pulumi.StringPtrInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

// ApplicationGroupArrayInput is an input type that accepts ApplicationGroupArray and ApplicationGroupArrayOutput values.
// You can construct a concrete instance of `ApplicationGroupArrayInput` via:
//
//	ApplicationGroupArray{ ApplicationGroupArgs{...} }
type ApplicationGroupArrayInput interface {
	pulumi.Input

	ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput
	ToApplicationGroupArrayOutputWithContext(context.Context) ApplicationGroupArrayOutput
}

type ApplicationGroupArray []ApplicationGroupInput

func (ApplicationGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGroup)(nil)).Elem()
}

func (i ApplicationGroupArray) ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput {
	return i.ToApplicationGroupArrayOutputWithContext(context.Background())
}

func (i ApplicationGroupArray) ToApplicationGroupArrayOutputWithContext(ctx context.Context) ApplicationGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupArrayOutput)
}

// ApplicationGroupMapInput is an input type that accepts ApplicationGroupMap and ApplicationGroupMapOutput values.
// You can construct a concrete instance of `ApplicationGroupMapInput` via:
//
//	ApplicationGroupMap{ "key": ApplicationGroupArgs{...} }
type ApplicationGroupMapInput interface {
	pulumi.Input

	ToApplicationGroupMapOutput() ApplicationGroupMapOutput
	ToApplicationGroupMapOutputWithContext(context.Context) ApplicationGroupMapOutput
}

type ApplicationGroupMap map[string]ApplicationGroupInput

func (ApplicationGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGroup)(nil)).Elem()
}

func (i ApplicationGroupMap) ToApplicationGroupMapOutput() ApplicationGroupMapOutput {
	return i.ToApplicationGroupMapOutputWithContext(context.Background())
}

func (i ApplicationGroupMap) ToApplicationGroupMapOutputWithContext(ctx context.Context) ApplicationGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupMapOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

// Application ID list. The structure of `application` block is documented below.
func (o ApplicationGroupOutput) Applications() ApplicationGroupApplicationArrayOutput {
	return o.ApplyT(func(v *ApplicationGroup) ApplicationGroupApplicationArrayOutput { return v.Applications }).(ApplicationGroupApplicationArrayOutput)
}

// Application behavior filter.
func (o ApplicationGroupOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Behavior }).(pulumi.StringOutput)
}

// Application category ID list. The structure of `category` block is documented below.
func (o ApplicationGroupOutput) Categories() ApplicationGroupCategoryArrayOutput {
	return o.ApplyT(func(v *ApplicationGroup) ApplicationGroupCategoryArrayOutput { return v.Categories }).(ApplicationGroupCategoryArrayOutput)
}

// Comment
func (o ApplicationGroupOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o ApplicationGroupOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Application group name.
func (o ApplicationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Application popularity filter (1 - 5, from least to most popular). Valid values: `1`, `2`, `3`, `4`, `5`.
func (o ApplicationGroupOutput) Popularity() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Popularity }).(pulumi.StringOutput)
}

// Application protocol filter.
func (o ApplicationGroupOutput) Protocols() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Protocols }).(pulumi.StringOutput)
}

// Risk, or impact, of allowing traffic from this application to occur (1 - 5; Low, Elevated, Medium, High, and Critical). The structure of `risk` block is documented below.
func (o ApplicationGroupOutput) Risks() ApplicationGroupRiskArrayOutput {
	return o.ApplyT(func(v *ApplicationGroup) ApplicationGroupRiskArrayOutput { return v.Risks }).(ApplicationGroupRiskArrayOutput)
}

// Application technology filter.
func (o ApplicationGroupOutput) Technology() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Technology }).(pulumi.StringOutput)
}

// Application group type.
func (o ApplicationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ApplicationGroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Application vendor filter.
func (o ApplicationGroupOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

type ApplicationGroupArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupArrayOutput) ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput {
	return o
}

func (o ApplicationGroupArrayOutput) ToApplicationGroupArrayOutputWithContext(ctx context.Context) ApplicationGroupArrayOutput {
	return o
}

func (o ApplicationGroupArrayOutput) Index(i pulumi.IntInput) ApplicationGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationGroup {
		return vs[0].([]*ApplicationGroup)[vs[1].(int)]
	}).(ApplicationGroupOutput)
}

type ApplicationGroupMapOutput struct{ *pulumi.OutputState }

func (ApplicationGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupMapOutput) ToApplicationGroupMapOutput() ApplicationGroupMapOutput {
	return o
}

func (o ApplicationGroupMapOutput) ToApplicationGroupMapOutputWithContext(ctx context.Context) ApplicationGroupMapOutput {
	return o
}

func (o ApplicationGroupMapOutput) MapIndex(k pulumi.StringInput) ApplicationGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationGroup {
		return vs[0].(map[string]*ApplicationGroup)[vs[1].(string)]
	}).(ApplicationGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupInput)(nil)).Elem(), &ApplicationGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupArrayInput)(nil)).Elem(), ApplicationGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupMapInput)(nil)).Elem(), ApplicationGroupMap{})
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
	pulumi.RegisterOutputType(ApplicationGroupArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGroupMapOutput{})
}
