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
    'BandwidthLimitsLimitResult',
    'InstancesInstanceListResult',
    'InstancesInstanceListAttachmentListResult',
]

@pulumi.output_type
class BandwidthLimitsLimitResult(dict):
    def __init__(__self__, *,
                 bandwidth_limit: int,
                 dst_region: str,
                 region: str):
        pulumi.set(__self__, "bandwidth_limit", bandwidth_limit)
        pulumi.set(__self__, "dst_region", dst_region)
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="bandwidthLimit")
    def bandwidth_limit(self) -> int:
        return pulumi.get(self, "bandwidth_limit")

    @property
    @pulumi.getter(name="dstRegion")
    def dst_region(self) -> str:
        return pulumi.get(self, "dst_region")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")


@pulumi.output_type
class InstancesInstanceListResult(dict):
    def __init__(__self__, *,
                 attachment_lists: Sequence['outputs.InstancesInstanceListAttachmentListResult'],
                 bandwidth_limit_type: str,
                 ccn_id: str,
                 charge_type: str,
                 create_time: str,
                 description: str,
                 name: str,
                 qos: str,
                 state: str):
        pulumi.set(__self__, "attachment_lists", attachment_lists)
        pulumi.set(__self__, "bandwidth_limit_type", bandwidth_limit_type)
        pulumi.set(__self__, "ccn_id", ccn_id)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "qos", qos)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="attachmentLists")
    def attachment_lists(self) -> Sequence['outputs.InstancesInstanceListAttachmentListResult']:
        return pulumi.get(self, "attachment_lists")

    @property
    @pulumi.getter(name="bandwidthLimitType")
    def bandwidth_limit_type(self) -> str:
        return pulumi.get(self, "bandwidth_limit_type")

    @property
    @pulumi.getter(name="ccnId")
    def ccn_id(self) -> str:
        return pulumi.get(self, "ccn_id")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def qos(self) -> str:
        return pulumi.get(self, "qos")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")


@pulumi.output_type
class InstancesInstanceListAttachmentListResult(dict):
    def __init__(__self__, *,
                 attached_time: str,
                 cidr_blocks: Sequence[str],
                 instance_id: str,
                 instance_region: str,
                 instance_type: str,
                 state: str):
        pulumi.set(__self__, "attached_time", attached_time)
        pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region", instance_region)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="attachedTime")
    def attached_time(self) -> str:
        return pulumi.get(self, "attached_time")

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Sequence[str]:
        return pulumi.get(self, "cidr_blocks")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceRegion")
    def instance_region(self) -> str:
        return pulumi.get(self, "instance_region")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")


