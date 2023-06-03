from typing import Optional
from pydantic import BaseModel
from app.models.enums import ChannelType,MessageType


class Channel(BaseModel):
    value:Optional[str]
    channel_type:ChannelType
    message_type:MessageType

