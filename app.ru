require 'rubygems'
require 'bundler'
require 'pry'
Bundler.require

conn = PG.connect( user: 'piotrek', password: 'pass', dbname: 'rowlf_development' )
#DB = Sequel.connect('postgres://piotrek:pass@localhost/rowlf_development', max_connections: 20)
app = proc do |env|
  #results = DB["select name, external_id FROM lodgings_lodgings LIMIT 100"]
  conn.exec( "select name, external_id FROM lodgings_lodgings LIMIT 100" ) do |results|
    json = Oj.dump({"lodgings" => results.to_a})
    [
      200,
      {
        'Content-Type' => 'application/json',
        'Content-Length' => json.bytesize.to_s,
      },
      [json]
    ]
  end
end

run app
