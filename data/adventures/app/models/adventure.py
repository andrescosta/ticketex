from typing import List, Optional
from app.models.address import Address
from tiklib.models.partner_entity import PartnerEntity

from app.models.ticket import Ticket

class Adventure(PartnerEntity):
    name: Optional[str]
    description: Optional[str]
    tickets: Optional[List[Ticket]]
    venue_id: Optional[str]