// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.
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
//			_, err := fortios.NewSwitchcontrollerCustomcommand(ctx, "trname", &fortios.SwitchcontrollerCustomcommandArgs{
//				Command:     pulumi.String("ls"),
//				CommandName: pulumi.String("1"),
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
// # SwitchController CustomCommand can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerCustomcommand:SwitchcontrollerCustomcommand labelname {{command_name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerCustomcommand:SwitchcontrollerCustomcommand labelname {{command_name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerCustomcommand struct {
	pulumi.CustomResourceState

	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command pulumi.StringOutput `pulumi:"command"`
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName pulumi.StringOutput `pulumi:"commandName"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerCustomcommand registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerCustomcommand(ctx *pulumi.Context,
	name string, args *SwitchcontrollerCustomcommandArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerCustomcommand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Command == nil {
		return nil, errors.New("invalid value for required argument 'Command'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerCustomcommand
	err := ctx.RegisterResource("fortios:index/switchcontrollerCustomcommand:SwitchcontrollerCustomcommand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerCustomcommand gets an existing SwitchcontrollerCustomcommand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerCustomcommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerCustomcommandState, opts ...pulumi.ResourceOption) (*SwitchcontrollerCustomcommand, error) {
	var resource SwitchcontrollerCustomcommand
	err := ctx.ReadResource("fortios:index/switchcontrollerCustomcommand:SwitchcontrollerCustomcommand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerCustomcommand resources.
type switchcontrollerCustomcommandState struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command *string `pulumi:"command"`
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName *string `pulumi:"commandName"`
	// Description.
	Description *string `pulumi:"description"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerCustomcommandState struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command pulumi.StringPtrInput
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerCustomcommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerCustomcommandState)(nil)).Elem()
}

type switchcontrollerCustomcommandArgs struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command string `pulumi:"command"`
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName *string `pulumi:"commandName"`
	// Description.
	Description *string `pulumi:"description"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerCustomcommand resource.
type SwitchcontrollerCustomcommandArgs struct {
	// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
	Command pulumi.StringInput
	// Command name called by the FortiGate switch controller in the execute command.
	CommandName pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerCustomcommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerCustomcommandArgs)(nil)).Elem()
}

type SwitchcontrollerCustomcommandInput interface {
	pulumi.Input

	ToSwitchcontrollerCustomcommandOutput() SwitchcontrollerCustomcommandOutput
	ToSwitchcontrollerCustomcommandOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandOutput
}

func (*SwitchcontrollerCustomcommand) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerCustomcommand)(nil)).Elem()
}

func (i *SwitchcontrollerCustomcommand) ToSwitchcontrollerCustomcommandOutput() SwitchcontrollerCustomcommandOutput {
	return i.ToSwitchcontrollerCustomcommandOutputWithContext(context.Background())
}

func (i *SwitchcontrollerCustomcommand) ToSwitchcontrollerCustomcommandOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerCustomcommandOutput)
}

// SwitchcontrollerCustomcommandArrayInput is an input type that accepts SwitchcontrollerCustomcommandArray and SwitchcontrollerCustomcommandArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerCustomcommandArrayInput` via:
//
//	SwitchcontrollerCustomcommandArray{ SwitchcontrollerCustomcommandArgs{...} }
type SwitchcontrollerCustomcommandArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerCustomcommandArrayOutput() SwitchcontrollerCustomcommandArrayOutput
	ToSwitchcontrollerCustomcommandArrayOutputWithContext(context.Context) SwitchcontrollerCustomcommandArrayOutput
}

type SwitchcontrollerCustomcommandArray []SwitchcontrollerCustomcommandInput

func (SwitchcontrollerCustomcommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerCustomcommand)(nil)).Elem()
}

func (i SwitchcontrollerCustomcommandArray) ToSwitchcontrollerCustomcommandArrayOutput() SwitchcontrollerCustomcommandArrayOutput {
	return i.ToSwitchcontrollerCustomcommandArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerCustomcommandArray) ToSwitchcontrollerCustomcommandArrayOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerCustomcommandArrayOutput)
}

// SwitchcontrollerCustomcommandMapInput is an input type that accepts SwitchcontrollerCustomcommandMap and SwitchcontrollerCustomcommandMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerCustomcommandMapInput` via:
//
//	SwitchcontrollerCustomcommandMap{ "key": SwitchcontrollerCustomcommandArgs{...} }
type SwitchcontrollerCustomcommandMapInput interface {
	pulumi.Input

	ToSwitchcontrollerCustomcommandMapOutput() SwitchcontrollerCustomcommandMapOutput
	ToSwitchcontrollerCustomcommandMapOutputWithContext(context.Context) SwitchcontrollerCustomcommandMapOutput
}

type SwitchcontrollerCustomcommandMap map[string]SwitchcontrollerCustomcommandInput

func (SwitchcontrollerCustomcommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerCustomcommand)(nil)).Elem()
}

func (i SwitchcontrollerCustomcommandMap) ToSwitchcontrollerCustomcommandMapOutput() SwitchcontrollerCustomcommandMapOutput {
	return i.ToSwitchcontrollerCustomcommandMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerCustomcommandMap) ToSwitchcontrollerCustomcommandMapOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerCustomcommandMapOutput)
}

type SwitchcontrollerCustomcommandOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerCustomcommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerCustomcommand)(nil)).Elem()
}

func (o SwitchcontrollerCustomcommandOutput) ToSwitchcontrollerCustomcommandOutput() SwitchcontrollerCustomcommandOutput {
	return o
}

func (o SwitchcontrollerCustomcommandOutput) ToSwitchcontrollerCustomcommandOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandOutput {
	return o
}

// String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).
func (o SwitchcontrollerCustomcommandOutput) Command() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerCustomcommand) pulumi.StringOutput { return v.Command }).(pulumi.StringOutput)
}

// Command name called by the FortiGate switch controller in the execute command.
func (o SwitchcontrollerCustomcommandOutput) CommandName() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerCustomcommand) pulumi.StringOutput { return v.CommandName }).(pulumi.StringOutput)
}

// Description.
func (o SwitchcontrollerCustomcommandOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerCustomcommand) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerCustomcommandOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerCustomcommand) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerCustomcommandArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerCustomcommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerCustomcommand)(nil)).Elem()
}

func (o SwitchcontrollerCustomcommandArrayOutput) ToSwitchcontrollerCustomcommandArrayOutput() SwitchcontrollerCustomcommandArrayOutput {
	return o
}

func (o SwitchcontrollerCustomcommandArrayOutput) ToSwitchcontrollerCustomcommandArrayOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandArrayOutput {
	return o
}

func (o SwitchcontrollerCustomcommandArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerCustomcommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerCustomcommand {
		return vs[0].([]*SwitchcontrollerCustomcommand)[vs[1].(int)]
	}).(SwitchcontrollerCustomcommandOutput)
}

type SwitchcontrollerCustomcommandMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerCustomcommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerCustomcommand)(nil)).Elem()
}

func (o SwitchcontrollerCustomcommandMapOutput) ToSwitchcontrollerCustomcommandMapOutput() SwitchcontrollerCustomcommandMapOutput {
	return o
}

func (o SwitchcontrollerCustomcommandMapOutput) ToSwitchcontrollerCustomcommandMapOutputWithContext(ctx context.Context) SwitchcontrollerCustomcommandMapOutput {
	return o
}

func (o SwitchcontrollerCustomcommandMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerCustomcommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerCustomcommand {
		return vs[0].(map[string]*SwitchcontrollerCustomcommand)[vs[1].(string)]
	}).(SwitchcontrollerCustomcommandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerCustomcommandInput)(nil)).Elem(), &SwitchcontrollerCustomcommand{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerCustomcommandArrayInput)(nil)).Elem(), SwitchcontrollerCustomcommandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerCustomcommandMapInput)(nil)).Elem(), SwitchcontrollerCustomcommandMap{})
	pulumi.RegisterOutputType(SwitchcontrollerCustomcommandOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerCustomcommandArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerCustomcommandMapOutput{})
}