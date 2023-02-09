-- Copyright 2023 The Mini Authors
--
-- Licensed under the Apache License, Version 2.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1-centos-docker
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : 10.211.55.3:3306
 Source Schema         : mini

 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001

 Date: 08/02/2023 12:52:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `username` varchar(255) NOT NULL,
                        `postID` varchar(256) NOT NULL,
                        `title` varchar(256) NOT NULL,
                        `content` longtext NOT NULL,
                        `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `postID` (`postID`),
                        KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=141 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of post
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `username` varchar(255) NOT NULL,
                        `password` varchar(255) NOT NULL,
                        `nickname` varchar(30) NOT NULL,
                        `email` varchar(256) NOT NULL,
                        `phone` varchar(16) NOT NULL,
                        `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `username` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
