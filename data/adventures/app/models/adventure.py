from typing import List, Optional
from app.models.address import Address
from tiklib.models.entity import Entity

from app.models.ticket import Ticket
from app.models.venue import Venue


class Adventure(Entity):
    name: Optional[str]
    description: Optional[str]
    tickets: Optional[List[Ticket]]
    venue: Optional[Venue]