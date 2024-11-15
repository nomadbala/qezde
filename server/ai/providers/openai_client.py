from g4f.client import Client
from tenacity import retry, stop_after_attempt, wait_exponential

class G4FClient:
    def __init__(self, model):
        self.model = model
        self.role = "assistant"
        self.client = Client()

    @retry(stop=stop_after_attempt(3), wait=wait_exponential(multiplier=1, min=4, max=10))
    def generate(self, request):
        response = self.client.chat.completions.create(
            model=self.model,
            messages=[{"role": self.role, "content": request}],
        )

        return response.choices[0].message.content