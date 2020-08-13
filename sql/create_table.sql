CREATE TABLE `strategy_bucket` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bucket_name` varchar(31) NOT NULL DEFAULT '' COMMENT '桶的名称',
  `exp_id` int(11) NOT NULL COMMENT '实验id',
  `state` varchar(15) DEFAULT 'OPEN' COMMENT '桶的状态,OPEN表示打开,CLOSE表示关闭',
  `creator` varchar(32) NOT NULL COMMENT '实验的操作人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `description` text COMMENT '描述说明',
  `allocation_percent` varchar(15) DEFAULT '0' COMMENT '桶的配比',
  `is_default` int(11) DEFAULT '0',
  `assignment` varchar(30) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `exp_id` (`exp_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `strategy_experiment` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '实验id',
  `exp_name` varchar(127) NOT NULL DEFAULT '' COMMENT '实验的名称',
  `business_id` int(11) DEFAULT NULL COMMENT '对应业务id',
  `start_time` datetime NOT NULL COMMENT '实验生效开始时间',
  `end_time` datetime NOT NULL COMMENT '实验生效结束时间',
  `exp_version` varchar(11) NOT NULL DEFAULT '1' COMMENT '实验的版本号',
  `state` varchar(15) DEFAULT 'DRAFT' COMMENT '实验的状态,DRAFT表示草稿,RUNNING表示正在运行,PAUSED表示暂停,TERMINATED表示实验已终止',
  `creator` varchar(31) NOT NULL COMMENT '实验的创建人',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `description` text COMMENT '描述说明',
  `salt` varchar(30) DEFAULT '',
  `new_user_only` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;