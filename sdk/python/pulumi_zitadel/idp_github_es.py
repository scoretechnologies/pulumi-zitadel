# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IdpGithubEsArgs', 'IdpGithubEs']

@pulumi.input_type
class IdpGithubEsArgs:
    def __init__(__self__, *,
                 authorization_endpoint: pulumi.Input[str],
                 client_id: pulumi.Input[str],
                 client_secret: pulumi.Input[str],
                 is_auto_creation: pulumi.Input[bool],
                 is_auto_update: pulumi.Input[bool],
                 is_creation_allowed: pulumi.Input[bool],
                 is_linking_allowed: pulumi.Input[bool],
                 token_endpoint: pulumi.Input[str],
                 user_endpoint: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a IdpGithubEs resource.
        :param pulumi.Input[str] authorization_endpoint: the providers authorization endpoint
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] token_endpoint: the providers token endpoint
        :param pulumi.Input[str] user_endpoint: the providers user endpoint
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        """
        pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "client_secret", client_secret)
        pulumi.set(__self__, "is_auto_creation", is_auto_creation)
        pulumi.set(__self__, "is_auto_update", is_auto_update)
        pulumi.set(__self__, "is_creation_allowed", is_creation_allowed)
        pulumi.set(__self__, "is_linking_allowed", is_linking_allowed)
        pulumi.set(__self__, "token_endpoint", token_endpoint)
        pulumi.set(__self__, "user_endpoint", user_endpoint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> pulumi.Input[str]:
        """
        the providers authorization endpoint
        """
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Input[str]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> pulumi.Input[bool]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @is_auto_creation.setter
    def is_auto_creation(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_auto_creation", value)

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> pulumi.Input[bool]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @is_auto_update.setter
    def is_auto_update(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_auto_update", value)

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> pulumi.Input[bool]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @is_creation_allowed.setter
    def is_creation_allowed(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_creation_allowed", value)

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> pulumi.Input[bool]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @is_linking_allowed.setter
    def is_linking_allowed(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_linking_allowed", value)

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> pulumi.Input[str]:
        """
        the providers token endpoint
        """
        return pulumi.get(self, "token_endpoint")

    @token_endpoint.setter
    def token_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "token_endpoint", value)

    @property
    @pulumi.getter(name="userEndpoint")
    def user_endpoint(self) -> pulumi.Input[str]:
        """
        the providers user endpoint
        """
        return pulumi.get(self, "user_endpoint")

    @user_endpoint.setter
    def user_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


@pulumi.input_type
class _IdpGithubEsState:
    def __init__(__self__, *,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_endpoint: Optional[pulumi.Input[str]] = None,
                 user_endpoint: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IdpGithubEs resources.
        :param pulumi.Input[str] authorization_endpoint: the providers authorization endpoint
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        :param pulumi.Input[str] token_endpoint: the providers token endpoint
        :param pulumi.Input[str] user_endpoint: the providers user endpoint
        """
        if authorization_endpoint is not None:
            pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if is_auto_creation is not None:
            pulumi.set(__self__, "is_auto_creation", is_auto_creation)
        if is_auto_update is not None:
            pulumi.set(__self__, "is_auto_update", is_auto_update)
        if is_creation_allowed is not None:
            pulumi.set(__self__, "is_creation_allowed", is_creation_allowed)
        if is_linking_allowed is not None:
            pulumi.set(__self__, "is_linking_allowed", is_linking_allowed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if token_endpoint is not None:
            pulumi.set(__self__, "token_endpoint", token_endpoint)
        if user_endpoint is not None:
            pulumi.set(__self__, "user_endpoint", user_endpoint)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        the providers authorization endpoint
        """
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @is_auto_creation.setter
    def is_auto_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_creation", value)

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @is_auto_update.setter
    def is_auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_auto_update", value)

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @is_creation_allowed.setter
    def is_creation_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_creation_allowed", value)

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @is_linking_allowed.setter
    def is_linking_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_linking_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        the providers token endpoint
        """
        return pulumi.get(self, "token_endpoint")

    @token_endpoint.setter
    def token_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_endpoint", value)

    @property
    @pulumi.getter(name="userEndpoint")
    def user_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        the providers user endpoint
        """
        return pulumi.get(self, "user_endpoint")

    @user_endpoint.setter
    def user_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_endpoint", value)


class IdpGithubEs(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_endpoint: Optional[pulumi.Input[str]] = None,
                 user_endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a GitHub Enterprise IDP on the instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zitadel as zitadel

        github_es = zitadel.IdpGithubEs("githubEs",
            authorization_endpoint="https://auth.endpoint",
            client_id="86a165...",
            client_secret="*****afdbac18",
            is_auto_creation=False,
            is_auto_update=True,
            is_creation_allowed=True,
            is_linking_allowed=False,
            scopes=[
                "openid",
                "profile",
                "email",
            ],
            token_endpoint="https://token.endpoint",
            user_endpoint="https://user.endpoint")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_endpoint: the providers authorization endpoint
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        :param pulumi.Input[str] token_endpoint: the providers token endpoint
        :param pulumi.Input[str] user_endpoint: the providers user endpoint
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdpGithubEsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a GitHub Enterprise IDP on the instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_zitadel as zitadel

        github_es = zitadel.IdpGithubEs("githubEs",
            authorization_endpoint="https://auth.endpoint",
            client_id="86a165...",
            client_secret="*****afdbac18",
            is_auto_creation=False,
            is_auto_update=True,
            is_creation_allowed=True,
            is_linking_allowed=False,
            scopes=[
                "openid",
                "profile",
                "email",
            ],
            token_endpoint="https://token.endpoint",
            user_endpoint="https://user.endpoint")
        ```

        :param str resource_name: The name of the resource.
        :param IdpGithubEsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdpGithubEsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 is_auto_creation: Optional[pulumi.Input[bool]] = None,
                 is_auto_update: Optional[pulumi.Input[bool]] = None,
                 is_creation_allowed: Optional[pulumi.Input[bool]] = None,
                 is_linking_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_endpoint: Optional[pulumi.Input[str]] = None,
                 user_endpoint: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdpGithubEsArgs.__new__(IdpGithubEsArgs)

            if authorization_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_endpoint'")
            __props__.__dict__["authorization_endpoint"] = authorization_endpoint
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if client_secret is None and not opts.urn:
                raise TypeError("Missing required property 'client_secret'")
            __props__.__dict__["client_secret"] = client_secret
            if is_auto_creation is None and not opts.urn:
                raise TypeError("Missing required property 'is_auto_creation'")
            __props__.__dict__["is_auto_creation"] = is_auto_creation
            if is_auto_update is None and not opts.urn:
                raise TypeError("Missing required property 'is_auto_update'")
            __props__.__dict__["is_auto_update"] = is_auto_update
            if is_creation_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'is_creation_allowed'")
            __props__.__dict__["is_creation_allowed"] = is_creation_allowed
            if is_linking_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'is_linking_allowed'")
            __props__.__dict__["is_linking_allowed"] = is_linking_allowed
            __props__.__dict__["name"] = name
            __props__.__dict__["scopes"] = scopes
            if token_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'token_endpoint'")
            __props__.__dict__["token_endpoint"] = token_endpoint
            if user_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'user_endpoint'")
            __props__.__dict__["user_endpoint"] = user_endpoint
        super(IdpGithubEs, __self__).__init__(
            'zitadel:index/idpGithubEs:IdpGithubEs',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_endpoint: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            is_auto_creation: Optional[pulumi.Input[bool]] = None,
            is_auto_update: Optional[pulumi.Input[bool]] = None,
            is_creation_allowed: Optional[pulumi.Input[bool]] = None,
            is_linking_allowed: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            token_endpoint: Optional[pulumi.Input[str]] = None,
            user_endpoint: Optional[pulumi.Input[str]] = None) -> 'IdpGithubEs':
        """
        Get an existing IdpGithubEs resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_endpoint: the providers authorization endpoint
        :param pulumi.Input[str] client_id: client id generated by the identity provider
        :param pulumi.Input[str] client_secret: client secret generated by the identity provider
        :param pulumi.Input[bool] is_auto_creation: enable if a new account in ZITADEL should be created automatically on login with an external account
        :param pulumi.Input[bool] is_auto_update: enable if a the ZITADEL account fields should be updated automatically on each login
        :param pulumi.Input[bool] is_creation_allowed: enable if users should be able to create a new account in ZITADEL when using an external account
        :param pulumi.Input[bool] is_linking_allowed: enable if users should be able to link an existing ZITADEL user with an external account
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: the scopes requested by ZITADEL during the request on the identity provider
        :param pulumi.Input[str] token_endpoint: the providers token endpoint
        :param pulumi.Input[str] user_endpoint: the providers user endpoint
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdpGithubEsState.__new__(_IdpGithubEsState)

        __props__.__dict__["authorization_endpoint"] = authorization_endpoint
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["is_auto_creation"] = is_auto_creation
        __props__.__dict__["is_auto_update"] = is_auto_update
        __props__.__dict__["is_creation_allowed"] = is_creation_allowed
        __props__.__dict__["is_linking_allowed"] = is_linking_allowed
        __props__.__dict__["name"] = name
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["token_endpoint"] = token_endpoint
        __props__.__dict__["user_endpoint"] = user_endpoint
        return IdpGithubEs(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> pulumi.Output[str]:
        """
        the providers authorization endpoint
        """
        return pulumi.get(self, "authorization_endpoint")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        client id generated by the identity provider
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[str]:
        """
        client secret generated by the identity provider
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="isAutoCreation")
    def is_auto_creation(self) -> pulumi.Output[bool]:
        """
        enable if a new account in ZITADEL should be created automatically on login with an external account
        """
        return pulumi.get(self, "is_auto_creation")

    @property
    @pulumi.getter(name="isAutoUpdate")
    def is_auto_update(self) -> pulumi.Output[bool]:
        """
        enable if a the ZITADEL account fields should be updated automatically on each login
        """
        return pulumi.get(self, "is_auto_update")

    @property
    @pulumi.getter(name="isCreationAllowed")
    def is_creation_allowed(self) -> pulumi.Output[bool]:
        """
        enable if users should be able to create a new account in ZITADEL when using an external account
        """
        return pulumi.get(self, "is_creation_allowed")

    @property
    @pulumi.getter(name="isLinkingAllowed")
    def is_linking_allowed(self) -> pulumi.Output[bool]:
        """
        enable if users should be able to link an existing ZITADEL user with an external account
        """
        return pulumi.get(self, "is_linking_allowed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        the scopes requested by ZITADEL during the request on the identity provider
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> pulumi.Output[str]:
        """
        the providers token endpoint
        """
        return pulumi.get(self, "token_endpoint")

    @property
    @pulumi.getter(name="userEndpoint")
    def user_endpoint(self) -> pulumi.Output[str]:
        """
        the providers user endpoint
        """
        return pulumi.get(self, "user_endpoint")

