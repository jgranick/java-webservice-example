const express = require('express');
const mysql = require('mysql2/promise');
require('dotenv').config();

const app = express();
const PORT = 8080;

app.get('/health', (req, res) => {
    res.send("Node.js App is running!");
});

app.get('/db-test', async (req, res) => {
    const dbConfig = {
        host: process.env.DB_HOST,
        port: process.env.DB_PORT || 3306,
        user: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME,
    };

    try {
        const connection = await mysql.createConnection(dbConfig);
        const [rows] = await connection.execute('SELECT NOW()');
        await connection.end();
        res.send(`DB Test: ${rows[0]['NOW()']}`);
    } catch (error) {
        console.error(error);
        res.status(500).send('DB connection failed!');
    }
});

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
