// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wanopt

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/internal"
)

// Configure a remote cache device as Web cache storage.
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
//	"github.com/pulumiverse/pulumi-fortios/sdk/go/fortios/wanopt"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wanopt.NewRemotestorage(ctx, "trname", &wanopt.RemotestorageArgs{
//				RemoteCacheIp: pulumi.String("0.0.0.0"),
//				Status:        pulumi.String("disable"),
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
// Wanopt RemoteStorage can be imported using any of these accepted formats:
//
// ```sh
// $ pulumi import fortios:wanopt/remotestorage:Remotestorage labelname WanoptRemoteStorage
// ```
//
// If you do not want to import arguments of block:
//
// $ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
// $ pulumi import fortios:wanopt/remotestorage:Remotestorage labelname WanoptRemoteStorage
// ```
//
// $ unset "FORTIOS_IMPORT_TABLE"
type Remotestorage struct {
	pulumi.CustomResourceState

	// ID that this device uses to connect to the remote device.
	LocalCacheId pulumi.StringOutput `pulumi:"localCacheId"`
	// ID of the remote device to which the device connects.
	RemoteCacheId pulumi.StringOutput `pulumi:"remoteCacheId"`
	// IP address of the remote device to which the device connects.
	RemoteCacheIp pulumi.StringOutput `pulumi:"remoteCacheIp"`
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewRemotestorage registers a new resource with the given unique name, arguments, and options.
func NewRemotestorage(ctx *pulumi.Context,
	name string, args *RemotestorageArgs, opts ...pulumi.ResourceOption) (*Remotestorage, error) {
	if args == nil {
		args = &RemotestorageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Remotestorage
	err := ctx.RegisterResource("fortios:wanopt/remotestorage:Remotestorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemotestorage gets an existing Remotestorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemotestorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemotestorageState, opts ...pulumi.ResourceOption) (*Remotestorage, error) {
	var resource Remotestorage
	err := ctx.ReadResource("fortios:wanopt/remotestorage:Remotestorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Remotestorage resources.
type remotestorageState struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId *string `pulumi:"localCacheId"`
	// ID of the remote device to which the device connects.
	RemoteCacheId *string `pulumi:"remoteCacheId"`
	// IP address of the remote device to which the device connects.
	RemoteCacheIp *string `pulumi:"remoteCacheIp"`
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type RemotestorageState struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId pulumi.StringPtrInput
	// ID of the remote device to which the device connects.
	RemoteCacheId pulumi.StringPtrInput
	// IP address of the remote device to which the device connects.
	RemoteCacheIp pulumi.StringPtrInput
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RemotestorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*remotestorageState)(nil)).Elem()
}

type remotestorageArgs struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId *string `pulumi:"localCacheId"`
	// ID of the remote device to which the device connects.
	RemoteCacheId *string `pulumi:"remoteCacheId"`
	// IP address of the remote device to which the device connects.
	RemoteCacheIp *string `pulumi:"remoteCacheIp"`
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Remotestorage resource.
type RemotestorageArgs struct {
	// ID that this device uses to connect to the remote device.
	LocalCacheId pulumi.StringPtrInput
	// ID of the remote device to which the device connects.
	RemoteCacheId pulumi.StringPtrInput
	// IP address of the remote device to which the device connects.
	RemoteCacheIp pulumi.StringPtrInput
	// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (RemotestorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remotestorageArgs)(nil)).Elem()
}

type RemotestorageInput interface {
	pulumi.Input

	ToRemotestorageOutput() RemotestorageOutput
	ToRemotestorageOutputWithContext(ctx context.Context) RemotestorageOutput
}

func (*Remotestorage) ElementType() reflect.Type {
	return reflect.TypeOf((**Remotestorage)(nil)).Elem()
}

func (i *Remotestorage) ToRemotestorageOutput() RemotestorageOutput {
	return i.ToRemotestorageOutputWithContext(context.Background())
}

func (i *Remotestorage) ToRemotestorageOutputWithContext(ctx context.Context) RemotestorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotestorageOutput)
}

// RemotestorageArrayInput is an input type that accepts RemotestorageArray and RemotestorageArrayOutput values.
// You can construct a concrete instance of `RemotestorageArrayInput` via:
//
//	RemotestorageArray{ RemotestorageArgs{...} }
type RemotestorageArrayInput interface {
	pulumi.Input

	ToRemotestorageArrayOutput() RemotestorageArrayOutput
	ToRemotestorageArrayOutputWithContext(context.Context) RemotestorageArrayOutput
}

type RemotestorageArray []RemotestorageInput

func (RemotestorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Remotestorage)(nil)).Elem()
}

func (i RemotestorageArray) ToRemotestorageArrayOutput() RemotestorageArrayOutput {
	return i.ToRemotestorageArrayOutputWithContext(context.Background())
}

func (i RemotestorageArray) ToRemotestorageArrayOutputWithContext(ctx context.Context) RemotestorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotestorageArrayOutput)
}

// RemotestorageMapInput is an input type that accepts RemotestorageMap and RemotestorageMapOutput values.
// You can construct a concrete instance of `RemotestorageMapInput` via:
//
//	RemotestorageMap{ "key": RemotestorageArgs{...} }
type RemotestorageMapInput interface {
	pulumi.Input

	ToRemotestorageMapOutput() RemotestorageMapOutput
	ToRemotestorageMapOutputWithContext(context.Context) RemotestorageMapOutput
}

type RemotestorageMap map[string]RemotestorageInput

func (RemotestorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Remotestorage)(nil)).Elem()
}

func (i RemotestorageMap) ToRemotestorageMapOutput() RemotestorageMapOutput {
	return i.ToRemotestorageMapOutputWithContext(context.Background())
}

func (i RemotestorageMap) ToRemotestorageMapOutputWithContext(ctx context.Context) RemotestorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemotestorageMapOutput)
}

type RemotestorageOutput struct{ *pulumi.OutputState }

func (RemotestorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Remotestorage)(nil)).Elem()
}

func (o RemotestorageOutput) ToRemotestorageOutput() RemotestorageOutput {
	return o
}

func (o RemotestorageOutput) ToRemotestorageOutputWithContext(ctx context.Context) RemotestorageOutput {
	return o
}

// ID that this device uses to connect to the remote device.
func (o RemotestorageOutput) LocalCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *Remotestorage) pulumi.StringOutput { return v.LocalCacheId }).(pulumi.StringOutput)
}

// ID of the remote device to which the device connects.
func (o RemotestorageOutput) RemoteCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v *Remotestorage) pulumi.StringOutput { return v.RemoteCacheId }).(pulumi.StringOutput)
}

// IP address of the remote device to which the device connects.
func (o RemotestorageOutput) RemoteCacheIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Remotestorage) pulumi.StringOutput { return v.RemoteCacheIp }).(pulumi.StringOutput)
}

// Enable/disable using remote device as Web cache storage. Valid values: `disable`, `enable`.
func (o RemotestorageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Remotestorage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o RemotestorageOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Remotestorage) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type RemotestorageArrayOutput struct{ *pulumi.OutputState }

func (RemotestorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Remotestorage)(nil)).Elem()
}

func (o RemotestorageArrayOutput) ToRemotestorageArrayOutput() RemotestorageArrayOutput {
	return o
}

func (o RemotestorageArrayOutput) ToRemotestorageArrayOutputWithContext(ctx context.Context) RemotestorageArrayOutput {
	return o
}

func (o RemotestorageArrayOutput) Index(i pulumi.IntInput) RemotestorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Remotestorage {
		return vs[0].([]*Remotestorage)[vs[1].(int)]
	}).(RemotestorageOutput)
}

type RemotestorageMapOutput struct{ *pulumi.OutputState }

func (RemotestorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Remotestorage)(nil)).Elem()
}

func (o RemotestorageMapOutput) ToRemotestorageMapOutput() RemotestorageMapOutput {
	return o
}

func (o RemotestorageMapOutput) ToRemotestorageMapOutputWithContext(ctx context.Context) RemotestorageMapOutput {
	return o
}

func (o RemotestorageMapOutput) MapIndex(k pulumi.StringInput) RemotestorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Remotestorage {
		return vs[0].(map[string]*Remotestorage)[vs[1].(string)]
	}).(RemotestorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemotestorageInput)(nil)).Elem(), &Remotestorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemotestorageArrayInput)(nil)).Elem(), RemotestorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemotestorageMapInput)(nil)).Elem(), RemotestorageMap{})
	pulumi.RegisterOutputType(RemotestorageOutput{})
	pulumi.RegisterOutputType(RemotestorageArrayOutput{})
	pulumi.RegisterOutputType(RemotestorageMapOutput{})
}