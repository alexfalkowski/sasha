# frozen_string_literal: true

module Sasha
  module V1
    class Bucket < Sinatra::Application
      set :public_folder, 'server'

      get '/error/article.yml' do
        status 500
      end
    end

    class Server < Nonnative::HTTPServer
      def initialize(service)
        super(Sinatra.new(Bucket), service)
      end
    end
  end
end
