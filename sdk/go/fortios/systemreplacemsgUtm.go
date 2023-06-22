// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Replacement messages.
//
// ## Import
//
// # SystemReplacemsg Utm can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgUtm:SystemreplacemsgUtm labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgUtm:SystemreplacemsgUtm labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgUtm struct {
	pulumi.CustomResourceState

	// Message string.
	Buffer pulumi.StringPtrOutput `pulumi:"buffer"`
	// Format flag.
	Format pulumi.StringOutput `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringOutput `pulumi:"header"`
	// Message type.
	MsgType pulumi.StringOutput `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemreplacemsgUtm registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgUtm(ctx *pulumi.Context,
	name string, args *SystemreplacemsgUtmArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgUtm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgUtm
	err := ctx.RegisterResource("fortios:index/systemreplacemsgUtm:SystemreplacemsgUtm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgUtm gets an existing SystemreplacemsgUtm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgUtm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgUtmState, opts ...pulumi.ResourceOption) (*SystemreplacemsgUtm, error) {
	var resource SystemreplacemsgUtm
	err := ctx.ReadResource("fortios:index/systemreplacemsgUtm:SystemreplacemsgUtm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgUtm resources.
type systemreplacemsgUtmState struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType *string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemreplacemsgUtmState struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgUtmState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgUtmState)(nil)).Elem()
}

type systemreplacemsgUtmArgs struct {
	// Message string.
	Buffer *string `pulumi:"buffer"`
	// Format flag.
	Format *string `pulumi:"format"`
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header *string `pulumi:"header"`
	// Message type.
	MsgType string `pulumi:"msgType"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemreplacemsgUtm resource.
type SystemreplacemsgUtmArgs struct {
	// Message string.
	Buffer pulumi.StringPtrInput
	// Format flag.
	Format pulumi.StringPtrInput
	// Header flag. Valid values: `none`, `http`, `8bit`.
	Header pulumi.StringPtrInput
	// Message type.
	MsgType pulumi.StringInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemreplacemsgUtmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgUtmArgs)(nil)).Elem()
}

type SystemreplacemsgUtmInput interface {
	pulumi.Input

	ToSystemreplacemsgUtmOutput() SystemreplacemsgUtmOutput
	ToSystemreplacemsgUtmOutputWithContext(ctx context.Context) SystemreplacemsgUtmOutput
}

func (*SystemreplacemsgUtm) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgUtm)(nil)).Elem()
}

func (i *SystemreplacemsgUtm) ToSystemreplacemsgUtmOutput() SystemreplacemsgUtmOutput {
	return i.ToSystemreplacemsgUtmOutputWithContext(context.Background())
}

func (i *SystemreplacemsgUtm) ToSystemreplacemsgUtmOutputWithContext(ctx context.Context) SystemreplacemsgUtmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgUtmOutput)
}

// SystemreplacemsgUtmArrayInput is an input type that accepts SystemreplacemsgUtmArray and SystemreplacemsgUtmArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgUtmArrayInput` via:
//
//	SystemreplacemsgUtmArray{ SystemreplacemsgUtmArgs{...} }
type SystemreplacemsgUtmArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgUtmArrayOutput() SystemreplacemsgUtmArrayOutput
	ToSystemreplacemsgUtmArrayOutputWithContext(context.Context) SystemreplacemsgUtmArrayOutput
}

type SystemreplacemsgUtmArray []SystemreplacemsgUtmInput

func (SystemreplacemsgUtmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgUtm)(nil)).Elem()
}

func (i SystemreplacemsgUtmArray) ToSystemreplacemsgUtmArrayOutput() SystemreplacemsgUtmArrayOutput {
	return i.ToSystemreplacemsgUtmArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgUtmArray) ToSystemreplacemsgUtmArrayOutputWithContext(ctx context.Context) SystemreplacemsgUtmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgUtmArrayOutput)
}

// SystemreplacemsgUtmMapInput is an input type that accepts SystemreplacemsgUtmMap and SystemreplacemsgUtmMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgUtmMapInput` via:
//
//	SystemreplacemsgUtmMap{ "key": SystemreplacemsgUtmArgs{...} }
type SystemreplacemsgUtmMapInput interface {
	pulumi.Input

	ToSystemreplacemsgUtmMapOutput() SystemreplacemsgUtmMapOutput
	ToSystemreplacemsgUtmMapOutputWithContext(context.Context) SystemreplacemsgUtmMapOutput
}

type SystemreplacemsgUtmMap map[string]SystemreplacemsgUtmInput

func (SystemreplacemsgUtmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgUtm)(nil)).Elem()
}

func (i SystemreplacemsgUtmMap) ToSystemreplacemsgUtmMapOutput() SystemreplacemsgUtmMapOutput {
	return i.ToSystemreplacemsgUtmMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgUtmMap) ToSystemreplacemsgUtmMapOutputWithContext(ctx context.Context) SystemreplacemsgUtmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgUtmMapOutput)
}

type SystemreplacemsgUtmOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgUtmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgUtm)(nil)).Elem()
}

func (o SystemreplacemsgUtmOutput) ToSystemreplacemsgUtmOutput() SystemreplacemsgUtmOutput {
	return o
}

func (o SystemreplacemsgUtmOutput) ToSystemreplacemsgUtmOutputWithContext(ctx context.Context) SystemreplacemsgUtmOutput {
	return o
}

// Message string.
func (o SystemreplacemsgUtmOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgUtm) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SystemreplacemsgUtmOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgUtm) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgUtmOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgUtm) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgUtmOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgUtm) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgUtmOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgUtm) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgUtmArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgUtmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgUtm)(nil)).Elem()
}

func (o SystemreplacemsgUtmArrayOutput) ToSystemreplacemsgUtmArrayOutput() SystemreplacemsgUtmArrayOutput {
	return o
}

func (o SystemreplacemsgUtmArrayOutput) ToSystemreplacemsgUtmArrayOutputWithContext(ctx context.Context) SystemreplacemsgUtmArrayOutput {
	return o
}

func (o SystemreplacemsgUtmArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgUtmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgUtm {
		return vs[0].([]*SystemreplacemsgUtm)[vs[1].(int)]
	}).(SystemreplacemsgUtmOutput)
}

type SystemreplacemsgUtmMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgUtmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgUtm)(nil)).Elem()
}

func (o SystemreplacemsgUtmMapOutput) ToSystemreplacemsgUtmMapOutput() SystemreplacemsgUtmMapOutput {
	return o
}

func (o SystemreplacemsgUtmMapOutput) ToSystemreplacemsgUtmMapOutputWithContext(ctx context.Context) SystemreplacemsgUtmMapOutput {
	return o
}

func (o SystemreplacemsgUtmMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgUtmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgUtm {
		return vs[0].(map[string]*SystemreplacemsgUtm)[vs[1].(string)]
	}).(SystemreplacemsgUtmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgUtmInput)(nil)).Elem(), &SystemreplacemsgUtm{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgUtmArrayInput)(nil)).Elem(), SystemreplacemsgUtmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgUtmMapInput)(nil)).Elem(), SystemreplacemsgUtmMap{})
	pulumi.RegisterOutputType(SystemreplacemsgUtmOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgUtmArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgUtmMapOutput{})
}