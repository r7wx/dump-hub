FROM node AS react-builder
WORKDIR /opt/dump-hub-app
COPY . .
RUN npm install
RUN npm run build

FROM nginx AS dump-hub-web
WORKDIR /usr/share/nginx/html
COPY --from=react-builder /opt/dump-hub-app/build .
WORKDIR /etc/nginx
RUN openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
    -subj "/C=AU/ST=None/L=None/O=r7wx/CN=dump-hub.priv" \
    -keyout dump-hub.key  -out dump-hub.crt