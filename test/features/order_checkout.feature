Feature: Order checkout

  Scenario: Create a checkout order successfully
    Given I have a valid order payload
    When I send a POST request to 
    Then the response code should be 200
    And the response should contain the order ID