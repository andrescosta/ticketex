from typing import Any
from motor.motor_asyncio import AsyncIOMotorClient
from odmantic import AIOEngine, Model


from tiklib.repository.entity_not_found_error import EntityNotFoundError


class EntityRepository:
    
    client = AsyncIOMotorClient("mongodb://root:example@localhost:27017/")
    engine = AIOEngine(client=client, database="ticketex")
    
    def __init__(self, entitytype) -> None:
        self.entitytype=entitytype

    async def save(self, entity:Model)->Model:
        return await self.engine.save(entity)
    
    async def get(self, entity:Model)->Model:
        entity = await self.engine.find_one(self.entitytype, self.entitytype.id == entity.id)
        if entity:
            return entity
        else:
            raise EntityNotFoundError(id, self.entitytype)

    async def update(self, entity:Model)->Model:
        return await self.save(entity)


    async def delete(self, entity:Model)->None:
        deleted_count = await self.engine.delete(entity)
        if deleted_count == 0:
            raise EntityNotFoundError(id, self.entitytype)