from fastapi import FastAPI
from app.api.v1_api import v1_api_router, v1_exception_handlers
from tiklib.db.client import init_motor_client

app = FastAPI(debug=True)
app.include_router(v1_api_router)
for handler in v1_exception_handlers:
    app.add_exception_handler(handler[0], handler[1])

init_motor_client(app)