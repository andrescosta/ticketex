from typing import List, Optional
from app.models.preference import Preference
from app.models.entity import Entity


class User(Entity):
    email: Optional[str]
    fullname: Optional[str]
    phone: Optional[str]
    ext_id: Optional[str]
    preferences: Optional[List[Preference]]

"""   class Config:
        schema_extra = {
            "example": {
                "id": "aaaa",
                "email": "user@example.com",
                "fullname": "John Doe",
                "phone": "1234567890",
                "jwt_sub": "abc123xyz",
                "preferences": [
                    {
                        {"name": "rock", "notif": True},
                        {"name": "pop", "notif": False},
                    },
                ],
            }
        }
"""