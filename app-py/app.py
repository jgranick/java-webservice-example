from flask import Flask
import mysql.connector
import os

app = Flask(__name__)

@app.route('/health', methods=['GET'])
def health():
    return "Python App is running!"

@app.route('/db-test', methods=['GET'])
def db_test():
    db_config = {
        'host': os.getenv('DB_HOST'),
        'port': os.getenv('DB_PORT', 3306),
        'database': os.getenv('DB_NAME'),
        'user': os.getenv('DB_USER'),
        'password': os.getenv('DB_PASSWORD')
    }
    try:
        connection = mysql.connector.connect(**db_config)
        cursor = connection.cursor()
        cursor.execute("SELECT NOW()")
        result = cursor.fetchone()
        cursor.close()
        connection.close()
        return f"DB Test: {result[0]}"
    except Exception as e:
        print(e)
        return "DB connection failed!", 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
