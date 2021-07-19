Feature: shorten url

  Scenario: successful redirect
    Given this url exists "www.abc.com" for short "def"
    When I send a "GET" to "/def"
    Then I should get a 301 HTTP response code

  Scenario: not found redirect
    When I send a "GET" to "/def"
    Then I should get a 404 HTTP response containing the following "StandardResponse":
      | StatusCode  | 404       |
      | Description | Not Found |

  Scenario: successful shorten then redirect
    When I send a "POST" to "/shorten" with the following:
      """
      {
      "original":"www.abc.com"
      }
      """
    Then I should get a 200 HTTP response containing the following "ShortResponse":
      | StatusCode  | 200                        |
      | Description | Ok                         |
      | Short       | http://localhost:8080/def |
    When I send a "GET" to "http://localhost:8080/def"
    Then I should get a 301 HTTP response code
