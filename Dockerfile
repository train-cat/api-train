FROM eraac/golang:debug

ADD api-train /api-train

WORKDIR /api-train

CMD ["bin/api-train", "-profile", "prod"]

