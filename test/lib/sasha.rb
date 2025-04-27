# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'nokogiri'

require 'sasha/v1/http'

module Sasha
  module V1
    class << self
      def http
        @http ||= Sasha::V1::HTTP.new(Nonnative.configuration.url)
      end
    end
  end
end
