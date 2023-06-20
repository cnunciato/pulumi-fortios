# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['Wirelesscontrollerhotspot20H2qpconncapabilityArgs', 'Wirelesscontrollerhotspot20H2qpconncapability']

@pulumi.input_type
class Wirelesscontrollerhotspot20H2qpconncapabilityArgs:
    def __init__(__self__, *,
                 esp_port: Optional[pulumi.Input[str]] = None,
                 ftp_port: Optional[pulumi.Input[str]] = None,
                 http_port: Optional[pulumi.Input[str]] = None,
                 icmp_port: Optional[pulumi.Input[str]] = None,
                 ikev2_port: Optional[pulumi.Input[str]] = None,
                 ikev2_xx_port: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pptp_vpn_port: Optional[pulumi.Input[str]] = None,
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 tls_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip_tcp_port: Optional[pulumi.Input[str]] = None,
                 voip_udp_port: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Wirelesscontrollerhotspot20H2qpconncapability resource.
        :param pulumi.Input[str] esp_port: Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ftp_port: Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] http_port: Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] icmp_port: Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_port: Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_xx_port: Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] name: Connection capability name.
        :param pulumi.Input[str] pptp_vpn_port: Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ssh_port: Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] tls_port: Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip_tcp_port: Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] voip_udp_port: Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        if esp_port is not None:
            pulumi.set(__self__, "esp_port", esp_port)
        if ftp_port is not None:
            pulumi.set(__self__, "ftp_port", ftp_port)
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if icmp_port is not None:
            pulumi.set(__self__, "icmp_port", icmp_port)
        if ikev2_port is not None:
            pulumi.set(__self__, "ikev2_port", ikev2_port)
        if ikev2_xx_port is not None:
            pulumi.set(__self__, "ikev2_xx_port", ikev2_xx_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pptp_vpn_port is not None:
            pulumi.set(__self__, "pptp_vpn_port", pptp_vpn_port)
        if ssh_port is not None:
            pulumi.set(__self__, "ssh_port", ssh_port)
        if tls_port is not None:
            pulumi.set(__self__, "tls_port", tls_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if voip_tcp_port is not None:
            pulumi.set(__self__, "voip_tcp_port", voip_tcp_port)
        if voip_udp_port is not None:
            pulumi.set(__self__, "voip_udp_port", voip_udp_port)

    @property
    @pulumi.getter(name="espPort")
    def esp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "esp_port")

    @esp_port.setter
    def esp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "esp_port", value)

    @property
    @pulumi.getter(name="ftpPort")
    def ftp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ftp_port")

    @ftp_port.setter
    def ftp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ftp_port", value)

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "http_port")

    @http_port.setter
    def http_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_port", value)

    @property
    @pulumi.getter(name="icmpPort")
    def icmp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "icmp_port")

    @icmp_port.setter
    def icmp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icmp_port", value)

    @property
    @pulumi.getter(name="ikev2Port")
    def ikev2_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ikev2_port")

    @ikev2_port.setter
    def ikev2_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ikev2_port", value)

    @property
    @pulumi.getter(name="ikev2XxPort")
    def ikev2_xx_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ikev2_xx_port")

    @ikev2_xx_port.setter
    def ikev2_xx_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ikev2_xx_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Connection capability name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pptpVpnPort")
    def pptp_vpn_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "pptp_vpn_port")

    @pptp_vpn_port.setter
    def pptp_vpn_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pptp_vpn_port", value)

    @property
    @pulumi.getter(name="sshPort")
    def ssh_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ssh_port")

    @ssh_port.setter
    def ssh_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_port", value)

    @property
    @pulumi.getter(name="tlsPort")
    def tls_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "tls_port")

    @tls_port.setter
    def tls_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_port", value)

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
    @pulumi.getter(name="voipTcpPort")
    def voip_tcp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "voip_tcp_port")

    @voip_tcp_port.setter
    def voip_tcp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip_tcp_port", value)

    @property
    @pulumi.getter(name="voipUdpPort")
    def voip_udp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "voip_udp_port")

    @voip_udp_port.setter
    def voip_udp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip_udp_port", value)


@pulumi.input_type
class _Wirelesscontrollerhotspot20H2qpconncapabilityState:
    def __init__(__self__, *,
                 esp_port: Optional[pulumi.Input[str]] = None,
                 ftp_port: Optional[pulumi.Input[str]] = None,
                 http_port: Optional[pulumi.Input[str]] = None,
                 icmp_port: Optional[pulumi.Input[str]] = None,
                 ikev2_port: Optional[pulumi.Input[str]] = None,
                 ikev2_xx_port: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pptp_vpn_port: Optional[pulumi.Input[str]] = None,
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 tls_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip_tcp_port: Optional[pulumi.Input[str]] = None,
                 voip_udp_port: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Wirelesscontrollerhotspot20H2qpconncapability resources.
        :param pulumi.Input[str] esp_port: Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ftp_port: Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] http_port: Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] icmp_port: Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_port: Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_xx_port: Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] name: Connection capability name.
        :param pulumi.Input[str] pptp_vpn_port: Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ssh_port: Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] tls_port: Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip_tcp_port: Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] voip_udp_port: Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        if esp_port is not None:
            pulumi.set(__self__, "esp_port", esp_port)
        if ftp_port is not None:
            pulumi.set(__self__, "ftp_port", ftp_port)
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if icmp_port is not None:
            pulumi.set(__self__, "icmp_port", icmp_port)
        if ikev2_port is not None:
            pulumi.set(__self__, "ikev2_port", ikev2_port)
        if ikev2_xx_port is not None:
            pulumi.set(__self__, "ikev2_xx_port", ikev2_xx_port)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pptp_vpn_port is not None:
            pulumi.set(__self__, "pptp_vpn_port", pptp_vpn_port)
        if ssh_port is not None:
            pulumi.set(__self__, "ssh_port", ssh_port)
        if tls_port is not None:
            pulumi.set(__self__, "tls_port", tls_port)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if voip_tcp_port is not None:
            pulumi.set(__self__, "voip_tcp_port", voip_tcp_port)
        if voip_udp_port is not None:
            pulumi.set(__self__, "voip_udp_port", voip_udp_port)

    @property
    @pulumi.getter(name="espPort")
    def esp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "esp_port")

    @esp_port.setter
    def esp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "esp_port", value)

    @property
    @pulumi.getter(name="ftpPort")
    def ftp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ftp_port")

    @ftp_port.setter
    def ftp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ftp_port", value)

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "http_port")

    @http_port.setter
    def http_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "http_port", value)

    @property
    @pulumi.getter(name="icmpPort")
    def icmp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "icmp_port")

    @icmp_port.setter
    def icmp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icmp_port", value)

    @property
    @pulumi.getter(name="ikev2Port")
    def ikev2_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ikev2_port")

    @ikev2_port.setter
    def ikev2_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ikev2_port", value)

    @property
    @pulumi.getter(name="ikev2XxPort")
    def ikev2_xx_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ikev2_xx_port")

    @ikev2_xx_port.setter
    def ikev2_xx_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ikev2_xx_port", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Connection capability name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pptpVpnPort")
    def pptp_vpn_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "pptp_vpn_port")

    @pptp_vpn_port.setter
    def pptp_vpn_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pptp_vpn_port", value)

    @property
    @pulumi.getter(name="sshPort")
    def ssh_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ssh_port")

    @ssh_port.setter
    def ssh_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_port", value)

    @property
    @pulumi.getter(name="tlsPort")
    def tls_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "tls_port")

    @tls_port.setter
    def tls_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_port", value)

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
    @pulumi.getter(name="voipTcpPort")
    def voip_tcp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "voip_tcp_port")

    @voip_tcp_port.setter
    def voip_tcp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip_tcp_port", value)

    @property
    @pulumi.getter(name="voipUdpPort")
    def voip_udp_port(self) -> Optional[pulumi.Input[str]]:
        """
        Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "voip_udp_port")

    @voip_udp_port.setter
    def voip_udp_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "voip_udp_port", value)


class Wirelesscontrollerhotspot20H2qpconncapability(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 esp_port: Optional[pulumi.Input[str]] = None,
                 ftp_port: Optional[pulumi.Input[str]] = None,
                 http_port: Optional[pulumi.Input[str]] = None,
                 icmp_port: Optional[pulumi.Input[str]] = None,
                 ikev2_port: Optional[pulumi.Input[str]] = None,
                 ikev2_xx_port: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pptp_vpn_port: Optional[pulumi.Input[str]] = None,
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 tls_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip_tcp_port: Optional[pulumi.Input[str]] = None,
                 voip_udp_port: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure connection capability.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.Wirelesscontrollerhotspot20H2qpconncapability("trname",
            esp_port="unknown",
            ftp_port="unknown",
            http_port="unknown",
            icmp_port="closed",
            ikev2_port="unknown",
            ikev2_xx_port="unknown",
            pptp_vpn_port="unknown",
            ssh_port="unknown",
            tls_port="unknown",
            voip_tcp_port="unknown",
            voip_udp_port="unknown")
        ```

        ## Import

        WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerhotspot20H2qpconncapability:Wirelesscontrollerhotspot20H2qpconncapability labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerhotspot20H2qpconncapability:Wirelesscontrollerhotspot20H2qpconncapability labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] esp_port: Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ftp_port: Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] http_port: Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] icmp_port: Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_port: Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_xx_port: Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] name: Connection capability name.
        :param pulumi.Input[str] pptp_vpn_port: Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ssh_port: Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] tls_port: Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip_tcp_port: Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] voip_udp_port: Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[Wirelesscontrollerhotspot20H2qpconncapabilityArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure connection capability.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.Wirelesscontrollerhotspot20H2qpconncapability("trname",
            esp_port="unknown",
            ftp_port="unknown",
            http_port="unknown",
            icmp_port="closed",
            ikev2_port="unknown",
            ikev2_xx_port="unknown",
            pptp_vpn_port="unknown",
            ssh_port="unknown",
            tls_port="unknown",
            voip_tcp_port="unknown",
            voip_udp_port="unknown")
        ```

        ## Import

        WirelessControllerHotspot20 H2QpConnCapability can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerhotspot20H2qpconncapability:Wirelesscontrollerhotspot20H2qpconncapability labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/wirelesscontrollerhotspot20H2qpconncapability:Wirelesscontrollerhotspot20H2qpconncapability labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param Wirelesscontrollerhotspot20H2qpconncapabilityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Wirelesscontrollerhotspot20H2qpconncapabilityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 esp_port: Optional[pulumi.Input[str]] = None,
                 ftp_port: Optional[pulumi.Input[str]] = None,
                 http_port: Optional[pulumi.Input[str]] = None,
                 icmp_port: Optional[pulumi.Input[str]] = None,
                 ikev2_port: Optional[pulumi.Input[str]] = None,
                 ikev2_xx_port: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pptp_vpn_port: Optional[pulumi.Input[str]] = None,
                 ssh_port: Optional[pulumi.Input[str]] = None,
                 tls_port: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 voip_tcp_port: Optional[pulumi.Input[str]] = None,
                 voip_udp_port: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Wirelesscontrollerhotspot20H2qpconncapabilityArgs.__new__(Wirelesscontrollerhotspot20H2qpconncapabilityArgs)

            __props__.__dict__["esp_port"] = esp_port
            __props__.__dict__["ftp_port"] = ftp_port
            __props__.__dict__["http_port"] = http_port
            __props__.__dict__["icmp_port"] = icmp_port
            __props__.__dict__["ikev2_port"] = ikev2_port
            __props__.__dict__["ikev2_xx_port"] = ikev2_xx_port
            __props__.__dict__["name"] = name
            __props__.__dict__["pptp_vpn_port"] = pptp_vpn_port
            __props__.__dict__["ssh_port"] = ssh_port
            __props__.__dict__["tls_port"] = tls_port
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["voip_tcp_port"] = voip_tcp_port
            __props__.__dict__["voip_udp_port"] = voip_udp_port
        super(Wirelesscontrollerhotspot20H2qpconncapability, __self__).__init__(
            'fortios:index/wirelesscontrollerhotspot20H2qpconncapability:Wirelesscontrollerhotspot20H2qpconncapability',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            esp_port: Optional[pulumi.Input[str]] = None,
            ftp_port: Optional[pulumi.Input[str]] = None,
            http_port: Optional[pulumi.Input[str]] = None,
            icmp_port: Optional[pulumi.Input[str]] = None,
            ikev2_port: Optional[pulumi.Input[str]] = None,
            ikev2_xx_port: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pptp_vpn_port: Optional[pulumi.Input[str]] = None,
            ssh_port: Optional[pulumi.Input[str]] = None,
            tls_port: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            voip_tcp_port: Optional[pulumi.Input[str]] = None,
            voip_udp_port: Optional[pulumi.Input[str]] = None) -> 'Wirelesscontrollerhotspot20H2qpconncapability':
        """
        Get an existing Wirelesscontrollerhotspot20H2qpconncapability resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] esp_port: Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ftp_port: Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] http_port: Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] icmp_port: Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_port: Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ikev2_xx_port: Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] name: Connection capability name.
        :param pulumi.Input[str] pptp_vpn_port: Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] ssh_port: Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] tls_port: Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] voip_tcp_port: Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        :param pulumi.Input[str] voip_udp_port: Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Wirelesscontrollerhotspot20H2qpconncapabilityState.__new__(_Wirelesscontrollerhotspot20H2qpconncapabilityState)

        __props__.__dict__["esp_port"] = esp_port
        __props__.__dict__["ftp_port"] = ftp_port
        __props__.__dict__["http_port"] = http_port
        __props__.__dict__["icmp_port"] = icmp_port
        __props__.__dict__["ikev2_port"] = ikev2_port
        __props__.__dict__["ikev2_xx_port"] = ikev2_xx_port
        __props__.__dict__["name"] = name
        __props__.__dict__["pptp_vpn_port"] = pptp_vpn_port
        __props__.__dict__["ssh_port"] = ssh_port
        __props__.__dict__["tls_port"] = tls_port
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["voip_tcp_port"] = voip_tcp_port
        __props__.__dict__["voip_udp_port"] = voip_udp_port
        return Wirelesscontrollerhotspot20H2qpconncapability(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="espPort")
    def esp_port(self) -> pulumi.Output[str]:
        """
        Set ESP port service (used by IPsec VPNs) status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "esp_port")

    @property
    @pulumi.getter(name="ftpPort")
    def ftp_port(self) -> pulumi.Output[str]:
        """
        Set FTP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ftp_port")

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> pulumi.Output[str]:
        """
        Set HTTP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "http_port")

    @property
    @pulumi.getter(name="icmpPort")
    def icmp_port(self) -> pulumi.Output[str]:
        """
        Set ICMP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "icmp_port")

    @property
    @pulumi.getter(name="ikev2Port")
    def ikev2_port(self) -> pulumi.Output[str]:
        """
        Set IKEv2 port service for IPsec VPN status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ikev2_port")

    @property
    @pulumi.getter(name="ikev2XxPort")
    def ikev2_xx_port(self) -> pulumi.Output[str]:
        """
        Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ikev2_xx_port")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Connection capability name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pptpVpnPort")
    def pptp_vpn_port(self) -> pulumi.Output[str]:
        """
        Set Point to Point Tunneling Protocol (PPTP) VPN port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "pptp_vpn_port")

    @property
    @pulumi.getter(name="sshPort")
    def ssh_port(self) -> pulumi.Output[str]:
        """
        Set SSH port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "ssh_port")

    @property
    @pulumi.getter(name="tlsPort")
    def tls_port(self) -> pulumi.Output[str]:
        """
        Set TLS VPN (HTTPS) port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "tls_port")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter(name="voipTcpPort")
    def voip_tcp_port(self) -> pulumi.Output[str]:
        """
        Set VoIP TCP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "voip_tcp_port")

    @property
    @pulumi.getter(name="voipUdpPort")
    def voip_udp_port(self) -> pulumi.Output[str]:
        """
        Set VoIP UDP port service status. Valid values: `closed`, `open`, `unknown`.
        """
        return pulumi.get(self, "voip_udp_port")

