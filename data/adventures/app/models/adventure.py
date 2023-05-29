from typing import List, Optional
from app.models.address import Address
from tiklib.models.entity import Entity


class Adventure(Entity):
    email: Optional[str]
    fullname: Optional[str]
    phone: Optional[str]
    ext_id: Optional[str]
    address: Optional[Address]