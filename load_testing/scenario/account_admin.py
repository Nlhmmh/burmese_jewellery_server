from locust import SequentialTaskSet, task


class AccountAdmin(SequentialTaskSet):
    @task
    def get(self):
        self.user.request_sender.send(
            method="GET",
            path=f"/api/admin/account_admin",
        )
