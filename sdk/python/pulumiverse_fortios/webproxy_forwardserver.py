# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['WebproxyForwardserverArgs', 'WebproxyForwardserver']

@pulumi.input_type
class WebproxyForwardserverArgs:
    def __init__(__self__, *,
                 addr_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_down_option: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebproxyForwardserver resource.
        :param pulumi.Input[str] addr_type: Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] fqdn: Forward server Fully Qualified Domain Name (FQDN).
        :param pulumi.Input[str] healthcheck: Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] ip: Forward proxy server IP address.
        :param pulumi.Input[str] monitor: URL for forward server health check monitoring (default = http://www.google.com).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] password: HTTP authentication password.
        :param pulumi.Input[int] port: Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        :param pulumi.Input[str] server_down_option: Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        :param pulumi.Input[str] username: HTTP authentication user name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if addr_type is not None:
            pulumi.set(__self__, "addr_type", addr_type)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if healthcheck is not None:
            pulumi.set(__self__, "healthcheck", healthcheck)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if monitor is not None:
            pulumi.set(__self__, "monitor", monitor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_down_option is not None:
            pulumi.set(__self__, "server_down_option", server_down_option)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> Optional[pulumi.Input[str]]:
        """
        Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        """
        return pulumi.get(self, "addr_type")

    @addr_type.setter
    def addr_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_type", value)

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
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        Forward server Fully Qualified Domain Name (FQDN).
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter
    def healthcheck(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "healthcheck")

    @healthcheck.setter
    def healthcheck(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "healthcheck", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Forward proxy server IP address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input[str]]:
        """
        URL for forward server health check monitoring (default = http://www.google.com).
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP authentication password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverDownOption")
    def server_down_option(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        """
        return pulumi.get(self, "server_down_option")

    @server_down_option.setter
    def server_down_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_down_option", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP authentication user name.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

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
class _WebproxyForwardserverState:
    def __init__(__self__, *,
                 addr_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_down_option: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebproxyForwardserver resources.
        :param pulumi.Input[str] addr_type: Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] fqdn: Forward server Fully Qualified Domain Name (FQDN).
        :param pulumi.Input[str] healthcheck: Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] ip: Forward proxy server IP address.
        :param pulumi.Input[str] monitor: URL for forward server health check monitoring (default = http://www.google.com).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] password: HTTP authentication password.
        :param pulumi.Input[int] port: Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        :param pulumi.Input[str] server_down_option: Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        :param pulumi.Input[str] username: HTTP authentication user name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if addr_type is not None:
            pulumi.set(__self__, "addr_type", addr_type)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if healthcheck is not None:
            pulumi.set(__self__, "healthcheck", healthcheck)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if monitor is not None:
            pulumi.set(__self__, "monitor", monitor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_down_option is not None:
            pulumi.set(__self__, "server_down_option", server_down_option)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> Optional[pulumi.Input[str]]:
        """
        Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        """
        return pulumi.get(self, "addr_type")

    @addr_type.setter
    def addr_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "addr_type", value)

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
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        Forward server Fully Qualified Domain Name (FQDN).
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter
    def healthcheck(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "healthcheck")

    @healthcheck.setter
    def healthcheck(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "healthcheck", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Forward proxy server IP address.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input[str]]:
        """
        URL for forward server health check monitoring (default = http://www.google.com).
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP authentication password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverDownOption")
    def server_down_option(self) -> Optional[pulumi.Input[str]]:
        """
        Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        """
        return pulumi.get(self, "server_down_option")

    @server_down_option.setter
    def server_down_option(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_down_option", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP authentication user name.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

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


class WebproxyForwardserver(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_down_option: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure forward-server addresses.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebproxyForwardserver("trname",
            addr_type="fqdn",
            healthcheck="disable",
            ip="0.0.0.0",
            monitor="http://www.google.com",
            port=3128,
            server_down_option="block")
        ```

        ## Import

        WebProxy ForwardServer can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webproxyForwardserver:WebproxyForwardserver labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webproxyForwardserver:WebproxyForwardserver labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] addr_type: Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] fqdn: Forward server Fully Qualified Domain Name (FQDN).
        :param pulumi.Input[str] healthcheck: Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] ip: Forward proxy server IP address.
        :param pulumi.Input[str] monitor: URL for forward server health check monitoring (default = http://www.google.com).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] password: HTTP authentication password.
        :param pulumi.Input[int] port: Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        :param pulumi.Input[str] server_down_option: Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        :param pulumi.Input[str] username: HTTP authentication user name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WebproxyForwardserverArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure forward-server addresses.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.WebproxyForwardserver("trname",
            addr_type="fqdn",
            healthcheck="disable",
            ip="0.0.0.0",
            monitor="http://www.google.com",
            port=3128,
            server_down_option="block")
        ```

        ## Import

        WebProxy ForwardServer can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/webproxyForwardserver:WebproxyForwardserver labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/webproxyForwardserver:WebproxyForwardserver labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param WebproxyForwardserverArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebproxyForwardserverArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addr_type: Optional[pulumi.Input[str]] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 healthcheck: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_down_option: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebproxyForwardserverArgs.__new__(WebproxyForwardserverArgs)

            __props__.__dict__["addr_type"] = addr_type
            __props__.__dict__["comment"] = comment
            __props__.__dict__["fqdn"] = fqdn
            __props__.__dict__["healthcheck"] = healthcheck
            __props__.__dict__["ip"] = ip
            __props__.__dict__["monitor"] = monitor
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["port"] = port
            __props__.__dict__["server_down_option"] = server_down_option
            __props__.__dict__["username"] = username
            __props__.__dict__["vdomparam"] = vdomparam
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(WebproxyForwardserver, __self__).__init__(
            'fortios:index/webproxyForwardserver:WebproxyForwardserver',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addr_type: Optional[pulumi.Input[str]] = None,
            comment: Optional[pulumi.Input[str]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            healthcheck: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            monitor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            server_down_option: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'WebproxyForwardserver':
        """
        Get an existing WebproxyForwardserver resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] addr_type: Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        :param pulumi.Input[str] comment: Comment.
        :param pulumi.Input[str] fqdn: Forward server Fully Qualified Domain Name (FQDN).
        :param pulumi.Input[str] healthcheck: Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        :param pulumi.Input[str] ip: Forward proxy server IP address.
        :param pulumi.Input[str] monitor: URL for forward server health check monitoring (default = http://www.google.com).
        :param pulumi.Input[str] name: Server name.
        :param pulumi.Input[str] password: HTTP authentication password.
        :param pulumi.Input[int] port: Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        :param pulumi.Input[str] server_down_option: Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        :param pulumi.Input[str] username: HTTP authentication user name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebproxyForwardserverState.__new__(_WebproxyForwardserverState)

        __props__.__dict__["addr_type"] = addr_type
        __props__.__dict__["comment"] = comment
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["healthcheck"] = healthcheck
        __props__.__dict__["ip"] = ip
        __props__.__dict__["monitor"] = monitor
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["port"] = port
        __props__.__dict__["server_down_option"] = server_down_option
        __props__.__dict__["username"] = username
        __props__.__dict__["vdomparam"] = vdomparam
        return WebproxyForwardserver(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addrType")
    def addr_type(self) -> pulumi.Output[str]:
        """
        Address type of the forwarding proxy server: IP or FQDN. Valid values: `ip`, `fqdn`.
        """
        return pulumi.get(self, "addr_type")

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[str]:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        Forward server Fully Qualified Domain Name (FQDN).
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def healthcheck(self) -> pulumi.Output[str]:
        """
        Enable/disable forward server health checking. Attempts to connect through the remote forwarding server to a destination to verify that the forwarding server is operating normally. Valid values: `disable`, `enable`.
        """
        return pulumi.get(self, "healthcheck")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        Forward proxy server IP address.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def monitor(self) -> pulumi.Output[str]:
        """
        URL for forward server health check monitoring (default = http://www.google.com).
        """
        return pulumi.get(self, "monitor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Server name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        HTTP authentication password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Port number that the forwarding server expects to receive HTTP sessions on (1 - 65535, default = 3128).
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serverDownOption")
    def server_down_option(self) -> pulumi.Output[str]:
        """
        Action to take when the forward server is found to be down: block sessions until the server is back up or pass sessions to their destination. Valid values: `block`, `pass`.
        """
        return pulumi.get(self, "server_down_option")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        HTTP authentication user name.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

