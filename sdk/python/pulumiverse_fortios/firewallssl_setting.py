# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FirewallsslSettingArgs', 'FirewallsslSetting']

@pulumi.input_type
class FirewallsslSettingArgs:
    def __init__(__self__, *,
                 cert_cache_capacity: pulumi.Input[int],
                 cert_cache_timeout: pulumi.Input[int],
                 no_matching_cipher_action: pulumi.Input[str],
                 proxy_connect_timeout: pulumi.Input[int],
                 session_cache_capacity: pulumi.Input[int],
                 session_cache_timeout: pulumi.Input[int],
                 ssl_dh_bits: pulumi.Input[str],
                 ssl_send_empty_frags: pulumi.Input[str],
                 abbreviate_handshake: Optional[pulumi.Input[str]] = None,
                 kxp_queue_threshold: Optional[pulumi.Input[int]] = None,
                 ssl_queue_threshold: Optional[pulumi.Input[int]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FirewallsslSetting resource.
        :param pulumi.Input[int] cert_cache_capacity: Maximum capacity of the host certificate cache (0 - 500, default = 200).
        :param pulumi.Input[int] cert_cache_timeout: Time limit to keep certificate cache (1 - 120 min, default = 10).
        :param pulumi.Input[str] no_matching_cipher_action: Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        :param pulumi.Input[int] proxy_connect_timeout: Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        :param pulumi.Input[int] session_cache_capacity: Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        :param pulumi.Input[int] session_cache_timeout: Time limit to keep SSL session state (1 - 60 min, default = 20).
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] abbreviate_handshake: Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] kxp_queue_threshold: Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        :param pulumi.Input[int] ssl_queue_threshold: Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "cert_cache_capacity", cert_cache_capacity)
        pulumi.set(__self__, "cert_cache_timeout", cert_cache_timeout)
        pulumi.set(__self__, "no_matching_cipher_action", no_matching_cipher_action)
        pulumi.set(__self__, "proxy_connect_timeout", proxy_connect_timeout)
        pulumi.set(__self__, "session_cache_capacity", session_cache_capacity)
        pulumi.set(__self__, "session_cache_timeout", session_cache_timeout)
        pulumi.set(__self__, "ssl_dh_bits", ssl_dh_bits)
        pulumi.set(__self__, "ssl_send_empty_frags", ssl_send_empty_frags)
        if abbreviate_handshake is not None:
            pulumi.set(__self__, "abbreviate_handshake", abbreviate_handshake)
        if kxp_queue_threshold is not None:
            pulumi.set(__self__, "kxp_queue_threshold", kxp_queue_threshold)
        if ssl_queue_threshold is not None:
            pulumi.set(__self__, "ssl_queue_threshold", ssl_queue_threshold)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="certCacheCapacity")
    def cert_cache_capacity(self) -> pulumi.Input[int]:
        """
        Maximum capacity of the host certificate cache (0 - 500, default = 200).
        """
        return pulumi.get(self, "cert_cache_capacity")

    @cert_cache_capacity.setter
    def cert_cache_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "cert_cache_capacity", value)

    @property
    @pulumi.getter(name="certCacheTimeout")
    def cert_cache_timeout(self) -> pulumi.Input[int]:
        """
        Time limit to keep certificate cache (1 - 120 min, default = 10).
        """
        return pulumi.get(self, "cert_cache_timeout")

    @cert_cache_timeout.setter
    def cert_cache_timeout(self, value: pulumi.Input[int]):
        pulumi.set(self, "cert_cache_timeout", value)

    @property
    @pulumi.getter(name="noMatchingCipherAction")
    def no_matching_cipher_action(self) -> pulumi.Input[str]:
        """
        Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        """
        return pulumi.get(self, "no_matching_cipher_action")

    @no_matching_cipher_action.setter
    def no_matching_cipher_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "no_matching_cipher_action", value)

    @property
    @pulumi.getter(name="proxyConnectTimeout")
    def proxy_connect_timeout(self) -> pulumi.Input[int]:
        """
        Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        """
        return pulumi.get(self, "proxy_connect_timeout")

    @proxy_connect_timeout.setter
    def proxy_connect_timeout(self, value: pulumi.Input[int]):
        pulumi.set(self, "proxy_connect_timeout", value)

    @property
    @pulumi.getter(name="sessionCacheCapacity")
    def session_cache_capacity(self) -> pulumi.Input[int]:
        """
        Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        """
        return pulumi.get(self, "session_cache_capacity")

    @session_cache_capacity.setter
    def session_cache_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "session_cache_capacity", value)

    @property
    @pulumi.getter(name="sessionCacheTimeout")
    def session_cache_timeout(self) -> pulumi.Input[int]:
        """
        Time limit to keep SSL session state (1 - 60 min, default = 20).
        """
        return pulumi.get(self, "session_cache_timeout")

    @session_cache_timeout.setter
    def session_cache_timeout(self, value: pulumi.Input[int]):
        pulumi.set(self, "session_cache_timeout", value)

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> pulumi.Input[str]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @ssl_dh_bits.setter
    def ssl_dh_bits(self, value: pulumi.Input[str]):
        pulumi.set(self, "ssl_dh_bits", value)

    @property
    @pulumi.getter(name="sslSendEmptyFrags")
    def ssl_send_empty_frags(self) -> pulumi.Input[str]:
        """
        Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_send_empty_frags")

    @ssl_send_empty_frags.setter
    def ssl_send_empty_frags(self, value: pulumi.Input[str]):
        pulumi.set(self, "ssl_send_empty_frags", value)

    @property
    @pulumi.getter(name="abbreviateHandshake")
    def abbreviate_handshake(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "abbreviate_handshake")

    @abbreviate_handshake.setter
    def abbreviate_handshake(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "abbreviate_handshake", value)

    @property
    @pulumi.getter(name="kxpQueueThreshold")
    def kxp_queue_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        """
        return pulumi.get(self, "kxp_queue_threshold")

    @kxp_queue_threshold.setter
    def kxp_queue_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kxp_queue_threshold", value)

    @property
    @pulumi.getter(name="sslQueueThreshold")
    def ssl_queue_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        """
        return pulumi.get(self, "ssl_queue_threshold")

    @ssl_queue_threshold.setter
    def ssl_queue_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ssl_queue_threshold", value)

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
class _FirewallsslSettingState:
    def __init__(__self__, *,
                 abbreviate_handshake: Optional[pulumi.Input[str]] = None,
                 cert_cache_capacity: Optional[pulumi.Input[int]] = None,
                 cert_cache_timeout: Optional[pulumi.Input[int]] = None,
                 kxp_queue_threshold: Optional[pulumi.Input[int]] = None,
                 no_matching_cipher_action: Optional[pulumi.Input[str]] = None,
                 proxy_connect_timeout: Optional[pulumi.Input[int]] = None,
                 session_cache_capacity: Optional[pulumi.Input[int]] = None,
                 session_cache_timeout: Optional[pulumi.Input[int]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_queue_threshold: Optional[pulumi.Input[int]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FirewallsslSetting resources.
        :param pulumi.Input[str] abbreviate_handshake: Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] cert_cache_capacity: Maximum capacity of the host certificate cache (0 - 500, default = 200).
        :param pulumi.Input[int] cert_cache_timeout: Time limit to keep certificate cache (1 - 120 min, default = 10).
        :param pulumi.Input[int] kxp_queue_threshold: Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        :param pulumi.Input[str] no_matching_cipher_action: Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        :param pulumi.Input[int] proxy_connect_timeout: Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        :param pulumi.Input[int] session_cache_capacity: Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        :param pulumi.Input[int] session_cache_timeout: Time limit to keep SSL session state (1 - 60 min, default = 20).
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[int] ssl_queue_threshold: Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if abbreviate_handshake is not None:
            pulumi.set(__self__, "abbreviate_handshake", abbreviate_handshake)
        if cert_cache_capacity is not None:
            pulumi.set(__self__, "cert_cache_capacity", cert_cache_capacity)
        if cert_cache_timeout is not None:
            pulumi.set(__self__, "cert_cache_timeout", cert_cache_timeout)
        if kxp_queue_threshold is not None:
            pulumi.set(__self__, "kxp_queue_threshold", kxp_queue_threshold)
        if no_matching_cipher_action is not None:
            pulumi.set(__self__, "no_matching_cipher_action", no_matching_cipher_action)
        if proxy_connect_timeout is not None:
            pulumi.set(__self__, "proxy_connect_timeout", proxy_connect_timeout)
        if session_cache_capacity is not None:
            pulumi.set(__self__, "session_cache_capacity", session_cache_capacity)
        if session_cache_timeout is not None:
            pulumi.set(__self__, "session_cache_timeout", session_cache_timeout)
        if ssl_dh_bits is not None:
            pulumi.set(__self__, "ssl_dh_bits", ssl_dh_bits)
        if ssl_queue_threshold is not None:
            pulumi.set(__self__, "ssl_queue_threshold", ssl_queue_threshold)
        if ssl_send_empty_frags is not None:
            pulumi.set(__self__, "ssl_send_empty_frags", ssl_send_empty_frags)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter(name="abbreviateHandshake")
    def abbreviate_handshake(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "abbreviate_handshake")

    @abbreviate_handshake.setter
    def abbreviate_handshake(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "abbreviate_handshake", value)

    @property
    @pulumi.getter(name="certCacheCapacity")
    def cert_cache_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum capacity of the host certificate cache (0 - 500, default = 200).
        """
        return pulumi.get(self, "cert_cache_capacity")

    @cert_cache_capacity.setter
    def cert_cache_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cert_cache_capacity", value)

    @property
    @pulumi.getter(name="certCacheTimeout")
    def cert_cache_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Time limit to keep certificate cache (1 - 120 min, default = 10).
        """
        return pulumi.get(self, "cert_cache_timeout")

    @cert_cache_timeout.setter
    def cert_cache_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cert_cache_timeout", value)

    @property
    @pulumi.getter(name="kxpQueueThreshold")
    def kxp_queue_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        """
        return pulumi.get(self, "kxp_queue_threshold")

    @kxp_queue_threshold.setter
    def kxp_queue_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kxp_queue_threshold", value)

    @property
    @pulumi.getter(name="noMatchingCipherAction")
    def no_matching_cipher_action(self) -> Optional[pulumi.Input[str]]:
        """
        Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        """
        return pulumi.get(self, "no_matching_cipher_action")

    @no_matching_cipher_action.setter
    def no_matching_cipher_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "no_matching_cipher_action", value)

    @property
    @pulumi.getter(name="proxyConnectTimeout")
    def proxy_connect_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        """
        return pulumi.get(self, "proxy_connect_timeout")

    @proxy_connect_timeout.setter
    def proxy_connect_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "proxy_connect_timeout", value)

    @property
    @pulumi.getter(name="sessionCacheCapacity")
    def session_cache_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        """
        return pulumi.get(self, "session_cache_capacity")

    @session_cache_capacity.setter
    def session_cache_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_cache_capacity", value)

    @property
    @pulumi.getter(name="sessionCacheTimeout")
    def session_cache_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Time limit to keep SSL session state (1 - 60 min, default = 20).
        """
        return pulumi.get(self, "session_cache_timeout")

    @session_cache_timeout.setter
    def session_cache_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_cache_timeout", value)

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> Optional[pulumi.Input[str]]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @ssl_dh_bits.setter
    def ssl_dh_bits(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_dh_bits", value)

    @property
    @pulumi.getter(name="sslQueueThreshold")
    def ssl_queue_threshold(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        """
        return pulumi.get(self, "ssl_queue_threshold")

    @ssl_queue_threshold.setter
    def ssl_queue_threshold(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ssl_queue_threshold", value)

    @property
    @pulumi.getter(name="sslSendEmptyFrags")
    def ssl_send_empty_frags(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_send_empty_frags")

    @ssl_send_empty_frags.setter
    def ssl_send_empty_frags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_send_empty_frags", value)

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


class FirewallsslSetting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 abbreviate_handshake: Optional[pulumi.Input[str]] = None,
                 cert_cache_capacity: Optional[pulumi.Input[int]] = None,
                 cert_cache_timeout: Optional[pulumi.Input[int]] = None,
                 kxp_queue_threshold: Optional[pulumi.Input[int]] = None,
                 no_matching_cipher_action: Optional[pulumi.Input[str]] = None,
                 proxy_connect_timeout: Optional[pulumi.Input[int]] = None,
                 session_cache_capacity: Optional[pulumi.Input[int]] = None,
                 session_cache_timeout: Optional[pulumi.Input[int]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_queue_threshold: Optional[pulumi.Input[int]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        SSL proxy settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.FirewallsslSetting("trname",
            abbreviate_handshake="enable",
            cert_cache_capacity=200,
            cert_cache_timeout=10,
            kxp_queue_threshold=16,
            no_matching_cipher_action="bypass",
            proxy_connect_timeout=30,
            session_cache_capacity=500,
            session_cache_timeout=20,
            ssl_dh_bits="2048",
            ssl_queue_threshold=32,
            ssl_send_empty_frags="enable")
        ```

        ## Import

        FirewallSsl Setting can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallsslSetting:FirewallsslSetting labelname FirewallSslSetting
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallsslSetting:FirewallsslSetting labelname FirewallSslSetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] abbreviate_handshake: Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] cert_cache_capacity: Maximum capacity of the host certificate cache (0 - 500, default = 200).
        :param pulumi.Input[int] cert_cache_timeout: Time limit to keep certificate cache (1 - 120 min, default = 10).
        :param pulumi.Input[int] kxp_queue_threshold: Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        :param pulumi.Input[str] no_matching_cipher_action: Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        :param pulumi.Input[int] proxy_connect_timeout: Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        :param pulumi.Input[int] session_cache_capacity: Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        :param pulumi.Input[int] session_cache_timeout: Time limit to keep SSL session state (1 - 60 min, default = 20).
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[int] ssl_queue_threshold: Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallsslSettingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        SSL proxy settings.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.FirewallsslSetting("trname",
            abbreviate_handshake="enable",
            cert_cache_capacity=200,
            cert_cache_timeout=10,
            kxp_queue_threshold=16,
            no_matching_cipher_action="bypass",
            proxy_connect_timeout=30,
            session_cache_capacity=500,
            session_cache_timeout=20,
            ssl_dh_bits="2048",
            ssl_queue_threshold=32,
            ssl_send_empty_frags="enable")
        ```

        ## Import

        FirewallSsl Setting can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/firewallsslSetting:FirewallsslSetting labelname FirewallSslSetting
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/firewallsslSetting:FirewallsslSetting labelname FirewallSslSetting
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param FirewallsslSettingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallsslSettingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 abbreviate_handshake: Optional[pulumi.Input[str]] = None,
                 cert_cache_capacity: Optional[pulumi.Input[int]] = None,
                 cert_cache_timeout: Optional[pulumi.Input[int]] = None,
                 kxp_queue_threshold: Optional[pulumi.Input[int]] = None,
                 no_matching_cipher_action: Optional[pulumi.Input[str]] = None,
                 proxy_connect_timeout: Optional[pulumi.Input[int]] = None,
                 session_cache_capacity: Optional[pulumi.Input[int]] = None,
                 session_cache_timeout: Optional[pulumi.Input[int]] = None,
                 ssl_dh_bits: Optional[pulumi.Input[str]] = None,
                 ssl_queue_threshold: Optional[pulumi.Input[int]] = None,
                 ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallsslSettingArgs.__new__(FirewallsslSettingArgs)

            __props__.__dict__["abbreviate_handshake"] = abbreviate_handshake
            if cert_cache_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'cert_cache_capacity'")
            __props__.__dict__["cert_cache_capacity"] = cert_cache_capacity
            if cert_cache_timeout is None and not opts.urn:
                raise TypeError("Missing required property 'cert_cache_timeout'")
            __props__.__dict__["cert_cache_timeout"] = cert_cache_timeout
            __props__.__dict__["kxp_queue_threshold"] = kxp_queue_threshold
            if no_matching_cipher_action is None and not opts.urn:
                raise TypeError("Missing required property 'no_matching_cipher_action'")
            __props__.__dict__["no_matching_cipher_action"] = no_matching_cipher_action
            if proxy_connect_timeout is None and not opts.urn:
                raise TypeError("Missing required property 'proxy_connect_timeout'")
            __props__.__dict__["proxy_connect_timeout"] = proxy_connect_timeout
            if session_cache_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'session_cache_capacity'")
            __props__.__dict__["session_cache_capacity"] = session_cache_capacity
            if session_cache_timeout is None and not opts.urn:
                raise TypeError("Missing required property 'session_cache_timeout'")
            __props__.__dict__["session_cache_timeout"] = session_cache_timeout
            if ssl_dh_bits is None and not opts.urn:
                raise TypeError("Missing required property 'ssl_dh_bits'")
            __props__.__dict__["ssl_dh_bits"] = ssl_dh_bits
            __props__.__dict__["ssl_queue_threshold"] = ssl_queue_threshold
            if ssl_send_empty_frags is None and not opts.urn:
                raise TypeError("Missing required property 'ssl_send_empty_frags'")
            __props__.__dict__["ssl_send_empty_frags"] = ssl_send_empty_frags
            __props__.__dict__["vdomparam"] = vdomparam
        super(FirewallsslSetting, __self__).__init__(
            'fortios:index/firewallsslSetting:FirewallsslSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            abbreviate_handshake: Optional[pulumi.Input[str]] = None,
            cert_cache_capacity: Optional[pulumi.Input[int]] = None,
            cert_cache_timeout: Optional[pulumi.Input[int]] = None,
            kxp_queue_threshold: Optional[pulumi.Input[int]] = None,
            no_matching_cipher_action: Optional[pulumi.Input[str]] = None,
            proxy_connect_timeout: Optional[pulumi.Input[int]] = None,
            session_cache_capacity: Optional[pulumi.Input[int]] = None,
            session_cache_timeout: Optional[pulumi.Input[int]] = None,
            ssl_dh_bits: Optional[pulumi.Input[str]] = None,
            ssl_queue_threshold: Optional[pulumi.Input[int]] = None,
            ssl_send_empty_frags: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'FirewallsslSetting':
        """
        Get an existing FirewallsslSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] abbreviate_handshake: Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        :param pulumi.Input[int] cert_cache_capacity: Maximum capacity of the host certificate cache (0 - 500, default = 200).
        :param pulumi.Input[int] cert_cache_timeout: Time limit to keep certificate cache (1 - 120 min, default = 10).
        :param pulumi.Input[int] kxp_queue_threshold: Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        :param pulumi.Input[str] no_matching_cipher_action: Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        :param pulumi.Input[int] proxy_connect_timeout: Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        :param pulumi.Input[int] session_cache_capacity: Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        :param pulumi.Input[int] session_cache_timeout: Time limit to keep SSL session state (1 - 60 min, default = 20).
        :param pulumi.Input[str] ssl_dh_bits: Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        :param pulumi.Input[int] ssl_queue_threshold: Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        :param pulumi.Input[str] ssl_send_empty_frags: Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallsslSettingState.__new__(_FirewallsslSettingState)

        __props__.__dict__["abbreviate_handshake"] = abbreviate_handshake
        __props__.__dict__["cert_cache_capacity"] = cert_cache_capacity
        __props__.__dict__["cert_cache_timeout"] = cert_cache_timeout
        __props__.__dict__["kxp_queue_threshold"] = kxp_queue_threshold
        __props__.__dict__["no_matching_cipher_action"] = no_matching_cipher_action
        __props__.__dict__["proxy_connect_timeout"] = proxy_connect_timeout
        __props__.__dict__["session_cache_capacity"] = session_cache_capacity
        __props__.__dict__["session_cache_timeout"] = session_cache_timeout
        __props__.__dict__["ssl_dh_bits"] = ssl_dh_bits
        __props__.__dict__["ssl_queue_threshold"] = ssl_queue_threshold
        __props__.__dict__["ssl_send_empty_frags"] = ssl_send_empty_frags
        __props__.__dict__["vdomparam"] = vdomparam
        return FirewallsslSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="abbreviateHandshake")
    def abbreviate_handshake(self) -> pulumi.Output[str]:
        """
        Enable/disable use of SSL abbreviated handshake. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "abbreviate_handshake")

    @property
    @pulumi.getter(name="certCacheCapacity")
    def cert_cache_capacity(self) -> pulumi.Output[int]:
        """
        Maximum capacity of the host certificate cache (0 - 500, default = 200).
        """
        return pulumi.get(self, "cert_cache_capacity")

    @property
    @pulumi.getter(name="certCacheTimeout")
    def cert_cache_timeout(self) -> pulumi.Output[int]:
        """
        Time limit to keep certificate cache (1 - 120 min, default = 10).
        """
        return pulumi.get(self, "cert_cache_timeout")

    @property
    @pulumi.getter(name="kxpQueueThreshold")
    def kxp_queue_threshold(self) -> pulumi.Output[int]:
        """
        Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).
        """
        return pulumi.get(self, "kxp_queue_threshold")

    @property
    @pulumi.getter(name="noMatchingCipherAction")
    def no_matching_cipher_action(self) -> pulumi.Output[str]:
        """
        Bypass or drop the connection when no matching cipher is found. Valid values: `bypass`, `drop`.
        """
        return pulumi.get(self, "no_matching_cipher_action")

    @property
    @pulumi.getter(name="proxyConnectTimeout")
    def proxy_connect_timeout(self) -> pulumi.Output[int]:
        """
        Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).
        """
        return pulumi.get(self, "proxy_connect_timeout")

    @property
    @pulumi.getter(name="sessionCacheCapacity")
    def session_cache_capacity(self) -> pulumi.Output[int]:
        """
        Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).
        """
        return pulumi.get(self, "session_cache_capacity")

    @property
    @pulumi.getter(name="sessionCacheTimeout")
    def session_cache_timeout(self) -> pulumi.Output[int]:
        """
        Time limit to keep SSL session state (1 - 60 min, default = 20).
        """
        return pulumi.get(self, "session_cache_timeout")

    @property
    @pulumi.getter(name="sslDhBits")
    def ssl_dh_bits(self) -> pulumi.Output[str]:
        """
        Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
        """
        return pulumi.get(self, "ssl_dh_bits")

    @property
    @pulumi.getter(name="sslQueueThreshold")
    def ssl_queue_threshold(self) -> pulumi.Output[int]:
        """
        Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).
        """
        return pulumi.get(self, "ssl_queue_threshold")

    @property
    @pulumi.getter(name="sslSendEmptyFrags")
    def ssl_send_empty_frags(self) -> pulumi.Output[str]:
        """
        Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only). Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "ssl_send_empty_frags")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

