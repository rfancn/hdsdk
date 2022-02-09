package dapr

import (
	"context"
	"encoding/json"
	"github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/hdget/hdsdk/utils"
	"github.com/pkg/errors"
)

const ContentTypeJson = "application/json"

// InvokeService 调用dapr服务
func InvokeService(appId, methodName string, data interface{}) ([]byte, error) {
	var value []byte
	switch t := data.(type) {
	case string:
		value = utils.StringToBytes(t)
	case []byte:
		value = t
	default:
		v, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "marshal invoke data")
		}
		value = v
	}

	daprClient, err := client.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "new dapr client")
	}
	if daprClient == nil {
		return nil, errors.New("dapr client is null, name resolution service may not started, please check it")
	}

	// IMPORTANT: daprClient是全局的连接, 不能关闭
	//defer daprClient.Close()

	content := &client.DataContent{
		ContentType: "application/json",
		Data:        value,
	}

	resp, err := daprClient.InvokeMethodWithContent(context.Background(), appId, methodName, "post", content)
	if err != nil {
		return nil, errors.Wrapf(err, "dapr invoke method with content, app:%s, method: %s, content: %s", appId, methodName, utils.BytesToString(value))
	}

	return resp, nil
}

// InvokeServiceWithClient 需要传入daprClient去调用
func InvokeServiceWithClient(daprClient client.Client, appId, methodName string, data interface{}) ([]byte, error) {
	if daprClient == nil {
		return nil, errors.New("dapr client is null, name resolution service may not started, please check it")
	}

	var value []byte
	switch t := data.(type) {
	case string:
		value = utils.StringToBytes(t)
	case []byte:
		value = t
	default:
		v, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "marshal invoke data")
		}
		value = v
	}

	content := &client.DataContent{
		ContentType: "application/json",
		Data:        value,
	}

	ret, err := daprClient.InvokeMethodWithContent(context.Background(), appId, methodName, "post", content)
	if err != nil {
		return nil, errors.Wrapf(err, "dapr invoke method with content, app:%s, method: %s, content: %s", appId, methodName, utils.BytesToString(value))
	}

	return ret, nil
}

// Reply dapr reply
func Reply(event *common.InvocationEvent, resp interface{}) *common.Content {
	data, err := json.Marshal(resp)
	if err != nil {
		return nil
	}

	return &common.Content{
		ContentType: ContentTypeJson,
		Data:        data,
		DataTypeURL: event.DataTypeURL,
	}
}
