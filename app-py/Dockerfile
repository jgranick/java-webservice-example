# Use the official Python image
FROM python:3.10-slim

# Set the working directory
WORKDIR /app

# Copy the application code
COPY app.py /app

# Install dependencies
RUN pip install flask mysql-connector-python

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["python", "app.py"]
