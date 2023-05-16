// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the membership of a user on an instance, defined with the given role.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const instanceMember = new zitadel.InstanceMember("instanceMember", {
 *     userId: zitadel_human_user.human_user.id,
 *     roles: ["IAM_OWNER"],
 * });
 * ```
 */
export class InstanceMember extends pulumi.CustomResource {
    /**
     * Get an existing InstanceMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceMemberState, opts?: pulumi.CustomResourceOptions): InstanceMember {
        return new InstanceMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/instanceMember:InstanceMember';

    /**
     * Returns true if the given object is an instance of InstanceMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceMember.__pulumiType;
    }

    /**
     * List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
     */
    public readonly roles!: pulumi.Output<string[]>;
    /**
     * ID of the user
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a InstanceMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceMemberArgs | InstanceMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceMemberState | undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as InstanceMemberArgs | undefined;
            if ((!args || args.roles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roles'");
            }
            if ((!args || args.userId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userId'");
            }
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceMember resources.
 */
export interface InstanceMemberState {
    /**
     * List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the user
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceMember resource.
 */
export interface InstanceMemberArgs {
    /**
     * List of roles granted, full list available here: https://zitadel.com/docs/guides/manage/console/managers#roles
     */
    roles: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the user
     */
    userId: pulumi.Input<string>;
}
