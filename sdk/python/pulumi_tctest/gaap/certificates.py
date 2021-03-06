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
    'CertificatesResult',
    'AwaitableCertificatesResult',
    'certificates',
    'certificates_output',
]

@pulumi.output_type
class CertificatesResult:
    """
    A collection of values returned by Certificates.
    """
    def __init__(__self__, certificates=None, id=None, name=None, result_output_file=None, type=None):
        if certificates and not isinstance(certificates, list):
            raise TypeError("Expected argument 'certificates' to be a list")
        pulumi.set(__self__, "certificates", certificates)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def certificates(self) -> Sequence['outputs.CertificatesCertificateResult']:
        return pulumi.get(self, "certificates")

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
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


class AwaitableCertificatesResult(CertificatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CertificatesResult(
            certificates=self.certificates,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            type=self.type)


def certificates(id: Optional[str] = None,
                 name: Optional[str] = None,
                 result_output_file: Optional[str] = None,
                 type: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCertificatesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Gaap/certificates:Certificates', __args__, opts=opts, typ=CertificatesResult).value

    return AwaitableCertificatesResult(
        certificates=__ret__.certificates,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        type=__ret__.type)


@_utilities.lift_output_func(certificates)
def certificates_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[Optional[str]]] = None,
                        result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        type: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[CertificatesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
