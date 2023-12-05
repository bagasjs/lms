import requests

user_data = {
    "name" : "Bagas Jonathan Sitanggang",
    "email" : "stgbgs@gmail.com",
    "password" : "test123",
    "password_confirmation" : "test123",
}

res = requests.post(url="http://localhost:6969/api/users", data=user_data)
if res.ok:
    print(res.json())
else:
    print("Request failed")
    print("{}".format(res.reason))
