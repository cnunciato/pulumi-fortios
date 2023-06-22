// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiCloud SSO admin users. Applies to FortiOS Version `>= 7.0.1`.
//
// ## Import
//
// # System SsoForticloudAdmin can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemSsoforticloudadmin struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// FortiCloud SSO admin name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms SystemSsoforticloudadminVdomArrayOutput `pulumi:"vdoms"`
}

// NewSystemSsoforticloudadmin registers a new resource with the given unique name, arguments, and options.
func NewSystemSsoforticloudadmin(ctx *pulumi.Context,
	name string, args *SystemSsoforticloudadminArgs, opts ...pulumi.ResourceOption) (*SystemSsoforticloudadmin, error) {
	if args == nil {
		args = &SystemSsoforticloudadminArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemSsoforticloudadmin
	err := ctx.RegisterResource("fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemSsoforticloudadmin gets an existing SystemSsoforticloudadmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemSsoforticloudadmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemSsoforticloudadminState, opts ...pulumi.ResourceOption) (*SystemSsoforticloudadmin, error) {
	var resource SystemSsoforticloudadmin
	err := ctx.ReadResource("fortios:index/systemSsoforticloudadmin:SystemSsoforticloudadmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemSsoforticloudadmin resources.
type systemSsoforticloudadminState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiCloud SSO admin name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms []SystemSsoforticloudadminVdom `pulumi:"vdoms"`
}

type SystemSsoforticloudadminState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiCloud SSO admin name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms SystemSsoforticloudadminVdomArrayInput
}

func (SystemSsoforticloudadminState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSsoforticloudadminState)(nil)).Elem()
}

type systemSsoforticloudadminArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// FortiCloud SSO admin name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms []SystemSsoforticloudadminVdom `pulumi:"vdoms"`
}

// The set of arguments for constructing a SystemSsoforticloudadmin resource.
type SystemSsoforticloudadminArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// FortiCloud SSO admin name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
	Vdoms SystemSsoforticloudadminVdomArrayInput
}

func (SystemSsoforticloudadminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemSsoforticloudadminArgs)(nil)).Elem()
}

type SystemSsoforticloudadminInput interface {
	pulumi.Input

	ToSystemSsoforticloudadminOutput() SystemSsoforticloudadminOutput
	ToSystemSsoforticloudadminOutputWithContext(ctx context.Context) SystemSsoforticloudadminOutput
}

func (*SystemSsoforticloudadmin) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSsoforticloudadmin)(nil)).Elem()
}

func (i *SystemSsoforticloudadmin) ToSystemSsoforticloudadminOutput() SystemSsoforticloudadminOutput {
	return i.ToSystemSsoforticloudadminOutputWithContext(context.Background())
}

func (i *SystemSsoforticloudadmin) ToSystemSsoforticloudadminOutputWithContext(ctx context.Context) SystemSsoforticloudadminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSsoforticloudadminOutput)
}

// SystemSsoforticloudadminArrayInput is an input type that accepts SystemSsoforticloudadminArray and SystemSsoforticloudadminArrayOutput values.
// You can construct a concrete instance of `SystemSsoforticloudadminArrayInput` via:
//
//	SystemSsoforticloudadminArray{ SystemSsoforticloudadminArgs{...} }
type SystemSsoforticloudadminArrayInput interface {
	pulumi.Input

	ToSystemSsoforticloudadminArrayOutput() SystemSsoforticloudadminArrayOutput
	ToSystemSsoforticloudadminArrayOutputWithContext(context.Context) SystemSsoforticloudadminArrayOutput
}

type SystemSsoforticloudadminArray []SystemSsoforticloudadminInput

func (SystemSsoforticloudadminArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSsoforticloudadmin)(nil)).Elem()
}

func (i SystemSsoforticloudadminArray) ToSystemSsoforticloudadminArrayOutput() SystemSsoforticloudadminArrayOutput {
	return i.ToSystemSsoforticloudadminArrayOutputWithContext(context.Background())
}

func (i SystemSsoforticloudadminArray) ToSystemSsoforticloudadminArrayOutputWithContext(ctx context.Context) SystemSsoforticloudadminArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSsoforticloudadminArrayOutput)
}

// SystemSsoforticloudadminMapInput is an input type that accepts SystemSsoforticloudadminMap and SystemSsoforticloudadminMapOutput values.
// You can construct a concrete instance of `SystemSsoforticloudadminMapInput` via:
//
//	SystemSsoforticloudadminMap{ "key": SystemSsoforticloudadminArgs{...} }
type SystemSsoforticloudadminMapInput interface {
	pulumi.Input

	ToSystemSsoforticloudadminMapOutput() SystemSsoforticloudadminMapOutput
	ToSystemSsoforticloudadminMapOutputWithContext(context.Context) SystemSsoforticloudadminMapOutput
}

type SystemSsoforticloudadminMap map[string]SystemSsoforticloudadminInput

func (SystemSsoforticloudadminMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSsoforticloudadmin)(nil)).Elem()
}

func (i SystemSsoforticloudadminMap) ToSystemSsoforticloudadminMapOutput() SystemSsoforticloudadminMapOutput {
	return i.ToSystemSsoforticloudadminMapOutputWithContext(context.Background())
}

func (i SystemSsoforticloudadminMap) ToSystemSsoforticloudadminMapOutputWithContext(ctx context.Context) SystemSsoforticloudadminMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemSsoforticloudadminMapOutput)
}

type SystemSsoforticloudadminOutput struct{ *pulumi.OutputState }

func (SystemSsoforticloudadminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemSsoforticloudadmin)(nil)).Elem()
}

func (o SystemSsoforticloudadminOutput) ToSystemSsoforticloudadminOutput() SystemSsoforticloudadminOutput {
	return o
}

func (o SystemSsoforticloudadminOutput) ToSystemSsoforticloudadminOutputWithContext(ctx context.Context) SystemSsoforticloudadminOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SystemSsoforticloudadminOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSsoforticloudadmin) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// FortiCloud SSO admin name.
func (o SystemSsoforticloudadminOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemSsoforticloudadmin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemSsoforticloudadminOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemSsoforticloudadmin) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Virtual domain(s) that the administrator can access. The structure of `vdom` block is documented below.
func (o SystemSsoforticloudadminOutput) Vdoms() SystemSsoforticloudadminVdomArrayOutput {
	return o.ApplyT(func(v *SystemSsoforticloudadmin) SystemSsoforticloudadminVdomArrayOutput { return v.Vdoms }).(SystemSsoforticloudadminVdomArrayOutput)
}

type SystemSsoforticloudadminArrayOutput struct{ *pulumi.OutputState }

func (SystemSsoforticloudadminArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemSsoforticloudadmin)(nil)).Elem()
}

func (o SystemSsoforticloudadminArrayOutput) ToSystemSsoforticloudadminArrayOutput() SystemSsoforticloudadminArrayOutput {
	return o
}

func (o SystemSsoforticloudadminArrayOutput) ToSystemSsoforticloudadminArrayOutputWithContext(ctx context.Context) SystemSsoforticloudadminArrayOutput {
	return o
}

func (o SystemSsoforticloudadminArrayOutput) Index(i pulumi.IntInput) SystemSsoforticloudadminOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemSsoforticloudadmin {
		return vs[0].([]*SystemSsoforticloudadmin)[vs[1].(int)]
	}).(SystemSsoforticloudadminOutput)
}

type SystemSsoforticloudadminMapOutput struct{ *pulumi.OutputState }

func (SystemSsoforticloudadminMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemSsoforticloudadmin)(nil)).Elem()
}

func (o SystemSsoforticloudadminMapOutput) ToSystemSsoforticloudadminMapOutput() SystemSsoforticloudadminMapOutput {
	return o
}

func (o SystemSsoforticloudadminMapOutput) ToSystemSsoforticloudadminMapOutputWithContext(ctx context.Context) SystemSsoforticloudadminMapOutput {
	return o
}

func (o SystemSsoforticloudadminMapOutput) MapIndex(k pulumi.StringInput) SystemSsoforticloudadminOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemSsoforticloudadmin {
		return vs[0].(map[string]*SystemSsoforticloudadmin)[vs[1].(string)]
	}).(SystemSsoforticloudadminOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSsoforticloudadminInput)(nil)).Elem(), &SystemSsoforticloudadmin{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSsoforticloudadminArrayInput)(nil)).Elem(), SystemSsoforticloudadminArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemSsoforticloudadminMapInput)(nil)).Elem(), SystemSsoforticloudadminMap{})
	pulumi.RegisterOutputType(SystemSsoforticloudadminOutput{})
	pulumi.RegisterOutputType(SystemSsoforticloudadminArrayOutput{})
	pulumi.RegisterOutputType(SystemSsoforticloudadminMapOutput{})
}