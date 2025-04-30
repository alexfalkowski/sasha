Feature: Sasha
  Sasha is a website for sasha-adventures.com

  Scenario Outline: Visit successfully sections
    When I visit "<section>" with layout "<layout>"
    Then I should see "<section>" with status <status>

    Examples:
      | layout  | section           | status |
      | full    | root              |    200 |
      | full    | home              |    200 |
      | full    | articles          |    200 |
      | full    | article           |    200 |
      | full    | article_not_found |    404 |
      | full    | article_error     |    500 |
      | partial | root              |    200 |
      | partial | home              |    200 |
      | partial | articles          |    200 |
      | partial | article           |    200 |
      | partial | article_not_found |    404 |
      | partial | article_error     |    500 |

  Scenario Outline: Visit unsuccessfully sections
    Given I set the proxy for server "<server>" to "<failure>"
    When I visit "<section>" with layout "<layout>"
    Then I should see an error with status <status>
    And I should reset the proxy for server "<server>"

    Examples:
      | layout  | section  | status | server | failure      |
      | full    | articles |    500 | bucket | close_all    |
      | full    | articles |    500 | bucket | invalid_data |
      | full    | article  |    500 | bucket | close_all    |
      | full    | article  |    500 | bucket | invalid_data |
      | partial | articles |    500 | bucket | close_all    |
      | partial | articles |    500 | bucket | invalid_data |
      | partial | article  |    500 | bucket | close_all    |
      | partial | article  |    500 | bucket | invalid_data |
