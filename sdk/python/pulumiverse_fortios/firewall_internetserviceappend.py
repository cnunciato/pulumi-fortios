# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallInternetserviceappendArgs', 'FirewallInternetserviceappend']

@pulumi.input_type
class FirewallInternetserviceappendArgs:
    def __init__(__self__, *,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallInternetserviceappend resource.
        :param pulumi.Input[str] addr_mode: Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        :param pulumi.Input[int] append_port: Appending TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[int] match_port: Matching TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if addr_mode is not None:
            pulumi.set(__self__, "addr_mode", addr_mode)
        if append_port is not None:
            pulumi.set(__self__, "append_port", append_port)
        if match_port is not None:
            pulumi.set(__self__, "match_port", match_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        """
        return pulumi.get(self, "addr_mode")

    @addr_mode.setter
    def addr_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_mode", value)

    @property
    @pulumi.getter(name="appendPort")
    def append_port(self) -> Optional[pulumi.Input[int]]:
        """
        Appending TCP/UDP/SCTP destination port (1 to 65535).
        """
        return pulumi.get(self, "append_port")

    @append_port.setter
    def append_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "append_port", value)

    @property
    @pulumi.getter(name="matchPort")
    def match_port(self) -> Optional[pulumi.Input[int]]:
        """
        Matching TCP/UDP/SCTP destination port (1 to 65535).
        """
        return pulumi.get(self, "match_port")

    @match_port.setter
    def match_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_port", value)

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
class _FirewallInternetserviceappendState:
    def __init__(__self__, *,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallInternetserviceappend resources.
        :param pulumi.Input[str] addr_mode: Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        :param pulumi.Input[int] append_port: Appending TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[int] match_port: Matching TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if addr_mode is not None:
            pulumi.set(__self__, "addr_mode", addr_mode)
        if append_port is not None:
            pulumi.set(__self__, "append_port", append_port)
        if match_port is not None:
            pulumi.set(__self__, "match_port", match_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        """
        return pulumi.get(self, "addr_mode")

    @addr_mode.setter
    def addr_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_mode", value)

    @property
    @pulumi.getter(name="appendPort")
    def append_port(self) -> Optional[pulumi.Input[int]]:
        """
        Appending TCP/UDP/SCTP destination port (1 to 65535).
        """
        return pulumi.get(self, "append_port")

    @append_port.setter
    def append_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "append_port", value)

    @property
    @pulumi.getter(name="matchPort")
    def match_port(self) -> Optional[pulumi.Input[int]]:
        """
        Matching TCP/UDP/SCTP destination port (1 to 65535).
        """
        return pulumi.get(self, "match_port")

    @match_port.setter
    def match_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "match_port", value)

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


class FirewallInternetserviceappend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure additional port mappings for Internet Services. Applies to FortiOS Version `6.2.4,6.2.6,6.4.1,6.4.2,6.4.10,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.2.0,7.2.1,7.2.2`.

        ## Import

        Firewall InternetServiceAppend can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallInternetserviceappend:FirewallInternetserviceappend labelname FirewallInternetServiceAppend
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallInternetserviceappend:FirewallInternetserviceappend labelname FirewallInternetServiceAppend
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] addr_mode: Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        :param pulumi.Input[int] append_port: Appending TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[int] match_port: Matching TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FirewallInternetserviceappendArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure additional port mappings for Internet Services. Applies to FortiOS Version `6.2.4,6.2.6,6.4.1,6.4.2,6.4.10,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.2.0,7.2.1,7.2.2`.

        ## Import

        Firewall InternetServiceAppend can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallInternetserviceappend:FirewallInternetserviceappend labelname FirewallInternetServiceAppend
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallInternetserviceappend:FirewallInternetserviceappend labelname FirewallInternetServiceAppend
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallInternetserviceappendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallInternetserviceappendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr_mode: Optional[pulumi.Input[str]] = None,
                 append_port: Optional[pulumi.Input[int]] = None,
                 match_port: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallInternetserviceappendArgs.__new__(FirewallInternetserviceappendArgs)

            __props__.__dict__["addr_mode"] = addr_mode
            __props__.__dict__["append_port"] = append_port
            __props__.__dict__["match_port"] = match_port
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallInternetserviceappend, __self__).__init__(
            'fortios:index/firewallInternetserviceappend:FirewallInternetserviceappend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addr_mode: Optional[pulumi.Input[str]] = None,
            append_port: Optional[pulumi.Input[int]] = None,
            match_port: Optional[pulumi.Input[int]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallInternetserviceappend':
        """
        Get an existing FirewallInternetserviceappend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] addr_mode: Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        :param pulumi.Input[int] append_port: Appending TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[int] match_port: Matching TCP/UDP/SCTP destination port (1 to 65535).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallInternetserviceappendState.__new__(_FirewallInternetserviceappendState)

        __props__.__dict__["addr_mode"] = addr_mode
        __props__.__dict__["append_port"] = append_port
        __props__.__dict__["match_port"] = match_port
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallInternetserviceappend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addrMode")
    def addr_mode(self) -> pulumi.Output[str]:
        """
        Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`, `both`.
        """
        return pulumi.get(self, "addr_mode")

    @property
    @pulumi.getter(name="appendPort")
    def append_port(self) -> pulumi.Output[int]:
        """
        Appending TCP/UDP/SCTP destination port (1 to 65535).
        """
        return pulumi.get(self, "append_port")

    @property
    @pulumi.getter(name="matchPort")
    def match_port(self) -> pulumi.Output[int]:
        """
        Matching TCP/UDP/SCTP destination port (1 to 65535).
        """
        return pulumi.get(self, "match_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

