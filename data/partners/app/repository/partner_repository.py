from app.models.partner import Partner
from tiklib.repository.entity_repository import EntityRepository
from motor.motor_asyncio import AsyncIOMotorClient

class PartnerRepository(EntityRepository):
    client = AsyncIOMotorClient("mongodb://root:example@localhost:27017")
    db = client["ticketex"]
    mycollection = db["partners"]

    def __init__(self) -> None:
        super().__init__(type(Partner()))

    @property
    def collection(self):
        return self.mycollection

partner = PartnerRepository()
