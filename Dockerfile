FROM iron/base
EXPOSE 8080

ADD dist/testbed /

ENTRYPOINT ["./testbed"]
