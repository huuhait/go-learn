package main

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/nats-io/nats.go"
	"gopkg.in/yaml.v3"
)

type Template struct {
	Subject      string `yaml:"subject"`
	TemplatePath string `yaml:"template_path"`
}

type Event struct {
	Name      string              `yaml:"name"`
	Key       string              `yaml:"key"`
	Templates map[string]Template `yaml:"templates"`
}

type Config struct {
	Events []Event `yaml:"events"`
}

var Mails Config

func LoadMailer() {
	b, _ := os.ReadFile("config.yaml")
	yaml.Unmarshal(b, &Mails)
}

func main() {
	////////// EMAIL /////////
	// LoadMailer()

	// msg, err := GetEventTemplate("email_confirmation", "vi")
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(string(msg))

	// SendEmail(msg)

	//////// ENV ////////////////
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// 	os.Getenv("NPM_TOKEN"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"),
	// )

	// log.Println(dsn)

	nc, err := nats.Connect("nats://127.0.0.1:4222")

	log.Println(err)

	nc.Publish("test-gui-di", []byte("toi dep trai"))
	nc.Publish("test-gui-di", []byte("toi dep trai"))
	nc.Publish("test-gui-di", []byte("toi dep trai"))
	nc.Publish("test-gui-di", []byte("toi dep trai"))
	nc.Flush()
}

func GetEventTemplate(key, lang string) ([]byte, error) {
	var template_path string

	for _, event := range Mails.Events {
		if event.Key == key {
			template_path = event.Templates[lang].TemplatePath
		}
	}

	if len(template_path) == 0 {
		return nil, errors.New("Can't find email")
	}

	content_tpl, err := template.ParseFiles(template_path)
	if err != nil {
		return nil, errors.New("Can't parse template")
	}

	content_buf := bytes.Buffer{}
	content_tpl.Execute(&content_buf, map[string]interface{}{
		"code": "132456",
	})

	msg := []byte("\r\n" + content_buf.String())

	return msg, nil
}

func SendEmail(body []byte) {
	auth := smtp.PlainAuth("", "business@zsmart.tech", "rNaUktf8VKLJ", "smtp.zoho.com")
	err := smtp.SendMail("smtp.zoho.com:587", auth, "business@zsmart.tech", []string{"huuhadz2k@gmail.com"}, body)
	if err != nil {
		log.Println(err)
	}
}
