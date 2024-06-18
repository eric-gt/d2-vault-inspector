FROM golang:1.22 AS setup
WORKDIR /app
COPY . .

FROM setup AS build
COPY --from=setup /app .
RUN go build -o /dist/d2-vault-inspector cmd/api/main.go

FROM build AS dev
COPY --from=build /app .

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mv tailwindcss-linux-x64 tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 3001
EXPOSE 2345

ENTRYPOINT ["./dev.sh"]

FROM build AS run
COPY --from=build /dist/d2-vault-inspector /dist/d2-vault-inspector
CMD ["/dist/d2-vault-inspector"]
