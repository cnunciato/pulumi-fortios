// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure IP in IP Tunneling.
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
//			_, err := fortios.NewSystemIpiptunnel(ctx, "trname", &fortios.SystemIpiptunnelArgs{
//				Interface: pulumi.String("port3"),
//				LocalGw:   pulumi.String("1.1.1.1"),
//				RemoteGw:  pulumi.String("2.2.2.2"),
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
// # System IpipTunnel can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemIpiptunnel:SystemIpiptunnel labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemIpiptunnel:SystemIpiptunnel labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemIpiptunnel struct {
	pulumi.CustomResourceState

	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringOutput `pulumi:"autoAsicOffload"`
	// Interface name that is associated with the incoming traffic from available options.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// IPv4 address for the local gateway.
	LocalGw pulumi.StringOutput `pulumi:"localGw"`
	// IPIP Tunnel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// IPv4 address for the remote gateway.
	RemoteGw pulumi.StringOutput `pulumi:"remoteGw"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringOutput `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemIpiptunnel registers a new resource with the given unique name, arguments, and options.
func NewSystemIpiptunnel(ctx *pulumi.Context,
	name string, args *SystemIpiptunnelArgs, opts ...pulumi.ResourceOption) (*SystemIpiptunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Interface == nil {
		return nil, errors.New("invalid value for required argument 'Interface'")
	}
	if args.LocalGw == nil {
		return nil, errors.New("invalid value for required argument 'LocalGw'")
	}
	if args.RemoteGw == nil {
		return nil, errors.New("invalid value for required argument 'RemoteGw'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemIpiptunnel
	err := ctx.RegisterResource("fortios:index/systemIpiptunnel:SystemIpiptunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemIpiptunnel gets an existing SystemIpiptunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemIpiptunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemIpiptunnelState, opts ...pulumi.ResourceOption) (*SystemIpiptunnel, error) {
	var resource SystemIpiptunnel
	err := ctx.ReadResource("fortios:index/systemIpiptunnel:SystemIpiptunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemIpiptunnel resources.
type systemIpiptunnelState struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Interface name that is associated with the incoming traffic from available options.
	Interface *string `pulumi:"interface"`
	// IPv4 address for the local gateway.
	LocalGw *string `pulumi:"localGw"`
	// IPIP Tunnel name.
	Name *string `pulumi:"name"`
	// IPv4 address for the remote gateway.
	RemoteGw *string `pulumi:"remoteGw"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan *string `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemIpiptunnelState struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Interface name that is associated with the incoming traffic from available options.
	Interface pulumi.StringPtrInput
	// IPv4 address for the local gateway.
	LocalGw pulumi.StringPtrInput
	// IPIP Tunnel name.
	Name pulumi.StringPtrInput
	// IPv4 address for the remote gateway.
	RemoteGw pulumi.StringPtrInput
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpiptunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpiptunnelState)(nil)).Elem()
}

type systemIpiptunnelArgs struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload *string `pulumi:"autoAsicOffload"`
	// Interface name that is associated with the incoming traffic from available options.
	Interface string `pulumi:"interface"`
	// IPv4 address for the local gateway.
	LocalGw string `pulumi:"localGw"`
	// IPIP Tunnel name.
	Name *string `pulumi:"name"`
	// IPv4 address for the remote gateway.
	RemoteGw string `pulumi:"remoteGw"`
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan *string `pulumi:"useSdwan"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemIpiptunnel resource.
type SystemIpiptunnelArgs struct {
	// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
	AutoAsicOffload pulumi.StringPtrInput
	// Interface name that is associated with the incoming traffic from available options.
	Interface pulumi.StringInput
	// IPv4 address for the local gateway.
	LocalGw pulumi.StringInput
	// IPIP Tunnel name.
	Name pulumi.StringPtrInput
	// IPv4 address for the remote gateway.
	RemoteGw pulumi.StringInput
	// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
	UseSdwan pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemIpiptunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemIpiptunnelArgs)(nil)).Elem()
}

type SystemIpiptunnelInput interface {
	pulumi.Input

	ToSystemIpiptunnelOutput() SystemIpiptunnelOutput
	ToSystemIpiptunnelOutputWithContext(ctx context.Context) SystemIpiptunnelOutput
}

func (*SystemIpiptunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpiptunnel)(nil)).Elem()
}

func (i *SystemIpiptunnel) ToSystemIpiptunnelOutput() SystemIpiptunnelOutput {
	return i.ToSystemIpiptunnelOutputWithContext(context.Background())
}

func (i *SystemIpiptunnel) ToSystemIpiptunnelOutputWithContext(ctx context.Context) SystemIpiptunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpiptunnelOutput)
}

// SystemIpiptunnelArrayInput is an input type that accepts SystemIpiptunnelArray and SystemIpiptunnelArrayOutput values.
// You can construct a concrete instance of `SystemIpiptunnelArrayInput` via:
//
//	SystemIpiptunnelArray{ SystemIpiptunnelArgs{...} }
type SystemIpiptunnelArrayInput interface {
	pulumi.Input

	ToSystemIpiptunnelArrayOutput() SystemIpiptunnelArrayOutput
	ToSystemIpiptunnelArrayOutputWithContext(context.Context) SystemIpiptunnelArrayOutput
}

type SystemIpiptunnelArray []SystemIpiptunnelInput

func (SystemIpiptunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpiptunnel)(nil)).Elem()
}

func (i SystemIpiptunnelArray) ToSystemIpiptunnelArrayOutput() SystemIpiptunnelArrayOutput {
	return i.ToSystemIpiptunnelArrayOutputWithContext(context.Background())
}

func (i SystemIpiptunnelArray) ToSystemIpiptunnelArrayOutputWithContext(ctx context.Context) SystemIpiptunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpiptunnelArrayOutput)
}

// SystemIpiptunnelMapInput is an input type that accepts SystemIpiptunnelMap and SystemIpiptunnelMapOutput values.
// You can construct a concrete instance of `SystemIpiptunnelMapInput` via:
//
//	SystemIpiptunnelMap{ "key": SystemIpiptunnelArgs{...} }
type SystemIpiptunnelMapInput interface {
	pulumi.Input

	ToSystemIpiptunnelMapOutput() SystemIpiptunnelMapOutput
	ToSystemIpiptunnelMapOutputWithContext(context.Context) SystemIpiptunnelMapOutput
}

type SystemIpiptunnelMap map[string]SystemIpiptunnelInput

func (SystemIpiptunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpiptunnel)(nil)).Elem()
}

func (i SystemIpiptunnelMap) ToSystemIpiptunnelMapOutput() SystemIpiptunnelMapOutput {
	return i.ToSystemIpiptunnelMapOutputWithContext(context.Background())
}

func (i SystemIpiptunnelMap) ToSystemIpiptunnelMapOutputWithContext(ctx context.Context) SystemIpiptunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemIpiptunnelMapOutput)
}

type SystemIpiptunnelOutput struct{ *pulumi.OutputState }

func (SystemIpiptunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemIpiptunnel)(nil)).Elem()
}

func (o SystemIpiptunnelOutput) ToSystemIpiptunnelOutput() SystemIpiptunnelOutput {
	return o
}

func (o SystemIpiptunnelOutput) ToSystemIpiptunnelOutputWithContext(ctx context.Context) SystemIpiptunnelOutput {
	return o
}

// Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
func (o SystemIpiptunnelOutput) AutoAsicOffload() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringOutput { return v.AutoAsicOffload }).(pulumi.StringOutput)
}

// Interface name that is associated with the incoming traffic from available options.
func (o SystemIpiptunnelOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// IPv4 address for the local gateway.
func (o SystemIpiptunnelOutput) LocalGw() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringOutput { return v.LocalGw }).(pulumi.StringOutput)
}

// IPIP Tunnel name.
func (o SystemIpiptunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IPv4 address for the remote gateway.
func (o SystemIpiptunnelOutput) RemoteGw() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringOutput { return v.RemoteGw }).(pulumi.StringOutput)
}

// Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
func (o SystemIpiptunnelOutput) UseSdwan() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringOutput { return v.UseSdwan }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemIpiptunnelOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemIpiptunnel) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemIpiptunnelArrayOutput struct{ *pulumi.OutputState }

func (SystemIpiptunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemIpiptunnel)(nil)).Elem()
}

func (o SystemIpiptunnelArrayOutput) ToSystemIpiptunnelArrayOutput() SystemIpiptunnelArrayOutput {
	return o
}

func (o SystemIpiptunnelArrayOutput) ToSystemIpiptunnelArrayOutputWithContext(ctx context.Context) SystemIpiptunnelArrayOutput {
	return o
}

func (o SystemIpiptunnelArrayOutput) Index(i pulumi.IntInput) SystemIpiptunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemIpiptunnel {
		return vs[0].([]*SystemIpiptunnel)[vs[1].(int)]
	}).(SystemIpiptunnelOutput)
}

type SystemIpiptunnelMapOutput struct{ *pulumi.OutputState }

func (SystemIpiptunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemIpiptunnel)(nil)).Elem()
}

func (o SystemIpiptunnelMapOutput) ToSystemIpiptunnelMapOutput() SystemIpiptunnelMapOutput {
	return o
}

func (o SystemIpiptunnelMapOutput) ToSystemIpiptunnelMapOutputWithContext(ctx context.Context) SystemIpiptunnelMapOutput {
	return o
}

func (o SystemIpiptunnelMapOutput) MapIndex(k pulumi.StringInput) SystemIpiptunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemIpiptunnel {
		return vs[0].(map[string]*SystemIpiptunnel)[vs[1].(string)]
	}).(SystemIpiptunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpiptunnelInput)(nil)).Elem(), &SystemIpiptunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpiptunnelArrayInput)(nil)).Elem(), SystemIpiptunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemIpiptunnelMapInput)(nil)).Elem(), SystemIpiptunnelMap{})
	pulumi.RegisterOutputType(SystemIpiptunnelOutput{})
	pulumi.RegisterOutputType(SystemIpiptunnelArrayOutput{})
	pulumi.RegisterOutputType(SystemIpiptunnelMapOutput{})
}
