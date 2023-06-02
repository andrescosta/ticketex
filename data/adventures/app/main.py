from fastapi import FastAPI
from app.api.v1_api import v1_api_router
from tiklib.init import init_tiklib
from tiklib.log.logging import add_request_timing_log
import logging


app = FastAPI(debug=True)
app.include_router(v1_api_router)

init_tiklib(app)

logger = logging.getLogger(__name__)
add_request_timing_log(app, logger)