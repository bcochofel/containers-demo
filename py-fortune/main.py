from flask import Flask
from flask_restful import Resource, Api
from flask_cors import CORS
import requests

app = Flask(__name__)
# To allow direct AJAX calls
CORS(app)


@app.route('/', methods=['GET'])
def home():
    r = requests.get('http://yerkee.com/api/fortune')

    return r.json()


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5001)
