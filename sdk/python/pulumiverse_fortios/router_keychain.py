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

__all__ = ['RouterKeychainArgs', 'RouterKeychain']

@pulumi.input_type
class RouterKeychainArgs:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouterKeychain resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]] keys: Configuration method to edit key settings. The structure of `key` block is documented below.
        :param pulumi.Input[str] name: Key-chain name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]]]:
        """
        Configuration method to edit key settings. The structure of `key` block is documented below.
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Key-chain name.
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
class _RouterKeychainState:
    def __init__(__self__, *,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouterKeychain resources.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]] keys: Configuration method to edit key settings. The structure of `key` block is documented below.
        :param pulumi.Input[str] name: Key-chain name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        if dynamic_sort_subtable is not None:
            pulumi.set(__self__, "dynamic_sort_subtable", dynamic_sort_subtable)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vdomparam is not None:
            pulumi.set(__self__, "vdomparam", vdomparam)

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
    def keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]]]:
        """
        Configuration method to edit key settings. The structure of `key` block is documented below.
        """
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RouterKeychainKeyArgs']]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Key-chain name.
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


class RouterKeychain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterKeychainKeyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure key-chain.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.RouterKeychain("trname", keys=[fortios.RouterKeychainKeyArgs(
            accept_lifetime="04:00:00 01 01 2008 04:00:00 01 01 2022",
            key_string="ewiwn3i23232s212",
            send_lifetime="04:00:00 01 01 2008 04:00:00 01 01 2022",
        )])
        ```

        ## Import

        Router KeyChain can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/routerKeychain:RouterKeychain labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/routerKeychain:RouterKeychain labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterKeychainKeyArgs']]]] keys: Configuration method to edit key settings. The structure of `key` block is documented below.
        :param pulumi.Input[str] name: Key-chain name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RouterKeychainArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure key-chain.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_fortios as fortios

        trname = fortios.RouterKeychain("trname", keys=[fortios.RouterKeychainKeyArgs(
            accept_lifetime="04:00:00 01 01 2008 04:00:00 01 01 2022",
            key_string="ewiwn3i23232s212",
            send_lifetime="04:00:00 01 01 2008 04:00:00 01 01 2022",
        )])
        ```

        ## Import

        Router KeyChain can be imported using any of these accepted formats

        ```sh
         $ pulumi import fortios:index/routerKeychain:RouterKeychain labelname {{name}}
        ```

         If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"

        ```sh
         $ pulumi import fortios:index/routerKeychain:RouterKeychain labelname {{name}}
        ```

         $ unset "FORTIOS_IMPORT_TABLE"

        :param str resource_name: The name of the resource.
        :param RouterKeychainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouterKeychainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterKeychainKeyArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vdomparam: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouterKeychainArgs.__new__(RouterKeychainArgs)

            __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
            __props__.__dict__["keys"] = keys
            __props__.__dict__["name"] = name
            __props__.__dict__["vdomparam"] = vdomparam
        super(RouterKeychain, __self__).__init__(
            'fortios:index/routerKeychain:RouterKeychain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dynamic_sort_subtable: Optional[pulumi.Input[str]] = None,
            keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterKeychainKeyArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vdomparam: Optional[pulumi.Input[str]] = None) -> 'RouterKeychain':
        """
        Get an existing RouterKeychain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dynamic_sort_subtable: Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouterKeychainKeyArgs']]]] keys: Configuration method to edit key settings. The structure of `key` block is documented below.
        :param pulumi.Input[str] name: Key-chain name.
        :param pulumi.Input[str] vdomparam: Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouterKeychainState.__new__(_RouterKeychainState)

        __props__.__dict__["dynamic_sort_subtable"] = dynamic_sort_subtable
        __props__.__dict__["keys"] = keys
        __props__.__dict__["name"] = name
        __props__.__dict__["vdomparam"] = vdomparam
        return RouterKeychain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dynamicSortSubtable")
    def dynamic_sort_subtable(self) -> pulumi.Output[Optional[str]]:
        """
        Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
        """
        return pulumi.get(self, "dynamic_sort_subtable")

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Output[Optional[Sequence['outputs.RouterKeychainKey']]]:
        """
        Configuration method to edit key settings. The structure of `key` block is documented below.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Key-chain name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def vdomparam(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        """
        return pulumi.get(self, "vdomparam")

