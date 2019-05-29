package msgstack
import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)
type Key struct {
	Username string `yaml:"user"`
	Id       string `yaml:"id"`
	Secret   string `yaml:"secret"`
}

func (c *Key) getConf() *Key {
	yamlFile, err := ioutil.ReadFile("./debug/cfg.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}


func main() {

}

func addSubDomain(subDomain string, ip string, domain string){
	var k Key
	k.getConf()
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", k.Id, k.Secret)

	request := alidns.CreateAddDomainRecordRequest()

	request.TTL = "600"
	request.Value = ip
	request.Type = "A"
	request.RR = subDomain
	request.DomainName = domain

	response, err := client.AddDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v/n", response)
}