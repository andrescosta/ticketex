from fastapi import Depends, FastAPI
from motor.motor_asyncio import AsyncIOMotorClient
from pymongo.database import Database
from tiklib.common import app_settings
import logging

logger = logging.getLogger(__name__)

__client:AsyncIOMotorClient = None

def init_motor_client(app:FastAPI):
    app.add_event_handler("startup", connect_client)
    app.add_event_handler("shutdown", close_client)

async def connect_client():
    global __client
    try:
        __client = AsyncIOMotorClient(app_settings.DATABASE_URL)
        logger.debug("Mongo connection was initialized.")
    except Exception as e:
        logger.error("Error initializing Mongo connection.")
        raise

def get_db()->Database:
    return __client[app_settings.DATABASE_NAME]

async def close_client():
    global __client
    if __client is None:
        logger.warn("Connection is None, nothing to close.")
        return
    __client.close()
    __client = None
    logger.debug("Mongo connection closed.")

