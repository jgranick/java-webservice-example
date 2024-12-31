# Use the official Ruby image
FROM ruby:3.2-slim

# Install necessary tools and libraries
RUN apt-get update && apt-get install -y \
  build-essential \
  default-libmysqlclient-dev \
  && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy Gemfile and Gemfile.lock
# COPY Gemfile Gemfile.lock ./
COPY Gemfile ./

# Install gems
RUN bundle install

# Copy application code
COPY app.rb ./

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["ruby", "app.rb"]
