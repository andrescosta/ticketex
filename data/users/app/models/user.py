from typing import List, Optional
from app.models.address import Address
from app.models.preference import Preference
from tiklib.models.entity import Entity


class User(Entity):
    email: Optional[str]
    fullname: Optional[str]
    phone: Optional[str]
    ext_id: Optional[str]
    preferences: Optional[List[Preference]]
    address: Optional[Address]