# frozen_string_literal: true

module Sasha
  module V1
    class OperationalBucket < Sinatra::Application
      set :public_folder, 'server'
    end

    class MissingArticlesBucket < Sinatra::Application
      get '/articles.yml' do
        status 404
      end
    end

    class MissingConfigBucket < Sinatra::Application
      get '/articles.yml' do
        content_type 'application/yaml'

        File.read('server/articles.yml')
      end

      get '/:article/article.yml' do
        status 404
      end
    end

    class MissingBodyBucket < Sinatra::Application
      get '/articles.yml' do
        content_type 'application/yaml'

        File.read('server/articles.yml')
      end

      get '/:article/article.yml' do |article|
        content_type 'application/yaml'

        File.read("server/#{article}/article.yml")
      end

      get '/:article/article.md' do
        status 404
      end
    end

    class ErroneousBucket < Sinatra::Application
      get '/articles.yml' do
        status 500
      end

      get '/:article/article.yml' do
        status 500
      end

      get '/:article/article.md' do
        status 500
      end
    end

    class Server < Nonnative::HTTPServer
      def initialize(service)
        if Sasha::V1.bucket_erroneous?
          super(Sinatra.new(ErroneousBucket), service)
        elsif Sasha::V1.bucket_missing_articles?
          super(Sinatra.new(MissingArticlesBucket), service)
        elsif Sasha::V1.bucket_missing_config?
          super(Sinatra.new(MissingConfigBucket), service)
        elsif Sasha::V1.bucket_missing_body?
          super(Sinatra.new(MissingBodyBucket), service)
        else
          super(Sinatra.new(OperationalBucket), service)
        end
      end
    end
  end
end
