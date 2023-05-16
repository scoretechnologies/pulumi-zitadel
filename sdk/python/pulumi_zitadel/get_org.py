# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOrgResult',
    'AwaitableGetOrgResult',
    'get_org',
    'get_org_output',
]

@pulumi.output_type
class GetOrgResult:
    """
    A collection of values returned by getOrg.
    """
    def __init__(__self__, id=None, name=None, org_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the org
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "org_id")


class AwaitableGetOrgResult(GetOrgResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgResult(
            id=self.id,
            name=self.name,
            org_id=self.org_id)


def get_org(org_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgResult:
    """
    Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    org_org = zitadel.get_org(org_id="177073608051458051")
    pulumi.export("org", org_org)
    ```


    :param str org_id: The ID of this resource.
    """
    __args__ = dict()
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('zitadel:index/getOrg:getOrg', __args__, opts=opts, typ=GetOrgResult).value

    return AwaitableGetOrgResult(
        id=__ret__.id,
        name=__ret__.name,
        org_id=__ret__.org_id)


@_utilities.lift_output_func(get_org)
def get_org_output(org_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrgResult]:
    """
    Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_zitadel as zitadel

    org_org = zitadel.get_org(org_id="177073608051458051")
    pulumi.export("org", org_org)
    ```


    :param str org_id: The ID of this resource.
    """
    ...