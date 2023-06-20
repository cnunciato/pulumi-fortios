// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP blacklist vendor. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Firewall InternetServiceIpblVendor can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetserviceipblvendor:FirewallInternetserviceipblvendor labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallInternetserviceipblvendor:FirewallInternetserviceipblvendor labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallInternetserviceipblvendor struct {
	pulumi.CustomResourceState

	// IP blacklist vendor ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// IP blacklist vendor name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallInternetserviceipblvendor registers a new resource with the given unique name, arguments, and options.
func NewFirewallInternetserviceipblvendor(ctx *pulumi.Context,
	name string, args *FirewallInternetserviceipblvendorArgs, opts ...pulumi.ResourceOption) (*FirewallInternetserviceipblvendor, error) {
	if args == nil {
		args = &FirewallInternetserviceipblvendorArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallInternetserviceipblvendor
	err := ctx.RegisterResource("fortios:index/firewallInternetserviceipblvendor:FirewallInternetserviceipblvendor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallInternetserviceipblvendor gets an existing FirewallInternetserviceipblvendor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallInternetserviceipblvendor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallInternetserviceipblvendorState, opts ...pulumi.ResourceOption) (*FirewallInternetserviceipblvendor, error) {
	var resource FirewallInternetserviceipblvendor
	err := ctx.ReadResource("fortios:index/firewallInternetserviceipblvendor:FirewallInternetserviceipblvendor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallInternetserviceipblvendor resources.
type firewallInternetserviceipblvendorState struct {
	// IP blacklist vendor ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist vendor name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallInternetserviceipblvendorState struct {
	// IP blacklist vendor ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist vendor name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetserviceipblvendorState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetserviceipblvendorState)(nil)).Elem()
}

type firewallInternetserviceipblvendorArgs struct {
	// IP blacklist vendor ID.
	Fosid *int `pulumi:"fosid"`
	// IP blacklist vendor name.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallInternetserviceipblvendor resource.
type FirewallInternetserviceipblvendorArgs struct {
	// IP blacklist vendor ID.
	Fosid pulumi.IntPtrInput
	// IP blacklist vendor name.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallInternetserviceipblvendorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallInternetserviceipblvendorArgs)(nil)).Elem()
}

type FirewallInternetserviceipblvendorInput interface {
	pulumi.Input

	ToFirewallInternetserviceipblvendorOutput() FirewallInternetserviceipblvendorOutput
	ToFirewallInternetserviceipblvendorOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorOutput
}

func (*FirewallInternetserviceipblvendor) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetserviceipblvendor)(nil)).Elem()
}

func (i *FirewallInternetserviceipblvendor) ToFirewallInternetserviceipblvendorOutput() FirewallInternetserviceipblvendorOutput {
	return i.ToFirewallInternetserviceipblvendorOutputWithContext(context.Background())
}

func (i *FirewallInternetserviceipblvendor) ToFirewallInternetserviceipblvendorOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetserviceipblvendorOutput)
}

// FirewallInternetserviceipblvendorArrayInput is an input type that accepts FirewallInternetserviceipblvendorArray and FirewallInternetserviceipblvendorArrayOutput values.
// You can construct a concrete instance of `FirewallInternetserviceipblvendorArrayInput` via:
//
//	FirewallInternetserviceipblvendorArray{ FirewallInternetserviceipblvendorArgs{...} }
type FirewallInternetserviceipblvendorArrayInput interface {
	pulumi.Input

	ToFirewallInternetserviceipblvendorArrayOutput() FirewallInternetserviceipblvendorArrayOutput
	ToFirewallInternetserviceipblvendorArrayOutputWithContext(context.Context) FirewallInternetserviceipblvendorArrayOutput
}

type FirewallInternetserviceipblvendorArray []FirewallInternetserviceipblvendorInput

func (FirewallInternetserviceipblvendorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetserviceipblvendor)(nil)).Elem()
}

func (i FirewallInternetserviceipblvendorArray) ToFirewallInternetserviceipblvendorArrayOutput() FirewallInternetserviceipblvendorArrayOutput {
	return i.ToFirewallInternetserviceipblvendorArrayOutputWithContext(context.Background())
}

func (i FirewallInternetserviceipblvendorArray) ToFirewallInternetserviceipblvendorArrayOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetserviceipblvendorArrayOutput)
}

// FirewallInternetserviceipblvendorMapInput is an input type that accepts FirewallInternetserviceipblvendorMap and FirewallInternetserviceipblvendorMapOutput values.
// You can construct a concrete instance of `FirewallInternetserviceipblvendorMapInput` via:
//
//	FirewallInternetserviceipblvendorMap{ "key": FirewallInternetserviceipblvendorArgs{...} }
type FirewallInternetserviceipblvendorMapInput interface {
	pulumi.Input

	ToFirewallInternetserviceipblvendorMapOutput() FirewallInternetserviceipblvendorMapOutput
	ToFirewallInternetserviceipblvendorMapOutputWithContext(context.Context) FirewallInternetserviceipblvendorMapOutput
}

type FirewallInternetserviceipblvendorMap map[string]FirewallInternetserviceipblvendorInput

func (FirewallInternetserviceipblvendorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetserviceipblvendor)(nil)).Elem()
}

func (i FirewallInternetserviceipblvendorMap) ToFirewallInternetserviceipblvendorMapOutput() FirewallInternetserviceipblvendorMapOutput {
	return i.ToFirewallInternetserviceipblvendorMapOutputWithContext(context.Background())
}

func (i FirewallInternetserviceipblvendorMap) ToFirewallInternetserviceipblvendorMapOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallInternetserviceipblvendorMapOutput)
}

type FirewallInternetserviceipblvendorOutput struct{ *pulumi.OutputState }

func (FirewallInternetserviceipblvendorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallInternetserviceipblvendor)(nil)).Elem()
}

func (o FirewallInternetserviceipblvendorOutput) ToFirewallInternetserviceipblvendorOutput() FirewallInternetserviceipblvendorOutput {
	return o
}

func (o FirewallInternetserviceipblvendorOutput) ToFirewallInternetserviceipblvendorOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorOutput {
	return o
}

// IP blacklist vendor ID.
func (o FirewallInternetserviceipblvendorOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallInternetserviceipblvendor) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// IP blacklist vendor name.
func (o FirewallInternetserviceipblvendorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallInternetserviceipblvendor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallInternetserviceipblvendorOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallInternetserviceipblvendor) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallInternetserviceipblvendorArrayOutput struct{ *pulumi.OutputState }

func (FirewallInternetserviceipblvendorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallInternetserviceipblvendor)(nil)).Elem()
}

func (o FirewallInternetserviceipblvendorArrayOutput) ToFirewallInternetserviceipblvendorArrayOutput() FirewallInternetserviceipblvendorArrayOutput {
	return o
}

func (o FirewallInternetserviceipblvendorArrayOutput) ToFirewallInternetserviceipblvendorArrayOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorArrayOutput {
	return o
}

func (o FirewallInternetserviceipblvendorArrayOutput) Index(i pulumi.IntInput) FirewallInternetserviceipblvendorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallInternetserviceipblvendor {
		return vs[0].([]*FirewallInternetserviceipblvendor)[vs[1].(int)]
	}).(FirewallInternetserviceipblvendorOutput)
}

type FirewallInternetserviceipblvendorMapOutput struct{ *pulumi.OutputState }

func (FirewallInternetserviceipblvendorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallInternetserviceipblvendor)(nil)).Elem()
}

func (o FirewallInternetserviceipblvendorMapOutput) ToFirewallInternetserviceipblvendorMapOutput() FirewallInternetserviceipblvendorMapOutput {
	return o
}

func (o FirewallInternetserviceipblvendorMapOutput) ToFirewallInternetserviceipblvendorMapOutputWithContext(ctx context.Context) FirewallInternetserviceipblvendorMapOutput {
	return o
}

func (o FirewallInternetserviceipblvendorMapOutput) MapIndex(k pulumi.StringInput) FirewallInternetserviceipblvendorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallInternetserviceipblvendor {
		return vs[0].(map[string]*FirewallInternetserviceipblvendor)[vs[1].(string)]
	}).(FirewallInternetserviceipblvendorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetserviceipblvendorInput)(nil)).Elem(), &FirewallInternetserviceipblvendor{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetserviceipblvendorArrayInput)(nil)).Elem(), FirewallInternetserviceipblvendorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInternetserviceipblvendorMapInput)(nil)).Elem(), FirewallInternetserviceipblvendorMap{})
	pulumi.RegisterOutputType(FirewallInternetserviceipblvendorOutput{})
	pulumi.RegisterOutputType(FirewallInternetserviceipblvendorArrayOutput{})
	pulumi.RegisterOutputType(FirewallInternetserviceipblvendorMapOutput{})
}
