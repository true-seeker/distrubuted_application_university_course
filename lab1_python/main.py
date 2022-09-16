from flask import Flask, request

from services import import_to_excel

app = Flask(__name__)


@app.route('/excel_import', methods=['POST'])
def excel_import():
    print(request.get_json())
    import_to_excel(request.get_json())
    return 'Hello World!'


if __name__ == '__main__':
    app.run(port=80, host='0.0.0.0')
