FROM eraac/golang:debug

COPY api-train/ /api-train

WORKDIR /api-train

CMD ["bin/api-train", "-profile", "prod"]

