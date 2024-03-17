// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure names for shaping classes. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// Firewall TrafficClass can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:firewall/trafficclass:Trafficclass labelname {{class_id}}
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Trafficclass struct {
	pulumi.CustomResourceState

	// Class ID to be named.
	ClassId pulumi.IntOutput `pulumi:"classId"`
	// Define the name for this class-id.
	ClassName pulumi.StringOutput `pulumi:"className"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewTrafficclass registers a new resource with the given unique name, arguments, and options.
func NewTrafficclass(ctx *pulumi.Context,
	name string, args *TrafficclassArgs, opts ...pulumi.ResourceOption) (*Trafficclass, error) {
	if args == nil {
		args = &TrafficclassArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Trafficclass
	err := ctx.RegisterResource("fortios:firewall/trafficclass:Trafficclass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficclass gets an existing Trafficclass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficclass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficclassState, opts ...pulumi.ResourceOption) (*Trafficclass, error) {
	var resource Trafficclass
	err := ctx.ReadResource("fortios:firewall/trafficclass:Trafficclass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trafficclass resources.
type trafficclassState struct {
	// Class ID to be named.
	ClassId *int `pulumi:"classId"`
	// Define the name for this class-id.
	ClassName *string `pulumi:"className"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type TrafficclassState struct {
	// Class ID to be named.
	ClassId pulumi.IntPtrInput
	// Define the name for this class-id.
	ClassName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TrafficclassState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficclassState)(nil)).Elem()
}

type trafficclassArgs struct {
	// Class ID to be named.
	ClassId *int `pulumi:"classId"`
	// Define the name for this class-id.
	ClassName *string `pulumi:"className"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Trafficclass resource.
type TrafficclassArgs struct {
	// Class ID to be named.
	ClassId pulumi.IntPtrInput
	// Define the name for this class-id.
	ClassName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (TrafficclassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficclassArgs)(nil)).Elem()
}

type TrafficclassInput interface {
	pulumi.Input

	ToTrafficclassOutput() TrafficclassOutput
	ToTrafficclassOutputWithContext(ctx context.Context) TrafficclassOutput
}

func (*Trafficclass) ElementType() reflect.Type {
	return reflect.TypeOf((**Trafficclass)(nil)).Elem()
}

func (i *Trafficclass) ToTrafficclassOutput() TrafficclassOutput {
	return i.ToTrafficclassOutputWithContext(context.Background())
}

func (i *Trafficclass) ToTrafficclassOutputWithContext(ctx context.Context) TrafficclassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficclassOutput)
}

// TrafficclassArrayInput is an input type that accepts TrafficclassArray and TrafficclassArrayOutput values.
// You can construct a concrete instance of `TrafficclassArrayInput` via:
//
//	TrafficclassArray{ TrafficclassArgs{...} }
type TrafficclassArrayInput interface {
	pulumi.Input

	ToTrafficclassArrayOutput() TrafficclassArrayOutput
	ToTrafficclassArrayOutputWithContext(context.Context) TrafficclassArrayOutput
}

type TrafficclassArray []TrafficclassInput

func (TrafficclassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trafficclass)(nil)).Elem()
}

func (i TrafficclassArray) ToTrafficclassArrayOutput() TrafficclassArrayOutput {
	return i.ToTrafficclassArrayOutputWithContext(context.Background())
}

func (i TrafficclassArray) ToTrafficclassArrayOutputWithContext(ctx context.Context) TrafficclassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficclassArrayOutput)
}

// TrafficclassMapInput is an input type that accepts TrafficclassMap and TrafficclassMapOutput values.
// You can construct a concrete instance of `TrafficclassMapInput` via:
//
//	TrafficclassMap{ "key": TrafficclassArgs{...} }
type TrafficclassMapInput interface {
	pulumi.Input

	ToTrafficclassMapOutput() TrafficclassMapOutput
	ToTrafficclassMapOutputWithContext(context.Context) TrafficclassMapOutput
}

type TrafficclassMap map[string]TrafficclassInput

func (TrafficclassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trafficclass)(nil)).Elem()
}

func (i TrafficclassMap) ToTrafficclassMapOutput() TrafficclassMapOutput {
	return i.ToTrafficclassMapOutputWithContext(context.Background())
}

func (i TrafficclassMap) ToTrafficclassMapOutputWithContext(ctx context.Context) TrafficclassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficclassMapOutput)
}

type TrafficclassOutput struct{ *pulumi.OutputState }

func (TrafficclassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trafficclass)(nil)).Elem()
}

func (o TrafficclassOutput) ToTrafficclassOutput() TrafficclassOutput {
	return o
}

func (o TrafficclassOutput) ToTrafficclassOutputWithContext(ctx context.Context) TrafficclassOutput {
	return o
}

// Class ID to be named.
func (o TrafficclassOutput) ClassId() pulumi.IntOutput {
	return o.ApplyT(func(v *Trafficclass) pulumi.IntOutput { return v.ClassId }).(pulumi.IntOutput)
}

// Define the name for this class-id.
func (o TrafficclassOutput) ClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *Trafficclass) pulumi.StringOutput { return v.ClassName }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o TrafficclassOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trafficclass) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type TrafficclassArrayOutput struct{ *pulumi.OutputState }

func (TrafficclassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trafficclass)(nil)).Elem()
}

func (o TrafficclassArrayOutput) ToTrafficclassArrayOutput() TrafficclassArrayOutput {
	return o
}

func (o TrafficclassArrayOutput) ToTrafficclassArrayOutputWithContext(ctx context.Context) TrafficclassArrayOutput {
	return o
}

func (o TrafficclassArrayOutput) Index(i pulumi.IntInput) TrafficclassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trafficclass {
		return vs[0].([]*Trafficclass)[vs[1].(int)]
	}).(TrafficclassOutput)
}

type TrafficclassMapOutput struct{ *pulumi.OutputState }

func (TrafficclassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trafficclass)(nil)).Elem()
}

func (o TrafficclassMapOutput) ToTrafficclassMapOutput() TrafficclassMapOutput {
	return o
}

func (o TrafficclassMapOutput) ToTrafficclassMapOutputWithContext(ctx context.Context) TrafficclassMapOutput {
	return o
}

func (o TrafficclassMapOutput) MapIndex(k pulumi.StringInput) TrafficclassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trafficclass {
		return vs[0].(map[string]*Trafficclass)[vs[1].(string)]
	}).(TrafficclassOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficclassInput)(nil)).Elem(), &Trafficclass{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficclassArrayInput)(nil)).Elem(), TrafficclassArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficclassMapInput)(nil)).Elem(), TrafficclassMap{})
	pulumi.RegisterOutputType(TrafficclassOutput{})
	pulumi.RegisterOutputType(TrafficclassArrayOutput{})
	pulumi.RegisterOutputType(TrafficclassMapOutput{})
}