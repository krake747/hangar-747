FROM node:22-alpine AS build
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build:ci

FROM nginx:alpine
COPY --from=build /app/dist/buenowind/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
