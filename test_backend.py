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

def create_user(ID, password):
    new_user = {'username': ID, 'password': password}
    test_post = requests.post(
        "http://localhost:8080/api/signup", 
        json=new_user, 
        headers={'Origin': ORIGIN}
    )

def login(ID, password):
    new_user = {'username': ID, 'password': password}
    test_login = requests.post(
        "http://localhost:8080/api/login", 
        json=new_user, 
        headers={'Origin': ORIGIN}
    )
    print(test_login.cookies)

create_user('RS', '1239120')
create_user('RS', '1239120')
login('RS', '1239120')
get_users()
get_users_by_id('RS')
