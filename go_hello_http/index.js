const spawnSync = require('child_process').spawnSync;

exports.Handler = function Handler(req, res) {
  var requestBody;

  switch (req.get('content-type')) {
    case 'application/json':
      requestBody = JSON.stringify(req.body);
      break;
    case 'application/octet-stream':
      requestBody = req.body;
      break;
    case 'text/plain':
      requestBody = req.body;
      break;
  }

  var fullUrl = req.protocol + '://' + req.get('host') + req.originalUrl;
  var httpRequest = {
    'url': fullUrl
    'method': req.method,
    'headers': req.headers,
    'body': requestBody,
    'remote_addr': req.ip,
  };

  var args = [
    '--event-type', 'http'
  ];

  result = spawnSync('./cloud-functions-go-shim', args, {
    input: JSON.stringify(httpRequest),
    stdio: 'pipe',
  });

  if (result.status !== 0) {
     console.log(result.stderr.toString());
     res.status(500);
     return;
  }

  data = JSON.parse(result.stdout);
  res.status(data.status_code);
  res.send(data.body);
};
