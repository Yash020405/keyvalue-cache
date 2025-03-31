import string
from locust import FastHttpUser, task, constant
import random
import uuid

# Configuration
KEY_POOL_SIZE = 10_000  # Shared key pool
VALUE_LENGTH = 256
PUT_RATIO = 0.5  # 50% PUT requests

class CacheUser(FastHttpUser):
    wait_time = constant(0)  # No wait time for maximum throughput

    # Pre-generate keys and values for better cache hit rates
    key_pool = [str(uuid.uuid4()) for _ in range(KEY_POOL_SIZE)]
    value_pool = [''.join(random.choices(string.printable, k=VALUE_LENGTH))
                  for _ in range(KEY_POOL_SIZE)]

    @task
    def mixed_load(self):
        """50/50 GET/PUT ratio."""
        if random.random() < PUT_RATIO:
            self.put_request()
        else:
            self.get_request()

    def put_request(self):
        key = random.choice(self.key_pool)
        value = random.choice(self.value_pool)
        self.client.post("/put", json={"key": key, "value": value}, name="/put")

    def get_request(self):
        key = random.choice(self.key_pool)
        self.client.get(f"/get?key={key}", name="/get")