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
    'SuperPlayerConfigsResult',
    'AwaitableSuperPlayerConfigsResult',
    'super_player_configs',
    'super_player_configs_output',
]

@pulumi.output_type
class SuperPlayerConfigsResult:
    """
    A collection of values returned by SuperPlayerConfigs.
    """
    def __init__(__self__, config_lists=None, id=None, name=None, result_output_file=None, sub_app_id=None, type=None):
        if config_lists and not isinstance(config_lists, list):
            raise TypeError("Expected argument 'config_lists' to be a list")
        pulumi.set(__self__, "config_lists", config_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if sub_app_id and not isinstance(sub_app_id, int):
            raise TypeError("Expected argument 'sub_app_id' to be a int")
        pulumi.set(__self__, "sub_app_id", sub_app_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="configLists")
    def config_lists(self) -> Sequence['outputs.SuperPlayerConfigsConfigListResult']:
        return pulumi.get(self, "config_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="subAppId")
    def sub_app_id(self) -> Optional[int]:
        return pulumi.get(self, "sub_app_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


class AwaitableSuperPlayerConfigsResult(SuperPlayerConfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SuperPlayerConfigsResult(
            config_lists=self.config_lists,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            sub_app_id=self.sub_app_id,
            type=self.type)


def super_player_configs(name: Optional[str] = None,
                         result_output_file: Optional[str] = None,
                         sub_app_id: Optional[int] = None,
                         type: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSuperPlayerConfigsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['subAppId'] = sub_app_id
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Vod/superPlayerConfigs:SuperPlayerConfigs', __args__, opts=opts, typ=SuperPlayerConfigsResult).value

    return AwaitableSuperPlayerConfigsResult(
        config_lists=__ret__.config_lists,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        sub_app_id=__ret__.sub_app_id,
        type=__ret__.type)


@_utilities.lift_output_func(super_player_configs)
def super_player_configs_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                sub_app_id: Optional[pulumi.Input[Optional[int]]] = None,
                                type: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SuperPlayerConfigsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
