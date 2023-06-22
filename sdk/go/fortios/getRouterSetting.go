// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router setting
func LookupRouterSetting(ctx *pulumi.Context, args *LookupRouterSettingArgs, opts ...pulumi.InvokeOption) (*LookupRouterSettingResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterSettingResult
	err := ctx.Invoke("fortios:index/getRouterSetting:getRouterSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterSetting.
type LookupRouterSettingArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterSetting.
type LookupRouterSettingResult struct {
	// bgp_debug_flags
	BgpDebugFlags string `pulumi:"bgpDebugFlags"`
	// Hostname for this virtual domain router.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// igmp_debug_flags
	IgmpDebugFlags string `pulumi:"igmpDebugFlags"`
	// imi_debug_flags
	ImiDebugFlags string `pulumi:"imiDebugFlags"`
	// isis_debug_flags
	IsisDebugFlags string `pulumi:"isisDebugFlags"`
	// ospf6_debug_events_flags
	Ospf6DebugEventsFlags string `pulumi:"ospf6DebugEventsFlags"`
	// ospf6_debug_ifsm_flags
	Ospf6DebugIfsmFlags string `pulumi:"ospf6DebugIfsmFlags"`
	// ospf6_debug_lsa_flags
	Ospf6DebugLsaFlags string `pulumi:"ospf6DebugLsaFlags"`
	// ospf6_debug_nfsm_flags
	Ospf6DebugNfsmFlags string `pulumi:"ospf6DebugNfsmFlags"`
	// ospf6_debug_nsm_flags
	Ospf6DebugNsmFlags string `pulumi:"ospf6DebugNsmFlags"`
	// ospf6_debug_packet_flags
	Ospf6DebugPacketFlags string `pulumi:"ospf6DebugPacketFlags"`
	// ospf6_debug_route_flags
	Ospf6DebugRouteFlags string `pulumi:"ospf6DebugRouteFlags"`
	// ospf_debug_events_flags
	OspfDebugEventsFlags string `pulumi:"ospfDebugEventsFlags"`
	// ospf_debug_ifsm_flags
	OspfDebugIfsmFlags string `pulumi:"ospfDebugIfsmFlags"`
	// ospf_debug_lsa_flags
	OspfDebugLsaFlags string `pulumi:"ospfDebugLsaFlags"`
	// ospf_debug_nfsm_flags
	OspfDebugNfsmFlags string `pulumi:"ospfDebugNfsmFlags"`
	// ospf_debug_nsm_flags
	OspfDebugNsmFlags string `pulumi:"ospfDebugNsmFlags"`
	// ospf_debug_packet_flags
	OspfDebugPacketFlags string `pulumi:"ospfDebugPacketFlags"`
	// ospf_debug_route_flags
	OspfDebugRouteFlags string `pulumi:"ospfDebugRouteFlags"`
	// pimdm_debug_flags
	PimdmDebugFlags string `pulumi:"pimdmDebugFlags"`
	// pimsm_debug_joinprune_flags
	PimsmDebugJoinpruneFlags string `pulumi:"pimsmDebugJoinpruneFlags"`
	// pimsm_debug_simple_flags
	PimsmDebugSimpleFlags string `pulumi:"pimsmDebugSimpleFlags"`
	// pimsm_debug_timer_flags
	PimsmDebugTimerFlags string `pulumi:"pimsmDebugTimerFlags"`
	// rip_debug_flags
	RipDebugFlags string `pulumi:"ripDebugFlags"`
	// ripng_debug_flags
	RipngDebugFlags string `pulumi:"ripngDebugFlags"`
	// Prefix-list as filter for showing routes.
	ShowFilter string  `pulumi:"showFilter"`
	Vdomparam  *string `pulumi:"vdomparam"`
}

func LookupRouterSettingOutput(ctx *pulumi.Context, args LookupRouterSettingOutputArgs, opts ...pulumi.InvokeOption) LookupRouterSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterSettingResult, error) {
			args := v.(LookupRouterSettingArgs)
			r, err := LookupRouterSetting(ctx, &args, opts...)
			var s LookupRouterSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterSettingResultOutput)
}

// A collection of arguments for invoking getRouterSetting.
type LookupRouterSettingOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterSettingArgs)(nil)).Elem()
}

// A collection of values returned by getRouterSetting.
type LookupRouterSettingResultOutput struct{ *pulumi.OutputState }

func (LookupRouterSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterSettingResult)(nil)).Elem()
}

func (o LookupRouterSettingResultOutput) ToLookupRouterSettingResultOutput() LookupRouterSettingResultOutput {
	return o
}

func (o LookupRouterSettingResultOutput) ToLookupRouterSettingResultOutputWithContext(ctx context.Context) LookupRouterSettingResultOutput {
	return o
}

// bgp_debug_flags
func (o LookupRouterSettingResultOutput) BgpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.BgpDebugFlags }).(pulumi.StringOutput)
}

// Hostname for this virtual domain router.
func (o LookupRouterSettingResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// igmp_debug_flags
func (o LookupRouterSettingResultOutput) IgmpDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.IgmpDebugFlags }).(pulumi.StringOutput)
}

// imi_debug_flags
func (o LookupRouterSettingResultOutput) ImiDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.ImiDebugFlags }).(pulumi.StringOutput)
}

// isis_debug_flags
func (o LookupRouterSettingResultOutput) IsisDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.IsisDebugFlags }).(pulumi.StringOutput)
}

// ospf6_debug_events_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugEventsFlags }).(pulumi.StringOutput)
}

// ospf6_debug_ifsm_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugIfsmFlags }).(pulumi.StringOutput)
}

// ospf6_debug_lsa_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugLsaFlags }).(pulumi.StringOutput)
}

// ospf6_debug_nfsm_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugNfsmFlags }).(pulumi.StringOutput)
}

// ospf6_debug_nsm_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugNsmFlags }).(pulumi.StringOutput)
}

// ospf6_debug_packet_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugPacketFlags }).(pulumi.StringOutput)
}

// ospf6_debug_route_flags
func (o LookupRouterSettingResultOutput) Ospf6DebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.Ospf6DebugRouteFlags }).(pulumi.StringOutput)
}

// ospf_debug_events_flags
func (o LookupRouterSettingResultOutput) OspfDebugEventsFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugEventsFlags }).(pulumi.StringOutput)
}

// ospf_debug_ifsm_flags
func (o LookupRouterSettingResultOutput) OspfDebugIfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugIfsmFlags }).(pulumi.StringOutput)
}

// ospf_debug_lsa_flags
func (o LookupRouterSettingResultOutput) OspfDebugLsaFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugLsaFlags }).(pulumi.StringOutput)
}

// ospf_debug_nfsm_flags
func (o LookupRouterSettingResultOutput) OspfDebugNfsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugNfsmFlags }).(pulumi.StringOutput)
}

// ospf_debug_nsm_flags
func (o LookupRouterSettingResultOutput) OspfDebugNsmFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugNsmFlags }).(pulumi.StringOutput)
}

// ospf_debug_packet_flags
func (o LookupRouterSettingResultOutput) OspfDebugPacketFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugPacketFlags }).(pulumi.StringOutput)
}

// ospf_debug_route_flags
func (o LookupRouterSettingResultOutput) OspfDebugRouteFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.OspfDebugRouteFlags }).(pulumi.StringOutput)
}

// pimdm_debug_flags
func (o LookupRouterSettingResultOutput) PimdmDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimdmDebugFlags }).(pulumi.StringOutput)
}

// pimsm_debug_joinprune_flags
func (o LookupRouterSettingResultOutput) PimsmDebugJoinpruneFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimsmDebugJoinpruneFlags }).(pulumi.StringOutput)
}

// pimsm_debug_simple_flags
func (o LookupRouterSettingResultOutput) PimsmDebugSimpleFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimsmDebugSimpleFlags }).(pulumi.StringOutput)
}

// pimsm_debug_timer_flags
func (o LookupRouterSettingResultOutput) PimsmDebugTimerFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.PimsmDebugTimerFlags }).(pulumi.StringOutput)
}

// rip_debug_flags
func (o LookupRouterSettingResultOutput) RipDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.RipDebugFlags }).(pulumi.StringOutput)
}

// ripng_debug_flags
func (o LookupRouterSettingResultOutput) RipngDebugFlags() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.RipngDebugFlags }).(pulumi.StringOutput)
}

// Prefix-list as filter for showing routes.
func (o LookupRouterSettingResultOutput) ShowFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) string { return v.ShowFilter }).(pulumi.StringOutput)
}

func (o LookupRouterSettingResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterSettingResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterSettingResultOutput{})
}