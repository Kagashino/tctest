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
    'CdnDomainHttpsConfig',
    'CdnDomainHttpsConfigClientCertificateConfig',
    'CdnDomainHttpsConfigForceRedirect',
    'CdnDomainHttpsConfigServerCertificateConfig',
    'CdnDomainOrigin',
    'CdnDomainRequestHeader',
    'CdnDomainRequestHeaderHeaderRule',
    'CdnDomainRuleCach',
    'DomainsDomainListResult',
    'DomainsDomainListHttpsConfigResult',
    'DomainsDomainListOriginResult',
    'DomainsDomainListRequestHeaderResult',
    'DomainsDomainListRequestHeaderHeaderRuleResult',
    'DomainsDomainListRuleCachResult',
]

@pulumi.output_type
class CdnDomainHttpsConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "httpsSwitch":
            suggest = "https_switch"
        elif key == "clientCertificateConfig":
            suggest = "client_certificate_config"
        elif key == "forceRedirect":
            suggest = "force_redirect"
        elif key == "http2Switch":
            suggest = "http2_switch"
        elif key == "ocspStaplingSwitch":
            suggest = "ocsp_stapling_switch"
        elif key == "serverCertificateConfig":
            suggest = "server_certificate_config"
        elif key == "spdySwitch":
            suggest = "spdy_switch"
        elif key == "verifyClient":
            suggest = "verify_client"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainHttpsConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainHttpsConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainHttpsConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 https_switch: str,
                 client_certificate_config: Optional['outputs.CdnDomainHttpsConfigClientCertificateConfig'] = None,
                 force_redirect: Optional['outputs.CdnDomainHttpsConfigForceRedirect'] = None,
                 http2_switch: Optional[str] = None,
                 ocsp_stapling_switch: Optional[str] = None,
                 server_certificate_config: Optional['outputs.CdnDomainHttpsConfigServerCertificateConfig'] = None,
                 spdy_switch: Optional[str] = None,
                 verify_client: Optional[str] = None):
        pulumi.set(__self__, "https_switch", https_switch)
        if client_certificate_config is not None:
            pulumi.set(__self__, "client_certificate_config", client_certificate_config)
        if force_redirect is not None:
            pulumi.set(__self__, "force_redirect", force_redirect)
        if http2_switch is not None:
            pulumi.set(__self__, "http2_switch", http2_switch)
        if ocsp_stapling_switch is not None:
            pulumi.set(__self__, "ocsp_stapling_switch", ocsp_stapling_switch)
        if server_certificate_config is not None:
            pulumi.set(__self__, "server_certificate_config", server_certificate_config)
        if spdy_switch is not None:
            pulumi.set(__self__, "spdy_switch", spdy_switch)
        if verify_client is not None:
            pulumi.set(__self__, "verify_client", verify_client)

    @property
    @pulumi.getter(name="httpsSwitch")
    def https_switch(self) -> str:
        return pulumi.get(self, "https_switch")

    @property
    @pulumi.getter(name="clientCertificateConfig")
    def client_certificate_config(self) -> Optional['outputs.CdnDomainHttpsConfigClientCertificateConfig']:
        return pulumi.get(self, "client_certificate_config")

    @property
    @pulumi.getter(name="forceRedirect")
    def force_redirect(self) -> Optional['outputs.CdnDomainHttpsConfigForceRedirect']:
        return pulumi.get(self, "force_redirect")

    @property
    @pulumi.getter(name="http2Switch")
    def http2_switch(self) -> Optional[str]:
        return pulumi.get(self, "http2_switch")

    @property
    @pulumi.getter(name="ocspStaplingSwitch")
    def ocsp_stapling_switch(self) -> Optional[str]:
        return pulumi.get(self, "ocsp_stapling_switch")

    @property
    @pulumi.getter(name="serverCertificateConfig")
    def server_certificate_config(self) -> Optional['outputs.CdnDomainHttpsConfigServerCertificateConfig']:
        return pulumi.get(self, "server_certificate_config")

    @property
    @pulumi.getter(name="spdySwitch")
    def spdy_switch(self) -> Optional[str]:
        return pulumi.get(self, "spdy_switch")

    @property
    @pulumi.getter(name="verifyClient")
    def verify_client(self) -> Optional[str]:
        return pulumi.get(self, "verify_client")


@pulumi.output_type
class CdnDomainHttpsConfigClientCertificateConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateContent":
            suggest = "certificate_content"
        elif key == "certificateName":
            suggest = "certificate_name"
        elif key == "deployTime":
            suggest = "deploy_time"
        elif key == "expireTime":
            suggest = "expire_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainHttpsConfigClientCertificateConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainHttpsConfigClientCertificateConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainHttpsConfigClientCertificateConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_content: str,
                 certificate_name: Optional[str] = None,
                 deploy_time: Optional[str] = None,
                 expire_time: Optional[str] = None):
        pulumi.set(__self__, "certificate_content", certificate_content)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if deploy_time is not None:
            pulumi.set(__self__, "deploy_time", deploy_time)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)

    @property
    @pulumi.getter(name="certificateContent")
    def certificate_content(self) -> str:
        return pulumi.get(self, "certificate_content")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[str]:
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter(name="deployTime")
    def deploy_time(self) -> Optional[str]:
        return pulumi.get(self, "deploy_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[str]:
        return pulumi.get(self, "expire_time")


@pulumi.output_type
class CdnDomainHttpsConfigForceRedirect(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "redirectStatusCode":
            suggest = "redirect_status_code"
        elif key == "redirectType":
            suggest = "redirect_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainHttpsConfigForceRedirect. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainHttpsConfigForceRedirect.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainHttpsConfigForceRedirect.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 redirect_status_code: Optional[int] = None,
                 redirect_type: Optional[str] = None,
                 switch: Optional[str] = None):
        if redirect_status_code is not None:
            pulumi.set(__self__, "redirect_status_code", redirect_status_code)
        if redirect_type is not None:
            pulumi.set(__self__, "redirect_type", redirect_type)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="redirectStatusCode")
    def redirect_status_code(self) -> Optional[int]:
        return pulumi.get(self, "redirect_status_code")

    @property
    @pulumi.getter(name="redirectType")
    def redirect_type(self) -> Optional[str]:
        return pulumi.get(self, "redirect_type")

    @property
    @pulumi.getter
    def switch(self) -> Optional[str]:
        return pulumi.get(self, "switch")


@pulumi.output_type
class CdnDomainHttpsConfigServerCertificateConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "certificateContent":
            suggest = "certificate_content"
        elif key == "certificateId":
            suggest = "certificate_id"
        elif key == "certificateName":
            suggest = "certificate_name"
        elif key == "deployTime":
            suggest = "deploy_time"
        elif key == "expireTime":
            suggest = "expire_time"
        elif key == "privateKey":
            suggest = "private_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainHttpsConfigServerCertificateConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainHttpsConfigServerCertificateConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainHttpsConfigServerCertificateConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 certificate_content: Optional[str] = None,
                 certificate_id: Optional[str] = None,
                 certificate_name: Optional[str] = None,
                 deploy_time: Optional[str] = None,
                 expire_time: Optional[str] = None,
                 message: Optional[str] = None,
                 private_key: Optional[str] = None):
        if certificate_content is not None:
            pulumi.set(__self__, "certificate_content", certificate_content)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if certificate_name is not None:
            pulumi.set(__self__, "certificate_name", certificate_name)
        if deploy_time is not None:
            pulumi.set(__self__, "deploy_time", deploy_time)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if message is not None:
            pulumi.set(__self__, "message", message)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)

    @property
    @pulumi.getter(name="certificateContent")
    def certificate_content(self) -> Optional[str]:
        return pulumi.get(self, "certificate_content")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[str]:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[str]:
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter(name="deployTime")
    def deploy_time(self) -> Optional[str]:
        return pulumi.get(self, "deploy_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[str]:
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        return pulumi.get(self, "message")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[str]:
        return pulumi.get(self, "private_key")


@pulumi.output_type
class CdnDomainOrigin(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "originLists":
            suggest = "origin_lists"
        elif key == "originType":
            suggest = "origin_type"
        elif key == "backupOriginLists":
            suggest = "backup_origin_lists"
        elif key == "backupOriginType":
            suggest = "backup_origin_type"
        elif key == "backupServerName":
            suggest = "backup_server_name"
        elif key == "cosPrivateAccess":
            suggest = "cos_private_access"
        elif key == "originPullProtocol":
            suggest = "origin_pull_protocol"
        elif key == "serverName":
            suggest = "server_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainOrigin. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainOrigin.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainOrigin.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 origin_lists: Sequence[str],
                 origin_type: str,
                 backup_origin_lists: Optional[Sequence[str]] = None,
                 backup_origin_type: Optional[str] = None,
                 backup_server_name: Optional[str] = None,
                 cos_private_access: Optional[str] = None,
                 origin_pull_protocol: Optional[str] = None,
                 server_name: Optional[str] = None):
        pulumi.set(__self__, "origin_lists", origin_lists)
        pulumi.set(__self__, "origin_type", origin_type)
        if backup_origin_lists is not None:
            pulumi.set(__self__, "backup_origin_lists", backup_origin_lists)
        if backup_origin_type is not None:
            pulumi.set(__self__, "backup_origin_type", backup_origin_type)
        if backup_server_name is not None:
            pulumi.set(__self__, "backup_server_name", backup_server_name)
        if cos_private_access is not None:
            pulumi.set(__self__, "cos_private_access", cos_private_access)
        if origin_pull_protocol is not None:
            pulumi.set(__self__, "origin_pull_protocol", origin_pull_protocol)
        if server_name is not None:
            pulumi.set(__self__, "server_name", server_name)

    @property
    @pulumi.getter(name="originLists")
    def origin_lists(self) -> Sequence[str]:
        return pulumi.get(self, "origin_lists")

    @property
    @pulumi.getter(name="originType")
    def origin_type(self) -> str:
        return pulumi.get(self, "origin_type")

    @property
    @pulumi.getter(name="backupOriginLists")
    def backup_origin_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "backup_origin_lists")

    @property
    @pulumi.getter(name="backupOriginType")
    def backup_origin_type(self) -> Optional[str]:
        return pulumi.get(self, "backup_origin_type")

    @property
    @pulumi.getter(name="backupServerName")
    def backup_server_name(self) -> Optional[str]:
        return pulumi.get(self, "backup_server_name")

    @property
    @pulumi.getter(name="cosPrivateAccess")
    def cos_private_access(self) -> Optional[str]:
        return pulumi.get(self, "cos_private_access")

    @property
    @pulumi.getter(name="originPullProtocol")
    def origin_pull_protocol(self) -> Optional[str]:
        return pulumi.get(self, "origin_pull_protocol")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> Optional[str]:
        return pulumi.get(self, "server_name")


@pulumi.output_type
class CdnDomainRequestHeader(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "headerRules":
            suggest = "header_rules"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainRequestHeader. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainRequestHeader.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainRequestHeader.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 header_rules: Optional[Sequence['outputs.CdnDomainRequestHeaderHeaderRule']] = None,
                 switch: Optional[str] = None):
        if header_rules is not None:
            pulumi.set(__self__, "header_rules", header_rules)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="headerRules")
    def header_rules(self) -> Optional[Sequence['outputs.CdnDomainRequestHeaderHeaderRule']]:
        return pulumi.get(self, "header_rules")

    @property
    @pulumi.getter
    def switch(self) -> Optional[str]:
        return pulumi.get(self, "switch")


@pulumi.output_type
class CdnDomainRequestHeaderHeaderRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "headerMode":
            suggest = "header_mode"
        elif key == "headerName":
            suggest = "header_name"
        elif key == "headerValue":
            suggest = "header_value"
        elif key == "rulePaths":
            suggest = "rule_paths"
        elif key == "ruleType":
            suggest = "rule_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainRequestHeaderHeaderRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainRequestHeaderHeaderRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainRequestHeaderHeaderRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 header_mode: str,
                 header_name: str,
                 header_value: str,
                 rule_paths: Sequence[str],
                 rule_type: str):
        pulumi.set(__self__, "header_mode", header_mode)
        pulumi.set(__self__, "header_name", header_name)
        pulumi.set(__self__, "header_value", header_value)
        pulumi.set(__self__, "rule_paths", rule_paths)
        pulumi.set(__self__, "rule_type", rule_type)

    @property
    @pulumi.getter(name="headerMode")
    def header_mode(self) -> str:
        return pulumi.get(self, "header_mode")

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> str:
        return pulumi.get(self, "header_name")

    @property
    @pulumi.getter(name="headerValue")
    def header_value(self) -> str:
        return pulumi.get(self, "header_value")

    @property
    @pulumi.getter(name="rulePaths")
    def rule_paths(self) -> Sequence[str]:
        return pulumi.get(self, "rule_paths")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> str:
        return pulumi.get(self, "rule_type")


@pulumi.output_type
class CdnDomainRuleCach(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "cacheTime":
            suggest = "cache_time"
        elif key == "compareMaxAge":
            suggest = "compare_max_age"
        elif key == "followOriginSwitch":
            suggest = "follow_origin_switch"
        elif key == "ignoreCacheControl":
            suggest = "ignore_cache_control"
        elif key == "ignoreSetCookie":
            suggest = "ignore_set_cookie"
        elif key == "noCacheSwitch":
            suggest = "no_cache_switch"
        elif key == "reValidate":
            suggest = "re_validate"
        elif key == "rulePaths":
            suggest = "rule_paths"
        elif key == "ruleType":
            suggest = "rule_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CdnDomainRuleCach. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CdnDomainRuleCach.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CdnDomainRuleCach.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cache_time: int,
                 compare_max_age: Optional[str] = None,
                 follow_origin_switch: Optional[str] = None,
                 ignore_cache_control: Optional[str] = None,
                 ignore_set_cookie: Optional[str] = None,
                 no_cache_switch: Optional[str] = None,
                 re_validate: Optional[str] = None,
                 rule_paths: Optional[Sequence[str]] = None,
                 rule_type: Optional[str] = None,
                 switch: Optional[str] = None):
        pulumi.set(__self__, "cache_time", cache_time)
        if compare_max_age is not None:
            pulumi.set(__self__, "compare_max_age", compare_max_age)
        if follow_origin_switch is not None:
            pulumi.set(__self__, "follow_origin_switch", follow_origin_switch)
        if ignore_cache_control is not None:
            pulumi.set(__self__, "ignore_cache_control", ignore_cache_control)
        if ignore_set_cookie is not None:
            pulumi.set(__self__, "ignore_set_cookie", ignore_set_cookie)
        if no_cache_switch is not None:
            pulumi.set(__self__, "no_cache_switch", no_cache_switch)
        if re_validate is not None:
            pulumi.set(__self__, "re_validate", re_validate)
        if rule_paths is not None:
            pulumi.set(__self__, "rule_paths", rule_paths)
        if rule_type is not None:
            pulumi.set(__self__, "rule_type", rule_type)
        if switch is not None:
            pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="cacheTime")
    def cache_time(self) -> int:
        return pulumi.get(self, "cache_time")

    @property
    @pulumi.getter(name="compareMaxAge")
    def compare_max_age(self) -> Optional[str]:
        return pulumi.get(self, "compare_max_age")

    @property
    @pulumi.getter(name="followOriginSwitch")
    def follow_origin_switch(self) -> Optional[str]:
        return pulumi.get(self, "follow_origin_switch")

    @property
    @pulumi.getter(name="ignoreCacheControl")
    def ignore_cache_control(self) -> Optional[str]:
        return pulumi.get(self, "ignore_cache_control")

    @property
    @pulumi.getter(name="ignoreSetCookie")
    def ignore_set_cookie(self) -> Optional[str]:
        return pulumi.get(self, "ignore_set_cookie")

    @property
    @pulumi.getter(name="noCacheSwitch")
    def no_cache_switch(self) -> Optional[str]:
        return pulumi.get(self, "no_cache_switch")

    @property
    @pulumi.getter(name="reValidate")
    def re_validate(self) -> Optional[str]:
        return pulumi.get(self, "re_validate")

    @property
    @pulumi.getter(name="rulePaths")
    def rule_paths(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "rule_paths")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> Optional[str]:
        return pulumi.get(self, "rule_type")

    @property
    @pulumi.getter
    def switch(self) -> Optional[str]:
        return pulumi.get(self, "switch")


@pulumi.output_type
class DomainsDomainListResult(dict):
    def __init__(__self__, *,
                 area: str,
                 cname: str,
                 create_time: str,
                 domain: str,
                 full_url_cache: bool,
                 https_configs: Sequence['outputs.DomainsDomainListHttpsConfigResult'],
                 id: str,
                 origins: Sequence['outputs.DomainsDomainListOriginResult'],
                 project_id: int,
                 range_origin_switch: str,
                 request_headers: Sequence['outputs.DomainsDomainListRequestHeaderResult'],
                 rule_caches: Sequence['outputs.DomainsDomainListRuleCachResult'],
                 service_type: str,
                 status: str,
                 tags: Mapping[str, Any],
                 update_time: str):
        pulumi.set(__self__, "area", area)
        pulumi.set(__self__, "cname", cname)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "full_url_cache", full_url_cache)
        pulumi.set(__self__, "https_configs", https_configs)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "origins", origins)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "range_origin_switch", range_origin_switch)
        pulumi.set(__self__, "request_headers", request_headers)
        pulumi.set(__self__, "rule_caches", rule_caches)
        pulumi.set(__self__, "service_type", service_type)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def area(self) -> str:
        return pulumi.get(self, "area")

    @property
    @pulumi.getter
    def cname(self) -> str:
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="fullUrlCache")
    def full_url_cache(self) -> bool:
        return pulumi.get(self, "full_url_cache")

    @property
    @pulumi.getter(name="httpsConfigs")
    def https_configs(self) -> Sequence['outputs.DomainsDomainListHttpsConfigResult']:
        return pulumi.get(self, "https_configs")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def origins(self) -> Sequence['outputs.DomainsDomainListOriginResult']:
        return pulumi.get(self, "origins")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> int:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="rangeOriginSwitch")
    def range_origin_switch(self) -> str:
        return pulumi.get(self, "range_origin_switch")

    @property
    @pulumi.getter(name="requestHeaders")
    def request_headers(self) -> Sequence['outputs.DomainsDomainListRequestHeaderResult']:
        return pulumi.get(self, "request_headers")

    @property
    @pulumi.getter(name="ruleCaches")
    def rule_caches(self) -> Sequence['outputs.DomainsDomainListRuleCachResult']:
        return pulumi.get(self, "rule_caches")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        return pulumi.get(self, "update_time")


@pulumi.output_type
class DomainsDomainListHttpsConfigResult(dict):
    def __init__(__self__, *,
                 http2_switch: str,
                 https_switch: str,
                 ocsp_stapling_switch: str,
                 spdy_switch: str,
                 verify_client: str):
        pulumi.set(__self__, "http2_switch", http2_switch)
        pulumi.set(__self__, "https_switch", https_switch)
        pulumi.set(__self__, "ocsp_stapling_switch", ocsp_stapling_switch)
        pulumi.set(__self__, "spdy_switch", spdy_switch)
        pulumi.set(__self__, "verify_client", verify_client)

    @property
    @pulumi.getter(name="http2Switch")
    def http2_switch(self) -> str:
        return pulumi.get(self, "http2_switch")

    @property
    @pulumi.getter(name="httpsSwitch")
    def https_switch(self) -> str:
        return pulumi.get(self, "https_switch")

    @property
    @pulumi.getter(name="ocspStaplingSwitch")
    def ocsp_stapling_switch(self) -> str:
        return pulumi.get(self, "ocsp_stapling_switch")

    @property
    @pulumi.getter(name="spdySwitch")
    def spdy_switch(self) -> str:
        return pulumi.get(self, "spdy_switch")

    @property
    @pulumi.getter(name="verifyClient")
    def verify_client(self) -> str:
        return pulumi.get(self, "verify_client")


@pulumi.output_type
class DomainsDomainListOriginResult(dict):
    def __init__(__self__, *,
                 backup_origin_lists: Sequence[str],
                 backup_origin_type: str,
                 backup_server_name: str,
                 cos_private_access: str,
                 origin_lists: Sequence[str],
                 origin_pull_protocol: str,
                 origin_type: str,
                 server_name: str):
        pulumi.set(__self__, "backup_origin_lists", backup_origin_lists)
        pulumi.set(__self__, "backup_origin_type", backup_origin_type)
        pulumi.set(__self__, "backup_server_name", backup_server_name)
        pulumi.set(__self__, "cos_private_access", cos_private_access)
        pulumi.set(__self__, "origin_lists", origin_lists)
        pulumi.set(__self__, "origin_pull_protocol", origin_pull_protocol)
        pulumi.set(__self__, "origin_type", origin_type)
        pulumi.set(__self__, "server_name", server_name)

    @property
    @pulumi.getter(name="backupOriginLists")
    def backup_origin_lists(self) -> Sequence[str]:
        return pulumi.get(self, "backup_origin_lists")

    @property
    @pulumi.getter(name="backupOriginType")
    def backup_origin_type(self) -> str:
        return pulumi.get(self, "backup_origin_type")

    @property
    @pulumi.getter(name="backupServerName")
    def backup_server_name(self) -> str:
        return pulumi.get(self, "backup_server_name")

    @property
    @pulumi.getter(name="cosPrivateAccess")
    def cos_private_access(self) -> str:
        return pulumi.get(self, "cos_private_access")

    @property
    @pulumi.getter(name="originLists")
    def origin_lists(self) -> Sequence[str]:
        return pulumi.get(self, "origin_lists")

    @property
    @pulumi.getter(name="originPullProtocol")
    def origin_pull_protocol(self) -> str:
        return pulumi.get(self, "origin_pull_protocol")

    @property
    @pulumi.getter(name="originType")
    def origin_type(self) -> str:
        return pulumi.get(self, "origin_type")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> str:
        return pulumi.get(self, "server_name")


@pulumi.output_type
class DomainsDomainListRequestHeaderResult(dict):
    def __init__(__self__, *,
                 header_rules: Sequence['outputs.DomainsDomainListRequestHeaderHeaderRuleResult'],
                 switch: str):
        pulumi.set(__self__, "header_rules", header_rules)
        pulumi.set(__self__, "switch", switch)

    @property
    @pulumi.getter(name="headerRules")
    def header_rules(self) -> Sequence['outputs.DomainsDomainListRequestHeaderHeaderRuleResult']:
        return pulumi.get(self, "header_rules")

    @property
    @pulumi.getter
    def switch(self) -> str:
        return pulumi.get(self, "switch")


@pulumi.output_type
class DomainsDomainListRequestHeaderHeaderRuleResult(dict):
    def __init__(__self__, *,
                 header_mode: str,
                 header_name: str,
                 header_value: str,
                 rule_paths: Sequence[str],
                 rule_type: str):
        pulumi.set(__self__, "header_mode", header_mode)
        pulumi.set(__self__, "header_name", header_name)
        pulumi.set(__self__, "header_value", header_value)
        pulumi.set(__self__, "rule_paths", rule_paths)
        pulumi.set(__self__, "rule_type", rule_type)

    @property
    @pulumi.getter(name="headerMode")
    def header_mode(self) -> str:
        return pulumi.get(self, "header_mode")

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> str:
        return pulumi.get(self, "header_name")

    @property
    @pulumi.getter(name="headerValue")
    def header_value(self) -> str:
        return pulumi.get(self, "header_value")

    @property
    @pulumi.getter(name="rulePaths")
    def rule_paths(self) -> Sequence[str]:
        return pulumi.get(self, "rule_paths")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> str:
        return pulumi.get(self, "rule_type")


@pulumi.output_type
class DomainsDomainListRuleCachResult(dict):
    def __init__(__self__, *,
                 cache_time: int,
                 follow_origin_switch: str,
                 ignore_set_cookie: str,
                 no_cache_switch: str,
                 re_validate: str,
                 rule_paths: Sequence[str],
                 rule_type: str,
                 switch: str,
                 compare_max_age: Optional[str] = None,
                 ignore_cache_control: Optional[str] = None):
        pulumi.set(__self__, "cache_time", cache_time)
        pulumi.set(__self__, "follow_origin_switch", follow_origin_switch)
        pulumi.set(__self__, "ignore_set_cookie", ignore_set_cookie)
        pulumi.set(__self__, "no_cache_switch", no_cache_switch)
        pulumi.set(__self__, "re_validate", re_validate)
        pulumi.set(__self__, "rule_paths", rule_paths)
        pulumi.set(__self__, "rule_type", rule_type)
        pulumi.set(__self__, "switch", switch)
        if compare_max_age is not None:
            pulumi.set(__self__, "compare_max_age", compare_max_age)
        if ignore_cache_control is not None:
            pulumi.set(__self__, "ignore_cache_control", ignore_cache_control)

    @property
    @pulumi.getter(name="cacheTime")
    def cache_time(self) -> int:
        return pulumi.get(self, "cache_time")

    @property
    @pulumi.getter(name="followOriginSwitch")
    def follow_origin_switch(self) -> str:
        return pulumi.get(self, "follow_origin_switch")

    @property
    @pulumi.getter(name="ignoreSetCookie")
    def ignore_set_cookie(self) -> str:
        return pulumi.get(self, "ignore_set_cookie")

    @property
    @pulumi.getter(name="noCacheSwitch")
    def no_cache_switch(self) -> str:
        return pulumi.get(self, "no_cache_switch")

    @property
    @pulumi.getter(name="reValidate")
    def re_validate(self) -> str:
        return pulumi.get(self, "re_validate")

    @property
    @pulumi.getter(name="rulePaths")
    def rule_paths(self) -> Sequence[str]:
        return pulumi.get(self, "rule_paths")

    @property
    @pulumi.getter(name="ruleType")
    def rule_type(self) -> str:
        return pulumi.get(self, "rule_type")

    @property
    @pulumi.getter
    def switch(self) -> str:
        return pulumi.get(self, "switch")

    @property
    @pulumi.getter(name="compareMaxAge")
    def compare_max_age(self) -> Optional[str]:
        return pulumi.get(self, "compare_max_age")

    @property
    @pulumi.getter(name="ignoreCacheControl")
    def ignore_cache_control(self) -> Optional[str]:
        return pulumi.get(self, "ignore_cache_control")


