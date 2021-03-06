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
    'GroupMembershipsResult',
    'AwaitableGroupMembershipsResult',
    'group_memberships',
    'group_memberships_output',
]

@pulumi.output_type
class GroupMembershipsResult:
    """
    A collection of values returned by GroupMemberships.
    """
    def __init__(__self__, group_id=None, id=None, membership_lists=None, result_output_file=None):
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if membership_lists and not isinstance(membership_lists, list):
            raise TypeError("Expected argument 'membership_lists' to be a list")
        pulumi.set(__self__, "membership_lists", membership_lists)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="membershipLists")
    def membership_lists(self) -> Sequence['outputs.GroupMembershipsMembershipListResult']:
        return pulumi.get(self, "membership_lists")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableGroupMembershipsResult(GroupMembershipsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GroupMembershipsResult(
            group_id=self.group_id,
            id=self.id,
            membership_lists=self.membership_lists,
            result_output_file=self.result_output_file)


def group_memberships(group_id: Optional[str] = None,
                      result_output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGroupMembershipsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['groupId'] = group_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Cam/groupMemberships:GroupMemberships', __args__, opts=opts, typ=GroupMembershipsResult).value

    return AwaitableGroupMembershipsResult(
        group_id=__ret__.group_id,
        id=__ret__.id,
        membership_lists=__ret__.membership_lists,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(group_memberships)
def group_memberships_output(group_id: Optional[pulumi.Input[Optional[str]]] = None,
                             result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GroupMembershipsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
