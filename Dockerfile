FROM golang:1.22-alpine

WORKDIR /app
COPY vikunja-emoji ./

CMD ["/app/vikunja-emoji"]

ENV OPENAI_API_KEY=""
ENV VIKUNJA_SERVICE_TOKEN=""
ENV VIKUNJA_SERVICE_URL=""