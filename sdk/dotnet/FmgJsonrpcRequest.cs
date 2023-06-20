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
    /// This resource supports handling JSON RPC request for FortiManager.
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
    ///     var test1 = new Fortios.FmgJsonrpcRequest("test1", new()
    ///     {
    ///         JsonContent = @"{
    ///   ""method"": ""add"",
    ///   ""params"": [
    ///     {
    ///       ""data"": [
    ///         {
    ///           ""action"": ""accept"",
    ///           ""dstaddr"": [""all""],
    ///           ""dstintf"": ""any"",
    ///           ""name"": ""policytest"",
    ///           ""schedule"": ""none"",
    ///           ""service"": ""ALL"",
    ///           ""srcaddr"": ""all"",
    ///           ""srcintf"": ""any"",
    ///           ""internet-service"": ""enable"",
    ///           ""internet-service-id"": ""Alibaba-Web"",
    ///           ""internet-service-src"": ""enable"",
    ///           ""internet-service-src-id"": ""Alibaba-Web"",
    ///           ""users"": ""guest"",
    ///           ""groups"": ""Guest-group""
    ///         }
    ///       ],
    ///       ""url"": ""/pm/config/adom/root/pkg/default/firewall/policy""
    ///     }
    ///   ]
    /// }
    /// 
    /// ",
    ///     });
    /// 
    ///     var test2 = new Fortios.FmgJsonrpcRequest("test2", new()
    ///     {
    ///         JsonContent = @"{
    ///   ""method"": ""add"",
    ///   ""params"": [
    ///     {
    ///       ""data"": [
    ///         {
    ///           ""ip"": ""192.168.1.2"",
    ///           ""name"": ""logserver4"",
    ///           ""port"": ""514""
    ///         }
    ///       ],
    ///       ""url"": ""/cli/global/system/syslog""
    ///     }
    ///   ]
    /// }
    /// 
    /// ",
    ///     });
    /// 
    ///     var test3 = new Fortios.FmgJsonrpcRequest("test3", new()
    ///     {
    ///         JsonContent = @"{
    ///   ""method"": ""get"",
    ///   ""params"": [
    ///     {
    ///       ""url"": ""/cli/global/system/admin/user/APIUser""
    ///     }
    ///   ]
    /// }
    /// 
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/fmgJsonrpcRequest:FmgJsonrpcRequest")]
    public partial class FmgJsonrpcRequest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// JSON RPC request, which should contain 'method' and 'params' parameters.
        /// </summary>
        [Output("jsonContent")]
        public Output<string> JsonContent { get; private set; } = null!;

        /// <summary>
        /// JSON RPC request response data.
        /// </summary>
        [Output("response")]
        public Output<string> Response { get; private set; } = null!;


        /// <summary>
        /// Create a FmgJsonrpcRequest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FmgJsonrpcRequest(string name, FmgJsonrpcRequestArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/fmgJsonrpcRequest:FmgJsonrpcRequest", name, args ?? new FmgJsonrpcRequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FmgJsonrpcRequest(string name, Input<string> id, FmgJsonrpcRequestState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/fmgJsonrpcRequest:FmgJsonrpcRequest", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FmgJsonrpcRequest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FmgJsonrpcRequest Get(string name, Input<string> id, FmgJsonrpcRequestState? state = null, CustomResourceOptions? options = null)
        {
            return new FmgJsonrpcRequest(name, id, state, options);
        }
    }

    public sealed class FmgJsonrpcRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// JSON RPC request, which should contain 'method' and 'params' parameters.
        /// </summary>
        [Input("jsonContent", required: true)]
        public Input<string> JsonContent { get; set; } = null!;

        public FmgJsonrpcRequestArgs()
        {
        }
        public static new FmgJsonrpcRequestArgs Empty => new FmgJsonrpcRequestArgs();
    }

    public sealed class FmgJsonrpcRequestState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// JSON RPC request, which should contain 'method' and 'params' parameters.
        /// </summary>
        [Input("jsonContent")]
        public Input<string>? JsonContent { get; set; }

        /// <summary>
        /// JSON RPC request response data.
        /// </summary>
        [Input("response")]
        public Input<string>? Response { get; set; }

        public FmgJsonrpcRequestState()
        {
        }
        public static new FmgJsonrpcRequestState Empty => new FmgJsonrpcRequestState();
    }
}
