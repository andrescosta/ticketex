from fastapi import APIRouter

from app.api.v1 import adventures
from tiklib.api.handlers import entity_not_found_error_exception_handler
from tiklib.repository.entity_not_found_error import EntityNotFoundError

v1_api_router = APIRouter()
v1_api_router.include_router(adventures.router, prefix="/v1/adventures", tags=["adventures"])

v1_exception_handlers = [
    (EntityNotFoundError, entity_not_found_error_exception_handler)
]
