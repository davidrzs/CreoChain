require 'faraday'

conn = Faraday.new(:url => 'http://localhost:8080') do |faraday|
  faraday.request  :url_encoded             # form-encode POST params
  faraday.response :logger                  # log requests to $stdout
  faraday.adapter  Faraday.default_adapter  # make requests with Net::HTTP
end



addChainResponse = conn.post '/v1/chain/', { :name => 'testChain', :globalauthcode => 'auth', :chainaccesstoken => 'chainauth'}

addBlock1Response = conn.post '/v1/chain/testChain', { :content => 'this is some example content', :authentication => 'chainauth'}

addBlock2Response = conn.post '/v1/chain/testChain', { :content => 'this is some more content', :authentication => 'chainauth'}

fullChainResponse = conn.get '/v1/chain/testChain/'


puts fullChainResponse
