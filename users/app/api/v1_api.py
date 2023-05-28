from fastapi import APIRouter

from app.api.v1 import users
from app.api.v1.handlers import entity_not_found_error_exception_handler
from app.repository.entity_not_found_error import EntityNotFoundError

v1_api_router = APIRouter()
v1_api_router.include_router(users.router, prefix="/v1/users", tags=["users"])

v1_exception_handlers = [
    (EntityNotFoundError, entity_not_found_error_exception_handler)
]
