# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NsxtSettingArgs', 'NsxtSetting']

@pulumi.input_type
class NsxtSettingArgs:
    def __init__(__self__, *,
                 liveness: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NsxtSetting resource.
        :param pulumi.Input[str] liveness: Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] service: Service name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if liveness is not None:
            pulumi.set(__self__, "liveness", liveness)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def liveness(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "liveness")

    @liveness.setter
    def liveness(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "liveness", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        Service name.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)

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
class _NsxtSettingState:
    def __init__(__self__, *,
                 liveness: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NsxtSetting resources.
        :param pulumi.Input[str] liveness: Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] service: Service name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if liveness is not None:
            pulumi.set(__self__, "liveness", liveness)
        if service is not None:
            pulumi.set(__self__, "service", service)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def liveness(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "liveness")

    @liveness.setter
    def liveness(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "liveness", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        Service name.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)

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


class NsxtSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 liveness: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure NSX-T setting. Applies to FortiOS Version `>= 6.4.10`.

        ## Import

        Nsxt Setting can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/nsxtSetting:NsxtSetting labelname NsxtSetting
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/nsxtSetting:NsxtSetting labelname NsxtSetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] liveness: Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] service: Service name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NsxtSettingArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure NSX-T setting. Applies to FortiOS Version `>= 6.4.10`.

        ## Import

        Nsxt Setting can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/nsxtSetting:NsxtSetting labelname NsxtSetting
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/nsxtSetting:NsxtSetting labelname NsxtSetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param NsxtSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxtSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 liveness: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NsxtSettingArgs.__new__(NsxtSettingArgs)

            __props__.__dict__["liveness"] = liveness
            __props__.__dict__["service"] = service
            __props__.__dict__["vdomparam"] = vdomparam
        super(NsxtSetting, __self__).__init__(
            'fortios:index/nsxtSetting:NsxtSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            liveness: Optional[pulumi.Input[str]] = None,
            service: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'NsxtSetting':
        """
        Get an existing NsxtSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] liveness: Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] service: Service name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxtSettingState.__new__(_NsxtSettingState)

        __props__.__dict__["liveness"] = liveness
        __props__.__dict__["service"] = service
        __props__.__dict__["vdomparam"] = vdomparam
        return NsxtSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def liveness(self) -> pulumi.Output[str]:
        """
        Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "liveness")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[str]:
        """
        Service name.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

