// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure sFlow per VDOM to add or change the IP address and UDP port that FortiGate sFlow agents in this VDOM use to send sFlow datagrams to an sFlow collector.
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
//			_, err := fortios.NewSystemVdomsflow(ctx, "trname", &fortios.SystemVdomsflowArgs{
//				CollectorIp:   pulumi.String("0.0.0.0"),
//				CollectorPort: pulumi.Int(6343),
//				SourceIp:      pulumi.String("0.0.0.0"),
//				VdomSflow:     pulumi.String("disable"),
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
// # System VdomSflow can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemVdomsflow:SystemVdomsflow labelname SystemVdomSflow
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemVdomsflow:SystemVdomsflow labelname SystemVdomSflow
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemVdomsflow struct {
	pulumi.CustomResourceState

	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp pulumi.StringOutput `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort pulumi.IntOutput `pulumi:"collectorPort"`
	// Specify outgoing interface to reach server.
	Interface pulumi.StringOutput `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringOutput `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable the sFlow configuration for the current VDOM. Valid values: `enable`, `disable`.
	VdomSflow pulumi.StringOutput `pulumi:"vdomSflow"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomsflow registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomsflow(ctx *pulumi.Context,
	name string, args *SystemVdomsflowArgs, opts ...pulumi.ResourceOption) (*SystemVdomsflow, error) {
	if args == nil {
		args = &SystemVdomsflowArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVdomsflow
	err := ctx.RegisterResource("fortios:index/systemVdomsflow:SystemVdomsflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomsflow gets an existing SystemVdomsflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomsflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdomsflowState, opts ...pulumi.ResourceOption) (*SystemVdomsflow, error) {
	var resource SystemVdomsflow
	err := ctx.ReadResource("fortios:index/systemVdomsflow:SystemVdomsflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomsflow resources.
type systemVdomsflowState struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp *string `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort *int `pulumi:"collectorPort"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable the sFlow configuration for the current VDOM. Valid values: `enable`, `disable`.
	VdomSflow *string `pulumi:"vdomSflow"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdomsflowState struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp pulumi.StringPtrInput
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Source IP address for sFlow agent.
	SourceIp pulumi.StringPtrInput
	// Enable/disable the sFlow configuration for the current VDOM. Valid values: `enable`, `disable`.
	VdomSflow pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomsflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomsflowState)(nil)).Elem()
}

type systemVdomsflowArgs struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp *string `pulumi:"collectorIp"`
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort *int `pulumi:"collectorPort"`
	// Specify outgoing interface to reach server.
	Interface *string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod *string `pulumi:"interfaceSelectMethod"`
	// Source IP address for sFlow agent.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable the sFlow configuration for the current VDOM. Valid values: `enable`, `disable`.
	VdomSflow *string `pulumi:"vdomSflow"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomsflow resource.
type SystemVdomsflowArgs struct {
	// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
	CollectorIp pulumi.StringPtrInput
	// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
	CollectorPort pulumi.IntPtrInput
	// Specify outgoing interface to reach server.
	Interface pulumi.StringPtrInput
	// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
	InterfaceSelectMethod pulumi.StringPtrInput
	// Source IP address for sFlow agent.
	SourceIp pulumi.StringPtrInput
	// Enable/disable the sFlow configuration for the current VDOM. Valid values: `enable`, `disable`.
	VdomSflow pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdomsflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdomsflowArgs)(nil)).Elem()
}

type SystemVdomsflowInput interface {
	pulumi.Input

	ToSystemVdomsflowOutput() SystemVdomsflowOutput
	ToSystemVdomsflowOutputWithContext(ctx context.Context) SystemVdomsflowOutput
}

func (*SystemVdomsflow) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomsflow)(nil)).Elem()
}

func (i *SystemVdomsflow) ToSystemVdomsflowOutput() SystemVdomsflowOutput {
	return i.ToSystemVdomsflowOutputWithContext(context.Background())
}

func (i *SystemVdomsflow) ToSystemVdomsflowOutputWithContext(ctx context.Context) SystemVdomsflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomsflowOutput)
}

// SystemVdomsflowArrayInput is an input type that accepts SystemVdomsflowArray and SystemVdomsflowArrayOutput values.
// You can construct a concrete instance of `SystemVdomsflowArrayInput` via:
//
//	SystemVdomsflowArray{ SystemVdomsflowArgs{...} }
type SystemVdomsflowArrayInput interface {
	pulumi.Input

	ToSystemVdomsflowArrayOutput() SystemVdomsflowArrayOutput
	ToSystemVdomsflowArrayOutputWithContext(context.Context) SystemVdomsflowArrayOutput
}

type SystemVdomsflowArray []SystemVdomsflowInput

func (SystemVdomsflowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomsflow)(nil)).Elem()
}

func (i SystemVdomsflowArray) ToSystemVdomsflowArrayOutput() SystemVdomsflowArrayOutput {
	return i.ToSystemVdomsflowArrayOutputWithContext(context.Background())
}

func (i SystemVdomsflowArray) ToSystemVdomsflowArrayOutputWithContext(ctx context.Context) SystemVdomsflowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomsflowArrayOutput)
}

// SystemVdomsflowMapInput is an input type that accepts SystemVdomsflowMap and SystemVdomsflowMapOutput values.
// You can construct a concrete instance of `SystemVdomsflowMapInput` via:
//
//	SystemVdomsflowMap{ "key": SystemVdomsflowArgs{...} }
type SystemVdomsflowMapInput interface {
	pulumi.Input

	ToSystemVdomsflowMapOutput() SystemVdomsflowMapOutput
	ToSystemVdomsflowMapOutputWithContext(context.Context) SystemVdomsflowMapOutput
}

type SystemVdomsflowMap map[string]SystemVdomsflowInput

func (SystemVdomsflowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomsflow)(nil)).Elem()
}

func (i SystemVdomsflowMap) ToSystemVdomsflowMapOutput() SystemVdomsflowMapOutput {
	return i.ToSystemVdomsflowMapOutputWithContext(context.Background())
}

func (i SystemVdomsflowMap) ToSystemVdomsflowMapOutputWithContext(ctx context.Context) SystemVdomsflowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdomsflowMapOutput)
}

type SystemVdomsflowOutput struct{ *pulumi.OutputState }

func (SystemVdomsflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomsflow)(nil)).Elem()
}

func (o SystemVdomsflowOutput) ToSystemVdomsflowOutput() SystemVdomsflowOutput {
	return o
}

func (o SystemVdomsflowOutput) ToSystemVdomsflowOutputWithContext(ctx context.Context) SystemVdomsflowOutput {
	return o
}

// IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
func (o SystemVdomsflowOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.StringOutput { return v.CollectorIp }).(pulumi.StringOutput)
}

// UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
func (o SystemVdomsflowOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.IntOutput { return v.CollectorPort }).(pulumi.IntOutput)
}

// Specify outgoing interface to reach server.
func (o SystemVdomsflowOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.StringOutput { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
func (o SystemVdomsflowOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.StringOutput { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Source IP address for sFlow agent.
func (o SystemVdomsflowOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable the sFlow configuration for the current VDOM. Valid values: `enable`, `disable`.
func (o SystemVdomsflowOutput) VdomSflow() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.StringOutput { return v.VdomSflow }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemVdomsflowOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVdomsflow) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemVdomsflowArrayOutput struct{ *pulumi.OutputState }

func (SystemVdomsflowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomsflow)(nil)).Elem()
}

func (o SystemVdomsflowArrayOutput) ToSystemVdomsflowArrayOutput() SystemVdomsflowArrayOutput {
	return o
}

func (o SystemVdomsflowArrayOutput) ToSystemVdomsflowArrayOutputWithContext(ctx context.Context) SystemVdomsflowArrayOutput {
	return o
}

func (o SystemVdomsflowArrayOutput) Index(i pulumi.IntInput) SystemVdomsflowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomsflow {
		return vs[0].([]*SystemVdomsflow)[vs[1].(int)]
	}).(SystemVdomsflowOutput)
}

type SystemVdomsflowMapOutput struct{ *pulumi.OutputState }

func (SystemVdomsflowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomsflow)(nil)).Elem()
}

func (o SystemVdomsflowMapOutput) ToSystemVdomsflowMapOutput() SystemVdomsflowMapOutput {
	return o
}

func (o SystemVdomsflowMapOutput) ToSystemVdomsflowMapOutputWithContext(ctx context.Context) SystemVdomsflowMapOutput {
	return o
}

func (o SystemVdomsflowMapOutput) MapIndex(k pulumi.StringInput) SystemVdomsflowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomsflow {
		return vs[0].(map[string]*SystemVdomsflow)[vs[1].(string)]
	}).(SystemVdomsflowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomsflowInput)(nil)).Elem(), &SystemVdomsflow{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomsflowArrayInput)(nil)).Elem(), SystemVdomsflowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdomsflowMapInput)(nil)).Elem(), SystemVdomsflowMap{})
	pulumi.RegisterOutputType(SystemVdomsflowOutput{})
	pulumi.RegisterOutputType(SystemVdomsflowArrayOutput{})
	pulumi.RegisterOutputType(SystemVdomsflowMapOutput{})
}
