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
from ._inputs import *

__all__ = ['SwitchcontrollerVlanArgs', 'SwitchcontrollerVlan']

@pulumi.input_type
class SwitchcontrollerVlanArgs:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_message_override_group: Optional[pulumi.Input[str]] = None,
                 portal_message_overrides: Optional[pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs']] = None,
                 radius_server: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 selected_usergroups: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]]] = None,
                 usergroup: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlanid: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SwitchcontrollerVlan resource.
        :param pulumi.Input[str] auth: Authentication. Valid values: `radius`, `usergroup`.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Switch VLAN name.
        :param pulumi.Input[str] portal_message_override_group: Specify captive portal replacement message override group.
        :param pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs'] portal_message_overrides: Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        :param pulumi.Input[str] radius_server: Authentication radius server.
        :param pulumi.Input[str] security: Security. Valid values: `open`, `captive-portal`, `8021x`.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]] selected_usergroups: Selected user group. The structure of `selected_usergroups` block is documented below.
        :param pulumi.Input[str] usergroup: Authentication usergroup.
        :param pulumi.Input[str] vdom: Virtual domain,
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlanid: VLAN ID.
        """
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portal_message_override_group is not None:
            pulumi.set(__self__, "portal_message_override_group", portal_message_override_group)
        if portal_message_overrides is not None:
            pulumi.set(__self__, "portal_message_overrides", portal_message_overrides)
        if radius_server is not None:
            pulumi.set(__self__, "radius_server", radius_server)
        if security is not None:
            pulumi.set(__self__, "security", security)
        if selected_usergroups is not None:
            pulumi.set(__self__, "selected_usergroups", selected_usergroups)
        if usergroup is not None:
            pulumi.set(__self__, "usergroup", usergroup)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlanid is not None:
            pulumi.set(__self__, "vlanid", vlanid)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication. Valid values: `radius`, `usergroup`.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

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
        Switch VLAN name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portalMessageOverrideGroup")
    def portal_message_override_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specify captive portal replacement message override group.
        """
        return pulumi.get(self, "portal_message_override_group")

    @portal_message_override_group.setter
    def portal_message_override_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal_message_override_group", value)

    @property
    @pulumi.getter(name="portalMessageOverrides")
    def portal_message_overrides(self) -> Optional[pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs']]:
        """
        Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        """
        return pulumi.get(self, "portal_message_overrides")

    @portal_message_overrides.setter
    def portal_message_overrides(self, value: Optional[pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs']]):
        pulumi.set(self, "portal_message_overrides", value)

    @property
    @pulumi.getter(name="radiusServer")
    def radius_server(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication radius server.
        """
        return pulumi.get(self, "radius_server")

    @radius_server.setter
    def radius_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "radius_server", value)

    @property
    @pulumi.getter
    def security(self) -> Optional[pulumi.Input[str]]:
        """
        Security. Valid values: `open`, `captive-portal`, `8021x`.
        """
        return pulumi.get(self, "security")

    @security.setter
    def security(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security", value)

    @property
    @pulumi.getter(name="selectedUsergroups")
    def selected_usergroups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]]]:
        """
        Selected user group. The structure of `selected_usergroups` block is documented below.
        """
        return pulumi.get(self, "selected_usergroups")

    @selected_usergroups.setter
    def selected_usergroups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]]]):
        pulumi.set(self, "selected_usergroups", value)

    @property
    @pulumi.getter
    def usergroup(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication usergroup.
        """
        return pulumi.get(self, "usergroup")

    @usergroup.setter
    def usergroup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usergroup", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual domain,
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)

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
    @pulumi.getter
    def vlanid(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN ID.
        """
        return pulumi.get(self, "vlanid")

    @vlanid.setter
    def vlanid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlanid", value)


@pulumi.input_type
class _SwitchcontrollerVlanState:
    def __init__(__self__, *,
                 auth: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_message_override_group: Optional[pulumi.Input[str]] = None,
                 portal_message_overrides: Optional[pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs']] = None,
                 radius_server: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 selected_usergroups: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]]] = None,
                 usergroup: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlanid: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SwitchcontrollerVlan resources.
        :param pulumi.Input[str] auth: Authentication. Valid values: `radius`, `usergroup`.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Switch VLAN name.
        :param pulumi.Input[str] portal_message_override_group: Specify captive portal replacement message override group.
        :param pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs'] portal_message_overrides: Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        :param pulumi.Input[str] radius_server: Authentication radius server.
        :param pulumi.Input[str] security: Security. Valid values: `open`, `captive-portal`, `8021x`.
        :param pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]] selected_usergroups: Selected user group. The structure of `selected_usergroups` block is documented below.
        :param pulumi.Input[str] usergroup: Authentication usergroup.
        :param pulumi.Input[str] vdom: Virtual domain,
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlanid: VLAN ID.
        """
        if auth is not None:
            pulumi.set(__self__, "auth", auth)
        if color is not None:
            pulumi.set(__self__, "color", color)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if portal_message_override_group is not None:
            pulumi.set(__self__, "portal_message_override_group", portal_message_override_group)
        if portal_message_overrides is not None:
            pulumi.set(__self__, "portal_message_overrides", portal_message_overrides)
        if radius_server is not None:
            pulumi.set(__self__, "radius_server", radius_server)
        if security is not None:
            pulumi.set(__self__, "security", security)
        if selected_usergroups is not None:
            pulumi.set(__self__, "selected_usergroups", selected_usergroups)
        if usergroup is not None:
            pulumi.set(__self__, "usergroup", usergroup)
        if vdom is not None:
            pulumi.set(__self__, "vdom", vdom)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vlanid is not None:
            pulumi.set(__self__, "vlanid", vlanid)

    @property
    @pulumi.getter
    def auth(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication. Valid values: `radius`, `usergroup`.
        """
        return pulumi.get(self, "auth")

    @auth.setter
    def auth(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth", value)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[int]]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

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
        Switch VLAN name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="portalMessageOverrideGroup")
    def portal_message_override_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specify captive portal replacement message override group.
        """
        return pulumi.get(self, "portal_message_override_group")

    @portal_message_override_group.setter
    def portal_message_override_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "portal_message_override_group", value)

    @property
    @pulumi.getter(name="portalMessageOverrides")
    def portal_message_overrides(self) -> Optional[pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs']]:
        """
        Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        """
        return pulumi.get(self, "portal_message_overrides")

    @portal_message_overrides.setter
    def portal_message_overrides(self, value: Optional[pulumi.Input['SwitchcontrollerVlanPortalMessageOverridesArgs']]):
        pulumi.set(self, "portal_message_overrides", value)

    @property
    @pulumi.getter(name="radiusServer")
    def radius_server(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication radius server.
        """
        return pulumi.get(self, "radius_server")

    @radius_server.setter
    def radius_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "radius_server", value)

    @property
    @pulumi.getter
    def security(self) -> Optional[pulumi.Input[str]]:
        """
        Security. Valid values: `open`, `captive-portal`, `8021x`.
        """
        return pulumi.get(self, "security")

    @security.setter
    def security(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security", value)

    @property
    @pulumi.getter(name="selectedUsergroups")
    def selected_usergroups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]]]:
        """
        Selected user group. The structure of `selected_usergroups` block is documented below.
        """
        return pulumi.get(self, "selected_usergroups")

    @selected_usergroups.setter
    def selected_usergroups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SwitchcontrollerVlanSelectedUsergroupArgs']]]]):
        pulumi.set(self, "selected_usergroups", value)

    @property
    @pulumi.getter
    def usergroup(self) -> Optional[pulumi.Input[str]]:
        """
        Authentication usergroup.
        """
        return pulumi.get(self, "usergroup")

    @usergroup.setter
    def usergroup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usergroup", value)

    @property
    @pulumi.getter
    def vdom(self) -> Optional[pulumi.Input[str]]:
        """
        Virtual domain,
        """
        return pulumi.get(self, "vdom")

    @vdom.setter
    def vdom(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdom", value)

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
    @pulumi.getter
    def vlanid(self) -> Optional[pulumi.Input[int]]:
        """
        VLAN ID.
        """
        return pulumi.get(self, "vlanid")

    @vlanid.setter
    def vlanid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vlanid", value)


class SwitchcontrollerVlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_message_override_group: Optional[pulumi.Input[str]] = None,
                 portal_message_overrides: Optional[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanPortalMessageOverridesArgs']]] = None,
                 radius_server: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 selected_usergroups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanSelectedUsergroupArgs']]]]] = None,
                 usergroup: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlanid: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configure VLANs for switch controller. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        SwitchController Vlan can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchcontrollerVlan:SwitchcontrollerVlan labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchcontrollerVlan:SwitchcontrollerVlan labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth: Authentication. Valid values: `radius`, `usergroup`.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Switch VLAN name.
        :param pulumi.Input[str] portal_message_override_group: Specify captive portal replacement message override group.
        :param pulumi.Input[pulumi.InputType['SwitchcontrollerVlanPortalMessageOverridesArgs']] portal_message_overrides: Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        :param pulumi.Input[str] radius_server: Authentication radius server.
        :param pulumi.Input[str] security: Security. Valid values: `open`, `captive-portal`, `8021x`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanSelectedUsergroupArgs']]]] selected_usergroups: Selected user group. The structure of `selected_usergroups` block is documented below.
        :param pulumi.Input[str] usergroup: Authentication usergroup.
        :param pulumi.Input[str] vdom: Virtual domain,
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlanid: VLAN ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchcontrollerVlanArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure VLANs for switch controller. Applies to FortiOS Version `<= 6.2.0`.

        ## Import

        SwitchController Vlan can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchcontrollerVlan:SwitchcontrollerVlan labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchcontrollerVlan:SwitchcontrollerVlan labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchcontrollerVlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchcontrollerVlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth: Optional[pulumi.Input[str]] = None,
                 color: Optional[pulumi.Input[int]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 portal_message_override_group: Optional[pulumi.Input[str]] = None,
                 portal_message_overrides: Optional[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanPortalMessageOverridesArgs']]] = None,
                 radius_server: Optional[pulumi.Input[str]] = None,
                 security: Optional[pulumi.Input[str]] = None,
                 selected_usergroups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanSelectedUsergroupArgs']]]]] = None,
                 usergroup: Optional[pulumi.Input[str]] = None,
                 vdom: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vlanid: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchcontrollerVlanArgs.__new__(SwitchcontrollerVlanArgs)

            __props__.__dict__["auth"] = auth
            __props__.__dict__["color"] = color
            __props__.__dict__["comments"] = comments
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["name"] = name
            __props__.__dict__["portal_message_override_group"] = portal_message_override_group
            __props__.__dict__["portal_message_overrides"] = portal_message_overrides
            __props__.__dict__["radius_server"] = radius_server
            __props__.__dict__["security"] = security
            __props__.__dict__["selected_usergroups"] = selected_usergroups
            __props__.__dict__["usergroup"] = usergroup
            __props__.__dict__["vdom"] = vdom
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vlanid"] = vlanid
        super(SwitchcontrollerVlan, __self__).__init__(
            'fortios:index/switchcontrollerVlan:SwitchcontrollerVlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth: Optional[pulumi.Input[str]] = None,
            color: Optional[pulumi.Input[int]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            portal_message_override_group: Optional[pulumi.Input[str]] = None,
            portal_message_overrides: Optional[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanPortalMessageOverridesArgs']]] = None,
            radius_server: Optional[pulumi.Input[str]] = None,
            security: Optional[pulumi.Input[str]] = None,
            selected_usergroups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanSelectedUsergroupArgs']]]]] = None,
            usergroup: Optional[pulumi.Input[str]] = None,
            vdom: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vlanid: Optional[pulumi.Input[int]] = None) -> 'SwitchcontrollerVlan':
        """
        Get an existing SwitchcontrollerVlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth: Authentication. Valid values: `radius`, `usergroup`.
        :param pulumi.Input[int] color: Color of icon on the GUI.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: Switch VLAN name.
        :param pulumi.Input[str] portal_message_override_group: Specify captive portal replacement message override group.
        :param pulumi.Input[pulumi.InputType['SwitchcontrollerVlanPortalMessageOverridesArgs']] portal_message_overrides: Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        :param pulumi.Input[str] radius_server: Authentication radius server.
        :param pulumi.Input[str] security: Security. Valid values: `open`, `captive-portal`, `8021x`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SwitchcontrollerVlanSelectedUsergroupArgs']]]] selected_usergroups: Selected user group. The structure of `selected_usergroups` block is documented below.
        :param pulumi.Input[str] usergroup: Authentication usergroup.
        :param pulumi.Input[str] vdom: Virtual domain,
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vlanid: VLAN ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchcontrollerVlanState.__new__(_SwitchcontrollerVlanState)

        __props__.__dict__["auth"] = auth
        __props__.__dict__["color"] = color
        __props__.__dict__["comments"] = comments
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["name"] = name
        __props__.__dict__["portal_message_override_group"] = portal_message_override_group
        __props__.__dict__["portal_message_overrides"] = portal_message_overrides
        __props__.__dict__["radius_server"] = radius_server
        __props__.__dict__["security"] = security
        __props__.__dict__["selected_usergroups"] = selected_usergroups
        __props__.__dict__["usergroup"] = usergroup
        __props__.__dict__["vdom"] = vdom
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vlanid"] = vlanid
        return SwitchcontrollerVlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def auth(self) -> pulumi.Output[str]:
        """
        Authentication. Valid values: `radius`, `usergroup`.
        """
        return pulumi.get(self, "auth")

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[int]:
        """
        Color of icon on the GUI.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Switch VLAN name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalMessageOverrideGroup")
    def portal_message_override_group(self) -> pulumi.Output[str]:
        """
        Specify captive portal replacement message override group.
        """
        return pulumi.get(self, "portal_message_override_group")

    @property
    @pulumi.getter(name="portalMessageOverrides")
    def portal_message_overrides(self) -> pulumi.Output['outputs.SwitchcontrollerVlanPortalMessageOverrides']:
        """
        Individual message overrides. The structure of `portal_message_overrides` block is documented below.
        """
        return pulumi.get(self, "portal_message_overrides")

    @property
    @pulumi.getter(name="radiusServer")
    def radius_server(self) -> pulumi.Output[str]:
        """
        Authentication radius server.
        """
        return pulumi.get(self, "radius_server")

    @property
    @pulumi.getter
    def security(self) -> pulumi.Output[str]:
        """
        Security. Valid values: `open`, `captive-portal`, `8021x`.
        """
        return pulumi.get(self, "security")

    @property
    @pulumi.getter(name="selectedUsergroups")
    def selected_usergroups(self) -> pulumi.Output[Optional[Sequence['outputs.SwitchcontrollerVlanSelectedUsergroup']]]:
        """
        Selected user group. The structure of `selected_usergroups` block is documented below.
        """
        return pulumi.get(self, "selected_usergroups")

    @property
    @pulumi.getter
    def usergroup(self) -> pulumi.Output[str]:
        """
        Authentication usergroup.
        """
        return pulumi.get(self, "usergroup")

    @property
    @pulumi.getter
    def vdom(self) -> pulumi.Output[str]:
        """
        Virtual domain,
        """
        return pulumi.get(self, "vdom")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vlanid(self) -> pulumi.Output[int]:
        """
        VLAN ID.
        """
        return pulumi.get(self, "vlanid")

