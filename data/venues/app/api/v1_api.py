from fastapi import APIRouter
from app.api.v1 import venues

v1_api_router = APIRouter()
v1_api_router.include_router(venues.router, prefix="/v1/venues", tags=["venues"])

