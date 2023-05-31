from app.models.user import User
from tiklib.repository.entity_repository import EntityRepository

class UserRepository(EntityRepository):
    def __init__(self) -> None:
        super().__init__(type(User()))

user = UserRepository()