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
    /// SSH proxy host public keys.
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
    ///     var trname = new Fortios.FirewallsshHostkey("trname", new()
    ///     {
    ///         Hostname = "testmachine",
    ///         Ip = "1.1.1.1",
    ///         Port = 22,
    ///         Status = "trusted",
    ///         Type = "RSA",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FirewallSsh HostKey can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallsshHostkey:FirewallsshHostkey labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/firewallsshHostkey:FirewallsshHostkey labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/firewallsshHostkey:FirewallsshHostkey")]
    public partial class FirewallsshHostkey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Hostname of the SSH server.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// IP address of the SSH server.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// SSH public key name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        /// </summary>
        [Output("nid")]
        public Output<string> Nid { get; private set; } = null!;

        /// <summary>
        /// Port of the SSH server.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// SSH public key.
        /// </summary>
        [Output("publicKey")]
        public Output<string?> PublicKey { get; private set; } = null!;

        /// <summary>
        /// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        /// </summary>
        [Output("usage")]
        public Output<string> Usage { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallsshHostkey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallsshHostkey(string name, FirewallsshHostkeyArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallsshHostkey:FirewallsshHostkey", name, args ?? new FirewallsshHostkeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallsshHostkey(string name, Input<string> id, FirewallsshHostkeyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/firewallsshHostkey:FirewallsshHostkey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
                AdditionalSecretOutputs =
                {
                    "publicKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallsshHostkey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallsshHostkey Get(string name, Input<string> id, FirewallsshHostkeyState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallsshHostkey(name, id, state, options);
        }
    }

    public sealed class FirewallsshHostkeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of the SSH server.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IP address of the SSH server.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// SSH public key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        /// </summary>
        [Input("nid")]
        public Input<string>? Nid { get; set; }

        /// <summary>
        /// Port of the SSH server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("publicKey")]
        private Input<string>? _publicKey;

        /// <summary>
        /// SSH public key.
        /// </summary>
        public Input<string>? PublicKey
        {
            get => _publicKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _publicKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        /// </summary>
        [Input("usage")]
        public Input<string>? Usage { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallsshHostkeyArgs()
        {
        }
        public static new FirewallsshHostkeyArgs Empty => new FirewallsshHostkeyArgs();
    }

    public sealed class FirewallsshHostkeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of the SSH server.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IP address of the SSH server.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// SSH public key name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
        /// </summary>
        [Input("nid")]
        public Input<string>? Nid { get; set; }

        /// <summary>
        /// Port of the SSH server.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("publicKey")]
        private Input<string>? _publicKey;

        /// <summary>
        /// SSH public key.
        /// </summary>
        public Input<string>? PublicKey
        {
            get => _publicKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _publicKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Set the trust status of the public key. Valid values: `trusted`, `revoked`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
        /// </summary>
        [Input("usage")]
        public Input<string>? Usage { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public FirewallsshHostkeyState()
        {
        }
        public static new FirewallsshHostkeyState Empty => new FirewallsshHostkeyState();
    }
}
