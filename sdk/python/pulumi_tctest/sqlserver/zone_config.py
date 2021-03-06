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
    'ZoneConfigResult',
    'AwaitableZoneConfigResult',
    'zone_config',
    'zone_config_output',
]

@pulumi.output_type
class ZoneConfigResult:
    """
    A collection of values returned by ZoneConfig.
    """
    def __init__(__self__, id=None, result_output_file=None, zone_lists=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if zone_lists and not isinstance(zone_lists, list):
            raise TypeError("Expected argument 'zone_lists' to be a list")
        pulumi.set(__self__, "zone_lists", zone_lists)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="zoneLists")
    def zone_lists(self) -> Sequence['outputs.ZoneConfigZoneListResult']:
        return pulumi.get(self, "zone_lists")


class AwaitableZoneConfigResult(ZoneConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ZoneConfigResult(
            id=self.id,
            result_output_file=self.result_output_file,
            zone_lists=self.zone_lists)


def zone_config(result_output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableZoneConfigResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Sqlserver/zoneConfig:ZoneConfig', __args__, opts=opts, typ=ZoneConfigResult).value

    return AwaitableZoneConfigResult(
        id=__ret__.id,
        result_output_file=__ret__.result_output_file,
        zone_lists=__ret__.zone_lists)


@_utilities.lift_output_func(zone_config)
def zone_config_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ZoneConfigResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
