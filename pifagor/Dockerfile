FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/bgoldovsky/pifagor/app
WORKDIR /go/src/github.com/bgoldovsky/pifagor/app

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	app; \
	fi
	
EXPOSE 8081