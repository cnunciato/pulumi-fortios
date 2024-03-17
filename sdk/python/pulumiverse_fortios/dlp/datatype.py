# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DatatypeArgs', 'Datatype']

@pulumi.input_type
class DatatypeArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 look_ahead: Optional[pulumi.Input[int]] = None,
                 look_back: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 transform: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 verify: Optional[pulumi.Input[str]] = None,
                 verify_transformed_pattern: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Datatype resource.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[int] look_ahead: Number of characters to obtain in advance for verification (1 - 255, default = 1).
        :param pulumi.Input[int] look_back: Number of characters required to save for verification (1 - 255, default = 1).
        :param pulumi.Input[str] name: Name of table containing the data type.
        :param pulumi.Input[str] pattern: Regular expression pattern string without look around.
        :param pulumi.Input[str] transform: Template to transform user input to a pattern using capture group from 'pattern'.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] verify: Regular expression pattern string used to verify the data type.
        :param pulumi.Input[str] verify_transformed_pattern: Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if look_ahead is not None:
            pulumi.set(__self__, "look_ahead", look_ahead)
        if look_back is not None:
            pulumi.set(__self__, "look_back", look_back)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if transform is not None:
            pulumi.set(__self__, "transform", transform)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if verify is not None:
            pulumi.set(__self__, "verify", verify)
        if verify_transformed_pattern is not None:
            pulumi.set(__self__, "verify_transformed_pattern", verify_transformed_pattern)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Optional comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="lookAhead")
    def look_ahead(self) -> Optional[pulumi.Input[int]]:
        """
        Number of characters to obtain in advance for verification (1 - 255, default = 1).
        """
        return pulumi.get(self, "look_ahead")

    @look_ahead.setter
    def look_ahead(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "look_ahead", value)

    @property
    @pulumi.getter(name="lookBack")
    def look_back(self) -> Optional[pulumi.Input[int]]:
        """
        Number of characters required to save for verification (1 - 255, default = 1).
        """
        return pulumi.get(self, "look_back")

    @look_back.setter
    def look_back(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "look_back", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of table containing the data type.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Regular expression pattern string without look around.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def transform(self) -> Optional[pulumi.Input[str]]:
        """
        Template to transform user input to a pattern using capture group from 'pattern'.
        """
        return pulumi.get(self, "transform")

    @transform.setter
    def transform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transform", value)

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
    def verify(self) -> Optional[pulumi.Input[str]]:
        """
        Regular expression pattern string used to verify the data type.
        """
        return pulumi.get(self, "verify")

    @verify.setter
    def verify(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verify", value)

    @property
    @pulumi.getter(name="verifyTransformedPattern")
    def verify_transformed_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "verify_transformed_pattern")

    @verify_transformed_pattern.setter
    def verify_transformed_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verify_transformed_pattern", value)


@pulumi.input_type
class _DatatypeState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 look_ahead: Optional[pulumi.Input[int]] = None,
                 look_back: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 transform: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 verify: Optional[pulumi.Input[str]] = None,
                 verify_transformed_pattern: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Datatype resources.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[int] look_ahead: Number of characters to obtain in advance for verification (1 - 255, default = 1).
        :param pulumi.Input[int] look_back: Number of characters required to save for verification (1 - 255, default = 1).
        :param pulumi.Input[str] name: Name of table containing the data type.
        :param pulumi.Input[str] pattern: Regular expression pattern string without look around.
        :param pulumi.Input[str] transform: Template to transform user input to a pattern using capture group from 'pattern'.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] verify: Regular expression pattern string used to verify the data type.
        :param pulumi.Input[str] verify_transformed_pattern: Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if look_ahead is not None:
            pulumi.set(__self__, "look_ahead", look_ahead)
        if look_back is not None:
            pulumi.set(__self__, "look_back", look_back)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if transform is not None:
            pulumi.set(__self__, "transform", transform)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)
        if verify is not None:
            pulumi.set(__self__, "verify", verify)
        if verify_transformed_pattern is not None:
            pulumi.set(__self__, "verify_transformed_pattern", verify_transformed_pattern)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Optional comments.
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="lookAhead")
    def look_ahead(self) -> Optional[pulumi.Input[int]]:
        """
        Number of characters to obtain in advance for verification (1 - 255, default = 1).
        """
        return pulumi.get(self, "look_ahead")

    @look_ahead.setter
    def look_ahead(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "look_ahead", value)

    @property
    @pulumi.getter(name="lookBack")
    def look_back(self) -> Optional[pulumi.Input[int]]:
        """
        Number of characters required to save for verification (1 - 255, default = 1).
        """
        return pulumi.get(self, "look_back")

    @look_back.setter
    def look_back(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "look_back", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of table containing the data type.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Regular expression pattern string without look around.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def transform(self) -> Optional[pulumi.Input[str]]:
        """
        Template to transform user input to a pattern using capture group from 'pattern'.
        """
        return pulumi.get(self, "transform")

    @transform.setter
    def transform(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transform", value)

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
    def verify(self) -> Optional[pulumi.Input[str]]:
        """
        Regular expression pattern string used to verify the data type.
        """
        return pulumi.get(self, "verify")

    @verify.setter
    def verify(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verify", value)

    @property
    @pulumi.getter(name="verifyTransformedPattern")
    def verify_transformed_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "verify_transformed_pattern")

    @verify_transformed_pattern.setter
    def verify_transformed_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "verify_transformed_pattern", value)


class Datatype(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 look_ahead: Optional[pulumi.Input[int]] = None,
                 look_back: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 transform: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 verify: Optional[pulumi.Input[str]] = None,
                 verify_transformed_pattern: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure predefined data type used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.

        ## Import

        Dlp DataType can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[int] look_ahead: Number of characters to obtain in advance for verification (1 - 255, default = 1).
        :param pulumi.Input[int] look_back: Number of characters required to save for verification (1 - 255, default = 1).
        :param pulumi.Input[str] name: Name of table containing the data type.
        :param pulumi.Input[str] pattern: Regular expression pattern string without look around.
        :param pulumi.Input[str] transform: Template to transform user input to a pattern using capture group from 'pattern'.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] verify: Regular expression pattern string used to verify the data type.
        :param pulumi.Input[str] verify_transformed_pattern: Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DatatypeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure predefined data type used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.

        ## Import

        Dlp DataType can be imported using any of these accepted formats:

        ```sh
        $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
        ```

        If you do not want to import arguments of block:

        $ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
        $ pulumi import fortios:dlp/datatype:Datatype labelname {{name}}
        ```

        $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param DatatypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatatypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 look_ahead: Optional[pulumi.Input[int]] = None,
                 look_back: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 transform: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 verify: Optional[pulumi.Input[str]] = None,
                 verify_transformed_pattern: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatatypeArgs.__new__(DatatypeArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["look_ahead"] = look_ahead
            __props__.__dict__["look_back"] = look_back
            __props__.__dict__["name"] = name
            __props__.__dict__["pattern"] = pattern
            __props__.__dict__["transform"] = transform
            __props__.__dict__["vdomparam"] = vdomparam
            __props__.__dict__["verify"] = verify
            __props__.__dict__["verify_transformed_pattern"] = verify_transformed_pattern
        super(Datatype, __self__).__init__(
            'fortios:dlp/datatype:Datatype',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            look_ahead: Optional[pulumi.Input[int]] = None,
            look_back: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            transform: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None,
            verify: Optional[pulumi.Input[str]] = None,
            verify_transformed_pattern: Optional[pulumi.Input[str]] = None) -> 'Datatype':
        """
        Get an existing Datatype resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Optional comments.
        :param pulumi.Input[int] look_ahead: Number of characters to obtain in advance for verification (1 - 255, default = 1).
        :param pulumi.Input[int] look_back: Number of characters required to save for verification (1 - 255, default = 1).
        :param pulumi.Input[str] name: Name of table containing the data type.
        :param pulumi.Input[str] pattern: Regular expression pattern string without look around.
        :param pulumi.Input[str] transform: Template to transform user input to a pattern using capture group from 'pattern'.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        :param pulumi.Input[str] verify: Regular expression pattern string used to verify the data type.
        :param pulumi.Input[str] verify_transformed_pattern: Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatatypeState.__new__(_DatatypeState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["look_ahead"] = look_ahead
        __props__.__dict__["look_back"] = look_back
        __props__.__dict__["name"] = name
        __props__.__dict__["pattern"] = pattern
        __props__.__dict__["transform"] = transform
        __props__.__dict__["vdomparam"] = vdomparam
        __props__.__dict__["verify"] = verify
        __props__.__dict__["verify_transformed_pattern"] = verify_transformed_pattern
        return Datatype(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Optional comments.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="lookAhead")
    def look_ahead(self) -> pulumi.Output[int]:
        """
        Number of characters to obtain in advance for verification (1 - 255, default = 1).
        """
        return pulumi.get(self, "look_ahead")

    @property
    @pulumi.getter(name="lookBack")
    def look_back(self) -> pulumi.Output[int]:
        """
        Number of characters required to save for verification (1 - 255, default = 1).
        """
        return pulumi.get(self, "look_back")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of table containing the data type.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[str]:
        """
        Regular expression pattern string without look around.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def transform(self) -> pulumi.Output[str]:
        """
        Template to transform user input to a pattern using capture group from 'pattern'.
        """
        return pulumi.get(self, "transform")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

    @property
    @pulumi.getter
    def verify(self) -> pulumi.Output[str]:
        """
        Regular expression pattern string used to verify the data type.
        """
        return pulumi.get(self, "verify")

    @property
    @pulumi.getter(name="verifyTransformedPattern")
    def verify_transformed_pattern(self) -> pulumi.Output[str]:
        """
        Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
        """
        return pulumi.get(self, "verify_transformed_pattern")
