FROM ruby:3.0.2

RUN apt-get update -qq \
    && apt-get install -y nodejs postgresql-client

ADD . /app
WORKDIR /app
RUN bundle install

EXPOSE 3000

CMD ["bash"]