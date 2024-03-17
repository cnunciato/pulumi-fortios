// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Filter.Web
{
    /// <summary>
    /// Configure web filter search engines.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Filter.Web.Searchengine("trname", new()
    ///     {
    ///         Charset = "utf-8",
    ///         Hostname = "sg.eiwuc.com",
    ///         Query = "sc=",
    ///         Safesearch = "disable",
    ///         Url = "^\\/f",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Webfilter SearchEngine can be imported using any of these accepted formats:
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/web/searchengine:Searchengine labelname {{name}}
    /// ```
    /// 
    /// If you do not want to import arguments of block:
    /// 
    /// $ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    /// $ pulumi import fortios:filter/web/searchengine:Searchengine labelname {{name}}
    /// ```
    /// 
    /// $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:filter/web/searchengine:Searchengine")]
    public partial class Searchengine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Search engine charset. Valid values: `utf-8`, `gb2312`.
        /// </summary>
        [Output("charset")]
        public Output<string> Charset { get; private set; } = null!;

        /// <summary>
        /// Hostname (regular expression).
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// Search engine name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Code used to prefix a query (must end with an equals character).
        /// </summary>
        [Output("query")]
        public Output<string> Query { get; private set; } = null!;

        /// <summary>
        /// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        /// </summary>
        [Output("safesearch")]
        public Output<string> Safesearch { get; private set; } = null!;

        /// <summary>
        /// Safe search parameter used in the URL.
        /// </summary>
        [Output("safesearchStr")]
        public Output<string> SafesearchStr { get; private set; } = null!;

        /// <summary>
        /// URL (regular expression).
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Searchengine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Searchengine(string name, SearchengineArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:filter/web/searchengine:Searchengine", name, args ?? new SearchengineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Searchengine(string name, Input<string> id, SearchengineState? state = null, CustomResourceOptions? options = null)
            : base("fortios:filter/web/searchengine:Searchengine", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Searchengine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Searchengine Get(string name, Input<string> id, SearchengineState? state = null, CustomResourceOptions? options = null)
        {
            return new Searchengine(name, id, state, options);
        }
    }

    public sealed class SearchengineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search engine charset. Valid values: `utf-8`, `gb2312`.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// Hostname (regular expression).
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Search engine name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Code used to prefix a query (must end with an equals character).
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        /// </summary>
        [Input("safesearch")]
        public Input<string>? Safesearch { get; set; }

        /// <summary>
        /// Safe search parameter used in the URL.
        /// </summary>
        [Input("safesearchStr")]
        public Input<string>? SafesearchStr { get; set; }

        /// <summary>
        /// URL (regular expression).
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SearchengineArgs()
        {
        }
        public static new SearchengineArgs Empty => new SearchengineArgs();
    }

    public sealed class SearchengineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search engine charset. Valid values: `utf-8`, `gb2312`.
        /// </summary>
        [Input("charset")]
        public Input<string>? Charset { get; set; }

        /// <summary>
        /// Hostname (regular expression).
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Search engine name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Code used to prefix a query (must end with an equals character).
        /// </summary>
        [Input("query")]
        public Input<string>? Query { get; set; }

        /// <summary>
        /// Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header.
        /// </summary>
        [Input("safesearch")]
        public Input<string>? Safesearch { get; set; }

        /// <summary>
        /// Safe search parameter used in the URL.
        /// </summary>
        [Input("safesearchStr")]
        public Input<string>? SafesearchStr { get; set; }

        /// <summary>
        /// URL (regular expression).
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SearchengineState()
        {
        }
        public static new SearchengineState Empty => new SearchengineState();
    }
}