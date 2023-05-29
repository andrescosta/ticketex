from fastapi import  Request
from fastapi.responses import JSONResponse

from tiklib.repository.entity_not_found_error import EntityNotFoundError

async def entity_not_found_error_exception_handler(
    request: Request, exc: EntityNotFoundError
) -> JSONResponse:
    return JSONResponse(
        status_code=404,
        content={"message": str(exc)},
    )
