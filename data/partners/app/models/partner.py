from typing import List, Optional
from odmantic import Model
from app.models.address import Address


class Partner(Model):
    email: Optional[str]
    name: Optional[str]
    phone: Optional[str]
    ext_id: Optional[str]
    addresses: Optional[List[Address]]