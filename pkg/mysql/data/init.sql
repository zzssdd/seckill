DROP DATABASE IF EXISTS user;
CREATE DATABASE  user;
USE user;
DROP TABLE IF EXISTS userInfo_1;
CREATE TABLE userInfo_1(
                           id int unsigned primary key  comment '主键',
                           email varchar(20) not null unique comment '邮箱',
                           password varchar(30) not null  comment '密码',
                           UNIQUE (email)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '用户表';
DROP TABLE IF EXISTS userInfo_2;
CREATE TABLE userInfo_2(
                           id int unsigned primary key  comment '主键',
                           email varchar(20) not null unique comment '邮箱',
                           password varchar(30) not null  comment '密码',
                           UNIQUE (email)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '用户表';
DROP TABLE IF EXISTS userInfo_3;
CREATE TABLE userInfo_3(
                           id int unsigned primary key  comment '主键',
                           email varchar(20) not null unique comment '邮箱',
                           password varchar(30) not null  comment '密码',
                           UNIQUE (email)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '用户表';
DROP TABLE IF EXISTS userInfo_4;
CREATE TABLE userInfo_4(
                           id int unsigned primary key  comment '主键',
                           email varchar(20) not null unique comment '邮箱',
                           password varchar(30) not null  comment '密码',
                           UNIQUE (email)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '用户表';
DROP TABLE IF EXISTS userInfo_5;
CREATE TABLE userInfo_5(
                           id int unsigned primary key  comment '主键',
                           email varchar(20) not null unique comment '邮箱',
                           password varchar(30) not null  comment '密码',
                           UNIQUE (email)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment '用户表';


DROP TABLE IF EXISTS user_order_1;
CREATE TABLE user_order_1(
                             id bigint unsigned primary key comment '订单id',
                             uid bigint unsigned default null comment '用户id',
                             pid int unsigned default null comment '商品id',
                             created_at datetime(3) default null,
                             UNIQUE (uid,pid,deleted_at),
                             UNIQUE (pid,uid,deleted_at),
                             KEY (deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment='用户订单表';
DROP TABLE IF EXISTS user_order_2;
CREATE TABLE user_order_2(
                             id bigint unsigned primary key comment '订单id',
                             uid bigint unsigned default null comment '用户id',
                             pid int unsigned default null comment '商品id',
                             created_at bigint unsigned default 0,
                             UNIQUE (uid,pid,deleted_at),
                             UNIQUE (pid,uid,deleted_at),
                             KEY (deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment='用户订单表';
DROP TABLE IF EXISTS user_order_3;
CREATE TABLE user_order_3(
                             id bigint unsigned primary key comment '订单id',
                             uid bigint unsigned default null comment '用户id',
                             pid int unsigned default null comment '商品id',
                             created_at bigint unsigned default 0,
                             UNIQUE (uid,pid,deleted_at),
                             UNIQUE (pid,uid,deleted_at),
                             KEY (deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment='用户订单表';
DROP TABLE IF EXISTS user_order_4;
CREATE TABLE user_order_4(
                             id bigint unsigned primary key comment '订单id',
                             uid bigint unsigned default null comment '用户id',
                             pid int unsigned default null comment '商品id',
                             created_at bigint unsigned default 0,
                             UNIQUE (uid,pid,deleted_at),
                             UNIQUE (pid,uid,deleted_at),
                             KEY (deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment='用户订单表';
DROP TABLE IF EXISTS user_order_5;
CREATE TABLE user_order_5(
                             id bigint unsigned primary key comment '订单id',
                             uid bigint unsigned default null comment '用户id',
                             pid int unsigned default null comment '商品id',
                             created_at bigint unsigned default 0,
                             UNIQUE (uid,pid,deleted_at),
                             UNIQUE (pid,uid,deleted_at),
                             KEY (deleted_at)
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin comment='用户订单表';

DROP DATABASE IF EXISTS product;
CREATE DATABASE  product;
USE product;
DROP TABLE IF EXISTS `productInfo`;
CREATE TABLE `productInfo` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `name` varchar(20) COLLATE utf8mb4_bin NOT NULL COMMENT '商品名称',
  `price` float(8,2) NOT NULL COMMENT '商品价格',
  `pic` varchar(30) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商品图片',
  `des` text COLLATE utf8mb4_bin COMMENT '商品描述',
  `num` int default 0  COMMENT '商品数量',
  `freezeNum` int default 0  COMMENT '冻结商品数量',
  PRIMARY KEY (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='商品表';

DROP DATABASE IF EXISTS orderStatus;
CREATE DATABASE  orderStatus;
USE orderStatus;
DROP TABLE IF EXISTS `orderStatus`;
CREATE TABLE `orderStatus` (
  `id` bigint(20) unsigned NOT NULL COMMENT '订单号',
  `uid` bigint(20) unsigned DEFAULT NULL COMMENT '用户号',
  `pid` int(10) unsigned DEFAULT NULL COMMENT '商品号',
  `status` enum('Created','Doing','Finished') COLLATE utf8mb4_bin DEFAULT 'Created' COMMENT '订单状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='订单状态';



