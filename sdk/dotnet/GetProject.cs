// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Zitadel
{
    public static class GetProject
    {
        /// <summary>
        /// Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
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
        ///     var projectProject = Zitadel.GetProject.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Org.Id,
        ///         ProjectId = "177073620768522243",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["project"] = projectProject.Apply(getProjectResult =&gt; getProjectResult),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("zitadel:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing the project, which can then be granted to different organizations or users directly, containing different applications.
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
        ///     var projectProject = Zitadel.GetProject.Invoke(new()
        ///     {
        ///         OrgId = data.Zitadel_org.Org.Id,
        ///         ProjectId = "177073620768522243",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["project"] = projectProject.Apply(getProjectResult =&gt; getProjectResult),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectResult>("zitadel:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization in which the project is located
        /// </summary>
        [Input("orgId", required: true)]
        public string OrgId { get; set; } = null!;

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        public GetProjectArgs()
        {
        }
        public static new GetProjectArgs Empty => new GetProjectArgs();
    }

    public sealed class GetProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization in which the project is located
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public GetProjectInvokeArgs()
        {
        }
        public static new GetProjectInvokeArgs Empty => new GetProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// ZITADEL checks if the org of the user has permission to this project
        /// </summary>
        public readonly bool HasProjectCheck;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the project
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Organization in which the project is located
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// Defines from where the private labeling should be triggered
        /// </summary>
        public readonly string PrivateLabelingSetting;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// describes if roles of user should be added in token
        /// </summary>
        public readonly bool ProjectRoleAssertion;
        /// <summary>
        /// ZITADEL checks if the user has at least one on this project
        /// </summary>
        public readonly bool ProjectRoleCheck;
        /// <summary>
        /// State of the project
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetProjectResult(
            bool hasProjectCheck,

            string id,

            string name,

            string orgId,

            string privateLabelingSetting,

            string projectId,

            bool projectRoleAssertion,

            bool projectRoleCheck,

            string state)
        {
            HasProjectCheck = hasProjectCheck;
            Id = id;
            Name = name;
            OrgId = orgId;
            PrivateLabelingSetting = privateLabelingSetting;
            ProjectId = projectId;
            ProjectRoleAssertion = projectRoleAssertion;
            ProjectRoleCheck = projectRoleCheck;
            State = state;
        }
    }
}