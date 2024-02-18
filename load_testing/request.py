from locust import FastHttpUser
from const import *
from requests_toolbelt import MultipartEncoder
import json


class RequestSender:
    fast_http_user = None
    access_token = ""

    def __init__(self, fast_http_user: FastHttpUser, access_token=""):
        self.fast_http_user = fast_http_user
        self.access_token = access_token

    def set_access_token(self, access_token):
        self.access_token = access_token

    def send(self, method, path, data=None):
        post_data = None
        if isinstance(data, dict):
            post_data = json.dumps(data)
        with self.fast_http_user.rest(
            method,
            get_url(path),
            headers={
                "content-type": "application/json;charset=UTF-8",
                "authorization": f"Bearer {self.access_token}",
                "content-length": 0
                if post_data is None
                else len(post_data.encode("utf-8")),
            },
            data=post_data,
        ) as res:
            try:
                # 400系とか500系なら except にいく
                res.raise_for_status()
                return res.json()
            except Exception as err:
                res.failure(f"{type(err).__name__} {str(err)}")

    def send_form_data(self, method, path, data=None):
        form_data = MultipartEncoder(fields=data)
        with self.fast_http_user.rest(
            method=method,
            url=get_url(path),
            headers={
                "content-type": form_data.content_type,
                "authorization": f"Bearer {self.access_token}",
                "content-length": form_data._calculate_length(),
            },
            data=form_data,
        ) as res:
            try:
                # 400系とか500系なら except にいく
                res.raise_for_status()
                return res.json()
            except Exception as err:
                res.failure(f"{type(err).__name__} {str(err)}")


def get_host():
    return f"http://{ENVIRONMENT}:{PORT}"


def get_url(path):
    host = get_host()
    return f"{host}{path}"
