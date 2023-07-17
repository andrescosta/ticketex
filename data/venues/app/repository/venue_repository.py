from app.models.venue import Venue
from tiklib.repository.entity_repository import EntityRepository

class VenueRepository(EntityRepository):
    def __init__(self) -> None:
        super().__init__(type(Venue()))

venue = VenueRepository()