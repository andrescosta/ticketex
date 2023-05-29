from typing import Optional

from pydantic import BaseModel


class Ticket(BaseModel):
    capacity: Optional[int]
    type: Optional[str]