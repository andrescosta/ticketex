from fastapi import APIRouter

from app.api.v1 import partners
from tiklib.api.handlers import entity_not_found_error_exception_handler
from tiklib.repository.entity_not_found_error import EntityNotFoundError

v1_api_router = APIRouter()
v1_api_router.include_router(partners.router, prefix="/v1/partners", tags=["partners"])

v1_exception_handlers = [
    (EntityNotFoundError, entity_not_found_error_exception_handler)
]
