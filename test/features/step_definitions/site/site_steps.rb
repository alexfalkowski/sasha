# frozen_string_literal: true

When('I visit {string} with layout {string}') do |section, layout|
  opts = {
    headers: { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0' },
    read_timeout: 10, open_timeout: 10
  }
  methods = {
    'full' => 'get',
    'partial' => 'put'
  }
  method = methods[layout]

  @response = Sasha::V1.client.send("#{method}_#{section}", opts)
end

Then('I should see {string} with status {int}') do |section, status|
  expect(@response.code).to eq(status)
  expect(@response.headers[:content_type]).to eq('text/html; charset=utf-8')

  expected = {
    'root' => 'Oprah Winfrey',
    'home' => 'Oprah Winfrey',
    'articles' => 'Articles',
    'article' => 'This is article 1.',
    'article_not_found' => 'Article not found!'
  }

  expected_section = expected[section]
  if expected_section
    body = @response.body
    html = Nokogiri::HTML.parse(body)

    expect(html.text).to include(expected[section])
  end
end

Then('I should see an error with status {int}') do |status|
  expect(@response.code).to eq(status)
end
