// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Cos
{
    [TctestResourceType("tctest:Cos/cosBucket:CosBucket")]
    public partial class CosBucket : Pulumi.CustomResource
    {
        /// <summary>
        /// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        /// <summary>
        /// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check
        /// https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
        /// </summary>
        [Output("aclBody")]
        public Output<string?> AclBody { get; private set; } = null!;

        /// <summary>
        /// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// A rule of Cross-Origin Resource Sharing (documented below).
        /// </summary>
        [Output("corsRules")]
        public Output<ImmutableArray<Outputs.CosBucketCorsRule>> CorsRules { get; private set; } = null!;

        /// <summary>
        /// The URL of this cos bucket.
        /// </summary>
        [Output("cosBucketUrl")]
        public Output<string> CosBucketUrl { get; private set; } = null!;

        /// <summary>
        /// The server-side encryption algorithm to use. Valid value is `AES256`.
        /// </summary>
        [Output("encryptionAlgorithm")]
        public Output<string?> EncryptionAlgorithm { get; private set; } = null!;

        /// <summary>
        /// A configuration of object lifecycle management (documented below).
        /// </summary>
        [Output("lifecycleRules")]
        public Output<ImmutableArray<Outputs.CosBucketLifecycleRule>> LifecycleRules { get; private set; } = null!;

        /// <summary>
        /// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be
        /// saved with `log_target_bucket`. To enable log, the full access of log service must be granted. [Full Access Role
        /// Policy](https://intl.cloud.tencent.com/document/product/436/16920).
        /// </summary>
        [Output("logEnable")]
        public Output<bool?> LogEnable { get; private set; } = null!;

        /// <summary>
        /// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file
        /// format is `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is
        /// `true`.
        /// </summary>
        [Output("logPrefix")]
        public Output<string> LogPrefix { get; private set; } = null!;

        /// <summary>
        /// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is
        /// `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is `true`.
        /// User must have full access on this bucket.
        /// </summary>
        [Output("logTargetBucket")]
        public Output<string> LogTargetBucket { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
        /// </summary>
        [Output("multiAz")]
        public Output<bool?> MultiAz { get; private set; } = null!;

        /// <summary>
        /// Bucket Origin Domain settings.
        /// </summary>
        [Output("originDomainRules")]
        public Output<ImmutableArray<Outputs.CosBucketOriginDomainRule>> OriginDomainRules { get; private set; } = null!;

        /// <summary>
        /// Bucket Origin-Pull settings.
        /// </summary>
        [Output("originPullRules")]
        public Output<ImmutableArray<Outputs.CosBucketOriginPullRule>> OriginPullRules { get; private set; } = null!;

        /// <summary>
        /// Request initiator identifier, format: `qcs::cam::uin/&lt;owneruin&gt;:uin/&lt;subuin&gt;`. NOTE: only `versioning_enable` is true
        /// can configure this argument.
        /// </summary>
        [Output("replicaRole")]
        public Output<string?> ReplicaRole { get; private set; } = null!;

        /// <summary>
        /// List of replica rule. NOTE: only `versioning_enable` is true and `replica_role` set can configure this argument.
        /// </summary>
        [Output("replicaRules")]
        public Output<ImmutableArray<Outputs.CosBucketReplicaRule>> ReplicaRules { get; private set; } = null!;

        /// <summary>
        /// The tags of a bucket.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Enable bucket versioning.
        /// </summary>
        [Output("versioningEnable")]
        public Output<bool?> VersioningEnable { get; private set; } = null!;

        /// <summary>
        /// A website object(documented below).
        /// </summary>
        [Output("website")]
        public Output<Outputs.CosBucketWebsite?> Website { get; private set; } = null!;


        /// <summary>
        /// Create a CosBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CosBucket(string name, CosBucketArgs args, CustomResourceOptions? options = null)
            : base("tctest:Cos/cosBucket:CosBucket", name, args ?? new CosBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CosBucket(string name, Input<string> id, CosBucketState? state = null, CustomResourceOptions? options = null)
            : base("tctest:Cos/cosBucket:CosBucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CosBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CosBucket Get(string name, Input<string> id, CosBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new CosBucket(name, id, state, options);
        }
    }

    public sealed class CosBucketArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check
        /// https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
        /// </summary>
        [Input("aclBody")]
        public Input<string>? AclBody { get; set; }

        /// <summary>
        /// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("corsRules")]
        private InputList<Inputs.CosBucketCorsRuleArgs>? _corsRules;

        /// <summary>
        /// A rule of Cross-Origin Resource Sharing (documented below).
        /// </summary>
        public InputList<Inputs.CosBucketCorsRuleArgs> CorsRules
        {
            get => _corsRules ?? (_corsRules = new InputList<Inputs.CosBucketCorsRuleArgs>());
            set => _corsRules = value;
        }

        /// <summary>
        /// The server-side encryption algorithm to use. Valid value is `AES256`.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        [Input("lifecycleRules")]
        private InputList<Inputs.CosBucketLifecycleRuleArgs>? _lifecycleRules;

        /// <summary>
        /// A configuration of object lifecycle management (documented below).
        /// </summary>
        public InputList<Inputs.CosBucketLifecycleRuleArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.CosBucketLifecycleRuleArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be
        /// saved with `log_target_bucket`. To enable log, the full access of log service must be granted. [Full Access Role
        /// Policy](https://intl.cloud.tencent.com/document/product/436/16920).
        /// </summary>
        [Input("logEnable")]
        public Input<bool>? LogEnable { get; set; }

        /// <summary>
        /// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file
        /// format is `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is
        /// `true`.
        /// </summary>
        [Input("logPrefix")]
        public Input<string>? LogPrefix { get; set; }

        /// <summary>
        /// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is
        /// `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is `true`.
        /// User must have full access on this bucket.
        /// </summary>
        [Input("logTargetBucket")]
        public Input<string>? LogTargetBucket { get; set; }

        /// <summary>
        /// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
        /// </summary>
        [Input("multiAz")]
        public Input<bool>? MultiAz { get; set; }

        [Input("originDomainRules")]
        private InputList<Inputs.CosBucketOriginDomainRuleArgs>? _originDomainRules;

        /// <summary>
        /// Bucket Origin Domain settings.
        /// </summary>
        public InputList<Inputs.CosBucketOriginDomainRuleArgs> OriginDomainRules
        {
            get => _originDomainRules ?? (_originDomainRules = new InputList<Inputs.CosBucketOriginDomainRuleArgs>());
            set => _originDomainRules = value;
        }

        [Input("originPullRules")]
        private InputList<Inputs.CosBucketOriginPullRuleArgs>? _originPullRules;

        /// <summary>
        /// Bucket Origin-Pull settings.
        /// </summary>
        public InputList<Inputs.CosBucketOriginPullRuleArgs> OriginPullRules
        {
            get => _originPullRules ?? (_originPullRules = new InputList<Inputs.CosBucketOriginPullRuleArgs>());
            set => _originPullRules = value;
        }

        /// <summary>
        /// Request initiator identifier, format: `qcs::cam::uin/&lt;owneruin&gt;:uin/&lt;subuin&gt;`. NOTE: only `versioning_enable` is true
        /// can configure this argument.
        /// </summary>
        [Input("replicaRole")]
        public Input<string>? ReplicaRole { get; set; }

        [Input("replicaRules")]
        private InputList<Inputs.CosBucketReplicaRuleArgs>? _replicaRules;

        /// <summary>
        /// List of replica rule. NOTE: only `versioning_enable` is true and `replica_role` set can configure this argument.
        /// </summary>
        public InputList<Inputs.CosBucketReplicaRuleArgs> ReplicaRules
        {
            get => _replicaRules ?? (_replicaRules = new InputList<Inputs.CosBucketReplicaRuleArgs>());
            set => _replicaRules = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of a bucket.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Enable bucket versioning.
        /// </summary>
        [Input("versioningEnable")]
        public Input<bool>? VersioningEnable { get; set; }

        /// <summary>
        /// A website object(documented below).
        /// </summary>
        [Input("website")]
        public Input<Inputs.CosBucketWebsiteArgs>? Website { get; set; }

        public CosBucketArgs()
        {
        }
    }

    public sealed class CosBucketState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        /// <summary>
        /// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check
        /// https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
        /// </summary>
        [Input("aclBody")]
        public Input<string>? AclBody { get; set; }

        /// <summary>
        /// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("corsRules")]
        private InputList<Inputs.CosBucketCorsRuleGetArgs>? _corsRules;

        /// <summary>
        /// A rule of Cross-Origin Resource Sharing (documented below).
        /// </summary>
        public InputList<Inputs.CosBucketCorsRuleGetArgs> CorsRules
        {
            get => _corsRules ?? (_corsRules = new InputList<Inputs.CosBucketCorsRuleGetArgs>());
            set => _corsRules = value;
        }

        /// <summary>
        /// The URL of this cos bucket.
        /// </summary>
        [Input("cosBucketUrl")]
        public Input<string>? CosBucketUrl { get; set; }

        /// <summary>
        /// The server-side encryption algorithm to use. Valid value is `AES256`.
        /// </summary>
        [Input("encryptionAlgorithm")]
        public Input<string>? EncryptionAlgorithm { get; set; }

        [Input("lifecycleRules")]
        private InputList<Inputs.CosBucketLifecycleRuleGetArgs>? _lifecycleRules;

        /// <summary>
        /// A configuration of object lifecycle management (documented below).
        /// </summary>
        public InputList<Inputs.CosBucketLifecycleRuleGetArgs> LifecycleRules
        {
            get => _lifecycleRules ?? (_lifecycleRules = new InputList<Inputs.CosBucketLifecycleRuleGetArgs>());
            set => _lifecycleRules = value;
        }

        /// <summary>
        /// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be
        /// saved with `log_target_bucket`. To enable log, the full access of log service must be granted. [Full Access Role
        /// Policy](https://intl.cloud.tencent.com/document/product/436/16920).
        /// </summary>
        [Input("logEnable")]
        public Input<bool>? LogEnable { get; set; }

        /// <summary>
        /// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file
        /// format is `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is
        /// `true`.
        /// </summary>
        [Input("logPrefix")]
        public Input<string>? LogPrefix { get; set; }

        /// <summary>
        /// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is
        /// `log_target_bucket`/`log_prefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `log_enable` is `true`.
        /// User must have full access on this bucket.
        /// </summary>
        [Input("logTargetBucket")]
        public Input<string>? LogTargetBucket { get; set; }

        /// <summary>
        /// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
        /// </summary>
        [Input("multiAz")]
        public Input<bool>? MultiAz { get; set; }

        [Input("originDomainRules")]
        private InputList<Inputs.CosBucketOriginDomainRuleGetArgs>? _originDomainRules;

        /// <summary>
        /// Bucket Origin Domain settings.
        /// </summary>
        public InputList<Inputs.CosBucketOriginDomainRuleGetArgs> OriginDomainRules
        {
            get => _originDomainRules ?? (_originDomainRules = new InputList<Inputs.CosBucketOriginDomainRuleGetArgs>());
            set => _originDomainRules = value;
        }

        [Input("originPullRules")]
        private InputList<Inputs.CosBucketOriginPullRuleGetArgs>? _originPullRules;

        /// <summary>
        /// Bucket Origin-Pull settings.
        /// </summary>
        public InputList<Inputs.CosBucketOriginPullRuleGetArgs> OriginPullRules
        {
            get => _originPullRules ?? (_originPullRules = new InputList<Inputs.CosBucketOriginPullRuleGetArgs>());
            set => _originPullRules = value;
        }

        /// <summary>
        /// Request initiator identifier, format: `qcs::cam::uin/&lt;owneruin&gt;:uin/&lt;subuin&gt;`. NOTE: only `versioning_enable` is true
        /// can configure this argument.
        /// </summary>
        [Input("replicaRole")]
        public Input<string>? ReplicaRole { get; set; }

        [Input("replicaRules")]
        private InputList<Inputs.CosBucketReplicaRuleGetArgs>? _replicaRules;

        /// <summary>
        /// List of replica rule. NOTE: only `versioning_enable` is true and `replica_role` set can configure this argument.
        /// </summary>
        public InputList<Inputs.CosBucketReplicaRuleGetArgs> ReplicaRules
        {
            get => _replicaRules ?? (_replicaRules = new InputList<Inputs.CosBucketReplicaRuleGetArgs>());
            set => _replicaRules = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tags of a bucket.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Enable bucket versioning.
        /// </summary>
        [Input("versioningEnable")]
        public Input<bool>? VersioningEnable { get; set; }

        /// <summary>
        /// A website object(documented below).
        /// </summary>
        [Input("website")]
        public Input<Inputs.CosBucketWebsiteGetArgs>? Website { get; set; }

        public CosBucketState()
        {
        }
    }
}
