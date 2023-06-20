// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the client with its MAC address. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6`.
//
// ## Import
//
// # WirelessController Address can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerAddress struct {
	pulumi.CustomResourceState

	// ID.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// MAC address.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerAddress registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerAddress(ctx *pulumi.Context,
	name string, args *WirelesscontrollerAddressArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerAddress, error) {
	if args == nil {
		args = &WirelesscontrollerAddressArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerAddress
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerAddress gets an existing WirelesscontrollerAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerAddressState, opts ...pulumi.ResourceOption) (*WirelesscontrollerAddress, error) {
	var resource WirelesscontrollerAddress
	err := ctx.ReadResource("fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerAddress resources.
type wirelesscontrollerAddressState struct {
	// ID.
	Fosid *string `pulumi:"fosid"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy *string `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelesscontrollerAddressState struct {
	// ID.
	Fosid pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerAddressState)(nil)).Elem()
}

type wirelesscontrollerAddressArgs struct {
	// ID.
	Fosid *string `pulumi:"fosid"`
	// MAC address.
	Mac *string `pulumi:"mac"`
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy *string `pulumi:"policy"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelesscontrollerAddress resource.
type WirelesscontrollerAddressArgs struct {
	// ID.
	Fosid pulumi.StringPtrInput
	// MAC address.
	Mac pulumi.StringPtrInput
	// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
	Policy pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerAddressArgs)(nil)).Elem()
}

type WirelesscontrollerAddressInput interface {
	pulumi.Input

	ToWirelesscontrollerAddressOutput() WirelesscontrollerAddressOutput
	ToWirelesscontrollerAddressOutputWithContext(ctx context.Context) WirelesscontrollerAddressOutput
}

func (*WirelesscontrollerAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerAddress)(nil)).Elem()
}

func (i *WirelesscontrollerAddress) ToWirelesscontrollerAddressOutput() WirelesscontrollerAddressOutput {
	return i.ToWirelesscontrollerAddressOutputWithContext(context.Background())
}

func (i *WirelesscontrollerAddress) ToWirelesscontrollerAddressOutputWithContext(ctx context.Context) WirelesscontrollerAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerAddressOutput)
}

// WirelesscontrollerAddressArrayInput is an input type that accepts WirelesscontrollerAddressArray and WirelesscontrollerAddressArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerAddressArrayInput` via:
//
//	WirelesscontrollerAddressArray{ WirelesscontrollerAddressArgs{...} }
type WirelesscontrollerAddressArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerAddressArrayOutput() WirelesscontrollerAddressArrayOutput
	ToWirelesscontrollerAddressArrayOutputWithContext(context.Context) WirelesscontrollerAddressArrayOutput
}

type WirelesscontrollerAddressArray []WirelesscontrollerAddressInput

func (WirelesscontrollerAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerAddress)(nil)).Elem()
}

func (i WirelesscontrollerAddressArray) ToWirelesscontrollerAddressArrayOutput() WirelesscontrollerAddressArrayOutput {
	return i.ToWirelesscontrollerAddressArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerAddressArray) ToWirelesscontrollerAddressArrayOutputWithContext(ctx context.Context) WirelesscontrollerAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerAddressArrayOutput)
}

// WirelesscontrollerAddressMapInput is an input type that accepts WirelesscontrollerAddressMap and WirelesscontrollerAddressMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerAddressMapInput` via:
//
//	WirelesscontrollerAddressMap{ "key": WirelesscontrollerAddressArgs{...} }
type WirelesscontrollerAddressMapInput interface {
	pulumi.Input

	ToWirelesscontrollerAddressMapOutput() WirelesscontrollerAddressMapOutput
	ToWirelesscontrollerAddressMapOutputWithContext(context.Context) WirelesscontrollerAddressMapOutput
}

type WirelesscontrollerAddressMap map[string]WirelesscontrollerAddressInput

func (WirelesscontrollerAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerAddress)(nil)).Elem()
}

func (i WirelesscontrollerAddressMap) ToWirelesscontrollerAddressMapOutput() WirelesscontrollerAddressMapOutput {
	return i.ToWirelesscontrollerAddressMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerAddressMap) ToWirelesscontrollerAddressMapOutputWithContext(ctx context.Context) WirelesscontrollerAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerAddressMapOutput)
}

type WirelesscontrollerAddressOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerAddress)(nil)).Elem()
}

func (o WirelesscontrollerAddressOutput) ToWirelesscontrollerAddressOutput() WirelesscontrollerAddressOutput {
	return o
}

func (o WirelesscontrollerAddressOutput) ToWirelesscontrollerAddressOutputWithContext(ctx context.Context) WirelesscontrollerAddressOutput {
	return o
}

// ID.
func (o WirelesscontrollerAddressOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerAddress) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// MAC address.
func (o WirelesscontrollerAddressOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerAddress) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
func (o WirelesscontrollerAddressOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerAddress) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerAddressOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerAddress) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelesscontrollerAddressArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerAddress)(nil)).Elem()
}

func (o WirelesscontrollerAddressArrayOutput) ToWirelesscontrollerAddressArrayOutput() WirelesscontrollerAddressArrayOutput {
	return o
}

func (o WirelesscontrollerAddressArrayOutput) ToWirelesscontrollerAddressArrayOutputWithContext(ctx context.Context) WirelesscontrollerAddressArrayOutput {
	return o
}

func (o WirelesscontrollerAddressArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerAddress {
		return vs[0].([]*WirelesscontrollerAddress)[vs[1].(int)]
	}).(WirelesscontrollerAddressOutput)
}

type WirelesscontrollerAddressMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerAddress)(nil)).Elem()
}

func (o WirelesscontrollerAddressMapOutput) ToWirelesscontrollerAddressMapOutput() WirelesscontrollerAddressMapOutput {
	return o
}

func (o WirelesscontrollerAddressMapOutput) ToWirelesscontrollerAddressMapOutputWithContext(ctx context.Context) WirelesscontrollerAddressMapOutput {
	return o
}

func (o WirelesscontrollerAddressMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerAddress {
		return vs[0].(map[string]*WirelesscontrollerAddress)[vs[1].(string)]
	}).(WirelesscontrollerAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerAddressInput)(nil)).Elem(), &WirelesscontrollerAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerAddressArrayInput)(nil)).Elem(), WirelesscontrollerAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerAddressMapInput)(nil)).Elem(), WirelesscontrollerAddressMap{})
	pulumi.RegisterOutputType(WirelesscontrollerAddressOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerAddressArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerAddressMapOutput{})
}
