from typing import List, Optional
from app.models.address import Address
from tiklib.models.entity import Entity


class Partner(Entity):
    email: Optional[str]
    name: Optional[str]
    phone: Optional[str]
    ext_id: Optional[str]
    addresses: Optional[List[Address]]