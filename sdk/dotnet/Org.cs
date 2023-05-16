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
    /// Resource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
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
    ///     var org = new Zitadel.Org("org");
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/org:Org")]
    public partial class Org : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the org
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Org resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Org(string name, OrgArgs? args = null, CustomResourceOptions? options = null)
            : base("zitadel:index/org:Org", name, args ?? new OrgArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Org(string name, Input<string> id, OrgState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/org:Org", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Org resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Org Get(string name, Input<string> id, OrgState? state = null, CustomResourceOptions? options = null)
        {
            return new Org(name, id, state, options);
        }
    }

    public sealed class OrgArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the org
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public OrgArgs()
        {
        }
        public static new OrgArgs Empty => new OrgArgs();
    }

    public sealed class OrgState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the org
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public OrgState()
        {
        }
        public static new OrgState Empty => new OrgState();
    }
}
