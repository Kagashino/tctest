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
    'ConfigExcludePathArgs',
    'ConfigExtraContainerFileArgs',
    'ConfigExtraContainerFileWorkloadArgs',
    'ConfigExtraContainerStdoutArgs',
    'ConfigExtraContainerStdoutWorkloadArgs',
    'ConfigExtraExcludePathArgs',
    'ConfigExtraExtractRuleArgs',
    'ConfigExtraExtractRuleFilterKeyRegexArgs',
    'ConfigExtraHostFileArgs',
    'ConfigExtractRuleArgs',
    'ConfigExtractRuleFilterKeyRegexArgs',
    'CosShipperCompressArgs',
    'CosShipperContentArgs',
    'CosShipperContentCsvArgs',
    'CosShipperContentJsonArgs',
    'CosShipperFilterRuleArgs',
    'MachineGroupMachineGroupTypeArgs',
]

@pulumi.input_type
class ConfigExcludePathArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigExtraContainerFileArgs:
    def __init__(__self__, *,
                 container: pulumi.Input[str],
                 file_pattern: pulumi.Input[str],
                 log_path: pulumi.Input[str],
                 namespace: pulumi.Input[str],
                 exclude_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_namespace: Optional[pulumi.Input[str]] = None,
                 include_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 workload: Optional[pulumi.Input['ConfigExtraContainerFileWorkloadArgs']] = None):
        pulumi.set(__self__, "container", container)
        pulumi.set(__self__, "file_pattern", file_pattern)
        pulumi.set(__self__, "log_path", log_path)
        pulumi.set(__self__, "namespace", namespace)
        if exclude_labels is not None:
            pulumi.set(__self__, "exclude_labels", exclude_labels)
        if exclude_namespace is not None:
            pulumi.set(__self__, "exclude_namespace", exclude_namespace)
        if include_labels is not None:
            pulumi.set(__self__, "include_labels", include_labels)
        if workload is not None:
            pulumi.set(__self__, "workload", workload)

    @property
    @pulumi.getter
    def container(self) -> pulumi.Input[str]:
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: pulumi.Input[str]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter(name="filePattern")
    def file_pattern(self) -> pulumi.Input[str]:
        return pulumi.get(self, "file_pattern")

    @file_pattern.setter
    def file_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_pattern", value)

    @property
    @pulumi.getter(name="logPath")
    def log_path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "log_path")

    @log_path.setter
    def log_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_path", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="excludeLabels")
    def exclude_labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "exclude_labels")

    @exclude_labels.setter
    def exclude_labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_labels", value)

    @property
    @pulumi.getter(name="excludeNamespace")
    def exclude_namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exclude_namespace")

    @exclude_namespace.setter
    def exclude_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude_namespace", value)

    @property
    @pulumi.getter(name="includeLabels")
    def include_labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "include_labels")

    @include_labels.setter
    def include_labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "include_labels", value)

    @property
    @pulumi.getter
    def workload(self) -> Optional[pulumi.Input['ConfigExtraContainerFileWorkloadArgs']]:
        return pulumi.get(self, "workload")

    @workload.setter
    def workload(self, value: Optional[pulumi.Input['ConfigExtraContainerFileWorkloadArgs']]):
        pulumi.set(self, "workload", value)


@pulumi.input_type
class ConfigExtraContainerFileWorkloadArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 name: pulumi.Input[str],
                 container: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)
        if container is not None:
            pulumi.set(__self__, "container", container)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def container(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class ConfigExtraContainerStdoutArgs:
    def __init__(__self__, *,
                 all_containers: pulumi.Input[bool],
                 exclude_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 exclude_namespace: Optional[pulumi.Input[str]] = None,
                 include_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 workloads: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtraContainerStdoutWorkloadArgs']]]] = None):
        pulumi.set(__self__, "all_containers", all_containers)
        if exclude_labels is not None:
            pulumi.set(__self__, "exclude_labels", exclude_labels)
        if exclude_namespace is not None:
            pulumi.set(__self__, "exclude_namespace", exclude_namespace)
        if include_labels is not None:
            pulumi.set(__self__, "include_labels", include_labels)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if workloads is not None:
            pulumi.set(__self__, "workloads", workloads)

    @property
    @pulumi.getter(name="allContainers")
    def all_containers(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "all_containers")

    @all_containers.setter
    def all_containers(self, value: pulumi.Input[bool]):
        pulumi.set(self, "all_containers", value)

    @property
    @pulumi.getter(name="excludeLabels")
    def exclude_labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "exclude_labels")

    @exclude_labels.setter
    def exclude_labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "exclude_labels", value)

    @property
    @pulumi.getter(name="excludeNamespace")
    def exclude_namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "exclude_namespace")

    @exclude_namespace.setter
    def exclude_namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "exclude_namespace", value)

    @property
    @pulumi.getter(name="includeLabels")
    def include_labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "include_labels")

    @include_labels.setter
    def include_labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "include_labels", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def workloads(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtraContainerStdoutWorkloadArgs']]]]:
        return pulumi.get(self, "workloads")

    @workloads.setter
    def workloads(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtraContainerStdoutWorkloadArgs']]]]):
        pulumi.set(self, "workloads", value)


@pulumi.input_type
class ConfigExtraContainerStdoutWorkloadArgs:
    def __init__(__self__, *,
                 kind: pulumi.Input[str],
                 name: pulumi.Input[str],
                 container: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "name", name)
        if container is not None:
            pulumi.set(__self__, "container", container)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def container(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container")

    @container.setter
    def container(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class ConfigExtraExcludePathArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConfigExtraExtractRuleArgs:
    def __init__(__self__, *,
                 backtracking: Optional[pulumi.Input[int]] = None,
                 begin_regex: Optional[pulumi.Input[str]] = None,
                 delimiter: Optional[pulumi.Input[str]] = None,
                 filter_key_regexes: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtraExtractRuleFilterKeyRegexArgs']]]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_regex: Optional[pulumi.Input[str]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 time_key: Optional[pulumi.Input[str]] = None,
                 un_match_log_key: Optional[pulumi.Input[str]] = None,
                 un_match_up_load_switch: Optional[pulumi.Input[bool]] = None):
        if backtracking is not None:
            pulumi.set(__self__, "backtracking", backtracking)
        if begin_regex is not None:
            pulumi.set(__self__, "begin_regex", begin_regex)
        if delimiter is not None:
            pulumi.set(__self__, "delimiter", delimiter)
        if filter_key_regexes is not None:
            pulumi.set(__self__, "filter_key_regexes", filter_key_regexes)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if log_regex is not None:
            pulumi.set(__self__, "log_regex", log_regex)
        if time_format is not None:
            pulumi.set(__self__, "time_format", time_format)
        if time_key is not None:
            pulumi.set(__self__, "time_key", time_key)
        if un_match_log_key is not None:
            pulumi.set(__self__, "un_match_log_key", un_match_log_key)
        if un_match_up_load_switch is not None:
            pulumi.set(__self__, "un_match_up_load_switch", un_match_up_load_switch)

    @property
    @pulumi.getter
    def backtracking(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "backtracking")

    @backtracking.setter
    def backtracking(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "backtracking", value)

    @property
    @pulumi.getter(name="beginRegex")
    def begin_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "begin_regex")

    @begin_regex.setter
    def begin_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "begin_regex", value)

    @property
    @pulumi.getter
    def delimiter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "delimiter")

    @delimiter.setter
    def delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delimiter", value)

    @property
    @pulumi.getter(name="filterKeyRegexes")
    def filter_key_regexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtraExtractRuleFilterKeyRegexArgs']]]]:
        return pulumi.get(self, "filter_key_regexes")

    @filter_key_regexes.setter
    def filter_key_regexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtraExtractRuleFilterKeyRegexArgs']]]]):
        pulumi.set(self, "filter_key_regexes", value)

    @property
    @pulumi.getter
    def keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter(name="logRegex")
    def log_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_regex")

    @log_regex.setter
    def log_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_regex", value)

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "time_format")

    @time_format.setter
    def time_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_format", value)

    @property
    @pulumi.getter(name="timeKey")
    def time_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "time_key")

    @time_key.setter
    def time_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_key", value)

    @property
    @pulumi.getter(name="unMatchLogKey")
    def un_match_log_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "un_match_log_key")

    @un_match_log_key.setter
    def un_match_log_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "un_match_log_key", value)

    @property
    @pulumi.getter(name="unMatchUpLoadSwitch")
    def un_match_up_load_switch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "un_match_up_load_switch")

    @un_match_up_load_switch.setter
    def un_match_up_load_switch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "un_match_up_load_switch", value)


@pulumi.input_type
class ConfigExtraExtractRuleFilterKeyRegexArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 regex: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class ConfigExtraHostFileArgs:
    def __init__(__self__, *,
                 file_pattern: pulumi.Input[str],
                 log_path: pulumi.Input[str],
                 custom_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "file_pattern", file_pattern)
        pulumi.set(__self__, "log_path", log_path)
        if custom_labels is not None:
            pulumi.set(__self__, "custom_labels", custom_labels)

    @property
    @pulumi.getter(name="filePattern")
    def file_pattern(self) -> pulumi.Input[str]:
        return pulumi.get(self, "file_pattern")

    @file_pattern.setter
    def file_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_pattern", value)

    @property
    @pulumi.getter(name="logPath")
    def log_path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "log_path")

    @log_path.setter
    def log_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_path", value)

    @property
    @pulumi.getter(name="customLabels")
    def custom_labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_labels")

    @custom_labels.setter
    def custom_labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_labels", value)


@pulumi.input_type
class ConfigExtractRuleArgs:
    def __init__(__self__, *,
                 backtracking: Optional[pulumi.Input[int]] = None,
                 begin_regex: Optional[pulumi.Input[str]] = None,
                 delimiter: Optional[pulumi.Input[str]] = None,
                 filter_key_regexes: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtractRuleFilterKeyRegexArgs']]]] = None,
                 keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 log_regex: Optional[pulumi.Input[str]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 time_key: Optional[pulumi.Input[str]] = None,
                 un_match_log_key: Optional[pulumi.Input[str]] = None,
                 un_match_up_load_switch: Optional[pulumi.Input[bool]] = None):
        if backtracking is not None:
            pulumi.set(__self__, "backtracking", backtracking)
        if begin_regex is not None:
            pulumi.set(__self__, "begin_regex", begin_regex)
        if delimiter is not None:
            pulumi.set(__self__, "delimiter", delimiter)
        if filter_key_regexes is not None:
            pulumi.set(__self__, "filter_key_regexes", filter_key_regexes)
        if keys is not None:
            pulumi.set(__self__, "keys", keys)
        if log_regex is not None:
            pulumi.set(__self__, "log_regex", log_regex)
        if time_format is not None:
            pulumi.set(__self__, "time_format", time_format)
        if time_key is not None:
            pulumi.set(__self__, "time_key", time_key)
        if un_match_log_key is not None:
            pulumi.set(__self__, "un_match_log_key", un_match_log_key)
        if un_match_up_load_switch is not None:
            pulumi.set(__self__, "un_match_up_load_switch", un_match_up_load_switch)

    @property
    @pulumi.getter
    def backtracking(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "backtracking")

    @backtracking.setter
    def backtracking(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "backtracking", value)

    @property
    @pulumi.getter(name="beginRegex")
    def begin_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "begin_regex")

    @begin_regex.setter
    def begin_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "begin_regex", value)

    @property
    @pulumi.getter
    def delimiter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "delimiter")

    @delimiter.setter
    def delimiter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delimiter", value)

    @property
    @pulumi.getter(name="filterKeyRegexes")
    def filter_key_regexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtractRuleFilterKeyRegexArgs']]]]:
        return pulumi.get(self, "filter_key_regexes")

    @filter_key_regexes.setter
    def filter_key_regexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConfigExtractRuleFilterKeyRegexArgs']]]]):
        pulumi.set(self, "filter_key_regexes", value)

    @property
    @pulumi.getter
    def keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter(name="logRegex")
    def log_regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_regex")

    @log_regex.setter
    def log_regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_regex", value)

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "time_format")

    @time_format.setter
    def time_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_format", value)

    @property
    @pulumi.getter(name="timeKey")
    def time_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "time_key")

    @time_key.setter
    def time_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_key", value)

    @property
    @pulumi.getter(name="unMatchLogKey")
    def un_match_log_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "un_match_log_key")

    @un_match_log_key.setter
    def un_match_log_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "un_match_log_key", value)

    @property
    @pulumi.getter(name="unMatchUpLoadSwitch")
    def un_match_up_load_switch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "un_match_up_load_switch")

    @un_match_up_load_switch.setter
    def un_match_up_load_switch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "un_match_up_load_switch", value)


@pulumi.input_type
class ConfigExtractRuleFilterKeyRegexArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 regex: Optional[pulumi.Input[str]] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if regex is not None:
            pulumi.set(__self__, "regex", regex)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def regex(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "regex", value)


@pulumi.input_type
class CosShipperCompressArgs:
    def __init__(__self__, *,
                 format: pulumi.Input[str]):
        pulumi.set(__self__, "format", format)

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: pulumi.Input[str]):
        pulumi.set(self, "format", value)


@pulumi.input_type
class CosShipperContentArgs:
    def __init__(__self__, *,
                 format: pulumi.Input[str],
                 csv: Optional[pulumi.Input['CosShipperContentCsvArgs']] = None,
                 json: Optional[pulumi.Input['CosShipperContentJsonArgs']] = None):
        pulumi.set(__self__, "format", format)
        if csv is not None:
            pulumi.set(__self__, "csv", csv)
        if json is not None:
            pulumi.set(__self__, "json", json)

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: pulumi.Input[str]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def csv(self) -> Optional[pulumi.Input['CosShipperContentCsvArgs']]:
        return pulumi.get(self, "csv")

    @csv.setter
    def csv(self, value: Optional[pulumi.Input['CosShipperContentCsvArgs']]):
        pulumi.set(self, "csv", value)

    @property
    @pulumi.getter
    def json(self) -> Optional[pulumi.Input['CosShipperContentJsonArgs']]:
        return pulumi.get(self, "json")

    @json.setter
    def json(self, value: Optional[pulumi.Input['CosShipperContentJsonArgs']]):
        pulumi.set(self, "json", value)


@pulumi.input_type
class CosShipperContentCsvArgs:
    def __init__(__self__, *,
                 delimiter: pulumi.Input[str],
                 escape_char: pulumi.Input[str],
                 keys: pulumi.Input[Sequence[pulumi.Input[str]]],
                 non_existing_field: pulumi.Input[str],
                 print_key: pulumi.Input[bool]):
        pulumi.set(__self__, "delimiter", delimiter)
        pulumi.set(__self__, "escape_char", escape_char)
        pulumi.set(__self__, "keys", keys)
        pulumi.set(__self__, "non_existing_field", non_existing_field)
        pulumi.set(__self__, "print_key", print_key)

    @property
    @pulumi.getter
    def delimiter(self) -> pulumi.Input[str]:
        return pulumi.get(self, "delimiter")

    @delimiter.setter
    def delimiter(self, value: pulumi.Input[str]):
        pulumi.set(self, "delimiter", value)

    @property
    @pulumi.getter(name="escapeChar")
    def escape_char(self) -> pulumi.Input[str]:
        return pulumi.get(self, "escape_char")

    @escape_char.setter
    def escape_char(self, value: pulumi.Input[str]):
        pulumi.set(self, "escape_char", value)

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "keys")

    @keys.setter
    def keys(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "keys", value)

    @property
    @pulumi.getter(name="nonExistingField")
    def non_existing_field(self) -> pulumi.Input[str]:
        return pulumi.get(self, "non_existing_field")

    @non_existing_field.setter
    def non_existing_field(self, value: pulumi.Input[str]):
        pulumi.set(self, "non_existing_field", value)

    @property
    @pulumi.getter(name="printKey")
    def print_key(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "print_key")

    @print_key.setter
    def print_key(self, value: pulumi.Input[bool]):
        pulumi.set(self, "print_key", value)


@pulumi.input_type
class CosShipperContentJsonArgs:
    def __init__(__self__, *,
                 enable_tag: pulumi.Input[bool],
                 meta_fields: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(__self__, "enable_tag", enable_tag)
        pulumi.set(__self__, "meta_fields", meta_fields)

    @property
    @pulumi.getter(name="enableTag")
    def enable_tag(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "enable_tag")

    @enable_tag.setter
    def enable_tag(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable_tag", value)

    @property
    @pulumi.getter(name="metaFields")
    def meta_fields(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "meta_fields")

    @meta_fields.setter
    def meta_fields(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "meta_fields", value)


@pulumi.input_type
class CosShipperFilterRuleArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 regex: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "regex", regex)
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
    def regex(self) -> pulumi.Input[str]:
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: pulumi.Input[str]):
        pulumi.set(self, "regex", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class MachineGroupMachineGroupTypeArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 values: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)


