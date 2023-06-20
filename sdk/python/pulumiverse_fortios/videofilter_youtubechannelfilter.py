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

__all__ = ['VideofilterYoutubechannelfilterArgs', 'VideofilterYoutubechannelfilter']

@pulumi.input_type
class VideofilterYoutubechannelfilterArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_category: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VideofilterYoutubechannelfilter resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] default_action: YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]] entries: YouTube filter entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] log: Eanble/disable logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] override_category: Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if override_category is not None:
            pulumi.set(__self__, "override_category", override_category)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[str]]:
        """
        YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action", value)

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
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]]]:
        """
        YouTube filter entries. The structure of `entries` block is documented below.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Eanble/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="overrideCategory")
    def override_category(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override_category")

    @override_category.setter
    def override_category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_category", value)

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


@pulumi.input_type
class _VideofilterYoutubechannelfilterState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_category: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VideofilterYoutubechannelfilter resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] default_action: YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]] entries: YouTube filter entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] log: Eanble/disable logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] override_category: Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if entries is not None:
            pulumi.set(__self__, "entries", entries)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if log is not None:
            pulumi.set(__self__, "log", log)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if override_category is not None:
            pulumi.set(__self__, "override_category", override_category)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[str]]:
        """
        YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action", value)

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
    def entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]]]:
        """
        YouTube filter entries. The structure of `entries` block is documented below.
        """
        return pulumi.get(self, "entries")

    @entries.setter
    def entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VideofilterYoutubechannelfilterEntryArgs']]]]):
        pulumi.set(self, "entries", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def log(self) -> Optional[pulumi.Input[str]]:
        """
        Eanble/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @log.setter
    def log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="overrideCategory")
    def override_category(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override_category")

    @override_category.setter
    def override_category(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_category", value)

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


class VideofilterYoutubechannelfilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VideofilterYoutubechannelfilterEntryArgs']]]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_category: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure YouTube channel filter. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        Videofilter YoutubeChannelFilter can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/videofilterYoutubechannelfilter:VideofilterYoutubechannelfilter labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/videofilterYoutubechannelfilter:VideofilterYoutubechannelfilter labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] default_action: YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VideofilterYoutubechannelfilterEntryArgs']]]] entries: YouTube filter entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] log: Eanble/disable logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] override_category: Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VideofilterYoutubechannelfilterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure YouTube channel filter. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        Videofilter YoutubeChannelFilter can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/videofilterYoutubechannelfilter:VideofilterYoutubechannelfilter labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/videofilterYoutubechannelfilter:VideofilterYoutubechannelfilter labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VideofilterYoutubechannelfilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VideofilterYoutubechannelfilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VideofilterYoutubechannelfilterEntryArgs']]]]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 log: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_category: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VideofilterYoutubechannelfilterArgs.__new__(VideofilterYoutubechannelfilterArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["entries"] = entries
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["log"] = log
            __props__.__dict__["name"] = name
            __props__.__dict__["override_category"] = override_category
            __props__.__dict__["vdomparam"] = vdomparam
        super(VideofilterYoutubechannelfilter, __self__).__init__(
            'fortios:index/videofilterYoutubechannelfilter:VideofilterYoutubechannelfilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            default_action: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VideofilterYoutubechannelfilterEntryArgs']]]]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            log: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            override_category: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'VideofilterYoutubechannelfilter':
        """
        Get an existing VideofilterYoutubechannelfilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] default_action: YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VideofilterYoutubechannelfilterEntryArgs']]]] entries: YouTube filter entries. The structure of `entries` block is documented below.
        :param pulumi.Input[int] fosid: ID.
        :param pulumi.Input[str] log: Eanble/disable logging. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] override_category: Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VideofilterYoutubechannelfilterState.__new__(_VideofilterYoutubechannelfilterState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["default_action"] = default_action
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["entries"] = entries
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["log"] = log
        __props__.__dict__["name"] = name
        __props__.__dict__["override_category"] = override_category
        __props__.__dict__["vdomparam"] = vdomparam
        return VideofilterYoutubechannelfilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output[str]:
        """
        YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def entries(self) -> pulumi.Output[Optional[Sequence['outputs.VideofilterYoutubechannelfilterEntry']]]:
        """
        YouTube filter entries. The structure of `entries` block is documented below.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def log(self) -> pulumi.Output[str]:
        """
        Eanble/disable logging. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "log")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overrideCategory")
    def override_category(self) -> pulumi.Output[str]:
        """
        Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "override_category")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

