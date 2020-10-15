package architecture
import (
	"fmt"
	"golibs/myutils"
)
import "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
//邮件发送
func SendMail(user string ,password string,to []string,cc []string,bcc []string,subject string,body string,mailtype string)(error)  {
	host := "smtpdm.aliyun.com:25"
	replyToAddress:="lyupeng@189.cn"
    err :=myutils.SendToMail(user, password, host, subject, body, mailtype, replyToAddress, to, cc, bcc)
	return err
}

//短信服务
func SendSMS(phone string,message string,accessKeyId string ,accessSecret string,signName string,templateCode string) (string,error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers =phone
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = message
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return response.Message,err
}
