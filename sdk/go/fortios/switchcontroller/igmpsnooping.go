// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package switchcontroller

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure FortiSwitch IGMP snooping global settings.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/switchcontroller"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := switchcontroller.NewIgmpsnooping(ctx, "trname", &switchcontroller.IgmpsnoopingArgs{
//				AgingTime:             pulumi.Int(300),
//				FloodUnknownMulticast: pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// SwitchController IgmpSnooping can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:switchcontroller/igmpsnooping:Igmpsnooping labelname SwitchControllerIgmpSnooping
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:switchcontroller/igmpsnooping:Igmpsnooping labelname SwitchControllerIgmpSnooping
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Igmpsnooping struct {
	pulumi.CustomResourceState

	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime pulumi.IntOutput `pulumi:"agingTime"`
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast pulumi.StringOutput `pulumi:"floodUnknownMulticast"`
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval pulumi.IntOutput `pulumi:"queryInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewIgmpsnooping registers a new resource with the given unique name, arguments, and options.
func NewIgmpsnooping(ctx *pulumi.Context,
	name string, args *IgmpsnoopingArgs, opts ...pulumi.ResourceOption) (*Igmpsnooping, error) {
	if args == nil {
		args = &IgmpsnoopingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Igmpsnooping
	err := ctx.RegisterResource("fortios:switchcontroller/igmpsnooping:Igmpsnooping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIgmpsnooping gets an existing Igmpsnooping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIgmpsnooping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IgmpsnoopingState, opts ...pulumi.ResourceOption) (*Igmpsnooping, error) {
	var resource Igmpsnooping
	err := ctx.ReadResource("fortios:switchcontroller/igmpsnooping:Igmpsnooping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Igmpsnooping resources.
type igmpsnoopingState struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime *int `pulumi:"agingTime"`
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast *string `pulumi:"floodUnknownMulticast"`
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval *int `pulumi:"queryInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type IgmpsnoopingState struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime pulumi.IntPtrInput
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast pulumi.StringPtrInput
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IgmpsnoopingState) ElementType() reflect.Type {
	return reflect.TypeOf((*igmpsnoopingState)(nil)).Elem()
}

type igmpsnoopingArgs struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime *int `pulumi:"agingTime"`
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast *string `pulumi:"floodUnknownMulticast"`
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval *int `pulumi:"queryInterval"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Igmpsnooping resource.
type IgmpsnoopingArgs struct {
	// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
	AgingTime pulumi.IntPtrInput
	// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
	FloodUnknownMulticast pulumi.StringPtrInput
	// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
	QueryInterval pulumi.IntPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (IgmpsnoopingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*igmpsnoopingArgs)(nil)).Elem()
}

type IgmpsnoopingInput interface {
	pulumi.Input

	ToIgmpsnoopingOutput() IgmpsnoopingOutput
	ToIgmpsnoopingOutputWithContext(ctx context.Context) IgmpsnoopingOutput
}

func (*Igmpsnooping) ElementType() reflect.Type {
	return reflect.TypeOf((**Igmpsnooping)(nil)).Elem()
}

func (i *Igmpsnooping) ToIgmpsnoopingOutput() IgmpsnoopingOutput {
	return i.ToIgmpsnoopingOutputWithContext(context.Background())
}

func (i *Igmpsnooping) ToIgmpsnoopingOutputWithContext(ctx context.Context) IgmpsnoopingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IgmpsnoopingOutput)
}

// IgmpsnoopingArrayInput is an input type that accepts IgmpsnoopingArray and IgmpsnoopingArrayOutput values.
// You can construct a concrete instance of `IgmpsnoopingArrayInput` via:
//
//	IgmpsnoopingArray{ IgmpsnoopingArgs{...} }
type IgmpsnoopingArrayInput interface {
	pulumi.Input

	ToIgmpsnoopingArrayOutput() IgmpsnoopingArrayOutput
	ToIgmpsnoopingArrayOutputWithContext(context.Context) IgmpsnoopingArrayOutput
}

type IgmpsnoopingArray []IgmpsnoopingInput

func (IgmpsnoopingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Igmpsnooping)(nil)).Elem()
}

func (i IgmpsnoopingArray) ToIgmpsnoopingArrayOutput() IgmpsnoopingArrayOutput {
	return i.ToIgmpsnoopingArrayOutputWithContext(context.Background())
}

func (i IgmpsnoopingArray) ToIgmpsnoopingArrayOutputWithContext(ctx context.Context) IgmpsnoopingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IgmpsnoopingArrayOutput)
}

// IgmpsnoopingMapInput is an input type that accepts IgmpsnoopingMap and IgmpsnoopingMapOutput values.
// You can construct a concrete instance of `IgmpsnoopingMapInput` via:
//
//	IgmpsnoopingMap{ "key": IgmpsnoopingArgs{...} }
type IgmpsnoopingMapInput interface {
	pulumi.Input

	ToIgmpsnoopingMapOutput() IgmpsnoopingMapOutput
	ToIgmpsnoopingMapOutputWithContext(context.Context) IgmpsnoopingMapOutput
}

type IgmpsnoopingMap map[string]IgmpsnoopingInput

func (IgmpsnoopingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Igmpsnooping)(nil)).Elem()
}

func (i IgmpsnoopingMap) ToIgmpsnoopingMapOutput() IgmpsnoopingMapOutput {
	return i.ToIgmpsnoopingMapOutputWithContext(context.Background())
}

func (i IgmpsnoopingMap) ToIgmpsnoopingMapOutputWithContext(ctx context.Context) IgmpsnoopingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IgmpsnoopingMapOutput)
}

type IgmpsnoopingOutput struct{ *pulumi.OutputState }

func (IgmpsnoopingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Igmpsnooping)(nil)).Elem()
}

func (o IgmpsnoopingOutput) ToIgmpsnoopingOutput() IgmpsnoopingOutput {
	return o
}

func (o IgmpsnoopingOutput) ToIgmpsnoopingOutputWithContext(ctx context.Context) IgmpsnoopingOutput {
	return o
}

// Maximum number of seconds to retain a multicast snooping entry for which no packets have been seen (15 - 3600 sec, default = 300).
func (o IgmpsnoopingOutput) AgingTime() pulumi.IntOutput {
	return o.ApplyT(func(v *Igmpsnooping) pulumi.IntOutput { return v.AgingTime }).(pulumi.IntOutput)
}

// Enable/disable unknown multicast flooding. Valid values: `enable`, `disable`.
func (o IgmpsnoopingOutput) FloodUnknownMulticast() pulumi.StringOutput {
	return o.ApplyT(func(v *Igmpsnooping) pulumi.StringOutput { return v.FloodUnknownMulticast }).(pulumi.StringOutput)
}

// Maximum time after which IGMP query will be sent (10 - 1200 sec, default = 125).
func (o IgmpsnoopingOutput) QueryInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *Igmpsnooping) pulumi.IntOutput { return v.QueryInterval }).(pulumi.IntOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o IgmpsnoopingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Igmpsnooping) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type IgmpsnoopingArrayOutput struct{ *pulumi.OutputState }

func (IgmpsnoopingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Igmpsnooping)(nil)).Elem()
}

func (o IgmpsnoopingArrayOutput) ToIgmpsnoopingArrayOutput() IgmpsnoopingArrayOutput {
	return o
}

func (o IgmpsnoopingArrayOutput) ToIgmpsnoopingArrayOutputWithContext(ctx context.Context) IgmpsnoopingArrayOutput {
	return o
}

func (o IgmpsnoopingArrayOutput) Index(i pulumi.IntInput) IgmpsnoopingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Igmpsnooping {
		return vs[0].([]*Igmpsnooping)[vs[1].(int)]
	}).(IgmpsnoopingOutput)
}

type IgmpsnoopingMapOutput struct{ *pulumi.OutputState }

func (IgmpsnoopingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Igmpsnooping)(nil)).Elem()
}

func (o IgmpsnoopingMapOutput) ToIgmpsnoopingMapOutput() IgmpsnoopingMapOutput {
	return o
}

func (o IgmpsnoopingMapOutput) ToIgmpsnoopingMapOutputWithContext(ctx context.Context) IgmpsnoopingMapOutput {
	return o
}

func (o IgmpsnoopingMapOutput) MapIndex(k pulumi.StringInput) IgmpsnoopingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Igmpsnooping {
		return vs[0].(map[string]*Igmpsnooping)[vs[1].(string)]
	}).(IgmpsnoopingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IgmpsnoopingInput)(nil)).Elem(), &Igmpsnooping{})
	pulumi.RegisterInputType(reflect.TypeOf((*IgmpsnoopingArrayInput)(nil)).Elem(), IgmpsnoopingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IgmpsnoopingMapInput)(nil)).Elem(), IgmpsnoopingMap{})
	pulumi.RegisterOutputType(IgmpsnoopingOutput{})
	pulumi.RegisterOutputType(IgmpsnoopingArrayOutput{})
	pulumi.RegisterOutputType(IgmpsnoopingMapOutput{})
}