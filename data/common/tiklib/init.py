from fastapi import FastAPI
from tiklib.api.handlers import add_handlers

from tiklib.db.client import init_motor_client
from tiklib.log.logging import init_log


def init_tiklib(app:FastAPI):
    init_log()
    init_motor_client(app)
    add_handlers(app)
    
