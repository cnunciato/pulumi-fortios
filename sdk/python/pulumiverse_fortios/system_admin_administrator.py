# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SystemAdminAdministratorArgs', 'SystemAdminAdministrator']

@pulumi.input_type
class SystemAdminAdministratorArgs:
    def __init__(__self__, *,
                 accprofile: pulumi.Input[str],
                 password: pulumi.Input[str],
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 trusthost1: Optional[pulumi.Input[str]] = None,
                 trusthost10: Optional[pulumi.Input[str]] = None,
                 trusthost2: Optional[pulumi.Input[str]] = None,
                 trusthost3: Optional[pulumi.Input[str]] = None,
                 trusthost4: Optional[pulumi.Input[str]] = None,
                 trusthost5: Optional[pulumi.Input[str]] = None,
                 trusthost6: Optional[pulumi.Input[str]] = None,
                 trusthost7: Optional[pulumi.Input[str]] = None,
                 trusthost8: Optional[pulumi.Input[str]] = None,
                 trusthost9: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SystemAdminAdministrator resource.
        :param pulumi.Input[str] accprofile: Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        :param pulumi.Input[str] password: Admin user password.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domain(s) that the administrator can access.
        """
        pulumi.set(__self__, "accprofile", accprofile)
        pulumi.set(__self__, "password", password)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if trusthost1 is not None:
            pulumi.set(__self__, "trusthost1", trusthost1)
        if trusthost10 is not None:
            pulumi.set(__self__, "trusthost10", trusthost10)
        if trusthost2 is not None:
            pulumi.set(__self__, "trusthost2", trusthost2)
        if trusthost3 is not None:
            pulumi.set(__self__, "trusthost3", trusthost3)
        if trusthost4 is not None:
            pulumi.set(__self__, "trusthost4", trusthost4)
        if trusthost5 is not None:
            pulumi.set(__self__, "trusthost5", trusthost5)
        if trusthost6 is not None:
            pulumi.set(__self__, "trusthost6", trusthost6)
        if trusthost7 is not None:
            pulumi.set(__self__, "trusthost7", trusthost7)
        if trusthost8 is not None:
            pulumi.set(__self__, "trusthost8", trusthost8)
        if trusthost9 is not None:
            pulumi.set(__self__, "trusthost9", trusthost9)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Input[str]:
        """
        Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        """
        return pulumi.get(self, "accprofile")

    @accprofile.setter
    def accprofile(self, value: pulumi.Input[str]):
        pulumi.set(self, "accprofile", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Admin user password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def trusthost1(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost1")

    @trusthost1.setter
    def trusthost1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost1", value)

    @property
    @pulumi.getter
    def trusthost10(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost10")

    @trusthost10.setter
    def trusthost10(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost10", value)

    @property
    @pulumi.getter
    def trusthost2(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost2")

    @trusthost2.setter
    def trusthost2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost2", value)

    @property
    @pulumi.getter
    def trusthost3(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost3")

    @trusthost3.setter
    def trusthost3(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost3", value)

    @property
    @pulumi.getter
    def trusthost4(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost4")

    @trusthost4.setter
    def trusthost4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost4", value)

    @property
    @pulumi.getter
    def trusthost5(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost5")

    @trusthost5.setter
    def trusthost5(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost5", value)

    @property
    @pulumi.getter
    def trusthost6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost6")

    @trusthost6.setter
    def trusthost6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost6", value)

    @property
    @pulumi.getter
    def trusthost7(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost7")

    @trusthost7.setter
    def trusthost7(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost7", value)

    @property
    @pulumi.getter
    def trusthost8(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost8")

    @trusthost8.setter
    def trusthost8(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost8", value)

    @property
    @pulumi.getter
    def trusthost9(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost9")

    @trusthost9.setter
    def trusthost9(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost9", value)

    @property
    @pulumi.getter
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Virtual domain(s) that the administrator can access.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vdoms", value)


@pulumi.input_type
class _SystemAdminAdministratorState:
    def __init__(__self__, *,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 trusthost1: Optional[pulumi.Input[str]] = None,
                 trusthost10: Optional[pulumi.Input[str]] = None,
                 trusthost2: Optional[pulumi.Input[str]] = None,
                 trusthost3: Optional[pulumi.Input[str]] = None,
                 trusthost4: Optional[pulumi.Input[str]] = None,
                 trusthost5: Optional[pulumi.Input[str]] = None,
                 trusthost6: Optional[pulumi.Input[str]] = None,
                 trusthost7: Optional[pulumi.Input[str]] = None,
                 trusthost8: Optional[pulumi.Input[str]] = None,
                 trusthost9: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SystemAdminAdministrator resources.
        :param pulumi.Input[str] accprofile: Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[str] password: Admin user password.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domain(s) that the administrator can access.
        """
        if accprofile is not None:
            pulumi.set(__self__, "accprofile", accprofile)
        if comments is not None:
            pulumi.set(__self__, "comments", comments)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if trusthost1 is not None:
            pulumi.set(__self__, "trusthost1", trusthost1)
        if trusthost10 is not None:
            pulumi.set(__self__, "trusthost10", trusthost10)
        if trusthost2 is not None:
            pulumi.set(__self__, "trusthost2", trusthost2)
        if trusthost3 is not None:
            pulumi.set(__self__, "trusthost3", trusthost3)
        if trusthost4 is not None:
            pulumi.set(__self__, "trusthost4", trusthost4)
        if trusthost5 is not None:
            pulumi.set(__self__, "trusthost5", trusthost5)
        if trusthost6 is not None:
            pulumi.set(__self__, "trusthost6", trusthost6)
        if trusthost7 is not None:
            pulumi.set(__self__, "trusthost7", trusthost7)
        if trusthost8 is not None:
            pulumi.set(__self__, "trusthost8", trusthost8)
        if trusthost9 is not None:
            pulumi.set(__self__, "trusthost9", trusthost9)
        if vdoms is not None:
            pulumi.set(__self__, "vdoms", vdoms)

    @property
    @pulumi.getter
    def accprofile(self) -> Optional[pulumi.Input[str]]:
        """
        Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        """
        return pulumi.get(self, "accprofile")

    @accprofile.setter
    def accprofile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accprofile", value)

    @property
    @pulumi.getter
    def comments(self) -> Optional[pulumi.Input[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @comments.setter
    def comments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comments", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Admin user password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def trusthost1(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost1")

    @trusthost1.setter
    def trusthost1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost1", value)

    @property
    @pulumi.getter
    def trusthost10(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost10")

    @trusthost10.setter
    def trusthost10(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost10", value)

    @property
    @pulumi.getter
    def trusthost2(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost2")

    @trusthost2.setter
    def trusthost2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost2", value)

    @property
    @pulumi.getter
    def trusthost3(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost3")

    @trusthost3.setter
    def trusthost3(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost3", value)

    @property
    @pulumi.getter
    def trusthost4(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost4")

    @trusthost4.setter
    def trusthost4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost4", value)

    @property
    @pulumi.getter
    def trusthost5(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost5")

    @trusthost5.setter
    def trusthost5(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost5", value)

    @property
    @pulumi.getter
    def trusthost6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost6")

    @trusthost6.setter
    def trusthost6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost6", value)

    @property
    @pulumi.getter
    def trusthost7(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost7")

    @trusthost7.setter
    def trusthost7(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost7", value)

    @property
    @pulumi.getter
    def trusthost8(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost8")

    @trusthost8.setter
    def trusthost8(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost8", value)

    @property
    @pulumi.getter
    def trusthost9(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "trusthost9")

    @trusthost9.setter
    def trusthost9(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusthost9", value)

    @property
    @pulumi.getter
    def vdoms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Virtual domain(s) that the administrator can access.
        """
        return pulumi.get(self, "vdoms")

    @vdoms.setter
    def vdoms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "vdoms", value)


class SystemAdminAdministrator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 trusthost1: Optional[pulumi.Input[str]] = None,
                 trusthost10: Optional[pulumi.Input[str]] = None,
                 trusthost2: Optional[pulumi.Input[str]] = None,
                 trusthost3: Optional[pulumi.Input[str]] = None,
                 trusthost4: Optional[pulumi.Input[str]] = None,
                 trusthost5: Optional[pulumi.Input[str]] = None,
                 trusthost6: Optional[pulumi.Input[str]] = None,
                 trusthost7: Optional[pulumi.Input[str]] = None,
                 trusthost8: Optional[pulumi.Input[str]] = None,
                 trusthost9: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to configure administrator accounts of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemAdmin`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        admintest = fortios.SystemAdminAdministrator("admintest",
            accprofile="3d3",
            comments="comments",
            password="cc37331AC1",
            trusthost1="1.1.1.0 255.255.255.0",
            trusthost2="2.2.2.0 255.255.255.0",
            vdoms=["root"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[str] password: Admin user password.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domain(s) that the administrator can access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SystemAdminAdministratorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to configure administrator accounts of FortiOS.

        !> **Warning:** The resource will be deprecated and replaced by new resource `SystemAdmin`, we recommend that you use the new resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        admintest = fortios.SystemAdminAdministrator("admintest",
            accprofile="3d3",
            comments="comments",
            password="cc37331AC1",
            trusthost1="1.1.1.0 255.255.255.0",
            trusthost2="2.2.2.0 255.255.255.0",
            vdoms=["root"])
        ```

        :param str resource_name: The name of the resource.
        :param SystemAdminAdministratorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SystemAdminAdministratorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accprofile: Optional[pulumi.Input[str]] = None,
                 comments: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 trusthost1: Optional[pulumi.Input[str]] = None,
                 trusthost10: Optional[pulumi.Input[str]] = None,
                 trusthost2: Optional[pulumi.Input[str]] = None,
                 trusthost3: Optional[pulumi.Input[str]] = None,
                 trusthost4: Optional[pulumi.Input[str]] = None,
                 trusthost5: Optional[pulumi.Input[str]] = None,
                 trusthost6: Optional[pulumi.Input[str]] = None,
                 trusthost7: Optional[pulumi.Input[str]] = None,
                 trusthost8: Optional[pulumi.Input[str]] = None,
                 trusthost9: Optional[pulumi.Input[str]] = None,
                 vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SystemAdminAdministratorArgs.__new__(SystemAdminAdministratorArgs)

            if accprofile is None and not opts.urn:
                raise TypeError("Missing required property 'accprofile'")
            __props__.__dict__["accprofile"] = accprofile
            __props__.__dict__["comments"] = comments
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            __props__.__dict__["trusthost1"] = trusthost1
            __props__.__dict__["trusthost10"] = trusthost10
            __props__.__dict__["trusthost2"] = trusthost2
            __props__.__dict__["trusthost3"] = trusthost3
            __props__.__dict__["trusthost4"] = trusthost4
            __props__.__dict__["trusthost5"] = trusthost5
            __props__.__dict__["trusthost6"] = trusthost6
            __props__.__dict__["trusthost7"] = trusthost7
            __props__.__dict__["trusthost8"] = trusthost8
            __props__.__dict__["trusthost9"] = trusthost9
            __props__.__dict__["vdoms"] = vdoms
        super(SystemAdminAdministrator, __self__).__init__(
            'fortios:index/systemAdminAdministrator:SystemAdminAdministrator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accprofile: Optional[pulumi.Input[str]] = None,
            comments: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            trusthost1: Optional[pulumi.Input[str]] = None,
            trusthost10: Optional[pulumi.Input[str]] = None,
            trusthost2: Optional[pulumi.Input[str]] = None,
            trusthost3: Optional[pulumi.Input[str]] = None,
            trusthost4: Optional[pulumi.Input[str]] = None,
            trusthost5: Optional[pulumi.Input[str]] = None,
            trusthost6: Optional[pulumi.Input[str]] = None,
            trusthost7: Optional[pulumi.Input[str]] = None,
            trusthost8: Optional[pulumi.Input[str]] = None,
            trusthost9: Optional[pulumi.Input[str]] = None,
            vdoms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SystemAdminAdministrator':
        """
        Get an existing SystemAdminAdministrator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accprofile: Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        :param pulumi.Input[str] comments: Comment.
        :param pulumi.Input[str] name: User name.
        :param pulumi.Input[str] password: Admin user password.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vdoms: Virtual domain(s) that the administrator can access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SystemAdminAdministratorState.__new__(_SystemAdminAdministratorState)

        __props__.__dict__["accprofile"] = accprofile
        __props__.__dict__["comments"] = comments
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["trusthost1"] = trusthost1
        __props__.__dict__["trusthost10"] = trusthost10
        __props__.__dict__["trusthost2"] = trusthost2
        __props__.__dict__["trusthost3"] = trusthost3
        __props__.__dict__["trusthost4"] = trusthost4
        __props__.__dict__["trusthost5"] = trusthost5
        __props__.__dict__["trusthost6"] = trusthost6
        __props__.__dict__["trusthost7"] = trusthost7
        __props__.__dict__["trusthost8"] = trusthost8
        __props__.__dict__["trusthost9"] = trusthost9
        __props__.__dict__["vdoms"] = vdoms
        return SystemAdminAdministrator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accprofile(self) -> pulumi.Output[str]:
        """
        Access profile for this administrator. Access profiles control administrator access to FortiGate features.
        """
        return pulumi.get(self, "accprofile")

    @property
    @pulumi.getter
    def comments(self) -> pulumi.Output[Optional[str]]:
        """
        Comment.
        """
        return pulumi.get(self, "comments")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Admin user password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def trusthost1(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost1")

    @property
    @pulumi.getter
    def trusthost10(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost10")

    @property
    @pulumi.getter
    def trusthost2(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost2")

    @property
    @pulumi.getter
    def trusthost3(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost3")

    @property
    @pulumi.getter
    def trusthost4(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost4")

    @property
    @pulumi.getter
    def trusthost5(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost5")

    @property
    @pulumi.getter
    def trusthost6(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost6")

    @property
    @pulumi.getter
    def trusthost7(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost7")

    @property
    @pulumi.getter
    def trusthost8(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost8")

    @property
    @pulumi.getter
    def trusthost9(self) -> pulumi.Output[str]:
        return pulumi.get(self, "trusthost9")

    @property
    @pulumi.getter
    def vdoms(self) -> pulumi.Output[Sequence[str]]:
        """
        Virtual domain(s) that the administrator can access.
        """
        return pulumi.get(self, "vdoms")

