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
    /// Configure device categories. Applies to FortiOS Version `&lt;= 6.2.0`.
    /// 
    /// ## Import
    /// 
    /// User DeviceCategory can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userDevicecategory:UserDevicecategory labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userDevicecategory:UserDevicecategory labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/userDevicecategory:UserDevicecategory")]
    public partial class UserDevicecategory : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// Device category description.
        /// </summary>
        [Output("desc")]
        public Output<string?> Desc { get; private set; } = null!;

        /// <summary>
        /// Device category name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a UserDevicecategory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserDevicecategory(string name, UserDevicecategoryArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/userDevicecategory:UserDevicecategory", name, args ?? new UserDevicecategoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserDevicecategory(string name, Input<string> id, UserDevicecategoryState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/userDevicecategory:UserDevicecategory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserDevicecategory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserDevicecategory Get(string name, Input<string> id, UserDevicecategoryState? state = null, CustomResourceOptions? options = null)
        {
            return new UserDevicecategory(name, id, state, options);
        }
    }

    public sealed class UserDevicecategoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Device category description.
        /// </summary>
        [Input("desc")]
        public Input<string>? Desc { get; set; }

        /// <summary>
        /// Device category name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserDevicecategoryArgs()
        {
        }
        public static new UserDevicecategoryArgs Empty => new UserDevicecategoryArgs();
    }

    public sealed class UserDevicecategoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Device category description.
        /// </summary>
        [Input("desc")]
        public Input<string>? Desc { get; set; }

        /// <summary>
        /// Device category name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserDevicecategoryState()
        {
        }
        public static new UserDevicecategoryState Empty => new UserDevicecategoryState();
    }
}
