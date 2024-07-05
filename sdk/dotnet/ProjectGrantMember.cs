// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Zitadel
{
    /// <summary>
    /// Resource representing the membership of a user on an granted project, defined with the given role.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zitadel = Pulumiverse.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Zitadel.ProjectGrantMember("default", new()
    ///     {
    ///         OrgId = data.Zitadel_org.Default.Id,
    ///         ProjectId = data.Zitadel_project.Default.Id,
    ///         UserId = data.Zitadel_human_user.Default.Id,
    ///         GrantId = "123456789012345678",
    ///         Roles = new[]
    ///         {
    ///             "PROJECT_GRANT_OWNER",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// bash The resource can be imported using the ID format `&lt;project_id:grant_id:user_id[:org_id]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/projectGrantMember:ProjectGrantMember imported '123456789012345678:123456789012345678:123456789012345678:123456789012345678'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/projectGrantMember:ProjectGrantMember")]
    public partial class ProjectGrantMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the grant
        /// </summary>
        [Output("grantId")]
        public Output<string> GrantId { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// ID of the project
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// List of roles granted
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        /// <summary>
        /// ID of the user
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectGrantMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectGrantMember(string name, ProjectGrantMemberArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/projectGrantMember:ProjectGrantMember", name, args ?? new ProjectGrantMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectGrantMember(string name, Input<string> id, ProjectGrantMemberState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/projectGrantMember:ProjectGrantMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProjectGrantMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectGrantMember Get(string name, Input<string> id, ProjectGrantMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectGrantMember(name, id, state, options);
        }
    }

    public sealed class ProjectGrantMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the grant
        /// </summary>
        [Input("grantId", required: true)]
        public Input<string> GrantId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public ProjectGrantMemberArgs()
        {
        }
        public static new ProjectGrantMemberArgs Empty => new ProjectGrantMemberArgs();
    }

    public sealed class ProjectGrantMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the grant
        /// </summary>
        [Input("grantId")]
        public Input<string>? GrantId { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// List of roles granted
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        /// <summary>
        /// ID of the user
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public ProjectGrantMemberState()
        {
        }
        public static new ProjectGrantMemberState Empty => new ProjectGrantMemberState();
    }
}
