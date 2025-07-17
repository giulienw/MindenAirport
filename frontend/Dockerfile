# Build stage
FROM node:18-alpine AS builder

ENV VITE_API_BASE_URL=/api

# Set working directory
WORKDIR /app

# Install Bun
RUN npm install -g bun

# Copy package files
COPY package.json bun.lockb ./

# Install dependencies
RUN bun install

# Copy source code
COPY . .

# Build the application for web (excluding electron)
RUN bun run build:web-docker

# Runtime stage
FROM nginx:alpine

# Copy built application to nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
