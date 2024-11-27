package otp

import (
	"bytes"
	_ "embed"
	"html/template"
	"io/ioutil"
	"math/rand/v2"
	"path/filepath"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"gopkg.in/gomail.v2"
)

type Email interface {
	SendEmail(to string) error
}

type GoogleEmail struct {
}

func NewEmail() Email {
	return &GoogleEmail{}
}

func (googleEmail *GoogleEmail) SendEmail(to string) error {

	templatePath := filepath.Join("internal", "email", "otp.html")

	// Đọc nội dung file template
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// internal/email/templateFile
	m := gomail.NewMessage()
	m.SetHeader("From", config.Global.GoogleEmail.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "OTP Verification")
	otpCode := rand.Int32N(999999-100000+1) + 100000

	tmpl, err := template.New("otp").Parse(string(templateContent))
	if err != nil {
		return err
	}

	data := struct {
		OTPCode int
	}{
		OTPCode: int(otpCode),
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, config.Global.GoogleEmail.From, config.Global.GoogleEmail.AppPassword)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
