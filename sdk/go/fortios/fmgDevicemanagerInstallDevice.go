// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource supports installing devicemanager script from FortiManager to the related device
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
//			_, err := fortios.NewFmgDevicemanagerInstallDevice(ctx, "test1", &fortios.FmgDevicemanagerInstallDeviceArgs{
//				TargetDevname: pulumi.String("FGVM64-test"),
//				Timeout:       pulumi.Int(5),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type FmgDevicemanagerInstallDevice struct {
	pulumi.CustomResourceState

	// Source ADOM name. default is 'root'
	Adom pulumi.StringPtrOutput `pulumi:"adom"`
	// Target device name.
	TargetDevname pulumi.StringOutput `pulumi:"targetDevname"`
	// Timeout for installing the script to the target, default: 3 minutes.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Vdom of managed device. default is 'root'
	Vdom pulumi.StringPtrOutput `pulumi:"vdom"`
}

// NewFmgDevicemanagerInstallDevice registers a new resource with the given unique name, arguments, and options.
func NewFmgDevicemanagerInstallDevice(ctx *pulumi.Context,
	name string, args *FmgDevicemanagerInstallDeviceArgs, opts ...pulumi.ResourceOption) (*FmgDevicemanagerInstallDevice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetDevname == nil {
		return nil, errors.New("invalid value for required argument 'TargetDevname'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FmgDevicemanagerInstallDevice
	err := ctx.RegisterResource("fortios:index/fmgDevicemanagerInstallDevice:FmgDevicemanagerInstallDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFmgDevicemanagerInstallDevice gets an existing FmgDevicemanagerInstallDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFmgDevicemanagerInstallDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FmgDevicemanagerInstallDeviceState, opts ...pulumi.ResourceOption) (*FmgDevicemanagerInstallDevice, error) {
	var resource FmgDevicemanagerInstallDevice
	err := ctx.ReadResource("fortios:index/fmgDevicemanagerInstallDevice:FmgDevicemanagerInstallDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FmgDevicemanagerInstallDevice resources.
type fmgDevicemanagerInstallDeviceState struct {
	// Source ADOM name. default is 'root'
	Adom *string `pulumi:"adom"`
	// Target device name.
	TargetDevname *string `pulumi:"targetDevname"`
	// Timeout for installing the script to the target, default: 3 minutes.
	Timeout *int `pulumi:"timeout"`
	// Vdom of managed device. default is 'root'
	Vdom *string `pulumi:"vdom"`
}

type FmgDevicemanagerInstallDeviceState struct {
	// Source ADOM name. default is 'root'
	Adom pulumi.StringPtrInput
	// Target device name.
	TargetDevname pulumi.StringPtrInput
	// Timeout for installing the script to the target, default: 3 minutes.
	Timeout pulumi.IntPtrInput
	// Vdom of managed device. default is 'root'
	Vdom pulumi.StringPtrInput
}

func (FmgDevicemanagerInstallDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgDevicemanagerInstallDeviceState)(nil)).Elem()
}

type fmgDevicemanagerInstallDeviceArgs struct {
	// Source ADOM name. default is 'root'
	Adom *string `pulumi:"adom"`
	// Target device name.
	TargetDevname string `pulumi:"targetDevname"`
	// Timeout for installing the script to the target, default: 3 minutes.
	Timeout *int `pulumi:"timeout"`
	// Vdom of managed device. default is 'root'
	Vdom *string `pulumi:"vdom"`
}

// The set of arguments for constructing a FmgDevicemanagerInstallDevice resource.
type FmgDevicemanagerInstallDeviceArgs struct {
	// Source ADOM name. default is 'root'
	Adom pulumi.StringPtrInput
	// Target device name.
	TargetDevname pulumi.StringInput
	// Timeout for installing the script to the target, default: 3 minutes.
	Timeout pulumi.IntPtrInput
	// Vdom of managed device. default is 'root'
	Vdom pulumi.StringPtrInput
}

func (FmgDevicemanagerInstallDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fmgDevicemanagerInstallDeviceArgs)(nil)).Elem()
}

type FmgDevicemanagerInstallDeviceInput interface {
	pulumi.Input

	ToFmgDevicemanagerInstallDeviceOutput() FmgDevicemanagerInstallDeviceOutput
	ToFmgDevicemanagerInstallDeviceOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceOutput
}

func (*FmgDevicemanagerInstallDevice) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgDevicemanagerInstallDevice)(nil)).Elem()
}

func (i *FmgDevicemanagerInstallDevice) ToFmgDevicemanagerInstallDeviceOutput() FmgDevicemanagerInstallDeviceOutput {
	return i.ToFmgDevicemanagerInstallDeviceOutputWithContext(context.Background())
}

func (i *FmgDevicemanagerInstallDevice) ToFmgDevicemanagerInstallDeviceOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgDevicemanagerInstallDeviceOutput)
}

// FmgDevicemanagerInstallDeviceArrayInput is an input type that accepts FmgDevicemanagerInstallDeviceArray and FmgDevicemanagerInstallDeviceArrayOutput values.
// You can construct a concrete instance of `FmgDevicemanagerInstallDeviceArrayInput` via:
//
//	FmgDevicemanagerInstallDeviceArray{ FmgDevicemanagerInstallDeviceArgs{...} }
type FmgDevicemanagerInstallDeviceArrayInput interface {
	pulumi.Input

	ToFmgDevicemanagerInstallDeviceArrayOutput() FmgDevicemanagerInstallDeviceArrayOutput
	ToFmgDevicemanagerInstallDeviceArrayOutputWithContext(context.Context) FmgDevicemanagerInstallDeviceArrayOutput
}

type FmgDevicemanagerInstallDeviceArray []FmgDevicemanagerInstallDeviceInput

func (FmgDevicemanagerInstallDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgDevicemanagerInstallDevice)(nil)).Elem()
}

func (i FmgDevicemanagerInstallDeviceArray) ToFmgDevicemanagerInstallDeviceArrayOutput() FmgDevicemanagerInstallDeviceArrayOutput {
	return i.ToFmgDevicemanagerInstallDeviceArrayOutputWithContext(context.Background())
}

func (i FmgDevicemanagerInstallDeviceArray) ToFmgDevicemanagerInstallDeviceArrayOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgDevicemanagerInstallDeviceArrayOutput)
}

// FmgDevicemanagerInstallDeviceMapInput is an input type that accepts FmgDevicemanagerInstallDeviceMap and FmgDevicemanagerInstallDeviceMapOutput values.
// You can construct a concrete instance of `FmgDevicemanagerInstallDeviceMapInput` via:
//
//	FmgDevicemanagerInstallDeviceMap{ "key": FmgDevicemanagerInstallDeviceArgs{...} }
type FmgDevicemanagerInstallDeviceMapInput interface {
	pulumi.Input

	ToFmgDevicemanagerInstallDeviceMapOutput() FmgDevicemanagerInstallDeviceMapOutput
	ToFmgDevicemanagerInstallDeviceMapOutputWithContext(context.Context) FmgDevicemanagerInstallDeviceMapOutput
}

type FmgDevicemanagerInstallDeviceMap map[string]FmgDevicemanagerInstallDeviceInput

func (FmgDevicemanagerInstallDeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgDevicemanagerInstallDevice)(nil)).Elem()
}

func (i FmgDevicemanagerInstallDeviceMap) ToFmgDevicemanagerInstallDeviceMapOutput() FmgDevicemanagerInstallDeviceMapOutput {
	return i.ToFmgDevicemanagerInstallDeviceMapOutputWithContext(context.Background())
}

func (i FmgDevicemanagerInstallDeviceMap) ToFmgDevicemanagerInstallDeviceMapOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FmgDevicemanagerInstallDeviceMapOutput)
}

type FmgDevicemanagerInstallDeviceOutput struct{ *pulumi.OutputState }

func (FmgDevicemanagerInstallDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FmgDevicemanagerInstallDevice)(nil)).Elem()
}

func (o FmgDevicemanagerInstallDeviceOutput) ToFmgDevicemanagerInstallDeviceOutput() FmgDevicemanagerInstallDeviceOutput {
	return o
}

func (o FmgDevicemanagerInstallDeviceOutput) ToFmgDevicemanagerInstallDeviceOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceOutput {
	return o
}

// Source ADOM name. default is 'root'
func (o FmgDevicemanagerInstallDeviceOutput) Adom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgDevicemanagerInstallDevice) pulumi.StringPtrOutput { return v.Adom }).(pulumi.StringPtrOutput)
}

// Target device name.
func (o FmgDevicemanagerInstallDeviceOutput) TargetDevname() pulumi.StringOutput {
	return o.ApplyT(func(v *FmgDevicemanagerInstallDevice) pulumi.StringOutput { return v.TargetDevname }).(pulumi.StringOutput)
}

// Timeout for installing the script to the target, default: 3 minutes.
func (o FmgDevicemanagerInstallDeviceOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FmgDevicemanagerInstallDevice) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// Vdom of managed device. default is 'root'
func (o FmgDevicemanagerInstallDeviceOutput) Vdom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FmgDevicemanagerInstallDevice) pulumi.StringPtrOutput { return v.Vdom }).(pulumi.StringPtrOutput)
}

type FmgDevicemanagerInstallDeviceArrayOutput struct{ *pulumi.OutputState }

func (FmgDevicemanagerInstallDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FmgDevicemanagerInstallDevice)(nil)).Elem()
}

func (o FmgDevicemanagerInstallDeviceArrayOutput) ToFmgDevicemanagerInstallDeviceArrayOutput() FmgDevicemanagerInstallDeviceArrayOutput {
	return o
}

func (o FmgDevicemanagerInstallDeviceArrayOutput) ToFmgDevicemanagerInstallDeviceArrayOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceArrayOutput {
	return o
}

func (o FmgDevicemanagerInstallDeviceArrayOutput) Index(i pulumi.IntInput) FmgDevicemanagerInstallDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FmgDevicemanagerInstallDevice {
		return vs[0].([]*FmgDevicemanagerInstallDevice)[vs[1].(int)]
	}).(FmgDevicemanagerInstallDeviceOutput)
}

type FmgDevicemanagerInstallDeviceMapOutput struct{ *pulumi.OutputState }

func (FmgDevicemanagerInstallDeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FmgDevicemanagerInstallDevice)(nil)).Elem()
}

func (o FmgDevicemanagerInstallDeviceMapOutput) ToFmgDevicemanagerInstallDeviceMapOutput() FmgDevicemanagerInstallDeviceMapOutput {
	return o
}

func (o FmgDevicemanagerInstallDeviceMapOutput) ToFmgDevicemanagerInstallDeviceMapOutputWithContext(ctx context.Context) FmgDevicemanagerInstallDeviceMapOutput {
	return o
}

func (o FmgDevicemanagerInstallDeviceMapOutput) MapIndex(k pulumi.StringInput) FmgDevicemanagerInstallDeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FmgDevicemanagerInstallDevice {
		return vs[0].(map[string]*FmgDevicemanagerInstallDevice)[vs[1].(string)]
	}).(FmgDevicemanagerInstallDeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FmgDevicemanagerInstallDeviceInput)(nil)).Elem(), &FmgDevicemanagerInstallDevice{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgDevicemanagerInstallDeviceArrayInput)(nil)).Elem(), FmgDevicemanagerInstallDeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FmgDevicemanagerInstallDeviceMapInput)(nil)).Elem(), FmgDevicemanagerInstallDeviceMap{})
	pulumi.RegisterOutputType(FmgDevicemanagerInstallDeviceOutput{})
	pulumi.RegisterOutputType(FmgDevicemanagerInstallDeviceArrayOutput{})
	pulumi.RegisterOutputType(FmgDevicemanagerInstallDeviceMapOutput{})
}
