// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the default lockout policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const lockoutPolicy = new zitadel.DefaultLockoutPolicy("lockout_policy", {
 *     maxPasswordAttempts: 5,
 * });
 * ```
 */
export class DefaultLockoutPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DefaultLockoutPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultLockoutPolicyState, opts?: pulumi.CustomResourceOptions): DefaultLockoutPolicy {
        return new DefaultLockoutPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/defaultLockoutPolicy:DefaultLockoutPolicy';

    /**
     * Returns true if the given object is an instance of DefaultLockoutPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultLockoutPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultLockoutPolicy.__pulumiType;
    }

    /**
     * Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
     */
    public readonly maxPasswordAttempts!: pulumi.Output<number>;

    /**
     * Create a DefaultLockoutPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultLockoutPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultLockoutPolicyArgs | DefaultLockoutPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultLockoutPolicyState | undefined;
            resourceInputs["maxPasswordAttempts"] = state ? state.maxPasswordAttempts : undefined;
        } else {
            const args = argsOrState as DefaultLockoutPolicyArgs | undefined;
            if ((!args || args.maxPasswordAttempts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxPasswordAttempts'");
            }
            resourceInputs["maxPasswordAttempts"] = args ? args.maxPasswordAttempts : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultLockoutPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultLockoutPolicy resources.
 */
export interface DefaultLockoutPolicyState {
    /**
     * Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
     */
    maxPasswordAttempts?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DefaultLockoutPolicy resource.
 */
export interface DefaultLockoutPolicyArgs {
    /**
     * Maximum password check attempts before the account gets locked. Attempts are reset as soon as the password is entered correctly or the password is reset.
     */
    maxPasswordAttempts: pulumi.Input<number>;
}
