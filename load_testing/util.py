import random
import string
import base64


def random_string(n):
    chars = string.ascii_letters + string.digits
    return "".join(random.choices(chars, k=n))


def random_number(min, max):
    return random.randint(min, max)


def profile_icon():
    with open("profile.png", "rb") as image:
        return base64.b64encode(bytearray(image.read())).decode("utf-8")
