Feature: Order creation

  Scenario: Create a new order successfully
    Given I have a valid order payload
    When I send a POST request to "/api/v1/order"
    Then the response code should be 200
    And the response should contain the order ID