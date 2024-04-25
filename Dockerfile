FROM node:16-alpine as builder

WORKDIR /admin-ui

COPY ./code-help-admin-ui .

RUN npm install
RUN npm run build

FROM nginx:stable as build

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /admin-ui/dist/ /usr/share/nginx/html/admin-ui/

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]