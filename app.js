var http = require('http');
var cluster = require('cluster');
var pg = require('pg');
var conString = "postgres://piotrek:pass@localhost/rowlf_development";

var numCPUs = require('os').cpus().length;

if (cluster.isMaster) {
  // Fork workers.
  for (var i = 0; i < numCPUs; i++) {
    cluster.fork();
  }

  cluster.on('exit', function(worker, code, signal) {
    console.log('worker ' + worker.process.pid + ' died');
  });
} else {

  pg.defaults.poolSize = 10;

  http.createServer(function(req, res, next) {
    pg.connect(conString, function(err, client, done) {
      var handleError = function(err) {
        if(!err) return false;
        done(client);
        next(err);
        return true;
      };
      client.query('select name, external_id FROM lodgings_lodgings LIMIT 100', null, function(err, result) {
        if(handleError(err)) return;
        done();
        res.writeHead(200, {'content-type': 'application/json'});
        res.end(JSON.stringify(result.rows));
      });
    });
  }).listen(3002);
}
