from typing import Optional
from pydantic import BaseModel


class Address(BaseModel):
    zipcode: Optional[str]
    street1:Optional[str]
    street2:Optional[str]
    country:Optional[str]
    state:Optional[str]
