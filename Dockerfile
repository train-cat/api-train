FROM eraac/golang:debug

ADD build/api-train.zip /api-train/

WORKDIR /api-train

CMD ["/api-train/bin/api-train", "-profile", "prod", "-config", "/api-train/config/aah.conf"]

