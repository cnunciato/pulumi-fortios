# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WirelesscontrollerAddressArgs', 'WirelesscontrollerAddress']

@pulumi.input_type
class WirelesscontrollerAddressArgs:
    def __init__(__self__, *,
                 fosid: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WirelesscontrollerAddress resource.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] policy: Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

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
class _WirelesscontrollerAddressState:
    def __init__(__self__, *,
                 fosid: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WirelesscontrollerAddress resources.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] policy: Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[str]]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

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


class WirelesscontrollerAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure the client with its MAC address. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6`.

        ## Import

        WirelessController Address can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] policy: Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WirelesscontrollerAddressArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure the client with its MAC address. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6`.

        ## Import

        WirelessController Address can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WirelesscontrollerAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WirelesscontrollerAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fosid: Optional[pulumi.Input[str]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WirelesscontrollerAddressArgs.__new__(WirelesscontrollerAddressArgs)

            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["mac"] = mac
            __props__.__dict__["policy"] = policy
            __props__.__dict__["vdomparam"] = vdomparam
        super(WirelesscontrollerAddress, __self__).__init__(
            'fortios:index/wirelesscontrollerAddress:WirelesscontrollerAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fosid: Optional[pulumi.Input[str]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WirelesscontrollerAddress':
        """
        Get an existing WirelesscontrollerAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fosid: ID.
        :param pulumi.Input[str] mac: MAC address.
        :param pulumi.Input[str] policy: Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WirelesscontrollerAddressState.__new__(_WirelesscontrollerAddressState)

        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["mac"] = mac
        __props__.__dict__["policy"] = policy
        __props__.__dict__["vdomparam"] = vdomparam
        return WirelesscontrollerAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[str]:
        """
        ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        MAC address.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        Allow or block the client with this MAC address. Valid values: `allow`, `deny`.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

