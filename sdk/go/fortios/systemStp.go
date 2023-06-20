// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Spanning Tree Protocol (STP). Applies to FortiOS Version `7.0.4`.
//
// ## Import
//
// # System Stp can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemStp:SystemStp labelname SystemStp
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemStp:SystemStp labelname SystemStp
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemStp struct {
	pulumi.CustomResourceState

	// Forward delay (4 - 30 sec, default = 15).
	ForwardDelay pulumi.IntOutput `pulumi:"forwardDelay"`
	// Hello time (1 - 10 sec, default = 2).
	HelloTime pulumi.IntOutput `pulumi:"helloTime"`
	// Maximum packet age (6 - 40 sec, default = 20).
	MaxAge pulumi.IntOutput `pulumi:"maxAge"`
	// Maximum number of hops (1 - 40, default = 20).
	MaxHops pulumi.IntOutput `pulumi:"maxHops"`
	// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
	SwitchPriority pulumi.StringOutput `pulumi:"switchPriority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemStp registers a new resource with the given unique name, arguments, and options.
func NewSystemStp(ctx *pulumi.Context,
	name string, args *SystemStpArgs, opts ...pulumi.ResourceOption) (*SystemStp, error) {
	if args == nil {
		args = &SystemStpArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemStp
	err := ctx.RegisterResource("fortios:index/systemStp:SystemStp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemStp gets an existing SystemStp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemStp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemStpState, opts ...pulumi.ResourceOption) (*SystemStp, error) {
	var resource SystemStp
	err := ctx.ReadResource("fortios:index/systemStp:SystemStp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemStp resources.
type systemStpState struct {
	// Forward delay (4 - 30 sec, default = 15).
	ForwardDelay *int `pulumi:"forwardDelay"`
	// Hello time (1 - 10 sec, default = 2).
	HelloTime *int `pulumi:"helloTime"`
	// Maximum packet age (6 - 40 sec, default = 20).
	MaxAge *int `pulumi:"maxAge"`
	// Maximum number of hops (1 - 40, default = 20).
	MaxHops *int `pulumi:"maxHops"`
	// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
	SwitchPriority *string `pulumi:"switchPriority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemStpState struct {
	// Forward delay (4 - 30 sec, default = 15).
	ForwardDelay pulumi.IntPtrInput
	// Hello time (1 - 10 sec, default = 2).
	HelloTime pulumi.IntPtrInput
	// Maximum packet age (6 - 40 sec, default = 20).
	MaxAge pulumi.IntPtrInput
	// Maximum number of hops (1 - 40, default = 20).
	MaxHops pulumi.IntPtrInput
	// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
	SwitchPriority pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemStpState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStpState)(nil)).Elem()
}

type systemStpArgs struct {
	// Forward delay (4 - 30 sec, default = 15).
	ForwardDelay *int `pulumi:"forwardDelay"`
	// Hello time (1 - 10 sec, default = 2).
	HelloTime *int `pulumi:"helloTime"`
	// Maximum packet age (6 - 40 sec, default = 20).
	MaxAge *int `pulumi:"maxAge"`
	// Maximum number of hops (1 - 40, default = 20).
	MaxHops *int `pulumi:"maxHops"`
	// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
	SwitchPriority *string `pulumi:"switchPriority"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemStp resource.
type SystemStpArgs struct {
	// Forward delay (4 - 30 sec, default = 15).
	ForwardDelay pulumi.IntPtrInput
	// Hello time (1 - 10 sec, default = 2).
	HelloTime pulumi.IntPtrInput
	// Maximum packet age (6 - 40 sec, default = 20).
	MaxAge pulumi.IntPtrInput
	// Maximum number of hops (1 - 40, default = 20).
	MaxHops pulumi.IntPtrInput
	// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
	SwitchPriority pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemStpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemStpArgs)(nil)).Elem()
}

type SystemStpInput interface {
	pulumi.Input

	ToSystemStpOutput() SystemStpOutput
	ToSystemStpOutputWithContext(ctx context.Context) SystemStpOutput
}

func (*SystemStp) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStp)(nil)).Elem()
}

func (i *SystemStp) ToSystemStpOutput() SystemStpOutput {
	return i.ToSystemStpOutputWithContext(context.Background())
}

func (i *SystemStp) ToSystemStpOutputWithContext(ctx context.Context) SystemStpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStpOutput)
}

// SystemStpArrayInput is an input type that accepts SystemStpArray and SystemStpArrayOutput values.
// You can construct a concrete instance of `SystemStpArrayInput` via:
//
//	SystemStpArray{ SystemStpArgs{...} }
type SystemStpArrayInput interface {
	pulumi.Input

	ToSystemStpArrayOutput() SystemStpArrayOutput
	ToSystemStpArrayOutputWithContext(context.Context) SystemStpArrayOutput
}

type SystemStpArray []SystemStpInput

func (SystemStpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStp)(nil)).Elem()
}

func (i SystemStpArray) ToSystemStpArrayOutput() SystemStpArrayOutput {
	return i.ToSystemStpArrayOutputWithContext(context.Background())
}

func (i SystemStpArray) ToSystemStpArrayOutputWithContext(ctx context.Context) SystemStpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStpArrayOutput)
}

// SystemStpMapInput is an input type that accepts SystemStpMap and SystemStpMapOutput values.
// You can construct a concrete instance of `SystemStpMapInput` via:
//
//	SystemStpMap{ "key": SystemStpArgs{...} }
type SystemStpMapInput interface {
	pulumi.Input

	ToSystemStpMapOutput() SystemStpMapOutput
	ToSystemStpMapOutputWithContext(context.Context) SystemStpMapOutput
}

type SystemStpMap map[string]SystemStpInput

func (SystemStpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStp)(nil)).Elem()
}

func (i SystemStpMap) ToSystemStpMapOutput() SystemStpMapOutput {
	return i.ToSystemStpMapOutputWithContext(context.Background())
}

func (i SystemStpMap) ToSystemStpMapOutputWithContext(ctx context.Context) SystemStpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemStpMapOutput)
}

type SystemStpOutput struct{ *pulumi.OutputState }

func (SystemStpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemStp)(nil)).Elem()
}

func (o SystemStpOutput) ToSystemStpOutput() SystemStpOutput {
	return o
}

func (o SystemStpOutput) ToSystemStpOutputWithContext(ctx context.Context) SystemStpOutput {
	return o
}

// Forward delay (4 - 30 sec, default = 15).
func (o SystemStpOutput) ForwardDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStp) pulumi.IntOutput { return v.ForwardDelay }).(pulumi.IntOutput)
}

// Hello time (1 - 10 sec, default = 2).
func (o SystemStpOutput) HelloTime() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStp) pulumi.IntOutput { return v.HelloTime }).(pulumi.IntOutput)
}

// Maximum packet age (6 - 40 sec, default = 20).
func (o SystemStpOutput) MaxAge() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStp) pulumi.IntOutput { return v.MaxAge }).(pulumi.IntOutput)
}

// Maximum number of hops (1 - 40, default = 20).
func (o SystemStpOutput) MaxHops() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemStp) pulumi.IntOutput { return v.MaxHops }).(pulumi.IntOutput)
}

// STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344). Valid values: `0`, `4096`, `8192`, `12288`, `16384`, `20480`, `24576`, `28672`, `32768`, `36864`, `40960`, `45056`, `49152`, `53248`, `57344`.
func (o SystemStpOutput) SwitchPriority() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemStp) pulumi.StringOutput { return v.SwitchPriority }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemStpOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemStp) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemStpArrayOutput struct{ *pulumi.OutputState }

func (SystemStpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemStp)(nil)).Elem()
}

func (o SystemStpArrayOutput) ToSystemStpArrayOutput() SystemStpArrayOutput {
	return o
}

func (o SystemStpArrayOutput) ToSystemStpArrayOutputWithContext(ctx context.Context) SystemStpArrayOutput {
	return o
}

func (o SystemStpArrayOutput) Index(i pulumi.IntInput) SystemStpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemStp {
		return vs[0].([]*SystemStp)[vs[1].(int)]
	}).(SystemStpOutput)
}

type SystemStpMapOutput struct{ *pulumi.OutputState }

func (SystemStpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemStp)(nil)).Elem()
}

func (o SystemStpMapOutput) ToSystemStpMapOutput() SystemStpMapOutput {
	return o
}

func (o SystemStpMapOutput) ToSystemStpMapOutputWithContext(ctx context.Context) SystemStpMapOutput {
	return o
}

func (o SystemStpMapOutput) MapIndex(k pulumi.StringInput) SystemStpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemStp {
		return vs[0].(map[string]*SystemStp)[vs[1].(string)]
	}).(SystemStpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStpInput)(nil)).Elem(), &SystemStp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStpArrayInput)(nil)).Elem(), SystemStpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemStpMapInput)(nil)).Elem(), SystemStpMap{})
	pulumi.RegisterOutputType(SystemStpOutput{})
	pulumi.RegisterOutputType(SystemStpArrayOutput{})
	pulumi.RegisterOutputType(SystemStpMapOutput{})
}
