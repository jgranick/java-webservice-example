require 'sinatra'
require 'mysql2'
require 'dotenv'

Dotenv.load

set :bind, '0.0.0.0'
set :port, 8080

get '/health' do
  "Ruby App is running!"
end

get '/db-test' do
  begin
    client = Mysql2::Client.new(
      host: ENV['DB_HOST'],
      port: ENV['DB_PORT'] || 3306,
      database: ENV['DB_NAME'],
      username: ENV['DB_USER'],
      password: ENV['DB_PASSWORD']
    )
    result = client.query('SELECT NOW()').first
    "DB Test: #{result['NOW()']}"
  rescue StandardError => e
    puts e.message
    status 500
    "DB connection failed!"
  ensure
    client&.close
  end
end
