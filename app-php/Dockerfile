# Use the official PHP image with Apache
FROM php:8.2-apache

# Install PDO MySQL extension
RUN docker-php-ext-install pdo pdo_mysql

# Copy custom Apache config to enable URL rewriting
COPY 000-default.conf /etc/apache2/sites-available/000-default.conf

# Enable mod_rewrite (for routing)
RUN a2enmod rewrite

# Set the working directory
WORKDIR /var/www/html

# Copy application files
COPY src/* /var/www/html/

# Expose port 80
EXPOSE 80
