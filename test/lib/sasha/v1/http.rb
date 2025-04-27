# frozen_string_literal: true

module Sasha
  module V1
    class HTTP < Nonnative::HTTPClient
      def root(opts = {})
        get('/', opts)
      end
    end
  end
end
