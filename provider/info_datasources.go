package tctest

import "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"

var DataSourceInfo = map[string]*tfbridge.DataSourceInfo{
	"tencentcloud_availability_regions":          {Tok: tfbridge.MakeDataSource(mainPkg, "Api", "AvailabilityRegions")},
	"tencentcloud_emr":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Emr", "Instances")},
	"tencentcloud_emr_nodes":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Emr", "Nodes")},
	"tencentcloud_availability_zones":            {Tok: tfbridge.MakeDataSource(mainPkg, "Availability", "Zones")},
	"tencentcloud_availability_zones_by_product": {Tok: tfbridge.MakeDataSource(mainPkg, "Availability", "ZonesByProduct")},
	"tencentcloud_instances":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Instances", "Instances")},
	"tencentcloud_reserved_instances":            {Tok: tfbridge.MakeDataSource(mainPkg, "Reserved", "Instances")},
	"tencentcloud_placement_groups":              {Tok: tfbridge.MakeDataSource(mainPkg, "Placement", "Groups")},
	"tencentcloud_key_pairs":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Key", "Pairs")},
	//"tencentcloud_image":                                    {Tok: tfbridge.MakeDataSource(mainPkg, "cvm", "Image")},
	"tencentcloud_images":                                   {Tok: tfbridge.MakeDataSource(mainPkg, "Cvm", "Images")},
	"tencentcloud_instance_types":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Cvm", "InstancesTypes")},
	"tencentcloud_reserved_instance_configs":                {Tok: tfbridge.MakeDataSource(mainPkg, "Cvm", "ReservedInstanceConfigs")},
	"tencentcloud_vpc_instances":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Vpc", "VpcInstances")},
	"tencentcloud_vpc_subnets":                              {Tok: tfbridge.MakeDataSource(mainPkg, "Vpc", "Subnets")},
	"tencentcloud_vpc_route_tables":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Vpc", "RouteTables")},
	"tencentcloud_vpc":                                      {Tok: tfbridge.MakeDataSource(mainPkg, "Vpc", "Instances")},
	"tencentcloud_vpc_acls":                                 {Tok: tfbridge.MakeDataSource(mainPkg, "Vpc", "Acls")},
	"tencentcloud_subnet":                                   {Tok: tfbridge.MakeDataSource(mainPkg, "Subnet", "Instances")},
	"tencentcloud_route_table":                              {Tok: tfbridge.MakeDataSource(mainPkg, "Route", "Table")},
	"tencentcloud_eip":                                      {Tok: tfbridge.MakeDataSource(mainPkg, "Eip", "Instances")},
	"tencentcloud_eips":                                     {Tok: tfbridge.MakeDataSource(mainPkg, "Eips", "Instances")},
	"tencentcloud_enis":                                     {Tok: tfbridge.MakeDataSource(mainPkg, "Enis", "Instances")},
	"tencentcloud_nats":                                     {Tok: tfbridge.MakeDataSource(mainPkg, "Nats", "Instances")},
	"tencentcloud_dnats":                                    {Tok: tfbridge.MakeDataSource(mainPkg, "Dnats", "Instances")},
	"tencentcloud_nat_gateways":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Nat", "Gateways")},
	"tencentcloud_nat_gateway_snats":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Nat", "GatewaySnats")},
	"tencentcloud_vpn_customer_gateways":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Vpn", "CustomerGateways")},
	"tencentcloud_vpn_gateways":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Vpn", "Gateways")},
	"tencentcloud_vpn_gateway_routes":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Vpn", "GatewayRoutes")},
	"tencentcloud_vpn_connections":                          {Tok: tfbridge.MakeDataSource(mainPkg, "Vpn", "Connections")},
	"tencentcloud_ha_vips":                                  {Tok: tfbridge.MakeDataSource(mainPkg, "Ha", "Vips")},
	"tencentcloud_ha_vip_eip_attachments":                   {Tok: tfbridge.MakeDataSource(mainPkg, "Ha", "VipEipAttachments")},
	"tencentcloud_ccn_instances":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Ccn", "Instances")},
	"tencentcloud_ccn_bandwidth_limits":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Ccn", "BandwidthLimits")},
	"tencentcloud_dc_instances":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Dc", "Instances")},
	"tencentcloud_dcx_instances":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Dcx", "Instances")},
	"tencentcloud_dc_gateway_instances":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Dc", "GatewayInstances")},
	"tencentcloud_dc_gateway_ccn_routes":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Dc", "GatewayCCNRoutes")},
	//"tencentcloud_security_group":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Security", "Group")},
	"tencentcloud_security_groups":                          {Tok: tfbridge.MakeDataSource(mainPkg, "Security", "Groups")},
	"tencentcloud_kubernetes_clusters":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Tke", "Clusters")},
	"tencentcloud_kubernetes_charts":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Tke", "Charts")},
	"tencentcloud_kubernetes_cluster_levels":                {Tok: tfbridge.MakeDataSource(mainPkg, "Tke", "ClusterLevels")},
	"tencentcloud_kubernetes_cluster_common_names":          {Tok: tfbridge.MakeDataSource(mainPkg, "Tke", "ClusterCommonNames")},
	"tencentcloud_eks_clusters":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Eks", "Clusters")},
	"tencentcloud_eks_cluster_credential":                   {Tok: tfbridge.MakeDataSource(mainPkg, "Eks", "ClusterCredential")},
	"tencentcloud_container_clusters":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Container", "Clusters")},
	"tencentcloud_container_cluster_instances":              {Tok: tfbridge.MakeDataSource(mainPkg, "Container", "ClusterInstances")},
	"tencentcloud_mysql_backup_list":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Mysql", "BackupList")},
	"tencentcloud_mysql_zone_config":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Mysql", "ZoneConfig")},
	"tencentcloud_mysql_parameter_list":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Mysql", "ParameterList")},
	"tencentcloud_mysql_default_params":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Mysql", "DefaultParams")},
	"tencentcloud_mysql_instance":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Mysql", "Instances")},
	"tencentcloud_cos_bucket_object":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Cos", "BucketObjects")},
	"tencentcloud_cos_buckets":                              {Tok: tfbridge.MakeDataSource(mainPkg, "Cos", "Buckets")},
	"tencentcloud_cfs_file_systems":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Cfs", "FileSystems")},
	"tencentcloud_cfs_access_groups":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Cfs", "AccessGroups")},
	"tencentcloud_cfs_access_rules":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Cfs", "AccessRules")},
	"tencentcloud_redis_zone_config":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Redis", "ZoneConfig")},
	"tencentcloud_redis_instances":                          {Tok: tfbridge.MakeDataSource(mainPkg, "Redis", "Instances")},
	"tencentcloud_as_scaling_configs":                       {Tok: tfbridge.MakeDataSource(mainPkg, "As", "ScalingConfigs")},
	"tencentcloud_as_scaling_groups":                        {Tok: tfbridge.MakeDataSource(mainPkg, "As", "ScalingGroups")},
	"tencentcloud_as_scaling_policies":                      {Tok: tfbridge.MakeDataSource(mainPkg, "As", "ScalingPolicies")},
	"tencentcloud_cbs_storages":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Cbs", "Storages")},
	"tencentcloud_cbs_snapshots":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Cbs", "Snapshots")},
	"tencentcloud_cbs_snapshot_policies":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Cbs", "SnapshotPolicies")},
	"tencentcloud_clb_instances":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Clb", "Instances")},
	"tencentcloud_clb_listeners":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Clb", "Listeners")},
	"tencentcloud_clb_listener_rules":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Clb", "ListenerRules")},
	"tencentcloud_clb_attachments":                          {Tok: tfbridge.MakeDataSource(mainPkg, "Clb", "ServerAttachments")},
	"tencentcloud_clb_redirections":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Clb", "Redirections")},
	"tencentcloud_clb_target_groups":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Clb", "TargetGroups")},
	"tencentcloud_mongodb_zone_config":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Mongodb", "ZoneConfig")},
	"tencentcloud_mongodb_instances":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Mongodb", "Instances")},
	"tencentcloud_dayu_cc_https_policies":                   {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "CCHttpsPolicies")},
	"tencentcloud_dayu_cc_http_policies":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "CCHttpPolicies")},
	"tencentcloud_dayu_ddos_policies":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "DdosPolicies")},
	"tencentcloud_dayu_ddos_policy_cases":                   {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "DdosPolicyCases")},
	"tencentcloud_dayu_ddos_policy_attachments":             {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "DdosPolicyAttachments")},
	"tencentcloud_dayu_l4_rules":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "L4Rules")},
	"tencentcloud_dayu_l4_rules_v2":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "L4RulesV2")},
	"tencentcloud_dayu_l7_rules":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "L7Rules")},
	"tencentcloud_dayu_l7_rules_v2":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "L7RulesV2")},
	"tencentcloud_gaap_proxies":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "Proxies")},
	"tencentcloud_gaap_realservers":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "Realservers")},
	"tencentcloud_gaap_layer4_listeners":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "Layer4Listeners")},
	"tencentcloud_gaap_layer7_listeners":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "Layer7Listeners")},
	"tencentcloud_gaap_http_domains":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "HttpDomains")},
	"tencentcloud_gaap_http_rules":                          {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "HttpRules")},
	"tencentcloud_gaap_security_policies":                   {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "SecurityPolices")},
	"tencentcloud_gaap_security_rules":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "SecurityRules")},
	"tencentcloud_gaap_certificates":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "Certificates")},
	"tencentcloud_gaap_domain_error_pages":                  {Tok: tfbridge.MakeDataSource(mainPkg, "Gaap", "DomainErrorPageInfoList")},
	"tencentcloud_ssl_certificates":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Ssl", "Certificates")},
	"tencentcloud_cam_roles":                                {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "Roles")},
	"tencentcloud_cam_users":                                {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "Users")},
	"tencentcloud_cam_groups":                               {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "Groups")},
	"tencentcloud_cam_group_memberships":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "GroupMemberships")},
	"tencentcloud_cam_policies":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "Policies")},
	"tencentcloud_cam_role_policy_attachments":              {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "RolePolicyAttachments")},
	"tencentcloud_cam_user_policy_attachments":              {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "UserPolicyAttachments")},
	"tencentcloud_cam_group_policy_attachments":             {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "GroupPolicyAttachments")},
	"tencentcloud_cam_saml_providers":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Cam", "SAMLProviders")},
	"tencentcloud_user_info":                                {Tok: tfbridge.MakeDataSource(mainPkg, "User", "Info")},
	"tencentcloud_cdn_domains":                              {Tok: tfbridge.MakeDataSource(mainPkg, "Cdn", "Domains")},
	"tencentcloud_scf_functions":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Scf", "Functions")},
	"tencentcloud_scf_namespaces":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Scf", "Namespaces")},
	"tencentcloud_scf_logs":                                 {Tok: tfbridge.MakeDataSource(mainPkg, "Scf", "Logs")},
	"tencentcloud_tcaplus_clusters":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Tcaplus", "Clusters")},
	"tencentcloud_tcaplus_tablegroups":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Tcaplus", "TableGroups")},
	"tencentcloud_tcaplus_tables":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Tcaplus", "Tables")},
	"tencentcloud_tcaplus_idls":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Tcaplus", "Idls")},
	"tencentcloud_monitor_policy_conditions":                {Tok: tfbridge.MakeDataSource(mainPkg, "Monitor", "PolicyConditions")},
	"tencentcloud_monitor_data":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Monitor", "Data")},
	"tencentcloud_monitor_product_event":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Monitor", "ProductEvent")},
	"tencentcloud_monitor_binding_objects":                  {Tok: tfbridge.MakeDataSource(mainPkg, "Monitor", "BindingObjects")},
	"tencentcloud_monitor_policy_groups":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Monitor", "PolicyGroups")},
	"tencentcloud_monitor_product_namespace":                {Tok: tfbridge.MakeDataSource(mainPkg, "Monitor", "ProductNamespace")},
	"tencentcloud_elasticsearch_instances":                  {Tok: tfbridge.MakeDataSource(mainPkg, "Elasticsearch", "Instances")},
	"tencentcloud_postgresql_instances":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Postgresql", "Instances")},
	"tencentcloud_postgresql_specinfos":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Postgresql", "Specinfos")},
	"tencentcloud_postgresql_xlogs":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Postgresql", "Xlogs")},
	"tencentcloud_sqlserver_zone_config":                    {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "ZoneConfig")},
	"tencentcloud_sqlserver_instances":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "Instances")},
	"tencentcloud_sqlserver_backups":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "Backups")},
	"tencentcloud_sqlserver_dbs":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "DBs")},
	"tencentcloud_sqlserver_accounts":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "Accounts")},
	"tencentcloud_sqlserver_account_db_attachments":         {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "AccountDBAttachments")},
	"tencentcloud_sqlserver_readonly_groups":                {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "ReadonlyGroups")},
	"tencentcloud_ckafka_users":                             {Tok: tfbridge.MakeDataSource(mainPkg, "Ckafka", "Users")},
	"tencentcloud_ckafka_acls":                              {Tok: tfbridge.MakeDataSource(mainPkg, "Ckafka", "Acls")},
	"tencentcloud_ckafka_topics":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Ckafka", "Topics")},
	"tencentcloud_audit_cos_regions":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Audit", "CosRegions")},
	"tencentcloud_audit_key_alias":                          {Tok: tfbridge.MakeDataSource(mainPkg, "Audit", "KeyAlias")},
	"tencentcloud_audits":                                   {Tok: tfbridge.MakeDataSource(mainPkg, "Audits", "Instances")},
	"tencentcloud_cynosdb_clusters":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Cynosdb", "Clusters")},
	"tencentcloud_cynosdb_instances":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Cynosdb", "Instances")},
	"tencentcloud_vod_adaptive_dynamic_streaming_templates": {Tok: tfbridge.MakeDataSource(mainPkg, "Vod", "AdaptiveDynamicStreamingTemplates")},
	"tencentcloud_vod_image_sprite_templates":               {Tok: tfbridge.MakeDataSource(mainPkg, "Vod", "ImageSpriteTemplates")},
	"tencentcloud_vod_procedure_templates":                  {Tok: tfbridge.MakeDataSource(mainPkg, "Vod", "ProcedureTemplates")},
	"tencentcloud_vod_snapshot_by_time_offset_templates":    {Tok: tfbridge.MakeDataSource(mainPkg, "Vod", "SnapshotByTimeOffsetTemplates")},
	"tencentcloud_vod_super_player_configs":                 {Tok: tfbridge.MakeDataSource(mainPkg, "Vod", "SuperPlayerConfigs")},
	"tencentcloud_sqlserver_publish_subscribes":             {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "PublishSubscribes")},
	"tencentcloud_api_gateway_usage_plans":                  {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayUsagePlans")},
	"tencentcloud_api_gateway_ip_strategies":                {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayIpStrategy")},
	"tencentcloud_api_gateway_customer_domains":             {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayCustomerDomains")},
	"tencentcloud_api_gateway_usage_plan_environments":      {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayUsagePlanEnvironments")},
	"tencentcloud_api_gateway_throttling_services":          {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayThrottlingServices")},
	"tencentcloud_api_gateway_throttling_apis":              {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayThrottlingApis")},
	"tencentcloud_api_gateway_apis":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayAPIs")},
	"tencentcloud_api_gateway_services":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayServices")},
	"tencentcloud_api_gateway_api_keys":                     {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "APIGatewayAPIKeys")},
	"tencentcloud_sqlserver_basic_instances":                {Tok: tfbridge.MakeDataSource(mainPkg, "Sqlserver", "BasicInstances")},
	"tencentcloud_tcr_instances":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "TCRInstances")},
	"tencentcloud_tcr_namespaces":                           {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "TCRNamespaces")},
	"tencentcloud_tcr_tokens":                               {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "TCRTokens")},
	"tencentcloud_tcr_vpc_attachments":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "TCRVPCAttachments")},
	"tencentcloud_tcr_repositories":                         {Tok: tfbridge.MakeDataSource(mainPkg, "Cloud", "TCRRepositories")},
	"tencentcloud_address_templates":                        {Tok: tfbridge.MakeDataSource(mainPkg, "Address", "Templates")},
	"tencentcloud_address_template_groups":                  {Tok: tfbridge.MakeDataSource(mainPkg, "Address", "TemplateGroups")},
	"tencentcloud_protocol_templates":                       {Tok: tfbridge.MakeDataSource(mainPkg, "Protocol", "Templates")},
	"tencentcloud_protocol_template_groups":                 {Tok: tfbridge.MakeDataSource(mainPkg, "Protocol", "TemplateGroups")},
	"tencentcloud_kms_keys":                                 {Tok: tfbridge.MakeDataSource(mainPkg, "Kms", "Keys")},
	"tencentcloud_ssm_secrets":                              {Tok: tfbridge.MakeDataSource(mainPkg, "Ssm", "Secrets")},
	"tencentcloud_ssm_secret_versions":                      {Tok: tfbridge.MakeDataSource(mainPkg, "Ssm", "SecretVersions")},
	"tencentcloud_cdh_instances":                            {Tok: tfbridge.MakeDataSource(mainPkg, "Cdh", "Instances")},
	//"tencentcloud_dayu_eip":                                 {Tok: tfbridge.MakeDataSource(mainPkg, "Dayu", "Eip")},
}
