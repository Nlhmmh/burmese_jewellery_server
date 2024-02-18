from locust import FastHttpUser, between
from const import *
from util import *
import scenario
from request import RequestSender


class User(FastHttpUser):
    tasks = {
        scenario.AccountAdmin: 1,
        scenario.Account: 1,
        scenario.Jewellery: 10,
    }

    wait_time = between(1.0, 3.0)

    def __init__(self, parent):
        super().__init__(parent)
        self.request_sender = RequestSender(fast_http_user=self)

    def on_start(self):
        self.account_id = ""
        self.email = f"{random_string(16)}@10antz.co.jp"
        self.password = "pass123@"
