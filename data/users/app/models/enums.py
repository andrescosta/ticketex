from enum import Enum


class ChannelType(str, Enum):
    EMAIL = "email"
    PHONE = "phone"

class MessageType(str, Enum):
    TICKETS="tickets"

class UserType(str, Enum):
    PARTNER = "partner"
    END_USER = "end_user"