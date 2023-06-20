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

__all__ = [
    'GetFirewallPolicy46Result',
    'AwaitableGetFirewallPolicy46Result',
    'get_firewall_policy46',
    'get_firewall_policy46_output',
]

@pulumi.output_type
class GetFirewallPolicy46Result:
    """
    A collection of values returned by getFirewallPolicy46.
    """
    def __init__(__self__, action=None, comments=None, dstaddrs=None, dstintf=None, fixedport=None, id=None, ippool=None, logtraffic=None, logtraffic_start=None, name=None, per_ip_shaper=None, permit_any_host=None, policyid=None, poolnames=None, schedule=None, services=None, srcaddrs=None, srcintf=None, status=None, tcp_mss_receiver=None, tcp_mss_sender=None, traffic_shaper=None, traffic_shaper_reverse=None, uuid=None, vdomparam=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if comments and not isinstance(comments, str):
            raise TypeError("Expected argument 'comments' to be a str")
        pulumi.set(__self__, "comments", comments)
        if dstaddrs and not isinstance(dstaddrs, list):
            raise TypeError("Expected argument 'dstaddrs' to be a list")
        pulumi.set(__self__, "dstaddrs", dstaddrs)
        if dstintf and not isinstance(dstintf, str):
            raise TypeError("Expected argument 'dstintf' to be a str")
        pulumi.set(__self__, "dstintf", dstintf)
        if fixedport and not isinstance(fixedport, str):
            raise TypeError("Expected argument 'fixedport' to be a str")
        pulumi.set(__self__, "fixedport", fixedport)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ippool and not isinstance(ippool, str):
            raise TypeError("Expected argument 'ippool' to be a str")
        pulumi.set(__self__, "ippool", ippool)
        if logtraffic and not isinstance(logtraffic, str):
            raise TypeError("Expected argument 'logtraffic' to be a str")
        pulumi.set(__self__, "logtraffic", logtraffic)
        if logtraffic_start and not isinstance(logtraffic_start, str):
            raise TypeError("Expected argument 'logtraffic_start' to be a str")
        pulumi.set(__self__, "logtraffic_start", logtraffic_start)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if per_ip_shaper and not isinstance(per_ip_shaper, str):
            raise TypeError("Expected argument 'per_ip_shaper' to be a str")
        pulumi.set(__self__, "per_ip_shaper", per_ip_shaper)
        if permit_any_host and not isinstance(permit_any_host, str):
            raise TypeError("Expected argument 'permit_any_host' to be a str")
        pulumi.set(__self__, "permit_any_host", permit_any_host)
        if policyid and not isinstance(policyid, int):
            raise TypeError("Expected argument 'policyid' to be a int")
        pulumi.set(__self__, "policyid", policyid)
        if poolnames and not isinstance(poolnames, list):
            raise TypeError("Expected argument 'poolnames' to be a list")
        pulumi.set(__self__, "poolnames", poolnames)
        if schedule and not isinstance(schedule, str):
            raise TypeError("Expected argument 'schedule' to be a str")
        pulumi.set(__self__, "schedule", schedule)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)
        if srcaddrs and not isinstance(srcaddrs, list):
            raise TypeError("Expected argument 'srcaddrs' to be a list")
        pulumi.set(__self__, "srcaddrs", srcaddrs)
        if srcintf and not isinstance(srcintf, str):
            raise TypeError("Expected argument 'srcintf' to be a str")
        pulumi.set(__self__, "srcintf", srcintf)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tcp_mss_receiver and not isinstance(tcp_mss_receiver, int):
            raise TypeError("Expected argument 'tcp_mss_receiver' to be a int")
        pulumi.set(__self__, "tcp_mss_receiver", tcp_mss_receiver)
        if tcp_mss_sender and not isinstance(tcp_mss_sender, int):
            raise TypeError("Expected argument 'tcp_mss_sender' to be a int")
        pulumi.set(__self__, "tcp_mss_sender", tcp_mss_sender)
        if traffic_shaper and not isinstance(traffic_shaper, str):
            raise TypeError("Expected argument 'traffic_shaper' to be a str")
        pulumi.set(__self__, "traffic_shaper", traffic_shaper)
        if traffic_shaper_reverse and not isinstance(traffic_shaper_reverse, str):
            raise TypeError("Expected argument 'traffic_shaper_reverse' to be a str")
        pulumi.set(__self__, "traffic_shaper_reverse", traffic_shaper_reverse)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdomparam and not isinstance(vdomparam, str):
            raise TypeError("Expected argument 'vdomparam' to be a str")
        pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Accept or deny traffic matching the policy.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def comments(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def dstaddrs(self) -> Sequence['outputs.GetFirewallPolicy46DstaddrResult']:
        """
        Destination address objects. The structure of `dstaddr` block is documented below.
        """
        return pulumi.get(self, "dstaddrs")

    @property
    @pulumi.getter
    def dstintf(self) -> str:
        """
        Destination interface name.
        """
        return pulumi.get(self, "dstintf")

    @property
    @pulumi.getter
    def fixedport(self) -> str:
        """
        Enable/disable fixed port for this policy.
        """
        return pulumi.get(self, "fixedport")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ippool(self) -> str:
        """
        Enable/disable use of IP Pools for source NAT.
        """
        return pulumi.get(self, "ippool")

    @property
    @pulumi.getter
    def logtraffic(self) -> str:
        """
        Enable/disable traffic logging for this policy.
        """
        return pulumi.get(self, "logtraffic")

    @property
    @pulumi.getter(name="logtrafficStart")
    def logtraffic_start(self) -> str:
        """
        Record logs when a session starts and ends.
        """
        return pulumi.get(self, "logtraffic_start")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        IP pool name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="perIpShaper")
    def per_ip_shaper(self) -> str:
        """
        Per IP traffic shaper.
        """
        return pulumi.get(self, "per_ip_shaper")

    @property
    @pulumi.getter(name="permitAnyHost")
    def permit_any_host(self) -> str:
        """
        Enable/disable allowing any host.
        """
        return pulumi.get(self, "permit_any_host")

    @property
    @pulumi.getter
    def policyid(self) -> int:
        """
        Policy ID.
        """
        return pulumi.get(self, "policyid")

    @property
    @pulumi.getter
    def poolnames(self) -> Sequence['outputs.GetFirewallPolicy46PoolnameResult']:
        """
        IP Pool names. The structure of `poolname` block is documented below.
        """
        return pulumi.get(self, "poolnames")

    @property
    @pulumi.getter
    def schedule(self) -> str:
        """
        Schedule name.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetFirewallPolicy46ServiceResult']:
        """
        Service name. The structure of `service` block is documented below.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def srcaddrs(self) -> Sequence['outputs.GetFirewallPolicy46SrcaddrResult']:
        """
        Source address objects. The structure of `srcaddr` block is documented below.
        """
        return pulumi.get(self, "srcaddrs")

    @property
    @pulumi.getter
    def srcintf(self) -> str:
        """
        Source interface name.
        """
        return pulumi.get(self, "srcintf")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Enable/disable this policy.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tcpMssReceiver")
    def tcp_mss_receiver(self) -> int:
        """
        TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
        """
        return pulumi.get(self, "tcp_mss_receiver")

    @property
    @pulumi.getter(name="tcpMssSender")
    def tcp_mss_sender(self) -> int:
        """
        TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
        """
        return pulumi.get(self, "tcp_mss_sender")

    @property
    @pulumi.getter(name="trafficShaper")
    def traffic_shaper(self) -> str:
        """
        Traffic shaper.
        """
        return pulumi.get(self, "traffic_shaper")

    @property
    @pulumi.getter(name="trafficShaperReverse")
    def traffic_shaper_reverse(self) -> str:
        """
        Reverse traffic shaper.
        """
        return pulumi.get(self, "traffic_shaper_reverse")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdomparam(self) -> Optional[str]:
        return pulumi.get(self, "vdomparam")


class AwaitableGetFirewallPolicy46Result(GetFirewallPolicy46Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallPolicy46Result(
            action=self.action,
            comments=self.comments,
            dstaddrs=self.dstaddrs,
            dstintf=self.dstintf,
            fixedport=self.fixedport,
            id=self.id,
            ippool=self.ippool,
            logtraffic=self.logtraffic,
            logtraffic_start=self.logtraffic_start,
            name=self.name,
            per_ip_shaper=self.per_ip_shaper,
            permit_any_host=self.permit_any_host,
            policyid=self.policyid,
            poolnames=self.poolnames,
            schedule=self.schedule,
            services=self.services,
            srcaddrs=self.srcaddrs,
            srcintf=self.srcintf,
            status=self.status,
            tcp_mss_receiver=self.tcp_mss_receiver,
            tcp_mss_sender=self.tcp_mss_sender,
            traffic_shaper=self.traffic_shaper,
            traffic_shaper_reverse=self.traffic_shaper_reverse,
            uuid=self.uuid,
            vdomparam=self.vdomparam)


def get_firewall_policy46(policyid: Optional[int] = None,
                          vdomparam: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallPolicy46Result:
    """
    Use this data source to get information on an fortios firewall policy46


    :param int policyid: Specify the policyid of the desired firewall policy46.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    __args__ = dict()
    __args__['policyid'] = policyid
    __args__['vdomparam'] = vdomparam
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fortios:index/getFirewallPolicy46:getFirewallPolicy46', __args__, opts=opts, typ=GetFirewallPolicy46Result).value

    return AwaitableGetFirewallPolicy46Result(
        action=__ret__.action,
        comments=__ret__.comments,
        dstaddrs=__ret__.dstaddrs,
        dstintf=__ret__.dstintf,
        fixedport=__ret__.fixedport,
        id=__ret__.id,
        ippool=__ret__.ippool,
        logtraffic=__ret__.logtraffic,
        logtraffic_start=__ret__.logtraffic_start,
        name=__ret__.name,
        per_ip_shaper=__ret__.per_ip_shaper,
        permit_any_host=__ret__.permit_any_host,
        policyid=__ret__.policyid,
        poolnames=__ret__.poolnames,
        schedule=__ret__.schedule,
        services=__ret__.services,
        srcaddrs=__ret__.srcaddrs,
        srcintf=__ret__.srcintf,
        status=__ret__.status,
        tcp_mss_receiver=__ret__.tcp_mss_receiver,
        tcp_mss_sender=__ret__.tcp_mss_sender,
        traffic_shaper=__ret__.traffic_shaper,
        traffic_shaper_reverse=__ret__.traffic_shaper_reverse,
        uuid=__ret__.uuid,
        vdomparam=__ret__.vdomparam)


@_utilities.lift_output_func(get_firewall_policy46)
def get_firewall_policy46_output(policyid: Optional[pulumi.Input[int]] = None,
                                 vdomparam: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFirewallPolicy46Result]:
    """
    Use this data source to get information on an fortios firewall policy46


    :param int policyid: Specify the policyid of the desired firewall policy46.
    :param str vdomparam: Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
    """
    ...
