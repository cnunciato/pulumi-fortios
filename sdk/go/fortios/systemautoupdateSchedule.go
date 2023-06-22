// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure update schedule.
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
//			_, err := fortios.NewSystemautoupdateSchedule(ctx, "trname", &fortios.SystemautoupdateScheduleArgs{
//				Day:       pulumi.String("Monday"),
//				Frequency: pulumi.String("every"),
//				Status:    pulumi.String("enable"),
//				Time:      pulumi.String("02:60"),
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
// # SystemAutoupdate Schedule can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemautoupdateSchedule:SystemautoupdateSchedule labelname SystemAutoupdateSchedule
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemautoupdateSchedule:SystemautoupdateSchedule labelname SystemAutoupdateSchedule
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemautoupdateSchedule struct {
	pulumi.CustomResourceState

	// Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
	Day pulumi.StringOutput `pulumi:"day"`
	// Update frequency.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// Enable/disable scheduled updates. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Update time.
	Time pulumi.StringOutput `pulumi:"time"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemautoupdateSchedule registers a new resource with the given unique name, arguments, and options.
func NewSystemautoupdateSchedule(ctx *pulumi.Context,
	name string, args *SystemautoupdateScheduleArgs, opts ...pulumi.ResourceOption) (*SystemautoupdateSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.Time == nil {
		return nil, errors.New("invalid value for required argument 'Time'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemautoupdateSchedule
	err := ctx.RegisterResource("fortios:index/systemautoupdateSchedule:SystemautoupdateSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemautoupdateSchedule gets an existing SystemautoupdateSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemautoupdateSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemautoupdateScheduleState, opts ...pulumi.ResourceOption) (*SystemautoupdateSchedule, error) {
	var resource SystemautoupdateSchedule
	err := ctx.ReadResource("fortios:index/systemautoupdateSchedule:SystemautoupdateSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemautoupdateSchedule resources.
type systemautoupdateScheduleState struct {
	// Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
	Day *string `pulumi:"day"`
	// Update frequency.
	Frequency *string `pulumi:"frequency"`
	// Enable/disable scheduled updates. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Update time.
	Time *string `pulumi:"time"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemautoupdateScheduleState struct {
	// Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
	Day pulumi.StringPtrInput
	// Update frequency.
	Frequency pulumi.StringPtrInput
	// Enable/disable scheduled updates. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Update time.
	Time pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemautoupdateScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemautoupdateScheduleState)(nil)).Elem()
}

type systemautoupdateScheduleArgs struct {
	// Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
	Day *string `pulumi:"day"`
	// Update frequency.
	Frequency string `pulumi:"frequency"`
	// Enable/disable scheduled updates. Valid values: `enable`, `disable`.
	Status string `pulumi:"status"`
	// Update time.
	Time string `pulumi:"time"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemautoupdateSchedule resource.
type SystemautoupdateScheduleArgs struct {
	// Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
	Day pulumi.StringPtrInput
	// Update frequency.
	Frequency pulumi.StringInput
	// Enable/disable scheduled updates. Valid values: `enable`, `disable`.
	Status pulumi.StringInput
	// Update time.
	Time pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemautoupdateScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemautoupdateScheduleArgs)(nil)).Elem()
}

type SystemautoupdateScheduleInput interface {
	pulumi.Input

	ToSystemautoupdateScheduleOutput() SystemautoupdateScheduleOutput
	ToSystemautoupdateScheduleOutputWithContext(ctx context.Context) SystemautoupdateScheduleOutput
}

func (*SystemautoupdateSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemautoupdateSchedule)(nil)).Elem()
}

func (i *SystemautoupdateSchedule) ToSystemautoupdateScheduleOutput() SystemautoupdateScheduleOutput {
	return i.ToSystemautoupdateScheduleOutputWithContext(context.Background())
}

func (i *SystemautoupdateSchedule) ToSystemautoupdateScheduleOutputWithContext(ctx context.Context) SystemautoupdateScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemautoupdateScheduleOutput)
}

// SystemautoupdateScheduleArrayInput is an input type that accepts SystemautoupdateScheduleArray and SystemautoupdateScheduleArrayOutput values.
// You can construct a concrete instance of `SystemautoupdateScheduleArrayInput` via:
//
//	SystemautoupdateScheduleArray{ SystemautoupdateScheduleArgs{...} }
type SystemautoupdateScheduleArrayInput interface {
	pulumi.Input

	ToSystemautoupdateScheduleArrayOutput() SystemautoupdateScheduleArrayOutput
	ToSystemautoupdateScheduleArrayOutputWithContext(context.Context) SystemautoupdateScheduleArrayOutput
}

type SystemautoupdateScheduleArray []SystemautoupdateScheduleInput

func (SystemautoupdateScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemautoupdateSchedule)(nil)).Elem()
}

func (i SystemautoupdateScheduleArray) ToSystemautoupdateScheduleArrayOutput() SystemautoupdateScheduleArrayOutput {
	return i.ToSystemautoupdateScheduleArrayOutputWithContext(context.Background())
}

func (i SystemautoupdateScheduleArray) ToSystemautoupdateScheduleArrayOutputWithContext(ctx context.Context) SystemautoupdateScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemautoupdateScheduleArrayOutput)
}

// SystemautoupdateScheduleMapInput is an input type that accepts SystemautoupdateScheduleMap and SystemautoupdateScheduleMapOutput values.
// You can construct a concrete instance of `SystemautoupdateScheduleMapInput` via:
//
//	SystemautoupdateScheduleMap{ "key": SystemautoupdateScheduleArgs{...} }
type SystemautoupdateScheduleMapInput interface {
	pulumi.Input

	ToSystemautoupdateScheduleMapOutput() SystemautoupdateScheduleMapOutput
	ToSystemautoupdateScheduleMapOutputWithContext(context.Context) SystemautoupdateScheduleMapOutput
}

type SystemautoupdateScheduleMap map[string]SystemautoupdateScheduleInput

func (SystemautoupdateScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemautoupdateSchedule)(nil)).Elem()
}

func (i SystemautoupdateScheduleMap) ToSystemautoupdateScheduleMapOutput() SystemautoupdateScheduleMapOutput {
	return i.ToSystemautoupdateScheduleMapOutputWithContext(context.Background())
}

func (i SystemautoupdateScheduleMap) ToSystemautoupdateScheduleMapOutputWithContext(ctx context.Context) SystemautoupdateScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemautoupdateScheduleMapOutput)
}

type SystemautoupdateScheduleOutput struct{ *pulumi.OutputState }

func (SystemautoupdateScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemautoupdateSchedule)(nil)).Elem()
}

func (o SystemautoupdateScheduleOutput) ToSystemautoupdateScheduleOutput() SystemautoupdateScheduleOutput {
	return o
}

func (o SystemautoupdateScheduleOutput) ToSystemautoupdateScheduleOutputWithContext(ctx context.Context) SystemautoupdateScheduleOutput {
	return o
}

// Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
func (o SystemautoupdateScheduleOutput) Day() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemautoupdateSchedule) pulumi.StringOutput { return v.Day }).(pulumi.StringOutput)
}

// Update frequency.
func (o SystemautoupdateScheduleOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemautoupdateSchedule) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

// Enable/disable scheduled updates. Valid values: `enable`, `disable`.
func (o SystemautoupdateScheduleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemautoupdateSchedule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Update time.
func (o SystemautoupdateScheduleOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemautoupdateSchedule) pulumi.StringOutput { return v.Time }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemautoupdateScheduleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemautoupdateSchedule) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemautoupdateScheduleArrayOutput struct{ *pulumi.OutputState }

func (SystemautoupdateScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemautoupdateSchedule)(nil)).Elem()
}

func (o SystemautoupdateScheduleArrayOutput) ToSystemautoupdateScheduleArrayOutput() SystemautoupdateScheduleArrayOutput {
	return o
}

func (o SystemautoupdateScheduleArrayOutput) ToSystemautoupdateScheduleArrayOutputWithContext(ctx context.Context) SystemautoupdateScheduleArrayOutput {
	return o
}

func (o SystemautoupdateScheduleArrayOutput) Index(i pulumi.IntInput) SystemautoupdateScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemautoupdateSchedule {
		return vs[0].([]*SystemautoupdateSchedule)[vs[1].(int)]
	}).(SystemautoupdateScheduleOutput)
}

type SystemautoupdateScheduleMapOutput struct{ *pulumi.OutputState }

func (SystemautoupdateScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemautoupdateSchedule)(nil)).Elem()
}

func (o SystemautoupdateScheduleMapOutput) ToSystemautoupdateScheduleMapOutput() SystemautoupdateScheduleMapOutput {
	return o
}

func (o SystemautoupdateScheduleMapOutput) ToSystemautoupdateScheduleMapOutputWithContext(ctx context.Context) SystemautoupdateScheduleMapOutput {
	return o
}

func (o SystemautoupdateScheduleMapOutput) MapIndex(k pulumi.StringInput) SystemautoupdateScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemautoupdateSchedule {
		return vs[0].(map[string]*SystemautoupdateSchedule)[vs[1].(string)]
	}).(SystemautoupdateScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemautoupdateScheduleInput)(nil)).Elem(), &SystemautoupdateSchedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemautoupdateScheduleArrayInput)(nil)).Elem(), SystemautoupdateScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemautoupdateScheduleMapInput)(nil)).Elem(), SystemautoupdateScheduleMap{})
	pulumi.RegisterOutputType(SystemautoupdateScheduleOutput{})
	pulumi.RegisterOutputType(SystemautoupdateScheduleArrayOutput{})
	pulumi.RegisterOutputType(SystemautoupdateScheduleMapOutput{})
}