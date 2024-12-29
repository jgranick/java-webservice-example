import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.Statement;
import java.util.logging.Logger;

import static spark.Spark.*;

public class Main {
    private static final Logger logger = Logger.getLogger(Main.class.getName());

    public static void main(String[] args) {
        port(8080);

        get("/health", (req, res) -> "Java App is running!");

        get("/db-test", (req, res) -> {
            String jdbcUrl = String.format("jdbc:mysql://%s:%s/%s", 
                System.getenv("DB_HOST"), 
                System.getenv("DB_PORT"), 
                System.getenv("DB_NAME"));

            try (Connection connection = DriverManager.getConnection(jdbcUrl, 
                                                                      System.getenv("DB_USER"), 
                                                                      System.getenv("DB_PASSWORD"))) {
                Statement statement = connection.createStatement();
                ResultSet resultSet = statement.executeQuery("SELECT NOW()");
                resultSet.next();
                return "DB Test: " + resultSet.getString(1);
            } catch (Exception e) {
                e.printStackTrace();
                return "DB connection failed!";
            }
        });
    }
}
