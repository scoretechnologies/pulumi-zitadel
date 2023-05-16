// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zitadel
{
    /// <summary>
    /// **Caution: Email can only be set verified if a password is set for the user, either with initial_password or during runtime**
    /// 
    /// Resource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Pulumi.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var humanUser = new Zitadel.HumanUser("humanUser", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         UserName = "humanfull@localhost.com",
    ///         FirstName = "firstname",
    ///         LastName = "lastname",
    ///         NickName = "nickname",
    ///         DisplayName = "displayname",
    ///         PreferredLanguage = "de",
    ///         Gender = "GENDER_MALE",
    ///         Phone = "+41799999999",
    ///         IsPhoneVerified = true,
    ///         Email = "test@zitadel.com",
    ///         IsEmailVerified = true,
    ///         InitialPassword = "Password1!",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/humanUser:HumanUser")]
    public partial class HumanUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Display name of the user
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Email of the user
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// First name of the user
        /// </summary>
        [Output("firstName")]
        public Output<string> FirstName { get; private set; } = null!;

        /// <summary>
        /// Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        /// </summary>
        [Output("gender")]
        public Output<string?> Gender { get; private set; } = null!;

        /// <summary>
        /// Initially set password for the user, not changeable after creation
        /// </summary>
        [Output("initialPassword")]
        public Output<string?> InitialPassword { get; private set; } = null!;

        /// <summary>
        /// Is the email verified of the user, can only be true if password of the user is set
        /// </summary>
        [Output("isEmailVerified")]
        public Output<bool?> IsEmailVerified { get; private set; } = null!;

        /// <summary>
        /// Is the phone verified of the user
        /// </summary>
        [Output("isPhoneVerified")]
        public Output<bool?> IsPhoneVerified { get; private set; } = null!;

        /// <summary>
        /// Last name of the user
        /// </summary>
        [Output("lastName")]
        public Output<string> LastName { get; private set; } = null!;

        /// <summary>
        /// Loginnames
        /// </summary>
        [Output("loginNames")]
        public Output<ImmutableArray<string>> LoginNames { get; private set; } = null!;

        /// <summary>
        /// Nick name of the user
        /// </summary>
        [Output("nickName")]
        public Output<string?> NickName { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Phone of the user
        /// </summary>
        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;

        /// <summary>
        /// Preferred language of the user
        /// </summary>
        [Output("preferredLanguage")]
        public Output<string?> PreferredLanguage { get; private set; } = null!;

        /// <summary>
        /// Preferred login name
        /// </summary>
        [Output("preferredLoginName")]
        public Output<string> PreferredLoginName { get; private set; } = null!;

        /// <summary>
        /// State of the user
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Username
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a HumanUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HumanUser(string name, HumanUserArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/humanUser:HumanUser", name, args ?? new HumanUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HumanUser(string name, Input<string> id, HumanUserState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/humanUser:HumanUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HumanUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HumanUser Get(string name, Input<string> id, HumanUserState? state = null, CustomResourceOptions? options = null)
        {
            return new HumanUser(name, id, state, options);
        }
    }

    public sealed class HumanUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name of the user
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Email of the user
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// First name of the user
        /// </summary>
        [Input("firstName", required: true)]
        public Input<string> FirstName { get; set; } = null!;

        /// <summary>
        /// Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        /// </summary>
        [Input("gender")]
        public Input<string>? Gender { get; set; }

        /// <summary>
        /// Initially set password for the user, not changeable after creation
        /// </summary>
        [Input("initialPassword")]
        public Input<string>? InitialPassword { get; set; }

        /// <summary>
        /// Is the email verified of the user, can only be true if password of the user is set
        /// </summary>
        [Input("isEmailVerified")]
        public Input<bool>? IsEmailVerified { get; set; }

        /// <summary>
        /// Is the phone verified of the user
        /// </summary>
        [Input("isPhoneVerified")]
        public Input<bool>? IsPhoneVerified { get; set; }

        /// <summary>
        /// Last name of the user
        /// </summary>
        [Input("lastName", required: true)]
        public Input<string> LastName { get; set; } = null!;

        /// <summary>
        /// Nick name of the user
        /// </summary>
        [Input("nickName")]
        public Input<string>? NickName { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Phone of the user
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        /// <summary>
        /// Preferred language of the user
        /// </summary>
        [Input("preferredLanguage")]
        public Input<string>? PreferredLanguage { get; set; }

        /// <summary>
        /// Username
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public HumanUserArgs()
        {
        }
        public static new HumanUserArgs Empty => new HumanUserArgs();
    }

    public sealed class HumanUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Display name of the user
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Email of the user
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// First name of the user
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Gender of the user, supported values: GENDER*UNSPECIFIED, GENDER*FEMALE, GENDER*MALE, GENDER*DIVERSE
        /// </summary>
        [Input("gender")]
        public Input<string>? Gender { get; set; }

        /// <summary>
        /// Initially set password for the user, not changeable after creation
        /// </summary>
        [Input("initialPassword")]
        public Input<string>? InitialPassword { get; set; }

        /// <summary>
        /// Is the email verified of the user, can only be true if password of the user is set
        /// </summary>
        [Input("isEmailVerified")]
        public Input<bool>? IsEmailVerified { get; set; }

        /// <summary>
        /// Is the phone verified of the user
        /// </summary>
        [Input("isPhoneVerified")]
        public Input<bool>? IsPhoneVerified { get; set; }

        /// <summary>
        /// Last name of the user
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("loginNames")]
        private InputList<string>? _loginNames;

        /// <summary>
        /// Loginnames
        /// </summary>
        public InputList<string> LoginNames
        {
            get => _loginNames ?? (_loginNames = new InputList<string>());
            set => _loginNames = value;
        }

        /// <summary>
        /// Nick name of the user
        /// </summary>
        [Input("nickName")]
        public Input<string>? NickName { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Phone of the user
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        /// <summary>
        /// Preferred language of the user
        /// </summary>
        [Input("preferredLanguage")]
        public Input<string>? PreferredLanguage { get; set; }

        /// <summary>
        /// Preferred login name
        /// </summary>
        [Input("preferredLoginName")]
        public Input<string>? PreferredLoginName { get; set; }

        /// <summary>
        /// State of the user
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Username
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public HumanUserState()
        {
        }
        public static new HumanUserState Empty => new HumanUserState();
    }
}
