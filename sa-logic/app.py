from flask import Flask, request, jsonify
from snownlp import SnowNLP

app = Flask(__name__)


@app.route('/analyse/sentiment', methods=['POST'])
def analyse_sentiment():
    s = SnowNLP(request.get_json()['sentence'])
    sentence = SnowNLP(s.sentences[0])
    return jsonify(
        code=0,
        data={
            'sentence': s.sentences[0],
            'sentiments': sentence.sentiments,
        }
    )


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
