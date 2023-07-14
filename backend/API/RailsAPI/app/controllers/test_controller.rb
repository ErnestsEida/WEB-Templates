class TestController < ApplicationController
  def test
    render json: { response: "pong" }, status: 200
  end
end