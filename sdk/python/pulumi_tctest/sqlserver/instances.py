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
    'InstancesResult',
    'AwaitableInstancesResult',
    'instances',
    'instances_output',
]

@pulumi.output_type
class InstancesResult:
    """
    A collection of values returned by Instances.
    """
    def __init__(__self__, id=None, instance_lists=None, name=None, project_id=None, result_output_file=None, subnet_id=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.InstancesInstanceListResult']:
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            id=self.id,
            instance_lists=self.instance_lists,
            name=self.name,
            project_id=self.project_id,
            result_output_file=self.result_output_file,
            subnet_id=self.subnet_id,
            vpc_id=self.vpc_id)


def instances(id: Optional[str] = None,
              name: Optional[str] = None,
              project_id: Optional[int] = None,
              result_output_file: Optional[str] = None,
              subnet_id: Optional[str] = None,
              vpc_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Sqlserver/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        id=__ret__.id,
        instance_lists=__ret__.instance_lists,
        name=__ret__.name,
        project_id=__ret__.project_id,
        result_output_file=__ret__.result_output_file,
        subnet_id=__ret__.subnet_id,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(instances)
def instances_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     project_id: Optional[pulumi.Input[Optional[int]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                     vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
