# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['OrgIdpJwtArgs', 'OrgIdpJwt']

@pulumi.input_type
class OrgIdpJwtArgs:
    def __init__(__self__, *,
                 auto_register: pulumi.Input[bool],
                 header_name: pulumi.Input[str],
                 issuer: pulumi.Input[str],
                 jwt_endpoint: pulumi.Input[str],
                 keys_endpoint: pulumi.Input[str],
                 org_id: pulumi.Input[str],
                 styling_type: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrgIdpJwt resource.
        :param pulumi.Input[bool] auto_register: auto register for users from this idp
        :param pulumi.Input[str] header_name: the name of the header where the JWT is sent in, default is authorization
        :param pulumi.Input[str] issuer: the issuer of the jwt (for validation)
        :param pulumi.Input[str] jwt_endpoint: the endpoint where the jwt can be extracted
        :param pulumi.Input[str] keys_endpoint: the endpoint to the key (JWK) which are used to sign the JWT with
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] styling_type: Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        :param pulumi.Input[str] name: Name of the IDP
        """
        pulumi.set(__self__, "auto_register", auto_register)
        pulumi.set(__self__, "header_name", header_name)
        pulumi.set(__self__, "issuer", issuer)
        pulumi.set(__self__, "jwt_endpoint", jwt_endpoint)
        pulumi.set(__self__, "keys_endpoint", keys_endpoint)
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "styling_type", styling_type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="autoRegister")
    def auto_register(self) -> pulumi.Input[bool]:
        """
        auto register for users from this idp
        """
        return pulumi.get(self, "auto_register")

    @auto_register.setter
    def auto_register(self, value: pulumi.Input[bool]):
        pulumi.set(self, "auto_register", value)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> pulumi.Input[str]:
        """
        the name of the header where the JWT is sent in, default is authorization
        """
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Input[str]:
        """
        the issuer of the jwt (for validation)
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="jwtEndpoint")
    def jwt_endpoint(self) -> pulumi.Input[str]:
        """
        the endpoint where the jwt can be extracted
        """
        return pulumi.get(self, "jwt_endpoint")

    @jwt_endpoint.setter
    def jwt_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "jwt_endpoint", value)

    @property
    @pulumi.getter(name="keysEndpoint")
    def keys_endpoint(self) -> pulumi.Input[str]:
        """
        the endpoint to the key (JWK) which are used to sign the JWT with
        """
        return pulumi.get(self, "keys_endpoint")

    @keys_endpoint.setter
    def keys_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "keys_endpoint", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="stylingType")
    def styling_type(self) -> pulumi.Input[str]:
        """
        Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        """
        return pulumi.get(self, "styling_type")

    @styling_type.setter
    def styling_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "styling_type", value)

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


@pulumi.input_type
class _OrgIdpJwtState:
    def __init__(__self__, *,
                 auto_register: Optional[pulumi.Input[bool]] = None,
                 header_name: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 jwt_endpoint: Optional[pulumi.Input[str]] = None,
                 keys_endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 styling_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrgIdpJwt resources.
        :param pulumi.Input[bool] auto_register: auto register for users from this idp
        :param pulumi.Input[str] header_name: the name of the header where the JWT is sent in, default is authorization
        :param pulumi.Input[str] issuer: the issuer of the jwt (for validation)
        :param pulumi.Input[str] jwt_endpoint: the endpoint where the jwt can be extracted
        :param pulumi.Input[str] keys_endpoint: the endpoint to the key (JWK) which are used to sign the JWT with
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] styling_type: Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        """
        if auto_register is not None:
            pulumi.set(__self__, "auto_register", auto_register)
        if header_name is not None:
            pulumi.set(__self__, "header_name", header_name)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if jwt_endpoint is not None:
            pulumi.set(__self__, "jwt_endpoint", jwt_endpoint)
        if keys_endpoint is not None:
            pulumi.set(__self__, "keys_endpoint", keys_endpoint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if styling_type is not None:
            pulumi.set(__self__, "styling_type", styling_type)

    @property
    @pulumi.getter(name="autoRegister")
    def auto_register(self) -> Optional[pulumi.Input[bool]]:
        """
        auto register for users from this idp
        """
        return pulumi.get(self, "auto_register")

    @auto_register.setter
    def auto_register(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_register", value)

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> Optional[pulumi.Input[str]]:
        """
        the name of the header where the JWT is sent in, default is authorization
        """
        return pulumi.get(self, "header_name")

    @header_name.setter
    def header_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "header_name", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        the issuer of the jwt (for validation)
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="jwtEndpoint")
    def jwt_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        the endpoint where the jwt can be extracted
        """
        return pulumi.get(self, "jwt_endpoint")

    @jwt_endpoint.setter
    def jwt_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "jwt_endpoint", value)

    @property
    @pulumi.getter(name="keysEndpoint")
    def keys_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        the endpoint to the key (JWK) which are used to sign the JWT with
        """
        return pulumi.get(self, "keys_endpoint")

    @keys_endpoint.setter
    def keys_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "keys_endpoint", value)

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
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="stylingType")
    def styling_type(self) -> Optional[pulumi.Input[str]]:
        """
        Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        """
        return pulumi.get(self, "styling_type")

    @styling_type.setter
    def styling_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "styling_type", value)


class OrgIdpJwt(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_register: Optional[pulumi.Input[bool]] = None,
                 header_name: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 jwt_endpoint: Optional[pulumi.Input[str]] = None,
                 keys_endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 styling_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a generic JWT IdP of the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        jwt_idp = zitadel.OrgIdpJwt("jwtIdp",
            org_id=zitadel_org["org"]["id"],
            styling_type="STYLING_TYPE_UNSPECIFIED",
            jwt_endpoint="https://jwtendpoint.com",
            issuer="https://google.com",
            keys_endpoint="https://jwtendpoint.com/keys",
            header_name="x-auth-token",
            auto_register=False)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_register: auto register for users from this idp
        :param pulumi.Input[str] header_name: the name of the header where the JWT is sent in, default is authorization
        :param pulumi.Input[str] issuer: the issuer of the jwt (for validation)
        :param pulumi.Input[str] jwt_endpoint: the endpoint where the jwt can be extracted
        :param pulumi.Input[str] keys_endpoint: the endpoint to the key (JWK) which are used to sign the JWT with
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] styling_type: Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrgIdpJwtArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a generic JWT IdP of the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        jwt_idp = zitadel.OrgIdpJwt("jwtIdp",
            org_id=zitadel_org["org"]["id"],
            styling_type="STYLING_TYPE_UNSPECIFIED",
            jwt_endpoint="https://jwtendpoint.com",
            issuer="https://google.com",
            keys_endpoint="https://jwtendpoint.com/keys",
            header_name="x-auth-token",
            auto_register=False)
        ```

        :param str resource_name: The name of the resource.
        :param OrgIdpJwtArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrgIdpJwtArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_register: Optional[pulumi.Input[bool]] = None,
                 header_name: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 jwt_endpoint: Optional[pulumi.Input[str]] = None,
                 keys_endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 styling_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrgIdpJwtArgs.__new__(OrgIdpJwtArgs)

            if auto_register is None and not opts.urn:
                raise TypeError("Missing required property 'auto_register'")
            __props__.__dict__["auto_register"] = auto_register
            if header_name is None and not opts.urn:
                raise TypeError("Missing required property 'header_name'")
            __props__.__dict__["header_name"] = header_name
            if issuer is None and not opts.urn:
                raise TypeError("Missing required property 'issuer'")
            __props__.__dict__["issuer"] = issuer
            if jwt_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'jwt_endpoint'")
            __props__.__dict__["jwt_endpoint"] = jwt_endpoint
            if keys_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'keys_endpoint'")
            __props__.__dict__["keys_endpoint"] = keys_endpoint
            __props__.__dict__["name"] = name
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            if styling_type is None and not opts.urn:
                raise TypeError("Missing required property 'styling_type'")
            __props__.__dict__["styling_type"] = styling_type
        super(OrgIdpJwt, __self__).__init__(
            'zitadel:index/orgIdpJwt:OrgIdpJwt',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_register: Optional[pulumi.Input[bool]] = None,
            header_name: Optional[pulumi.Input[str]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            jwt_endpoint: Optional[pulumi.Input[str]] = None,
            keys_endpoint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            styling_type: Optional[pulumi.Input[str]] = None) -> 'OrgIdpJwt':
        """
        Get an existing OrgIdpJwt resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_register: auto register for users from this idp
        :param pulumi.Input[str] header_name: the name of the header where the JWT is sent in, default is authorization
        :param pulumi.Input[str] issuer: the issuer of the jwt (for validation)
        :param pulumi.Input[str] jwt_endpoint: the endpoint where the jwt can be extracted
        :param pulumi.Input[str] keys_endpoint: the endpoint to the key (JWK) which are used to sign the JWT with
        :param pulumi.Input[str] name: Name of the IDP
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] styling_type: Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrgIdpJwtState.__new__(_OrgIdpJwtState)

        __props__.__dict__["auto_register"] = auto_register
        __props__.__dict__["header_name"] = header_name
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["jwt_endpoint"] = jwt_endpoint
        __props__.__dict__["keys_endpoint"] = keys_endpoint
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["styling_type"] = styling_type
        return OrgIdpJwt(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRegister")
    def auto_register(self) -> pulumi.Output[bool]:
        """
        auto register for users from this idp
        """
        return pulumi.get(self, "auto_register")

    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> pulumi.Output[str]:
        """
        the name of the header where the JWT is sent in, default is authorization
        """
        return pulumi.get(self, "header_name")

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[str]:
        """
        the issuer of the jwt (for validation)
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="jwtEndpoint")
    def jwt_endpoint(self) -> pulumi.Output[str]:
        """
        the endpoint where the jwt can be extracted
        """
        return pulumi.get(self, "jwt_endpoint")

    @property
    @pulumi.getter(name="keysEndpoint")
    def keys_endpoint(self) -> pulumi.Output[str]:
        """
        the endpoint to the key (JWK) which are used to sign the JWT with
        """
        return pulumi.get(self, "keys_endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the IDP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="stylingType")
    def styling_type(self) -> pulumi.Output[str]:
        """
        Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        """
        return pulumi.get(self, "styling_type")
