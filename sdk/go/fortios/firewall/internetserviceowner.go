// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Internet Service owner. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Firewall InternetServiceOwner can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/internetserviceowner:Internetserviceowner labelname {{fosid}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/internetserviceowner:Internetserviceowner labelname {{fosid}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Internetserviceowner struct {
	pulumi.CustomResourceState

	// Internet Service owner ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Internet Service owner name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewInternetserviceowner registers a new resource with the given unique name, arguments, and options.
func NewInternetserviceowner(ctx *pulumi.Context,
	name string, args *InternetserviceownerArgs, opts ...pulumi.ResourceOption) (*Internetserviceowner, error) {
	if args == nil {
		args = &InternetserviceownerArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Internetserviceowner
	err := ctx.RegisterResource("fortios:firewall/internetserviceowner:Internetserviceowner", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetserviceowner gets an existing Internetserviceowner resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetserviceowner(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetserviceownerState, opts ...pulumi.ResourceOption) (*Internetserviceowner, error) {
	var resource Internetserviceowner
	err := ctx.ReadResource("fortios:firewall/internetserviceowner:Internetserviceowner", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Internetserviceowner resources.
type internetserviceownerState struct {
	// Internet Service owner ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service owner name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type InternetserviceownerState struct {
	// Internet Service owner ID.
	Fosid pulumi.IntPtrInput
	// Internet Service owner name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceownerState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceownerState)(nil)).Elem()
}

type internetserviceownerArgs struct {
	// Internet Service owner ID.
	Fosid *int `pulumi:"fosid"`
	// Internet Service owner name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Internetserviceowner resource.
type InternetserviceownerArgs struct {
	// Internet Service owner ID.
	Fosid pulumi.IntPtrInput
	// Internet Service owner name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (InternetserviceownerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetserviceownerArgs)(nil)).Elem()
}

type InternetserviceownerInput interface {
	pulumi.Input

	ToInternetserviceownerOutput() InternetserviceownerOutput
	ToInternetserviceownerOutputWithContext(ctx context.Context) InternetserviceownerOutput
}

func (*Internetserviceowner) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetserviceowner)(nil)).Elem()
}

func (i *Internetserviceowner) ToInternetserviceownerOutput() InternetserviceownerOutput {
	return i.ToInternetserviceownerOutputWithContext(context.Background())
}

func (i *Internetserviceowner) ToInternetserviceownerOutputWithContext(ctx context.Context) InternetserviceownerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceownerOutput)
}

// InternetserviceownerArrayInput is an input type that accepts InternetserviceownerArray and InternetserviceownerArrayOutput values.
// You can construct a concrete instance of `InternetserviceownerArrayInput` via:
//
//	InternetserviceownerArray{ InternetserviceownerArgs{...} }
type InternetserviceownerArrayInput interface {
	pulumi.Input

	ToInternetserviceownerArrayOutput() InternetserviceownerArrayOutput
	ToInternetserviceownerArrayOutputWithContext(context.Context) InternetserviceownerArrayOutput
}

type InternetserviceownerArray []InternetserviceownerInput

func (InternetserviceownerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetserviceowner)(nil)).Elem()
}

func (i InternetserviceownerArray) ToInternetserviceownerArrayOutput() InternetserviceownerArrayOutput {
	return i.ToInternetserviceownerArrayOutputWithContext(context.Background())
}

func (i InternetserviceownerArray) ToInternetserviceownerArrayOutputWithContext(ctx context.Context) InternetserviceownerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceownerArrayOutput)
}

// InternetserviceownerMapInput is an input type that accepts InternetserviceownerMap and InternetserviceownerMapOutput values.
// You can construct a concrete instance of `InternetserviceownerMapInput` via:
//
//	InternetserviceownerMap{ "key": InternetserviceownerArgs{...} }
type InternetserviceownerMapInput interface {
	pulumi.Input

	ToInternetserviceownerMapOutput() InternetserviceownerMapOutput
	ToInternetserviceownerMapOutputWithContext(context.Context) InternetserviceownerMapOutput
}

type InternetserviceownerMap map[string]InternetserviceownerInput

func (InternetserviceownerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetserviceowner)(nil)).Elem()
}

func (i InternetserviceownerMap) ToInternetserviceownerMapOutput() InternetserviceownerMapOutput {
	return i.ToInternetserviceownerMapOutputWithContext(context.Background())
}

func (i InternetserviceownerMap) ToInternetserviceownerMapOutputWithContext(ctx context.Context) InternetserviceownerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetserviceownerMapOutput)
}

type InternetserviceownerOutput struct{ *pulumi.OutputState }

func (InternetserviceownerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Internetserviceowner)(nil)).Elem()
}

func (o InternetserviceownerOutput) ToInternetserviceownerOutput() InternetserviceownerOutput {
	return o
}

func (o InternetserviceownerOutput) ToInternetserviceownerOutputWithContext(ctx context.Context) InternetserviceownerOutput {
	return o
}

// Internet Service owner ID.
func (o InternetserviceownerOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *Internetserviceowner) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Internet Service owner name.
func (o InternetserviceownerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Internetserviceowner) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o InternetserviceownerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Internetserviceowner) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type InternetserviceownerArrayOutput struct{ *pulumi.OutputState }

func (InternetserviceownerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Internetserviceowner)(nil)).Elem()
}

func (o InternetserviceownerArrayOutput) ToInternetserviceownerArrayOutput() InternetserviceownerArrayOutput {
	return o
}

func (o InternetserviceownerArrayOutput) ToInternetserviceownerArrayOutputWithContext(ctx context.Context) InternetserviceownerArrayOutput {
	return o
}

func (o InternetserviceownerArrayOutput) Index(i pulumi.IntInput) InternetserviceownerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Internetserviceowner {
		return vs[0].([]*Internetserviceowner)[vs[1].(int)]
	}).(InternetserviceownerOutput)
}

type InternetserviceownerMapOutput struct{ *pulumi.OutputState }

func (InternetserviceownerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Internetserviceowner)(nil)).Elem()
}

func (o InternetserviceownerMapOutput) ToInternetserviceownerMapOutput() InternetserviceownerMapOutput {
	return o
}

func (o InternetserviceownerMapOutput) ToInternetserviceownerMapOutputWithContext(ctx context.Context) InternetserviceownerMapOutput {
	return o
}

func (o InternetserviceownerMapOutput) MapIndex(k pulumi.StringInput) InternetserviceownerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Internetserviceowner {
		return vs[0].(map[string]*Internetserviceowner)[vs[1].(string)]
	}).(InternetserviceownerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceownerInput)(nil)).Elem(), &Internetserviceowner{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceownerArrayInput)(nil)).Elem(), InternetserviceownerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InternetserviceownerMapInput)(nil)).Elem(), InternetserviceownerMap{})
	pulumi.RegisterOutputType(InternetserviceownerOutput{})
	pulumi.RegisterOutputType(InternetserviceownerArrayOutput{})
	pulumi.RegisterOutputType(InternetserviceownerMapOutput{})
}