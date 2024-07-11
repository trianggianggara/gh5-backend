############################
# STEP 1 build executable binary
############################
FROM golang:1.22.4-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o gh5-backend

############################
# STEP 2 build a small image
############################
FROM scratch
WORKDIR /app
COPY --from=builder /app/gh5-backend /app
COPY --from=builder /app/.env* /app

CMD ["./gh5-backend"]