FROM golang:1.22-alpine

WORKDIR /app
COPY vikunja-emoji ./

CMD ["/app/app"]

ENV OPENAI_API_KEY=""
ENV VIKUNJA_SERVICE_TOKEN=""
ENV VIKUNJA_SERVICE_URL=""