# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VirtualwirepairArgs', 'Virtualwirepair']

@pulumi.input_type
class VirtualwirepairArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]],
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Virtualwirepair resource.
        :param pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]] members: Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Virtual-wire-pair name. Must be a unique interface name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan_filter: Set VLAN filters.
        :param pulumi.Input[str] wildcard_vlan: Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        pulumi.set(__self__, "members", members)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan_filter is not None:
            pulumi.set(__self__, "vlan_filter", vlan_filter)
        if wildcard_vlan is not None:
            pulumi.set(__self__, "wildcard_vlan", wildcard_vlan)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]]:
        """
        Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual-wire-pair name. Must be a unique interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="vlanFilter")
    def vlan_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Set VLAN filters.
        """
        return pulumi.get(self, "vlan_filter")

    @vlan_filter.setter
    def vlan_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_filter", value)

    @property
    @pulumi.getter(name="wildcardVlan")
    def wildcard_vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "wildcard_vlan")

    @wildcard_vlan.setter
    def wildcard_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wildcard_vlan", value)


@pulumi.input_type
class _VirtualwirepairState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Virtualwirepair resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]] members: Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Virtual-wire-pair name. Must be a unique interface name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan_filter: Set VLAN filters.
        :param pulumi.Input[str] wildcard_vlan: Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlan_filter is not None:
            pulumi.set(__self__, "vlan_filter", vlan_filter)
        if wildcard_vlan is not None:
            pulumi.set(__self__, "wildcard_vlan", wildcard_vlan)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]]]:
        """
        Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VirtualwirepairMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual-wire-pair name. Must be a unique interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @vdomparam.setter
    def vdomparam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdomparam", value)

    @property
    @pulumi.getter(name="vlanFilter")
    def vlan_filter(self) -> Optional[pulumi.Input[str]]:
        """
        Set VLAN filters.
        """
        return pulumi.get(self, "vlan_filter")

    @vlan_filter.setter
    def vlan_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vlan_filter", value)

    @property
    @pulumi.getter(name="wildcardVlan")
    def wildcard_vlan(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "wildcard_vlan")

    @wildcard_vlan.setter
    def wildcard_vlan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wildcard_vlan", value)


class Virtualwirepair(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualwirepairMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure virtual wire pairs.

        ## Import

        System VirtualWirePair can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualwirepairMemberArgs']]]] members: Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Virtual-wire-pair name. Must be a unique interface name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan_filter: Set VLAN filters.
        :param pulumi.Input[str] wildcard_vlan: Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VirtualwirepairArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure virtual wire pairs.

        ## Import

        System VirtualWirePair can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:system/virtualwirepair:Virtualwirepair labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VirtualwirepairArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VirtualwirepairArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualwirepairMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlan_filter: Optional[pulumi.Input[str]] = None,
                 wildcard_vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VirtualwirepairArgs.__new__(VirtualwirepairArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vlan_filter"] = vlan_filter
            __props__.__dict__["wildcard_vlan"] = wildcard_vlan
        super(Virtualwirepair, __self__).__init__(
            'fortios:system/virtualwirepair:Virtualwirepair',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualwirepairMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vlan_filter: Optional[pulumi.Input[str]] = None,
            wildcard_vlan: Optional[pulumi.Input[str]] = None) -> 'Virtualwirepair':
        """
        Get an existing Virtualwirepair resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VirtualwirepairMemberArgs']]]] members: Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: Virtual-wire-pair name. Must be a unique interface name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vlan_filter: Set VLAN filters.
        :param pulumi.Input[str] wildcard_vlan: Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VirtualwirepairState.__new__(_VirtualwirepairState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vlan_filter"] = vlan_filter
        __props__.__dict__["wildcard_vlan"] = wildcard_vlan
        return Virtualwirepair(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.VirtualwirepairMember']]:
        """
        Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Virtual-wire-pair name. Must be a unique interface name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="vlanFilter")
    def vlan_filter(self) -> pulumi.Output[str]:
        """
        Set VLAN filters.
        """
        return pulumi.get(self, "vlan_filter")

    @property
    @pulumi.getter(name="wildcardVlan")
    def wildcard_vlan(self) -> pulumi.Output[str]:
        """
        Enable/disable wildcard VLAN. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "wildcard_vlan")
