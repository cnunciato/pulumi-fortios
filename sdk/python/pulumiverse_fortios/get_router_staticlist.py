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
    'GetRouterStaticlistResult',
    'AwaitableGetRouterStaticlistResult',
    'get_router_staticlist',
    'get_router_staticlist_output',
]

@pulumi.output_type
class GetRouterStaticlistResult:
    """
    A collection of values returned by getRouterStaticlist.
    """
    def __init__(__self__, filter=None, id=None, seq_numlists=None, vdomparam=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if seq_numlists and not isinstance(seq_numlists, list):
            raise TypeError("Expected argument 'seq_numlists' to be a list")
        pulumi.set(__self__, "seq_numlists", seq_numlists)
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
    @pulumi.getter(name="seqNumlists")
    def seq_numlists(self) -> Sequence[int]:
        """
        A list of the `RouterStatic`.
        """
        return pulumi.get(self, "seq_numlists")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetRouterStaticlistResult(GetRouterStaticlistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterStaticlistResult(
            filter=self.filter,
            id=self.id,
            seq_numlists=self.seq_numlists,
            vdomparam=self.vdomparam)


def get_router_staticlist(filter: Optional[str] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouterStaticlistResult:
    """
    Provides a list of `RouterStatic`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fortios as fortios

    sample1 = fortios.get_router_staticlist(filter="seq_num>1")
    pulumi.export("output1", sample1.seq_numlists)
    ```


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getRouterStaticlist:getRouterStaticlist', __args__, opts=opts, typ=GetRouterStaticlistResult).value

    return AwaitableGetRouterStaticlistResult(
        filter=__ret__.filter,
        id=__ret__.id,
        seq_numlists=__ret__.seq_numlists,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_router_staticlist)
def get_router_staticlist_output(filter: Optional[pulumi.Input[Optional[str]]] = None,
                                 vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouterStaticlistResult]:
    """
    Provides a list of `RouterStatic`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fortios as fortios

    sample1 = fortios.get_router_staticlist(filter="seq_num>1")
    pulumi.export("output1", sample1.seq_numlists)
    ```


    :param str filter: A filter used to scope the list. See Filter results of datasource.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
