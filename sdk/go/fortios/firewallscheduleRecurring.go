// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Recurring schedule configuration.
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
//			_, err := fortios.NewFirewallscheduleRecurring(ctx, "trname", &fortios.FirewallscheduleRecurringArgs{
//				Color: pulumi.Int(0),
//				Day:   pulumi.String("sunday"),
//				End:   pulumi.String("00:00"),
//				Start: pulumi.String("00:00"),
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
// # FirewallSchedule Recurring can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallscheduleRecurring:FirewallscheduleRecurring labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallscheduleRecurring:FirewallscheduleRecurring labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallscheduleRecurring struct {
	pulumi.CustomResourceState

	// Color of icon on the GUI.
	Color pulumi.IntOutput `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day pulumi.StringOutput `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End pulumi.StringOutput `pulumi:"end"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringOutput `pulumi:"fabricObject"`
	// Recurring schedule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start pulumi.StringOutput `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallscheduleRecurring registers a new resource with the given unique name, arguments, and options.
func NewFirewallscheduleRecurring(ctx *pulumi.Context,
	name string, args *FirewallscheduleRecurringArgs, opts ...pulumi.ResourceOption) (*FirewallscheduleRecurring, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.End == nil {
		return nil, errors.New("invalid value for required argument 'End'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallscheduleRecurring
	err := ctx.RegisterResource("fortios:index/firewallscheduleRecurring:FirewallscheduleRecurring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallscheduleRecurring gets an existing FirewallscheduleRecurring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallscheduleRecurring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallscheduleRecurringState, opts ...pulumi.ResourceOption) (*FirewallscheduleRecurring, error) {
	var resource FirewallscheduleRecurring
	err := ctx.ReadResource("fortios:index/firewallscheduleRecurring:FirewallscheduleRecurring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallscheduleRecurring resources.
type firewallscheduleRecurringState struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day *string `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End *string `pulumi:"end"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Recurring schedule name.
	Name *string `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start *string `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallscheduleRecurringState struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day pulumi.StringPtrInput
	// Time of day to end the schedule, format hh:mm.
	End pulumi.StringPtrInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Recurring schedule name.
	Name pulumi.StringPtrInput
	// Time of day to start the schedule, format hh:mm.
	Start pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallscheduleRecurringState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallscheduleRecurringState)(nil)).Elem()
}

type firewallscheduleRecurringArgs struct {
	// Color of icon on the GUI.
	Color *int `pulumi:"color"`
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day *string `pulumi:"day"`
	// Time of day to end the schedule, format hh:mm.
	End string `pulumi:"end"`
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject *string `pulumi:"fabricObject"`
	// Recurring schedule name.
	Name *string `pulumi:"name"`
	// Time of day to start the schedule, format hh:mm.
	Start string `pulumi:"start"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallscheduleRecurring resource.
type FirewallscheduleRecurringArgs struct {
	// Color of icon on the GUI.
	Color pulumi.IntPtrInput
	// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
	Day pulumi.StringPtrInput
	// Time of day to end the schedule, format hh:mm.
	End pulumi.StringInput
	// Security Fabric global object setting. Valid values: `enable`, `disable`.
	FabricObject pulumi.StringPtrInput
	// Recurring schedule name.
	Name pulumi.StringPtrInput
	// Time of day to start the schedule, format hh:mm.
	Start pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallscheduleRecurringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallscheduleRecurringArgs)(nil)).Elem()
}

type FirewallscheduleRecurringInput interface {
	pulumi.Input

	ToFirewallscheduleRecurringOutput() FirewallscheduleRecurringOutput
	ToFirewallscheduleRecurringOutputWithContext(ctx context.Context) FirewallscheduleRecurringOutput
}

func (*FirewallscheduleRecurring) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallscheduleRecurring)(nil)).Elem()
}

func (i *FirewallscheduleRecurring) ToFirewallscheduleRecurringOutput() FirewallscheduleRecurringOutput {
	return i.ToFirewallscheduleRecurringOutputWithContext(context.Background())
}

func (i *FirewallscheduleRecurring) ToFirewallscheduleRecurringOutputWithContext(ctx context.Context) FirewallscheduleRecurringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallscheduleRecurringOutput)
}

// FirewallscheduleRecurringArrayInput is an input type that accepts FirewallscheduleRecurringArray and FirewallscheduleRecurringArrayOutput values.
// You can construct a concrete instance of `FirewallscheduleRecurringArrayInput` via:
//
//	FirewallscheduleRecurringArray{ FirewallscheduleRecurringArgs{...} }
type FirewallscheduleRecurringArrayInput interface {
	pulumi.Input

	ToFirewallscheduleRecurringArrayOutput() FirewallscheduleRecurringArrayOutput
	ToFirewallscheduleRecurringArrayOutputWithContext(context.Context) FirewallscheduleRecurringArrayOutput
}

type FirewallscheduleRecurringArray []FirewallscheduleRecurringInput

func (FirewallscheduleRecurringArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallscheduleRecurring)(nil)).Elem()
}

func (i FirewallscheduleRecurringArray) ToFirewallscheduleRecurringArrayOutput() FirewallscheduleRecurringArrayOutput {
	return i.ToFirewallscheduleRecurringArrayOutputWithContext(context.Background())
}

func (i FirewallscheduleRecurringArray) ToFirewallscheduleRecurringArrayOutputWithContext(ctx context.Context) FirewallscheduleRecurringArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallscheduleRecurringArrayOutput)
}

// FirewallscheduleRecurringMapInput is an input type that accepts FirewallscheduleRecurringMap and FirewallscheduleRecurringMapOutput values.
// You can construct a concrete instance of `FirewallscheduleRecurringMapInput` via:
//
//	FirewallscheduleRecurringMap{ "key": FirewallscheduleRecurringArgs{...} }
type FirewallscheduleRecurringMapInput interface {
	pulumi.Input

	ToFirewallscheduleRecurringMapOutput() FirewallscheduleRecurringMapOutput
	ToFirewallscheduleRecurringMapOutputWithContext(context.Context) FirewallscheduleRecurringMapOutput
}

type FirewallscheduleRecurringMap map[string]FirewallscheduleRecurringInput

func (FirewallscheduleRecurringMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallscheduleRecurring)(nil)).Elem()
}

func (i FirewallscheduleRecurringMap) ToFirewallscheduleRecurringMapOutput() FirewallscheduleRecurringMapOutput {
	return i.ToFirewallscheduleRecurringMapOutputWithContext(context.Background())
}

func (i FirewallscheduleRecurringMap) ToFirewallscheduleRecurringMapOutputWithContext(ctx context.Context) FirewallscheduleRecurringMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallscheduleRecurringMapOutput)
}

type FirewallscheduleRecurringOutput struct{ *pulumi.OutputState }

func (FirewallscheduleRecurringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallscheduleRecurring)(nil)).Elem()
}

func (o FirewallscheduleRecurringOutput) ToFirewallscheduleRecurringOutput() FirewallscheduleRecurringOutput {
	return o
}

func (o FirewallscheduleRecurringOutput) ToFirewallscheduleRecurringOutputWithContext(ctx context.Context) FirewallscheduleRecurringOutput {
	return o
}

// Color of icon on the GUI.
func (o FirewallscheduleRecurringOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
func (o FirewallscheduleRecurringOutput) Day() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.StringOutput { return v.Day }).(pulumi.StringOutput)
}

// Time of day to end the schedule, format hh:mm.
func (o FirewallscheduleRecurringOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.StringOutput { return v.End }).(pulumi.StringOutput)
}

// Security Fabric global object setting. Valid values: `enable`, `disable`.
func (o FirewallscheduleRecurringOutput) FabricObject() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.StringOutput { return v.FabricObject }).(pulumi.StringOutput)
}

// Recurring schedule name.
func (o FirewallscheduleRecurringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Time of day to start the schedule, format hh:mm.
func (o FirewallscheduleRecurringOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallscheduleRecurringOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallscheduleRecurring) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallscheduleRecurringArrayOutput struct{ *pulumi.OutputState }

func (FirewallscheduleRecurringArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallscheduleRecurring)(nil)).Elem()
}

func (o FirewallscheduleRecurringArrayOutput) ToFirewallscheduleRecurringArrayOutput() FirewallscheduleRecurringArrayOutput {
	return o
}

func (o FirewallscheduleRecurringArrayOutput) ToFirewallscheduleRecurringArrayOutputWithContext(ctx context.Context) FirewallscheduleRecurringArrayOutput {
	return o
}

func (o FirewallscheduleRecurringArrayOutput) Index(i pulumi.IntInput) FirewallscheduleRecurringOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallscheduleRecurring {
		return vs[0].([]*FirewallscheduleRecurring)[vs[1].(int)]
	}).(FirewallscheduleRecurringOutput)
}

type FirewallscheduleRecurringMapOutput struct{ *pulumi.OutputState }

func (FirewallscheduleRecurringMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallscheduleRecurring)(nil)).Elem()
}

func (o FirewallscheduleRecurringMapOutput) ToFirewallscheduleRecurringMapOutput() FirewallscheduleRecurringMapOutput {
	return o
}

func (o FirewallscheduleRecurringMapOutput) ToFirewallscheduleRecurringMapOutputWithContext(ctx context.Context) FirewallscheduleRecurringMapOutput {
	return o
}

func (o FirewallscheduleRecurringMapOutput) MapIndex(k pulumi.StringInput) FirewallscheduleRecurringOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallscheduleRecurring {
		return vs[0].(map[string]*FirewallscheduleRecurring)[vs[1].(string)]
	}).(FirewallscheduleRecurringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallscheduleRecurringInput)(nil)).Elem(), &FirewallscheduleRecurring{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallscheduleRecurringArrayInput)(nil)).Elem(), FirewallscheduleRecurringArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallscheduleRecurringMapInput)(nil)).Elem(), FirewallscheduleRecurringMap{})
	pulumi.RegisterOutputType(FirewallscheduleRecurringOutput{})
	pulumi.RegisterOutputType(FirewallscheduleRecurringArrayOutput{})
	pulumi.RegisterOutputType(FirewallscheduleRecurringMapOutput{})
}
