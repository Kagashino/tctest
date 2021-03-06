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
    'TemplateGroupsResult',
    'AwaitableTemplateGroupsResult',
    'template_groups',
    'template_groups_output',
]

@pulumi.output_type
class TemplateGroupsResult:
    """
    A collection of values returned by TemplateGroups.
    """
    def __init__(__self__, group_lists=None, id=None, name=None, result_output_file=None):
        if group_lists and not isinstance(group_lists, list):
            raise TypeError("Expected argument 'group_lists' to be a list")
        pulumi.set(__self__, "group_lists", group_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="groupLists")
    def group_lists(self) -> Sequence['outputs.TemplateGroupsGroupListResult']:
        return pulumi.get(self, "group_lists")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableTemplateGroupsResult(TemplateGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TemplateGroupsResult(
            group_lists=self.group_lists,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file)


def template_groups(id: Optional[str] = None,
                    name: Optional[str] = None,
                    result_output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTemplateGroupsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Address/templateGroups:TemplateGroups', __args__, opts=opts, typ=TemplateGroupsResult).value

    return AwaitableTemplateGroupsResult(
        group_lists=__ret__.group_lists,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(template_groups)
def template_groups_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[TemplateGroupsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
