// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka
{
    public static class Users
    {
        public static Task<UsersResult> InvokeAsync(UsersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<UsersResult>("tencentcloud:Ckafka/users:Users", args ?? new UsersArgs(), options.WithDefaults());

        public static Output<UsersResult> Invoke(UsersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<UsersResult>("tencentcloud:Ckafka/users:Users", args ?? new UsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class UsersArgs : Pulumi.InvokeArgs
    {
        [Input("accountName")]
        public string? AccountName { get; set; }

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public UsersArgs()
        {
        }
    }

    public sealed class UsersInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public UsersInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class UsersResult
    {
        public readonly string? AccountName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<Outputs.UsersUserListResult> UserLists;

        [OutputConstructor]
        private UsersResult(
            string? accountName,

            string id,

            string instanceId,

            string? resultOutputFile,

            ImmutableArray<Outputs.UsersUserListResult> userLists)
        {
            AccountName = accountName;
            Id = id;
            InstanceId = instanceId;
            ResultOutputFile = resultOutputFile;
            UserLists = userLists;
        }
    }
}