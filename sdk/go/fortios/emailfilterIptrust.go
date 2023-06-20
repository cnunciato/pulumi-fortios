// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure AntiSpam IP trust. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # Emailfilter Iptrust can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/emailfilterIptrust:EmailfilterIptrust labelname {{fosid}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/emailfilterIptrust:EmailfilterIptrust labelname {{fosid}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type EmailfilterIptrust struct {
	pulumi.CustomResourceState

	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries EmailfilterIptrustEntryArrayOutput `pulumi:"entries"`
	// ID.
	Fosid pulumi.IntOutput `pulumi:"fosid"`
	// Name of table.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewEmailfilterIptrust registers a new resource with the given unique name, arguments, and options.
func NewEmailfilterIptrust(ctx *pulumi.Context,
	name string, args *EmailfilterIptrustArgs, opts ...pulumi.ResourceOption) (*EmailfilterIptrust, error) {
	if args == nil {
		args = &EmailfilterIptrustArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource EmailfilterIptrust
	err := ctx.RegisterResource("fortios:index/emailfilterIptrust:EmailfilterIptrust", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmailfilterIptrust gets an existing EmailfilterIptrust resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmailfilterIptrust(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmailfilterIptrustState, opts ...pulumi.ResourceOption) (*EmailfilterIptrust, error) {
	var resource EmailfilterIptrust
	err := ctx.ReadResource("fortios:index/emailfilterIptrust:EmailfilterIptrust", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmailfilterIptrust resources.
type emailfilterIptrustState struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries []EmailfilterIptrustEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type EmailfilterIptrustState struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries EmailfilterIptrustEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterIptrustState) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterIptrustState)(nil)).Elem()
}

type emailfilterIptrustArgs struct {
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries []EmailfilterIptrustEntry `pulumi:"entries"`
	// ID.
	Fosid *int `pulumi:"fosid"`
	// Name of table.
	Name *string `pulumi:"name"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a EmailfilterIptrust resource.
type EmailfilterIptrustArgs struct {
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
	Entries EmailfilterIptrustEntryArrayInput
	// ID.
	Fosid pulumi.IntPtrInput
	// Name of table.
	Name pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (EmailfilterIptrustArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*emailfilterIptrustArgs)(nil)).Elem()
}

type EmailfilterIptrustInput interface {
	pulumi.Input

	ToEmailfilterIptrustOutput() EmailfilterIptrustOutput
	ToEmailfilterIptrustOutputWithContext(ctx context.Context) EmailfilterIptrustOutput
}

func (*EmailfilterIptrust) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterIptrust)(nil)).Elem()
}

func (i *EmailfilterIptrust) ToEmailfilterIptrustOutput() EmailfilterIptrustOutput {
	return i.ToEmailfilterIptrustOutputWithContext(context.Background())
}

func (i *EmailfilterIptrust) ToEmailfilterIptrustOutputWithContext(ctx context.Context) EmailfilterIptrustOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterIptrustOutput)
}

// EmailfilterIptrustArrayInput is an input type that accepts EmailfilterIptrustArray and EmailfilterIptrustArrayOutput values.
// You can construct a concrete instance of `EmailfilterIptrustArrayInput` via:
//
//	EmailfilterIptrustArray{ EmailfilterIptrustArgs{...} }
type EmailfilterIptrustArrayInput interface {
	pulumi.Input

	ToEmailfilterIptrustArrayOutput() EmailfilterIptrustArrayOutput
	ToEmailfilterIptrustArrayOutputWithContext(context.Context) EmailfilterIptrustArrayOutput
}

type EmailfilterIptrustArray []EmailfilterIptrustInput

func (EmailfilterIptrustArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterIptrust)(nil)).Elem()
}

func (i EmailfilterIptrustArray) ToEmailfilterIptrustArrayOutput() EmailfilterIptrustArrayOutput {
	return i.ToEmailfilterIptrustArrayOutputWithContext(context.Background())
}

func (i EmailfilterIptrustArray) ToEmailfilterIptrustArrayOutputWithContext(ctx context.Context) EmailfilterIptrustArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterIptrustArrayOutput)
}

// EmailfilterIptrustMapInput is an input type that accepts EmailfilterIptrustMap and EmailfilterIptrustMapOutput values.
// You can construct a concrete instance of `EmailfilterIptrustMapInput` via:
//
//	EmailfilterIptrustMap{ "key": EmailfilterIptrustArgs{...} }
type EmailfilterIptrustMapInput interface {
	pulumi.Input

	ToEmailfilterIptrustMapOutput() EmailfilterIptrustMapOutput
	ToEmailfilterIptrustMapOutputWithContext(context.Context) EmailfilterIptrustMapOutput
}

type EmailfilterIptrustMap map[string]EmailfilterIptrustInput

func (EmailfilterIptrustMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterIptrust)(nil)).Elem()
}

func (i EmailfilterIptrustMap) ToEmailfilterIptrustMapOutput() EmailfilterIptrustMapOutput {
	return i.ToEmailfilterIptrustMapOutputWithContext(context.Background())
}

func (i EmailfilterIptrustMap) ToEmailfilterIptrustMapOutputWithContext(ctx context.Context) EmailfilterIptrustMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailfilterIptrustMapOutput)
}

type EmailfilterIptrustOutput struct{ *pulumi.OutputState }

func (EmailfilterIptrustOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailfilterIptrust)(nil)).Elem()
}

func (o EmailfilterIptrustOutput) ToEmailfilterIptrustOutput() EmailfilterIptrustOutput {
	return o
}

func (o EmailfilterIptrustOutput) ToEmailfilterIptrustOutputWithContext(ctx context.Context) EmailfilterIptrustOutput {
	return o
}

// Optional comments.
func (o EmailfilterIptrustOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterIptrust) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o EmailfilterIptrustOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterIptrust) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Spam filter trusted IP addresses. The structure of `entries` block is documented below.
func (o EmailfilterIptrustOutput) Entries() EmailfilterIptrustEntryArrayOutput {
	return o.ApplyT(func(v *EmailfilterIptrust) EmailfilterIptrustEntryArrayOutput { return v.Entries }).(EmailfilterIptrustEntryArrayOutput)
}

// ID.
func (o EmailfilterIptrustOutput) Fosid() pulumi.IntOutput {
	return o.ApplyT(func(v *EmailfilterIptrust) pulumi.IntOutput { return v.Fosid }).(pulumi.IntOutput)
}

// Name of table.
func (o EmailfilterIptrustOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EmailfilterIptrust) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o EmailfilterIptrustOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailfilterIptrust) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type EmailfilterIptrustArrayOutput struct{ *pulumi.OutputState }

func (EmailfilterIptrustArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmailfilterIptrust)(nil)).Elem()
}

func (o EmailfilterIptrustArrayOutput) ToEmailfilterIptrustArrayOutput() EmailfilterIptrustArrayOutput {
	return o
}

func (o EmailfilterIptrustArrayOutput) ToEmailfilterIptrustArrayOutputWithContext(ctx context.Context) EmailfilterIptrustArrayOutput {
	return o
}

func (o EmailfilterIptrustArrayOutput) Index(i pulumi.IntInput) EmailfilterIptrustOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmailfilterIptrust {
		return vs[0].([]*EmailfilterIptrust)[vs[1].(int)]
	}).(EmailfilterIptrustOutput)
}

type EmailfilterIptrustMapOutput struct{ *pulumi.OutputState }

func (EmailfilterIptrustMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmailfilterIptrust)(nil)).Elem()
}

func (o EmailfilterIptrustMapOutput) ToEmailfilterIptrustMapOutput() EmailfilterIptrustMapOutput {
	return o
}

func (o EmailfilterIptrustMapOutput) ToEmailfilterIptrustMapOutputWithContext(ctx context.Context) EmailfilterIptrustMapOutput {
	return o
}

func (o EmailfilterIptrustMapOutput) MapIndex(k pulumi.StringInput) EmailfilterIptrustOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmailfilterIptrust {
		return vs[0].(map[string]*EmailfilterIptrust)[vs[1].(string)]
	}).(EmailfilterIptrustOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterIptrustInput)(nil)).Elem(), &EmailfilterIptrust{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterIptrustArrayInput)(nil)).Elem(), EmailfilterIptrustArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailfilterIptrustMapInput)(nil)).Elem(), EmailfilterIptrustMap{})
	pulumi.RegisterOutputType(EmailfilterIptrustOutput{})
	pulumi.RegisterOutputType(EmailfilterIptrustArrayOutput{})
	pulumi.RegisterOutputType(EmailfilterIptrustMapOutput{})
}
