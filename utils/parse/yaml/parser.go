/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yaml

import (
	"github.com/xuxiaofan1101/agollo/v4/utils"
	"gopkg.in/yaml.v2"
	"strings"
)

// Parser properties转换器
type Parser struct {
}

func (d *Parser) Parse(configContent interface{}) (map[string]interface{}, error) {
	content, ok := configContent.(string)
	if !ok {
		return nil, nil
	}
	if utils.Empty == content {
		return nil, nil
	}

	var result map[interface{}]interface{}

	// 使用 yaml.v2 进行解析
	err := yaml.Unmarshal([]byte(content), &result)
	if err != nil {
		return nil, err
	}

	// 将 map[interface{}]interface{} 转换为 map[string]interface{}
	return convertToStringKeyMap(result), nil
}

// 递归将 map[interface{}]interface{} 转换为 map[string]interface{}
func convertToStringKeyMap(input map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range input {
		strKey := key.(string) // 假设键是字符串类型
		switch v := value.(type) {
		case map[interface{}]interface{}:
			result[strKey] = convertToStringKeyMap(v)
		default:
			result[strKey] = v
		}
	}
	return result
}

// 递归获取嵌套值的函数
func getNestedValue(data map[string]interface{}, key string) interface{} {
	keys := strings.Split(key, ".")
	var result interface{} = data
	for _, k := range keys {
		if m, ok := result.(map[string]interface{}); ok {
			result = m[k]
		} else {
			return nil
		}
	}
	return result
}
