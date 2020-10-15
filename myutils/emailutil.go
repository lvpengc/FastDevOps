package myutils

import (
	"net/smtp"
	"strings"
)

func SendToMail(user, password, host, subject, body, mailtype, replyToAddress string, to, cc, bcc []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	cc_address := strings.Join(cc, ";")
	bcc_address := strings.Join(bcc, ";")
	to_address := strings.Join(to, ";")
	msg := []byte("To: " + to_address + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + cc_address + "\r\nBcc: " + bcc_address + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := mergeSlice(to, cc)
	send_to = mergeSlice(send_to, bcc)
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func mergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}