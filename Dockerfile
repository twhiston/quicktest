FROM tomwhiston/micro-golang:test

EXPOSE 8020

USER 1001

ENTRYPOINT ["/go/scripts/hold.sh"]
