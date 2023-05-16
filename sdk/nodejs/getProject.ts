// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const projectProject = zitadel.getProject({
 *     orgId: data.zitadel_org.org.id,
 *     projectId: "177073620768522243",
 * });
 * export const project = projectProject;
 * ```
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("zitadel:index/getProject:getProject", {
        "orgId": args.orgId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * Organization in which the project is located
     */
    orgId: string;
    /**
     * The ID of this resource.
     */
    projectId: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * ZITADEL checks if the org of the user has permission to this project
     */
    readonly hasProjectCheck: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the project
     */
    readonly name: string;
    /**
     * Organization in which the project is located
     */
    readonly orgId: string;
    /**
     * Defines from where the private labeling should be triggered
     */
    readonly privateLabelingSetting: string;
    /**
     * The ID of this resource.
     */
    readonly projectId: string;
    /**
     * describes if roles of user should be added in token
     */
    readonly projectRoleAssertion: boolean;
    /**
     * ZITADEL checks if the user has at least one on this project
     */
    readonly projectRoleCheck: boolean;
    /**
     * State of the project
     */
    readonly state: string;
}

export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply(a => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * Organization in which the project is located
     */
    orgId: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    projectId: pulumi.Input<string>;
}
