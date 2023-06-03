from enum import Enum


class ChannelType(str, Enum):
    EMAIL = "email"
    PHONE = "phone"

class MessageType(str, Enum):
    TICKETS="tickets"