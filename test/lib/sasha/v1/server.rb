# frozen_string_literal: true

module Sasha
  module V1
    class OperationalBucket < Sinatra::Application
      set :public_folder, 'server'
    end

    class MissingBucket < Sinatra::Application
    end

    class ErroneousBucket < Sinatra::Application
      get '/articles.yml' do
        status 500
      end

      put '/articles.yml' do
        status 500
      end

      get '/:article/article.yml' do
        status 500
      end

      put '/:article/article.yml' do
        status 500
      end
    end

    class Server < Nonnative::HTTPServer
      def initialize(service)
        if Sasha::V1.bucket_erroneous?
          super(Sinatra.new(ErroneousBucket), service)

          return
        end

        if Sasha::V1.bucket_missing?
          super(Sinatra.new(MissingBucket), service)

          return
        end

        super(Sinatra.new(OperationalBucket), service)
      end
    end
  end
end
