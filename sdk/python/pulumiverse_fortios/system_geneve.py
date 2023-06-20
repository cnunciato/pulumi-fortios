# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemGeneveArgs', 'SystemGeneve']

@pulumi.input_type
class SystemGeneveArgs:
    def __init__(__self__, *,
                 interface: pulumi.Input[str],
                 ip_version: pulumi.Input[str],
                 remote_ip: pulumi.Input[str],
                 vni: pulumi.Input[int],
                 dstport: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_ip6: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemGeneve resource.
        :param pulumi.Input[str] interface: Outgoing interface for GENEVE encapsulated traffic.
        :param pulumi.Input[str] ip_version: IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        :param pulumi.Input[str] remote_ip: IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[int] vni: GENEVE network ID.
        :param pulumi.Input[int] dstport: GENEVE destination port (1 - 65535, default = 6081).
        :param pulumi.Input[str] name: GENEVE device or interface name. Must be an unique interface name.
        :param pulumi.Input[str] remote_ip6: IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] type: GENEVE type. Valid values: `ethernet`, `ppp`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "interface", interface)
        pulumi.set(__self__, "ip_version", ip_version)
        pulumi.set(__self__, "remote_ip", remote_ip)
        pulumi.set(__self__, "vni", vni)
        if dstport is not None:
            pulumi.set(__self__, "dstport", dstport)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote_ip6 is not None:
            pulumi.set(__self__, "remote_ip6", remote_ip6)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Input[str]:
        """
        Outgoing interface for GENEVE encapsulated traffic.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Input[str]:
        """
        IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter(name="remoteIp")
    def remote_ip(self) -> pulumi.Input[str]:
        """
        IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        """
        return pulumi.get(self, "remote_ip")

    @remote_ip.setter
    def remote_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_ip", value)

    @property
    @pulumi.getter
    def vni(self) -> pulumi.Input[int]:
        """
        GENEVE network ID.
        """
        return pulumi.get(self, "vni")

    @vni.setter
    def vni(self, value: pulumi.Input[int]):
        pulumi.set(self, "vni", value)

    @property
    @pulumi.getter
    def dstport(self) -> Optional[pulumi.Input[int]]:
        """
        GENEVE destination port (1 - 65535, default = 6081).
        """
        return pulumi.get(self, "dstport")

    @dstport.setter
    def dstport(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dstport", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        GENEVE device or interface name. Must be an unique interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="remoteIp6")
    def remote_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        """
        return pulumi.get(self, "remote_ip6")

    @remote_ip6.setter
    def remote_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_ip6", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        GENEVE type. Valid values: `ethernet`, `ppp`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
class _SystemGeneveState:
    def __init__(__self__, *,
                 dstport: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_ip: Optional[pulumi.Input[str]] = None,
                 remote_ip6: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SystemGeneve resources.
        :param pulumi.Input[int] dstport: GENEVE destination port (1 - 65535, default = 6081).
        :param pulumi.Input[str] interface: Outgoing interface for GENEVE encapsulated traffic.
        :param pulumi.Input[str] ip_version: IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        :param pulumi.Input[str] name: GENEVE device or interface name. Must be an unique interface name.
        :param pulumi.Input[str] remote_ip: IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] remote_ip6: IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] type: GENEVE type. Valid values: `ethernet`, `ppp`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vni: GENEVE network ID.
        """
        if dstport is not None:
            pulumi.set(__self__, "dstport", dstport)
        if interface is not None:
            pulumi.set(__self__, "interface", interface)
        if ip_version is not None:
            pulumi.set(__self__, "ip_version", ip_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote_ip is not None:
            pulumi.set(__self__, "remote_ip", remote_ip)
        if remote_ip6 is not None:
            pulumi.set(__self__, "remote_ip6", remote_ip6)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if vni is not None:
            pulumi.set(__self__, "vni", vni)

    @property
    @pulumi.getter
    def dstport(self) -> Optional[pulumi.Input[int]]:
        """
        GENEVE destination port (1 - 65535, default = 6081).
        """
        return pulumi.get(self, "dstport")

    @dstport.setter
    def dstport(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dstport", value)

    @property
    @pulumi.getter
    def interface(self) -> Optional[pulumi.Input[str]]:
        """
        Outgoing interface for GENEVE encapsulated traffic.
        """
        return pulumi.get(self, "interface")

    @interface.setter
    def interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface", value)

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> Optional[pulumi.Input[str]]:
        """
        IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        """
        return pulumi.get(self, "ip_version")

    @ip_version.setter
    def ip_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        GENEVE device or interface name. Must be an unique interface name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="remoteIp")
    def remote_ip(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        """
        return pulumi.get(self, "remote_ip")

    @remote_ip.setter
    def remote_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_ip", value)

    @property
    @pulumi.getter(name="remoteIp6")
    def remote_ip6(self) -> Optional[pulumi.Input[str]]:
        """
        IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        """
        return pulumi.get(self, "remote_ip6")

    @remote_ip6.setter
    def remote_ip6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_ip6", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        GENEVE type. Valid values: `ethernet`, `ppp`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
    def vni(self) -> Optional[pulumi.Input[int]]:
        """
        GENEVE network ID.
        """
        return pulumi.get(self, "vni")

    @vni.setter
    def vni(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vni", value)


class SystemGeneve(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dstport: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_ip: Optional[pulumi.Input[str]] = None,
                 remote_ip6: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configure GENEVE devices. Applies to FortiOS Version `>= 6.2.4`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemGeneve("trname",
            dstport=22,
            interface="port2",
            ip_version="ipv4-unicast",
            remote_ip="1.1.1.1",
            remote_ip6="::",
            vni=0)
        ```

        ## Import

        System Geneve can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemGeneve:SystemGeneve labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemGeneve:SystemGeneve labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] dstport: GENEVE destination port (1 - 65535, default = 6081).
        :param pulumi.Input[str] interface: Outgoing interface for GENEVE encapsulated traffic.
        :param pulumi.Input[str] ip_version: IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        :param pulumi.Input[str] name: GENEVE device or interface name. Must be an unique interface name.
        :param pulumi.Input[str] remote_ip: IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] remote_ip6: IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] type: GENEVE type. Valid values: `ethernet`, `ppp`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vni: GENEVE network ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemGeneveArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure GENEVE devices. Applies to FortiOS Version `>= 6.2.4`.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.SystemGeneve("trname",
            dstport=22,
            interface="port2",
            ip_version="ipv4-unicast",
            remote_ip="1.1.1.1",
            remote_ip6="::",
            vni=0)
        ```

        ## Import

        System Geneve can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemGeneve:SystemGeneve labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemGeneve:SystemGeneve labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemGeneveArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemGeneveArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dstport: Optional[pulumi.Input[int]] = None,
                 interface: Optional[pulumi.Input[str]] = None,
                 ip_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_ip: Optional[pulumi.Input[str]] = None,
                 remote_ip6: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemGeneveArgs.__new__(SystemGeneveArgs)

            __props__.__dict__["dstport"] = dstport
            if interface is None and not opts.urn:
                raise TypeError("Missing required property 'interface'")
            __props__.__dict__["interface"] = interface
            if ip_version is None and not opts.urn:
                raise TypeError("Missing required property 'ip_version'")
            __props__.__dict__["ip_version"] = ip_version
            __props__.__dict__["name"] = name
            if remote_ip is None and not opts.urn:
                raise TypeError("Missing required property 'remote_ip'")
            __props__.__dict__["remote_ip"] = remote_ip
            __props__.__dict__["remote_ip6"] = remote_ip6
            __props__.__dict__["type"] = type
            __props__.__dict__["vdomparam"] = vdomparam
            if vni is None and not opts.urn:
                raise TypeError("Missing required property 'vni'")
            __props__.__dict__["vni"] = vni
        super(SystemGeneve, __self__).__init__(
            'fortios:index/systemGeneve:SystemGeneve',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dstport: Optional[pulumi.Input[int]] = None,
            interface: Optional[pulumi.Input[str]] = None,
            ip_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote_ip: Optional[pulumi.Input[str]] = None,
            remote_ip6: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            vni: Optional[pulumi.Input[int]] = None) -> 'SystemGeneve':
        """
        Get an existing SystemGeneve resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] dstport: GENEVE destination port (1 - 65535, default = 6081).
        :param pulumi.Input[str] interface: Outgoing interface for GENEVE encapsulated traffic.
        :param pulumi.Input[str] ip_version: IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        :param pulumi.Input[str] name: GENEVE device or interface name. Must be an unique interface name.
        :param pulumi.Input[str] remote_ip: IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] remote_ip6: IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        :param pulumi.Input[str] type: GENEVE type. Valid values: `ethernet`, `ppp`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[int] vni: GENEVE network ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemGeneveState.__new__(_SystemGeneveState)

        __props__.__dict__["dstport"] = dstport
        __props__.__dict__["interface"] = interface
        __props__.__dict__["ip_version"] = ip_version
        __props__.__dict__["name"] = name
        __props__.__dict__["remote_ip"] = remote_ip
        __props__.__dict__["remote_ip6"] = remote_ip6
        __props__.__dict__["type"] = type
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["vni"] = vni
        return SystemGeneve(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dstport(self) -> pulumi.Output[int]:
        """
        GENEVE destination port (1 - 65535, default = 6081).
        """
        return pulumi.get(self, "dstport")

    @property
    @pulumi.getter
    def interface(self) -> pulumi.Output[str]:
        """
        Outgoing interface for GENEVE encapsulated traffic.
        """
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> pulumi.Output[str]:
        """
        IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        GENEVE device or interface name. Must be an unique interface name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteIp")
    def remote_ip(self) -> pulumi.Output[str]:
        """
        IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
        """
        return pulumi.get(self, "remote_ip")

    @property
    @pulumi.getter(name="remoteIp6")
    def remote_ip6(self) -> pulumi.Output[str]:
        """
        IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
        """
        return pulumi.get(self, "remote_ip6")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        GENEVE type. Valid values: `ethernet`, `ppp`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def vni(self) -> pulumi.Output[int]:
        """
        GENEVE network ID.
        """
        return pulumi.get(self, "vni")

