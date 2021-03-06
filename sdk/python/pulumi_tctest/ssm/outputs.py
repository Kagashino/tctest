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
    'SecretVersionsSecretVersionListResult',
    'SecretsSecretListResult',
]

@pulumi.output_type
class SecretVersionsSecretVersionListResult(dict):
    def __init__(__self__, *,
                 secret_binary: str,
                 secret_string: str,
                 version_id: str):
        pulumi.set(__self__, "secret_binary", secret_binary)
        pulumi.set(__self__, "secret_string", secret_string)
        pulumi.set(__self__, "version_id", version_id)

    @property
    @pulumi.getter(name="secretBinary")
    def secret_binary(self) -> str:
        return pulumi.get(self, "secret_binary")

    @property
    @pulumi.getter(name="secretString")
    def secret_string(self) -> str:
        return pulumi.get(self, "secret_string")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> str:
        return pulumi.get(self, "version_id")


@pulumi.output_type
class SecretsSecretListResult(dict):
    def __init__(__self__, *,
                 create_time: int,
                 create_uin: int,
                 delete_time: int,
                 description: str,
                 kms_key_id: str,
                 secret_name: str,
                 status: str):
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "create_uin", create_uin)
        pulumi.set(__self__, "delete_time", delete_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        pulumi.set(__self__, "secret_name", secret_name)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> int:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="createUin")
    def create_uin(self) -> int:
        return pulumi.get(self, "create_uin")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> int:
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> str:
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


