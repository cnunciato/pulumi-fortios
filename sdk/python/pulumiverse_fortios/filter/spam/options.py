# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['OptionsArgs', 'Options']

@pulumi.input_type
class OptionsArgs:
    def __init__(__self__, *,
                 dns_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Options resource.
        :param pulumi.Input[int] dns_timeout: DNS query time out (1 - 30 sec).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dns_timeout is not None:
            pulumi.set(__self__, "dns_timeout", dns_timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsTimeout")
    def dns_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        DNS query time out (1 - 30 sec).
        """
        return pulumi.get(self, "dns_timeout")

    @dns_timeout.setter
    def dns_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dns_timeout", value)

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
class _OptionsState:
    def __init__(__self__, *,
                 dns_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Options resources.
        :param pulumi.Input[int] dns_timeout: DNS query time out (1 - 30 sec).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dns_timeout is not None:
            pulumi.set(__self__, "dns_timeout", dns_timeout)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="dnsTimeout")
    def dns_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        DNS query time out (1 - 30 sec).
        """
        return pulumi.get(self, "dns_timeout")

    @dns_timeout.setter
    def dns_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dns_timeout", value)

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


class Options(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure AntiSpam options. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.filter.spam.Options("trname", dns_timeout=7)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Spamfilter Options can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:filter/spam/options:Options labelname SpamfilterOptions
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:filter/spam/options:Options labelname SpamfilterOptions
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] dns_timeout: DNS query time out (1 - 30 sec).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OptionsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure AntiSpam options. Applies to FortiOS Version `<= 6.2.0`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.filter.spam.Options("trname", dns_timeout=7)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Spamfilter Options can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:filter/spam/options:Options labelname SpamfilterOptions
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:filter/spam/options:Options labelname SpamfilterOptions
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param OptionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OptionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_timeout: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OptionsArgs.__new__(OptionsArgs)

            __props__.__dict__["dns_timeout"] = dns_timeout
            __props__.__dict__["vdomparam"] = vdomparam
        super(Options, __self__).__init__(
            'fortios:filter/spam/options:Options',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_timeout: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'Options':
        """
        Get an existing Options resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] dns_timeout: DNS query time out (1 - 30 sec).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OptionsState.__new__(_OptionsState)

        __props__.__dict__["dns_timeout"] = dns_timeout
        __props__.__dict__["vdomparam"] = vdomparam
        return Options(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsTimeout")
    def dns_timeout(self) -> pulumi.Output[int]:
        """
        DNS query time out (1 - 30 sec).
        """
        return pulumi.get(self, "dns_timeout")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

