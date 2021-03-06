var net = require('net'),
    util = require('util');

var server = net.createServer(function (socket) {
  var n = 0;
  socket.on('data', function (data) {
    socket.write(data);
    n += data.length;
  });
  socket.on('error', function(err) {
    util.log(err)
  });
  socket.on('end', function() {
    util.log("echoed "+n+" byte to "+ socket.remoteAddress+":"+socket.remotePort);
  });
});

server.listen(3640, "0.0.0.0");
