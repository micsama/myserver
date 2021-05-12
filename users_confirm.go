package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //引入sms
)

const etime = "3"

func Rsend(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(99999)
	i = i%8900 + 1000
	mytype := c.Query("type")
	phone := "+86" + c.Query("phone")
	Mysms(phone, strconv.Itoa(i), mytype, c)
}
func Mysms(num string, code string, mytype string, c *gin.Context) (bool, bool) {
	// vcode3 := sha256.Sum256([]byte(code))
	// fmt.Println("hi3")
	// vcode4 := (*string)(unsafe.Pointer(&vcode3))
	// fmt.Println("hi4")
	// c.SetCookie("vcode", *vcode4, 7200, "/", "", false, true)
	c.SetCookie("vcode", code, 7200, "/", "", false, true)
	file, _ := os.Open("mykey.json")
	Body, _ := io.ReadAll(file)
	any := jsoniter.Get(Body, "mail", "secretId")
	secretId := any.ToString()
	any = jsoniter.Get(Body, "mail", "secretKey")
	secretKey := any.ToString()
	credential := common.NewCredential(secretId, secretKey)
	fmt.Println(secretId, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf)
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = common.StringPtrs([]string{num})
	request.TemplateID = common.StringPtr("918173")
	request.Sign = common.StringPtr("无调网络")
	request.TemplateParamSet = common.StringPtrs([]string{mytype, code, etime})
	request.SmsSdkAppid = common.StringPtr("1400504872")
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return false, false
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
	return true, true
}
