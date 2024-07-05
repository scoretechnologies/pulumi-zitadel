// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing an action belonging to an organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumiverse/zitadel";
 *
 * const _default = new zitadel.Action("default", {
 *     orgId: data.zitadel_org["default"].id,
 *     script: "testscript",
 *     timeout: "10s",
 *     allowedToFail: true,
 * });
 * ```
 *
 * ## Import
 *
 * bash The resource can be imported using the ID format `<id[:org_id]>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/action:Action imported '123456789012345678:123456789012345678'
 * ```
 */
export class Action extends pulumi.CustomResource {
    /**
     * Get an existing Action resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionState, opts?: pulumi.CustomResourceOptions): Action {
        return new Action(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/action:Action';

    /**
     * Returns true if the given object is an instance of Action.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Action {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Action.__pulumiType;
    }

    /**
     * when true, the next action will be called even if this action fails
     */
    public readonly allowedToFail!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the organization
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    public readonly script!: pulumi.Output<string>;
    /**
     * the state of the action
     */
    public /*out*/ readonly state!: pulumi.Output<number>;
    /**
     * after which time the action will be terminated if not finished
     */
    public readonly timeout!: pulumi.Output<string>;

    /**
     * Create a Action resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionArgs | ActionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionState | undefined;
            resourceInputs["allowedToFail"] = state ? state.allowedToFail : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["script"] = state ? state.script : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as ActionArgs | undefined;
            if ((!args || args.allowedToFail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allowedToFail'");
            }
            if ((!args || args.script === undefined) && !opts.urn) {
                throw new Error("Missing required property 'script'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            resourceInputs["allowedToFail"] = args ? args.allowedToFail : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["script"] = args ? args.script : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Action.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Action resources.
 */
export interface ActionState {
    /**
     * when true, the next action will be called even if this action fails
     */
    allowedToFail?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    script?: pulumi.Input<string>;
    /**
     * the state of the action
     */
    state?: pulumi.Input<number>;
    /**
     * after which time the action will be terminated if not finished
     */
    timeout?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Action resource.
 */
export interface ActionArgs {
    /**
     * when true, the next action will be called even if this action fails
     */
    allowedToFail: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    script: pulumi.Input<string>;
    /**
     * after which time the action will be terminated if not finished
     */
    timeout: pulumi.Input<string>;
}
