FROM alpine

EXPOSE 4600 

ADD ./captcha /

CMD ["/captcha"]