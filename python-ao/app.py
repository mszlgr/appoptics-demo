# app.py - a minimal flask api using flask_restful
from flask import Flask
from flask_restful import Resource, Api
import requests
import redis
import appoptics_apm
import appoptics_apm.middleware
# the flask app

class LogAO(object):
  def __init__(self, app):
    self.app = app

  def __call__(self, environ, start_response):
#    app.logger.info(environ)
    app.logger.info(appoptics_apm.get_log_trace_id())
    return self.app(environ, start_response)


app = Flask(__name__)
app.wsgi_app = LogAO(app.wsgi_app)
app.wsgi_app = appoptics_apm.middleware.AppOpticsApmMiddleware(app.wsgi_app)
api = Api(app)

@app.route('/')
def hello_world():
    appoptics_apm.custom_metrics_increment("python-my-counter-metric", 1)
    return 'Hello, World! - from python-ao\n'

@app.route('/custom')
def custom():
    app.logger.info(appoptics_apm.get_log_trace_id())
    appoptics_apm.set_transaction_name('custom_transaction_name')
    appoptics_apm.log_exception()
    return 'Hello, World!'

@app.route('/fail')
def fail():
    appoptics_apm.log_error(object, 'unexpected result!')
    return 'Not found', 500

@app.route('/remote')
def remote():
    r = requests.get('http://golang-ao:8000/')
    return r.text, 200

@app.route('/redis')
def redis_handler():
    r = redis.Redis(host='redis')
    return r.info(section='CPU')

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
