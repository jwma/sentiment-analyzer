# Build stage
FROM node:10-alpine as builder

ENV VUE_APP_API_HOST=http://localhost:8080

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Final stage
FROM nginx:1.16-alpine
COPY --from=builder /app/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]