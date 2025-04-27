# frozen_string_literal: true

When('I visit {string}') do |section|
  opts = {
    headers: { request_id: SecureRandom.uuid, user_agent: 'Web-client/1.0 HTTP/1.0' },
    read_timeout: 10, open_timeout: 10
  }

  @response = Sasha::V1.http.send(section, opts)
end

Then('I should see {string}') do |section|
  expect(@response.code).to eq(200)
  expect(@response.headers[:content_type]).to eq('text/html; charset=utf-8')

  m = { 'root' => "Sasha's Adventures" }
  html = Nokogiri::HTML.parse(@response.body)

  expect(html.text).to include(m[section])
end
