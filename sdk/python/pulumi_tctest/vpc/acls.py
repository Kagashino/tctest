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
    'AclsResult',
    'AwaitableAclsResult',
    'acls',
    'acls_output',
]

@pulumi.output_type
class AclsResult:
    """
    A collection of values returned by Acls.
    """
    def __init__(__self__, acl_lists=None, id=None, name=None, result_output_file=None, vpc_id=None):
        if acl_lists and not isinstance(acl_lists, list):
            raise TypeError("Expected argument 'acl_lists' to be a list")
        pulumi.set(__self__, "acl_lists", acl_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="aclLists")
    def acl_lists(self) -> Sequence['outputs.AclsAclListResult']:
        return pulumi.get(self, "acl_lists")

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

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableAclsResult(AclsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AclsResult(
            acl_lists=self.acl_lists,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            vpc_id=self.vpc_id)


def acls(id: Optional[str] = None,
         name: Optional[str] = None,
         result_output_file: Optional[str] = None,
         vpc_id: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAclsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Vpc/acls:Acls', __args__, opts=opts, typ=AclsResult).value

    return AwaitableAclsResult(
        acl_lists=__ret__.acl_lists,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(acls)
def acls_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                name: Optional[pulumi.Input[Optional[str]]] = None,
                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AclsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
