from app.models.partner import Partner
from tiklib.repository.entity_repository import EntityRepository

class PartnerRepository(EntityRepository):
    def __init__(self) -> None:
        super().__init__(type(Partner()))

partner = PartnerRepository()