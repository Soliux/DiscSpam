# FIRST CONNECTION:
# {"op":0,"d":{"server_id":"921208832754061323","user_id":"883879836983115827","session_id":"e9065797acdbc7528dce9be76145459c","token":"669949b8642a090d","video":true,"streams":[{"type":"video","rid":"100","quality":100}]}}

# SECOND CONNECTION
# "op":8,"d":{"v":6,"heartbeat_interval":13750}}




import asyncio
import websockets
import json
import time


GATEWAY_URL = 'wss://gateway.discord.gg/?v=6&encoding=json'
TOKEN = 'ODgzODc5ODM2OTgzMTE1ODI3.YTQXQw.yPEExLMl5u29Ehs254nQcNrdpd4'

async def connection():
    async with websockets.connect(GATEWAY_URL) as websocket:
        # heartbea
        await websocket.send(json.dumps({
            "op": 2,
            "d": {
                "token": TOKEN,
                "v": 6,
                "compress": False,
                "large_threshold": 250,
                "properties": {
                    "$os": "linux",
                    "$browser": "discord.py",
                    "$device": "discord.py",
                    "$referrer": "",
                    "$referring_domain": ""
                }
            }
        }))
        print('Sent heartbeat')

        # identify
        await websocket.send(json.dumps({
            "op": 14,
            "d": {
                "guild_id": "921208832754061323",
                "channels": {
                    "921208833207050279": [
                        0,99
                    ]
                }
            }
        }))
        print('Grabbed User IDS???')
        b = await websocket.recv()
        print(b)


            


asyncio.get_event_loop().run_until_complete(connection())