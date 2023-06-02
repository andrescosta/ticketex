from fastapi import APIRouter
from app.api.v1 import partners

v1_api_router = APIRouter()
v1_api_router.include_router(partners.router, prefix="/v1/partners", tags=["partners"])

