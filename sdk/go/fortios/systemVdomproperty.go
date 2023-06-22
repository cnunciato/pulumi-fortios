// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure VDOM property.
//
// ## Import
//
// # System VdomProperty can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/systemVdomproperty:SystemVdomproperty labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/systemVdomproperty:SystemVdomproperty labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SystemVdomproperty struct {
	pulumi.CustomResourceState

	// Maximum guaranteed number of firewall custom services.
	CustomService pulumi.StringOutput `pulumi:"customService"`
	// Description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel pulumi.StringOutput `pulumi:"dialupTunnel"`
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress pulumi.StringOutput `pulumi:"firewallAddress"`
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp pulumi.StringOutput `pulumi:"firewallAddrgrp"`
	// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 pulumi.StringOutput `pulumi:"ipsecPhase1"`
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface pulumi.StringOutput `pulumi:"ipsecPhase1Interface"`
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 pulumi.StringOutput `pulumi:"ipsecPhase2"`
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface pulumi.StringOutput `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in MB (range depends on how much disk space is available).
	LogDiskQuota pulumi.StringOutput `pulumi:"logDiskQuota"`
	// VDOM name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule pulumi.StringOutput `pulumi:"onetimeSchedule"`
	// Maximum guaranteed number of concurrent proxy users.
	Proxy pulumi.StringOutput `pulumi:"proxy"`
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule pulumi.StringOutput `pulumi:"recurringSchedule"`
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup pulumi.StringOutput `pulumi:"serviceGroup"`
	// Maximum guaranteed number of sessions.
	Session pulumi.StringOutput `pulumi:"session"`
	// Permanent SNMP Index of the virtual domain (0 - 4294967295).
	SnmpIndex pulumi.IntOutput `pulumi:"snmpIndex"`
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn pulumi.StringOutput `pulumi:"sslvpn"`
	// Maximum guaranteed number of local users.
	User pulumi.StringOutput `pulumi:"user"`
	// Maximum guaranteed number of user groups.
	UserGroup pulumi.StringOutput `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSystemVdomproperty registers a new resource with the given unique name, arguments, and options.
func NewSystemVdomproperty(ctx *pulumi.Context,
	name string, args *SystemVdompropertyArgs, opts ...pulumi.ResourceOption) (*SystemVdomproperty, error) {
	if args == nil {
		args = &SystemVdompropertyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SystemVdomproperty
	err := ctx.RegisterResource("fortios:index/systemVdomproperty:SystemVdomproperty", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSystemVdomproperty gets an existing SystemVdomproperty resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSystemVdomproperty(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemVdompropertyState, opts ...pulumi.ResourceOption) (*SystemVdomproperty, error) {
	var resource SystemVdomproperty
	err := ctx.ReadResource("fortios:index/systemVdomproperty:SystemVdomproperty", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SystemVdomproperty resources.
type systemVdompropertyState struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService *string `pulumi:"customService"`
	// Description.
	Description *string `pulumi:"description"`
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel *string `pulumi:"dialupTunnel"`
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress *string `pulumi:"firewallAddress"`
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp *string `pulumi:"firewallAddrgrp"`
	// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 *string `pulumi:"ipsecPhase1"`
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface *string `pulumi:"ipsecPhase1Interface"`
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 *string `pulumi:"ipsecPhase2"`
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface *string `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in MB (range depends on how much disk space is available).
	LogDiskQuota *string `pulumi:"logDiskQuota"`
	// VDOM name.
	Name *string `pulumi:"name"`
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule *string `pulumi:"onetimeSchedule"`
	// Maximum guaranteed number of concurrent proxy users.
	Proxy *string `pulumi:"proxy"`
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule *string `pulumi:"recurringSchedule"`
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup *string `pulumi:"serviceGroup"`
	// Maximum guaranteed number of sessions.
	Session *string `pulumi:"session"`
	// Permanent SNMP Index of the virtual domain (0 - 4294967295).
	SnmpIndex *int `pulumi:"snmpIndex"`
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn *string `pulumi:"sslvpn"`
	// Maximum guaranteed number of local users.
	User *string `pulumi:"user"`
	// Maximum guaranteed number of user groups.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SystemVdompropertyState struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel pulumi.StringPtrInput
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress pulumi.StringPtrInput
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp pulumi.StringPtrInput
	// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface pulumi.StringPtrInput
	// Log disk quota in MB (range depends on how much disk space is available).
	LogDiskQuota pulumi.StringPtrInput
	// VDOM name.
	Name pulumi.StringPtrInput
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of concurrent proxy users.
	Proxy pulumi.StringPtrInput
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup pulumi.StringPtrInput
	// Maximum guaranteed number of sessions.
	Session pulumi.StringPtrInput
	// Permanent SNMP Index of the virtual domain (0 - 4294967295).
	SnmpIndex pulumi.IntPtrInput
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn pulumi.StringPtrInput
	// Maximum guaranteed number of local users.
	User pulumi.StringPtrInput
	// Maximum guaranteed number of user groups.
	UserGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdompropertyState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdompropertyState)(nil)).Elem()
}

type systemVdompropertyArgs struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService *string `pulumi:"customService"`
	// Description.
	Description *string `pulumi:"description"`
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel *string `pulumi:"dialupTunnel"`
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress *string `pulumi:"firewallAddress"`
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp *string `pulumi:"firewallAddrgrp"`
	// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 *string `pulumi:"ipsecPhase1"`
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface *string `pulumi:"ipsecPhase1Interface"`
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 *string `pulumi:"ipsecPhase2"`
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface *string `pulumi:"ipsecPhase2Interface"`
	// Log disk quota in MB (range depends on how much disk space is available).
	LogDiskQuota *string `pulumi:"logDiskQuota"`
	// VDOM name.
	Name *string `pulumi:"name"`
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule *string `pulumi:"onetimeSchedule"`
	// Maximum guaranteed number of concurrent proxy users.
	Proxy *string `pulumi:"proxy"`
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule *string `pulumi:"recurringSchedule"`
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup *string `pulumi:"serviceGroup"`
	// Maximum guaranteed number of sessions.
	Session *string `pulumi:"session"`
	// Permanent SNMP Index of the virtual domain (0 - 4294967295).
	SnmpIndex *int `pulumi:"snmpIndex"`
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn *string `pulumi:"sslvpn"`
	// Maximum guaranteed number of local users.
	User *string `pulumi:"user"`
	// Maximum guaranteed number of user groups.
	UserGroup *string `pulumi:"userGroup"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SystemVdomproperty resource.
type SystemVdompropertyArgs struct {
	// Maximum guaranteed number of firewall custom services.
	CustomService pulumi.StringPtrInput
	// Description.
	Description pulumi.StringPtrInput
	// Maximum guaranteed number of dial-up tunnels.
	DialupTunnel pulumi.StringPtrInput
	// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
	FirewallAddress pulumi.StringPtrInput
	// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
	FirewallAddrgrp pulumi.StringPtrInput
	// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
	FirewallPolicy pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
	IpsecPhase1 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
	IpsecPhase1Interface pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
	IpsecPhase2 pulumi.StringPtrInput
	// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
	IpsecPhase2Interface pulumi.StringPtrInput
	// Log disk quota in MB (range depends on how much disk space is available).
	LogDiskQuota pulumi.StringPtrInput
	// VDOM name.
	Name pulumi.StringPtrInput
	// Maximum guaranteed number of firewall one-time schedules.
	OnetimeSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of concurrent proxy users.
	Proxy pulumi.StringPtrInput
	// Maximum guaranteed number of firewall recurring schedules.
	RecurringSchedule pulumi.StringPtrInput
	// Maximum guaranteed number of firewall service groups.
	ServiceGroup pulumi.StringPtrInput
	// Maximum guaranteed number of sessions.
	Session pulumi.StringPtrInput
	// Permanent SNMP Index of the virtual domain (0 - 4294967295).
	SnmpIndex pulumi.IntPtrInput
	// Maximum guaranteed number of SSL-VPNs.
	Sslvpn pulumi.StringPtrInput
	// Maximum guaranteed number of local users.
	User pulumi.StringPtrInput
	// Maximum guaranteed number of user groups.
	UserGroup pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SystemVdompropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemVdompropertyArgs)(nil)).Elem()
}

type SystemVdompropertyInput interface {
	pulumi.Input

	ToSystemVdompropertyOutput() SystemVdompropertyOutput
	ToSystemVdompropertyOutputWithContext(ctx context.Context) SystemVdompropertyOutput
}

func (*SystemVdomproperty) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomproperty)(nil)).Elem()
}

func (i *SystemVdomproperty) ToSystemVdompropertyOutput() SystemVdompropertyOutput {
	return i.ToSystemVdompropertyOutputWithContext(context.Background())
}

func (i *SystemVdomproperty) ToSystemVdompropertyOutputWithContext(ctx context.Context) SystemVdompropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdompropertyOutput)
}

// SystemVdompropertyArrayInput is an input type that accepts SystemVdompropertyArray and SystemVdompropertyArrayOutput values.
// You can construct a concrete instance of `SystemVdompropertyArrayInput` via:
//
//	SystemVdompropertyArray{ SystemVdompropertyArgs{...} }
type SystemVdompropertyArrayInput interface {
	pulumi.Input

	ToSystemVdompropertyArrayOutput() SystemVdompropertyArrayOutput
	ToSystemVdompropertyArrayOutputWithContext(context.Context) SystemVdompropertyArrayOutput
}

type SystemVdompropertyArray []SystemVdompropertyInput

func (SystemVdompropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomproperty)(nil)).Elem()
}

func (i SystemVdompropertyArray) ToSystemVdompropertyArrayOutput() SystemVdompropertyArrayOutput {
	return i.ToSystemVdompropertyArrayOutputWithContext(context.Background())
}

func (i SystemVdompropertyArray) ToSystemVdompropertyArrayOutputWithContext(ctx context.Context) SystemVdompropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdompropertyArrayOutput)
}

// SystemVdompropertyMapInput is an input type that accepts SystemVdompropertyMap and SystemVdompropertyMapOutput values.
// You can construct a concrete instance of `SystemVdompropertyMapInput` via:
//
//	SystemVdompropertyMap{ "key": SystemVdompropertyArgs{...} }
type SystemVdompropertyMapInput interface {
	pulumi.Input

	ToSystemVdompropertyMapOutput() SystemVdompropertyMapOutput
	ToSystemVdompropertyMapOutputWithContext(context.Context) SystemVdompropertyMapOutput
}

type SystemVdompropertyMap map[string]SystemVdompropertyInput

func (SystemVdompropertyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomproperty)(nil)).Elem()
}

func (i SystemVdompropertyMap) ToSystemVdompropertyMapOutput() SystemVdompropertyMapOutput {
	return i.ToSystemVdompropertyMapOutputWithContext(context.Background())
}

func (i SystemVdompropertyMap) ToSystemVdompropertyMapOutputWithContext(ctx context.Context) SystemVdompropertyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemVdompropertyMapOutput)
}

type SystemVdompropertyOutput struct{ *pulumi.OutputState }

func (SystemVdompropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemVdomproperty)(nil)).Elem()
}

func (o SystemVdompropertyOutput) ToSystemVdompropertyOutput() SystemVdompropertyOutput {
	return o
}

func (o SystemVdompropertyOutput) ToSystemVdompropertyOutputWithContext(ctx context.Context) SystemVdompropertyOutput {
	return o
}

// Maximum guaranteed number of firewall custom services.
func (o SystemVdompropertyOutput) CustomService() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.CustomService }).(pulumi.StringOutput)
}

// Description.
func (o SystemVdompropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Maximum guaranteed number of dial-up tunnels.
func (o SystemVdompropertyOutput) DialupTunnel() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.DialupTunnel }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
func (o SystemVdompropertyOutput) FirewallAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.FirewallAddress }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall address groups (IPv4, IPv6).
func (o SystemVdompropertyOutput) FirewallAddrgrp() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.FirewallAddrgrp }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
func (o SystemVdompropertyOutput) FirewallPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.FirewallPolicy }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase 1 tunnels.
func (o SystemVdompropertyOutput) IpsecPhase1() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.IpsecPhase1 }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
func (o SystemVdompropertyOutput) IpsecPhase1Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.IpsecPhase1Interface }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase 2 tunnels.
func (o SystemVdompropertyOutput) IpsecPhase2() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.IpsecPhase2 }).(pulumi.StringOutput)
}

// Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
func (o SystemVdompropertyOutput) IpsecPhase2Interface() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.IpsecPhase2Interface }).(pulumi.StringOutput)
}

// Log disk quota in MB (range depends on how much disk space is available).
func (o SystemVdompropertyOutput) LogDiskQuota() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.LogDiskQuota }).(pulumi.StringOutput)
}

// VDOM name.
func (o SystemVdompropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall one-time schedules.
func (o SystemVdompropertyOutput) OnetimeSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.OnetimeSchedule }).(pulumi.StringOutput)
}

// Maximum guaranteed number of concurrent proxy users.
func (o SystemVdompropertyOutput) Proxy() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.Proxy }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall recurring schedules.
func (o SystemVdompropertyOutput) RecurringSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.RecurringSchedule }).(pulumi.StringOutput)
}

// Maximum guaranteed number of firewall service groups.
func (o SystemVdompropertyOutput) ServiceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.ServiceGroup }).(pulumi.StringOutput)
}

// Maximum guaranteed number of sessions.
func (o SystemVdompropertyOutput) Session() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.Session }).(pulumi.StringOutput)
}

// Permanent SNMP Index of the virtual domain (0 - 4294967295).
func (o SystemVdompropertyOutput) SnmpIndex() pulumi.IntOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.IntOutput { return v.SnmpIndex }).(pulumi.IntOutput)
}

// Maximum guaranteed number of SSL-VPNs.
func (o SystemVdompropertyOutput) Sslvpn() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.Sslvpn }).(pulumi.StringOutput)
}

// Maximum guaranteed number of local users.
func (o SystemVdompropertyOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

// Maximum guaranteed number of user groups.
func (o SystemVdompropertyOutput) UserGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringOutput { return v.UserGroup }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SystemVdompropertyOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemVdomproperty) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SystemVdompropertyArrayOutput struct{ *pulumi.OutputState }

func (SystemVdompropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SystemVdomproperty)(nil)).Elem()
}

func (o SystemVdompropertyArrayOutput) ToSystemVdompropertyArrayOutput() SystemVdompropertyArrayOutput {
	return o
}

func (o SystemVdompropertyArrayOutput) ToSystemVdompropertyArrayOutputWithContext(ctx context.Context) SystemVdompropertyArrayOutput {
	return o
}

func (o SystemVdompropertyArrayOutput) Index(i pulumi.IntInput) SystemVdompropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SystemVdomproperty {
		return vs[0].([]*SystemVdomproperty)[vs[1].(int)]
	}).(SystemVdompropertyOutput)
}

type SystemVdompropertyMapOutput struct{ *pulumi.OutputState }

func (SystemVdompropertyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SystemVdomproperty)(nil)).Elem()
}

func (o SystemVdompropertyMapOutput) ToSystemVdompropertyMapOutput() SystemVdompropertyMapOutput {
	return o
}

func (o SystemVdompropertyMapOutput) ToSystemVdompropertyMapOutputWithContext(ctx context.Context) SystemVdompropertyMapOutput {
	return o
}

func (o SystemVdompropertyMapOutput) MapIndex(k pulumi.StringInput) SystemVdompropertyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SystemVdomproperty {
		return vs[0].(map[string]*SystemVdomproperty)[vs[1].(string)]
	}).(SystemVdompropertyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdompropertyInput)(nil)).Elem(), &SystemVdomproperty{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdompropertyArrayInput)(nil)).Elem(), SystemVdompropertyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemVdompropertyMapInput)(nil)).Elem(), SystemVdompropertyMap{})
	pulumi.RegisterOutputType(SystemVdompropertyOutput{})
	pulumi.RegisterOutputType(SystemVdompropertyArrayOutput{})
	pulumi.RegisterOutputType(SystemVdompropertyMapOutput{})
}