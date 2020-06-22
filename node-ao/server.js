const ao = require('appoptics-apm')
const axios = require('axios')
const redis = require('redis')
const express = require('express')
const app = express()

app.all('*', logAO);

function logAO(req, res, next) {
  console.log(ao.getFormattedTraceId())
  next();
}

app.get('/', function (req, res) {
  res.send('hello world - from node\n')
})

app.get('/redis', function (req, res) {
  r = redis.createClient('redis://redis')
  r.info('CPU', function (err, result) {
    res.send(result)
  })
})

app.get('/fail', function (req, res) {
  res.sendStatus(500)
})

app.get('/metric', function (req, res) {
  res.send('custom metric triggered')
})

app.get('/remote', function (req, res) {

axios.get('http://python-ao:5000/')
.then((response) => {
  console.log(`status code: ${response.statusCode}`)
  res.send(response.data)
})
.catch((error) => {
  console.error(error)
})

})

app.listen(3000, () => console.log('Example app listening on port 3000!'))
