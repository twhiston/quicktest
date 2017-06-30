FROM tomwhiston/micro-golang:test

EXPOSE 8020

ENTRYPOINT ["/go/scripts/hold.sh"]
