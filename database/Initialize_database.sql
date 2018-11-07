-- Initialize database for project seuxw.cc.

-- 创建数据库
create database if not exists `seuxw`;
use `seuxw`;

-- ---------------------
--  数据表前缀说明
--  sd-     Service Dimension   服务维度表
--  sf-     Service Fact        服务事实表
--
--
--
-- ---------------------

-- ---------------------
--
--  用户项目表
--
-- ---------------------

-- 创建天气预报表 sf_weather
drop table if exists `sf_weather`;
create table `sf_weather` (
    `weather_id` INT AUTO_INCREMENT NOT NULL COMMENT '天气log ID',
    `date`       VARCHAR(20) NOT NULL DEFAULT '' COMMENT '日期',
    `sun_rise_time` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '日出时间',
    `sun_down_time` VARCHAR(10) NOT NULL DEFAULT '' COMMENT '日落时间',
    `pm25_json`     JSON NOT NULL COMMENT 'pm2.5 JSON',
    `life_json`     JSON NOT NULL COMMENT '生活指数 JSON',
    `temp_day`      VARCHAR(4) NOT NULL DEFAULT '' COMMENT '白天温度',
    `temp_night`    VARCHAR(4) NOT NULL DEFAULT '' COMMENT '夜间温度',
    `weather`       VARCHAR(11) NOT NULL DEFAULT '' COMMENT '天气',
    `week`          VARCHAR(4) NOT NULL DEFAULT '' COMMENT '周几',
    
    `deleted`   INT DEFAULT 0 COMMENT '删除标记 0-未删除 1-已删除',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',

    PRIMARY KEY (`weather_id`),
    INDEX (`date`, `deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 创建用户表 sd_user
drop table if exists `sd_user`;
create table `sd_user` (
    `user_id`   INT AUTO_INCREMENT NOT NULL COMMENT '用户 ID 用户唯一标识符',
    `card_id`   INT NOT NULL DEFAULT 0 COMMENT '学生一卡通编号',
    `qq_id`     INT NOT NULL DEFAULT 0 COMMENT '用户绑定 QQ 账号',
    `wechat_id` INT NOT NULL DEFAULT 0 COMMENT '用户绑定微信账号',
    `stu_no`    VARCHAR(15) NOT NULL DEFAULT '' COMMENT '学生学号',
    `real_name` VARCHAR(15) NOT NULL DEFAULT '' COMMENT '用户真实姓名',
    `nick_name` VARCHAR(63) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `gender`    INT NOT NULL DEFAULT 0 COMMENT '用户性别 0-未知 1-男 2-女',
    `user_type` INT NOT NULL DEFAULT 0 COMMENT '用户类别 0-普通 10-VIP 20-管理 30-超级管理',
    `pwd`       VARCHAR(31) NOT NULL DEFAULT '' COMMENT '用户密码',
    `session`   VARCHAR(31) NOT NULL DEFAULT '' COMMENT '用户当前有效 session',
    `mobile`    INT(12) NOT NULL DEFAULT 0 COMMENT '用户手机号码',

    `deleted`   INT DEFAULT 0 COMMENT '删除标记 0-未删除 1-已删除',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 创建一卡通信息表 sd_card
drop table if exists `sd_card`;
create table `sd_card` (
    `card_id`   INT NOT NULL DEFAULT 0 COMMENT '学生一卡通编号',
    `real_name` VARCHAR(15) NOT NULL DEFAULT '' COMMENT '用户真实姓名',
    `stu_no`    VARCHAR(15) NOT NULL DEFAULT '' COMMENT '学生学号',
    `dept_no`   VARCHAR(7) NOT NULL DEFAULT '' COMMENT '学院编号',
    `dept_name` VARCHAR(15) NOT NULL DEFAULT '' COMMENT '学院名称',
    `major_name`VARCHAR(7) NOT NULL DEFAULT '' COMMENT '专业名称',
    `grade`     INT(4) NOT NULL DEFAULT 0 COMMENT '年级',
    `pwd_card`  VARCHAR(31) NOT NULL DEFAULT '' COMMENT '一卡通密码',
    `pwd_money` VARCHAR(31) NOT NULL DEFAULT '' COMMENT '消费密码',
    `gender`    INT NOT NULL DEFAULT 0 COMMENT '用户性别 0-未知 1-男 2-女',

    `deleted`   INT DEFAULT 0 COMMENT '删除标记 0-未删除 1-已删除',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`card_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 创建 QQ 信息表 sd_qq
drop table if exists `sd_qq`;
create table `sd_qq` (
    `qq_id`     INT NOT NULL DEFAULT 0 COMMENT '用户绑定 QQ 账号',
    `nick_name` VARCHAR(63) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `gender`    INT NOT NULL DEFAULT 0 COMMENT '用户性别 0-未知 1-男 2-女',
    `vip`       INT NOT NULL DEFAULT 0 COMMENT 'QQ 会员 0-非会员 1-普通会员 2-超级会员',
    `vip_level` INT NOT NULL DEFAULT 0 COMMENT 'QQ 会员等级',
    `rmk_name`  VARCHAR(63) NOT NULL DEFAULT '' COMMENT '备注名称',
    `hometown`  VARCHAR(15) NOT NULL DEFAULT '' COMMENT 'QQ 家乡区域',
    `address`   VARCHAR(15) NOT NULL DEFAULT '' COMMENT 'QQ 当前区域',
    `birthday`  VARCHAR(15) NOT NULL DEFAULT '' COMMENT 'QQ 出生日',

    `deleted`   INT DEFAULT 0 COMMENT '删除标记 0-未删除 1-已删除',
    `modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近修改时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`qq_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 创建用户脱敏信息视图 v_insensitive_userinfo
drop view if exists `v_insensitive_userinfo`;
create view v_insensitive_userinfo as
select
    u.user_id, u.card_id, u.qq_id, u.wechat_id, u.stu_no,
    u.real_name, u.nick_name, 
    case 
        when c.gender <> 0 then c.gender 
        when q.gender <> 0 then q.gender 
        else 0 
    end as gender,
    u.user_type, c.dept_name, c.major_name, c.grade, q.rmk_name,
    q.address, q.hometown, q.birthday, q.vip, q.vip_level
from
    sd_user as u
inner join sd_card as c on c.card_id = u.card_id and c.deleted = 0
inner join sd_qq as q on q.qq_id = u.qq_id and q.deleted = 0
where
    u.deleted = 0;

-- 创建用户敏感信息视图 v_sensitive_userinfo
drop view if exists `v_sensitive_userinfo`;
create view v_sensitive_userinfo as
select
    u.user_id, u.pwd, u.mobile, u.session, c.pwd_card, c.pwd_money
from
    sd_user as u
inner join sd_card as c on c.card_id = u.card_id and c.deleted = 0
where
    u.deleted = 0;