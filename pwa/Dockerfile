FROM nginx:alpine

ADD ./nginx.conf /etc/nginx/conf.d/default.conf
#ADD ./cert/fullchain.cer ./cert/iyou.city.key /etc/
ADD ./dist /usr/share/nginx/html

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
