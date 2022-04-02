# curl -X POST -d "{\"name\":\"felix\",\"url\":\"abcd.com\",\"comment\":\"kannur\"}" http://localhost:8080/index


import requests
import json
import asyncio
from websocket import create_connection

data={"name":"felix","url":"abcd.com","comment":"kannur"}
number=1

def httpConnection():
	r = requests.post("ws://localhost:8080/index", data=json.dumps(data))
	print(r, r.text)

def socketConnection():
	ws = create_connection("ws://localhost:8080/index")
	ws.send(json.dumps(data))
	while True:
		try:
			result =  ws.recv()
			print(result)
		except:
			ws.close()
			break

socketConnection()