# build vue
FROM node:16 as builder
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

# generate certs
FROM alpine as certs-gen
RUN apk add curl openssl
RUN curl https://get.acme.sh | sh -s email=jmzwcn@qq.com
ENV CF_Key 8c71d824811a387925e02195889cc72bee5b5
ENV CF_Email jmzwcn@qq.com
RUN /root/.acme.sh/acme.sh  --issue  -d iyou.city -d www.iyou.city  --dns dns_cf

# copy the static dist files 
FROM nginx:alpine-slim
ADD ./nginx.conf /etc/nginx/conf.d/default.conf
#ADD ./cert/fullchain.cer ./cert/iyou.city.key /etc/
COPY --from=certs-gen /root/.acme.sh/iyou.city_ecc/fullchain.cer /root/.acme.sh/iyou.city_ecc/iyou.city.key /etc/
#ADD ./dist /usr/share/nginx/html
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
