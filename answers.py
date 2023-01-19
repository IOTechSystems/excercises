import base64
import json
from datetime import datetime

with open("exercise-02/data/data.json") as data_file:
  json_data = json.load(data_file)

def convert_base64(base64_message):
  base64_bytes = base64_message.encode('ascii')
  message_bytes = base64.b64decode(base64_bytes)
  message = message_bytes.decode('ascii')
  return  int(message)

def get_time(unix_time):
  return datetime.utcfromtimestamp(int(unix_time)).strftime('%Y-%m-%d %H:%M:%S')

total = 0
uuids = []
current_time = datetime.now().replace(microsecond=0)

for data in json_data['Devices']:
  timestamp = datetime.strptime(get_time(data['timestamp']), '%Y-%m-%d %H:%M:%S')

  if  timestamp > current_time:
    total += convert_base64(data['value'])
    uuids.append(data['Info'].split(':')[1].split(',')[0])

dictionary = {
"ValueTotal":total,
"UUIDS": uuids
}

with open("answers.json", "w+") as outfile:
  json.dump(dictionary, outfile)

