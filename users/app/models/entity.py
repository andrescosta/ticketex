from typing import Optional
from pydantic import BaseModel

class Entity(BaseModel):
    id: Optional[str]
    jwt_sub: Optional[str]


