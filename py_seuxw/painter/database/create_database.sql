create database if not exists `xiaowei_seu`;

-- sf = ServiceFact 服务事实表 - 树洞数据记录
create table if not exists `sf_treehole`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `tree_hole_type` int(2) COMMENT '树洞类型 1-树洞 2-传话 3-酒话 4-建议',
    `from_user_id` int(11) COMMENT '来源用户ID',
    `from_user` varchar(31) NOT NULL COMMENT '来源用户姓名/昵称 用于显示不可空',
    `to_user` varchar(31) COMMENT '指向用户姓名/昵称 类型为2的需要该值',
    `to_user_id` varchar(31) COMMENT '指向用户ID 类型为2的需要该值 该字段暂时不用',
    `content` varchar(511) COMMENT '树洞内容',
    `content_img` varchar(255) COMMENT '树洞图片URL',
    `reply` varchar(256) COMMENT '回复内容',
    `replyer_id` int(11) COMMENT '运营ID 不对外显示',
    `result_uuid` varchar(31) COMMENT '最终生成图片的UUID',

    `send_date` varchar(11) COMMENT '生成日期',
    `reply_date` varchar(11) COMMENT '回复日期',
    `deleted` int DEFAULT 0 COMMENT '删除标记',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- sd - ServiceDimension 服务维度表 - 树洞回复运营信息
create table if not exists `sd_treehole_replyer`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `qq` int(11) NOT NULL COMMENT '运营QQ',
    `username` varchar(31) NOT NULL COMMENT '运营姓名',
    `nickname` varchar(31) NOT NULL COMMENT '运营昵称（对外展示）' ,

    `deleted` int DEFAULT 0 COMMENT '删除标记',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;