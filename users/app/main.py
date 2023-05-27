from fastapi import FastAPI
from app.api.v1_api import api_router

app = FastAPI(debug=True)
app.include_router(api_router)

