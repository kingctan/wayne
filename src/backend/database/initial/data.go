package initial

// init wayne data
var InitialData = []string{
	// -- config
	`INSERT INTO  config  ( name,  value ) VALUES ('system.api-name-generate-rule', 'join');`,
	`INSERT INTO  config  ( name,  value ) VALUES ('system.oauth2-title', 'OAuth 2.0 Login');`,
	// -- user
	`INSERT INTO  user  ( id, name, email, display, comment, type, deleted, admin, last_login, last_ip, create_time, update_time, password, salt ) VALUES (1,'admin','admin@gmail.com','管理员','',0,0,1,now(),'127.0.0.1',now(),now(),'e7cadd50397b88397045bf1b7f406b34dc8dc6b8f79d470c0a80cf7aad08690748bf5e6c2d0881bb8bb9c96045b08318fa2b','BZoWKqwaQ6');`,
	// -- namespace
	`INSERT INTO  namespace  ( id, name, meta_data, create_time, update_time, user, deleted ) VALUES (1,'demo','{\"namespace\":\"default\"}',now(),now(),'admin',0);`,
	// -- permission
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('46', 'APPUSER_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('47', 'APPUSER_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('48', 'APPUSER_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('49', 'APPUSER_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('50', 'NAMESPACEUSER_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('51', 'NAMESPACEUSER_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('52', 'NAMESPACEUSER_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('53', 'NAMESPACEUSER_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('54', 'DEPLOYMENT_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('55', 'DEPLOYMENT_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('56', 'DEPLOYMENT_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('57', 'DEPLOYMENT_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('58', 'DEPLOYMENT_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('59', 'SERVICE_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('60', 'SERVICE_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('61', 'SERVICE_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('62', 'SERVICE_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('63', 'SERVICE_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('64', 'CONFIGMAP_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('65', 'CONFIGMAP_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('66', 'CONFIGMAP_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('67', 'CONFIGMAP_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('68', 'CONFIGMAP_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('69', 'PVC_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('70', 'PVC_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('71', 'PVC_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('72', 'PVC_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('74', 'APP_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('75', 'APP_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('76', 'APP_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('77', 'APP_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('79', 'SECRET_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('80', 'SECRET_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('81', 'SECRET_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('82', 'SECRET_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('83', 'SECRET_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('84', 'NAMESPACE_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('85', 'NAMESPACE_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('86', 'NAMESPACE_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('87', 'NAMESPACE_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('88', 'APP_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('89', 'PVC_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('90', 'CRONJOB_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('91', 'CRONJOB_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('92', 'CRONJOB_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('93', 'CRONJOB_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('94', 'CRONJOB_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('95', 'WEBHOOK_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('96', 'WEBHOOK_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('97', 'WEBHOOK_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('98', 'WEBHOOK_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('99', 'WEBHOOK_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('100', 'APIKEY_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('101', 'APIKEY_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('102', 'APIKEY_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('103', 'APIKEY_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('104', 'STATEFULSET_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('105', 'STATEFULSET_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('106', 'STATEFULSET_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('107', 'STATEFULSET_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('108', 'STATEFULSET_DEPLOY', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('109', 'DAEMONSET_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('110', 'DAEMONSET_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('111', 'DAEMONSET_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('112', 'DAEMONSET_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('113', 'DAEMONSET_DEPLOY', '', now(), now());`,

	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('114', 'DEPLOYMENT_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('115', 'DAEMONSET_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('116', 'CRONJOB_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('117', 'STATEFULSET_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('118', 'SERVICE_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('119', 'CONFIGMAP_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('120', 'SECRET_OFFLINE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('121', 'PVC_OFFLINE', '', now(), now());`,

	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('122', 'INGRESS_CREATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('123', 'INGRESS_UPDATE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('124', 'INGRESS_READ', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('125', 'INGRESS_DELETE', '', now(), now());`,
	`INSERT INTO  permission  ( id,  name,  comment,  create_time,  update_time ) VALUES ('126', 'INGRESS_DEPLOY', '', now(), now());`,

	// -- group
	// group 名称前加点可以解决group与mysql内置对象重名的问题
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('1', '访客', '访客', '1', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('8', '组管理员', '组管理员', '1', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('9', '组成员', '组成员', '1', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('10', '项目负责人', '项目负责人', '0', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('11', '项目开发', '项目开发', '0', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('12', '项目测试', '项目测试', '0', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('20', '项目运维', '运维相关负责人', '0', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('21', 'API_VIEWER', 'API只读权限', '2', now(), now());`,
	`INSERT INTO  .group  ( id,  name,  comment,  type,  create_time,  update_time ) VALUES ('22', 'API_EDITOR', 'API读写权限', '2', now(), now());`,

	// -- group_permissions
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('846', '10', '46');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('847', '10', '47');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('848', '10', '48');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('849', '10', '49');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('850', '10', '54');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('851', '10', '55');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('852', '10', '56');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('853', '10', '57');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('854', '10', '58');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('855', '10', '59');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('856', '10', '60');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('857', '10', '61');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('858', '10', '62');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('859', '10', '63');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('860', '10', '64');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('861', '10', '65');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('862', '10', '66');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('863', '10', '67');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('864', '10', '68');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('865', '10', '71');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('866', '10', '79');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('867', '10', '80');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('868', '10', '81');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('869', '10', '82');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('870', '10', '83');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('871', '10', '90');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('872', '10', '91');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('873', '10', '92');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('874', '10', '93');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('875', '10', '94');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3876', '10', '95');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3877', '10', '96');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3878', '10', '97');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3879', '10', '98');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3880', '10', '99');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3881', '10', '100');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3882', '10', '101');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3883', '10', '102');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3884', '10', '103');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3885', '10', '122');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3886', '10', '123');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3887', '10', '124');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3888', '10', '125');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('3889', '10', '126');`,

	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('876', '11', '54');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('877', '11', '55');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('878', '11', '56');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('879', '11', '58');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('880', '11', '59');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('881', '11', '60');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('882', '11', '61');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('883', '11', '63');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('884', '11', '64');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('885', '11', '65');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('886', '11', '66');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('887', '11', '68');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('888', '11', '69');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('889', '11', '70');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('890', '11', '71');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('891', '11', '74');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('892', '11', '75');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('893', '11', '76');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('894', '11', '79');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('895', '11', '80');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('896', '11', '81');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('897', '11', '83');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('898', '11', '84');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('899', '11', '85');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('900', '11', '86');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('901', '11', '89');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('902', '11', '91');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('903', '11', '92');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('904', '11', '94');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1905', '11', '96');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1906', '11', '97');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1907', '11', '99');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1908', '11', '100');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1909', '11', '123');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1910', '11', '124');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('1911', '11', '126');`,

	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('905', '12', '48');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('906', '12', '54');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('907', '12', '56');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('908', '12', '58');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('909', '12', '61');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('910', '12', '63');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('911', '12', '66');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('912', '12', '68');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('913', '12', '71');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('914', '12', '76');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('915', '12', '81');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('916', '12', '83');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('917', '12', '92');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('918', '12', '94');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('2917', '12', '124');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('2918', '12', '126');`,

	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('919', '20', '92');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('920', '20', '94');`,

	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('921', '1', '56');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('922', '1', '61');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('923', '1', '66');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('924', '1', '71');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('925', '1', '76');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('926', '1', '81');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('927', '1', '86');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('928', '1', '92');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('929', '8', '46');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('930', '8', '47');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('931', '8', '48');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('932', '8', '49');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('933', '8', '50');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('934', '8', '51');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('935', '8', '52');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('936', '8', '53');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('937', '8', '74');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('938', '8', '75');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('939', '8', '76');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('940', '8', '86');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('941', '8', '97');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('942', '9', '48');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('943', '9', '76');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('944', '9', '86');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('945', '9', '97');`,
	`INSERT INTO  group_permissions  ( id,  group_id,  permission_id ) VALUES ('946', '1', '52');`,
}
