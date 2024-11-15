from fastapi import FastAPI, HTTPException
import os
from dotenv import load_dotenv

from models.question_request import QuestionRequest
from providers.openai_client import G4FClient

load_dotenv()

app = FastAPI()

g4f_client = G4FClient(os.getenv("GPT_MODEL"))

@app.post("/api/v1/ask", status_code=200)
async def gpt(request: QuestionRequest):
    if not request:
        raise HTTPException(status_code=400, detail="Empty Request")

    response = g4f_client.generate(request.question)

    if not response:
        raise HTTPException(status_code=404, detail="No response received")

    return {"message": response}