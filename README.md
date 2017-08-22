# captcha

Create captcha by random digits

```
curl -XGET 'http://127.0.0.1:4600/captchas?len=6&width=120&height=40'

{
  "code":"813770",
  "data":"iVBORw0KGgoAAAANSUhEUgAAAHgAAAAoCAMAAAACNM4XAAAAP1BMVEUAAAAMgCNc0HNMwGN98ZQ8sFNYzG8wpEcbjzJSxmly5olXy25v44Y1qUwvo0YThyoRhShv44Z06Ity5olXy277IxW1AAAAAXRSTlMAQObYZgAAAbZJREFUeJzcl1vO6yAMhBkkqlRKBVX2v9ejXADbGNJDaCP9vDTl9s2AcYj5ZQF+isvcRlu4ixsO8vOLAnTyqszj+RxNrm9vaoH3HoO5deySWwFvcC3+3pJb7bksS2rfmZ+DHyX3zchNE0tWBpx3p9yHQq5iX5VZyBJH46cCSm69vF418mb6MA64i3td7hjnWtF1o61c5y6C26OtpWRgd7v+OLLwHRJOR1jRHxGMsm4UV8vNGwFlYH9OPj+QQQFHtwqlfcQ9H9/WpoADjghThpIAkGXy3uc96YjJEELNMAEXs0/T5GMi+ODwK8UGy2zJ7Id0ymGoisnQrNdzCu0WWxGc8u5RyUQYYT5GRx8Z2Q11LBcXNA7I7/9legVdLljBNSVX1nSha8dJ1cBrr+T4PY1AJ/Ocpj52g8vJDHljo9BimM4LYMdhhk4fa7xYgxGWnSNkYhGY0/TrrZCLG7HJjv3LM87znOo8JE4+pLL06sgZK3PZHc0Ydq4Feb9NdoJ5gobcVaNHfST3cjkatZvBvgzjvwrpm6GpbjR4n7t5F/ka9sbv65+Wuzx+cd/OyDdx/1L5FwAA///hjwYtRQHMlQAAAABJRU5ErkJggg=="
}
```

## run 

docker run -d --restart=always -p 4600:4600 vicanso/captcha
