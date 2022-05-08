FROM node:16-alpine3.14 as builder

WORKDIR /app

ARG VITE_API_URL="https://countdown-staging-api.longphung.me"
ENV VITE_API_URL=${VITE_API_URL}
# Enable pnpm
RUN corepack enable

COPY package.json .
COPY pnpm-lock.yaml .
RUN pnpm install --frozen-lockfile

COPY . .
RUN pnpm build

FROM nginx:stable-alpine

RUN apk add --update nodejs npm

COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
# Patch esbuild for Vite
# COPY ./scripts/entrypoint.sh /entrypoint.sh
# RUN chmod +x /entrypoint.sh
# ENTRYPOINT ["/entrypoint.sh"]
