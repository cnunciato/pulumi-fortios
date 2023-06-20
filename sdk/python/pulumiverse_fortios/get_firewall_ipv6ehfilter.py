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
    'GetFirewallIpv6ehfilterResult',
    'AwaitableGetFirewallIpv6ehfilterResult',
    'get_firewall_ipv6ehfilter',
    'get_firewall_ipv6ehfilter_output',
]

@pulumi.output_type
class GetFirewallIpv6ehfilterResult:
    """
    A collection of values returned by getFirewallIpv6ehfilter.
    """
    def __init__(__self__, auth=None, dest_opt=None, fragment=None, hdopt_type=None, hop_opt=None, id=None, no_next=None, routing=None, routing_type=None, vdomparam=None):
        if auth and not isinstance(auth, str):
            raise TypeError("Expected argument 'auth' to be a str")
        pulumi.set(__self__, "auth", auth)
        if dest_opt and not isinstance(dest_opt, str):
            raise TypeError("Expected argument 'dest_opt' to be a str")
        pulumi.set(__self__, "dest_opt", dest_opt)
        if fragment and not isinstance(fragment, str):
            raise TypeError("Expected argument 'fragment' to be a str")
        pulumi.set(__self__, "fragment", fragment)
        if hdopt_type and not isinstance(hdopt_type, int):
            raise TypeError("Expected argument 'hdopt_type' to be a int")
        pulumi.set(__self__, "hdopt_type", hdopt_type)
        if hop_opt and not isinstance(hop_opt, str):
            raise TypeError("Expected argument 'hop_opt' to be a str")
        pulumi.set(__self__, "hop_opt", hop_opt)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if no_next and not isinstance(no_next, str):
            raise TypeError("Expected argument 'no_next' to be a str")
        pulumi.set(__self__, "no_next", no_next)
        if routing and not isinstance(routing, str):
            raise TypeError("Expected argument 'routing' to be a str")
        pulumi.set(__self__, "routing", routing)
        if routing_type and not isinstance(routing_type, int):
            raise TypeError("Expected argument 'routing_type' to be a int")
        pulumi.set(__self__, "routing_type", routing_type)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def auth(self) -> str:
        """
        Enable/disable blocking packets with the Authentication header (default = disable).
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter(name="destOpt")
    def dest_opt(self) -> str:
        """
        Enable/disable blocking packets with Destination Options headers (default = disable).
        """
        return pulumi.get(self, "dest_opt")

    @property
    @pulumi.getter
    def fragment(self) -> str:
        """
        Enable/disable blocking packets with the Fragment header (default = disable).
        """
        return pulumi.get(self, "fragment")

    @property
    @pulumi.getter(name="hdoptType")
    def hdopt_type(self) -> int:
        """
        Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
        """
        return pulumi.get(self, "hdopt_type")

    @property
    @pulumi.getter(name="hopOpt")
    def hop_opt(self) -> str:
        """
        Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
        """
        return pulumi.get(self, "hop_opt")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="noNext")
    def no_next(self) -> str:
        """
        Enable/disable blocking packets with the No Next header (default = disable)
        """
        return pulumi.get(self, "no_next")

    @property
    @pulumi.getter
    def routing(self) -> str:
        """
        Enable/disable blocking packets with Routing headers (default = enable).
        """
        return pulumi.get(self, "routing")

    @property
    @pulumi.getter(name="routingType")
    def routing_type(self) -> int:
        """
        Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
        """
        return pulumi.get(self, "routing_type")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallIpv6ehfilterResult(GetFirewallIpv6ehfilterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallIpv6ehfilterResult(
            auth=self.auth,
            dest_opt=self.dest_opt,
            fragment=self.fragment,
            hdopt_type=self.hdopt_type,
            hop_opt=self.hop_opt,
            id=self.id,
            no_next=self.no_next,
            routing=self.routing,
            routing_type=self.routing_type,
            vdomparam=self.vdomparam)


def get_firewall_ipv6ehfilter(vdomparam: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallIpv6ehfilterResult:
    """
    Use this data source to get information on fortios firewall ipv6ehfilter


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallIpv6ehfilter:getFirewallIpv6ehfilter', __args__, opts=opts, typ=GetFirewallIpv6ehfilterResult).value

    return AwaitableGetFirewallIpv6ehfilterResult(
        auth=__ret__.auth,
        dest_opt=__ret__.dest_opt,
        fragment=__ret__.fragment,
        hdopt_type=__ret__.hdopt_type,
        hop_opt=__ret__.hop_opt,
        id=__ret__.id,
        no_next=__ret__.no_next,
        routing=__ret__.routing,
        routing_type=__ret__.routing_type,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_firewall_ipv6ehfilter)
def get_firewall_ipv6ehfilter_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallIpv6ehfilterResult]:
    """
    Use this data source to get information on fortios firewall ipv6ehfilter


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
