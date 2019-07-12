import time
import random


from celery_app import app


@app.task
def add(x, y):
    time.sleep(3 + int(random.random() * 10))
    return x + y
