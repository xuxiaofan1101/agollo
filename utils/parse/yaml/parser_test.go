package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestParser_ParseAndGetNestedValue 测试解析YAML并获取嵌套值
func TestParser_ParseAndGetNestedValue(t *testing.T) {
	// 初始化解析器
	parser := &Parser{}

	// 定义一个YAML字符串，包含嵌套的配置项
	yamlContent := `
appConfig:
  appName: TestApp
  appVersion: 1.0
serverConfig:
  enableSsl: true
  maxConnections: 100
loggerConfig:
  logLevel: debug
  logOutput: /var/log/test.log
`

	// 调用解析方法
	result, err := parser.Parse(yamlContent)

	// 确保解析没有错误
	assert.NoError(t, err, "解析YAML时出错")

	// 验证 appConfig 的嵌套值
	appName := getNestedValue(result, "appConfig.appName")
	assert.Equal(t, "TestApp", appName, "appName 不匹配")

	appVersion := getNestedValue(result, "appConfig.appVersion")
	assert.Equal(t, 1.0, appVersion, "appVersion 不匹配")

	// 验证 serverConfig 的嵌套值
	enableSsl := getNestedValue(result, "serverConfig.enableSsl")
	assert.Equal(t, true, enableSsl, "enableSsl 不匹配")

	maxConnections := getNestedValue(result, "serverConfig.maxConnections")
	assert.Equal(t, 100, maxConnections, "maxConnections 不匹配")

	// 验证 loggerConfig 的嵌套值
	logLevel := getNestedValue(result, "loggerConfig.logLevel")
	assert.Equal(t, "debug", logLevel, "logLevel 不匹配")

	logOutput := getNestedValue(result, "loggerConfig.logOutput")
	assert.Equal(t, "/var/log/test.log", logOutput, "logOutput 不匹配")
}
