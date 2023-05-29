from typing import Optional

from pydantic import BaseModel

from app.models.address import Address


class Venue (BaseModel):
    name: Optional[str]
    address: Optional[Address]
    phone: Optional[str]