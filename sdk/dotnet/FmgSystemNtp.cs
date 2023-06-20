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
    /// This resource supports modifying system ntp setting for FortiManager.
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
    ///     var test1 = new Fortios.FmgSystemNtp("test1", new()
    ///     {
    ///         Server = "ntp1.fortinet.com",
    ///         Status = "enable",
    ///         SyncInterval = 30,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [FortiosResourceType("fortios:index/fmgSystemNtp:FmgSystemNtp")]
    public partial class FmgSystemNtp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP address/hostname of NTP Server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Enable/disable NTP.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// NTP sync interval (minute).
        /// </summary>
        [Output("syncInterval")]
        public Output<int?> SyncInterval { get; private set; } = null!;


        /// <summary>
        /// Create a FmgSystemNtp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FmgSystemNtp(string name, FmgSystemNtpArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/fmgSystemNtp:FmgSystemNtp", name, args ?? new FmgSystemNtpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FmgSystemNtp(string name, Input<string> id, FmgSystemNtpState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/fmgSystemNtp:FmgSystemNtp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FmgSystemNtp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FmgSystemNtp Get(string name, Input<string> id, FmgSystemNtpState? state = null, CustomResourceOptions? options = null)
        {
            return new FmgSystemNtp(name, id, state, options);
        }
    }

    public sealed class FmgSystemNtpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address/hostname of NTP Server.
        /// </summary>
        [Input("server", required: true)]
        public Input<string> Server { get; set; } = null!;

        /// <summary>
        /// Enable/disable NTP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NTP sync interval (minute).
        /// </summary>
        [Input("syncInterval")]
        public Input<int>? SyncInterval { get; set; }

        public FmgSystemNtpArgs()
        {
        }
        public static new FmgSystemNtpArgs Empty => new FmgSystemNtpArgs();
    }

    public sealed class FmgSystemNtpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address/hostname of NTP Server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Enable/disable NTP.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// NTP sync interval (minute).
        /// </summary>
        [Input("syncInterval")]
        public Input<int>? SyncInterval { get; set; }

        public FmgSystemNtpState()
        {
        }
        public static new FmgSystemNtpState Empty => new FmgSystemNtpState();
    }
}
