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
    /// Resource representing a GitHub Enterprise IDP on the instance.
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
    ///     var @default = new Zitadel.IdpGithubEs("default", new()
    ///     {
    ///         AuthorizationEndpoint = "https://auth.endpoint",
    ///         ClientId = "86a165...",
    ///         ClientSecret = "*****afdbac18",
    ///         IsAutoCreation = false,
    ///         IsAutoUpdate = true,
    ///         IsCreationAllowed = true,
    ///         IsLinkingAllowed = false,
    ///         Scopes = new[]
    ///         {
    ///             "openid",
    ///             "profile",
    ///             "email",
    ///         },
    ///         TokenEndpoint = "https://token.endpoint",
    ///         UserEndpoint = "https://user.endpoint",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// bash The resource can be imported using the ID format `&lt;id[:client_secret]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/idpGithubEs:IdpGithubEs imported '123456789012345678:1234567890123456781234567890123456787890'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/idpGithubEs:IdpGithubEs")]
    public partial class IdpGithubEs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the providers authorization endpoint
        /// </summary>
        [Output("authorizationEndpoint")]
        public Output<string> AuthorizationEndpoint { get; private set; } = null!;

        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Output("isAutoCreation")]
        public Output<bool> IsAutoCreation { get; private set; } = null!;

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Output("isAutoUpdate")]
        public Output<bool> IsAutoUpdate { get; private set; } = null!;

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Output("isCreationAllowed")]
        public Output<bool> IsCreationAllowed { get; private set; } = null!;

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Output("isLinkingAllowed")]
        public Output<bool> IsLinkingAllowed { get; private set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// the providers token endpoint
        /// </summary>
        [Output("tokenEndpoint")]
        public Output<string> TokenEndpoint { get; private set; } = null!;

        /// <summary>
        /// the providers user endpoint
        /// </summary>
        [Output("userEndpoint")]
        public Output<string> UserEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a IdpGithubEs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdpGithubEs(string name, IdpGithubEsArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/idpGithubEs:IdpGithubEs", name, args ?? new IdpGithubEsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdpGithubEs(string name, Input<string> id, IdpGithubEsState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/idpGithubEs:IdpGithubEs", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdpGithubEs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdpGithubEs Get(string name, Input<string> id, IdpGithubEsState? state = null, CustomResourceOptions? options = null)
        {
            return new IdpGithubEs(name, id, state, options);
        }
    }

    public sealed class IdpGithubEsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the providers authorization endpoint
        /// </summary>
        [Input("authorizationEndpoint", required: true)]
        public Input<string> AuthorizationEndpoint { get; set; } = null!;

        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("clientSecret", required: true)]
        private Input<string>? _clientSecret;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Input("isAutoCreation", required: true)]
        public Input<bool> IsAutoCreation { get; set; } = null!;

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Input("isAutoUpdate", required: true)]
        public Input<bool> IsAutoUpdate { get; set; } = null!;

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Input("isCreationAllowed", required: true)]
        public Input<bool> IsCreationAllowed { get; set; } = null!;

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Input("isLinkingAllowed", required: true)]
        public Input<bool> IsLinkingAllowed { get; set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// the providers token endpoint
        /// </summary>
        [Input("tokenEndpoint", required: true)]
        public Input<string> TokenEndpoint { get; set; } = null!;

        /// <summary>
        /// the providers user endpoint
        /// </summary>
        [Input("userEndpoint", required: true)]
        public Input<string> UserEndpoint { get; set; } = null!;

        public IdpGithubEsArgs()
        {
        }
        public static new IdpGithubEsArgs Empty => new IdpGithubEsArgs();
    }

    public sealed class IdpGithubEsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the providers authorization endpoint
        /// </summary>
        [Input("authorizationEndpoint")]
        public Input<string>? AuthorizationEndpoint { get; set; }

        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Input("isAutoCreation")]
        public Input<bool>? IsAutoCreation { get; set; }

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Input("isAutoUpdate")]
        public Input<bool>? IsAutoUpdate { get; set; }

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Input("isCreationAllowed")]
        public Input<bool>? IsCreationAllowed { get; set; }

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Input("isLinkingAllowed")]
        public Input<bool>? IsLinkingAllowed { get; set; }

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// the providers token endpoint
        /// </summary>
        [Input("tokenEndpoint")]
        public Input<string>? TokenEndpoint { get; set; }

        /// <summary>
        /// the providers user endpoint
        /// </summary>
        [Input("userEndpoint")]
        public Input<string>? UserEndpoint { get; set; }

        public IdpGithubEsState()
        {
        }
        public static new IdpGithubEsState Empty => new IdpGithubEsState();
    }
}
