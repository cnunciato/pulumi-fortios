// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure inter wireless controller operation.
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
//			_, err := fortios.NewWirelesscontrollerIntercontroller(ctx, "trname", &fortios.WirelesscontrollerIntercontrollerArgs{
//				FastFailoverMax:     pulumi.Int(10),
//				FastFailoverWait:    pulumi.Int(10),
//				InterControllerKey:  pulumi.String("ENC XXXX"),
//				InterControllerMode: pulumi.String("disable"),
//				InterControllerPri:  pulumi.String("primary"),
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
// # WirelessController InterController can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerIntercontroller:WirelesscontrollerIntercontroller labelname WirelessControllerInterController
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerIntercontroller:WirelesscontrollerIntercontroller labelname WirelessControllerInterController
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WirelesscontrollerIntercontroller struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax pulumi.IntOutput `pulumi:"fastFailoverMax"`
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait pulumi.IntOutput `pulumi:"fastFailoverWait"`
	// Secret key for inter-controller communications.
	InterControllerKey pulumi.StringPtrOutput `pulumi:"interControllerKey"`
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode pulumi.StringOutput `pulumi:"interControllerMode"`
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers WirelesscontrollerIntercontrollerInterControllerPeerArrayOutput `pulumi:"interControllerPeers"`
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri pulumi.StringOutput `pulumi:"interControllerPri"`
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming pulumi.StringOutput `pulumi:"l3Roaming"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerIntercontroller registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerIntercontroller(ctx *pulumi.Context,
	name string, args *WirelesscontrollerIntercontrollerArgs, opts ...pulumi.ResourceOption) (*WirelesscontrollerIntercontroller, error) {
	if args == nil {
		args = &WirelesscontrollerIntercontrollerArgs{}
	}

	if args.InterControllerKey != nil {
		args.InterControllerKey = pulumi.ToSecret(args.InterControllerKey).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"interControllerKey",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource WirelesscontrollerIntercontroller
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerIntercontroller:WirelesscontrollerIntercontroller", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerIntercontroller gets an existing WirelesscontrollerIntercontroller resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerIntercontroller(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WirelesscontrollerIntercontrollerState, opts ...pulumi.ResourceOption) (*WirelesscontrollerIntercontroller, error) {
	var resource WirelesscontrollerIntercontroller
	err := ctx.ReadResource("fortios:index/wirelesscontrollerIntercontroller:WirelesscontrollerIntercontroller", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WirelesscontrollerIntercontroller resources.
type wirelesscontrollerIntercontrollerState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax *int `pulumi:"fastFailoverMax"`
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait *int `pulumi:"fastFailoverWait"`
	// Secret key for inter-controller communications.
	InterControllerKey *string `pulumi:"interControllerKey"`
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode *string `pulumi:"interControllerMode"`
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers []WirelesscontrollerIntercontrollerInterControllerPeer `pulumi:"interControllerPeers"`
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri *string `pulumi:"interControllerPri"`
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming *string `pulumi:"l3Roaming"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WirelesscontrollerIntercontrollerState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax pulumi.IntPtrInput
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait pulumi.IntPtrInput
	// Secret key for inter-controller communications.
	InterControllerKey pulumi.StringPtrInput
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode pulumi.StringPtrInput
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers WirelesscontrollerIntercontrollerInterControllerPeerArrayInput
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri pulumi.StringPtrInput
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerIntercontrollerState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerIntercontrollerState)(nil)).Elem()
}

type wirelesscontrollerIntercontrollerArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax *int `pulumi:"fastFailoverMax"`
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait *int `pulumi:"fastFailoverWait"`
	// Secret key for inter-controller communications.
	InterControllerKey *string `pulumi:"interControllerKey"`
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode *string `pulumi:"interControllerMode"`
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers []WirelesscontrollerIntercontrollerInterControllerPeer `pulumi:"interControllerPeers"`
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri *string `pulumi:"interControllerPri"`
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming *string `pulumi:"l3Roaming"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WirelesscontrollerIntercontroller resource.
type WirelesscontrollerIntercontrollerArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
	FastFailoverMax pulumi.IntPtrInput
	// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
	FastFailoverWait pulumi.IntPtrInput
	// Secret key for inter-controller communications.
	InterControllerKey pulumi.StringPtrInput
	// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
	InterControllerMode pulumi.StringPtrInput
	// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
	InterControllerPeers WirelesscontrollerIntercontrollerInterControllerPeerArrayInput
	// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
	InterControllerPri pulumi.StringPtrInput
	// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
	L3Roaming pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WirelesscontrollerIntercontrollerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerIntercontrollerArgs)(nil)).Elem()
}

type WirelesscontrollerIntercontrollerInput interface {
	pulumi.Input

	ToWirelesscontrollerIntercontrollerOutput() WirelesscontrollerIntercontrollerOutput
	ToWirelesscontrollerIntercontrollerOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerOutput
}

func (*WirelesscontrollerIntercontroller) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerIntercontroller)(nil)).Elem()
}

func (i *WirelesscontrollerIntercontroller) ToWirelesscontrollerIntercontrollerOutput() WirelesscontrollerIntercontrollerOutput {
	return i.ToWirelesscontrollerIntercontrollerOutputWithContext(context.Background())
}

func (i *WirelesscontrollerIntercontroller) ToWirelesscontrollerIntercontrollerOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerIntercontrollerOutput)
}

// WirelesscontrollerIntercontrollerArrayInput is an input type that accepts WirelesscontrollerIntercontrollerArray and WirelesscontrollerIntercontrollerArrayOutput values.
// You can construct a concrete instance of `WirelesscontrollerIntercontrollerArrayInput` via:
//
//	WirelesscontrollerIntercontrollerArray{ WirelesscontrollerIntercontrollerArgs{...} }
type WirelesscontrollerIntercontrollerArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerIntercontrollerArrayOutput() WirelesscontrollerIntercontrollerArrayOutput
	ToWirelesscontrollerIntercontrollerArrayOutputWithContext(context.Context) WirelesscontrollerIntercontrollerArrayOutput
}

type WirelesscontrollerIntercontrollerArray []WirelesscontrollerIntercontrollerInput

func (WirelesscontrollerIntercontrollerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerIntercontroller)(nil)).Elem()
}

func (i WirelesscontrollerIntercontrollerArray) ToWirelesscontrollerIntercontrollerArrayOutput() WirelesscontrollerIntercontrollerArrayOutput {
	return i.ToWirelesscontrollerIntercontrollerArrayOutputWithContext(context.Background())
}

func (i WirelesscontrollerIntercontrollerArray) ToWirelesscontrollerIntercontrollerArrayOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerIntercontrollerArrayOutput)
}

// WirelesscontrollerIntercontrollerMapInput is an input type that accepts WirelesscontrollerIntercontrollerMap and WirelesscontrollerIntercontrollerMapOutput values.
// You can construct a concrete instance of `WirelesscontrollerIntercontrollerMapInput` via:
//
//	WirelesscontrollerIntercontrollerMap{ "key": WirelesscontrollerIntercontrollerArgs{...} }
type WirelesscontrollerIntercontrollerMapInput interface {
	pulumi.Input

	ToWirelesscontrollerIntercontrollerMapOutput() WirelesscontrollerIntercontrollerMapOutput
	ToWirelesscontrollerIntercontrollerMapOutputWithContext(context.Context) WirelesscontrollerIntercontrollerMapOutput
}

type WirelesscontrollerIntercontrollerMap map[string]WirelesscontrollerIntercontrollerInput

func (WirelesscontrollerIntercontrollerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerIntercontroller)(nil)).Elem()
}

func (i WirelesscontrollerIntercontrollerMap) ToWirelesscontrollerIntercontrollerMapOutput() WirelesscontrollerIntercontrollerMapOutput {
	return i.ToWirelesscontrollerIntercontrollerMapOutputWithContext(context.Background())
}

func (i WirelesscontrollerIntercontrollerMap) ToWirelesscontrollerIntercontrollerMapOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WirelesscontrollerIntercontrollerMapOutput)
}

type WirelesscontrollerIntercontrollerOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerIntercontrollerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WirelesscontrollerIntercontroller)(nil)).Elem()
}

func (o WirelesscontrollerIntercontrollerOutput) ToWirelesscontrollerIntercontrollerOutput() WirelesscontrollerIntercontrollerOutput {
	return o
}

func (o WirelesscontrollerIntercontrollerOutput) ToWirelesscontrollerIntercontrollerOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o WirelesscontrollerIntercontrollerOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
func (o WirelesscontrollerIntercontrollerOutput) FastFailoverMax() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.IntOutput { return v.FastFailoverMax }).(pulumi.IntOutput)
}

// Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
func (o WirelesscontrollerIntercontrollerOutput) FastFailoverWait() pulumi.IntOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.IntOutput { return v.FastFailoverWait }).(pulumi.IntOutput)
}

// Secret key for inter-controller communications.
func (o WirelesscontrollerIntercontrollerOutput) InterControllerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.StringPtrOutput { return v.InterControllerKey }).(pulumi.StringPtrOutput)
}

// Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
func (o WirelesscontrollerIntercontrollerOutput) InterControllerMode() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.StringOutput { return v.InterControllerMode }).(pulumi.StringOutput)
}

// Fast failover peer wireless controller list. The structure of `interControllerPeer` block is documented below.
func (o WirelesscontrollerIntercontrollerOutput) InterControllerPeers() WirelesscontrollerIntercontrollerInterControllerPeerArrayOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) WirelesscontrollerIntercontrollerInterControllerPeerArrayOutput {
		return v.InterControllerPeers
	}).(WirelesscontrollerIntercontrollerInterControllerPeerArrayOutput)
}

// Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
func (o WirelesscontrollerIntercontrollerOutput) InterControllerPri() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.StringOutput { return v.InterControllerPri }).(pulumi.StringOutput)
}

// Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
func (o WirelesscontrollerIntercontrollerOutput) L3Roaming() pulumi.StringOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.StringOutput { return v.L3Roaming }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WirelesscontrollerIntercontrollerOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WirelesscontrollerIntercontroller) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WirelesscontrollerIntercontrollerArrayOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerIntercontrollerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WirelesscontrollerIntercontroller)(nil)).Elem()
}

func (o WirelesscontrollerIntercontrollerArrayOutput) ToWirelesscontrollerIntercontrollerArrayOutput() WirelesscontrollerIntercontrollerArrayOutput {
	return o
}

func (o WirelesscontrollerIntercontrollerArrayOutput) ToWirelesscontrollerIntercontrollerArrayOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerArrayOutput {
	return o
}

func (o WirelesscontrollerIntercontrollerArrayOutput) Index(i pulumi.IntInput) WirelesscontrollerIntercontrollerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WirelesscontrollerIntercontroller {
		return vs[0].([]*WirelesscontrollerIntercontroller)[vs[1].(int)]
	}).(WirelesscontrollerIntercontrollerOutput)
}

type WirelesscontrollerIntercontrollerMapOutput struct{ *pulumi.OutputState }

func (WirelesscontrollerIntercontrollerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WirelesscontrollerIntercontroller)(nil)).Elem()
}

func (o WirelesscontrollerIntercontrollerMapOutput) ToWirelesscontrollerIntercontrollerMapOutput() WirelesscontrollerIntercontrollerMapOutput {
	return o
}

func (o WirelesscontrollerIntercontrollerMapOutput) ToWirelesscontrollerIntercontrollerMapOutputWithContext(ctx context.Context) WirelesscontrollerIntercontrollerMapOutput {
	return o
}

func (o WirelesscontrollerIntercontrollerMapOutput) MapIndex(k pulumi.StringInput) WirelesscontrollerIntercontrollerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WirelesscontrollerIntercontroller {
		return vs[0].(map[string]*WirelesscontrollerIntercontroller)[vs[1].(string)]
	}).(WirelesscontrollerIntercontrollerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerIntercontrollerInput)(nil)).Elem(), &WirelesscontrollerIntercontroller{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerIntercontrollerArrayInput)(nil)).Elem(), WirelesscontrollerIntercontrollerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WirelesscontrollerIntercontrollerMapInput)(nil)).Elem(), WirelesscontrollerIntercontrollerMap{})
	pulumi.RegisterOutputType(WirelesscontrollerIntercontrollerOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerIntercontrollerArrayOutput{})
	pulumi.RegisterOutputType(WirelesscontrollerIntercontrollerMapOutput{})
}
