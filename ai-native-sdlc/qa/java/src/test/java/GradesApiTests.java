import io.restassured.RestAssured;
import org.testng.annotations.BeforeClass;
import org.testng.annotations.Test;
import static io.restassured.RestAssured.*;
import static org.hamcrest.Matchers.*;

public class GradesApiTests {
  @BeforeClass
  public void setup() { RestAssured.baseURI = "http://localhost:8080"; }

  @Test
  public void listGrades() {
    given().when().get("/grades").then().statusCode(200).body("$", is(notNullValue()));
  }

  @Test
  public void crudGrade() {
    int id =
      given().contentType("application/json")
        .body("{\"student_id\":1,\"course\":\"History\",\"score\":78}")
        .when().post("/grades").then().statusCode(201)
        .extract().path("id");

    given().when().get("/grades/" + id).then().statusCode(200);

    given().contentType("application/json")
      .body("{\"student_id\":1,\"course\":\"History\",\"score\":80}")
      .when().put("/grades/" + id).then().statusCode(200);

    given().when().delete("/grades/" + id).then().statusCode(204);
  }
}
