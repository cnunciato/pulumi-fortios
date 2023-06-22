// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure SSL servers.
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
//			_, err := fortios.NewFirewallSslserver(ctx, "trname", &fortios.FirewallSslserverArgs{
//				AddHeaderXForwardedProto: pulumi.String("enable"),
//				Ip:                       pulumi.String("1.1.1.1"),
//				MappedPort:               pulumi.Int(2234),
//				Port:                     pulumi.Int(32321),
//				SslAlgorithm:             pulumi.String("high"),
//				SslCert:                  pulumi.String("Fortinet_CA_SSL"),
//				SslClientRenegotiation:   pulumi.String("allow"),
//				SslDhBits:                pulumi.String("2048"),
//				SslMaxVersion:            pulumi.String("tls-1.2"),
//				SslMinVersion:            pulumi.String("tls-1.1"),
//				SslMode:                  pulumi.String("half"),
//				SslSendEmptyFrags:        pulumi.String("enable"),
//				UrlRewrite:               pulumi.String("disable"),
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
// # Firewall SslServer can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallSslserver:FirewallSslserver labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallSslserver:FirewallSslserver labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallSslserver struct {
	pulumi.CustomResourceState

	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto pulumi.StringOutput `pulumi:"addHeaderXForwardedProto"`
	// IPv4 address of the SSL server.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort pulumi.IntOutput `pulumi:"mappedPort"`
	// Server name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Server service port (1 - 65535, default = 443).
	Port pulumi.IntOutput `pulumi:"port"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringOutput `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert pulumi.StringOutput `pulumi:"sslCert"`
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation pulumi.StringOutput `pulumi:"sslClientRenegotiation"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringOutput `pulumi:"sslDhBits"`
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion pulumi.StringOutput `pulumi:"sslMaxVersion"`
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion pulumi.StringOutput `pulumi:"sslMinVersion"`
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode pulumi.StringOutput `pulumi:"sslMode"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringOutput `pulumi:"sslSendEmptyFrags"`
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite pulumi.StringOutput `pulumi:"urlRewrite"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewFirewallSslserver registers a new resource with the given unique name, arguments, and options.
func NewFirewallSslserver(ctx *pulumi.Context,
	name string, args *FirewallSslserverArgs, opts ...pulumi.ResourceOption) (*FirewallSslserver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.SslCert == nil {
		return nil, errors.New("invalid value for required argument 'SslCert'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallSslserver
	err := ctx.RegisterResource("fortios:index/firewallSslserver:FirewallSslserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallSslserver gets an existing FirewallSslserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallSslserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallSslserverState, opts ...pulumi.ResourceOption) (*FirewallSslserver, error) {
	var resource FirewallSslserver
	err := ctx.ReadResource("fortios:index/firewallSslserver:FirewallSslserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallSslserver resources.
type firewallSslserverState struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto *string `pulumi:"addHeaderXForwardedProto"`
	// IPv4 address of the SSL server.
	Ip *string `pulumi:"ip"`
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort *int `pulumi:"mappedPort"`
	// Server name.
	Name *string `pulumi:"name"`
	// Server service port (1 - 65535, default = 443).
	Port *int `pulumi:"port"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert *string `pulumi:"sslCert"`
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation *string `pulumi:"sslClientRenegotiation"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion *string `pulumi:"sslMaxVersion"`
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion *string `pulumi:"sslMinVersion"`
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode *string `pulumi:"sslMode"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags *string `pulumi:"sslSendEmptyFrags"`
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite *string `pulumi:"urlRewrite"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type FirewallSslserverState struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto pulumi.StringPtrInput
	// IPv4 address of the SSL server.
	Ip pulumi.StringPtrInput
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort pulumi.IntPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// Server service port (1 - 65535, default = 443).
	Port pulumi.IntPtrInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert pulumi.StringPtrInput
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion pulumi.StringPtrInput
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion pulumi.StringPtrInput
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode pulumi.StringPtrInput
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringPtrInput
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallSslserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSslserverState)(nil)).Elem()
}

type firewallSslserverArgs struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto *string `pulumi:"addHeaderXForwardedProto"`
	// IPv4 address of the SSL server.
	Ip string `pulumi:"ip"`
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort *int `pulumi:"mappedPort"`
	// Server name.
	Name *string `pulumi:"name"`
	// Server service port (1 - 65535, default = 443).
	Port int `pulumi:"port"`
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm *string `pulumi:"sslAlgorithm"`
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert string `pulumi:"sslCert"`
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation *string `pulumi:"sslClientRenegotiation"`
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits *string `pulumi:"sslDhBits"`
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion *string `pulumi:"sslMaxVersion"`
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion *string `pulumi:"sslMinVersion"`
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode *string `pulumi:"sslMode"`
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags *string `pulumi:"sslSendEmptyFrags"`
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite *string `pulumi:"urlRewrite"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a FirewallSslserver resource.
type FirewallSslserverArgs struct {
	// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
	AddHeaderXForwardedProto pulumi.StringPtrInput
	// IPv4 address of the SSL server.
	Ip pulumi.StringInput
	// Mapped server service port (1 - 65535, default = 80).
	MappedPort pulumi.IntPtrInput
	// Server name.
	Name pulumi.StringPtrInput
	// Server service port (1 - 65535, default = 443).
	Port pulumi.IntInput
	// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
	SslAlgorithm pulumi.StringPtrInput
	// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
	SslCert pulumi.StringInput
	// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
	SslClientRenegotiation pulumi.StringPtrInput
	// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
	SslDhBits pulumi.StringPtrInput
	// Highest SSL/TLS version to negotiate.
	SslMaxVersion pulumi.StringPtrInput
	// Lowest SSL/TLS version to negotiate.
	SslMinVersion pulumi.StringPtrInput
	// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
	SslMode pulumi.StringPtrInput
	// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
	SslSendEmptyFrags pulumi.StringPtrInput
	// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
	UrlRewrite pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (FirewallSslserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallSslserverArgs)(nil)).Elem()
}

type FirewallSslserverInput interface {
	pulumi.Input

	ToFirewallSslserverOutput() FirewallSslserverOutput
	ToFirewallSslserverOutputWithContext(ctx context.Context) FirewallSslserverOutput
}

func (*FirewallSslserver) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSslserver)(nil)).Elem()
}

func (i *FirewallSslserver) ToFirewallSslserverOutput() FirewallSslserverOutput {
	return i.ToFirewallSslserverOutputWithContext(context.Background())
}

func (i *FirewallSslserver) ToFirewallSslserverOutputWithContext(ctx context.Context) FirewallSslserverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslserverOutput)
}

// FirewallSslserverArrayInput is an input type that accepts FirewallSslserverArray and FirewallSslserverArrayOutput values.
// You can construct a concrete instance of `FirewallSslserverArrayInput` via:
//
//	FirewallSslserverArray{ FirewallSslserverArgs{...} }
type FirewallSslserverArrayInput interface {
	pulumi.Input

	ToFirewallSslserverArrayOutput() FirewallSslserverArrayOutput
	ToFirewallSslserverArrayOutputWithContext(context.Context) FirewallSslserverArrayOutput
}

type FirewallSslserverArray []FirewallSslserverInput

func (FirewallSslserverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSslserver)(nil)).Elem()
}

func (i FirewallSslserverArray) ToFirewallSslserverArrayOutput() FirewallSslserverArrayOutput {
	return i.ToFirewallSslserverArrayOutputWithContext(context.Background())
}

func (i FirewallSslserverArray) ToFirewallSslserverArrayOutputWithContext(ctx context.Context) FirewallSslserverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslserverArrayOutput)
}

// FirewallSslserverMapInput is an input type that accepts FirewallSslserverMap and FirewallSslserverMapOutput values.
// You can construct a concrete instance of `FirewallSslserverMapInput` via:
//
//	FirewallSslserverMap{ "key": FirewallSslserverArgs{...} }
type FirewallSslserverMapInput interface {
	pulumi.Input

	ToFirewallSslserverMapOutput() FirewallSslserverMapOutput
	ToFirewallSslserverMapOutputWithContext(context.Context) FirewallSslserverMapOutput
}

type FirewallSslserverMap map[string]FirewallSslserverInput

func (FirewallSslserverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSslserver)(nil)).Elem()
}

func (i FirewallSslserverMap) ToFirewallSslserverMapOutput() FirewallSslserverMapOutput {
	return i.ToFirewallSslserverMapOutputWithContext(context.Background())
}

func (i FirewallSslserverMap) ToFirewallSslserverMapOutputWithContext(ctx context.Context) FirewallSslserverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallSslserverMapOutput)
}

type FirewallSslserverOutput struct{ *pulumi.OutputState }

func (FirewallSslserverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallSslserver)(nil)).Elem()
}

func (o FirewallSslserverOutput) ToFirewallSslserverOutput() FirewallSslserverOutput {
	return o
}

func (o FirewallSslserverOutput) ToFirewallSslserverOutputWithContext(ctx context.Context) FirewallSslserverOutput {
	return o
}

// Enable/disable adding an X-Forwarded-Proto header to forwarded requests. Valid values: `enable`, `disable`.
func (o FirewallSslserverOutput) AddHeaderXForwardedProto() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.AddHeaderXForwardedProto }).(pulumi.StringOutput)
}

// IPv4 address of the SSL server.
func (o FirewallSslserverOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Mapped server service port (1 - 65535, default = 80).
func (o FirewallSslserverOutput) MappedPort() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.IntOutput { return v.MappedPort }).(pulumi.IntOutput)
}

// Server name.
func (o FirewallSslserverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Server service port (1 - 65535, default = 443).
func (o FirewallSslserverOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
func (o FirewallSslserverOutput) SslAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslAlgorithm }).(pulumi.StringOutput)
}

// Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
func (o FirewallSslserverOutput) SslCert() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslCert }).(pulumi.StringOutput)
}

// Allow or block client renegotiation by server. Valid values: `allow`, `deny`, `secure`.
func (o FirewallSslserverOutput) SslClientRenegotiation() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslClientRenegotiation }).(pulumi.StringOutput)
}

// Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
func (o FirewallSslserverOutput) SslDhBits() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslDhBits }).(pulumi.StringOutput)
}

// Highest SSL/TLS version to negotiate.
func (o FirewallSslserverOutput) SslMaxVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslMaxVersion }).(pulumi.StringOutput)
}

// Lowest SSL/TLS version to negotiate.
func (o FirewallSslserverOutput) SslMinVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslMinVersion }).(pulumi.StringOutput)
}

// SSL/TLS mode for encryption and decryption of traffic. Valid values: `half`, `full`.
func (o FirewallSslserverOutput) SslMode() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslMode }).(pulumi.StringOutput)
}

// Enable/disable sending empty fragments to avoid attack on CBC IV. Valid values: `enable`, `disable`.
func (o FirewallSslserverOutput) SslSendEmptyFrags() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.SslSendEmptyFrags }).(pulumi.StringOutput)
}

// Enable/disable rewriting the URL. Valid values: `enable`, `disable`.
func (o FirewallSslserverOutput) UrlRewrite() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringOutput { return v.UrlRewrite }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallSslserverOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallSslserver) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type FirewallSslserverArrayOutput struct{ *pulumi.OutputState }

func (FirewallSslserverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallSslserver)(nil)).Elem()
}

func (o FirewallSslserverArrayOutput) ToFirewallSslserverArrayOutput() FirewallSslserverArrayOutput {
	return o
}

func (o FirewallSslserverArrayOutput) ToFirewallSslserverArrayOutputWithContext(ctx context.Context) FirewallSslserverArrayOutput {
	return o
}

func (o FirewallSslserverArrayOutput) Index(i pulumi.IntInput) FirewallSslserverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallSslserver {
		return vs[0].([]*FirewallSslserver)[vs[1].(int)]
	}).(FirewallSslserverOutput)
}

type FirewallSslserverMapOutput struct{ *pulumi.OutputState }

func (FirewallSslserverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallSslserver)(nil)).Elem()
}

func (o FirewallSslserverMapOutput) ToFirewallSslserverMapOutput() FirewallSslserverMapOutput {
	return o
}

func (o FirewallSslserverMapOutput) ToFirewallSslserverMapOutputWithContext(ctx context.Context) FirewallSslserverMapOutput {
	return o
}

func (o FirewallSslserverMapOutput) MapIndex(k pulumi.StringInput) FirewallSslserverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallSslserver {
		return vs[0].(map[string]*FirewallSslserver)[vs[1].(string)]
	}).(FirewallSslserverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslserverInput)(nil)).Elem(), &FirewallSslserver{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslserverArrayInput)(nil)).Elem(), FirewallSslserverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallSslserverMapInput)(nil)).Elem(), FirewallSslserverMap{})
	pulumi.RegisterOutputType(FirewallSslserverOutput{})
	pulumi.RegisterOutputType(FirewallSslserverArrayOutput{})
	pulumi.RegisterOutputType(FirewallSslserverMapOutput{})
}