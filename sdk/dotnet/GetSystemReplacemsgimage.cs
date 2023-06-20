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
    public static class GetSystemReplacemsgimage
    {
        /// <summary>
        /// Use this data source to get information on an fortios system replacemsgimage
        /// </summary>
        public static Task<GetSystemReplacemsgimageResult> InvokeAsync(GetSystemReplacemsgimageArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemReplacemsgimageResult>("fortios:index/getSystemReplacemsgimage:getSystemReplacemsgimage", args ?? new GetSystemReplacemsgimageArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios system replacemsgimage
        /// </summary>
        public static Output<GetSystemReplacemsgimageResult> Invoke(GetSystemReplacemsgimageInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemReplacemsgimageResult>("fortios:index/getSystemReplacemsgimage:getSystemReplacemsgimage", args ?? new GetSystemReplacemsgimageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemReplacemsgimageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsgimage.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetSystemReplacemsgimageArgs()
        {
        }
        public static new GetSystemReplacemsgimageArgs Empty => new GetSystemReplacemsgimageArgs();
    }

    public sealed class GetSystemReplacemsgimageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired system replacemsgimage.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetSystemReplacemsgimageInvokeArgs()
        {
        }
        public static new GetSystemReplacemsgimageInvokeArgs Empty => new GetSystemReplacemsgimageInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemReplacemsgimageResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Image data.
        /// </summary>
        public readonly string ImageBase64;
        /// <summary>
        /// Image type.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// Image name.
        /// </summary>
        public readonly string Name;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetSystemReplacemsgimageResult(
            string id,

            string imageBase64,

            string imageType,

            string name,

            string? vdomparam)
        {
            Id = id;
            ImageBase64 = imageBase64;
            ImageType = imageType;
            Name = name;
            Vdomparam = vdomparam;
        }
    }
}
