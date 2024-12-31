# Use the official Node.js image
FROM node:18-slim

# Set the working directory
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the application code
COPY app.js ./

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["node", "app.js"]
