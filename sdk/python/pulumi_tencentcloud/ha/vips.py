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
    'VipsResult',
    'AwaitableVipsResult',
    'vips',
    'vips_output',
]

@pulumi.output_type
class VipsResult:
    """
    A collection of values returned by Vips.
    """
    def __init__(__self__, address_ip=None, ha_vip_lists=None, id=None, name=None, result_output_file=None, subnet_id=None, vpc_id=None):
        if address_ip and not isinstance(address_ip, str):
            raise TypeError("Expected argument 'address_ip' to be a str")
        pulumi.set(__self__, "address_ip", address_ip)
        if ha_vip_lists and not isinstance(ha_vip_lists, list):
            raise TypeError("Expected argument 'ha_vip_lists' to be a list")
        pulumi.set(__self__, "ha_vip_lists", ha_vip_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="addressIp")
    def address_ip(self) -> Optional[str]:
        return pulumi.get(self, "address_ip")

    @property
    @pulumi.getter(name="haVipLists")
    def ha_vip_lists(self) -> Sequence['outputs.VipsHaVipListResult']:
        return pulumi.get(self, "ha_vip_lists")

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
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")


class AwaitableVipsResult(VipsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return VipsResult(
            address_ip=self.address_ip,
            ha_vip_lists=self.ha_vip_lists,
            id=self.id,
            name=self.name,
            result_output_file=self.result_output_file,
            subnet_id=self.subnet_id,
            vpc_id=self.vpc_id)


def vips(address_ip: Optional[str] = None,
         id: Optional[str] = None,
         name: Optional[str] = None,
         result_output_file: Optional[str] = None,
         subnet_id: Optional[str] = None,
         vpc_id: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableVipsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressIp'] = address_ip
    __args__['id'] = id
    __args__['name'] = name
    __args__['resultOutputFile'] = result_output_file
    __args__['subnetId'] = subnet_id
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tencentcloud:Ha/vips:Vips', __args__, opts=opts, typ=VipsResult).value

    return AwaitableVipsResult(
        address_ip=__ret__.address_ip,
        ha_vip_lists=__ret__.ha_vip_lists,
        id=__ret__.id,
        name=__ret__.name,
        result_output_file=__ret__.result_output_file,
        subnet_id=__ret__.subnet_id,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(vips)
def vips_output(address_ip: Optional[pulumi.Input[Optional[str]]] = None,
                id: Optional[pulumi.Input[Optional[str]]] = None,
                name: Optional[pulumi.Input[Optional[str]]] = None,
                result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[VipsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...