package clouds

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tat "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat/v20201028"
)

type TencentCloud struct {
	cloud *Cloud
	Client *tat.Client
}


// init initializes the TencentCloud struct
func (c *TencentCloud) init() error {
	credential := common.NewCredential(
		"SecretId",
		"SecretKey",)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tat.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := tat.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tat.NewCancelInvocationRequest()

	// 返回的resp是一个CancelInvocationResponse的实例，与请求对象对应
	response, err := client.CancelInvocation(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return nil
	}
	if err != nil {
			panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
	c.Client = tat.NewClient(credential, "", cpf)
	return nil
}

// close closes the TencentCloud struct
func (c *TencentCloud) close() error {
	c.Client = nil
	return nil
}