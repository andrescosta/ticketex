from typing import Optional

from pydantic import BaseModel

from tiklib.models.address import Address
from tiklib.models.partner_entity import PartnerEntity

class Venue (PartnerEntity):
    name: Optional[str]
    address: Optional[Address]
    phone: Optional[str]