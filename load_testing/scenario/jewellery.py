from locust import SequentialTaskSet, task


class Jewellery(SequentialTaskSet):
    # @task
    # def getAdminJewellery(self):
    #     self.user.request_sender.send(
    #         method="GET",
    #         path=f"/api/admin/jewellery",
    #     )

    @task
    def getCategory(self):
        self.user.request_sender.send(
            method="GET",
            path=f"/api/category",
        )

    @task
    def getGem(self):
        self.user.request_sender.send(
            method="GET",
            path=f"/api/gem",
        )

    @task
    def getMaterial(self):
        self.user.request_sender.send(
            method="GET",
            path=f"/api/material",
        )
