FROM node:20-alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app

FROM base AS prod-deps
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod --frozen-lockfile

FROM base AS build

ARG VITE_ADMIN_API_URL
ENV VITE_ADMIN_API_URL=$VITE_ADMIN_API_URL
ARG VITE_KEYCLOAK_URL
ENV VITE_KEYCLOAK_URL=$VITE_KEYCLOAK_URL
ARG VITE_KEYCLOAK_REALM
ENV VITE_KEYCLOAK_REALM=$VITE_KEYCLOAK_REALM
ARG VITE_KEYCLOAK_CLIENTID
ENV VITE_KEYCLOAK_CLIENTID=$VITE_KEYCLOAK_CLIENTID
ARG VITE_BASE_ROUTE
ENV VITE_BASE_ROUTE=$VITE_BASE_ROUTE

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

FROM nginx:stable

RUN rm -rf /usr/share/nginx/html/*

# COPY --from=build /app/dist  /usr/share/nginx/html/admin-ui/

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]