# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetFirewallDoSpolicy6Result',
    'AwaitableGetFirewallDoSpolicy6Result',
    'get_firewall_do_spolicy6',
    'get_firewall_do_spolicy6_output',
]

@pulumi.output_type
class GetFirewallDoSpolicy6Result:
    """
    A collection of values returned by getFirewallDoSpolicy6.
    """
    def __init__(__self__, anomalies=None, comments=None, dstaddrs=None, id=None, interface=None, name=None, policyid=None, services=None, srcaddrs=None, status=None, vdomparam=None):
        if anomalies and not isinstance(anomalies, list):
            raise TypeError("Expected argument 'anomalies' to be a list")
        pulumi.set(__self__, "anomalies", anomalies)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if dstaddrs and not isinstance(dstaddrs, list):
            raise TypeError("Expected argument 'dstaddrs' to be a list")
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policyid and not isinstance(policyid, int):
            raise TypeError("Expected argument 'policyid' to be a int")
        pulumi.set(__self__, "policyid", policyid)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if srcaddrs and not isinstance(srcaddrs, list):
            raise TypeError("Expected argument 'srcaddrs' to be a list")
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def anomalies(self) -> Sequence['outputs.GetFirewallDoSpolicy6AnomalyResult']:
        """
        Anomaly name. The structure of `anomaly` block is documented below.
        """
        return pulumi.get(self, "anomalies")

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dstaddrs(self) -> Sequence['outputs.GetFirewallDoSpolicy6DstaddrResult']:
        """
        Destination address name from available addresses. The structure of `dstaddr` block is documented below.
        """
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interface(self) -> str:
        """
        Incoming interface name from available interfaces.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Anomaly name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policyid(self) -> int:
        """
        Policy ID.
        """
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetFirewallDoSpolicy6ServiceResult']:
        """
        Service object from available options. The structure of `service` block is documented below.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def srcaddrs(self) -> Sequence['outputs.GetFirewallDoSpolicy6SrcaddrResult']:
        """
        Source address name from available addresses. The structure of `srcaddr` block is documented below.
        """
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this anomaly.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallDoSpolicy6Result(GetFirewallDoSpolicy6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallDoSpolicy6Result(
            anomalies=self.anomalies,
            comments=self.comments,
            dstaddrs=self.dstaddrs,
            id=self.id,
            interface=self.interface,
            name=self.name,
            policyid=self.policyid,
            services=self.services,
            srcaddrs=self.srcaddrs,
            status=self.status,
            vdomparam=self.vdomparam)


def get_firewall_do_spolicy6(policyid: Optional[int] = None,
                             vdomparam: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallDoSpolicy6Result:
    """
    Use this data source to get information on an fortios firewall DoSpolicy6


    :param int policyid: Specify the policyid of the desired firewall DoSpolicy6.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['policyid'] = policyid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallDoSpolicy6:getFirewallDoSpolicy6', __args__, opts=opts, typ=GetFirewallDoSpolicy6Result).value

    return AwaitableGetFirewallDoSpolicy6Result(
        anomalies=__ret__.anomalies,
        comments=__ret__.comments,
        dstaddrs=__ret__.dstaddrs,
        id=__ret__.id,
        interface=__ret__.interface,
        name=__ret__.name,
        policyid=__ret__.policyid,
        services=__ret__.services,
        srcaddrs=__ret__.srcaddrs,
        status=__ret__.status,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_firewall_do_spolicy6)
def get_firewall_do_spolicy6_output(policyid: Optional[pulumi.Input[int]] = None,
                                    vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallDoSpolicy6Result]:
    """
    Use this data source to get information on an fortios firewall DoSpolicy6


    :param int policyid: Specify the policyid of the desired firewall DoSpolicy6.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
