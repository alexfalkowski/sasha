# frozen_string_literal: true

module Sasha
  module V1
    class HTTP < Nonnative::HTTPClient
      def get_root(opts = {})
        get('/', opts)
      end

      def put_root(opts = {})
        put('/', opts)
      end

      def get_articles(opts = {})
        get('/articles', opts)
      end

      def put_articles(opts = {})
        put('/articles', opts)
      end

      def get_article(opts = {})
        get('/article/article-1', opts)
      end

      def put_article(opts = {})
        put('/article/article-1', opts)
      end

      def get_article_not_found(opts = {})
        get('/article/none', opts)
      end

      def put_article_not_found(opts = {})
        put('/article/none', opts)
      end

      alias get_home get_root
      alias put_home put_root
    end
  end
end
