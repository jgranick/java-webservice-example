<?php
// Health check endpoint
if ($_SERVER['REQUEST_URI'] === '/health') {
    http_response_code(200);
    echo "PHP App is running!";
    exit;
}

// Database test endpoint
if ($_SERVER['REQUEST_URI'] === '/db-test') {
    $dbHost = getenv('DB_HOST');
    $dbPort = getenv('DB_PORT');
    $dbName = getenv('DB_NAME');
    $dbUser = getenv('DB_USER');
    $dbPassword = getenv('DB_PASSWORD');

    $dsn = "mysql:host=$dbHost;port=$dbPort;dbname=$dbName";

    try {
        $pdo = new PDO($dsn, $dbUser, $dbPassword);
        $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

        $stmt = $pdo->query("SELECT NOW()");
        $now = $stmt->fetchColumn();

        http_response_code(200);
        echo "DB Test: " . $now;
    } catch (PDOException $e) {
        http_response_code(500);
        error_log("Error connecting to the database: " . $e->getMessage());
        echo "DB connection failed!";
    }
    exit;
}

// Fallback for undefined routes
http_response_code(404);
echo "Not Found";
