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
    'NamespacesResult',
    'AwaitableNamespacesResult',
    'namespaces',
    'namespaces_output',
]

@pulumi.output_type
class NamespacesResult:
    """
    A collection of values returned by Namespaces.
    """
    def __init__(__self__, description=None, id=None, namespace=None, namespaces=None, result_output_file=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if namespaces and not isinstance(namespaces, list):
            raise TypeError("Expected argument 'namespaces' to be a list")
        pulumi.set(__self__, "namespaces", namespaces)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def namespaces(self) -> Sequence['outputs.NamespacesNamespaceResult']:
        return pulumi.get(self, "namespaces")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")


class AwaitableNamespacesResult(NamespacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return NamespacesResult(
            description=self.description,
            id=self.id,
            namespace=self.namespace,
            namespaces=self.namespaces,
            result_output_file=self.result_output_file)


def namespaces(description: Optional[str] = None,
               namespace: Optional[str] = None,
               result_output_file: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableNamespacesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['namespace'] = namespace
    __args__['resultOutputFile'] = result_output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Scf/namespaces:Namespaces', __args__, opts=opts, typ=NamespacesResult).value

    return AwaitableNamespacesResult(
        description=__ret__.description,
        id=__ret__.id,
        namespace=__ret__.namespace,
        namespaces=__ret__.namespaces,
        result_output_file=__ret__.result_output_file)


@_utilities.lift_output_func(namespaces)
def namespaces_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                      namespace: Optional[pulumi.Input[Optional[str]]] = None,
                      result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[NamespacesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
