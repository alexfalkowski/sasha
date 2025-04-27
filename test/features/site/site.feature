Feature: Sasha
  Sasha is a website for sasha-adventures.com

  Scenario Outline: Visit sections
    When I visit "<section>"
    Then I should see "<section>"

    Examples:
      | section |
      | root    |
