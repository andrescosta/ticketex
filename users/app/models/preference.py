from typing import Optional

from pydantic import BaseModel


class Preference(BaseModel):
    name: Optional[str]
    notif: Optional[bool]