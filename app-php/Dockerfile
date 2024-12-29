# Use the official PHP image with PHP-FPM
FROM php:8.2-fpm

# Install PDO MySQL extension
RUN docker-php-ext-install pdo pdo_mysql

# Set the working directory
WORKDIR /var/www/html

# Copy the PHP application (index.php)
COPY index.php /var/www/html/

# Expose port 9000 for PHP-FPM to communicate with Nginx
EXPOSE 9000