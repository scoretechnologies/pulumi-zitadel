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
    public static class GetApplicationApi
    {
        /// <summary>
        /// Datasource representing an API application belonging to a project, with all configuration possibilities.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetApplicationApi.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Default.Id,
        ///         ProjectId = data.Zitadel_project.Default.Id,
        ///         AppId = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationApiResult> InvokeAsync(GetApplicationApiArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationApiResult>("zitadel:index/getApplicationApi:getApplicationApi", args ?? new GetApplicationApiArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an API application belonging to a project, with all configuration possibilities.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetApplicationApi.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Default.Id,
        ///         ProjectId = data.Zitadel_project.Default.Id,
        ///         AppId = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationApiResult> Invoke(GetApplicationApiInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationApiResult>("zitadel:index/getApplicationApi:getApplicationApi", args ?? new GetApplicationApiInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationApiArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public string? OrgId { get; set; }

        /// <summary>
        /// ID of the project
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetApplicationApiArgs()
        {
        }
        public static new GetApplicationApiArgs Empty => new GetApplicationApiArgs();
    }

    public sealed class GetApplicationApiInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

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

        public GetApplicationApiInvokeArgs()
        {
        }
        public static new GetApplicationApiInvokeArgs Empty => new GetApplicationApiInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationApiResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// Auth method type
        /// </summary>
        public readonly string AuthMethodType;
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the application
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of the organization
        /// </summary>
        public readonly string? OrgId;
        /// <summary>
        /// ID of the project
        /// </summary>
        public readonly string ProjectId;

        [OutputConstructor]
        private GetApplicationApiResult(
            string appId,

            string authMethodType,

            string clientId,

            string id,

            string name,

            string? orgId,

            string projectId)
        {
            AppId = appId;
            AuthMethodType = authMethodType;
            ClientId = clientId;
            Id = id;
            Name = name;
            OrgId = orgId;
            ProjectId = projectId;
        }
    }
}
