# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRouterSettingResult',
    'AwaitableGetRouterSettingResult',
    'get_router_setting',
    'get_router_setting_output',
]

@pulumi.output_type
class GetRouterSettingResult:
    """
    A collection of values returned by getRouterSetting.
    """
    def __init__(__self__, bgp_debug_flags=None, hostname=None, id=None, igmp_debug_flags=None, imi_debug_flags=None, isis_debug_flags=None, ospf6_debug_events_flags=None, ospf6_debug_ifsm_flags=None, ospf6_debug_lsa_flags=None, ospf6_debug_nfsm_flags=None, ospf6_debug_nsm_flags=None, ospf6_debug_packet_flags=None, ospf6_debug_route_flags=None, ospf_debug_events_flags=None, ospf_debug_ifsm_flags=None, ospf_debug_lsa_flags=None, ospf_debug_nfsm_flags=None, ospf_debug_nsm_flags=None, ospf_debug_packet_flags=None, ospf_debug_route_flags=None, pimdm_debug_flags=None, pimsm_debug_joinprune_flags=None, pimsm_debug_simple_flags=None, pimsm_debug_timer_flags=None, rip_debug_flags=None, ripng_debug_flags=None, show_filter=None, vdomparam=None):
        if bgp_debug_flags and not isinstance(bgp_debug_flags, str):
            raise TypeError("Expected argument 'bgp_debug_flags' to be a str")
        pulumi.set(__self__, "bgp_debug_flags", bgp_debug_flags)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if igmp_debug_flags and not isinstance(igmp_debug_flags, str):
            raise TypeError("Expected argument 'igmp_debug_flags' to be a str")
        pulumi.set(__self__, "igmp_debug_flags", igmp_debug_flags)
        if imi_debug_flags and not isinstance(imi_debug_flags, str):
            raise TypeError("Expected argument 'imi_debug_flags' to be a str")
        pulumi.set(__self__, "imi_debug_flags", imi_debug_flags)
        if isis_debug_flags and not isinstance(isis_debug_flags, str):
            raise TypeError("Expected argument 'isis_debug_flags' to be a str")
        pulumi.set(__self__, "isis_debug_flags", isis_debug_flags)
        if ospf6_debug_events_flags and not isinstance(ospf6_debug_events_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_events_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_events_flags", ospf6_debug_events_flags)
        if ospf6_debug_ifsm_flags and not isinstance(ospf6_debug_ifsm_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_ifsm_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_ifsm_flags", ospf6_debug_ifsm_flags)
        if ospf6_debug_lsa_flags and not isinstance(ospf6_debug_lsa_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_lsa_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_lsa_flags", ospf6_debug_lsa_flags)
        if ospf6_debug_nfsm_flags and not isinstance(ospf6_debug_nfsm_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_nfsm_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_nfsm_flags", ospf6_debug_nfsm_flags)
        if ospf6_debug_nsm_flags and not isinstance(ospf6_debug_nsm_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_nsm_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_nsm_flags", ospf6_debug_nsm_flags)
        if ospf6_debug_packet_flags and not isinstance(ospf6_debug_packet_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_packet_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_packet_flags", ospf6_debug_packet_flags)
        if ospf6_debug_route_flags and not isinstance(ospf6_debug_route_flags, str):
            raise TypeError("Expected argument 'ospf6_debug_route_flags' to be a str")
        pulumi.set(__self__, "ospf6_debug_route_flags", ospf6_debug_route_flags)
        if ospf_debug_events_flags and not isinstance(ospf_debug_events_flags, str):
            raise TypeError("Expected argument 'ospf_debug_events_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_events_flags", ospf_debug_events_flags)
        if ospf_debug_ifsm_flags and not isinstance(ospf_debug_ifsm_flags, str):
            raise TypeError("Expected argument 'ospf_debug_ifsm_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_ifsm_flags", ospf_debug_ifsm_flags)
        if ospf_debug_lsa_flags and not isinstance(ospf_debug_lsa_flags, str):
            raise TypeError("Expected argument 'ospf_debug_lsa_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_lsa_flags", ospf_debug_lsa_flags)
        if ospf_debug_nfsm_flags and not isinstance(ospf_debug_nfsm_flags, str):
            raise TypeError("Expected argument 'ospf_debug_nfsm_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_nfsm_flags", ospf_debug_nfsm_flags)
        if ospf_debug_nsm_flags and not isinstance(ospf_debug_nsm_flags, str):
            raise TypeError("Expected argument 'ospf_debug_nsm_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_nsm_flags", ospf_debug_nsm_flags)
        if ospf_debug_packet_flags and not isinstance(ospf_debug_packet_flags, str):
            raise TypeError("Expected argument 'ospf_debug_packet_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_packet_flags", ospf_debug_packet_flags)
        if ospf_debug_route_flags and not isinstance(ospf_debug_route_flags, str):
            raise TypeError("Expected argument 'ospf_debug_route_flags' to be a str")
        pulumi.set(__self__, "ospf_debug_route_flags", ospf_debug_route_flags)
        if pimdm_debug_flags and not isinstance(pimdm_debug_flags, str):
            raise TypeError("Expected argument 'pimdm_debug_flags' to be a str")
        pulumi.set(__self__, "pimdm_debug_flags", pimdm_debug_flags)
        if pimsm_debug_joinprune_flags and not isinstance(pimsm_debug_joinprune_flags, str):
            raise TypeError("Expected argument 'pimsm_debug_joinprune_flags' to be a str")
        pulumi.set(__self__, "pimsm_debug_joinprune_flags", pimsm_debug_joinprune_flags)
        if pimsm_debug_simple_flags and not isinstance(pimsm_debug_simple_flags, str):
            raise TypeError("Expected argument 'pimsm_debug_simple_flags' to be a str")
        pulumi.set(__self__, "pimsm_debug_simple_flags", pimsm_debug_simple_flags)
        if pimsm_debug_timer_flags and not isinstance(pimsm_debug_timer_flags, str):
            raise TypeError("Expected argument 'pimsm_debug_timer_flags' to be a str")
        pulumi.set(__self__, "pimsm_debug_timer_flags", pimsm_debug_timer_flags)
        if rip_debug_flags and not isinstance(rip_debug_flags, str):
            raise TypeError("Expected argument 'rip_debug_flags' to be a str")
        pulumi.set(__self__, "rip_debug_flags", rip_debug_flags)
        if ripng_debug_flags and not isinstance(ripng_debug_flags, str):
            raise TypeError("Expected argument 'ripng_debug_flags' to be a str")
        pulumi.set(__self__, "ripng_debug_flags", ripng_debug_flags)
        if show_filter and not isinstance(show_filter, str):
            raise TypeError("Expected argument 'show_filter' to be a str")
        pulumi.set(__self__, "show_filter", show_filter)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="bgpDebugFlags")
    def bgp_debug_flags(self) -> str:
        """
        bgp_debug_flags
        """
        return pulumi.get(self, "bgp_debug_flags")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        Hostname for this virtual domain router.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="igmpDebugFlags")
    def igmp_debug_flags(self) -> str:
        """
        igmp_debug_flags
        """
        return pulumi.get(self, "igmp_debug_flags")

    @property
    @pulumi.getter(name="imiDebugFlags")
    def imi_debug_flags(self) -> str:
        """
        imi_debug_flags
        """
        return pulumi.get(self, "imi_debug_flags")

    @property
    @pulumi.getter(name="isisDebugFlags")
    def isis_debug_flags(self) -> str:
        """
        isis_debug_flags
        """
        return pulumi.get(self, "isis_debug_flags")

    @property
    @pulumi.getter(name="ospf6DebugEventsFlags")
    def ospf6_debug_events_flags(self) -> str:
        """
        ospf6_debug_events_flags
        """
        return pulumi.get(self, "ospf6_debug_events_flags")

    @property
    @pulumi.getter(name="ospf6DebugIfsmFlags")
    def ospf6_debug_ifsm_flags(self) -> str:
        """
        ospf6_debug_ifsm_flags
        """
        return pulumi.get(self, "ospf6_debug_ifsm_flags")

    @property
    @pulumi.getter(name="ospf6DebugLsaFlags")
    def ospf6_debug_lsa_flags(self) -> str:
        """
        ospf6_debug_lsa_flags
        """
        return pulumi.get(self, "ospf6_debug_lsa_flags")

    @property
    @pulumi.getter(name="ospf6DebugNfsmFlags")
    def ospf6_debug_nfsm_flags(self) -> str:
        """
        ospf6_debug_nfsm_flags
        """
        return pulumi.get(self, "ospf6_debug_nfsm_flags")

    @property
    @pulumi.getter(name="ospf6DebugNsmFlags")
    def ospf6_debug_nsm_flags(self) -> str:
        """
        ospf6_debug_nsm_flags
        """
        return pulumi.get(self, "ospf6_debug_nsm_flags")

    @property
    @pulumi.getter(name="ospf6DebugPacketFlags")
    def ospf6_debug_packet_flags(self) -> str:
        """
        ospf6_debug_packet_flags
        """
        return pulumi.get(self, "ospf6_debug_packet_flags")

    @property
    @pulumi.getter(name="ospf6DebugRouteFlags")
    def ospf6_debug_route_flags(self) -> str:
        """
        ospf6_debug_route_flags
        """
        return pulumi.get(self, "ospf6_debug_route_flags")

    @property
    @pulumi.getter(name="ospfDebugEventsFlags")
    def ospf_debug_events_flags(self) -> str:
        """
        ospf_debug_events_flags
        """
        return pulumi.get(self, "ospf_debug_events_flags")

    @property
    @pulumi.getter(name="ospfDebugIfsmFlags")
    def ospf_debug_ifsm_flags(self) -> str:
        """
        ospf_debug_ifsm_flags
        """
        return pulumi.get(self, "ospf_debug_ifsm_flags")

    @property
    @pulumi.getter(name="ospfDebugLsaFlags")
    def ospf_debug_lsa_flags(self) -> str:
        """
        ospf_debug_lsa_flags
        """
        return pulumi.get(self, "ospf_debug_lsa_flags")

    @property
    @pulumi.getter(name="ospfDebugNfsmFlags")
    def ospf_debug_nfsm_flags(self) -> str:
        """
        ospf_debug_nfsm_flags
        """
        return pulumi.get(self, "ospf_debug_nfsm_flags")

    @property
    @pulumi.getter(name="ospfDebugNsmFlags")
    def ospf_debug_nsm_flags(self) -> str:
        """
        ospf_debug_nsm_flags
        """
        return pulumi.get(self, "ospf_debug_nsm_flags")

    @property
    @pulumi.getter(name="ospfDebugPacketFlags")
    def ospf_debug_packet_flags(self) -> str:
        """
        ospf_debug_packet_flags
        """
        return pulumi.get(self, "ospf_debug_packet_flags")

    @property
    @pulumi.getter(name="ospfDebugRouteFlags")
    def ospf_debug_route_flags(self) -> str:
        """
        ospf_debug_route_flags
        """
        return pulumi.get(self, "ospf_debug_route_flags")

    @property
    @pulumi.getter(name="pimdmDebugFlags")
    def pimdm_debug_flags(self) -> str:
        """
        pimdm_debug_flags
        """
        return pulumi.get(self, "pimdm_debug_flags")

    @property
    @pulumi.getter(name="pimsmDebugJoinpruneFlags")
    def pimsm_debug_joinprune_flags(self) -> str:
        """
        pimsm_debug_joinprune_flags
        """
        return pulumi.get(self, "pimsm_debug_joinprune_flags")

    @property
    @pulumi.getter(name="pimsmDebugSimpleFlags")
    def pimsm_debug_simple_flags(self) -> str:
        """
        pimsm_debug_simple_flags
        """
        return pulumi.get(self, "pimsm_debug_simple_flags")

    @property
    @pulumi.getter(name="pimsmDebugTimerFlags")
    def pimsm_debug_timer_flags(self) -> str:
        """
        pimsm_debug_timer_flags
        """
        return pulumi.get(self, "pimsm_debug_timer_flags")

    @property
    @pulumi.getter(name="ripDebugFlags")
    def rip_debug_flags(self) -> str:
        """
        rip_debug_flags
        """
        return pulumi.get(self, "rip_debug_flags")

    @property
    @pulumi.getter(name="ripngDebugFlags")
    def ripng_debug_flags(self) -> str:
        """
        ripng_debug_flags
        """
        return pulumi.get(self, "ripng_debug_flags")

    @property
    @pulumi.getter(name="showFilter")
    def show_filter(self) -> str:
        """
        Prefix-list as filter for showing routes.
        """
        return pulumi.get(self, "show_filter")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterSettingResult(GetRouterSettingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterSettingResult(
            bgp_debug_flags=self.bgp_debug_flags,
            hostname=self.hostname,
            id=self.id,
            igmp_debug_flags=self.igmp_debug_flags,
            imi_debug_flags=self.imi_debug_flags,
            isis_debug_flags=self.isis_debug_flags,
            ospf6_debug_events_flags=self.ospf6_debug_events_flags,
            ospf6_debug_ifsm_flags=self.ospf6_debug_ifsm_flags,
            ospf6_debug_lsa_flags=self.ospf6_debug_lsa_flags,
            ospf6_debug_nfsm_flags=self.ospf6_debug_nfsm_flags,
            ospf6_debug_nsm_flags=self.ospf6_debug_nsm_flags,
            ospf6_debug_packet_flags=self.ospf6_debug_packet_flags,
            ospf6_debug_route_flags=self.ospf6_debug_route_flags,
            ospf_debug_events_flags=self.ospf_debug_events_flags,
            ospf_debug_ifsm_flags=self.ospf_debug_ifsm_flags,
            ospf_debug_lsa_flags=self.ospf_debug_lsa_flags,
            ospf_debug_nfsm_flags=self.ospf_debug_nfsm_flags,
            ospf_debug_nsm_flags=self.ospf_debug_nsm_flags,
            ospf_debug_packet_flags=self.ospf_debug_packet_flags,
            ospf_debug_route_flags=self.ospf_debug_route_flags,
            pimdm_debug_flags=self.pimdm_debug_flags,
            pimsm_debug_joinprune_flags=self.pimsm_debug_joinprune_flags,
            pimsm_debug_simple_flags=self.pimsm_debug_simple_flags,
            pimsm_debug_timer_flags=self.pimsm_debug_timer_flags,
            rip_debug_flags=self.rip_debug_flags,
            ripng_debug_flags=self.ripng_debug_flags,
            show_filter=self.show_filter,
            vdomparam=self.vdomparam)


def get_router_setting(vdomparam: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterSettingResult:
    """
    Use this data source to get information on fortios router setting


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterSetting:getRouterSetting', __args__, opts=opts, typ=GetRouterSettingResult).value

    return AwaitableGetRouterSettingResult(
        bgp_debug_flags=__ret__.bgp_debug_flags,
        hostname=__ret__.hostname,
        id=__ret__.id,
        igmp_debug_flags=__ret__.igmp_debug_flags,
        imi_debug_flags=__ret__.imi_debug_flags,
        isis_debug_flags=__ret__.isis_debug_flags,
        ospf6_debug_events_flags=__ret__.ospf6_debug_events_flags,
        ospf6_debug_ifsm_flags=__ret__.ospf6_debug_ifsm_flags,
        ospf6_debug_lsa_flags=__ret__.ospf6_debug_lsa_flags,
        ospf6_debug_nfsm_flags=__ret__.ospf6_debug_nfsm_flags,
        ospf6_debug_nsm_flags=__ret__.ospf6_debug_nsm_flags,
        ospf6_debug_packet_flags=__ret__.ospf6_debug_packet_flags,
        ospf6_debug_route_flags=__ret__.ospf6_debug_route_flags,
        ospf_debug_events_flags=__ret__.ospf_debug_events_flags,
        ospf_debug_ifsm_flags=__ret__.ospf_debug_ifsm_flags,
        ospf_debug_lsa_flags=__ret__.ospf_debug_lsa_flags,
        ospf_debug_nfsm_flags=__ret__.ospf_debug_nfsm_flags,
        ospf_debug_nsm_flags=__ret__.ospf_debug_nsm_flags,
        ospf_debug_packet_flags=__ret__.ospf_debug_packet_flags,
        ospf_debug_route_flags=__ret__.ospf_debug_route_flags,
        pimdm_debug_flags=__ret__.pimdm_debug_flags,
        pimsm_debug_joinprune_flags=__ret__.pimsm_debug_joinprune_flags,
        pimsm_debug_simple_flags=__ret__.pimsm_debug_simple_flags,
        pimsm_debug_timer_flags=__ret__.pimsm_debug_timer_flags,
        rip_debug_flags=__ret__.rip_debug_flags,
        ripng_debug_flags=__ret__.ripng_debug_flags,
        show_filter=__ret__.show_filter,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_setting)
def get_router_setting_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterSettingResult]:
    """
    Use this data source to get information on fortios router setting


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
