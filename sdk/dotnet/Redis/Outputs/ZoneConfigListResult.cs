// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tctest.Redis.Outputs
{

    [OutputType]
    public sealed class ZoneConfigListResult
    {
        public readonly ImmutableArray<int> MemSizes;
        public readonly ImmutableArray<int> RedisReplicasNums;
        public readonly ImmutableArray<int> RedisShardNums;
        public readonly ImmutableArray<int> ShardMemories;
        public readonly string Type;
        public readonly int TypeId;
        public readonly string Version;
        public readonly string Zone;

        [OutputConstructor]
        private ZoneConfigListResult(
            ImmutableArray<int> memSizes,

            ImmutableArray<int> redisReplicasNums,

            ImmutableArray<int> redisShardNums,

            ImmutableArray<int> shardMemories,

            string type,

            int typeId,

            string version,

            string zone)
        {
            MemSizes = memSizes;
            RedisReplicasNums = redisReplicasNums;
            RedisShardNums = redisShardNums;
            ShardMemories = shardMemories;
            Type = type;
            TypeId = typeId;
            Version = version;
            Zone = zone;
        }
    }
}
