FROM zkrhm/golang:1.10-alpine-with-glide

ARG app_name
ARG package_name

ENV APP_NAME=${app_name}
ENV PACKAGE=${package_name}

RUN mkdir -p $GOPATH/src/$PACKAGE
WORKDIR $GOPATH/src/$PACKAGE/

COPY . .

RUN glide install

RUN make install
RUN apk del git make
ENTRYPOINT [ ${app_name} serve ]

# EXPOSE 8000