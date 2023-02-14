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





table=${1:?}
db2struct --gorm --no-json -H 10.211.55.3 -d mini -t "$table" --package model --struct "$table"  -u root -p g8azxYLFm2 --target="$table".go