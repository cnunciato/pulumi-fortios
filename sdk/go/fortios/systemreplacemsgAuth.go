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
// # SystemReplacemsg Auth can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgAuth:SystemreplacemsgAuth labelname {{msg_type}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemreplacemsgAuth:SystemreplacemsgAuth labelname {{msg_type}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemreplacemsgAuth struct {
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

// NewSystemreplacemsgAuth registers a new resource with the given unique name, arguments, and options.
func NewSystemreplacemsgAuth(ctx *pulumi.Context,
	name string, args *SystemreplacemsgAuthArgs, opts ...pulumi.ResourceOption) (*SystemreplacemsgAuth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MsgType == nil {
		return nil, errors.New("invalid value for required argument 'MsgType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SystemreplacemsgAuth
	err := ctx.RegisterResource("fortios:index/systemreplacemsgAuth:SystemreplacemsgAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemreplacemsgAuth gets an existing SystemreplacemsgAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemreplacemsgAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemreplacemsgAuthState, opts ...pulumi.ResourceOption) (*SystemreplacemsgAuth, error) {
	var resource SystemreplacemsgAuth
	err := ctx.ReadResource("fortios:index/systemreplacemsgAuth:SystemreplacemsgAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemreplacemsgAuth resources.
type systemreplacemsgAuthState struct {
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

type SystemreplacemsgAuthState struct {
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

func (SystemreplacemsgAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgAuthState)(nil)).Elem()
}

type systemreplacemsgAuthArgs struct {
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

// The set of arguments for constructing a SystemreplacemsgAuth resource.
type SystemreplacemsgAuthArgs struct {
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

func (SystemreplacemsgAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemreplacemsgAuthArgs)(nil)).Elem()
}

type SystemreplacemsgAuthInput interface {
	pulumi.Input

	ToSystemreplacemsgAuthOutput() SystemreplacemsgAuthOutput
	ToSystemreplacemsgAuthOutputWithContext(ctx context.Context) SystemreplacemsgAuthOutput
}

func (*SystemreplacemsgAuth) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgAuth)(nil)).Elem()
}

func (i *SystemreplacemsgAuth) ToSystemreplacemsgAuthOutput() SystemreplacemsgAuthOutput {
	return i.ToSystemreplacemsgAuthOutputWithContext(context.Background())
}

func (i *SystemreplacemsgAuth) ToSystemreplacemsgAuthOutputWithContext(ctx context.Context) SystemreplacemsgAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgAuthOutput)
}

// SystemreplacemsgAuthArrayInput is an input type that accepts SystemreplacemsgAuthArray and SystemreplacemsgAuthArrayOutput values.
// You can construct a concrete instance of `SystemreplacemsgAuthArrayInput` via:
//
//	SystemreplacemsgAuthArray{ SystemreplacemsgAuthArgs{...} }
type SystemreplacemsgAuthArrayInput interface {
	pulumi.Input

	ToSystemreplacemsgAuthArrayOutput() SystemreplacemsgAuthArrayOutput
	ToSystemreplacemsgAuthArrayOutputWithContext(context.Context) SystemreplacemsgAuthArrayOutput
}

type SystemreplacemsgAuthArray []SystemreplacemsgAuthInput

func (SystemreplacemsgAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgAuth)(nil)).Elem()
}

func (i SystemreplacemsgAuthArray) ToSystemreplacemsgAuthArrayOutput() SystemreplacemsgAuthArrayOutput {
	return i.ToSystemreplacemsgAuthArrayOutputWithContext(context.Background())
}

func (i SystemreplacemsgAuthArray) ToSystemreplacemsgAuthArrayOutputWithContext(ctx context.Context) SystemreplacemsgAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgAuthArrayOutput)
}

// SystemreplacemsgAuthMapInput is an input type that accepts SystemreplacemsgAuthMap and SystemreplacemsgAuthMapOutput values.
// You can construct a concrete instance of `SystemreplacemsgAuthMapInput` via:
//
//	SystemreplacemsgAuthMap{ "key": SystemreplacemsgAuthArgs{...} }
type SystemreplacemsgAuthMapInput interface {
	pulumi.Input

	ToSystemreplacemsgAuthMapOutput() SystemreplacemsgAuthMapOutput
	ToSystemreplacemsgAuthMapOutputWithContext(context.Context) SystemreplacemsgAuthMapOutput
}

type SystemreplacemsgAuthMap map[string]SystemreplacemsgAuthInput

func (SystemreplacemsgAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgAuth)(nil)).Elem()
}

func (i SystemreplacemsgAuthMap) ToSystemreplacemsgAuthMapOutput() SystemreplacemsgAuthMapOutput {
	return i.ToSystemreplacemsgAuthMapOutputWithContext(context.Background())
}

func (i SystemreplacemsgAuthMap) ToSystemreplacemsgAuthMapOutputWithContext(ctx context.Context) SystemreplacemsgAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemreplacemsgAuthMapOutput)
}

type SystemreplacemsgAuthOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemreplacemsgAuth)(nil)).Elem()
}

func (o SystemreplacemsgAuthOutput) ToSystemreplacemsgAuthOutput() SystemreplacemsgAuthOutput {
	return o
}

func (o SystemreplacemsgAuthOutput) ToSystemreplacemsgAuthOutputWithContext(ctx context.Context) SystemreplacemsgAuthOutput {
	return o
}

// Message string.
func (o SystemreplacemsgAuthOutput) Buffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgAuth) pulumi.StringPtrOutput { return v.Buffer }).(pulumi.StringPtrOutput)
}

// Format flag.
func (o SystemreplacemsgAuthOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgAuth) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Header flag. Valid values: `none`, `http`, `8bit`.
func (o SystemreplacemsgAuthOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgAuth) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// Message type.
func (o SystemreplacemsgAuthOutput) MsgType() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemreplacemsgAuth) pulumi.StringOutput { return v.MsgType }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemreplacemsgAuthOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemreplacemsgAuth) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemreplacemsgAuthArrayOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemreplacemsgAuth)(nil)).Elem()
}

func (o SystemreplacemsgAuthArrayOutput) ToSystemreplacemsgAuthArrayOutput() SystemreplacemsgAuthArrayOutput {
	return o
}

func (o SystemreplacemsgAuthArrayOutput) ToSystemreplacemsgAuthArrayOutputWithContext(ctx context.Context) SystemreplacemsgAuthArrayOutput {
	return o
}

func (o SystemreplacemsgAuthArrayOutput) Index(i pulumi.IntInput) SystemreplacemsgAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemreplacemsgAuth {
		return vs[0].([]*SystemreplacemsgAuth)[vs[1].(int)]
	}).(SystemreplacemsgAuthOutput)
}

type SystemreplacemsgAuthMapOutput struct{ *pulumi.OutputState }

func (SystemreplacemsgAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemreplacemsgAuth)(nil)).Elem()
}

func (o SystemreplacemsgAuthMapOutput) ToSystemreplacemsgAuthMapOutput() SystemreplacemsgAuthMapOutput {
	return o
}

func (o SystemreplacemsgAuthMapOutput) ToSystemreplacemsgAuthMapOutputWithContext(ctx context.Context) SystemreplacemsgAuthMapOutput {
	return o
}

func (o SystemreplacemsgAuthMapOutput) MapIndex(k pulumi.StringInput) SystemreplacemsgAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemreplacemsgAuth {
		return vs[0].(map[string]*SystemreplacemsgAuth)[vs[1].(string)]
	}).(SystemreplacemsgAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgAuthInput)(nil)).Elem(), &SystemreplacemsgAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgAuthArrayInput)(nil)).Elem(), SystemreplacemsgAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemreplacemsgAuthMapInput)(nil)).Elem(), SystemreplacemsgAuthMap{})
	pulumi.RegisterOutputType(SystemreplacemsgAuthOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgAuthArrayOutput{})
	pulumi.RegisterOutputType(SystemreplacemsgAuthMapOutput{})
}
