// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system dnsdatabase
func LookupSystemDnsdatabase(ctx *pulumi.Context, args *LookupSystemDnsdatabaseArgs, opts ...pulumi.InvokeOption) (*LookupSystemDnsdatabaseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemDnsdatabaseResult
	err := ctx.Invoke("fortios:index/getSystemDnsdatabase:getSystemDnsdatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemDnsdatabase.
type LookupSystemDnsdatabaseArgs struct {
	// Specify the name of the desired system dnsdatabase.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemDnsdatabase.
type LookupSystemDnsdatabaseResult struct {
	// DNS zone transfer IP address list.
	AllowTransfer string `pulumi:"allowTransfer"`
	// Enable/disable authoritative zone.
	Authoritative string `pulumi:"authoritative"`
	// Email address of the administrator for this zone.
	// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
	// When using a simple username, the domain of the email will be this zone.
	Contact string `pulumi:"contact"`
	// DNS entry. The structure of `dnsEntry` block is documented below.
	DnsEntries []GetSystemDnsdatabaseDnsEntry `pulumi:"dnsEntries"`
	// Domain name.
	Domain string `pulumi:"domain"`
	// DNS zone forwarder IP address list.
	Forwarder string `pulumi:"forwarder"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
	IpMaster string `pulumi:"ipMaster"`
	// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
	IpPrimary string `pulumi:"ipPrimary"`
	// Zone name.
	Name string `pulumi:"name"`
	// Domain name of the default DNS server for this zone.
	PrimaryName string `pulumi:"primaryName"`
	// Maximum number of resource records (10 - 65536, 0 means infinite).
	RrMax int `pulumi:"rrMax"`
	// Source IP for forwarding to DNS server.
	SourceIp string `pulumi:"sourceIp"`
	// Enable/disable resource record status.
	Status string `pulumi:"status"`
	// Time-to-live for this entry (0 to 2147483647 sec, default = 0).
	Ttl int `pulumi:"ttl"`
	// Resource record type.
	Type      string  `pulumi:"type"`
	Vdomparam *string `pulumi:"vdomparam"`
	// Zone view (public to serve public clients, shadow to serve internal clients).
	View string `pulumi:"view"`
}

func LookupSystemDnsdatabaseOutput(ctx *pulumi.Context, args LookupSystemDnsdatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupSystemDnsdatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemDnsdatabaseResult, error) {
			args := v.(LookupSystemDnsdatabaseArgs)
			r, err := LookupSystemDnsdatabase(ctx, &args, opts...)
			var s LookupSystemDnsdatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemDnsdatabaseResultOutput)
}

// A collection of arguments for invoking getSystemDnsdatabase.
type LookupSystemDnsdatabaseOutputArgs struct {
	// Specify the name of the desired system dnsdatabase.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemDnsdatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDnsdatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getSystemDnsdatabase.
type LookupSystemDnsdatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupSystemDnsdatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemDnsdatabaseResult)(nil)).Elem()
}

func (o LookupSystemDnsdatabaseResultOutput) ToLookupSystemDnsdatabaseResultOutput() LookupSystemDnsdatabaseResultOutput {
	return o
}

func (o LookupSystemDnsdatabaseResultOutput) ToLookupSystemDnsdatabaseResultOutputWithContext(ctx context.Context) LookupSystemDnsdatabaseResultOutput {
	return o
}

// DNS zone transfer IP address list.
func (o LookupSystemDnsdatabaseResultOutput) AllowTransfer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.AllowTransfer }).(pulumi.StringOutput)
}

// Enable/disable authoritative zone.
func (o LookupSystemDnsdatabaseResultOutput) Authoritative() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Authoritative }).(pulumi.StringOutput)
}

// Email address of the administrator for this zone.
// You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com)
// When using a simple username, the domain of the email will be this zone.
func (o LookupSystemDnsdatabaseResultOutput) Contact() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Contact }).(pulumi.StringOutput)
}

// DNS entry. The structure of `dnsEntry` block is documented below.
func (o LookupSystemDnsdatabaseResultOutput) DnsEntries() GetSystemDnsdatabaseDnsEntryArrayOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) []GetSystemDnsdatabaseDnsEntry { return v.DnsEntries }).(GetSystemDnsdatabaseDnsEntryArrayOutput)
}

// Domain name.
func (o LookupSystemDnsdatabaseResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Domain }).(pulumi.StringOutput)
}

// DNS zone forwarder IP address list.
func (o LookupSystemDnsdatabaseResultOutput) Forwarder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Forwarder }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemDnsdatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
func (o LookupSystemDnsdatabaseResultOutput) IpMaster() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.IpMaster }).(pulumi.StringOutput)
}

// IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
func (o LookupSystemDnsdatabaseResultOutput) IpPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.IpPrimary }).(pulumi.StringOutput)
}

// Zone name.
func (o LookupSystemDnsdatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Domain name of the default DNS server for this zone.
func (o LookupSystemDnsdatabaseResultOutput) PrimaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.PrimaryName }).(pulumi.StringOutput)
}

// Maximum number of resource records (10 - 65536, 0 means infinite).
func (o LookupSystemDnsdatabaseResultOutput) RrMax() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) int { return v.RrMax }).(pulumi.IntOutput)
}

// Source IP for forwarding to DNS server.
func (o LookupSystemDnsdatabaseResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable resource record status.
func (o LookupSystemDnsdatabaseResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Status }).(pulumi.StringOutput)
}

// Time-to-live for this entry (0 to 2147483647 sec, default = 0).
func (o LookupSystemDnsdatabaseResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) int { return v.Ttl }).(pulumi.IntOutput)
}

// Resource record type.
func (o LookupSystemDnsdatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSystemDnsdatabaseResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Zone view (public to serve public clients, shadow to serve internal clients).
func (o LookupSystemDnsdatabaseResultOutput) View() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemDnsdatabaseResult) string { return v.View }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemDnsdatabaseResultOutput{})
}