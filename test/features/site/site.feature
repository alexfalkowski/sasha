Feature: Sasha
  Sasha is a website for sasha-adventures.com

  Scenario Outline: Visit sections
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>" with status <status>

    Examples:
      | layout  | section           | status |
      | full    | root              |    200 |
      | full    | home              |    200 |
      | full    | articles          |    200 |
      | full    | article           |    200 |
      | full    | article_not_found |    404 |
      | partial | root              |    200 |
      | partial | home              |    200 |
      | partial | articles          |    200 |
      | partial | article           |    200 |
      | partial | article_not_found |    404 |
