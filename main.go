package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/dchest/captcha"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ImageCaptcha struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")

func genImageCaptcha(length, width, height int) (*ImageCaptcha, error) {
	var buf bytes.Buffer
	digits := captcha.RandomDigits(length)
	image := captcha.NewImage("", digits, width, height)
	s := make([]string, length)
	_, err := image.WriteTo(&buf)
	if err != nil {
		return nil, err
	}
	data := base64.StdEncoding.EncodeToString(buf.Bytes())
	for index, v := range digits {
		s[index] = strconv.Itoa(int(v))
	}
	code := strings.Join(s, "")
	return &ImageCaptcha{code, data}, nil
}

func parseInt(value string, defaultValue int) int {
	if len(value) != 0 {
		v, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return v
	}
	return defaultValue
}

func imageCaptchaServe(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	if req.Header.Get("X-Token") != ACCESS_TOKEN {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "token is invalid"}`))
		return
	}

	query := req.URL.Query()
	length := parseInt(query.Get("len"), 4)
	height := parseInt(query.Get("height"), 40)
	width := parseInt(query.Get("width"), 120)
	data, err := genImageCaptcha(length, width, height)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "get captcha fail"}`))
		return
	}
	jsons, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "json marshal fail"}`))
		return
	}
	w.Write(jsons)
}

func pingServe(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", pingServe)
	http.HandleFunc("/captchas", imageCaptchaServe)
	log.Println("server is at :4600")
	if err := http.ListenAndServe(":4600", nil); err != nil {
		log.Fatal(err)
	}
}
