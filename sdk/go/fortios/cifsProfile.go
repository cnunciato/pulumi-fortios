// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure CIFS profile. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1`.
//
// ## Import
//
// # Cifs Profile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/cifsProfile:CifsProfile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/cifsProfile:CifsProfile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type CifsProfile struct {
	pulumi.CustomResourceState

	// Domain for which to decrypt CIFS traffic.
	DomainController pulumi.StringOutput `pulumi:"domainController"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter CifsProfileFileFilterOutput `pulumi:"fileFilter"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
	ServerCredentialType pulumi.StringOutput `pulumi:"serverCredentialType"`
	// Server keytab. The structure of `serverKeytab` block is documented below.
	ServerKeytabs CifsProfileServerKeytabArrayOutput `pulumi:"serverKeytabs"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewCifsProfile registers a new resource with the given unique name, arguments, and options.
func NewCifsProfile(ctx *pulumi.Context,
	name string, args *CifsProfileArgs, opts ...pulumi.ResourceOption) (*CifsProfile, error) {
	if args == nil {
		args = &CifsProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource CifsProfile
	err := ctx.RegisterResource("fortios:index/cifsProfile:CifsProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCifsProfile gets an existing CifsProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCifsProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CifsProfileState, opts ...pulumi.ResourceOption) (*CifsProfile, error) {
	var resource CifsProfile
	err := ctx.ReadResource("fortios:index/cifsProfile:CifsProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CifsProfile resources.
type cifsProfileState struct {
	// Domain for which to decrypt CIFS traffic.
	DomainController *string `pulumi:"domainController"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *CifsProfileFileFilter `pulumi:"fileFilter"`
	// Profile name.
	Name *string `pulumi:"name"`
	// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
	ServerCredentialType *string `pulumi:"serverCredentialType"`
	// Server keytab. The structure of `serverKeytab` block is documented below.
	ServerKeytabs []CifsProfileServerKeytab `pulumi:"serverKeytabs"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type CifsProfileState struct {
	// Domain for which to decrypt CIFS traffic.
	DomainController pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter CifsProfileFileFilterPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
	ServerCredentialType pulumi.StringPtrInput
	// Server keytab. The structure of `serverKeytab` block is documented below.
	ServerKeytabs CifsProfileServerKeytabArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CifsProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*cifsProfileState)(nil)).Elem()
}

type cifsProfileArgs struct {
	// Domain for which to decrypt CIFS traffic.
	DomainController *string `pulumi:"domainController"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter *CifsProfileFileFilter `pulumi:"fileFilter"`
	// Profile name.
	Name *string `pulumi:"name"`
	// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
	ServerCredentialType *string `pulumi:"serverCredentialType"`
	// Server keytab. The structure of `serverKeytab` block is documented below.
	ServerKeytabs []CifsProfileServerKeytab `pulumi:"serverKeytabs"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a CifsProfile resource.
type CifsProfileArgs struct {
	// Domain for which to decrypt CIFS traffic.
	DomainController pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// File filter. The structure of `fileFilter` block is documented below.
	FileFilter CifsProfileFileFilterPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
	ServerCredentialType pulumi.StringPtrInput
	// Server keytab. The structure of `serverKeytab` block is documented below.
	ServerKeytabs CifsProfileServerKeytabArrayInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (CifsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cifsProfileArgs)(nil)).Elem()
}

type CifsProfileInput interface {
	pulumi.Input

	ToCifsProfileOutput() CifsProfileOutput
	ToCifsProfileOutputWithContext(ctx context.Context) CifsProfileOutput
}

func (*CifsProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**CifsProfile)(nil)).Elem()
}

func (i *CifsProfile) ToCifsProfileOutput() CifsProfileOutput {
	return i.ToCifsProfileOutputWithContext(context.Background())
}

func (i *CifsProfile) ToCifsProfileOutputWithContext(ctx context.Context) CifsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CifsProfileOutput)
}

// CifsProfileArrayInput is an input type that accepts CifsProfileArray and CifsProfileArrayOutput values.
// You can construct a concrete instance of `CifsProfileArrayInput` via:
//
//	CifsProfileArray{ CifsProfileArgs{...} }
type CifsProfileArrayInput interface {
	pulumi.Input

	ToCifsProfileArrayOutput() CifsProfileArrayOutput
	ToCifsProfileArrayOutputWithContext(context.Context) CifsProfileArrayOutput
}

type CifsProfileArray []CifsProfileInput

func (CifsProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CifsProfile)(nil)).Elem()
}

func (i CifsProfileArray) ToCifsProfileArrayOutput() CifsProfileArrayOutput {
	return i.ToCifsProfileArrayOutputWithContext(context.Background())
}

func (i CifsProfileArray) ToCifsProfileArrayOutputWithContext(ctx context.Context) CifsProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CifsProfileArrayOutput)
}

// CifsProfileMapInput is an input type that accepts CifsProfileMap and CifsProfileMapOutput values.
// You can construct a concrete instance of `CifsProfileMapInput` via:
//
//	CifsProfileMap{ "key": CifsProfileArgs{...} }
type CifsProfileMapInput interface {
	pulumi.Input

	ToCifsProfileMapOutput() CifsProfileMapOutput
	ToCifsProfileMapOutputWithContext(context.Context) CifsProfileMapOutput
}

type CifsProfileMap map[string]CifsProfileInput

func (CifsProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CifsProfile)(nil)).Elem()
}

func (i CifsProfileMap) ToCifsProfileMapOutput() CifsProfileMapOutput {
	return i.ToCifsProfileMapOutputWithContext(context.Background())
}

func (i CifsProfileMap) ToCifsProfileMapOutputWithContext(ctx context.Context) CifsProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CifsProfileMapOutput)
}

type CifsProfileOutput struct{ *pulumi.OutputState }

func (CifsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CifsProfile)(nil)).Elem()
}

func (o CifsProfileOutput) ToCifsProfileOutput() CifsProfileOutput {
	return o
}

func (o CifsProfileOutput) ToCifsProfileOutputWithContext(ctx context.Context) CifsProfileOutput {
	return o
}

// Domain for which to decrypt CIFS traffic.
func (o CifsProfileOutput) DomainController() pulumi.StringOutput {
	return o.ApplyT(func(v *CifsProfile) pulumi.StringOutput { return v.DomainController }).(pulumi.StringOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o CifsProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CifsProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// File filter. The structure of `fileFilter` block is documented below.
func (o CifsProfileOutput) FileFilter() CifsProfileFileFilterOutput {
	return o.ApplyT(func(v *CifsProfile) CifsProfileFileFilterOutput { return v.FileFilter }).(CifsProfileFileFilterOutput)
}

// Profile name.
func (o CifsProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CifsProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// CIFS server credential type. Valid values: `none`, `credential-replication`, `credential-keytab`.
func (o CifsProfileOutput) ServerCredentialType() pulumi.StringOutput {
	return o.ApplyT(func(v *CifsProfile) pulumi.StringOutput { return v.ServerCredentialType }).(pulumi.StringOutput)
}

// Server keytab. The structure of `serverKeytab` block is documented below.
func (o CifsProfileOutput) ServerKeytabs() CifsProfileServerKeytabArrayOutput {
	return o.ApplyT(func(v *CifsProfile) CifsProfileServerKeytabArrayOutput { return v.ServerKeytabs }).(CifsProfileServerKeytabArrayOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o CifsProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CifsProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type CifsProfileArrayOutput struct{ *pulumi.OutputState }

func (CifsProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CifsProfile)(nil)).Elem()
}

func (o CifsProfileArrayOutput) ToCifsProfileArrayOutput() CifsProfileArrayOutput {
	return o
}

func (o CifsProfileArrayOutput) ToCifsProfileArrayOutputWithContext(ctx context.Context) CifsProfileArrayOutput {
	return o
}

func (o CifsProfileArrayOutput) Index(i pulumi.IntInput) CifsProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CifsProfile {
		return vs[0].([]*CifsProfile)[vs[1].(int)]
	}).(CifsProfileOutput)
}

type CifsProfileMapOutput struct{ *pulumi.OutputState }

func (CifsProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CifsProfile)(nil)).Elem()
}

func (o CifsProfileMapOutput) ToCifsProfileMapOutput() CifsProfileMapOutput {
	return o
}

func (o CifsProfileMapOutput) ToCifsProfileMapOutputWithContext(ctx context.Context) CifsProfileMapOutput {
	return o
}

func (o CifsProfileMapOutput) MapIndex(k pulumi.StringInput) CifsProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CifsProfile {
		return vs[0].(map[string]*CifsProfile)[vs[1].(string)]
	}).(CifsProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CifsProfileInput)(nil)).Elem(), &CifsProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*CifsProfileArrayInput)(nil)).Elem(), CifsProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CifsProfileMapInput)(nil)).Elem(), CifsProfileMap{})
	pulumi.RegisterOutputType(CifsProfileOutput{})
	pulumi.RegisterOutputType(CifsProfileArrayOutput{})
	pulumi.RegisterOutputType(CifsProfileMapOutput{})
}