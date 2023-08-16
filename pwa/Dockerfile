# first stage builds vue
FROM node:16 as builder
WORKDIR /app
COPY . .
RUN npm install && npm run build

# second stage copies the static dist files 
FROM nginx:alpine-slim
ADD ./nginx.conf /etc/nginx/conf.d/default.conf
ADD ./cert/fullchain.cer ./cert/iyou.city.key /etc/
#ADD ./dist /usr/share/nginx/html
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
