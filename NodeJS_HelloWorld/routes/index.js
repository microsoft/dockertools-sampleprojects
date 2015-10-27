var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { 
    title: 'NodeJS Hello World', 
    platform: req.app.get('platform'), 
    hostname: req.app.get('hostname') })
});

module.exports = router;
