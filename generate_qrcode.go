package main

import (
	qrcode "github.com/skip2/go-qrcode"
)

func generateQRCode(ip string) (string, error) {
	if q, err := qrcode.New("http://"+ip+":7684/web", 1); err != nil {
		return "", err
	} else {
		q.DisableBorder = true
		s := q.ToString(true)
		return s, nil
	}
}
