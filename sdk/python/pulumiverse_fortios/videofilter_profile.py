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

__all__ = ['VideofilterProfileArgs', 'VideofilterProfile']

@pulumi.input_type
class VideofilterProfileArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dailymotion: Optional[pulumi.Input[str]] = None,
                 fortiguard_category: Optional[pulumi.Input['VideofilterProfileFortiguardCategoryArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vimeo: Optional[pulumi.Input[str]] = None,
                 youtube: Optional[pulumi.Input[str]] = None,
                 youtube_channel_filter: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a VideofilterProfile resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dailymotion: Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        :param pulumi.Input['VideofilterProfileFortiguardCategoryArgs'] fortiguard_category: Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] replacemsg_group: Replacement message group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vimeo: Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] youtube: Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] youtube_channel_filter: Set YouTube channel filter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dailymotion is not None:
            pulumi.set(__self__, "dailymotion", dailymotion)
        if fortiguard_category is not None:
            pulumi.set(__self__, "fortiguard_category", fortiguard_category)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if replacemsg_group is not None:
            pulumi.set(__self__, "replacemsg_group", replacemsg_group)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vimeo is not None:
            pulumi.set(__self__, "vimeo", vimeo)
        if youtube is not None:
            pulumi.set(__self__, "youtube", youtube)
        if youtube_channel_filter is not None:
            pulumi.set(__self__, "youtube_channel_filter", youtube_channel_filter)

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
    @pulumi.getter
    def dailymotion(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dailymotion")

    @dailymotion.setter
    def dailymotion(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dailymotion", value)

    @property
    @pulumi.getter(name="fortiguardCategory")
    def fortiguard_category(self) -> Optional[pulumi.Input['VideofilterProfileFortiguardCategoryArgs']]:
        """
        Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        """
        return pulumi.get(self, "fortiguard_category")

    @fortiguard_category.setter
    def fortiguard_category(self, value: Optional[pulumi.Input['VideofilterProfileFortiguardCategoryArgs']]):
        pulumi.set(self, "fortiguard_category", value)

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
    @pulumi.getter(name="replacemsgGroup")
    def replacemsg_group(self) -> Optional[pulumi.Input[str]]:
        """
        Replacement message group.
        """
        return pulumi.get(self, "replacemsg_group")

    @replacemsg_group.setter
    def replacemsg_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replacemsg_group", value)

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
    def vimeo(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "vimeo")

    @vimeo.setter
    def vimeo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vimeo", value)

    @property
    @pulumi.getter
    def youtube(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "youtube")

    @youtube.setter
    def youtube(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "youtube", value)

    @property
    @pulumi.getter(name="youtubeChannelFilter")
    def youtube_channel_filter(self) -> Optional[pulumi.Input[int]]:
        """
        Set YouTube channel filter.
        """
        return pulumi.get(self, "youtube_channel_filter")

    @youtube_channel_filter.setter
    def youtube_channel_filter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "youtube_channel_filter", value)


@pulumi.input_type
class _VideofilterProfileState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 dailymotion: Optional[pulumi.Input[str]] = None,
                 fortiguard_category: Optional[pulumi.Input['VideofilterProfileFortiguardCategoryArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vimeo: Optional[pulumi.Input[str]] = None,
                 youtube: Optional[pulumi.Input[str]] = None,
                 youtube_channel_filter: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering VideofilterProfile resources.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dailymotion: Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        :param pulumi.Input['VideofilterProfileFortiguardCategoryArgs'] fortiguard_category: Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] replacemsg_group: Replacement message group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vimeo: Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] youtube: Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] youtube_channel_filter: Set YouTube channel filter.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if dailymotion is not None:
            pulumi.set(__self__, "dailymotion", dailymotion)
        if fortiguard_category is not None:
            pulumi.set(__self__, "fortiguard_category", fortiguard_category)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if replacemsg_group is not None:
            pulumi.set(__self__, "replacemsg_group", replacemsg_group)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vimeo is not None:
            pulumi.set(__self__, "vimeo", vimeo)
        if youtube is not None:
            pulumi.set(__self__, "youtube", youtube)
        if youtube_channel_filter is not None:
            pulumi.set(__self__, "youtube_channel_filter", youtube_channel_filter)

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
    @pulumi.getter
    def dailymotion(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dailymotion")

    @dailymotion.setter
    def dailymotion(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dailymotion", value)

    @property
    @pulumi.getter(name="fortiguardCategory")
    def fortiguard_category(self) -> Optional[pulumi.Input['VideofilterProfileFortiguardCategoryArgs']]:
        """
        Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        """
        return pulumi.get(self, "fortiguard_category")

    @fortiguard_category.setter
    def fortiguard_category(self, value: Optional[pulumi.Input['VideofilterProfileFortiguardCategoryArgs']]):
        pulumi.set(self, "fortiguard_category", value)

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
    @pulumi.getter(name="replacemsgGroup")
    def replacemsg_group(self) -> Optional[pulumi.Input[str]]:
        """
        Replacement message group.
        """
        return pulumi.get(self, "replacemsg_group")

    @replacemsg_group.setter
    def replacemsg_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "replacemsg_group", value)

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
    def vimeo(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "vimeo")

    @vimeo.setter
    def vimeo(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vimeo", value)

    @property
    @pulumi.getter
    def youtube(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "youtube")

    @youtube.setter
    def youtube(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "youtube", value)

    @property
    @pulumi.getter(name="youtubeChannelFilter")
    def youtube_channel_filter(self) -> Optional[pulumi.Input[int]]:
        """
        Set YouTube channel filter.
        """
        return pulumi.get(self, "youtube_channel_filter")

    @youtube_channel_filter.setter
    def youtube_channel_filter(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "youtube_channel_filter", value)


class VideofilterProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dailymotion: Optional[pulumi.Input[str]] = None,
                 fortiguard_category: Optional[pulumi.Input[pulumi.InputType['VideofilterProfileFortiguardCategoryArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vimeo: Optional[pulumi.Input[str]] = None,
                 youtube: Optional[pulumi.Input[str]] = None,
                 youtube_channel_filter: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configure VideoFilter profile. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        Videofilter Profile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/videofilterProfile:VideofilterProfile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/videofilterProfile:VideofilterProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dailymotion: Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[pulumi.InputType['VideofilterProfileFortiguardCategoryArgs']] fortiguard_category: Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] replacemsg_group: Replacement message group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vimeo: Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] youtube: Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] youtube_channel_filter: Set YouTube channel filter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VideofilterProfileArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure VideoFilter profile. Applies to FortiOS Version `>= 7.0.1`.

        ## Import

        Videofilter Profile can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/videofilterProfile:VideofilterProfile labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/videofilterProfile:VideofilterProfile labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param VideofilterProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VideofilterProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 dailymotion: Optional[pulumi.Input[str]] = None,
                 fortiguard_category: Optional[pulumi.Input[pulumi.InputType['VideofilterProfileFortiguardCategoryArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 replacemsg_group: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vimeo: Optional[pulumi.Input[str]] = None,
                 youtube: Optional[pulumi.Input[str]] = None,
                 youtube_channel_filter: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VideofilterProfileArgs.__new__(VideofilterProfileArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["dailymotion"] = dailymotion
            __props__.__dict__["fortiguard_category"] = fortiguard_category
            __props__.__dict__["name"] = name
            __props__.__dict__["replacemsg_group"] = replacemsg_group
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["vimeo"] = vimeo
            __props__.__dict__["youtube"] = youtube
            __props__.__dict__["youtube_channel_filter"] = youtube_channel_filter
        super(VideofilterProfile, __self__).__init__(
            'fortios:index/videofilterProfile:VideofilterProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            dailymotion: Optional[pulumi.Input[str]] = None,
            fortiguard_category: Optional[pulumi.Input[pulumi.InputType['VideofilterProfileFortiguardCategoryArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            replacemsg_group: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vimeo: Optional[pulumi.Input[str]] = None,
            youtube: Optional[pulumi.Input[str]] = None,
            youtube_channel_filter: Optional[pulumi.Input[int]] = None) -> 'VideofilterProfile':
        """
        Get an existing VideofilterProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] dailymotion: Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[pulumi.InputType['VideofilterProfileFortiguardCategoryArgs']] fortiguard_category: Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        :param pulumi.Input[str] name: Name.
        :param pulumi.Input[str] replacemsg_group: Replacement message group.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] vimeo: Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] youtube: Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] youtube_channel_filter: Set YouTube channel filter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VideofilterProfileState.__new__(_VideofilterProfileState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["dailymotion"] = dailymotion
        __props__.__dict__["fortiguard_category"] = fortiguard_category
        __props__.__dict__["name"] = name
        __props__.__dict__["replacemsg_group"] = replacemsg_group
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vimeo"] = vimeo
        __props__.__dict__["youtube"] = youtube
        __props__.__dict__["youtube_channel_filter"] = youtube_channel_filter
        return VideofilterProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def dailymotion(self) -> pulumi.Output[str]:
        """
        Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "dailymotion")

    @property
    @pulumi.getter(name="fortiguardCategory")
    def fortiguard_category(self) -> pulumi.Output['outputs.VideofilterProfileFortiguardCategory']:
        """
        Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
        """
        return pulumi.get(self, "fortiguard_category")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="replacemsgGroup")
    def replacemsg_group(self) -> pulumi.Output[str]:
        """
        Replacement message group.
        """
        return pulumi.get(self, "replacemsg_group")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vimeo(self) -> pulumi.Output[str]:
        """
        Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "vimeo")

    @property
    @pulumi.getter
    def youtube(self) -> pulumi.Output[str]:
        """
        Enable/disable YouTube video source. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "youtube")

    @property
    @pulumi.getter(name="youtubeChannelFilter")
    def youtube_channel_filter(self) -> pulumi.Output[int]:
        """
        Set YouTube channel filter.
        """
        return pulumi.get(self, "youtube_channel_filter")

