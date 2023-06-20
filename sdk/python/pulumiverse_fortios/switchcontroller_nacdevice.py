# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SwitchcontrollerNacdeviceArgs', 'SwitchcontrollerNacdevice']

@pulumi.input_type
class SwitchcontrollerNacdeviceArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 last_known_port: Optional[pulumi.Input[str]] = None,
                 last_known_switch: Optional[pulumi.Input[str]] = None,
                 last_seen: Optional[pulumi.Input[int]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mac_policy: Optional[pulumi.Input[str]] = None,
                 matched_nac_policy: Optional[pulumi.Input[str]] = None,
                 port_policy: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SwitchcontrollerNacdevice resource.
        :param pulumi.Input[str] description: Description for the learned NAC device.
        :param pulumi.Input[int] fosid: Device ID.
        :param pulumi.Input[str] last_known_port: Managed FortiSwitch port where NAC device is last learned.
        :param pulumi.Input[str] last_known_switch: Managed FortiSwitch where NAC device is last learned.
        :param pulumi.Input[int] last_seen: Device last seen.
        :param pulumi.Input[str] mac: MAC address of the learned NAC device.
        :param pulumi.Input[str] mac_policy: MAC policy to be applied on this learned NAC device.
        :param pulumi.Input[str] matched_nac_policy: Matched NAC policy for the learned NAC device.
        :param pulumi.Input[str] port_policy: Port policy to be applied on this learned NAC device.
        :param pulumi.Input[str] status: Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if last_known_port is not None:
            pulumi.set(__self__, "last_known_port", last_known_port)
        if last_known_switch is not None:
            pulumi.set(__self__, "last_known_switch", last_known_switch)
        if last_seen is not None:
            pulumi.set(__self__, "last_seen", last_seen)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if mac_policy is not None:
            pulumi.set(__self__, "mac_policy", mac_policy)
        if matched_nac_policy is not None:
            pulumi.set(__self__, "matched_nac_policy", matched_nac_policy)
        if port_policy is not None:
            pulumi.set(__self__, "port_policy", port_policy)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the learned NAC device.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Device ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="lastKnownPort")
    def last_known_port(self) -> Optional[pulumi.Input[str]]:
        """
        Managed FortiSwitch port where NAC device is last learned.
        """
        return pulumi.get(self, "last_known_port")

    @last_known_port.setter
    def last_known_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_known_port", value)

    @property
    @pulumi.getter(name="lastKnownSwitch")
    def last_known_switch(self) -> Optional[pulumi.Input[str]]:
        """
        Managed FortiSwitch where NAC device is last learned.
        """
        return pulumi.get(self, "last_known_switch")

    @last_known_switch.setter
    def last_known_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_known_switch", value)

    @property
    @pulumi.getter(name="lastSeen")
    def last_seen(self) -> Optional[pulumi.Input[int]]:
        """
        Device last seen.
        """
        return pulumi.get(self, "last_seen")

    @last_seen.setter
    def last_seen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_seen", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address of the learned NAC device.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="macPolicy")
    def mac_policy(self) -> Optional[pulumi.Input[str]]:
        """
        MAC policy to be applied on this learned NAC device.
        """
        return pulumi.get(self, "mac_policy")

    @mac_policy.setter
    def mac_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_policy", value)

    @property
    @pulumi.getter(name="matchedNacPolicy")
    def matched_nac_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Matched NAC policy for the learned NAC device.
        """
        return pulumi.get(self, "matched_nac_policy")

    @matched_nac_policy.setter
    def matched_nac_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "matched_nac_policy", value)

    @property
    @pulumi.getter(name="portPolicy")
    def port_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Port policy to be applied on this learned NAC device.
        """
        return pulumi.get(self, "port_policy")

    @port_policy.setter
    def port_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_policy", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
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
class _SwitchcontrollerNacdeviceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 last_known_port: Optional[pulumi.Input[str]] = None,
                 last_known_switch: Optional[pulumi.Input[str]] = None,
                 last_seen: Optional[pulumi.Input[int]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mac_policy: Optional[pulumi.Input[str]] = None,
                 matched_nac_policy: Optional[pulumi.Input[str]] = None,
                 port_policy: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SwitchcontrollerNacdevice resources.
        :param pulumi.Input[str] description: Description for the learned NAC device.
        :param pulumi.Input[int] fosid: Device ID.
        :param pulumi.Input[str] last_known_port: Managed FortiSwitch port where NAC device is last learned.
        :param pulumi.Input[str] last_known_switch: Managed FortiSwitch where NAC device is last learned.
        :param pulumi.Input[int] last_seen: Device last seen.
        :param pulumi.Input[str] mac: MAC address of the learned NAC device.
        :param pulumi.Input[str] mac_policy: MAC policy to be applied on this learned NAC device.
        :param pulumi.Input[str] matched_nac_policy: Matched NAC policy for the learned NAC device.
        :param pulumi.Input[str] port_policy: Port policy to be applied on this learned NAC device.
        :param pulumi.Input[str] status: Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fosid is not None:
            pulumi.set(__self__, "fosid", fosid)
        if last_known_port is not None:
            pulumi.set(__self__, "last_known_port", last_known_port)
        if last_known_switch is not None:
            pulumi.set(__self__, "last_known_switch", last_known_switch)
        if last_seen is not None:
            pulumi.set(__self__, "last_seen", last_seen)
        if mac is not None:
            pulumi.set(__self__, "mac", mac)
        if mac_policy is not None:
            pulumi.set(__self__, "mac_policy", mac_policy)
        if matched_nac_policy is not None:
            pulumi.set(__self__, "matched_nac_policy", matched_nac_policy)
        if port_policy is not None:
            pulumi.set(__self__, "port_policy", port_policy)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for the learned NAC device.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def fosid(self) -> Optional[pulumi.Input[int]]:
        """
        Device ID.
        """
        return pulumi.get(self, "fosid")

    @fosid.setter
    def fosid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "fosid", value)

    @property
    @pulumi.getter(name="lastKnownPort")
    def last_known_port(self) -> Optional[pulumi.Input[str]]:
        """
        Managed FortiSwitch port where NAC device is last learned.
        """
        return pulumi.get(self, "last_known_port")

    @last_known_port.setter
    def last_known_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_known_port", value)

    @property
    @pulumi.getter(name="lastKnownSwitch")
    def last_known_switch(self) -> Optional[pulumi.Input[str]]:
        """
        Managed FortiSwitch where NAC device is last learned.
        """
        return pulumi.get(self, "last_known_switch")

    @last_known_switch.setter
    def last_known_switch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_known_switch", value)

    @property
    @pulumi.getter(name="lastSeen")
    def last_seen(self) -> Optional[pulumi.Input[int]]:
        """
        Device last seen.
        """
        return pulumi.get(self, "last_seen")

    @last_seen.setter
    def last_seen(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_seen", value)

    @property
    @pulumi.getter
    def mac(self) -> Optional[pulumi.Input[str]]:
        """
        MAC address of the learned NAC device.
        """
        return pulumi.get(self, "mac")

    @mac.setter
    def mac(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac", value)

    @property
    @pulumi.getter(name="macPolicy")
    def mac_policy(self) -> Optional[pulumi.Input[str]]:
        """
        MAC policy to be applied on this learned NAC device.
        """
        return pulumi.get(self, "mac_policy")

    @mac_policy.setter
    def mac_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mac_policy", value)

    @property
    @pulumi.getter(name="matchedNacPolicy")
    def matched_nac_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Matched NAC policy for the learned NAC device.
        """
        return pulumi.get(self, "matched_nac_policy")

    @matched_nac_policy.setter
    def matched_nac_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "matched_nac_policy", value)

    @property
    @pulumi.getter(name="portPolicy")
    def port_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Port policy to be applied on this learned NAC device.
        """
        return pulumi.get(self, "port_policy")

    @port_policy.setter
    def port_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port_policy", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
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


class SwitchcontrollerNacdevice(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 last_known_port: Optional[pulumi.Input[str]] = None,
                 last_known_switch: Optional[pulumi.Input[str]] = None,
                 last_seen: Optional[pulumi.Input[int]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mac_policy: Optional[pulumi.Input[str]] = None,
                 matched_nac_policy: Optional[pulumi.Input[str]] = None,
                 port_policy: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.

        ## Import

        SwitchController NacDevice can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the learned NAC device.
        :param pulumi.Input[int] fosid: Device ID.
        :param pulumi.Input[str] last_known_port: Managed FortiSwitch port where NAC device is last learned.
        :param pulumi.Input[str] last_known_switch: Managed FortiSwitch where NAC device is last learned.
        :param pulumi.Input[int] last_seen: Device last seen.
        :param pulumi.Input[str] mac: MAC address of the learned NAC device.
        :param pulumi.Input[str] mac_policy: MAC policy to be applied on this learned NAC device.
        :param pulumi.Input[str] matched_nac_policy: Matched NAC policy for the learned NAC device.
        :param pulumi.Input[str] port_policy: Port policy to be applied on this learned NAC device.
        :param pulumi.Input[str] status: Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SwitchcontrollerNacdeviceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure/list NAC devices learned on the managed FortiSwitch ports which matches NAC policy. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,7.0.0`.

        ## Import

        SwitchController NacDevice can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice labelname {{fosid}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice labelname {{fosid}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param SwitchcontrollerNacdeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SwitchcontrollerNacdeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fosid: Optional[pulumi.Input[int]] = None,
                 last_known_port: Optional[pulumi.Input[str]] = None,
                 last_known_switch: Optional[pulumi.Input[str]] = None,
                 last_seen: Optional[pulumi.Input[int]] = None,
                 mac: Optional[pulumi.Input[str]] = None,
                 mac_policy: Optional[pulumi.Input[str]] = None,
                 matched_nac_policy: Optional[pulumi.Input[str]] = None,
                 port_policy: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SwitchcontrollerNacdeviceArgs.__new__(SwitchcontrollerNacdeviceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["fosid"] = fosid
            __props__.__dict__["last_known_port"] = last_known_port
            __props__.__dict__["last_known_switch"] = last_known_switch
            __props__.__dict__["last_seen"] = last_seen
            __props__.__dict__["mac"] = mac
            __props__.__dict__["mac_policy"] = mac_policy
            __props__.__dict__["matched_nac_policy"] = matched_nac_policy
            __props__.__dict__["port_policy"] = port_policy
            __props__.__dict__["status"] = status
            __props__.__dict__["vdomparam"] = vdomparam
        super(SwitchcontrollerNacdevice, __self__).__init__(
            'fortios:index/switchcontrollerNacdevice:SwitchcontrollerNacdevice',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            fosid: Optional[pulumi.Input[int]] = None,
            last_known_port: Optional[pulumi.Input[str]] = None,
            last_known_switch: Optional[pulumi.Input[str]] = None,
            last_seen: Optional[pulumi.Input[int]] = None,
            mac: Optional[pulumi.Input[str]] = None,
            mac_policy: Optional[pulumi.Input[str]] = None,
            matched_nac_policy: Optional[pulumi.Input[str]] = None,
            port_policy: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'SwitchcontrollerNacdevice':
        """
        Get an existing SwitchcontrollerNacdevice resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the learned NAC device.
        :param pulumi.Input[int] fosid: Device ID.
        :param pulumi.Input[str] last_known_port: Managed FortiSwitch port where NAC device is last learned.
        :param pulumi.Input[str] last_known_switch: Managed FortiSwitch where NAC device is last learned.
        :param pulumi.Input[int] last_seen: Device last seen.
        :param pulumi.Input[str] mac: MAC address of the learned NAC device.
        :param pulumi.Input[str] mac_policy: MAC policy to be applied on this learned NAC device.
        :param pulumi.Input[str] matched_nac_policy: Matched NAC policy for the learned NAC device.
        :param pulumi.Input[str] port_policy: Port policy to be applied on this learned NAC device.
        :param pulumi.Input[str] status: Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SwitchcontrollerNacdeviceState.__new__(_SwitchcontrollerNacdeviceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["fosid"] = fosid
        __props__.__dict__["last_known_port"] = last_known_port
        __props__.__dict__["last_known_switch"] = last_known_switch
        __props__.__dict__["last_seen"] = last_seen
        __props__.__dict__["mac"] = mac
        __props__.__dict__["mac_policy"] = mac_policy
        __props__.__dict__["matched_nac_policy"] = matched_nac_policy
        __props__.__dict__["port_policy"] = port_policy
        __props__.__dict__["status"] = status
        __props__.__dict__["vdomparam"] = vdomparam
        return SwitchcontrollerNacdevice(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description for the learned NAC device.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fosid(self) -> pulumi.Output[int]:
        """
        Device ID.
        """
        return pulumi.get(self, "fosid")

    @property
    @pulumi.getter(name="lastKnownPort")
    def last_known_port(self) -> pulumi.Output[str]:
        """
        Managed FortiSwitch port where NAC device is last learned.
        """
        return pulumi.get(self, "last_known_port")

    @property
    @pulumi.getter(name="lastKnownSwitch")
    def last_known_switch(self) -> pulumi.Output[str]:
        """
        Managed FortiSwitch where NAC device is last learned.
        """
        return pulumi.get(self, "last_known_switch")

    @property
    @pulumi.getter(name="lastSeen")
    def last_seen(self) -> pulumi.Output[int]:
        """
        Device last seen.
        """
        return pulumi.get(self, "last_seen")

    @property
    @pulumi.getter
    def mac(self) -> pulumi.Output[str]:
        """
        MAC address of the learned NAC device.
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter(name="macPolicy")
    def mac_policy(self) -> pulumi.Output[str]:
        """
        MAC policy to be applied on this learned NAC device.
        """
        return pulumi.get(self, "mac_policy")

    @property
    @pulumi.getter(name="matchedNacPolicy")
    def matched_nac_policy(self) -> pulumi.Output[str]:
        """
        Matched NAC policy for the learned NAC device.
        """
        return pulumi.get(self, "matched_nac_policy")

    @property
    @pulumi.getter(name="portPolicy")
    def port_policy(self) -> pulumi.Output[str]:
        """
        Port policy to be applied on this learned NAC device.
        """
        return pulumi.get(self, "port_policy")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the learned NAC device. Set enable to authorize the NAC device. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

