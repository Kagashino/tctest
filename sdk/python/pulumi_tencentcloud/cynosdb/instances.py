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
    def __init__(__self__, cluster_id=None, db_type=None, id=None, instance_id=None, instance_lists=None, instance_name=None, project_id=None, result_output_file=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if db_type and not isinstance(db_type, str):
            raise TypeError("Expected argument 'db_type' to be a str")
        pulumi.set(__self__, "db_type", db_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_lists and not isinstance(instance_lists, list):
            raise TypeError("Expected argument 'instance_lists' to be a list")
        pulumi.set(__self__, "instance_lists", instance_lists)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="dbType")
    def db_type(self) -> Optional[str]:
        return pulumi.get(self, "db_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceLists")
    def instance_lists(self) -> Sequence['outputs.InstancesInstanceListResult']:
        return pulumi.get(self, "instance_lists")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[str]:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[int]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            cluster_id=self.cluster_id,
            db_type=self.db_type,
            id=self.id,
            instance_id=self.instance_id,
            instance_lists=self.instance_lists,
            instance_name=self.instance_name,
            project_id=self.project_id,
            result_output_file=self.result_output_file)


def instances(cluster_id: Optional[str] = None,
              db_type: Optional[str] = None,
              instance_id: Optional[str] = None,
              instance_name: Optional[str] = None,
              project_id: Optional[int] = None,
              result_output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['dbType'] = db_type
    __args__['instanceId'] = instance_id
    __args__['instanceName'] = instance_name
    __args__['projectId'] = project_id
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Cynosdb/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        cluster_id=__ret__.cluster_id,
        db_type=__ret__.db_type,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_lists=__ret__.instance_lists,
        instance_name=__ret__.instance_name,
        project_id=__ret__.project_id,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(instances)
def instances_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                     db_type: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_name: Optional[pulumi.Input[Optional[str]]] = None,
                     project_id: Optional[pulumi.Input[Optional[int]]] = None,
                     result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...