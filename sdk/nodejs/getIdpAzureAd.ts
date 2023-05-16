// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing an Azure AD IDP on the instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const azureAd = pulumi.output(zitadel.getIdpAzureAd({
 *     id: "177073614158299139",
 * }));
 * ```
 */
export function getIdpAzureAd(args: GetIdpAzureAdArgs, opts?: pulumi.InvokeOptions): Promise<GetIdpAzureAdResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("zitadel:index/getIdpAzureAd:getIdpAzureAd", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdpAzureAd.
 */
export interface GetIdpAzureAdArgs {
    /**
     * The ID of this resource.
     */
    id: string;
}

/**
 * A collection of values returned by getIdpAzureAd.
 */
export interface GetIdpAzureAdResult {
    /**
     * client id generated by the identity provider
     */
    readonly clientId: string;
    /**
     * client secret generated by the identity provider
     */
    readonly clientSecret: string;
    /**
     * automatically mark emails as verified
     */
    readonly emailVerified: boolean;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * enabled if a new account in ZITADEL are created automatically on login with an external account
     */
    readonly isAutoCreation: boolean;
    /**
     * enabled if a the ZITADEL account fields are updated automatically on each login
     */
    readonly isAutoUpdate: boolean;
    /**
     * enabled if users are able to create a new account in ZITADEL when using an external account
     */
    readonly isCreationAllowed: boolean;
    /**
     * enabled if users are able to link an existing ZITADEL user with an external account
     */
    readonly isLinkingAllowed: boolean;
    /**
     * Name of the IDP
     */
    readonly name: string;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    readonly scopes: string[];
    /**
     * the azure ad tenant id
     */
    readonly tenantId: string;
    /**
     * the azure ad tenant type
     */
    readonly tenantType: string;
}

export function getIdpAzureAdOutput(args: GetIdpAzureAdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdpAzureAdResult> {
    return pulumi.output(args).apply(a => getIdpAzureAd(a, opts))
}

/**
 * A collection of arguments for invoking getIdpAzureAd.
 */
export interface GetIdpAzureAdOutputArgs {
    /**
     * The ID of this resource.
     */
    id: pulumi.Input<string>;
}