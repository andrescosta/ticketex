from app.models.user import User
from tiklib.repository.entity_repository import EntityRepository
from motor.motor_asyncio import AsyncIOMotorClient

class UserRepository(EntityRepository):
    client = AsyncIOMotorClient("mongodb://root:example@localhost:27017")
    db = client["ticketex"]
    mycollection = db["users"]

    def __init__(self) -> None:
        super().__init__(type(User()))

    @property
    def collection(self):
        return self.mycollection


user = UserRepository()