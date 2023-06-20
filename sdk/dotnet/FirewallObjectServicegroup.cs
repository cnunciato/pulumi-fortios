// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    /// <summary>
    /// Provides a resource to configure firewall service group of FortiOS.
    /// 
    /// !&gt; **Warning:** The resource will be deprecated and replaced by new resource `fortios.FirewallserviceGroup`, we recommend that you use the new resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var v11 = new Fortios.FirewallObjectServicegroup("v11", new()
    ///     {
    ///         Comment = "fdsafdsa",
    ///         Members = new[]
    ///         {
    ///             "DCE-RPC",
    ///             "DNS",
    ///             "HTTPS",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/firewallObjectServicegroup:FirewallObjectServicegroup")]
    public partial class FirewallObjectServicegroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Service objects contained within the group.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Service group name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallObjectServicegroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallObjectServicegroup(string name, FirewallObjectServicegroupArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/firewallObjectServicegroup:FirewallObjectServicegroup", name, args ?? new FirewallObjectServicegroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallObjectServicegroup(string name, Input<string> id, FirewallObjectServicegroupState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallObjectServicegroup:FirewallObjectServicegroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallObjectServicegroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallObjectServicegroup Get(string name, Input<string> id, FirewallObjectServicegroupState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallObjectServicegroup(name, id, state, options);
        }
    }

    public sealed class FirewallObjectServicegroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// Service objects contained within the group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallObjectServicegroupArgs()
        {
        }
        public static new FirewallObjectServicegroupArgs Empty => new FirewallObjectServicegroupArgs();
    }

    public sealed class FirewallObjectServicegroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Service objects contained within the group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Service group name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FirewallObjectServicegroupState()
        {
        }
        public static new FirewallObjectServicegroupState Empty => new FirewallObjectServicegroupState();
    }
}
