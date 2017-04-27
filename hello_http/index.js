/**
 * Handles HTTP request with custom header, method, query-string,
 * and JSON body
 *
 * @param {Object} req Cloud Function request context.
 * @param {Object} res Cloud Function response context.
 */
exports.helloHttp = function helloHttp (req, res) {
  let name;

  console.log("Method = " + req.method + 
    ", Header = " + req.get('x-myheader') +
    ", Query = " + req.query.foo +
    ", JSON Field = " + req.body.foo);

  switch (req.method) {
    case 'GET':
    case 'POST':
      break;
    default:
      return res.status(500).send({ error: 'Something blew up!' });
      break;
  }

  switch (req.get('content-type')) {
    // '{"name":"Mark"}'
    case 'application/json':
      name = req.body.name;
      break;

    // 'Mark' stored in a Buffer
    case 'application/octet-stream':
      name = req.body.toString(); // Convert buffer to a string
      break;

    // 'Mark' string
    case 'text/plain':
      name = req.body;
      break;

    // 'name=Mark'
    case 'application/x-www-form-urlencoded':
      name = req.body.name;
      break;
  }

  res.status(200).send(`Howdy ${name || 'World'}!`);
};
