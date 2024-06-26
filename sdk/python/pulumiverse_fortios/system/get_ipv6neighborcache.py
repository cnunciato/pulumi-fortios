# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetIpv6neighborcacheResult',
    'AwaitableGetIpv6neighborcacheResult',
    'get_ipv6neighborcache',
    'get_ipv6neighborcache_output',
]

@pulumi.output_type
class GetIpv6neighborcacheResult:
    """
    A collection of values returned by getIpv6neighborcache.
    """
    def __init__(__self__, fosid=None, id=None, interface=None, ipv6=None, mac=None, vdomparam=None):
        if fosid and not isinstance(fosid, int):
            raise TypeError("Expected argument 'fosid' to be a int")
        pulumi.set(__self__, "fosid", fosid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if ipv6 and not isinstance(ipv6, str):
            raise TypeError("Expected argument 'ipv6' to be a str")
        pulumi.set(__self__, "ipv6", ipv6)
        if mac and not isinstance(mac, str):
            raise TypeError("Expected argument 'mac' to be a str")
        pulumi.set(__self__, "mac", mac)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> int:
        """
        Unique integer ID of the entry.
        """
        return pulumi.get(self, "fosid")

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
        Select the associated interface name from available options.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        """
        IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def mac(self) -> str:
        """
        MAC address (format: xx:xx:xx:xx:xx:xx).
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetIpv6neighborcacheResult(GetIpv6neighborcacheResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpv6neighborcacheResult(
            fosid=self.fosid,
            id=self.id,
            interface=self.interface,
            ipv6=self.ipv6,
            mac=self.mac,
            vdomparam=self.vdomparam)


def get_ipv6neighborcache(fosid: Optional[int] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpv6neighborcacheResult:
    """
    Use this data source to get information on an fortios system ipv6neighborcache


    :param int fosid: Specify the fosid of the desired system ipv6neighborcache.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['fosid'] = fosid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:system/getIpv6neighborcache:getIpv6neighborcache', __args__, opts=opts, typ=GetIpv6neighborcacheResult).value

    return AwaitableGetIpv6neighborcacheResult(
        fosid=pulumi.get(__ret__, 'fosid'),
        id=pulumi.get(__ret__, 'id'),
        interface=pulumi.get(__ret__, 'interface'),
        ipv6=pulumi.get(__ret__, 'ipv6'),
        mac=pulumi.get(__ret__, 'mac'),
        vdomparam=pulumi.get(__ret__, 'vdomparam'))


@_utilities.lift_output_func(get_ipv6neighborcache)
def get_ipv6neighborcache_output(fosid: Optional[pulumi.Input[int]] = None,
                                 vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpv6neighborcacheResult]:
    """
    Use this data source to get information on an fortios system ipv6neighborcache


    :param int fosid: Specify the fosid of the desired system ipv6neighborcache.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
