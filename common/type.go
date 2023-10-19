package common

import (
	"time"

	"github.com/0xcregis/easynode_monitor/common/config"
)

/*
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

/*
*
CREATE TABLE `chain_node` (

	`id` bigint NOT NULL AUTO_INCREMENT,
	`chain_code` varchar(255) DEFAULT NULL COMMENT '区块链码',
	`chain_name` varchar(255) DEFAULT NULL COMMENT '区块链简称',
	`chain_icon` varchar(255) DEFAULT NULL COMMENT '区块链icon',
	`chain_uri` json DEFAULT NULL COMMENT '本链的访问方式，不同访问协议 json 数组',
	`node_uri` json DEFAULT NULL COMMENT 'Fullnode 节点的访问方式，json 数组',
	`node_jwt` varchar(255) DEFAULT NULL COMMENT 'fullnode 节点jwt',
	`latest_blocknumber` varchar(255) DEFAULT NULL COMMENT '最新的区块高度',
	`net_status` int DEFAULT NULL COMMENT '网络状态 0:正常 1: 较慢 2: 不可用',
	`net_duration` int DEFAULT NULL COMMENT '请求耗时，单位毫秒',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
*/
type FullNode struct {
	ID                int64     `json:"id" gorm:"column:id"`
	ChainCode         string    `json:"chain_code" gorm:"column:chain_code"`                 // 区块链码
	ChainName         string    `json:"chain_name" gorm:"column:chain_name"`                 // 区块链简称
	ChainIcon         string    `json:"chain_icon" gorm:"column:chain_icon"`                 // 区块链icon
	ChainUri          string    `json:"chain_uri" gorm:"column:chain_uri"`                   // 本链的访问方式，不同访问协议 json 数组
	NodeUri           string    `json:"node_uri" gorm:"column:node_uri"`                     // Fullnode 节点的访问方式，json 数组
	NodeJwt           string    `json:"node_jwt" gorm:"column:node_jwt"`                     // fullnode 节点jwt
	LatestBlocknumber string    `json:"latest_blocknumber" gorm:"column:latest_blocknumber"` // 最新的区块高度
	NetStatus         int       `json:"net_status" gorm:"column:net_status"`                 // 网络状态 0:正常 1: 较慢 2: 不可用
	NetDuration       int64     `json:"net_duration" gorm:"column:net_duration"`             // 请求耗时，单位毫秒
	CreateTime        time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime        time.Time `json:"update_time" gorm:"column:update_time"`
}

type ChainNode struct {
	Chain    *config.Chain `json:"chain"`
	FullNode []*FullNode   `json:"fullNode"`
}
