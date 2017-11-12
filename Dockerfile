FROM eraac/golang

ADD build/api-train.zip /

WORKDIR /api-train

CMD ["/api-train/bin/api-train", "-profile", "prod", "-config", "/api-train/config/aah.conf"]

