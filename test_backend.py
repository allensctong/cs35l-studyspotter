#!/usr/bin/python3
import requests

ORIGIN = "http://localhost:5173"

def get_users():
    test_get = requests.get(
        "http://localhost:8080/api/user", 
        headers={'Origin': ORIGIN}
    )
    print(test_get.text)

def get_users_by_id(ID):
    test_get = requests.get(
        f"http://localhost:8080/api/user/{ID}", 
        headers={'Origin': ORIGIN}
    )
    print(test_get.text)


new_user = {'id': 'uwu', 'password': 'owo'}
test_post = requests.post(
    "http://localhost:8080/api/user", 
    json=new_user, 
    headers={'Origin': ORIGIN}
)

get_users()
get_users_by_id('uwu')
