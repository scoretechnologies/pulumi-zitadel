// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zitadel
{
    public static class GetOrgIdpAzureAd
    {
        /// <summary>
        /// Datasource representing an Azure AD IdP of the organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var azureAd = Zitadel.GetOrgIdpAzureAd.Invoke(new()
        ///     {
        ///         Id = "177073614158299139",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetOrgIdpAzureAdResult> InvokeAsync(GetOrgIdpAzureAdArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrgIdpAzureAdResult>("zitadel:index/getOrgIdpAzureAd:getOrgIdpAzureAd", args ?? new GetOrgIdpAzureAdArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an Azure AD IdP of the organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var azureAd = Zitadel.GetOrgIdpAzureAd.Invoke(new()
        ///     {
        ///         Id = "177073614158299139",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetOrgIdpAzureAdResult> Invoke(GetOrgIdpAzureAdInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrgIdpAzureAdResult>("zitadel:index/getOrgIdpAzureAd:getOrgIdpAzureAd", args ?? new GetOrgIdpAzureAdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgIdpAzureAdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public string OrgId { get; set; } = null!;

        public GetOrgIdpAzureAdArgs()
        {
        }
        public static new GetOrgIdpAzureAdArgs Empty => new GetOrgIdpAzureAdArgs();
    }

    public sealed class GetOrgIdpAzureAdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        public GetOrgIdpAzureAdInvokeArgs()
        {
        }
        public static new GetOrgIdpAzureAdInvokeArgs Empty => new GetOrgIdpAzureAdInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgIdpAzureAdResult
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// automatically mark emails as verified
        /// </summary>
        public readonly bool EmailVerified;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// enabled if a new account in ZITADEL are created automatically on login with an external account
        /// </summary>
        public readonly bool IsAutoCreation;
        /// <summary>
        /// enabled if a the ZITADEL account fields are updated automatically on each login
        /// </summary>
        public readonly bool IsAutoUpdate;
        /// <summary>
        /// enabled if users are able to create a new account in ZITADEL when using an external account
        /// </summary>
        public readonly bool IsCreationAllowed;
        /// <summary>
        /// enabled if users are able to link an existing ZITADEL user with an external account
        /// </summary>
        public readonly bool IsLinkingAllowed;
        /// <summary>
        /// Name of the IDP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public readonly ImmutableArray<string> Scopes;
        /// <summary>
        /// the azure ad tenant id
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// the azure ad tenant type
        /// </summary>
        public readonly string TenantType;

        [OutputConstructor]
        private GetOrgIdpAzureAdResult(
            string clientId,

            string clientSecret,

            bool emailVerified,

            string id,

            bool isAutoCreation,

            bool isAutoUpdate,

            bool isCreationAllowed,

            bool isLinkingAllowed,

            string name,

            string orgId,

            ImmutableArray<string> scopes,

            string tenantId,

            string tenantType)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            EmailVerified = emailVerified;
            Id = id;
            IsAutoCreation = isAutoCreation;
            IsAutoUpdate = isAutoUpdate;
            IsCreationAllowed = isCreationAllowed;
            IsLinkingAllowed = isLinkingAllowed;
            Name = name;
            OrgId = orgId;
            Scopes = scopes;
            TenantId = tenantId;
            TenantType = tenantType;
        }
    }
}