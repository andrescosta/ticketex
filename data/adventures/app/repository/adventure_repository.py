from tiklib.repository.entity_repository import EntityRepository
from app.models.adventure import Adventure

class AdventureRepository(EntityRepository):
    def __init__(self) -> None:
        super().__init__(type(Adventure()))

adventure = AdventureRepository()