import os, subprocess

db_name = "practica"
path = os.path.split(os.path.abspath(__file__))[0]

for filename in os.listdir(path):

    if ".json" in filename and os.path.isfile(os.path.join(path, filename)):

        collection = filename.split(".")[0]

        subprocess.run([
            'mongoimport',
            '--db', db_name,
            '--collection', collection,
            '--drop',
            '--file', filename
        ])

