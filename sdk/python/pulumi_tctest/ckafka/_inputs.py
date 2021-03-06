# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceConfigArgs',
    'InstanceDynamicRetentionConfigArgs',
    'InstanceTagArgs',
]

@pulumi.input_type
class InstanceConfigArgs:
    def __init__(__self__, *,
                 auto_create_topic_enable: pulumi.Input[bool],
                 default_num_partitions: pulumi.Input[int],
                 default_replication_factor: pulumi.Input[int]):
        pulumi.set(__self__, "auto_create_topic_enable", auto_create_topic_enable)
        pulumi.set(__self__, "default_num_partitions", default_num_partitions)
        pulumi.set(__self__, "default_replication_factor", default_replication_factor)

    @property
    @pulumi.getter(name="autoCreateTopicEnable")
    def auto_create_topic_enable(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "auto_create_topic_enable")

    @auto_create_topic_enable.setter
    def auto_create_topic_enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "auto_create_topic_enable", value)

    @property
    @pulumi.getter(name="defaultNumPartitions")
    def default_num_partitions(self) -> pulumi.Input[int]:
        return pulumi.get(self, "default_num_partitions")

    @default_num_partitions.setter
    def default_num_partitions(self, value: pulumi.Input[int]):
        pulumi.set(self, "default_num_partitions", value)

    @property
    @pulumi.getter(name="defaultReplicationFactor")
    def default_replication_factor(self) -> pulumi.Input[int]:
        return pulumi.get(self, "default_replication_factor")

    @default_replication_factor.setter
    def default_replication_factor(self, value: pulumi.Input[int]):
        pulumi.set(self, "default_replication_factor", value)


@pulumi.input_type
class InstanceDynamicRetentionConfigArgs:
    def __init__(__self__, *,
                 bottom_retention: Optional[pulumi.Input[int]] = None,
                 disk_quota_percentage: Optional[pulumi.Input[int]] = None,
                 enable: Optional[pulumi.Input[int]] = None,
                 step_forward_percentage: Optional[pulumi.Input[int]] = None):
        if bottom_retention is not None:
            pulumi.set(__self__, "bottom_retention", bottom_retention)
        if disk_quota_percentage is not None:
            pulumi.set(__self__, "disk_quota_percentage", disk_quota_percentage)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if step_forward_percentage is not None:
            pulumi.set(__self__, "step_forward_percentage", step_forward_percentage)

    @property
    @pulumi.getter(name="bottomRetention")
    def bottom_retention(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "bottom_retention")

    @bottom_retention.setter
    def bottom_retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bottom_retention", value)

    @property
    @pulumi.getter(name="diskQuotaPercentage")
    def disk_quota_percentage(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disk_quota_percentage")

    @disk_quota_percentage.setter
    def disk_quota_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_quota_percentage", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="stepForwardPercentage")
    def step_forward_percentage(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "step_forward_percentage")

    @step_forward_percentage.setter
    def step_forward_percentage(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "step_forward_percentage", value)


@pulumi.input_type
class InstanceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


