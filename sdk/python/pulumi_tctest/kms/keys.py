# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'KeysResult',
    'AwaitableKeysResult',
    'keys',
    'keys_output',
]

@pulumi.output_type
class KeysResult:
    """
    A collection of values returned by Keys.
    """
    def __init__(__self__, id=None, key_lists=None, key_state=None, key_usage=None, order_type=None, origin=None, result_output_file=None, role=None, search_key_alias=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_lists and not isinstance(key_lists, list):
            raise TypeError("Expected argument 'key_lists' to be a list")
        pulumi.set(__self__, "key_lists", key_lists)
        if key_state and not isinstance(key_state, int):
            raise TypeError("Expected argument 'key_state' to be a int")
        pulumi.set(__self__, "key_state", key_state)
        if key_usage and not isinstance(key_usage, str):
            raise TypeError("Expected argument 'key_usage' to be a str")
        pulumi.set(__self__, "key_usage", key_usage)
        if order_type and not isinstance(order_type, int):
            raise TypeError("Expected argument 'order_type' to be a int")
        pulumi.set(__self__, "order_type", order_type)
        if origin and not isinstance(origin, str):
            raise TypeError("Expected argument 'origin' to be a str")
        pulumi.set(__self__, "origin", origin)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if role and not isinstance(role, int):
            raise TypeError("Expected argument 'role' to be a int")
        pulumi.set(__self__, "role", role)
        if search_key_alias and not isinstance(search_key_alias, str):
            raise TypeError("Expected argument 'search_key_alias' to be a str")
        pulumi.set(__self__, "search_key_alias", search_key_alias)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyLists")
    def key_lists(self) -> Sequence['outputs.KeysKeyListResult']:
        return pulumi.get(self, "key_lists")

    @property
    @pulumi.getter(name="keyState")
    def key_state(self) -> Optional[int]:
        return pulumi.get(self, "key_state")

    @property
    @pulumi.getter(name="keyUsage")
    def key_usage(self) -> Optional[str]:
        return pulumi.get(self, "key_usage")

    @property
    @pulumi.getter(name="orderType")
    def order_type(self) -> Optional[int]:
        return pulumi.get(self, "order_type")

    @property
    @pulumi.getter
    def origin(self) -> Optional[str]:
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter
    def role(self) -> Optional[int]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="searchKeyAlias")
    def search_key_alias(self) -> Optional[str]:
        return pulumi.get(self, "search_key_alias")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


class AwaitableKeysResult(KeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return KeysResult(
            id=self.id,
            key_lists=self.key_lists,
            key_state=self.key_state,
            key_usage=self.key_usage,
            order_type=self.order_type,
            origin=self.origin,
            result_output_file=self.result_output_file,
            role=self.role,
            search_key_alias=self.search_key_alias,
            tags=self.tags)


def keys(key_state: Optional[int] = None,
         key_usage: Optional[str] = None,
         order_type: Optional[int] = None,
         origin: Optional[str] = None,
         result_output_file: Optional[str] = None,
         role: Optional[int] = None,
         search_key_alias: Optional[str] = None,
         tags: Optional[Mapping[str, Any]] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableKeysResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['keyState'] = key_state
    __args__['keyUsage'] = key_usage
    __args__['orderType'] = order_type
    __args__['origin'] = origin
    __args__['resultOutputFile'] = result_output_file
    __args__['role'] = role
    __args__['searchKeyAlias'] = search_key_alias
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Kms/keys:Keys', __args__, opts=opts, typ=KeysResult).value

    return AwaitableKeysResult(
        id=__ret__.id,
        key_lists=__ret__.key_lists,
        key_state=__ret__.key_state,
        key_usage=__ret__.key_usage,
        order_type=__ret__.order_type,
        origin=__ret__.origin,
        result_output_file=__ret__.result_output_file,
        role=__ret__.role,
        search_key_alias=__ret__.search_key_alias,
        tags=__ret__.tags)


@_utilities.lift_output_func(keys)
def keys_output(key_state: Optional[pulumi.Input[Optional[int]]] = None,
                key_usage: Optional[pulumi.Input[Optional[str]]] = None,
                order_type: Optional[pulumi.Input[Optional[int]]] = None,
                origin: Optional[pulumi.Input[Optional[str]]] = None,
                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                role: Optional[pulumi.Input[Optional[int]]] = None,
                search_key_alias: Optional[pulumi.Input[Optional[str]]] = None,
                tags: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[KeysResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
