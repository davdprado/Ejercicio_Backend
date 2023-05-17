import requests
from flask import Flask,jsonify


app = Flask(__name__)

@app.route('/getList', methods=['GET'])
def getRandomJokes():
    jokes =[]
    while len(jokes)<25:
        response = requests.get('https://api.chucknorris.io/jokes/random')
        if response.status_code ==200:
            data = response.json()
            if data['id'] not in [joke.get('id') for joke in jokes]:
                jokes.append({'id':data['id'], 'value':data['value']})
    #verificar que no hay IDS Repetidos
    ids_repeat = []
    ids = set()
    for item in jokes:
        if item['id'] in jokes:
            ids_repeat.append(item['id'])
        else:
             ids.add(item['id'])

    if ids_repeat:
        print("Se encontraron ids repetidos")
    else:
        print("no existen IDS repetidos")

    return jsonify(jokes)

if __name__ == '__main__':
    app.run()
    
    