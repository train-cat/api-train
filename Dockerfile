FROM eraac/golang:debug

COPY api-train/ /api-train

WORKDIR /api-train

CMD ["/api-train/bin/api-train", "-profile", "prod", "-config", "/api-train/config/aah.conf"]

