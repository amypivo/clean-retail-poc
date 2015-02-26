var fs = require('fs');
var http = require('http');
var url = require('url');

var fileContents = fs.readFileSync(process.argv[2]).toString();
var destinationURL = process.argv[3];
var productsArray = JSON.parse(fileContents);

var requestBody = "";
var elasticOperation = "{\"index\": {}}";

for (i=0; i<productsArray.length; i++) {
    requestBody = requestBody.concat(elasticOperation).concat("\n");
    requestBody = requestBody.concat(JSON.stringify(productsArray[i])).concat("\n");
}

var postOptions = url.parse(destinationURL);
postOptions.method = 'POST';

var request = http.request(postOptions, function (response) {
    console.log("Response: " + response.statusCode);

    response.on('data', function (chunk) {
    });
});

request.write(requestBody);
request.end();
