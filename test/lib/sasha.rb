# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'nokogiri'

require 'sasha/v1/client'
require 'sasha/v1/server'

module Sasha
  module V1
    class << self
      def client
        @client ||= Sasha::V1::Client.new(Nonnative.configuration.url)
      end
    end
  end
end
