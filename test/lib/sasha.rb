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
      def bucket_operational?
        bucket_state == 'operational'
      end

      def bucket_missing_articles?
        bucket_state == 'missing_articles'
      end

      def bucket_missing_config?
        bucket_state == 'missing_config'
      end

      def bucket_missing_body?
        bucket_state == 'missing_body'
      end

      def bucket_erroneous?
        bucket_state == 'erroneous'
      end

      def apply_bucket_state(state)
        ENV['BUCKET_STATE'] = state
      end

      def bucket_state
        ENV.fetch('BUCKET_STATE', nil)
      end

      def client
        @client ||= Sasha::V1::Client.new(Nonnative.configuration.url)
      end
    end
  end
end
