using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using MySql.Data.MySqlClient;

var builder = WebApplication.CreateBuilder(args);

var app = builder.Build();

app.MapGet("/health", () => "C# App is running!");

app.MapGet("/db-test", () =>
{
    var connectionString = $"Server={Environment.GetEnvironmentVariable("DB_HOST")};" +
                           $"Port={Environment.GetEnvironmentVariable("DB_PORT")};" +
                           $"Database={Environment.GetEnvironmentVariable("DB_NAME")};" +
                           $"User={Environment.GetEnvironmentVariable("DB_USER")};" +
                           $"Password={Environment.GetEnvironmentVariable("DB_PASSWORD")};";

    try
    {
        using var connection = new MySqlConnection(connectionString);
        connection.Open();

        using var command = new MySqlCommand("SELECT NOW()", connection);
        var result = command.ExecuteScalar();
        return Results.Ok($"DB Test: {result}");
    }
    catch (Exception ex)
    {
        Console.WriteLine(ex);
        return Results.Problem("DB connection failed!", statusCode: 500);
    }
});

app.Run("http://0.0.0.0:8080");
