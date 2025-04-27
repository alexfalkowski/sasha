# frozen_string_literal: true

module Sasha
  module V1
    class HTTP < Nonnative::HTTPClient
      def root(opts = {})
        get('/', opts)
      end

      def articles(opts = {})
        get('/articles', opts)
      end

      def article(opts = {})
        get('/article/article-1', opts)
      end

      def article_not_found(opts = {})
        get('/article/none', opts)
      end

      alias home root
    end
  end
end
