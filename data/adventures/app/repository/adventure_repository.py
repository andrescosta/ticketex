from tiklib.repository.entity_repository import EntityRepository
from motor.motor_asyncio import AsyncIOMotorClient
from app.models.adventure import Adventure

class AdventureRepository(EntityRepository):
    client = AsyncIOMotorClient("mongodb://root:example@localhost:27017")
    db = client["ticketex"]
    mycollection = db["adventures"]

    def __init__(self) -> None:
        super().__init__(type(Adventure()))

    @property
    def collection(self):
        return self.mycollection


adventure = AdventureRepository()