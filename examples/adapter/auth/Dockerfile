FROM alpine:3.2
RUN apk --no-cache add ca-certificates
WORKDIR /bin/
COPY ./bin/auth .
ENTRYPOINT [ "/bin/auth" ]
CMD [ "44225" ]
EXPOSE 44225