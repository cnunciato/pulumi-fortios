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

__all__ = ['SystemIpsecaggregateArgs', 'SystemIpsecaggregate']

@pulumi.input_type
class SystemIpsecaggregateArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]],
                 algorithm: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SystemIpsecaggregate resource.
        :param pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]] members: Member tunnels of the aggregate. The structure of `member` block is documented below.
        :param pulumi.Input[str] algorithm: Frame distribution algorithm.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[str] name: IPsec aggregate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        pulumi.set(__self__, "members", members)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]]:
        """
        Member tunnels of the aggregate. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Frame distribution algorithm.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPsec aggregate name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _SystemIpsecaggregateState:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SystemIpsecaggregate resources.
        :param pulumi.Input[str] algorithm: Frame distribution algorithm.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]] members: Member tunnels of the aggregate. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: IPsec aggregate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Frame distribution algorithm.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> Optional[pulumi.Input[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @dynamic_sort_subtable.setter
    def dynamic_sort_subtable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dynamic_sort_subtable", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]]]:
        """
        Member tunnels of the aggregate. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SystemIpsecaggregateMemberArgs']]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IPsec aggregate name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class SystemIpsecaggregate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemIpsecaggregateMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure an aggregate of IPsec tunnels.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname1_vpnipsec_phase1interface = fortios.VpnipsecPhase1interface("trname1VpnipsecPhase1interface",
            acct_verify="disable",
            add_gw_route="disable",
            add_route="enable",
            assign_ip="enable",
            assign_ip_from="range",
            authmethod="psk",
            auto_discovery_forwarder="disable",
            auto_discovery_psk="disable",
            auto_discovery_receiver="disable",
            auto_discovery_sender="disable",
            auto_negotiate="enable",
            cert_id_validation="enable",
            childless_ike="disable",
            client_auto_negotiate="disable",
            client_keep_alive="disable",
            default_gw="0.0.0.0",
            default_gw_priority=0,
            dhgrp="14 5",
            digital_signature_auth="disable",
            distance=15,
            dns_mode="manual",
            dpd="on-demand",
            dpd_retrycount=3,
            dpd_retryinterval="20",
            eap="disable",
            eap_identity="use-id-payload",
            encap_local_gw4="0.0.0.0",
            encap_local_gw6="::",
            encap_remote_gw4="0.0.0.0",
            encap_remote_gw6="::",
            encapsulation="none",
            encapsulation_address="ike",
            enforce_unique_id="disable",
            exchange_interface_ip="disable",
            exchange_ip_addr4="0.0.0.0",
            exchange_ip_addr6="::",
            forticlient_enforcement="disable",
            fragmentation="enable",
            fragmentation_mtu=1200,
            group_authentication="disable",
            ha_sync_esp_seqno="enable",
            idle_timeout="disable",
            idle_timeoutinterval=15,
            ike_version="1",
            include_local_lan="disable",
            interface="port3",
            ip_version="4",
            ipv4_dns_server1="0.0.0.0",
            ipv4_dns_server2="0.0.0.0",
            ipv4_dns_server3="0.0.0.0",
            ipv4_end_ip="0.0.0.0",
            ipv4_netmask="255.255.255.255",
            ipv4_start_ip="0.0.0.0",
            ipv4_wins_server1="0.0.0.0",
            ipv4_wins_server2="0.0.0.0",
            ipv6_dns_server1="::",
            ipv6_dns_server2="::",
            ipv6_dns_server3="::",
            ipv6_end_ip="::",
            ipv6_prefix=128,
            ipv6_start_ip="::",
            keepalive=10,
            keylife=86400,
            local_gw="0.0.0.0",
            local_gw6="::",
            localid_type="auto",
            mesh_selector_type="disable",
            mode="main",
            mode_cfg="disable",
            monitor_hold_down_delay=0,
            monitor_hold_down_time="00:00",
            monitor_hold_down_type="immediate",
            monitor_hold_down_weekday="sunday",
            nattraversal="enable",
            negotiate_timeout=30,
            net_device="disable",
            passive_mode="disable",
            peertype="any",
            ppk="disable",
            priority=0,
            proposal="aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1",
            psksecret="eweeeeeeeecee",
            reauth="disable",
            rekey="enable",
            remote_gw="2.2.2.2",
            remote_gw6="::",
            rsa_signature_format="pkcs1",
            save_password="disable",
            send_cert_chain="enable",
            signature_hash_alg="sha2-512 sha2-384 sha2-256 sha1",
            suite_b="disable",
            tunnel_search="selectors",
            type="static",
            unity_support="enable",
            wizard_type="custom",
            xauthtype="disable")
        trname1_vpnipsec_phase2interface = fortios.VpnipsecPhase2interface("trname1VpnipsecPhase2interface",
            add_route="phase1",
            auto_discovery_forwarder="phase1",
            auto_discovery_sender="phase1",
            auto_negotiate="disable",
            dhcp_ipsec="disable",
            dhgrp="14 5",
            dst_addr_type="subnet",
            dst_end_ip6="::",
            dst_port=0,
            dst_subnet="0.0.0.0 0.0.0.0",
            encapsulation="tunnel-mode",
            keepalive="disable",
            keylife_type="seconds",
            keylifekbs=5120,
            keylifeseconds=43200,
            l2tp="disable",
            pfs="enable",
            phase1name=trname1_vpnipsec_phase1interface.name,
            proposal="aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305",
            protocol=0,
            replay="enable",
            route_overlap="use-new",
            single_source="disable",
            src_addr_type="subnet",
            src_end_ip6="::",
            src_port=0,
            src_subnet="0.0.0.0 0.0.0.0")
        trname = fortios.SystemIpsecaggregate("trname",
            algorithm="round-robin",
            members=[fortios.SystemIpsecaggregateMemberArgs(
                tunnel_name=trname1_vpnipsec_phase1interface.name,
            )])
        ```

        ## Import

        System IpsecAggregate can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemIpsecaggregate:SystemIpsecaggregate labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemIpsecaggregate:SystemIpsecaggregate labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: Frame distribution algorithm.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemIpsecaggregateMemberArgs']]]] members: Member tunnels of the aggregate. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: IPsec aggregate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemIpsecaggregateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure an aggregate of IPsec tunnels.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname1_vpnipsec_phase1interface = fortios.VpnipsecPhase1interface("trname1VpnipsecPhase1interface",
            acct_verify="disable",
            add_gw_route="disable",
            add_route="enable",
            assign_ip="enable",
            assign_ip_from="range",
            authmethod="psk",
            auto_discovery_forwarder="disable",
            auto_discovery_psk="disable",
            auto_discovery_receiver="disable",
            auto_discovery_sender="disable",
            auto_negotiate="enable",
            cert_id_validation="enable",
            childless_ike="disable",
            client_auto_negotiate="disable",
            client_keep_alive="disable",
            default_gw="0.0.0.0",
            default_gw_priority=0,
            dhgrp="14 5",
            digital_signature_auth="disable",
            distance=15,
            dns_mode="manual",
            dpd="on-demand",
            dpd_retrycount=3,
            dpd_retryinterval="20",
            eap="disable",
            eap_identity="use-id-payload",
            encap_local_gw4="0.0.0.0",
            encap_local_gw6="::",
            encap_remote_gw4="0.0.0.0",
            encap_remote_gw6="::",
            encapsulation="none",
            encapsulation_address="ike",
            enforce_unique_id="disable",
            exchange_interface_ip="disable",
            exchange_ip_addr4="0.0.0.0",
            exchange_ip_addr6="::",
            forticlient_enforcement="disable",
            fragmentation="enable",
            fragmentation_mtu=1200,
            group_authentication="disable",
            ha_sync_esp_seqno="enable",
            idle_timeout="disable",
            idle_timeoutinterval=15,
            ike_version="1",
            include_local_lan="disable",
            interface="port3",
            ip_version="4",
            ipv4_dns_server1="0.0.0.0",
            ipv4_dns_server2="0.0.0.0",
            ipv4_dns_server3="0.0.0.0",
            ipv4_end_ip="0.0.0.0",
            ipv4_netmask="255.255.255.255",
            ipv4_start_ip="0.0.0.0",
            ipv4_wins_server1="0.0.0.0",
            ipv4_wins_server2="0.0.0.0",
            ipv6_dns_server1="::",
            ipv6_dns_server2="::",
            ipv6_dns_server3="::",
            ipv6_end_ip="::",
            ipv6_prefix=128,
            ipv6_start_ip="::",
            keepalive=10,
            keylife=86400,
            local_gw="0.0.0.0",
            local_gw6="::",
            localid_type="auto",
            mesh_selector_type="disable",
            mode="main",
            mode_cfg="disable",
            monitor_hold_down_delay=0,
            monitor_hold_down_time="00:00",
            monitor_hold_down_type="immediate",
            monitor_hold_down_weekday="sunday",
            nattraversal="enable",
            negotiate_timeout=30,
            net_device="disable",
            passive_mode="disable",
            peertype="any",
            ppk="disable",
            priority=0,
            proposal="aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1",
            psksecret="eweeeeeeeecee",
            reauth="disable",
            rekey="enable",
            remote_gw="2.2.2.2",
            remote_gw6="::",
            rsa_signature_format="pkcs1",
            save_password="disable",
            send_cert_chain="enable",
            signature_hash_alg="sha2-512 sha2-384 sha2-256 sha1",
            suite_b="disable",
            tunnel_search="selectors",
            type="static",
            unity_support="enable",
            wizard_type="custom",
            xauthtype="disable")
        trname1_vpnipsec_phase2interface = fortios.VpnipsecPhase2interface("trname1VpnipsecPhase2interface",
            add_route="phase1",
            auto_discovery_forwarder="phase1",
            auto_discovery_sender="phase1",
            auto_negotiate="disable",
            dhcp_ipsec="disable",
            dhgrp="14 5",
            dst_addr_type="subnet",
            dst_end_ip6="::",
            dst_port=0,
            dst_subnet="0.0.0.0 0.0.0.0",
            encapsulation="tunnel-mode",
            keepalive="disable",
            keylife_type="seconds",
            keylifekbs=5120,
            keylifeseconds=43200,
            l2tp="disable",
            pfs="enable",
            phase1name=trname1_vpnipsec_phase1interface.name,
            proposal="aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305",
            protocol=0,
            replay="enable",
            route_overlap="use-new",
            single_source="disable",
            src_addr_type="subnet",
            src_end_ip6="::",
            src_port=0,
            src_subnet="0.0.0.0 0.0.0.0")
        trname = fortios.SystemIpsecaggregate("trname",
            algorithm="round-robin",
            members=[fortios.SystemIpsecaggregateMemberArgs(
                tunnel_name=trname1_vpnipsec_phase1interface.name,
            )])
        ```

        ## Import

        System IpsecAggregate can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/systemIpsecaggregate:SystemIpsecaggregate labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/systemIpsecaggregate:SystemIpsecaggregate labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SystemIpsecaggregateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemIpsecaggregateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemIpsecaggregateMemberArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemIpsecaggregateArgs.__new__(SystemIpsecaggregateArgs)

            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(SystemIpsecaggregate, __self__).__init__(
            'fortios:index/systemIpsecaggregate:SystemIpsecaggregate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm: Optional[pulumi.Input[str]] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemIpsecaggregateMemberArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SystemIpsecaggregate':
        """
        Get an existing SystemIpsecaggregate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: Frame distribution algorithm.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SystemIpsecaggregateMemberArgs']]]] members: Member tunnels of the aggregate. The structure of `member` block is documented below.
        :param pulumi.Input[str] name: IPsec aggregate name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemIpsecaggregateState.__new__(_SystemIpsecaggregateState)

        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return SystemIpsecaggregate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[str]:
        """
        Frame distribution algorithm.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence['outputs.SystemIpsecaggregateMember']]:
        """
        Member tunnels of the aggregate. The structure of `member` block is documented below.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        IPsec aggregate name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

