#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:curso-go-web .
#docker run -p 8080:8080 -d mercurius:curso-go-web

FROM scratch

ADD curso-go-web /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/curso-go-web" ]