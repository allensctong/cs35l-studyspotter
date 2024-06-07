#!/usr/bin/python3
import requests

ORIGIN = "http://localhost:5173"
cookies = {}

def get_users():
    test_get = requests.get(
        "http://localhost:8080/api/user", 
        headers={'Origin': ORIGIN}, cookies=cookies
    )
    print(test_get.text)

def get_users_by_id(ID):
    test_get = requests.get(
        f"http://localhost:8080/api/user/{ID}", 
        headers={'Origin': ORIGIN}, cookies=cookies
    )
    print(f"http://localhost:8080/api/user/{ID}")
    print(test_get.text)

def create_user(ID, password):
    new_user = {'username': ID, 'password': password}
    test_post = requests.post(
        "http://localhost:8080/api/signup", 
        json=new_user, 
        headers={'Origin': ORIGIN}, cookies=cookies
    )
    print(test_post.text)

def login(ID, password):
    new_user = {'username': ID, 'password': password}
    test_login = requests.post(
        "http://localhost:8080/api/login", 
        json=new_user, 
        headers={'Origin': ORIGIN}, cookies=cookies
    )
    return test_login.cookies

def like():
    body = {'username': 'uwu'}
    test_post = requests.put(
        "http://localhost:8080/api/post/2/like", 
        json=body, 
        headers={'Origin': ORIGIN}, cookies=cookies
    )
    print(test_post.text)

def friend():
    body = {'username': 'uwu'}
    test_post = requests.put(
        "http://localhost:8080/api/user/t1/friend", 
        json=body, 
        headers={'Origin': ORIGIN}, cookies=cookies
    )
    print(test_post.text)

login('uwu', 'owo')
friend()
'''
create_user('RS', '1239120')
create_user('R', '1239120')
create_user('S', '1239120')
create_user('RSA', '1239120')
create_user('uwu', 'owo')
create_user('owo', '1239120')
create_user('iwi', '1239120')
create_user('QAQ', '1239120')
get_users_by_id('RS') '''
'''get_users_by_id('RS')
create_user('RS', '1239120')
create_user('RS', '1239120')
cookies = login('RS', '1239120')
print(cookies)
get_users_by_id('RS')
'''
