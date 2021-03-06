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
    'GroupRuleAddressTemplate',
    'GroupRuleProtocolTemplate',
    'GroupsSecurityGroupResult',
]

@pulumi.output_type
class GroupRuleAddressTemplate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "groupId":
            suggest = "group_id"
        elif key == "templateId":
            suggest = "template_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GroupRuleAddressTemplate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GroupRuleAddressTemplate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GroupRuleAddressTemplate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 group_id: Optional[str] = None,
                 template_id: Optional[str] = None):
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[str]:
        return pulumi.get(self, "template_id")


@pulumi.output_type
class GroupRuleProtocolTemplate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "groupId":
            suggest = "group_id"
        elif key == "templateId":
            suggest = "template_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GroupRuleProtocolTemplate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GroupRuleProtocolTemplate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GroupRuleProtocolTemplate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 group_id: Optional[str] = None,
                 template_id: Optional[str] = None):
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if template_id is not None:
            pulumi.set(__self__, "template_id", template_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[str]:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> Optional[str]:
        return pulumi.get(self, "template_id")


@pulumi.output_type
class GroupsSecurityGroupResult(dict):
    def __init__(__self__, *,
                 be_associate_count: int,
                 create_time: str,
                 description: str,
                 egresses: Sequence[str],
                 ingresses: Sequence[str],
                 name: str,
                 project_id: int,
                 security_group_id: str,
                 tags: Mapping[str, Any]):
        pulumi.set(__self__, "be_associate_count", be_associate_count)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "egresses", egresses)
        pulumi.set(__self__, "ingresses", ingresses)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="beAssociateCount")
    def be_associate_count(self) -> int:
        return pulumi.get(self, "be_associate_count")

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
    def egresses(self) -> Sequence[str]:
        return pulumi.get(self, "egresses")

    @property
    @pulumi.getter
    def ingresses(self) -> Sequence[str]:
        return pulumi.get(self, "ingresses")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        return pulumi.get(self, "tags")


