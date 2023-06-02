from fastapi import  Request
from fastapi.responses import JSONResponse
from  pymongo.errors import PyMongoError
from fastapi import FastAPI
from tiklib.repository.entity_not_found_error import EntityNotFoundError
import logging

logger = logging.getLogger(__name__)

async def entity_not_found_error_exception_handler(
    request: Request, exc: EntityNotFoundError
) -> JSONResponse:
    return JSONResponse(
        status_code=404,
        content={"message": str(exc)},
    )

async def mongo_error_exception_handler(
    request: Request, exc: PyMongoError
) -> JSONResponse:
    logger.exception(f"Error connecting mongo: {str(exc)}")
    return JSONResponse(
        status_code=500,
        content={"message": "Internal error."},
    )

v1_exception_handlers = [
    (EntityNotFoundError, entity_not_found_error_exception_handler),
    (PyMongoError, mongo_error_exception_handler)
]

def add_handlers(app:FastAPI):
    for handler in v1_exception_handlers:
        app.add_exception_handler(handler[0], handler[1])