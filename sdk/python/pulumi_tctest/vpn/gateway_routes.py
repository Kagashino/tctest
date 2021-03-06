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
    'GatewayRoutesResult',
    'AwaitableGatewayRoutesResult',
    'gateway_routes',
    'gateway_routes_output',
]

@pulumi.output_type
class GatewayRoutesResult:
    """
    A collection of values returned by GatewayRoutes.
    """
    def __init__(__self__, destination_cidr=None, id=None, instance_id=None, instance_type=None, result_output_file=None, vpn_gateway_id=None, vpn_gateway_route_lists=None):
        if destination_cidr and not isinstance(destination_cidr, str):
            raise TypeError("Expected argument 'destination_cidr' to be a str")
        pulumi.set(__self__, "destination_cidr", destination_cidr)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if result_output_file and not isinstance(result_output_file, str):
            raise TypeError("Expected argument 'result_output_file' to be a str")
        pulumi.set(__self__, "result_output_file", result_output_file)
        if vpn_gateway_id and not isinstance(vpn_gateway_id, str):
            raise TypeError("Expected argument 'vpn_gateway_id' to be a str")
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)
        if vpn_gateway_route_lists and not isinstance(vpn_gateway_route_lists, list):
            raise TypeError("Expected argument 'vpn_gateway_route_lists' to be a list")
        pulumi.set(__self__, "vpn_gateway_route_lists", vpn_gateway_route_lists)

    @property
    @pulumi.getter(name="destinationCidr")
    def destination_cidr(self) -> Optional[str]:
        return pulumi.get(self, "destination_cidr")

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
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="resultOutputFile")
    def result_output_file(self) -> Optional[str]:
        return pulumi.get(self, "result_output_file")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> str:
        return pulumi.get(self, "vpn_gateway_id")

    @property
    @pulumi.getter(name="vpnGatewayRouteLists")
    def vpn_gateway_route_lists(self) -> Sequence['outputs.GatewayRoutesVpnGatewayRouteListResult']:
        return pulumi.get(self, "vpn_gateway_route_lists")


class AwaitableGatewayRoutesResult(GatewayRoutesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GatewayRoutesResult(
            destination_cidr=self.destination_cidr,
            id=self.id,
            instance_id=self.instance_id,
            instance_type=self.instance_type,
            result_output_file=self.result_output_file,
            vpn_gateway_id=self.vpn_gateway_id,
            vpn_gateway_route_lists=self.vpn_gateway_route_lists)


def gateway_routes(destination_cidr: Optional[str] = None,
                   instance_id: Optional[str] = None,
                   instance_type: Optional[str] = None,
                   result_output_file: Optional[str] = None,
                   vpn_gateway_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGatewayRoutesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['destinationCidr'] = destination_cidr
    __args__['instanceId'] = instance_id
    __args__['instanceType'] = instance_type
    __args__['resultOutputFile'] = result_output_file
    __args__['vpnGatewayId'] = vpn_gateway_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('tctest:Vpn/gatewayRoutes:GatewayRoutes', __args__, opts=opts, typ=GatewayRoutesResult).value

    return AwaitableGatewayRoutesResult(
        destination_cidr=__ret__.destination_cidr,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_type=__ret__.instance_type,
        result_output_file=__ret__.result_output_file,
        vpn_gateway_id=__ret__.vpn_gateway_id,
        vpn_gateway_route_lists=__ret__.vpn_gateway_route_lists)


@_utilities.lift_output_func(gateway_routes)
def gateway_routes_output(destination_cidr: Optional[pulumi.Input[Optional[str]]] = None,
                          instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                          instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                          result_output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GatewayRoutesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
