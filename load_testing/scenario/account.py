from locust import SequentialTaskSet, task


class Account(SequentialTaskSet):
    @task
    def get(self):
        self.user.request_sender.send(
            method="GET",
            path=f"/api/admin/account",
        )
