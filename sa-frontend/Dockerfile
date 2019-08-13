# Build stage
FROM node:10-alpine as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Final stage
FROM nginx:1.16-alpine
COPY --from=builder /app/dist /usr/share/nginx/html