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
    'GetFirewallPolicylistResult',
    'AwaitableGetFirewallPolicylistResult',
    'get_firewall_policylist',
    'get_firewall_policylist_output',
]

@pulumi.output_type
class GetFirewallPolicylistResult:
    """
    A collection of values returned by getFirewallPolicylist.
    """
    def __init__(__self__, filter=None, id=None, policyidlists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policyidlists and not isinstance(policyidlists, list):
            raise TypeError("Expected argument 'policyidlists' to be a list")
        pulumi.set(__self__, "policyidlists", policyidlists)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def policyidlists(self) -> Sequence[int]:
        """
        A list of the `FirewallPolicy`.
        """
        return pulumi.get(self, "policyidlists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallPolicylistResult(GetFirewallPolicylistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicylistResult(
            filter=self.filter,
            id=self.id,
            policyidlists=self.policyidlists,
            vdomparam=self.vdomparam)


def get_firewall_policylist(filter: Optional[str] = None,
                            vdomparam: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallPolicylistResult:
    """
    Provides a list of `FirewallPolicy`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fortios as fortios

    sample1 = fortios.get_firewall_policylist()
    pulumi.export("output1", sample1)
    sample2 = fortios.get_firewall_policylist(filter="schedule==always&action==accept,action==deny")
    pulumi.export("sample2Output", sample2.policyidlists)
    ```


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallPolicylist:getFirewallPolicylist', __args__, opts=opts, typ=GetFirewallPolicylistResult).value

    return AwaitableGetFirewallPolicylistResult(
        filter=__ret__.filter,
        id=__ret__.id,
        policyidlists=__ret__.policyidlists,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_firewall_policylist)
def get_firewall_policylist_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                   vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallPolicylistResult]:
    """
    Provides a list of `FirewallPolicy`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fortios as fortios

    sample1 = fortios.get_firewall_policylist()
    pulumi.export("output1", sample1)
    sample2 = fortios.get_firewall_policylist(filter="schedule==always&action==accept,action==deny")
    pulumi.export("sample2Output", sample2.policyidlists)
    ```


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
