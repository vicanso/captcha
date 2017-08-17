#!/bin/sh

GOOS=linux go build

docker build -t vicanso/captcha .

rm ./captcha