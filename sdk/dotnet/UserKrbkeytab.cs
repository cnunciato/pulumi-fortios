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
    /// Configure Kerberos keytab entries.
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
    ///     var trname2 = new Fortios.UserLdap("trname2", new()
    ///     {
    ///         AccountKeyFilter = "(&amp;(userPrincipalName=%s)(!(UserAccountControl:1.2.840.113556.1.4.803:=2)))",
    ///         AccountKeyProcessing = "same",
    ///         Cnid = "cn",
    ///         Dn = "EIWNCIEW",
    ///         GroupMemberCheck = "user-attr",
    ///         GroupObjectFilter = "(&amp;(objectcategory=group)(member=*))",
    ///         MemberAttr = "memberOf",
    ///         PasswordExpiryWarning = "disable",
    ///         PasswordRenewal = "disable",
    ///         Port = 389,
    ///         Secure = "disable",
    ///         Server = "1.1.1.1",
    ///         ServerIdentityCheck = "disable",
    ///         SourceIp = "0.0.0.0",
    ///         SslMinProtoVersion = "default",
    ///         Type = "simple",
    ///     });
    /// 
    ///     var trname = new Fortios.UserKrbkeytab("trname", new()
    ///     {
    ///         Keytab = "ZXdlY2VxcmVxd3Jld3E=",
    ///         LdapServer = trname2.Name,
    ///         Principal = "testprin",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User KrbKeytab can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userKrbkeytab:UserKrbkeytab labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/userKrbkeytab:UserKrbkeytab labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/userKrbkeytab:UserKrbkeytab")]
    public partial class UserKrbkeytab : global::Pulumi.CustomResource
    {
        /// <summary>
        /// base64 coded keytab file containing a pre-shared key.
        /// </summary>
        [Output("keytab")]
        public Output<string> Keytab { get; private set; } = null!;

        /// <summary>
        /// LDAP server name.
        /// </summary>
        [Output("ldapServer")]
        public Output<string> LdapServer { get; private set; } = null!;

        /// <summary>
        /// Kerberos keytab entry name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("pacData")]
        public Output<string> PacData { get; private set; } = null!;

        /// <summary>
        /// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a UserKrbkeytab resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserKrbkeytab(string name, UserKrbkeytabArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/userKrbkeytab:UserKrbkeytab", name, args ?? new UserKrbkeytabArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserKrbkeytab(string name, Input<string> id, UserKrbkeytabState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/userKrbkeytab:UserKrbkeytab", name, state, MakeResourceOptions(options, id))
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
                    "keytab",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserKrbkeytab resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserKrbkeytab Get(string name, Input<string> id, UserKrbkeytabState? state = null, CustomResourceOptions? options = null)
        {
            return new UserKrbkeytab(name, id, state, options);
        }
    }

    public sealed class UserKrbkeytabArgs : global::Pulumi.ResourceArgs
    {
        [Input("keytab", required: true)]
        private Input<string>? _keytab;

        /// <summary>
        /// base64 coded keytab file containing a pre-shared key.
        /// </summary>
        public Input<string>? Keytab
        {
            get => _keytab;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keytab = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// LDAP server name.
        /// </summary>
        [Input("ldapServer", required: true)]
        public Input<string> LdapServer { get; set; } = null!;

        /// <summary>
        /// Kerberos keytab entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pacData")]
        public Input<string>? PacData { get; set; }

        /// <summary>
        /// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserKrbkeytabArgs()
        {
        }
        public static new UserKrbkeytabArgs Empty => new UserKrbkeytabArgs();
    }

    public sealed class UserKrbkeytabState : global::Pulumi.ResourceArgs
    {
        [Input("keytab")]
        private Input<string>? _keytab;

        /// <summary>
        /// base64 coded keytab file containing a pre-shared key.
        /// </summary>
        public Input<string>? Keytab
        {
            get => _keytab;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keytab = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// LDAP server name.
        /// </summary>
        [Input("ldapServer")]
        public Input<string>? LdapServer { get; set; }

        /// <summary>
        /// Kerberos keytab entry name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enable/disable parsing PAC data in the ticket. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("pacData")]
        public Input<string>? PacData { get; set; }

        /// <summary>
        /// Kerberos service principal, e.g. HTTP/fgt.example.com@EXAMPLE.COM.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public UserKrbkeytabState()
        {
        }
        public static new UserKrbkeytabState Empty => new UserKrbkeytabState();
    }
}