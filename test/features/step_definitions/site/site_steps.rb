# frozen_string_literal: true

Before('@operational') do
  Sasha::V1.apply_bucket_state 'operational'
end

Before('@missing_articles') do
  Sasha::V1.apply_bucket_state 'missing_articles'
end

Before('@missing_config') do
  Sasha::V1.apply_bucket_state 'missing_config'
end

Before('@missing_body') do
  Sasha::V1.apply_bucket_state 'missing_body'
end

Before('@erroneous') do
  Sasha::V1.apply_bucket_state 'erroneous'
end

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

Then('I should see {string} succesfully') do |section|
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/html; charset=utf-8')

  expected = {
    'root' => 'Oprah Winfrey',
    'home' => 'Oprah Winfrey',
    'articles' => 'Articles',
    'article' => 'This is article 1.',
    'article_not_found' => 'repository: get article: not found'
  }

  expected_section = expected[section]
  if expected_section
    body = @response.body
    html = Nokogiri::HTML.parse(body)

    expect(html.text).to include(expected[section])
  end
end

Then('the {string} is not found') do |_string|
  expect(@response.code).to eq(404)
end

Then('the {string} is erroneous') do |_string|
  expect(@response.code).to eq(500)
end
