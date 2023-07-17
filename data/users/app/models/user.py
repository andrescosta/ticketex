from typing import List, Optional
from tiklib.models.address import Address
from app.models.preference import Preference
from tiklib.models.entity import Entity
from app.models.enums import UserType

class User(Entity):
    email: Optional[str]
    fullname: Optional[str]
    phone: Optional[str]
    ext_id: Optional[str]
    preferences: Optional[List[Preference]]
    address: Optional[Address]
    type: Optional[UserType]