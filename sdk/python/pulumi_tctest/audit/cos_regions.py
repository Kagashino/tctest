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
    'CosRegionsResult',
    'AwaitableCosRegionsResult',
    'cos_regions',
    'cos_regions_output',
]

@pulumi.output_type
class CosRegionsResult:
    """
    A collection of values returned by CosRegions.
    """
    def __init__(__self__, audit_cos_region_lists=None, id=None, result_output_file=None):
        if audit_cos_region_lists and not isinstance(audit_cos_region_lists, list):
            raise TypeError("Expected argument 'audit_cos_region_lists' to be a list")
        pulumi.set(__self__, "audit_cos_region_lists", audit_cos_region_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="auditCosRegionLists")
    def audit_cos_region_lists(self) -> Sequence['outputs.CosRegionsAuditCosRegionListResult']:
        return pulumi.get(self, "audit_cos_region_lists")

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


class AwaitableCosRegionsResult(CosRegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CosRegionsResult(
            audit_cos_region_lists=self.audit_cos_region_lists,
            id=self.id,
            result_output_file=self.result_output_file)


def cos_regions(result_output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCosRegionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Audit/cosRegions:CosRegions', __args__, opts=opts, typ=CosRegionsResult).value

    return AwaitableCosRegionsResult(
        audit_cos_region_lists=__ret__.audit_cos_region_lists,
        id=__ret__.id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(cos_regions)
def cos_regions_output(result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[CosRegionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
