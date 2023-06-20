// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure names for shaping classes. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Firewall TrafficClass can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallTrafficclass:FirewallTrafficclass labelname {{class_id}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallTrafficclass:FirewallTrafficclass labelname {{class_id}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallTrafficclass struct {
	pulumi.CustomResourceState

	// Class ID to be named.
	ClassId pulumi.IntOutput `pulumi:"classId"`
	// Define the name for this class-id.
	ClassName pulumi.StringOutput `pulumi:"className"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallTrafficclass registers a new resource with the given unique name, arguments, and options.
func NewFirewallTrafficclass(ctx *pulumi.Context,
	name string, args *FirewallTrafficclassArgs, opts ...pulumi.ResourceOption) (*FirewallTrafficclass, error) {
	if args == nil {
		args = &FirewallTrafficclassArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallTrafficclass
	err := ctx.RegisterResource("fortios:index/firewallTrafficclass:FirewallTrafficclass", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallTrafficclass gets an existing FirewallTrafficclass resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallTrafficclass(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallTrafficclassState, opts ...pulumi.ResourceOption) (*FirewallTrafficclass, error) {
	var resource FirewallTrafficclass
	err := ctx.ReadResource("fortios:index/firewallTrafficclass:FirewallTrafficclass", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallTrafficclass resources.
type firewallTrafficclassState struct {
	// Class ID to be named.
	ClassId *int `pulumi:"classId"`
	// Define the name for this class-id.
	ClassName *string `pulumi:"className"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallTrafficclassState struct {
	// Class ID to be named.
	ClassId pulumi.IntPtrInput
	// Define the name for this class-id.
	ClassName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallTrafficclassState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallTrafficclassState)(nil)).Elem()
}

type firewallTrafficclassArgs struct {
	// Class ID to be named.
	ClassId *int `pulumi:"classId"`
	// Define the name for this class-id.
	ClassName *string `pulumi:"className"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallTrafficclass resource.
type FirewallTrafficclassArgs struct {
	// Class ID to be named.
	ClassId pulumi.IntPtrInput
	// Define the name for this class-id.
	ClassName pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallTrafficclassArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallTrafficclassArgs)(nil)).Elem()
}

type FirewallTrafficclassInput interface {
	pulumi.Input

	ToFirewallTrafficclassOutput() FirewallTrafficclassOutput
	ToFirewallTrafficclassOutputWithContext(ctx context.Context) FirewallTrafficclassOutput
}

func (*FirewallTrafficclass) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallTrafficclass)(nil)).Elem()
}

func (i *FirewallTrafficclass) ToFirewallTrafficclassOutput() FirewallTrafficclassOutput {
	return i.ToFirewallTrafficclassOutputWithContext(context.Background())
}

func (i *FirewallTrafficclass) ToFirewallTrafficclassOutputWithContext(ctx context.Context) FirewallTrafficclassOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallTrafficclassOutput)
}

// FirewallTrafficclassArrayInput is an input type that accepts FirewallTrafficclassArray and FirewallTrafficclassArrayOutput values.
// You can construct a concrete instance of `FirewallTrafficclassArrayInput` via:
//
//	FirewallTrafficclassArray{ FirewallTrafficclassArgs{...} }
type FirewallTrafficclassArrayInput interface {
	pulumi.Input

	ToFirewallTrafficclassArrayOutput() FirewallTrafficclassArrayOutput
	ToFirewallTrafficclassArrayOutputWithContext(context.Context) FirewallTrafficclassArrayOutput
}

type FirewallTrafficclassArray []FirewallTrafficclassInput

func (FirewallTrafficclassArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallTrafficclass)(nil)).Elem()
}

func (i FirewallTrafficclassArray) ToFirewallTrafficclassArrayOutput() FirewallTrafficclassArrayOutput {
	return i.ToFirewallTrafficclassArrayOutputWithContext(context.Background())
}

func (i FirewallTrafficclassArray) ToFirewallTrafficclassArrayOutputWithContext(ctx context.Context) FirewallTrafficclassArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallTrafficclassArrayOutput)
}

// FirewallTrafficclassMapInput is an input type that accepts FirewallTrafficclassMap and FirewallTrafficclassMapOutput values.
// You can construct a concrete instance of `FirewallTrafficclassMapInput` via:
//
//	FirewallTrafficclassMap{ "key": FirewallTrafficclassArgs{...} }
type FirewallTrafficclassMapInput interface {
	pulumi.Input

	ToFirewallTrafficclassMapOutput() FirewallTrafficclassMapOutput
	ToFirewallTrafficclassMapOutputWithContext(context.Context) FirewallTrafficclassMapOutput
}

type FirewallTrafficclassMap map[string]FirewallTrafficclassInput

func (FirewallTrafficclassMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallTrafficclass)(nil)).Elem()
}

func (i FirewallTrafficclassMap) ToFirewallTrafficclassMapOutput() FirewallTrafficclassMapOutput {
	return i.ToFirewallTrafficclassMapOutputWithContext(context.Background())
}

func (i FirewallTrafficclassMap) ToFirewallTrafficclassMapOutputWithContext(ctx context.Context) FirewallTrafficclassMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallTrafficclassMapOutput)
}

type FirewallTrafficclassOutput struct{ *pulumi.OutputState }

func (FirewallTrafficclassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallTrafficclass)(nil)).Elem()
}

func (o FirewallTrafficclassOutput) ToFirewallTrafficclassOutput() FirewallTrafficclassOutput {
	return o
}

func (o FirewallTrafficclassOutput) ToFirewallTrafficclassOutputWithContext(ctx context.Context) FirewallTrafficclassOutput {
	return o
}

// Class ID to be named.
func (o FirewallTrafficclassOutput) ClassId() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallTrafficclass) pulumi.IntOutput { return v.ClassId }).(pulumi.IntOutput)
}

// Define the name for this class-id.
func (o FirewallTrafficclassOutput) ClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallTrafficclass) pulumi.StringOutput { return v.ClassName }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallTrafficclassOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallTrafficclass) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallTrafficclassArrayOutput struct{ *pulumi.OutputState }

func (FirewallTrafficclassArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallTrafficclass)(nil)).Elem()
}

func (o FirewallTrafficclassArrayOutput) ToFirewallTrafficclassArrayOutput() FirewallTrafficclassArrayOutput {
	return o
}

func (o FirewallTrafficclassArrayOutput) ToFirewallTrafficclassArrayOutputWithContext(ctx context.Context) FirewallTrafficclassArrayOutput {
	return o
}

func (o FirewallTrafficclassArrayOutput) Index(i pulumi.IntInput) FirewallTrafficclassOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallTrafficclass {
		return vs[0].([]*FirewallTrafficclass)[vs[1].(int)]
	}).(FirewallTrafficclassOutput)
}

type FirewallTrafficclassMapOutput struct{ *pulumi.OutputState }

func (FirewallTrafficclassMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallTrafficclass)(nil)).Elem()
}

func (o FirewallTrafficclassMapOutput) ToFirewallTrafficclassMapOutput() FirewallTrafficclassMapOutput {
	return o
}

func (o FirewallTrafficclassMapOutput) ToFirewallTrafficclassMapOutputWithContext(ctx context.Context) FirewallTrafficclassMapOutput {
	return o
}

func (o FirewallTrafficclassMapOutput) MapIndex(k pulumi.StringInput) FirewallTrafficclassOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallTrafficclass {
		return vs[0].(map[string]*FirewallTrafficclass)[vs[1].(string)]
	}).(FirewallTrafficclassOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallTrafficclassInput)(nil)).Elem(), &FirewallTrafficclass{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallTrafficclassArrayInput)(nil)).Elem(), FirewallTrafficclassArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallTrafficclassMapInput)(nil)).Elem(), FirewallTrafficclassMap{})
	pulumi.RegisterOutputType(FirewallTrafficclassOutput{})
	pulumi.RegisterOutputType(FirewallTrafficclassArrayOutput{})
	pulumi.RegisterOutputType(FirewallTrafficclassMapOutput{})
}
