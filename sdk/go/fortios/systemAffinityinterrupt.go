// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure interrupt affinity.
//
// ## Import
//
// # System AffinityInterrupt can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemAffinityinterrupt:SystemAffinityinterrupt labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemAffinityinterrupt:SystemAffinityinterrupt labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemAffinityinterrupt struct {
	pulumi.CustomResourceState

	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringOutput `pulumi:"affinityCpumask"`
	// ID of the interrupt affinity setting.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Interrupt name.
	Interrupt pulumi.StringOutput `pulumi:"interrupt"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemAffinityinterrupt registers a new resource with the given unique name, arguments, and options.
func NewSystemAffinityinterrupt(ctx *pulumi.Context,
	name string, args *SystemAffinityinterruptArgs, opts ...pulumi.ResourceOption) (*SystemAffinityinterrupt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AffinityCpumask == nil {
		return nil, errors.New("invalid value for required argument 'AffinityCpumask'")
	}
	if args.Fosid == nil {
		return nil, errors.New("invalid value for required argument 'Fosid'")
	}
	if args.Interrupt == nil {
		return nil, errors.New("invalid value for required argument 'Interrupt'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemAffinityinterrupt
	err := ctx.RegisterResource("fortios:index/systemAffinityinterrupt:SystemAffinityinterrupt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemAffinityinterrupt gets an existing SystemAffinityinterrupt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemAffinityinterrupt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemAffinityinterruptState, opts ...pulumi.ResourceOption) (*SystemAffinityinterrupt, error) {
	var resource SystemAffinityinterrupt
	err := ctx.ReadResource("fortios:index/systemAffinityinterrupt:SystemAffinityinterrupt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemAffinityinterrupt resources.
type systemAffinityinterruptState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask *string `pulumi:"affinityCpumask"`
	// ID of the interrupt affinity setting.
	Fosid *int `pulumi:"fosid"`
	// Interrupt name.
	Interrupt *string `pulumi:"interrupt"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemAffinityinterruptState struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringPtrInput
	// ID of the interrupt affinity setting.
	Fosid pulumi.IntPtrInput
	// Interrupt name.
	Interrupt pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAffinityinterruptState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAffinityinterruptState)(nil)).Elem()
}

type systemAffinityinterruptArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask string `pulumi:"affinityCpumask"`
	// ID of the interrupt affinity setting.
	Fosid int `pulumi:"fosid"`
	// Interrupt name.
	Interrupt string `pulumi:"interrupt"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemAffinityinterrupt resource.
type SystemAffinityinterruptArgs struct {
	// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
	AffinityCpumask pulumi.StringInput
	// ID of the interrupt affinity setting.
	Fosid pulumi.IntInput
	// Interrupt name.
	Interrupt pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemAffinityinterruptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemAffinityinterruptArgs)(nil)).Elem()
}

type SystemAffinityinterruptInput interface {
	pulumi.Input

	ToSystemAffinityinterruptOutput() SystemAffinityinterruptOutput
	ToSystemAffinityinterruptOutputWithContext(ctx context.Context) SystemAffinityinterruptOutput
}

func (*SystemAffinityinterrupt) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAffinityinterrupt)(nil)).Elem()
}

func (i *SystemAffinityinterrupt) ToSystemAffinityinterruptOutput() SystemAffinityinterruptOutput {
	return i.ToSystemAffinityinterruptOutputWithContext(context.Background())
}

func (i *SystemAffinityinterrupt) ToSystemAffinityinterruptOutputWithContext(ctx context.Context) SystemAffinityinterruptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityinterruptOutput)
}

// SystemAffinityinterruptArrayInput is an input type that accepts SystemAffinityinterruptArray and SystemAffinityinterruptArrayOutput values.
// You can construct a concrete instance of `SystemAffinityinterruptArrayInput` via:
//
//	SystemAffinityinterruptArray{ SystemAffinityinterruptArgs{...} }
type SystemAffinityinterruptArrayInput interface {
	pulumi.Input

	ToSystemAffinityinterruptArrayOutput() SystemAffinityinterruptArrayOutput
	ToSystemAffinityinterruptArrayOutputWithContext(context.Context) SystemAffinityinterruptArrayOutput
}

type SystemAffinityinterruptArray []SystemAffinityinterruptInput

func (SystemAffinityinterruptArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAffinityinterrupt)(nil)).Elem()
}

func (i SystemAffinityinterruptArray) ToSystemAffinityinterruptArrayOutput() SystemAffinityinterruptArrayOutput {
	return i.ToSystemAffinityinterruptArrayOutputWithContext(context.Background())
}

func (i SystemAffinityinterruptArray) ToSystemAffinityinterruptArrayOutputWithContext(ctx context.Context) SystemAffinityinterruptArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityinterruptArrayOutput)
}

// SystemAffinityinterruptMapInput is an input type that accepts SystemAffinityinterruptMap and SystemAffinityinterruptMapOutput values.
// You can construct a concrete instance of `SystemAffinityinterruptMapInput` via:
//
//	SystemAffinityinterruptMap{ "key": SystemAffinityinterruptArgs{...} }
type SystemAffinityinterruptMapInput interface {
	pulumi.Input

	ToSystemAffinityinterruptMapOutput() SystemAffinityinterruptMapOutput
	ToSystemAffinityinterruptMapOutputWithContext(context.Context) SystemAffinityinterruptMapOutput
}

type SystemAffinityinterruptMap map[string]SystemAffinityinterruptInput

func (SystemAffinityinterruptMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAffinityinterrupt)(nil)).Elem()
}

func (i SystemAffinityinterruptMap) ToSystemAffinityinterruptMapOutput() SystemAffinityinterruptMapOutput {
	return i.ToSystemAffinityinterruptMapOutputWithContext(context.Background())
}

func (i SystemAffinityinterruptMap) ToSystemAffinityinterruptMapOutputWithContext(ctx context.Context) SystemAffinityinterruptMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAffinityinterruptMapOutput)
}

type SystemAffinityinterruptOutput struct{ *pulumi.OutputState }

func (SystemAffinityinterruptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAffinityinterrupt)(nil)).Elem()
}

func (o SystemAffinityinterruptOutput) ToSystemAffinityinterruptOutput() SystemAffinityinterruptOutput {
	return o
}

func (o SystemAffinityinterruptOutput) ToSystemAffinityinterruptOutputWithContext(ctx context.Context) SystemAffinityinterruptOutput {
	return o
}

// Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
func (o SystemAffinityinterruptOutput) AffinityCpumask() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAffinityinterrupt) pulumi.StringOutput { return v.AffinityCpumask }).(pulumi.StringOutput)
}

// ID of the interrupt affinity setting.
func (o SystemAffinityinterruptOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemAffinityinterrupt) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Interrupt name.
func (o SystemAffinityinterruptOutput) Interrupt() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemAffinityinterrupt) pulumi.StringOutput { return v.Interrupt }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemAffinityinterruptOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAffinityinterrupt) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemAffinityinterruptArrayOutput struct{ *pulumi.OutputState }

func (SystemAffinityinterruptArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemAffinityinterrupt)(nil)).Elem()
}

func (o SystemAffinityinterruptArrayOutput) ToSystemAffinityinterruptArrayOutput() SystemAffinityinterruptArrayOutput {
	return o
}

func (o SystemAffinityinterruptArrayOutput) ToSystemAffinityinterruptArrayOutputWithContext(ctx context.Context) SystemAffinityinterruptArrayOutput {
	return o
}

func (o SystemAffinityinterruptArrayOutput) Index(i pulumi.IntInput) SystemAffinityinterruptOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemAffinityinterrupt {
		return vs[0].([]*SystemAffinityinterrupt)[vs[1].(int)]
	}).(SystemAffinityinterruptOutput)
}

type SystemAffinityinterruptMapOutput struct{ *pulumi.OutputState }

func (SystemAffinityinterruptMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemAffinityinterrupt)(nil)).Elem()
}

func (o SystemAffinityinterruptMapOutput) ToSystemAffinityinterruptMapOutput() SystemAffinityinterruptMapOutput {
	return o
}

func (o SystemAffinityinterruptMapOutput) ToSystemAffinityinterruptMapOutputWithContext(ctx context.Context) SystemAffinityinterruptMapOutput {
	return o
}

func (o SystemAffinityinterruptMapOutput) MapIndex(k pulumi.StringInput) SystemAffinityinterruptOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemAffinityinterrupt {
		return vs[0].(map[string]*SystemAffinityinterrupt)[vs[1].(string)]
	}).(SystemAffinityinterruptOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAffinityinterruptInput)(nil)).Elem(), &SystemAffinityinterrupt{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAffinityinterruptArrayInput)(nil)).Elem(), SystemAffinityinterruptArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemAffinityinterruptMapInput)(nil)).Elem(), SystemAffinityinterruptMap{})
	pulumi.RegisterOutputType(SystemAffinityinterruptOutput{})
	pulumi.RegisterOutputType(SystemAffinityinterruptArrayOutput{})
	pulumi.RegisterOutputType(SystemAffinityinterruptMapOutput{})
}