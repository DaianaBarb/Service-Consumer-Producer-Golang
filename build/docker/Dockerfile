FROM registry.b2w.io/b2wdigital/appsec/base-images/golang1.21-alpine3

ARG ARTIFACT_NAME
ARG TOKEN_EXTERNAL_MATCH
ARG USERNAME=root
USER $USERNAME

WORKDIR /app

COPY $ARTIFACT_NAME /app/main
COPY /config/.stag.env /config/.stag.env
COPY /config/.prod.env /config/.prod.env

RUN chmod +x /app/main

ENTRYPOINT ["./main"]