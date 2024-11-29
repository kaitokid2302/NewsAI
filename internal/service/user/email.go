package userservice

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand/v2"
	"path/filepath"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"gopkg.in/gomail.v2"
)

func (s *UserServiceImpl) SendEmail(to string) (int, error) {
	templatePath := filepath.Join("internal", "email", "otp.html")

	// Đọc nội dung file template
	templateContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Print("before before before here")
		return 0, err
	}

	// internal/email/templateFile
	m := gomail.NewMessage()
	m.SetHeader("From", config.Global.GoogleEmail.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "OTP Verification")
	otpCode := rand.Int32N(999999-100000+1) + 100000

	tmpl, err := template.New("otp").Parse(string(templateContent))
	if err != nil {
		fmt.Print("before before here")
		return 0, err
	}

	data := struct {
		OTPCode int
	}{
		OTPCode: int(otpCode),
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		fmt.Print("before here")
		return 0, err
	}
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 465, config.Global.GoogleEmail.From, config.Global.GoogleEmail.AppPassword)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("here")
		return 0, err
	}
	s.SetOTPCode(to, int(otpCode))

	return int(otpCode), nil
}

func (s *UserServiceImpl) ResendOTP(email string) (int, error) {
	if s.userRepo.ExistUser(email) {
		return 0, fmt.Errorf("user already exists")
	}
	user, er := s.GetTempUser(email)
	if er != nil || user.Email == "" {
		return 0, fmt.Errorf("user not found, register again")
	}
	otpCode, er := s.SendEmail(email)
	if er != nil {
		return 0, er
	}
	return otpCode, nil
}
