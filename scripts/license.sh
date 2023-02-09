#!/usr/bin/env bash
# Copyright 2023 The Mini Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# 此脚本使用license工具生成LIENSE ,并为代码添加LIENSE声明

license -name "The Mini Authors"  -project "Mini" -o LICENSE apache-2.0 ; addlicense -c "The Mini Authors" -l apache  .
