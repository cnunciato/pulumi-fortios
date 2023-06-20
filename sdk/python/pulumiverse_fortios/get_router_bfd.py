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
    'GetRouterBfdResult',
    'AwaitableGetRouterBfdResult',
    'get_router_bfd',
    'get_router_bfd_output',
]

@pulumi.output_type
class GetRouterBfdResult:
    """
    A collection of values returned by getRouterBfd.
    """
    def __init__(__self__, id=None, multihop_templates=None, neighbors=None, vdomparam=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if multihop_templates and not isinstance(multihop_templates, list):
            raise TypeError("Expected argument 'multihop_templates' to be a list")
        pulumi.set(__self__, "multihop_templates", multihop_templates)
        if neighbors and not isinstance(neighbors, list):
            raise TypeError("Expected argument 'neighbors' to be a list")
        pulumi.set(__self__, "neighbors", neighbors)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="multihopTemplates")
    def multihop_templates(self) -> Sequence['outputs.GetRouterBfdMultihopTemplateResult']:
        """
        BFD multi-hop template table. The structure of `multihop_template` block is documented below.
        """
        return pulumi.get(self, "multihop_templates")

    @property
    @pulumi.getter
    def neighbors(self) -> Sequence['outputs.GetRouterBfdNeighborResult']:
        """
        neighbor The structure of `neighbor` block is documented below.
        """
        return pulumi.get(self, "neighbors")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterBfdResult(GetRouterBfdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterBfdResult(
            id=self.id,
            multihop_templates=self.multihop_templates,
            neighbors=self.neighbors,
            vdomparam=self.vdomparam)


def get_router_bfd(vdomparam: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterBfdResult:
    """
    Use this data source to get information on fortios router bfd


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterBfd:getRouterBfd', __args__, opts=opts, typ=GetRouterBfdResult).value

    return AwaitableGetRouterBfdResult(
        id=__ret__.id,
        multihop_templates=__ret__.multihop_templates,
        neighbors=__ret__.neighbors,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_bfd)
def get_router_bfd_output(vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterBfdResult]:
    """
    Use this data source to get information on fortios router bfd


    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
