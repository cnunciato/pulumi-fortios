// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure explicit FTP proxy settings.
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
//			_, err := fortios.NewFtpproxyExplicit(ctx, "trname", &fortios.FtpproxyExplicitArgs{
//				IncomingIp:       pulumi.String("0.0.0.0"),
//				SecDefaultAction: pulumi.String("deny"),
//				Status:           pulumi.String("disable"),
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
// # FtpProxy Explicit can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/ftpproxyExplicit:FtpproxyExplicit labelname FtpProxyExplicit
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/ftpproxyExplicit:FtpproxyExplicit labelname FtpProxyExplicit
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FtpproxyExplicit struct {
	pulumi.CustomResourceState

	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringOutput `pulumi:"incomingIp"`
	// Accept incoming FTP requests on one or more ports.
	IncomingPort pulumi.StringOutput `pulumi:"incomingPort"`
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp pulumi.StringOutput `pulumi:"outgoingIp"`
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringOutput `pulumi:"secDefaultAction"`
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl pulumi.StringOutput `pulumi:"ssl"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringOutput `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert pulumi.StringOutput `pulumi:"sslCert"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringOutput `pulumi:"sslDhBits"`
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFtpproxyExplicit registers a new resource with the given unique name, arguments, and options.
func NewFtpproxyExplicit(ctx *pulumi.Context,
	name string, args *FtpproxyExplicitArgs, opts ...pulumi.ResourceOption) (*FtpproxyExplicit, error) {
	if args == nil {
		args = &FtpproxyExplicitArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FtpproxyExplicit
	err := ctx.RegisterResource("fortios:index/ftpproxyExplicit:FtpproxyExplicit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFtpproxyExplicit gets an existing FtpproxyExplicit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFtpproxyExplicit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FtpproxyExplicitState, opts ...pulumi.ResourceOption) (*FtpproxyExplicit, error) {
	var resource FtpproxyExplicit
	err := ctx.ReadResource("fortios:index/ftpproxyExplicit:FtpproxyExplicit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FtpproxyExplicit resources.
type ftpproxyExplicitState struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp *string `pulumi:"incomingIp"`
	// Accept incoming FTP requests on one or more ports.
	IncomingPort *string `pulumi:"incomingPort"`
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp *string `pulumi:"outgoingIp"`
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl *string `pulumi:"ssl"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert *string `pulumi:"sslCert"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FtpproxyExplicitState struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringPtrInput
	// Accept incoming FTP requests on one or more ports.
	IncomingPort pulumi.StringPtrInput
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp pulumi.StringPtrInput
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringPtrInput
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl pulumi.StringPtrInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FtpproxyExplicitState) ElementType() reflect.Type {
	return reflect.TypeOf((*ftpproxyExplicitState)(nil)).Elem()
}

type ftpproxyExplicitArgs struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp *string `pulumi:"incomingIp"`
	// Accept incoming FTP requests on one or more ports.
	IncomingPort *string `pulumi:"incomingPort"`
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp *string `pulumi:"outgoingIp"`
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction *string `pulumi:"secDefaultAction"`
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl *string `pulumi:"ssl"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert *string `pulumi:"sslCert"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FtpproxyExplicit resource.
type FtpproxyExplicitArgs struct {
	// Accept incoming FTP requests from this IP address. An interface must have this IP address.
	IncomingIp pulumi.StringPtrInput
	// Accept incoming FTP requests on one or more ports.
	IncomingPort pulumi.StringPtrInput
	// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
	OutgoingIp pulumi.StringPtrInput
	// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
	SecDefaultAction pulumi.StringPtrInput
	// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
	Ssl pulumi.StringPtrInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FtpproxyExplicitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ftpproxyExplicitArgs)(nil)).Elem()
}

type FtpproxyExplicitInput interface {
	pulumi.Input

	ToFtpproxyExplicitOutput() FtpproxyExplicitOutput
	ToFtpproxyExplicitOutputWithContext(ctx context.Context) FtpproxyExplicitOutput
}

func (*FtpproxyExplicit) ElementType() reflect.Type {
	return reflect.TypeOf((**FtpproxyExplicit)(nil)).Elem()
}

func (i *FtpproxyExplicit) ToFtpproxyExplicitOutput() FtpproxyExplicitOutput {
	return i.ToFtpproxyExplicitOutputWithContext(context.Background())
}

func (i *FtpproxyExplicit) ToFtpproxyExplicitOutputWithContext(ctx context.Context) FtpproxyExplicitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpproxyExplicitOutput)
}

// FtpproxyExplicitArrayInput is an input type that accepts FtpproxyExplicitArray and FtpproxyExplicitArrayOutput values.
// You can construct a concrete instance of `FtpproxyExplicitArrayInput` via:
//
//	FtpproxyExplicitArray{ FtpproxyExplicitArgs{...} }
type FtpproxyExplicitArrayInput interface {
	pulumi.Input

	ToFtpproxyExplicitArrayOutput() FtpproxyExplicitArrayOutput
	ToFtpproxyExplicitArrayOutputWithContext(context.Context) FtpproxyExplicitArrayOutput
}

type FtpproxyExplicitArray []FtpproxyExplicitInput

func (FtpproxyExplicitArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FtpproxyExplicit)(nil)).Elem()
}

func (i FtpproxyExplicitArray) ToFtpproxyExplicitArrayOutput() FtpproxyExplicitArrayOutput {
	return i.ToFtpproxyExplicitArrayOutputWithContext(context.Background())
}

func (i FtpproxyExplicitArray) ToFtpproxyExplicitArrayOutputWithContext(ctx context.Context) FtpproxyExplicitArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpproxyExplicitArrayOutput)
}

// FtpproxyExplicitMapInput is an input type that accepts FtpproxyExplicitMap and FtpproxyExplicitMapOutput values.
// You can construct a concrete instance of `FtpproxyExplicitMapInput` via:
//
//	FtpproxyExplicitMap{ "key": FtpproxyExplicitArgs{...} }
type FtpproxyExplicitMapInput interface {
	pulumi.Input

	ToFtpproxyExplicitMapOutput() FtpproxyExplicitMapOutput
	ToFtpproxyExplicitMapOutputWithContext(context.Context) FtpproxyExplicitMapOutput
}

type FtpproxyExplicitMap map[string]FtpproxyExplicitInput

func (FtpproxyExplicitMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FtpproxyExplicit)(nil)).Elem()
}

func (i FtpproxyExplicitMap) ToFtpproxyExplicitMapOutput() FtpproxyExplicitMapOutput {
	return i.ToFtpproxyExplicitMapOutputWithContext(context.Background())
}

func (i FtpproxyExplicitMap) ToFtpproxyExplicitMapOutputWithContext(ctx context.Context) FtpproxyExplicitMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FtpproxyExplicitMapOutput)
}

type FtpproxyExplicitOutput struct{ *pulumi.OutputState }

func (FtpproxyExplicitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FtpproxyExplicit)(nil)).Elem()
}

func (o FtpproxyExplicitOutput) ToFtpproxyExplicitOutput() FtpproxyExplicitOutput {
	return o
}

func (o FtpproxyExplicitOutput) ToFtpproxyExplicitOutputWithContext(ctx context.Context) FtpproxyExplicitOutput {
	return o
}

// Accept incoming FTP requests from this IP address. An interface must have this IP address.
func (o FtpproxyExplicitOutput) IncomingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.IncomingIp }).(pulumi.StringOutput)
}

// Accept incoming FTP requests on one or more ports.
func (o FtpproxyExplicitOutput) IncomingPort() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.IncomingPort }).(pulumi.StringOutput)
}

// Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
func (o FtpproxyExplicitOutput) OutgoingIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.OutgoingIp }).(pulumi.StringOutput)
}

// Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
func (o FtpproxyExplicitOutput) SecDefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.SecDefaultAction }).(pulumi.StringOutput)
}

// Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
func (o FtpproxyExplicitOutput) Ssl() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.Ssl }).(pulumi.StringOutput)
}

// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
func (o FtpproxyExplicitOutput) SslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.SslAlgorithm }).(pulumi.StringOutput)
}

// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
func (o FtpproxyExplicitOutput) SslCert() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.SslCert }).(pulumi.StringOutput)
}

// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
func (o FtpproxyExplicitOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

// Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
func (o FtpproxyExplicitOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FtpproxyExplicitOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FtpproxyExplicit) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FtpproxyExplicitArrayOutput struct{ *pulumi.OutputState }

func (FtpproxyExplicitArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FtpproxyExplicit)(nil)).Elem()
}

func (o FtpproxyExplicitArrayOutput) ToFtpproxyExplicitArrayOutput() FtpproxyExplicitArrayOutput {
	return o
}

func (o FtpproxyExplicitArrayOutput) ToFtpproxyExplicitArrayOutputWithContext(ctx context.Context) FtpproxyExplicitArrayOutput {
	return o
}

func (o FtpproxyExplicitArrayOutput) Index(i pulumi.IntInput) FtpproxyExplicitOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FtpproxyExplicit {
		return vs[0].([]*FtpproxyExplicit)[vs[1].(int)]
	}).(FtpproxyExplicitOutput)
}

type FtpproxyExplicitMapOutput struct{ *pulumi.OutputState }

func (FtpproxyExplicitMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FtpproxyExplicit)(nil)).Elem()
}

func (o FtpproxyExplicitMapOutput) ToFtpproxyExplicitMapOutput() FtpproxyExplicitMapOutput {
	return o
}

func (o FtpproxyExplicitMapOutput) ToFtpproxyExplicitMapOutputWithContext(ctx context.Context) FtpproxyExplicitMapOutput {
	return o
}

func (o FtpproxyExplicitMapOutput) MapIndex(k pulumi.StringInput) FtpproxyExplicitOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FtpproxyExplicit {
		return vs[0].(map[string]*FtpproxyExplicit)[vs[1].(string)]
	}).(FtpproxyExplicitOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FtpproxyExplicitInput)(nil)).Elem(), &FtpproxyExplicit{})
	pulumi.RegisterInputType(reflect.TypeOf((*FtpproxyExplicitArrayInput)(nil)).Elem(), FtpproxyExplicitArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FtpproxyExplicitMapInput)(nil)).Elem(), FtpproxyExplicitMap{})
	pulumi.RegisterOutputType(FtpproxyExplicitOutput{})
	pulumi.RegisterOutputType(FtpproxyExplicitArrayOutput{})
	pulumi.RegisterOutputType(FtpproxyExplicitMapOutput{})
}