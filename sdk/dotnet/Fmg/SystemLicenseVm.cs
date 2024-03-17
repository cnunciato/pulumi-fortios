// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Fmg
{
    /// <summary>
    /// This resource supports uploading VM license to FortiGate through FortiManager.
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
    ///     var test1 = new Fortios.Fmg.SystemLicenseVm("test1", new()
    ///     {
    ///         FileContent = "XXX",
    ///         Target = "fortigate-test",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [FortiosResourceType("fortios:fmg/systemLicenseVm:SystemLicenseVm")]
    public partial class SystemLicenseVm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ADOM that the target device belongs to. default is 'root'.
        /// </summary>
        [Output("adom")]
        public Output<string?> Adom { get; private set; } = null!;

        /// <summary>
        /// The license file, it needs to be base64 encoded.
        /// </summary>
        [Output("fileContent")]
        public Output<string> FileContent { get; private set; } = null!;

        /// <summary>
        /// Target name, which is managed by FortiManager.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;


        /// <summary>
        /// Create a SystemLicenseVm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SystemLicenseVm(string name, SystemLicenseVmArgs args, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemLicenseVm:SystemLicenseVm", name, args ?? new SystemLicenseVmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SystemLicenseVm(string name, Input<string> id, SystemLicenseVmState? state = null, CustomResourceOptions? options = null)
            : base("fortios:fmg/systemLicenseVm:SystemLicenseVm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SystemLicenseVm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SystemLicenseVm Get(string name, Input<string> id, SystemLicenseVmState? state = null, CustomResourceOptions? options = null)
        {
            return new SystemLicenseVm(name, id, state, options);
        }
    }

    public sealed class SystemLicenseVmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM that the target device belongs to. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// The license file, it needs to be base64 encoded.
        /// </summary>
        [Input("fileContent", required: true)]
        public Input<string> FileContent { get; set; } = null!;

        /// <summary>
        /// Target name, which is managed by FortiManager.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        public SystemLicenseVmArgs()
        {
        }
        public static new SystemLicenseVmArgs Empty => new SystemLicenseVmArgs();
    }

    public sealed class SystemLicenseVmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ADOM that the target device belongs to. default is 'root'.
        /// </summary>
        [Input("adom")]
        public Input<string>? Adom { get; set; }

        /// <summary>
        /// The license file, it needs to be base64 encoded.
        /// </summary>
        [Input("fileContent")]
        public Input<string>? FileContent { get; set; }

        /// <summary>
        /// Target name, which is managed by FortiManager.
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        public SystemLicenseVmState()
        {
        }
        public static new SystemLicenseVmState Empty => new SystemLicenseVmState();
    }
}