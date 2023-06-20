// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// FortiGate controller configuration. Applies to FortiOS Version `>= 7.2.1`.
//
// ## Import
//
// # ExtensionController Fortigate can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/extensioncontrollerFortigate:ExtensioncontrollerFortigate labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/extensioncontrollerFortigate:ExtensioncontrollerFortigate labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type ExtensioncontrollerFortigate struct {
	pulumi.CustomResourceState

	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized pulumi.StringOutput `pulumi:"authorized"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// device-id
	DeviceId pulumi.IntOutput `pulumi:"deviceId"`
	// FortiGate serial number.
	Fosid pulumi.StringOutput `pulumi:"fosid"`
	// FortiGate hostname.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// FortiGate entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// FortiGate profile configuration.
	Profile pulumi.StringOutput `pulumi:"profile"`
	// VDOM.
	Vdom pulumi.IntOutput `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewExtensioncontrollerFortigate registers a new resource with the given unique name, arguments, and options.
func NewExtensioncontrollerFortigate(ctx *pulumi.Context,
	name string, args *ExtensioncontrollerFortigateArgs, opts ...pulumi.ResourceOption) (*ExtensioncontrollerFortigate, error) {
	if args == nil {
		args = &ExtensioncontrollerFortigateArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ExtensioncontrollerFortigate
	err := ctx.RegisterResource("fortios:index/extensioncontrollerFortigate:ExtensioncontrollerFortigate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtensioncontrollerFortigate gets an existing ExtensioncontrollerFortigate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtensioncontrollerFortigate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensioncontrollerFortigateState, opts ...pulumi.ResourceOption) (*ExtensioncontrollerFortigate, error) {
	var resource ExtensioncontrollerFortigate
	err := ctx.ReadResource("fortios:index/extensioncontrollerFortigate:ExtensioncontrollerFortigate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExtensioncontrollerFortigate resources.
type extensioncontrollerFortigateState struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// Description.
	Description *string `pulumi:"description"`
	// device-id
	DeviceId *int `pulumi:"deviceId"`
	// FortiGate serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiGate hostname.
	Hostname *string `pulumi:"hostname"`
	// FortiGate entry name.
	Name *string `pulumi:"name"`
	// FortiGate profile configuration.
	Profile *string `pulumi:"profile"`
	// VDOM.
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type ExtensioncontrollerFortigateState struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// device-id
	DeviceId pulumi.IntPtrInput
	// FortiGate serial number.
	Fosid pulumi.StringPtrInput
	// FortiGate hostname.
	Hostname pulumi.StringPtrInput
	// FortiGate entry name.
	Name pulumi.StringPtrInput
	// FortiGate profile configuration.
	Profile pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtensioncontrollerFortigateState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensioncontrollerFortigateState)(nil)).Elem()
}

type extensioncontrollerFortigateArgs struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized *string `pulumi:"authorized"`
	// Description.
	Description *string `pulumi:"description"`
	// device-id
	DeviceId *int `pulumi:"deviceId"`
	// FortiGate serial number.
	Fosid *string `pulumi:"fosid"`
	// FortiGate hostname.
	Hostname *string `pulumi:"hostname"`
	// FortiGate entry name.
	Name *string `pulumi:"name"`
	// FortiGate profile configuration.
	Profile *string `pulumi:"profile"`
	// VDOM.
	Vdom *int `pulumi:"vdom"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a ExtensioncontrollerFortigate resource.
type ExtensioncontrollerFortigateArgs struct {
	// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
	Authorized pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// device-id
	DeviceId pulumi.IntPtrInput
	// FortiGate serial number.
	Fosid pulumi.StringPtrInput
	// FortiGate hostname.
	Hostname pulumi.StringPtrInput
	// FortiGate entry name.
	Name pulumi.StringPtrInput
	// FortiGate profile configuration.
	Profile pulumi.StringPtrInput
	// VDOM.
	Vdom pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (ExtensioncontrollerFortigateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensioncontrollerFortigateArgs)(nil)).Elem()
}

type ExtensioncontrollerFortigateInput interface {
	pulumi.Input

	ToExtensioncontrollerFortigateOutput() ExtensioncontrollerFortigateOutput
	ToExtensioncontrollerFortigateOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateOutput
}

func (*ExtensioncontrollerFortigate) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensioncontrollerFortigate)(nil)).Elem()
}

func (i *ExtensioncontrollerFortigate) ToExtensioncontrollerFortigateOutput() ExtensioncontrollerFortigateOutput {
	return i.ToExtensioncontrollerFortigateOutputWithContext(context.Background())
}

func (i *ExtensioncontrollerFortigate) ToExtensioncontrollerFortigateOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerFortigateOutput)
}

// ExtensioncontrollerFortigateArrayInput is an input type that accepts ExtensioncontrollerFortigateArray and ExtensioncontrollerFortigateArrayOutput values.
// You can construct a concrete instance of `ExtensioncontrollerFortigateArrayInput` via:
//
//	ExtensioncontrollerFortigateArray{ ExtensioncontrollerFortigateArgs{...} }
type ExtensioncontrollerFortigateArrayInput interface {
	pulumi.Input

	ToExtensioncontrollerFortigateArrayOutput() ExtensioncontrollerFortigateArrayOutput
	ToExtensioncontrollerFortigateArrayOutputWithContext(context.Context) ExtensioncontrollerFortigateArrayOutput
}

type ExtensioncontrollerFortigateArray []ExtensioncontrollerFortigateInput

func (ExtensioncontrollerFortigateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensioncontrollerFortigate)(nil)).Elem()
}

func (i ExtensioncontrollerFortigateArray) ToExtensioncontrollerFortigateArrayOutput() ExtensioncontrollerFortigateArrayOutput {
	return i.ToExtensioncontrollerFortigateArrayOutputWithContext(context.Background())
}

func (i ExtensioncontrollerFortigateArray) ToExtensioncontrollerFortigateArrayOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerFortigateArrayOutput)
}

// ExtensioncontrollerFortigateMapInput is an input type that accepts ExtensioncontrollerFortigateMap and ExtensioncontrollerFortigateMapOutput values.
// You can construct a concrete instance of `ExtensioncontrollerFortigateMapInput` via:
//
//	ExtensioncontrollerFortigateMap{ "key": ExtensioncontrollerFortigateArgs{...} }
type ExtensioncontrollerFortigateMapInput interface {
	pulumi.Input

	ToExtensioncontrollerFortigateMapOutput() ExtensioncontrollerFortigateMapOutput
	ToExtensioncontrollerFortigateMapOutputWithContext(context.Context) ExtensioncontrollerFortigateMapOutput
}

type ExtensioncontrollerFortigateMap map[string]ExtensioncontrollerFortigateInput

func (ExtensioncontrollerFortigateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensioncontrollerFortigate)(nil)).Elem()
}

func (i ExtensioncontrollerFortigateMap) ToExtensioncontrollerFortigateMapOutput() ExtensioncontrollerFortigateMapOutput {
	return i.ToExtensioncontrollerFortigateMapOutputWithContext(context.Background())
}

func (i ExtensioncontrollerFortigateMap) ToExtensioncontrollerFortigateMapOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtensioncontrollerFortigateMapOutput)
}

type ExtensioncontrollerFortigateOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerFortigateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtensioncontrollerFortigate)(nil)).Elem()
}

func (o ExtensioncontrollerFortigateOutput) ToExtensioncontrollerFortigateOutput() ExtensioncontrollerFortigateOutput {
	return o
}

func (o ExtensioncontrollerFortigateOutput) ToExtensioncontrollerFortigateOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateOutput {
	return o
}

// Enable/disable FortiGate administration. Valid values: `disable`, `enable`.
func (o ExtensioncontrollerFortigateOutput) Authorized() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringOutput { return v.Authorized }).(pulumi.StringOutput)
}

// Description.
func (o ExtensioncontrollerFortigateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// device-id
func (o ExtensioncontrollerFortigateOutput) DeviceId() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.IntOutput { return v.DeviceId }).(pulumi.IntOutput)
}

// FortiGate serial number.
func (o ExtensioncontrollerFortigateOutput) Fosid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringOutput { return v.Fosid }).(pulumi.StringOutput)
}

// FortiGate hostname.
func (o ExtensioncontrollerFortigateOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// FortiGate entry name.
func (o ExtensioncontrollerFortigateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// FortiGate profile configuration.
func (o ExtensioncontrollerFortigateOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringOutput { return v.Profile }).(pulumi.StringOutput)
}

// VDOM.
func (o ExtensioncontrollerFortigateOutput) Vdom() pulumi.IntOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.IntOutput { return v.Vdom }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o ExtensioncontrollerFortigateOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtensioncontrollerFortigate) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type ExtensioncontrollerFortigateArrayOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerFortigateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExtensioncontrollerFortigate)(nil)).Elem()
}

func (o ExtensioncontrollerFortigateArrayOutput) ToExtensioncontrollerFortigateArrayOutput() ExtensioncontrollerFortigateArrayOutput {
	return o
}

func (o ExtensioncontrollerFortigateArrayOutput) ToExtensioncontrollerFortigateArrayOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateArrayOutput {
	return o
}

func (o ExtensioncontrollerFortigateArrayOutput) Index(i pulumi.IntInput) ExtensioncontrollerFortigateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExtensioncontrollerFortigate {
		return vs[0].([]*ExtensioncontrollerFortigate)[vs[1].(int)]
	}).(ExtensioncontrollerFortigateOutput)
}

type ExtensioncontrollerFortigateMapOutput struct{ *pulumi.OutputState }

func (ExtensioncontrollerFortigateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExtensioncontrollerFortigate)(nil)).Elem()
}

func (o ExtensioncontrollerFortigateMapOutput) ToExtensioncontrollerFortigateMapOutput() ExtensioncontrollerFortigateMapOutput {
	return o
}

func (o ExtensioncontrollerFortigateMapOutput) ToExtensioncontrollerFortigateMapOutputWithContext(ctx context.Context) ExtensioncontrollerFortigateMapOutput {
	return o
}

func (o ExtensioncontrollerFortigateMapOutput) MapIndex(k pulumi.StringInput) ExtensioncontrollerFortigateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExtensioncontrollerFortigate {
		return vs[0].(map[string]*ExtensioncontrollerFortigate)[vs[1].(string)]
	}).(ExtensioncontrollerFortigateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerFortigateInput)(nil)).Elem(), &ExtensioncontrollerFortigate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerFortigateArrayInput)(nil)).Elem(), ExtensioncontrollerFortigateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExtensioncontrollerFortigateMapInput)(nil)).Elem(), ExtensioncontrollerFortigateMap{})
	pulumi.RegisterOutputType(ExtensioncontrollerFortigateOutput{})
	pulumi.RegisterOutputType(ExtensioncontrollerFortigateArrayOutput{})
	pulumi.RegisterOutputType(ExtensioncontrollerFortigateMapOutput{})
}
