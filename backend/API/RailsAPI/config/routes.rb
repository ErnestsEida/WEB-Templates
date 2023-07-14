Rails.application.routes.draw do
  scope :api, defaults: { format: :json } do
    get :ping, to: "test#test"
  end
end
