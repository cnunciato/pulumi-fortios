// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios system gretunnel
func LookupSystemGretunnel(ctx *pulumi.Context, args *LookupSystemGretunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemGretunnelResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemGretunnelResult
	err := ctx.Invoke("fortios:index/getSystemGretunnel:getSystemGretunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemGretunnel.
type LookupSystemGretunnelArgs struct {
	// Specify the name of the desired system gretunnel.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemGretunnel.
type LookupSystemGretunnelResult struct {
	// Enable/disable validating checksums in received GRE packets.
	ChecksumReception string `pulumi:"checksumReception"`
	// Enable/disable including checksums in transmitted GRE packets.
	ChecksumTransmission string `pulumi:"checksumTransmission"`
	// DiffServ setting to be applied to GRE tunnel outer IP header.
	Diffservcode string `pulumi:"diffservcode"`
	// Enable/disable DSCP copying.
	DscpCopying string `pulumi:"dscpCopying"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Interface name.
	Interface string `pulumi:"interface"`
	// IP version to use for VPN interface.
	IpVersion string `pulumi:"ipVersion"`
	// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
	KeepaliveFailtimes int `pulumi:"keepaliveFailtimes"`
	// Keepalive message interval (0 - 32767, 0 = disabled).
	KeepaliveInterval int `pulumi:"keepaliveInterval"`
	// Require received GRE packets contain this key (0 - 4294967295).
	KeyInbound int `pulumi:"keyInbound"`
	// Include this key in transmitted GRE packets (0 - 4294967295).
	KeyOutbound int `pulumi:"keyOutbound"`
	// IP address of the local gateway.
	LocalGw string `pulumi:"localGw"`
	// IPv6 address of the local gateway.
	LocalGw6 string `pulumi:"localGw6"`
	// Tunnel name.
	Name string `pulumi:"name"`
	// IP address of the remote gateway.
	RemoteGw string `pulumi:"remoteGw"`
	// IPv6 address of the remote gateway.
	RemoteGw6 string `pulumi:"remoteGw6"`
	// Enable/disable validating sequence numbers in received GRE packets.
	SequenceNumberReception string `pulumi:"sequenceNumberReception"`
	// Enable/disable including of sequence numbers in transmitted GRE packets.
	SequenceNumberTransmission string `pulumi:"sequenceNumberTransmission"`
	// Enable/disable use of SD-WAN to reach remote gateway.
	UseSdwan  string  `pulumi:"useSdwan"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemGretunnelOutput(ctx *pulumi.Context, args LookupSystemGretunnelOutputArgs, opts ...pulumi.InvokeOption) LookupSystemGretunnelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemGretunnelResult, error) {
			args := v.(LookupSystemGretunnelArgs)
			r, err := LookupSystemGretunnel(ctx, &args, opts...)
			var s LookupSystemGretunnelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemGretunnelResultOutput)
}

// A collection of arguments for invoking getSystemGretunnel.
type LookupSystemGretunnelOutputArgs struct {
	// Specify the name of the desired system gretunnel.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemGretunnelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemGretunnelArgs)(nil)).Elem()
}

// A collection of values returned by getSystemGretunnel.
type LookupSystemGretunnelResultOutput struct{ *pulumi.OutputState }

func (LookupSystemGretunnelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemGretunnelResult)(nil)).Elem()
}

func (o LookupSystemGretunnelResultOutput) ToLookupSystemGretunnelResultOutput() LookupSystemGretunnelResultOutput {
	return o
}

func (o LookupSystemGretunnelResultOutput) ToLookupSystemGretunnelResultOutputWithContext(ctx context.Context) LookupSystemGretunnelResultOutput {
	return o
}

// Enable/disable validating checksums in received GRE packets.
func (o LookupSystemGretunnelResultOutput) ChecksumReception() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.ChecksumReception }).(pulumi.StringOutput)
}

// Enable/disable including checksums in transmitted GRE packets.
func (o LookupSystemGretunnelResultOutput) ChecksumTransmission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.ChecksumTransmission }).(pulumi.StringOutput)
}

// DiffServ setting to be applied to GRE tunnel outer IP header.
func (o LookupSystemGretunnelResultOutput) Diffservcode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.Diffservcode }).(pulumi.StringOutput)
}

// Enable/disable DSCP copying.
func (o LookupSystemGretunnelResultOutput) DscpCopying() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.DscpCopying }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemGretunnelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface name.
func (o LookupSystemGretunnelResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.Interface }).(pulumi.StringOutput)
}

// IP version to use for VPN interface.
func (o LookupSystemGretunnelResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

// Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
func (o LookupSystemGretunnelResultOutput) KeepaliveFailtimes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) int { return v.KeepaliveFailtimes }).(pulumi.IntOutput)
}

// Keepalive message interval (0 - 32767, 0 = disabled).
func (o LookupSystemGretunnelResultOutput) KeepaliveInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) int { return v.KeepaliveInterval }).(pulumi.IntOutput)
}

// Require received GRE packets contain this key (0 - 4294967295).
func (o LookupSystemGretunnelResultOutput) KeyInbound() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) int { return v.KeyInbound }).(pulumi.IntOutput)
}

// Include this key in transmitted GRE packets (0 - 4294967295).
func (o LookupSystemGretunnelResultOutput) KeyOutbound() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) int { return v.KeyOutbound }).(pulumi.IntOutput)
}

// IP address of the local gateway.
func (o LookupSystemGretunnelResultOutput) LocalGw() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.LocalGw }).(pulumi.StringOutput)
}

// IPv6 address of the local gateway.
func (o LookupSystemGretunnelResultOutput) LocalGw6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.LocalGw6 }).(pulumi.StringOutput)
}

// Tunnel name.
func (o LookupSystemGretunnelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.Name }).(pulumi.StringOutput)
}

// IP address of the remote gateway.
func (o LookupSystemGretunnelResultOutput) RemoteGw() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.RemoteGw }).(pulumi.StringOutput)
}

// IPv6 address of the remote gateway.
func (o LookupSystemGretunnelResultOutput) RemoteGw6() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.RemoteGw6 }).(pulumi.StringOutput)
}

// Enable/disable validating sequence numbers in received GRE packets.
func (o LookupSystemGretunnelResultOutput) SequenceNumberReception() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.SequenceNumberReception }).(pulumi.StringOutput)
}

// Enable/disable including of sequence numbers in transmitted GRE packets.
func (o LookupSystemGretunnelResultOutput) SequenceNumberTransmission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.SequenceNumberTransmission }).(pulumi.StringOutput)
}

// Enable/disable use of SD-WAN to reach remote gateway.
func (o LookupSystemGretunnelResultOutput) UseSdwan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) string { return v.UseSdwan }).(pulumi.StringOutput)
}

func (o LookupSystemGretunnelResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemGretunnelResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemGretunnelResultOutput{})
}
