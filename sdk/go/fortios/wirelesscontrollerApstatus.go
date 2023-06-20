// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure access point status (rogue | accepted | suppressed).
//
// ## Import
//
// # WirelessController ApStatus can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerApstatus struct {
	pulumi.CustomResourceState

	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringOutput `pulumi:"bssid"`
	// AP ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringOutput `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerApstatus registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerApstatus(ctx *pulumi.Context,
	name string, args *WirelesscontrollerApstatusArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerApstatus, error) {
	if args == nil {
		args = &WirelesscontrollerApstatusArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerApstatus
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerApstatus gets an existing WirelesscontrollerApstatus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerApstatus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerApstatusState, opts ...pulumi.ResourceOption) (*WirelesscontrollerApstatus, error) {
	var resource WirelesscontrollerApstatus
	err := ctx.ReadResource("fortios:index/wirelesscontrollerApstatus:WirelesscontrollerApstatus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerApstatus resources.
type wirelesscontrollerApstatusState struct {
	// Access Point's (AP's) BSSID.
	Bssid *string `pulumi:"bssid"`
	// AP ID.
	Fosid *int `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid *string `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelesscontrollerApstatusState struct {
	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringPtrInput
	// AP ID.
	Fosid pulumi.IntPtrInput
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringPtrInput
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerApstatusState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerApstatusState)(nil)).Elem()
}

type wirelesscontrollerApstatusArgs struct {
	// Access Point's (AP's) BSSID.
	Bssid *string `pulumi:"bssid"`
	// AP ID.
	Fosid *int `pulumi:"fosid"`
	// Access Point's (AP's) SSID.
	Ssid *string `pulumi:"ssid"`
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelesscontrollerApstatus resource.
type WirelesscontrollerApstatusArgs struct {
	// Access Point's (AP's) BSSID.
	Bssid pulumi.StringPtrInput
	// AP ID.
	Fosid pulumi.IntPtrInput
	// Access Point's (AP's) SSID.
	Ssid pulumi.StringPtrInput
	// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerApstatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerApstatusArgs)(nil)).Elem()
}

type WirelesscontrollerApstatusInput interface {
	pulumi.Input

	ToWirelesscontrollerApstatusOutput() WirelesscontrollerApstatusOutput
	ToWirelesscontrollerApstatusOutputWithContext(ctx context.Context) WirelesscontrollerApstatusOutput
}

func (*WirelesscontrollerApstatus) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerApstatus)(nil)).Elem()
}

func (i *WirelesscontrollerApstatus) ToWirelesscontrollerApstatusOutput() WirelesscontrollerApstatusOutput {
	return i.ToWirelesscontrollerApstatusOutputWithContext(context.Background())
}

func (i *WirelesscontrollerApstatus) ToWirelesscontrollerApstatusOutputWithContext(ctx context.Context) WirelesscontrollerApstatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerApstatusOutput)
}

// WirelesscontrollerApstatusArrayInput is an input type that accepts WirelesscontrollerApstatusArray and WirelesscontrollerApstatusArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerApstatusArrayInput` via:
//
//	WirelesscontrollerApstatusArray{ WirelesscontrollerApstatusArgs{...} }
type WirelesscontrollerApstatusArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerApstatusArrayOutput() WirelesscontrollerApstatusArrayOutput
	ToWirelesscontrollerApstatusArrayOutputWithContext(context.Context) WirelesscontrollerApstatusArrayOutput
}

type WirelesscontrollerApstatusArray []WirelesscontrollerApstatusInput

func (WirelesscontrollerApstatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerApstatus)(nil)).Elem()
}

func (i WirelesscontrollerApstatusArray) ToWirelesscontrollerApstatusArrayOutput() WirelesscontrollerApstatusArrayOutput {
	return i.ToWirelesscontrollerApstatusArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerApstatusArray) ToWirelesscontrollerApstatusArrayOutputWithContext(ctx context.Context) WirelesscontrollerApstatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerApstatusArrayOutput)
}

// WirelesscontrollerApstatusMapInput is an input type that accepts WirelesscontrollerApstatusMap and WirelesscontrollerApstatusMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerApstatusMapInput` via:
//
//	WirelesscontrollerApstatusMap{ "key": WirelesscontrollerApstatusArgs{...} }
type WirelesscontrollerApstatusMapInput interface {
	pulumi.Input

	ToWirelesscontrollerApstatusMapOutput() WirelesscontrollerApstatusMapOutput
	ToWirelesscontrollerApstatusMapOutputWithContext(context.Context) WirelesscontrollerApstatusMapOutput
}

type WirelesscontrollerApstatusMap map[string]WirelesscontrollerApstatusInput

func (WirelesscontrollerApstatusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerApstatus)(nil)).Elem()
}

func (i WirelesscontrollerApstatusMap) ToWirelesscontrollerApstatusMapOutput() WirelesscontrollerApstatusMapOutput {
	return i.ToWirelesscontrollerApstatusMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerApstatusMap) ToWirelesscontrollerApstatusMapOutputWithContext(ctx context.Context) WirelesscontrollerApstatusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerApstatusMapOutput)
}

type WirelesscontrollerApstatusOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerApstatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerApstatus)(nil)).Elem()
}

func (o WirelesscontrollerApstatusOutput) ToWirelesscontrollerApstatusOutput() WirelesscontrollerApstatusOutput {
	return o
}

func (o WirelesscontrollerApstatusOutput) ToWirelesscontrollerApstatusOutputWithContext(ctx context.Context) WirelesscontrollerApstatusOutput {
	return o
}

// Access Point's (AP's) BSSID.
func (o WirelesscontrollerApstatusOutput) Bssid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerApstatus) pulumi.StringOutput { return v.Bssid }).(pulumi.StringOutput)
}

// AP ID.
func (o WirelesscontrollerApstatusOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerApstatus) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Access Point's (AP's) SSID.
func (o WirelesscontrollerApstatusOutput) Ssid() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerApstatus) pulumi.StringOutput { return v.Ssid }).(pulumi.StringOutput)
}

// Access Point's (AP's) status: rogue, accepted, or supressed. Valid values: `rogue`, `accepted`, `suppressed`.
func (o WirelesscontrollerApstatusOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerApstatus) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerApstatusOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerApstatus) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelesscontrollerApstatusArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerApstatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerApstatus)(nil)).Elem()
}

func (o WirelesscontrollerApstatusArrayOutput) ToWirelesscontrollerApstatusArrayOutput() WirelesscontrollerApstatusArrayOutput {
	return o
}

func (o WirelesscontrollerApstatusArrayOutput) ToWirelesscontrollerApstatusArrayOutputWithContext(ctx context.Context) WirelesscontrollerApstatusArrayOutput {
	return o
}

func (o WirelesscontrollerApstatusArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerApstatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerApstatus {
		return vs[0].([]*WirelesscontrollerApstatus)[vs[1].(int)]
	}).(WirelesscontrollerApstatusOutput)
}

type WirelesscontrollerApstatusMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerApstatusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerApstatus)(nil)).Elem()
}

func (o WirelesscontrollerApstatusMapOutput) ToWirelesscontrollerApstatusMapOutput() WirelesscontrollerApstatusMapOutput {
	return o
}

func (o WirelesscontrollerApstatusMapOutput) ToWirelesscontrollerApstatusMapOutputWithContext(ctx context.Context) WirelesscontrollerApstatusMapOutput {
	return o
}

func (o WirelesscontrollerApstatusMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerApstatusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerApstatus {
		return vs[0].(map[string]*WirelesscontrollerApstatus)[vs[1].(string)]
	}).(WirelesscontrollerApstatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerApstatusInput)(nil)).Elem(), &WirelesscontrollerApstatus{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerApstatusArrayInput)(nil)).Elem(), WirelesscontrollerApstatusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerApstatusMapInput)(nil)).Elem(), WirelesscontrollerApstatusMap{})
	pulumi.RegisterOutputType(WirelesscontrollerApstatusOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerApstatusArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerApstatusMapOutput{})
}
