// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure network authentication type.
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
//			_, err := fortios.NewWirelesscontrollerhotspot20Anqpnetworkauthtype(ctx, "trname", &fortios.Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs{
//				AuthType: pulumi.String("acceptance-of-terms"),
//				Url:      pulumi.String("www.example.com"),
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
// # WirelessControllerHotspot20 AnqpNetworkAuthType can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpnetworkauthtype:Wirelesscontrollerhotspot20Anqpnetworkauthtype labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/wirelesscontrollerhotspot20Anqpnetworkauthtype:Wirelesscontrollerhotspot20Anqpnetworkauthtype labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type Wirelesscontrollerhotspot20Anqpnetworkauthtype struct {
	pulumi.CustomResourceState

	// Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// Authentication type name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Redirect URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWirelesscontrollerhotspot20Anqpnetworkauthtype registers a new resource with the given unique name, arguments, and options.
func NewWirelesscontrollerhotspot20Anqpnetworkauthtype(ctx *pulumi.Context,
	name string, args *Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20Anqpnetworkauthtype, error) {
	if args == nil {
		args = &Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Wirelesscontrollerhotspot20Anqpnetworkauthtype
	err := ctx.RegisterResource("fortios:index/wirelesscontrollerhotspot20Anqpnetworkauthtype:Wirelesscontrollerhotspot20Anqpnetworkauthtype", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWirelesscontrollerhotspot20Anqpnetworkauthtype gets an existing Wirelesscontrollerhotspot20Anqpnetworkauthtype resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWirelesscontrollerhotspot20Anqpnetworkauthtype(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Wirelesscontrollerhotspot20AnqpnetworkauthtypeState, opts ...pulumi.ResourceOption) (*Wirelesscontrollerhotspot20Anqpnetworkauthtype, error) {
	var resource Wirelesscontrollerhotspot20Anqpnetworkauthtype
	err := ctx.ReadResource("fortios:index/wirelesscontrollerhotspot20Anqpnetworkauthtype:Wirelesscontrollerhotspot20Anqpnetworkauthtype", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Wirelesscontrollerhotspot20Anqpnetworkauthtype resources.
type wirelesscontrollerhotspot20AnqpnetworkauthtypeState struct {
	// Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
	AuthType *string `pulumi:"authType"`
	// Authentication type name.
	Name *string `pulumi:"name"`
	// Redirect URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeState struct {
	// Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
	AuthType pulumi.StringPtrInput
	// Authentication type name.
	Name pulumi.StringPtrInput
	// Redirect URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20AnqpnetworkauthtypeState)(nil)).Elem()
}

type wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs struct {
	// Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
	AuthType *string `pulumi:"authType"`
	// Authentication type name.
	Name *string `pulumi:"name"`
	// Redirect URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a Wirelesscontrollerhotspot20Anqpnetworkauthtype resource.
type Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs struct {
	// Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
	AuthType pulumi.StringPtrInput
	// Authentication type name.
	Name pulumi.StringPtrInput
	// Redirect URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs)(nil)).Elem()
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput
	ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput
}

func (*Wirelesscontrollerhotspot20Anqpnetworkauthtype) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20Anqpnetworkauthtype)(nil)).Elem()
}

func (i *Wirelesscontrollerhotspot20Anqpnetworkauthtype) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput {
	return i.ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutputWithContext(context.Background())
}

func (i *Wirelesscontrollerhotspot20Anqpnetworkauthtype) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput)
}

// Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayInput is an input type that accepts Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray and Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayInput` via:
//
//	Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray{ Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs{...} }
type Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput
	ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutputWithContext(context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray []Wirelesscontrollerhotspot20AnqpnetworkauthtypeInput

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20Anqpnetworkauthtype)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput {
	return i.ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput)
}

// Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapInput is an input type that accepts Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap and Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput values.
// You can construct a concrete instance of `Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapInput` via:
//
//	Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap{ "key": Wirelesscontrollerhotspot20AnqpnetworkauthtypeArgs{...} }
type Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapInput interface {
	pulumi.Input

	ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput
	ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutputWithContext(context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap map[string]Wirelesscontrollerhotspot20AnqpnetworkauthtypeInput

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20Anqpnetworkauthtype)(nil)).Elem()
}

func (i Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput {
	return i.ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutputWithContext(context.Background())
}

func (i Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput)
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Wirelesscontrollerhotspot20Anqpnetworkauthtype)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput {
	return o
}

// Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnetworkauthtype) pulumi.StringOutput { return v.AuthType }).(pulumi.StringOutput)
}

// Authentication type name.
func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnetworkauthtype) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Redirect URL.
func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnetworkauthtype) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Wirelesscontrollerhotspot20Anqpnetworkauthtype) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Wirelesscontrollerhotspot20Anqpnetworkauthtype)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput) Index(i pulumi.IntInput) Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20Anqpnetworkauthtype {
		return vs[0].([]*Wirelesscontrollerhotspot20Anqpnetworkauthtype)[vs[1].(int)]
	}).(Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput)
}

type Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput struct{ *pulumi.OutputState }

func (Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Wirelesscontrollerhotspot20Anqpnetworkauthtype)(nil)).Elem()
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput() Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput) ToWirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutputWithContext(ctx context.Context) Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput {
	return o
}

func (o Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput) MapIndex(k pulumi.StringInput) Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Wirelesscontrollerhotspot20Anqpnetworkauthtype {
		return vs[0].(map[string]*Wirelesscontrollerhotspot20Anqpnetworkauthtype)[vs[1].(string)]
	}).(Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpnetworkauthtypeInput)(nil)).Elem(), &Wirelesscontrollerhotspot20Anqpnetworkauthtype{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayInput)(nil)).Elem(), Wirelesscontrollerhotspot20AnqpnetworkauthtypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapInput)(nil)).Elem(), Wirelesscontrollerhotspot20AnqpnetworkauthtypeMap{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpnetworkauthtypeOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpnetworkauthtypeArrayOutput{})
	pulumi.RegisterOutputType(Wirelesscontrollerhotspot20AnqpnetworkauthtypeMapOutput{})
}
