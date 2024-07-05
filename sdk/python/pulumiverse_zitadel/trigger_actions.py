# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TriggerActionsArgs', 'TriggerActions']

@pulumi.input_type
class TriggerActionsArgs:
    def __init__(__self__, *,
                 action_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 flow_type: pulumi.Input[str],
                 trigger_type: pulumi.Input[str],
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TriggerActions resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] action_ids: IDs of the triggered actions
        :param pulumi.Input[str] flow_type: Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        :param pulumi.Input[str] trigger_type: Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        :param pulumi.Input[str] org_id: ID of the organization
        """
        TriggerActionsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action_ids=action_ids,
            flow_type=flow_type,
            trigger_type=trigger_type,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
             flow_type: pulumi.Input[str],
             trigger_type: pulumi.Input[str],
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'actionIds' in kwargs:
            action_ids = kwargs['actionIds']
        if 'flowType' in kwargs:
            flow_type = kwargs['flowType']
        if 'triggerType' in kwargs:
            trigger_type = kwargs['triggerType']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        _setter("action_ids", action_ids)
        _setter("flow_type", flow_type)
        _setter("trigger_type", trigger_type)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="actionIds")
    def action_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        IDs of the triggered actions
        """
        return pulumi.get(self, "action_ids")

    @action_ids.setter
    def action_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "action_ids", value)

    @property
    @pulumi.getter(name="flowType")
    def flow_type(self) -> pulumi.Input[str]:
        """
        Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        """
        return pulumi.get(self, "flow_type")

    @flow_type.setter
    def flow_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "flow_type", value)

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> pulumi.Input[str]:
        """
        Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        """
        return pulumi.get(self, "trigger_type")

    @trigger_type.setter
    def trigger_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "trigger_type", value)

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


@pulumi.input_type
class _TriggerActionsState:
    def __init__(__self__, *,
                 action_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 flow_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TriggerActions resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] action_ids: IDs of the triggered actions
        :param pulumi.Input[str] flow_type: Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] trigger_type: Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        """
        _TriggerActionsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action_ids=action_ids,
            flow_type=flow_type,
            org_id=org_id,
            trigger_type=trigger_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             flow_type: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             trigger_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'actionIds' in kwargs:
            action_ids = kwargs['actionIds']
        if 'flowType' in kwargs:
            flow_type = kwargs['flowType']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'triggerType' in kwargs:
            trigger_type = kwargs['triggerType']

        if action_ids is not None:
            _setter("action_ids", action_ids)
        if flow_type is not None:
            _setter("flow_type", flow_type)
        if org_id is not None:
            _setter("org_id", org_id)
        if trigger_type is not None:
            _setter("trigger_type", trigger_type)

    @property
    @pulumi.getter(name="actionIds")
    def action_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of the triggered actions
        """
        return pulumi.get(self, "action_ids")

    @action_ids.setter
    def action_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "action_ids", value)

    @property
    @pulumi.getter(name="flowType")
    def flow_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        """
        return pulumi.get(self, "flow_type")

    @flow_type.setter
    def flow_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_type", value)

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
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> Optional[pulumi.Input[str]]:
        """
        Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        """
        return pulumi.get(self, "trigger_type")

    @trigger_type.setter
    def trigger_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_type", value)


class TriggerActions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 flow_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing triggers, when actions get started

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.TriggerActions("default",
            org_id=data["zitadel_org"]["default"]["id"],
            flow_type="FLOW_TYPE_CUSTOMISE_TOKEN",
            trigger_type="TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION",
            action_ids=[data["zitadel_action"]["default"]["id"]])
        ```

        ## Import

        bash The resource can be imported using the ID format `<flow_type:trigger_type[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/triggerActions:TriggerActions imported 'FLOW_TYPE_EXTERNAL_AUTHENTICATION:TRIGGER_TYPE_POST_CREATION:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] action_ids: IDs of the triggered actions
        :param pulumi.Input[str] flow_type: Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] trigger_type: Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TriggerActionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing triggers, when actions get started

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.TriggerActions("default",
            org_id=data["zitadel_org"]["default"]["id"],
            flow_type="FLOW_TYPE_CUSTOMISE_TOKEN",
            trigger_type="TRIGGER_TYPE_PRE_ACCESS_TOKEN_CREATION",
            action_ids=[data["zitadel_action"]["default"]["id"]])
        ```

        ## Import

        bash The resource can be imported using the ID format `<flow_type:trigger_type[:org_id]>`, e.g.

        ```sh
         $ pulumi import zitadel:index/triggerActions:TriggerActions imported 'FLOW_TYPE_EXTERNAL_AUTHENTICATION:TRIGGER_TYPE_POST_CREATION:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param TriggerActionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TriggerActionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TriggerActionsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 flow_type: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 trigger_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TriggerActionsArgs.__new__(TriggerActionsArgs)

            if action_ids is None and not opts.urn:
                raise TypeError("Missing required property 'action_ids'")
            __props__.__dict__["action_ids"] = action_ids
            if flow_type is None and not opts.urn:
                raise TypeError("Missing required property 'flow_type'")
            __props__.__dict__["flow_type"] = flow_type
            __props__.__dict__["org_id"] = org_id
            if trigger_type is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_type'")
            __props__.__dict__["trigger_type"] = trigger_type
        super(TriggerActions, __self__).__init__(
            'zitadel:index/triggerActions:TriggerActions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            flow_type: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            trigger_type: Optional[pulumi.Input[str]] = None) -> 'TriggerActions':
        """
        Get an existing TriggerActions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] action_ids: IDs of the triggered actions
        :param pulumi.Input[str] flow_type: Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[str] trigger_type: Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TriggerActionsState.__new__(_TriggerActionsState)

        __props__.__dict__["action_ids"] = action_ids
        __props__.__dict__["flow_type"] = flow_type
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["trigger_type"] = trigger_type
        return TriggerActions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="actionIds")
    def action_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        IDs of the triggered actions
        """
        return pulumi.get(self, "action_ids")

    @property
    @pulumi.getter(name="flowType")
    def flow_type(self) -> pulumi.Output[str]:
        """
        Type of the flow to which the action triggers belong, supported values: FLOW*TYPE*EXTERNAL*AUTHENTICATION, FLOW*TYPE*CUSTOMISE*TOKEN, FLOW*TYPE*INTERNAL*AUTHENTICATION, FLOW*TYPE*SAML*RESPONSE
        """
        return pulumi.get(self, "flow_type")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="triggerType")
    def trigger_type(self) -> pulumi.Output[str]:
        """
        Trigger type on when the actions get triggered, supported values: TRIGGER*TYPE*POST*AUTHENTICATION, TRIGGER*TYPE*PRE*CREATION, TRIGGER*TYPE*POST*CREATION, TRIGGER*TYPE*PRE*USERINFO*CREATION, TRIGGER*TYPE*PRE*ACCESS*TOKEN*CREATION, TRIGGER*TYPE*PRE*SAML*RESPONSE_CREATION
        """
        return pulumi.get(self, "trigger_type")

