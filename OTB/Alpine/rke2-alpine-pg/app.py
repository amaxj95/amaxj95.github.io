from flask import Flask, request
from psycopg2 import connect

app = Flask(__name__)

@app.route("/")
def hello_world():
  return "Hello, World!"

if __name__ == "__main__":
  connection = connect("dbname='my_database' user='my_user' password='my_password'")
  cursor = connection.cursor()
  cursor.execute("SELECT * FROM my_table")
  rows = cursor.fetchall()
  for row in rows:
    print(row)
  connection.close()