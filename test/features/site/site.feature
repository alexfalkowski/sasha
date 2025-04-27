Feature: Sasha
  Sasha is a website for sasha-adventures.com

  Scenario Outline: Visit sections
    When I visit "<section>"
    Then I should see "<section>" with status <status>

    Examples:
      | section           | status |
      | root              |    200 |
      | home              |    200 |
      | articles          |    200 |
      | article           |    200 |
      | article_not_found |    404 |
