Feature: shorten url

  Scenario: successful shorten a url
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