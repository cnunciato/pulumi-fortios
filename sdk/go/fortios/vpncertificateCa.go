// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CA certificate.
//
// ## Import
//
// # VpnCertificate Ca can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/vpncertificateCa:VpncertificateCa labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/vpncertificateCa:VpncertificateCa labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type VpncertificateCa struct {
	pulumi.CustomResourceState

	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays pulumi.IntOutput `pulumi:"autoUpdateDays"`
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning pulumi.IntOutput `pulumi:"autoUpdateDaysWarning"`
	// CA certificate as a PEM file.
	Ca pulumi.StringOutput `pulumi:"ca"`
	// CA identifier of the SCEP server.
	CaIdentifier pulumi.StringOutput `pulumi:"caIdentifier"`
	// Time at which CA was last updated.
	LastUpdated pulumi.IntOutput `pulumi:"lastUpdated"`
	// Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete pulumi.StringOutput `pulumi:"obsolete"`
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringOutput `pulumi:"range"`
	// URL of the SCEP server.
	ScepUrl pulumi.StringOutput `pulumi:"scepUrl"`
	// CA certificate source type.
	Source pulumi.StringOutput `pulumi:"source"`
	// Source IP address for communications to the SCEP server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted pulumi.StringOutput `pulumi:"sslInspectionTrusted"`
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted pulumi.StringOutput `pulumi:"trusted"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpncertificateCa registers a new resource with the given unique name, arguments, and options.
func NewVpncertificateCa(ctx *pulumi.Context,
	name string, args *VpncertificateCaArgs, opts ...pulumi.ResourceOption) (*VpncertificateCa, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ca == nil {
		return nil, errors.New("invalid value for required argument 'Ca'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpncertificateCa
	err := ctx.RegisterResource("fortios:index/vpncertificateCa:VpncertificateCa", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpncertificateCa gets an existing VpncertificateCa resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpncertificateCa(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpncertificateCaState, opts ...pulumi.ResourceOption) (*VpncertificateCa, error) {
	var resource VpncertificateCa
	err := ctx.ReadResource("fortios:index/vpncertificateCa:VpncertificateCa", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpncertificateCa resources.
type vpncertificateCaState struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays *int `pulumi:"autoUpdateDays"`
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning *int `pulumi:"autoUpdateDaysWarning"`
	// CA certificate as a PEM file.
	Ca *string `pulumi:"ca"`
	// CA identifier of the SCEP server.
	CaIdentifier *string `pulumi:"caIdentifier"`
	// Time at which CA was last updated.
	LastUpdated *int `pulumi:"lastUpdated"`
	// Name.
	Name *string `pulumi:"name"`
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete *string `pulumi:"obsolete"`
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// URL of the SCEP server.
	ScepUrl *string `pulumi:"scepUrl"`
	// CA certificate source type.
	Source *string `pulumi:"source"`
	// Source IP address for communications to the SCEP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted *string `pulumi:"sslInspectionTrusted"`
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted *string `pulumi:"trusted"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpncertificateCaState struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays pulumi.IntPtrInput
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning pulumi.IntPtrInput
	// CA certificate as a PEM file.
	Ca pulumi.StringPtrInput
	// CA identifier of the SCEP server.
	CaIdentifier pulumi.StringPtrInput
	// Time at which CA was last updated.
	LastUpdated pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete pulumi.StringPtrInput
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// URL of the SCEP server.
	ScepUrl pulumi.StringPtrInput
	// CA certificate source type.
	Source pulumi.StringPtrInput
	// Source IP address for communications to the SCEP server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted pulumi.StringPtrInput
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpncertificateCaState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpncertificateCaState)(nil)).Elem()
}

type vpncertificateCaArgs struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays *int `pulumi:"autoUpdateDays"`
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning *int `pulumi:"autoUpdateDaysWarning"`
	// CA certificate as a PEM file.
	Ca string `pulumi:"ca"`
	// CA identifier of the SCEP server.
	CaIdentifier *string `pulumi:"caIdentifier"`
	// Time at which CA was last updated.
	LastUpdated *int `pulumi:"lastUpdated"`
	// Name.
	Name *string `pulumi:"name"`
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete *string `pulumi:"obsolete"`
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range *string `pulumi:"range"`
	// URL of the SCEP server.
	ScepUrl *string `pulumi:"scepUrl"`
	// CA certificate source type.
	Source *string `pulumi:"source"`
	// Source IP address for communications to the SCEP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted *string `pulumi:"sslInspectionTrusted"`
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted *string `pulumi:"trusted"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpncertificateCa resource.
type VpncertificateCaArgs struct {
	// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
	AutoUpdateDays pulumi.IntPtrInput
	// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
	AutoUpdateDaysWarning pulumi.IntPtrInput
	// CA certificate as a PEM file.
	Ca pulumi.StringInput
	// CA identifier of the SCEP server.
	CaIdentifier pulumi.StringPtrInput
	// Time at which CA was last updated.
	LastUpdated pulumi.IntPtrInput
	// Name.
	Name pulumi.StringPtrInput
	// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
	Obsolete pulumi.StringPtrInput
	// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
	Range pulumi.StringPtrInput
	// URL of the SCEP server.
	ScepUrl pulumi.StringPtrInput
	// CA certificate source type.
	Source pulumi.StringPtrInput
	// Source IP address for communications to the SCEP server.
	SourceIp pulumi.StringPtrInput
	// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
	SslInspectionTrusted pulumi.StringPtrInput
	// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
	Trusted pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpncertificateCaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpncertificateCaArgs)(nil)).Elem()
}

type VpncertificateCaInput interface {
	pulumi.Input

	ToVpncertificateCaOutput() VpncertificateCaOutput
	ToVpncertificateCaOutputWithContext(ctx context.Context) VpncertificateCaOutput
}

func (*VpncertificateCa) ElementType() reflect.Type {
	return reflect.TypeOf((**VpncertificateCa)(nil)).Elem()
}

func (i *VpncertificateCa) ToVpncertificateCaOutput() VpncertificateCaOutput {
	return i.ToVpncertificateCaOutputWithContext(context.Background())
}

func (i *VpncertificateCa) ToVpncertificateCaOutputWithContext(ctx context.Context) VpncertificateCaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateCaOutput)
}

// VpncertificateCaArrayInput is an input type that accepts VpncertificateCaArray and VpncertificateCaArrayOutput values.
// You can construct a concrete instance of `VpncertificateCaArrayInput` via:
//
//	VpncertificateCaArray{ VpncertificateCaArgs{...} }
type VpncertificateCaArrayInput interface {
	pulumi.Input

	ToVpncertificateCaArrayOutput() VpncertificateCaArrayOutput
	ToVpncertificateCaArrayOutputWithContext(context.Context) VpncertificateCaArrayOutput
}

type VpncertificateCaArray []VpncertificateCaInput

func (VpncertificateCaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpncertificateCa)(nil)).Elem()
}

func (i VpncertificateCaArray) ToVpncertificateCaArrayOutput() VpncertificateCaArrayOutput {
	return i.ToVpncertificateCaArrayOutputWithContext(context.Background())
}

func (i VpncertificateCaArray) ToVpncertificateCaArrayOutputWithContext(ctx context.Context) VpncertificateCaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateCaArrayOutput)
}

// VpncertificateCaMapInput is an input type that accepts VpncertificateCaMap and VpncertificateCaMapOutput values.
// You can construct a concrete instance of `VpncertificateCaMapInput` via:
//
//	VpncertificateCaMap{ "key": VpncertificateCaArgs{...} }
type VpncertificateCaMapInput interface {
	pulumi.Input

	ToVpncertificateCaMapOutput() VpncertificateCaMapOutput
	ToVpncertificateCaMapOutputWithContext(context.Context) VpncertificateCaMapOutput
}

type VpncertificateCaMap map[string]VpncertificateCaInput

func (VpncertificateCaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpncertificateCa)(nil)).Elem()
}

func (i VpncertificateCaMap) ToVpncertificateCaMapOutput() VpncertificateCaMapOutput {
	return i.ToVpncertificateCaMapOutputWithContext(context.Background())
}

func (i VpncertificateCaMap) ToVpncertificateCaMapOutputWithContext(ctx context.Context) VpncertificateCaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateCaMapOutput)
}

type VpncertificateCaOutput struct{ *pulumi.OutputState }

func (VpncertificateCaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpncertificateCa)(nil)).Elem()
}

func (o VpncertificateCaOutput) ToVpncertificateCaOutput() VpncertificateCaOutput {
	return o
}

func (o VpncertificateCaOutput) ToVpncertificateCaOutputWithContext(ctx context.Context) VpncertificateCaOutput {
	return o
}

// Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).
func (o VpncertificateCaOutput) AutoUpdateDays() pulumi.IntOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.IntOutput { return v.AutoUpdateDays }).(pulumi.IntOutput)
}

// Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).
func (o VpncertificateCaOutput) AutoUpdateDaysWarning() pulumi.IntOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.IntOutput { return v.AutoUpdateDaysWarning }).(pulumi.IntOutput)
}

// CA certificate as a PEM file.
func (o VpncertificateCaOutput) Ca() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.Ca }).(pulumi.StringOutput)
}

// CA identifier of the SCEP server.
func (o VpncertificateCaOutput) CaIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.CaIdentifier }).(pulumi.StringOutput)
}

// Time at which CA was last updated.
func (o VpncertificateCaOutput) LastUpdated() pulumi.IntOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.IntOutput { return v.LastUpdated }).(pulumi.IntOutput)
}

// Name.
func (o VpncertificateCaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable this CA as obsoleted. Valid values: `disable`, `enable`.
func (o VpncertificateCaOutput) Obsolete() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.Obsolete }).(pulumi.StringOutput)
}

// Either global or VDOM IP address range for the CA certificate. Valid values: `global`, `vdom`.
func (o VpncertificateCaOutput) Range() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.Range }).(pulumi.StringOutput)
}

// URL of the SCEP server.
func (o VpncertificateCaOutput) ScepUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.ScepUrl }).(pulumi.StringOutput)
}

// CA certificate source type.
func (o VpncertificateCaOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Source IP address for communications to the SCEP server.
func (o VpncertificateCaOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable this CA as a trusted CA for SSL inspection. Valid values: `enable`, `disable`.
func (o VpncertificateCaOutput) SslInspectionTrusted() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.SslInspectionTrusted }).(pulumi.StringOutput)
}

// Enable/disable as a trusted CA. Valid values: `enable`, `disable`.
func (o VpncertificateCaOutput) Trusted() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringOutput { return v.Trusted }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VpncertificateCaOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpncertificateCa) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpncertificateCaArrayOutput struct{ *pulumi.OutputState }

func (VpncertificateCaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpncertificateCa)(nil)).Elem()
}

func (o VpncertificateCaArrayOutput) ToVpncertificateCaArrayOutput() VpncertificateCaArrayOutput {
	return o
}

func (o VpncertificateCaArrayOutput) ToVpncertificateCaArrayOutputWithContext(ctx context.Context) VpncertificateCaArrayOutput {
	return o
}

func (o VpncertificateCaArrayOutput) Index(i pulumi.IntInput) VpncertificateCaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpncertificateCa {
		return vs[0].([]*VpncertificateCa)[vs[1].(int)]
	}).(VpncertificateCaOutput)
}

type VpncertificateCaMapOutput struct{ *pulumi.OutputState }

func (VpncertificateCaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpncertificateCa)(nil)).Elem()
}

func (o VpncertificateCaMapOutput) ToVpncertificateCaMapOutput() VpncertificateCaMapOutput {
	return o
}

func (o VpncertificateCaMapOutput) ToVpncertificateCaMapOutputWithContext(ctx context.Context) VpncertificateCaMapOutput {
	return o
}

func (o VpncertificateCaMapOutput) MapIndex(k pulumi.StringInput) VpncertificateCaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpncertificateCa {
		return vs[0].(map[string]*VpncertificateCa)[vs[1].(string)]
	}).(VpncertificateCaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateCaInput)(nil)).Elem(), &VpncertificateCa{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateCaArrayInput)(nil)).Elem(), VpncertificateCaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateCaMapInput)(nil)).Elem(), VpncertificateCaMap{})
	pulumi.RegisterOutputType(VpncertificateCaOutput{})
	pulumi.RegisterOutputType(VpncertificateCaArrayOutput{})
	pulumi.RegisterOutputType(VpncertificateCaMapOutput{})
}