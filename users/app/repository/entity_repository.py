from typing import Any
from motor.motor_asyncio import AsyncIOMotorClient
import copy
from bson import ObjectId, json_util


from app.models.entity import Entity
from app.repository.entity_not_found_error import EntityNotFoundError


class EntityRepository:
    
    client = AsyncIOMotorClient("mongodb://root:example@localhost:27017")
    db = client["tick_user"]
    collection = db["users"]
    entitytype=None

    def __init__(self, entitytype) -> None:
        self.entitytype=entitytype

    async def save(self, entity)->Entity:
        result = await self.collection.insert_one(entity.dict())
        nentity = copy.copy(entity)
        nentity.id = str(result.inserted_id); 
        return nentity
    
    async def get(self, id)->Entity:
        entity_dic = await self.collection.find_one({"_id": ObjectId(id)})
        if entity_dic:
            entity_dic["id"] = str(entity_dic.pop("_id"))
            entity = json_util.dumps(entity_dic)
            return entity
        else:
            raise EntityNotFoundError(id, self.entitytype)

    async def update(self, entity)->None:
        entity_dict = entity.dict(exclude_unset=True)
        result = await self.collection.update_one({"_id": ObjectId(entity.id)}, {"$set": entity_dict})
        if result.modified_count == 0:
            raise EntityNotFoundError(entity.id, self.entitytype)


    async def delete(self, id)->None:
        result = await self.collection.delete_one({"_id": ObjectId(id)})
        if result.deleted_count == 0:
            raise EntityNotFoundError(id, self.entitytype)
