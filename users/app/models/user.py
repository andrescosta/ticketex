from typing import Optional
from app.models.entity import Entity


class User(Entity):
    email: Optional[str]
    fullname: Optional[str]
    phone: Optional[str]

    class Config:
        schema_extra = {
            "example": {
                "id": "aaaa",
                "email": "user@example.com",
                "fullname": "John Doe",
                "phone": "1234567890",
                "jwt_sub": "abc123xyz"
            }
        }
