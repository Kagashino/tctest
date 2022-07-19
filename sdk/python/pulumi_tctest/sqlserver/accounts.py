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
    'AccountsResult',
    'AwaitableAccountsResult',
    'accounts',
    'accounts_output',
]

@pulumi.output_type
class AccountsResult:
    """
    A collection of values returned by Accounts.
    """
    def __init__(__self__, id=None, instance_id=None, lists=None, name=None, result_output_file=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if lists and not isinstance(lists, list):
            raise TypeError("Expected argument 'lists' to be a list")
        pulumi.set(__self__, "lists", lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def lists(self) -> Sequence['outputs.AccountsListResult']:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableAccountsResult(AccountsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AccountsResult(
            id=self.id,
            instance_id=self.instance_id,
            lists=self.lists,
            name=self.name,
            result_output_file=self.result_output_file)


def accounts(instance_id: Optional[str] = None,
             name: Optional[str] = None,
             result_output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAccountsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Sqlserver/accounts:Accounts', __args__, opts=opts, typ=AccountsResult).value

    return AwaitableAccountsResult(
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        lists=__ret__.lists,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(accounts)
def accounts_output(instance_id: Optional[pulumi.Input[str]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AccountsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...