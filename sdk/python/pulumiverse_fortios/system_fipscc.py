# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemFipsccArgs', 'SystemFipscc']

@pulumi.input_type
class SystemFipsccArgs:
    def __init__(__self__, *,
                 entropy_token: Optional[pulumi.Input[str]] = None,
                 key_generation_self_test: Optional[pulumi.Input[str]] = None,
                 self_test_period: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemFipscc resource.
        :param pulumi.Input[str] entropy_token: Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        :param pulumi.Input[str] key_generation_self_test: Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] self_test_period: Self test period.
        :param pulumi.Input[str] status: Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if entropy_token is not None:
            pulumi.set(__self__, "entropy_token", entropy_token)
        if key_generation_self_test is not None:
            pulumi.set(__self__, "key_generation_self_test", key_generation_self_test)
        if self_test_period is not None:
            pulumi.set(__self__, "self_test_period", self_test_period)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="entropyToken")
    def entropy_token(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        """
        return pulumi.get(self, "entropy_token")

    @entropy_token.setter
    def entropy_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entropy_token", value)

    @property
    @pulumi.getter(name="keyGenerationSelfTest")
    def key_generation_self_test(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "key_generation_self_test")

    @key_generation_self_test.setter
    def key_generation_self_test(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_generation_self_test", value)

    @property
    @pulumi.getter(name="selfTestPeriod")
    def self_test_period(self) -> Optional[pulumi.Input[int]]:
        """
        Self test period.
        """
        return pulumi.get(self, "self_test_period")

    @self_test_period.setter
    def self_test_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "self_test_period", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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
class _SystemFipsccState:
    def __init__(__self__, *,
                 entropy_token: Optional[pulumi.Input[str]] = None,
                 key_generation_self_test: Optional[pulumi.Input[str]] = None,
                 self_test_period: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemFipscc resources.
        :param pulumi.Input[str] entropy_token: Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        :param pulumi.Input[str] key_generation_self_test: Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] self_test_period: Self test period.
        :param pulumi.Input[str] status: Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if entropy_token is not None:
            pulumi.set(__self__, "entropy_token", entropy_token)
        if key_generation_self_test is not None:
            pulumi.set(__self__, "key_generation_self_test", key_generation_self_test)
        if self_test_period is not None:
            pulumi.set(__self__, "self_test_period", self_test_period)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="entropyToken")
    def entropy_token(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        """
        return pulumi.get(self, "entropy_token")

    @entropy_token.setter
    def entropy_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entropy_token", value)

    @property
    @pulumi.getter(name="keyGenerationSelfTest")
    def key_generation_self_test(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "key_generation_self_test")

    @key_generation_self_test.setter
    def key_generation_self_test(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_generation_self_test", value)

    @property
    @pulumi.getter(name="selfTestPeriod")
    def self_test_period(self) -> Optional[pulumi.Input[int]]:
        """
        Self test period.
        """
        return pulumi.get(self, "self_test_period")

    @self_test_period.setter
    def self_test_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "self_test_period", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

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


class SystemFipscc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entropy_token: Optional[pulumi.Input[str]] = None,
                 key_generation_self_test: Optional[pulumi.Input[str]] = None,
                 self_test_period: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure FIPS-CC mode.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemFipscc("trname",
            entropy_token="enable",
            key_generation_self_test="disable",
            self_test_period=1440,
            status="disable")
        ```

        ## Import

        System FipsCc can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemFipscc:SystemFipscc labelname SystemFipsCc
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemFipscc:SystemFipscc labelname SystemFipsCc
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entropy_token: Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        :param pulumi.Input[str] key_generation_self_test: Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] self_test_period: Self test period.
        :param pulumi.Input[str] status: Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SystemFipsccArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure FIPS-CC mode.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemFipscc("trname",
            entropy_token="enable",
            key_generation_self_test="disable",
            self_test_period=1440,
            status="disable")
        ```

        ## Import

        System FipsCc can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemFipscc:SystemFipscc labelname SystemFipsCc
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemFipscc:SystemFipscc labelname SystemFipsCc
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemFipsccArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemFipsccArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entropy_token: Optional[pulumi.Input[str]] = None,
                 key_generation_self_test: Optional[pulumi.Input[str]] = None,
                 self_test_period: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemFipsccArgs.__new__(SystemFipsccArgs)

            __props__.__dict__["entropy_token"] = entropy_token
            __props__.__dict__["key_generation_self_test"] = key_generation_self_test
            __props__.__dict__["self_test_period"] = self_test_period
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemFipscc, __self__).__init__(
            'fortios:index/systemFipscc:SystemFipscc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            entropy_token: Optional[pulumi.Input[str]] = None,
            key_generation_self_test: Optional[pulumi.Input[str]] = None,
            self_test_period: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemFipscc':
        """
        Get an existing SystemFipscc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entropy_token: Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        :param pulumi.Input[str] key_generation_self_test: Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] self_test_period: Self test period.
        :param pulumi.Input[str] status: Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemFipsccState.__new__(_SystemFipsccState)

        __props__.__dict__["entropy_token"] = entropy_token
        __props__.__dict__["key_generation_self_test"] = key_generation_self_test
        __props__.__dict__["self_test_period"] = self_test_period
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemFipscc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="entropyToken")
    def entropy_token(self) -> pulumi.Output[str]:
        """
        Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
        """
        return pulumi.get(self, "entropy_token")

    @property
    @pulumi.getter(name="keyGenerationSelfTest")
    def key_generation_self_test(self) -> pulumi.Output[str]:
        """
        Enable/disable self tests after key generation. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "key_generation_self_test")

    @property
    @pulumi.getter(name="selfTestPeriod")
    def self_test_period(self) -> pulumi.Output[int]:
        """
        Self test period.
        """
        return pulumi.get(self, "self_test_period")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

