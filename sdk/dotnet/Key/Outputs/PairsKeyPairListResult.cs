// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Key.Outputs
{

    [OutputType]
    public sealed class PairsKeyPairListResult
    {
        public readonly string CreateTime;
        public readonly string KeyId;
        public readonly string KeyName;
        public readonly int ProjectId;
        public readonly string PublicKey;

        [OutputConstructor]
        private PairsKeyPairListResult(
            string createTime,

            string keyId,

            string keyName,

            int projectId,

            string publicKey)
        {
            CreateTime = createTime;
            KeyId = keyId;
            KeyName = keyName;
            ProjectId = projectId;
            PublicKey = publicKey;
        }
    }
}
