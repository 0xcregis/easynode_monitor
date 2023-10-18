package common

import "time"

/*
*
CREATE TABLE `sys_user` (

	`id` bigint NOT NULL AUTO_INCREMENT,
	`nick_name` varchar(255) DEFAULT NULL,
	`account` varchar(255) NOT NULL,
	`role` int DEFAULT '1' COMMENT '1:admin 2:user',
	`status` int DEFAULT '0' COMMENT '0:禁用 1:可用',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `account` (`account`) USING BTREE

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统用户表';
*/
type SysUser struct {
	ID         int64     `json:"id" gorm:"column:id"`
	NickName   string    `json:"nick_name" gorm:"column:nick_name"`
	Account    string    `json:"account" gorm:"column:account"`
	Password   string    `json:"password" gorm:"column:password"`
	Role       int       `json:"role" gorm:"column:role"`     // 1:admin 2:user
	Status     int       `json:"status" gorm:"column:status"` // 0:禁用 1:可用
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time"`
}
