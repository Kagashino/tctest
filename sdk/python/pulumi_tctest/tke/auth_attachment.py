# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthAttachmentArgs', 'AuthAttachment']

@pulumi.input_type
class AuthAttachmentArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 issuer: pulumi.Input[str],
                 auto_create_discovery_anonymous_auth: Optional[pulumi.Input[bool]] = None,
                 jwks_uri: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthAttachment resource.
        :param pulumi.Input[str] cluster_id: ID of clusters.
        :param pulumi.Input[str] issuer: Specify service-account-issuer.
        :param pulumi.Input[bool] auto_create_discovery_anonymous_auth: If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
               '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        :param pulumi.Input[str] jwks_uri: Specify service-account-jwks-uri.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "issuer", issuer)
        if auto_create_discovery_anonymous_auth is not None:
            pulumi.set(__self__, "auto_create_discovery_anonymous_auth", auto_create_discovery_anonymous_auth)
        if jwks_uri is not None:
            pulumi.set(__self__, "jwks_uri", jwks_uri)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        ID of clusters.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Input[str]:
        """
        Specify service-account-issuer.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="autoCreateDiscoveryAnonymousAuth")
    def auto_create_discovery_anonymous_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
        '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        """
        return pulumi.get(self, "auto_create_discovery_anonymous_auth")

    @auto_create_discovery_anonymous_auth.setter
    def auto_create_discovery_anonymous_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_discovery_anonymous_auth", value)

    @property
    @pulumi.getter(name="jwksUri")
    def jwks_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Specify service-account-jwks-uri.
        """
        return pulumi.get(self, "jwks_uri")

    @jwks_uri.setter
    def jwks_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "jwks_uri", value)


@pulumi.input_type
class _AuthAttachmentState:
    def __init__(__self__, *,
                 auto_create_discovery_anonymous_auth: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 jwks_uri: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthAttachment resources.
        :param pulumi.Input[bool] auto_create_discovery_anonymous_auth: If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
               '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        :param pulumi.Input[str] cluster_id: ID of clusters.
        :param pulumi.Input[str] issuer: Specify service-account-issuer.
        :param pulumi.Input[str] jwks_uri: Specify service-account-jwks-uri.
        """
        if auto_create_discovery_anonymous_auth is not None:
            pulumi.set(__self__, "auto_create_discovery_anonymous_auth", auto_create_discovery_anonymous_auth)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if jwks_uri is not None:
            pulumi.set(__self__, "jwks_uri", jwks_uri)

    @property
    @pulumi.getter(name="autoCreateDiscoveryAnonymousAuth")
    def auto_create_discovery_anonymous_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
        '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        """
        return pulumi.get(self, "auto_create_discovery_anonymous_auth")

    @auto_create_discovery_anonymous_auth.setter
    def auto_create_discovery_anonymous_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_create_discovery_anonymous_auth", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of clusters.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        Specify service-account-issuer.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="jwksUri")
    def jwks_uri(self) -> Optional[pulumi.Input[str]]:
        """
        Specify service-account-jwks-uri.
        """
        return pulumi.get(self, "jwks_uri")

    @jwks_uri.setter
    def jwks_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "jwks_uri", value)


class AuthAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_discovery_anonymous_auth: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 jwks_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a AuthAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_discovery_anonymous_auth: If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
               '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        :param pulumi.Input[str] cluster_id: ID of clusters.
        :param pulumi.Input[str] issuer: Specify service-account-issuer.
        :param pulumi.Input[str] jwks_uri: Specify service-account-jwks-uri.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a AuthAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param AuthAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_create_discovery_anonymous_auth: Optional[pulumi.Input[bool]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 jwks_uri: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        else:
            opts = copy.copy(opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthAttachmentArgs.__new__(AuthAttachmentArgs)

            __props__.__dict__["auto_create_discovery_anonymous_auth"] = auto_create_discovery_anonymous_auth
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if issuer is None and not opts.urn:
                raise TypeError("Missing required property 'issuer'")
            __props__.__dict__["issuer"] = issuer
            __props__.__dict__["jwks_uri"] = jwks_uri
        super(AuthAttachment, __self__).__init__(
            'tctest:Tke/authAttachment:AuthAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_create_discovery_anonymous_auth: Optional[pulumi.Input[bool]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            jwks_uri: Optional[pulumi.Input[str]] = None) -> 'AuthAttachment':
        """
        Get an existing AuthAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_create_discovery_anonymous_auth: If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
               '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        :param pulumi.Input[str] cluster_id: ID of clusters.
        :param pulumi.Input[str] issuer: Specify service-account-issuer.
        :param pulumi.Input[str] jwks_uri: Specify service-account-jwks-uri.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthAttachmentState.__new__(_AuthAttachmentState)

        __props__.__dict__["auto_create_discovery_anonymous_auth"] = auto_create_discovery_anonymous_auth
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["jwks_uri"] = jwks_uri
        return AuthAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoCreateDiscoveryAnonymousAuth")
    def auto_create_discovery_anonymous_auth(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, the rbac rule will be created automatically which allow anonymous user to access
        '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        """
        return pulumi.get(self, "auto_create_discovery_anonymous_auth")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        ID of clusters.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[str]:
        """
        Specify service-account-issuer.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="jwksUri")
    def jwks_uri(self) -> pulumi.Output[Optional[str]]:
        """
        Specify service-account-jwks-uri.
        """
        return pulumi.get(self, "jwks_uri")

