FROM node:14-alpine as build

WORKDIR /app

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .

RUN  npm run build --prod
RUN ls dist/client
FROM nginx:stable-alpine
COPY --from=build /app/dist/client/ /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
