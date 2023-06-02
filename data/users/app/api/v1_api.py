from fastapi import APIRouter

from app.api.v1 import users

v1_api_router = APIRouter()
v1_api_router.include_router(users.router, prefix="/v1/users", tags=["users"])