var fs = require('fs');
var axios = require('axios');
var cheerio = require('cheerio');
var url = 'https://gobyexample.com';
var baseDir = process.argv[2] || 'learn-go-by-example';

function check(error) {
  if (error) {
    throw error;
  }
}

fs.mkdir(baseDir, function(error) {
  check(error);
  axios.get(url).then(function(response) {
    var $ = cheerio.load(response.data);
    $('li > a').each(function(index, element) {
      let num = index < 10 ? '0' + index : index;
      let filePath = `${num}-${$(element).attr('href')}`;
      let fileName = `${filePath}.go`;
      let partialPath = `${baseDir}/${filePath}`;
      let fullPath = `${baseDir}/${filePath}/${fileName}`;
      fs.mkdir(partialPath, function(error) {
        check(error);
        fs.writeFile(fullPath, '', function(error) {
          check(error);
        });
      });
    });
  }).catch(function(error) {
    check(error);
  });
});
