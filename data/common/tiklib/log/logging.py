import logging
import yaml
import coloredlogs
import random
import string
import time
from fastapi import FastAPI, Request


def init_log(logfile:str="logging.yaml"):
    try:
        with open(logfile,"rt") as stream:
            config = yaml.safe_load(stream)
            logging.config.dictConfig(config)
        coloredlogs.install()
    except Exception as e:
        print(f'Error reading {logfile}: "{e}". Setting log level as info')
        logging.basicConfig(level=logging.INFO)
        coloredlogs.install(level=logging.INFO)

def add_request_timing_log(app:FastAPI, logger:logging.Logger):
    @app.middleware("http")
    async def log_requests(request: Request, call_next):
        idem = ''.join(random.choices(string.ascii_uppercase + string.digits, k=6))
        logger.debug(f"rid={idem} start request path={request.url.path}")
        start_time = time.time()
    
        response = await call_next(request)
    
        process_time = (time.time() - start_time) * 1000
        formatted_process_time = '{0:.2f}'.format(process_time)
        logger.debug(f"rid={idem} completed_in={formatted_process_time}ms status_code={response.status_code}")
    
        return response
