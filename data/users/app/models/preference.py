from typing import Optional, List

from pydantic import BaseModel

from app.models.channel import Channel

class Preference(BaseModel):
    channels: Optional[List[Channel]]
