# Build vue
FROM node:16 AS builder
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

# Generate certs
FROM neilpang/acme.sh AS certs
ENV CF_Key=8c71d824811a387925e02195889cc72bee5b5 CF_Email=jmzwcn@qq.com
RUN acme.sh --register-account -m jmzwcn@qq.com && --issue -d iyou.city -d *.iyou.city --dns dns_cf

# Copy static files 
FROM nginx:alpine-slim
# ENV CF_Key=8c71d824811a387925e02195889cc72bee5b5 CF_Email=jmzwcn@qq.com
# RUN apk add curl openssl && \
#     curl https://get.acme.sh | sh -s email=jmzwcn@qq.com && cd /root/.acme.sh && \
#     ./acme.sh --issue -d iyou.city -d *.iyou.city  --dns dns_cf && \
#     ./acme.sh --install-cert -d iyou.city --key-file /etc/iyou.city.key --fullchain-file /etc/fullchain.cer

ADD ./nginx.conf /etc/nginx/conf.d/default.conf
#ADD ./cert/fullchain.cer ./cert/iyou.city.key /etc/
COPY --from=certs /acme.sh/iyou.city_ecc/fullchain.cer /acme.sh/iyou.city_ecc/iyou.city.key /etc/
#ADD ./dist /usr/share/nginx/html
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
