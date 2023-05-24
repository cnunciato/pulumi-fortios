# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFirewallDoSpolicyAnomalyResult',
    'GetFirewallDoSpolicyDstaddrResult',
    'GetFirewallDoSpolicyServiceResult',
    'GetFirewallDoSpolicySrcaddrResult',
]

@pulumi.output_type
class GetFirewallDoSpolicyAnomalyResult(dict):
    def __init__(__self__, *,
                 action: str,
                 log: str,
                 name: str,
                 quarantine: str,
                 quarantine_expiry: str,
                 quarantine_log: str,
                 status: str,
                 threshold: int,
                 thresholddefault: int):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "log", log)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "quarantine", quarantine)
        pulumi.set(__self__, "quarantine_expiry", quarantine_expiry)
        pulumi.set(__self__, "quarantine_log", quarantine_log)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "threshold", threshold)
        pulumi.set(__self__, "thresholddefault", thresholddefault)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def log(self) -> str:
        return pulumi.get(self, "log")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def quarantine(self) -> str:
        return pulumi.get(self, "quarantine")

    @property
    @pulumi.getter(name="quarantineExpiry")
    def quarantine_expiry(self) -> str:
        return pulumi.get(self, "quarantine_expiry")

    @property
    @pulumi.getter(name="quarantineLog")
    def quarantine_log(self) -> str:
        return pulumi.get(self, "quarantine_log")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def threshold(self) -> int:
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter
    def thresholddefault(self) -> int:
        return pulumi.get(self, "thresholddefault")


@pulumi.output_type
class GetFirewallDoSpolicyDstaddrResult(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallDoSpolicyServiceResult(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class GetFirewallDoSpolicySrcaddrResult(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

